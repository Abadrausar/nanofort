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

// Rubble Web GUI Interface
package main

import "net/http"

import "os"
import "fmt"
import "bytes"
import "runtime/pprof"
import "time"
import "io/ioutil"
import "strings"
import "strconv"
import "runtime"

import "flag"

import "image"
import "image/png"

import "rubble7"
import "rubble7/rblutil"
import "rubble7/rblutil/rparse"
import "rubble7/rblutil/brender"
import "rubble7/rblutil/addon"

import _ "rubble7/scripting/dctech_lua"
import _ "rubble7/scripting/otto"

import "dctech/axis"

var Addr string

var RblDir string
var DFDir string
var OutputDir string
var AddonsDir *rblutil.ArgList

var AddonsList *rblutil.ArgList
var ConfigList *rblutil.ArgList

var Profile string

func main() {
	logBuffer := new(bytes.Buffer)
	err, log := rblutil.NewLogger(logBuffer)
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

	Addr = "127.0.0.1:2120"

	RblDir = "."
	DFDir = ".."
	OutputDir = "df:raw"
	AddonsDir = new(rblutil.ArgList)

	// Load defaults from config if present
	log.Println("Attempting to Read Config File: ./rubble7.ini")
	file, err := ioutil.ReadFile("./rubble7.ini")
	if err == nil {
		log.Println("  Read OK, loading options from file.")
		rblutil.ParseINI(string(file), "\n", func(key, value string) {
			switch key {
			case "addr":
				Addr = value
			case "rbldir":
				RblDir = value
			case "dfdir":
				DFDir = value
			case "outputdir":
				OutputDir = value
			case "addonsdir":
				AddonsDir.Set(value)
			case "profile":
				Profile = value
			default:
				log.Println("  Invalid config option:", key, ", skipping.")
			}
		})
	} else {
		log.Println("  Read failed (this is most likely ok)\n  Error:", err)
		log.Println("    Using hardcoded defaults.")
	}

	flags.StringVar(&Addr, "addr", Addr, "The server address.")

	flags.StringVar(&RblDir, "rbldir", RblDir, "The path to Rubble's directory.")
	flags.StringVar(&DFDir, "dfdir", DFDir, "The path to the DF directory. May be an AXIS path (only the 'rubble' location ID works).")
	flags.StringVar(&OutputDir, "outputdir", OutputDir, "Where should Rubble write the generated raw files? May be an AXIS path (only the 'rubble' and 'df' location IDs work).")
	flags.Var(AddonsDir, "addonsdir", "Rubble addons directory. May be an AXIS path (only the 'rubble', 'df', and 'out' location IDs work).")

	flags.StringVar(&Profile, "profile", "", "Output CPU profile information to specified file.")

	flags.Parse(os.Args[1:])

	if AddonsDir.Empty() {
		AddonsDir.Set("rubble:addons")
	}

	if Profile != "" {
		log.Println("Writing CPU profiling data to file:", Profile)
		f, err := os.Create(Profile)
		if err != nil {
			log.Println("  Profile Error:", err)
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	err, state := rubble7.NewState(RblDir, DFDir, OutputDir, *AddonsDir, log)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	tmpl, err := LoadHTMLTemplates(state.FS)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	css := getCSS(state.FS)

	// Note: http.Redirect acts "funny" when you use http.StatusInternalServerError,
	// so I use http.StatusFound even for error-triggered redirects to the log page.
	// (this may just be my browser, but better safe than sorry)

	// Main menu
	http.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/menu\"")

		timeStart := time.Now()
		var err error
		err, state = rubble7.NewState(RblDir, DFDir, OutputDir, *AddonsDir, log)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}

		err = state.Load([]string{}, []string{})
		log.Println("Load time: ", time.Since(timeStart))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}

		err = tmpl.ExecuteTemplate(w, "menu", state)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Redirecting:", r.URL, "to main menu.")

		http.Redirect(w, r, "./menu", http.StatusFound)
	})

	// Addon List

	http.HandleFunc("/addons", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/addons\"")

		err = tmpl.ExecuteTemplate(w, "addons", state)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
	})

	// Normal Generation
	
	http.HandleFunc("/genaddons", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/genaddons\"")

		err = tmpl.ExecuteTemplate(w, "genaddons", state)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
	})

	http.HandleFunc("/genvars", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/genvars\"")

		addons := []string{}
		for _, addon := range state.Addons {
			if addon.Active {
				addons = append(addons, addon.Meta.Name)
			}
		}

		// This makes sure all addons are in the correct state, even though the non-library addons are more or
		// less correct there is a lot more to activation than just that.
		err = state.Activate(addons, nil)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}

		// If the user elected to skip selecting config vars, go straight to generation after activation.
		if r.FormValue("hidden") != "" {
			http.Redirect(w, r, "/pleasewait?to=/genrun", http.StatusFound)
			return
		}

		vars := make(map[string]*addon.MetaVar)
		vals := make(map[string]string)
		for _, addon := range state.Addons {
			if !addon.Active {
				continue
			}

			for i := range addon.Meta.Vars {
				vars[i] = addon.Meta.Vars[i]
				data, ok := state.VariableData[i]
				if ok {
					vals[i] = data
				} else if len(vars[i].Values[0]) > 0 {
					vals[i] = vars[i].Values[0]
				}
			}
		}

		err = tmpl.ExecuteTemplate(w, "genvars", struct {
			Vars map[string]*addon.MetaVar
			Vals map[string]string
		}{vars, vals})
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
	})

	http.HandleFunc("/genrun", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/genrun\"")

		timeStart := time.Now()
		err := state.RunActivated()
		log.Println("Run time: ", time.Since(timeStart))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
		log.Println("Done.")
		
		if log.WC > 0 {
			http.Redirect(w, r, "/log?state=warn", http.StatusFound)
			return
		}
		
		http.Redirect(w, r, "/log?header=true", http.StatusFound)
	})

	// Tileset

	http.HandleFunc("/tsetaddons", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/tsetaddons\"")

		regions := axis.ListDir(state.FS, "df:data/save")
		for i := range regions {
			if regions[i] == "current" {
				regions = append(regions[:i], regions[i+1:]...)
				break
			}
		}

		for _, addon := range state.Addons {
			if !addon.Meta.Tags["TileSet"] {
				addon.Active = false
			}
		}

		state.VariableData["_RUBBLE_IA_REGION_"] = ""

		err = tmpl.ExecuteTemplate(w, "iaaddons", struct {
			Regions  []string
			State    *rubble7.State
			Tag      string
			URL      string
			Back_URL string
			Name     string
		}{regions, state, "TileSet", "/tsetrun", "/tsetaddons", "Tileset Application"})
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
	})

	http.HandleFunc("/tsetrun", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/tsetrun\"")

		region := state.VariableData["_RUBBLE_IA_REGION_"]
		if region == "" {
			region = "raw"
		}

		addons := []string{}
		for _, addon := range state.Addons {
			if addon.Active {
				addons = append(addons, addon.Meta.Name)
			}
		}

		timeStart := time.Now()
		err = rubble7.TSetModeRun(region, RblDir, DFDir, *AddonsDir, addons, log)
		log.Println("Run time: ", time.Since(timeStart))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
		log.Println("Done.")
		
		if log.WC > 0 {
			http.Redirect(w, r, "/log?state=warn", http.StatusFound)
			return
		}
		
		http.Redirect(w, r, "/log?header=true", http.StatusFound)
	})

	// Independent Apply

	http.HandleFunc("/iaaddons", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/iaaddons\"")

		regions := axis.ListDir(state.FS, "df:data/save")
		for i := range regions {
			if regions[i] == "current" {
				regions = append(regions[:i], regions[i+1:]...)
				break
			}
		}

		for _, addon := range state.Addons {
			if !addon.Meta.Tags["SaveSafe"] {
				addon.Active = false
			}
		}

		state.VariableData["_RUBBLE_IA_REGION_"] = ""

		err = tmpl.ExecuteTemplate(w, "iaaddons", struct {
			Regions  []string
			State    *rubble7.State
			Tag      string
			URL      string
			Back_URL string
			Name     string
		}{regions, state, "SaveSafe", "/iarun", "/iaaddons", "Independent Application"})
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
	})

	http.HandleFunc("/iarun", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/iarun\"")

		region := state.VariableData["_RUBBLE_IA_REGION_"]
		if region == "" {
			region = "raw"
		}

		addons := []string{}
		for _, addon := range state.Addons {
			if addon.Active {
				addons = append(addons, addon.Meta.Name)
			}
		}

		timeStart := time.Now()
		err = rubble7.IAModeRun(region, RblDir, DFDir, *AddonsDir, addons, log)
		log.Println("Run time: ", time.Since(timeStart))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
		log.Println("Done.")
		
		if log.WC > 0 {
			http.Redirect(w, r, "/log?state=warn", http.StatusFound)
			return
		}
		
		http.Redirect(w, r, "/log?header=true", http.StatusFound)
	})

	// Common

	// Toggle an addon's activation state.
	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		addon := r.FormValue("addon")
		log.Println("Toggling addon: \""+addon+"\" to", r.FormValue("state"))

		if _, ok := state.AddonsTbl[addon]; !ok {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Attempt to toggle non-existent addon: " + addon)
			return
		}
		state.AddonsTbl[addon].Active = r.FormValue("state") == "true"

		w.WriteHeader(http.StatusOK)
	})

	// Set a configuration variable.
	http.HandleFunc("/setvar", func(w http.ResponseWriter, r *http.Request) {
		vname := r.FormValue("key")
		log.Println("Setting variable: \""+vname+"\" to", r.FormValue("val"))

		state.VariableData[vname] = r.FormValue("val")

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/pleasewait", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.ExecuteTemplate(w, "pleasewait", r.FormValue("to"))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
	})

	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/log\"")
		es := r.FormValue("state")
		header := r.FormValue("header") != "" || es != ""
		err = tmpl.ExecuteTemplate(w, "log", struct {
			State  string
			Header bool
			Log    string
		}{es, header, logBuffer.String()})
		if err != nil {
			http.Error(w, err.Error(), http.StatusFound)
		}
	})

	http.HandleFunc("/addondata", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/addondata\"")

		addref, ok := state.AddonsTbl[r.FormValue("addon")]
		if !ok {
			log.Println("Error: Invalid addon name in query.")
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}

		url := r.FormValue("fromurl")
		if url == "" {
			url = "/addons"
		}
		name := r.FormValue("fromname")
		if name == "" {
			name = "Addon List"
		}

		err = tmpl.ExecuteTemplate(w, "addondata", struct {
			Addon    *addon.Addon
			FromUrl  string
			FromName string
		}{addref, url, name})
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
	})

	http.HandleFunc("/css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		fmt.Fprint(w, css)
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/ico")
		fmt.Fprint(w, ico)
	})

	http.HandleFunc("/axis", func(w http.ResponseWriter, r *http.Request) {
		log.Println("AXIS file request:", r.FormValue("path"))

		content, err := axis.ReadAll(state.FS, r.FormValue("path"))
		if err != nil {
			log.Println("AXIS error (non-fatal):", err)
			http.Error(w, "AXIS error: "+err.Error(), http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/addonfile", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("addon")
		fname := r.FormValue("file")
		log.Println("Addon file request: Addon:", name, "File:", fname)

		addon, ok := state.AddonsTbl[name]
		if !ok {
			log.Println("Nonexistent addon: \"" + name + "\".")
			http.Error(w, "Nonexistent addon: \""+name+"\".", http.StatusNotFound)
			return
		}

		file, ok := addon.Files[fname]
		if !ok {
			log.Println("Nonexistent file: \"" + fname + "\" in addon: \"" + name + "\".")
			http.Error(w, "Nonexistent file: \""+fname+"\" in addon: \""+name+"\".", http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, "%s", file.Content)
	})

	http.HandleFunc("/wshop", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Rendering workshop:", r.FormValue("id"))

		name := r.FormValue("addon")
		fname := r.FormValue("file")

		addon, ok := state.AddonsTbl[name]
		if !ok {
			log.Println("  Nonexistent addon: \"" + name + "\".")
			http.Error(w, "Nonexistent addon: \""+name+"\".", http.StatusNotFound)
			return
		}

		file, ok := addon.Files[fname]
		if !ok {
			log.Println("  Nonexistent file: \"" + fname + "\" in addon: \"" + name + "\".")
			http.Error(w, "Nonexistent file: \""+fname+"\" in addon: \""+name+"\".", http.StatusNotFound)
			return
		}

		raws := rparse.ParseRaws([]byte(brender.Fix(string(file.Content))))
		raws = brender.Isolate(r.FormValue("id"), raws)
		if len(raws) == 0 {
			log.Println("  Could not find building to render.")
			http.Error(w, "Could not find building to render.", http.StatusNotFound)
			return
		}

		wshop, err := brender.Parse(raws, 7, 0)
		if err != nil {
			log.Println("  Building parse error:", err)
			http.Error(w, "Building parse error: "+err.Error(), http.StatusNotFound)
			return
		}

		stg, err := strconv.ParseInt(r.FormValue("stage"), 10, 8)
		if err != nil {
			stg = 3
		}

		tsetName := ""
		init := rparse.ParseRaws(state.Init)
		for _, tag := range init {
			if tag.ID == "GRAPHICS_FULLFONT" {
				tsetName = tag.Params[0]
				break
			}
		}
		if tsetName == "" {
			log.Println("  Could not find current tileset name.")
			http.Error(w, "Could not find current tileset name.", http.StatusNotFound)
			return
		}

		tset, err := axis.Read(state.FS, "df:data/art/"+tsetName)
		if err != nil {
			log.Println("  AXIS error (non-fatal):", err)
			http.Error(w, "AXIS error: "+err.Error(), http.StatusNotFound)
			return
		}

		tileset, _, err := image.Decode(tset)
		tset.Close()
		if err != nil {
			log.Println("  Image decode error:", err)
			http.Error(w, "Image decode error: "+err.Error(), http.StatusNotFound)
			return
		}

		img := brender.Render(wshop, int(stg), tileset)

		err = png.Encode(w, img)
		if err != nil {
			log.Println("  Image encode error:", err)
			//http.Error(w, "Image encode error: " + err.Error(), http.StatusNotFound)
			return
		}
	})

	http.HandleFunc("/doclist", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/doclist\"")

		err := tmpl.ExecuteTemplate(w, "doclist", state.DocPacks)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
	})

	http.HandleFunc("/doc/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"" + r.URL.Path + "\"")

		file := strings.TrimPrefix(r.URL.Path, "/doc/")
		content, err := axis.ReadAll(state.FS, "rubble:other/"+file+".md")
		if err != nil {
			content, err = axis.ReadAll(state.FS, "addons:dir:"+file+".md")
			if err != nil {
				content, err = axis.ReadAll(state.FS, "addons:zip:"+file+".md")
				if err != nil {
					log.Println("Doc file: " + file + " not found.")
					http.Redirect(w, r, "/log?state=error", http.StatusFound)
					return
				}
			}
		}

		err = tmpl.ExecuteTemplate(w, "docpage", rblutil.MarkdownToHTML(string(content)))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}
	})

	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/kill\"")

		err := tmpl.ExecuteTemplate(w, "kill", nil)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log?state=error", http.StatusFound)
			return
		}

		d, _ := time.ParseDuration("1s")
		time.AfterFunc(d, func() { os.Exit(0) })
	})

	log.Separator()
	LaunchBrowser(log, RblDir, Addr)

	log.Println("Starting Server...")
	err = http.ListenAndServe(Addr, nil)
	if err != nil {
		log.Println("  Server Startup Failed:\n    Error:", err)
		os.Exit(1)
	}
}

// isAbs returns true if the path is not a relative path (includes no "." or ".." parts).
func isAbs(path string) bool {
	dirs := strings.Split(path, "/")

	for i := range dirs {
		if dirs[i] == ".." || dirs[i] == "." {
			return false
		}
	}
	return true
}
