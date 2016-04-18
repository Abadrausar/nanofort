/*
Copyright 2013-2016 by Milo Christiansen

This software is provided 'as-is', without any express or implied warranty. In
no event will the authors be held liable for any damages arising from the use of
this software.

Permission is granted to anyone to use this software for any purpose, including
commercial applications, and to alter it and redistribute it freely, subject to
the following restrictions:

1. The origin of this software must not be misrepresented; you must not claim
that you wrote the original software. If you use this software in a product, an
acknowledgment in the product documentation would be appreciated but is not
required.

2. Altered source versions must be plainly marked as such, and must not be
misrepresented as being the original software.

3. This notice may not be removed or altered from any source distribution.
*/

// Rubble CLI Interface
package main

import "os"
import "fmt"
import "io/ioutil"
import "strconv"
import "bytes"
import "strings"
import "runtime"
import "sort"

import "flag"

import "rubble7"
import "rubble7/rblutil"
import "rubble7/rblutil/test"
import "rubble7/rblutil/addon"
import "rubble7/rblutil/parse"
import "rubble7/rblutil/errors"

import _ "rubble7/scripting/dctech_lua"
import _ "rubble7/scripting/otto"

var RblDir string
var DFDir string
var OutputDir string
var AddonsDir *rblutil.ArgList

var ConfigList *rblutil.ArgList

var ZapConfig bool

var TestID string
var AddonName string

var Verbose bool

func main() {
	err, log := rblutil.NewLogger()
	if err != nil {
		fmt.Println("Fatal Error:", err)
		os.Exit(1)
	}

	flags := flag.NewFlagSet("rubble", flag.ExitOnError)
	flags.SetOutput(log)

	log.Header(rubble7.Version)

	defer func() {
		err := recover()
		if err != nil {
			log.Println("Unrecovered Error:")
			log.Println("  The following error was not properly recovered, please report this ASAP!")
			log.Printf("  %#v\n", err)
			log.Println("Stack Trace:")
			buf := make([]byte, 4096)
			buf = buf[:runtime.Stack(buf, true)]
			log.Printf("%s\n", buf)
			os.Exit(1)
		}
	}()

	// Used for the -zapxxx options.
	configCount := 0

	// Setting hardcoded defaults
	RblDir = "."
	DFDir = ".."
	OutputDir = "df:raw"
	AddonsDir = new(rblutil.ArgList)

	ConfigList = new(rblutil.ArgList)

	// Load defaults from config if present
	log.Println("Attempting to Read Config File: ./rubble.ini")
	file, err := ioutil.ReadFile("./rubble.ini")
	if err == nil {
		log.Println("  Read OK, loading options from file.")
		rblutil.ParseINI(string(file), "\n", func(key, value string) {
			switch key {
			case "rbldir":
				RblDir = value
			case "dfdir":
				DFDir = value
			case "outputdir":
				OutputDir = value
			case "addonsdir":
				AddonsDir.Set(value)
			case "config":
				configCount++
				ConfigList.Set(value)
			case "zapconfig":
				ZapConfig, _ = strconv.ParseBool(value)
			case "testaddon":
				AddonName = value
			case "testid":
				TestID = value
			case "verbose":
				Verbose, _ = strconv.ParseBool(value)
			default:
				log.Println("  Invalid config option:", key, ", skipping.")
			}
		})
	} else {
		log.Println("  Read failed (this is most likely ok)\n  Error:", err)
		log.Println("    Using hardcoded defaults.")
	}

	flags.StringVar(&RblDir, "rbldir", RblDir, "The path to Rubble's directory.")
	flags.StringVar(&DFDir, "dfdir", DFDir, "The path to the DF directory. May be an AXIS path (only the 'rubble' location ID works).")
	flags.StringVar(&OutputDir, "outputdir", OutputDir, "Where should Rubble write the generated raw files? May be an AXIS path (only the 'rubble' and 'df' location IDs work).")
	flags.Var(AddonsDir, "addonsdir", "Rubble addons directory. May be an AXIS path (only the 'rubble', 'df', and 'out' location IDs work).")

	flags.Var(ConfigList, "config", "List of config variables. This is optional. If the value is a file path then the file is read as an ini file containing config variables. May be specified more than once.")

	flags.BoolVar(&ZapConfig, "zapconfig", ZapConfig, "Ignore any -config flags loaded from rubble7.ini.")
	
	flags.StringVar(&AddonName, "testaddon", AddonName, "Only run tests that are in a certain addon.")
	flags.StringVar(&TestID, "testid", TestID, "Only run tests that have the given ID (more than one test can have the same ID).")
	flags.BoolVar(&Verbose, "verbose", Verbose, "Print verbose output.")

	flags.Parse(os.Args[1:])

	if AddonsDir.Empty() {
		AddonsDir.Set("rubble:addons")
	}

	if ZapConfig {
		if configCount > 0 && configCount < len(*ConfigList) {
			*ConfigList = (*ConfigList)[configCount:]
		}
	}

	log.Separator()
	log.Println("Getting List of Addons with Tests...")
	
	err, state := rubble7.NewState(RblDir, DFDir, OutputDir, *AddonsDir, log)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = state.Load(nil, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	addons := []string{}
	for _, addon := range state.Addons { // DO NOT USE state.AddonMeta! This needs to produce a list in load order.
		if addon.Meta.Tags["HasTests"] {
			addons = append(addons, addon.Meta.Name)
		}
	}
	
	testRunCount := 0
	testFailCount := 0
	resultBuffer := new(bytes.Buffer)
	
	state.Log.Separator()
	log.Println("Running Specified Tests...")
	log.Println("  This will produce a *lot* of log file spam for every addon containing tests!")
	log.Println("  Results will be summarized at the end.")
	
	for _, name := range addons {
		if AddonName != "" && AddonName != name {
			continue
		}
		
		err, state := rubble7.NewState(RblDir, DFDir, OutputDir, *AddonsDir, log)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		
		runScript := func(file *addon.File) error {
			state.CurrentFile = file.Name
			state.Log.Println("  " + file.Name)
			
			_, err := state.Script.RunScriptFile(file)
			if err != nil {
				return err
			}
			return nil
		}
		
		err = state.Load(nil, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		
		err = state.ActivateAny([]string{name}, *ConfigList)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		
		state.Log.Println("Generating Sorted Active File List...")
		state.Files.Update(state.Addons)
		
		state.Log.Separator()
		state.Log.Println("Running Init Scripts...")
		err = state.GlobalFiles.RunAction(map[string]bool{
			"Skip":       false,
			"InitScript": true,
		}, runScript)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	
		state.Log.Separator()
		state.Log.Println("Running Prescripts...")
		err = state.Files.RunAction(map[string]bool{
			"Skip":      false,
			"PreScript": true,
		}, runScript)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		
		state.Log.Println("Running Test Files...")
		order := []string{}
		for fname := range state.AddonsTbl[name].Files {
			order = append(order, fname)
		}
		sort.Strings(order)
		for _, fname := range order {
			file := state.AddonsTbl[name].Files[fname]
			
			if file.Tags["Skip"] || !file.Tags["TemplateTest"] {
				continue
			}
			
			state.CurrentFile = file.Name
			state.Log.Println("  " + file.Name)
			
			tests := test.Parse(file.Content, file.Name, 1)
			for _, test := range tests {
				if TestID != "" && TestID != test.ID {
					continue
				}
				testRunCount++
				
				err := func() (err error) {
					defer errors.TrapError(&err, state.Log)
					
					state.Stage = parse.StgPreParse
					result := parse.Parse([]byte(test.In), file.Name, test.InLine, parse.StgPreParse, state.Dispatcher)
					state.Stage = parse.StgParse
					result = parse.Parse(result, file.Name, test.InLine, parse.StgParse, state.Dispatcher)
					state.Stage = parse.StgPostParse
					result = parse.Parse(result, file.Name, test.InLine, parse.StgPostParse, state.Dispatcher)
					
					final := strings.TrimSpace(string(result))
					
					if final != test.Out {
						testFailCount++
						msg := test.FailMsg(name, final)
						fmt.Fprintln(resultBuffer, msg)
						state.Log.Println(msg)
					} else if Verbose {
						msg := test.PassMsg(name, final)
						fmt.Fprintln(resultBuffer, msg)
						state.Log.Println(msg)
					}
					return nil
				}()
				if err != nil {
					testFailCount++
					msg := "    Unnamed Test in Addon: \"" + name + "\" Failed:\n"
					if test.ID != "" {
						msg = "    Test: " + test.ID + " in Addon: \"" + name + "\" Failed:\n"
					}
					msg += "      " + err.Error()
					fmt.Fprintln(resultBuffer, msg)
					state.Log.Println(msg)
				}
			}
		}
	}
	
	state.Log.Separator()
	state.Log.Printf("%v Tests Run.\n", testRunCount)
	if testFailCount > 0 {
		state.Log.Printf("  %v Tests Failed.\n", testFailCount)
		state.Log.Println("  Result Summary:")
		state.Log.Print(resultBuffer.String())
	} else {
		state.Log.Println("  All Tests Passed!")
		if Verbose {
			state.Log.Println("  Result Summary:")
			state.Log.Print(resultBuffer.String())
		}
	}
	log.Println("Done.")
}

