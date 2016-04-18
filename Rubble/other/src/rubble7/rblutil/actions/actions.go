/*
Copyright 2016 by Milo Christiansen

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

// The actions registry for the universal interface.
package actions

import "rubble7/rblutil"

import "flag"
import "os"
import "strconv"

// Option is a command line option descriptor.
type Option struct {
	Name string
	Help string
	
	DS string // Default for string flags
	DB bool // Default for bool flags.
	// []string flags have no defaults.
	
	// If true the result is a *bool, otherwise the result is a *string (or a *rblutil.ArgList (*[]string) if Multiple is true).
	Flag     bool
	Multiple bool // Ignored if Flag is true.
}

type ActionFunc func(log *rblutil.Logger, rblDir, dfDir, outDir string, addonDirs []string, options []interface{}) bool

// This API is very specific to the universal interface, so it's OK to use globals for this kind of stuff.
var actions = map[string]Action{}
var optDefs = map[string][]Option{}
var modeHelp = map[string]string{}
var modes string

func commonOpt(f *flag.FlagSet, rblini map[string][]string, name, value, help string) interface{} {
	iniv, ok := rblini[name]
	if !ok || len(iniv) == 0 {
		return f.String(name, value, help)
	}
	return f.String(name, iniv[len(iniv)-1], help)
}

func makeFlags(mode string, rblini map[string][]string, log *rblutil.Logger, hidecommon bool) (*flag.FlagSet, []interface{}, []interface{}) {
	def, ok := optDefs[mode]
	if !ok {
		return nil, nil, nil
	}
	
	f := flag.NewFlagSet(mode, flag.ExitOnError)
	options := make([]interface{}, len(def))
	f.SetOutput(log)
	
	for i, opt := range def {
		iniv, ok := rblini[opt.Name]
		if ok && len(iniv) == 0 {
			ok = false
		}
		if opt.Flag {
			d := opt.DB
			if ok {
				d, _ = strconv.ParseBool(iniv[len(iniv)-1])
			}
			options[i] = f.Bool(opt.Name, d, opt.Help)
			continue
		}
		if opt.Multiple {
			v := new(rblutil.ArgList)
			f.Var(v, opt.Name, opt.Help)
			if ok {
				for i := range iniv {
					v.Set(iniv[i])
				}
			}
			options[i] = v
			continue
		}
		d := opt.DS
		if ok {
			d = iniv[len(iniv)-1]
		}
		options[i] = f.String(opt.Name, d, opt.Help)
	}
	
	common := make([]interface{}, 0, 4)
	if !hidecommon {
		// Add the common options
		common = append(common, commonOpt(f, rblini, "rbldir", ".", "If you are seeing this it's an error!"))
		common = append(common, commonOpt(f, rblini, "dfdir", "..", "If you are seeing this it's an error!"))
		common = append(common, commonOpt(f, rblini, "outputdir", "df:raw", "If you are seeing this it's an error!"))
		v := new(rblutil.ArgList)
		common = append(common, v)
		f.Var(v, "addonsdir", "If you are seeing this it's an error!")
		iniv, ok := rblini["addonsdir"]
		if ok {
			for i := range iniv {
				v.Set(iniv[i])
			}
		}
	}
	
	f.Usage = func() {
		log.Println("Run \"rubble help\" for usage.")
	}
	
	return f, options, common
}

// Register an action with this API.
func Register(mode string, action Action, help string, options []Option) {
	actions[mode] = action
	modeHelp[mode] = help
	optDefs[mode] = options
	if modes != "" {
		modes += ", " + mode
		return
	}
	modes = mode
}

// Register an action function with this API. The passed in value is automatically wrapped.
func RegisterFunc(mode string, action ActionFunc, help string, options []Option) {
	Register(mode, funcWrapper(action), help, options)
}

// This includes usage for the common arguments. I do some real backflips to keep these out of the usage
// statements for individual modes so as not to clutter them.
var usage = `Usage:

  rubble mode [arguments]

The available modes are:

  %v

Run "rubble help [mode]" for more information about a mode.

Some shared options are available in all modes. These options are as follows:

  -addonsdir value
        Rubble addons directory. May be an AXIS path (only the 'rubble', 'df',
        and 'out' location IDs work). May be specified multiple times. If no
        values are specified this defaults to "rubble:addons".
  -dfdir string
        The path to the DF directory. May be an AXIS path (only the 'rubble'
        location ID works). (default "..")
  -outputdir string
        Where should Rubble write the generated raw files? May be an AXIS path
        (only the 'rubble' and 'df' location IDs work). (default "df:raw")
  -rbldir string
        The path to Rubble's directory. (default ".")
`

// Exec runs the action associated with the given mode. If the mode has no action then an error message will be
// printed to the log (listing the valid modes).
// This function will never return!
func Exec(mode string, args []string, log *rblutil.Logger, rblini map[string][]string) {
	if mode == "help" {
		log.Separator()
		if len(args) == 0 {
			log.Printf(usage, modes)
			os.Exit(0)
		}
		help, ok := modeHelp[args[0]]
		if !ok {
			log.Println("Unknown mode \"" + args[0] + "\" passed to help.\nValid modes are: " + modes)
			os.Exit(0)
		}
		log.Println("Help for mode:", args[0])
		log.Println()
		log.Println(help)
		f, _, _ := makeFlags(args[0], rblini, log, true)
		// Horrid hack
		fc := 0
		f.VisitAll(func(fl *flag.Flag) {
			fc++
		})
		if fc != 0 {
			log.Println("\nOptions:\n")
			f.PrintDefaults()
		}
		log.Println("\nFor options common to all modes run \"rubble help\"")
		os.Exit(0)
	}
	
	f, options, common := makeFlags(mode, rblini, log, false)
	if f == nil {
		log.Println("Unknown mode, valid modes are: help, " + modes)
		os.Exit(3)
	}
	
	f.Parse(args) // os.Exit(2) on error
	
	rblDir := *common[0].(*string)
	dfDir := *common[1].(*string)
	outDir := *common[2].(*string)
	addonDirs := common[3].(*rblutil.ArgList)
	if addonDirs.Empty() {
		addonDirs.Set("rubble:addons")
	}
	
	action := actions[mode]
	ok := action.Run(log, rblDir, dfDir, outDir, *addonDirs, options)
	if !ok {
		// Operation failed!
		// The action should have already logged an error message.
		os.Exit(1)
	}
	os.Exit(0)
}

// Wrapper type for action functions.
type funcWrapper ActionFunc

func (f funcWrapper) Run(log *rblutil.Logger, rblDir, dfDir, outDir string, addonDirs []string, options []interface{}) bool {
	return f(log, rblDir, dfDir, outDir, addonDirs, options)
}

// Action is the interface types providing actions to the universal interface must satisfy.
type Action interface {
	// Carry out the action.
	// The only tricky bit is `options`, this is a slice of a mix of `*bool`, `*string`, and `*rblutil.ArgList` values.
	// Which ones are where is determined by the option specifiers you pass to Register.
	// Most of the other arguments are set from the "common options" all actions support.
	Run(log *rblutil.Logger, rblDir, dfDir, outDir string, addonDirs []string, options []interface{}) bool
}

