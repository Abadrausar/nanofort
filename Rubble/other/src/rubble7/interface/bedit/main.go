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

// Rubble Building Editor Interface
package main

import "net/http"

import "os"
import "fmt"
import "bytes"
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

import "dctech/axis"

var Addr string

var RblDir string
var DFDir string
var OutputDir string
var AddonsDir *rblutil.ArgList

var AddonsList *rblutil.ArgList
var ConfigList *rblutil.ArgList

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

	Addr = "127.0.0.1:2121"

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

	flags.Parse(os.Args[1:])

	if AddonsDir.Empty() {
		AddonsDir.Set("rubble:addons")
	}

	err, state := rubble7.NewState(RblDir, DFDir, OutputDir, *AddonsDir, log)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	tileSheet := ""
	init := rparse.ParseRaws(state.Init)
	for _, tag := range init {
		if tag.ID == "GRAPHICS_FULLFONT" {
			tileSheet = tag.Params[0]
			break
		}
	}
	if tileSheet == "" {
		log.Println("Could not find current tileset name in df:data/init.txt.")
		http.Redirect(w, r, "/log", http.StatusFound)
		return
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

		err = tmpl.ExecuteTemplate(w, "menu", state)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log", http.StatusFound)
			return
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Redirecting:", r.URL, "to main menu.")

		http.Redirect(w, r, "./menu", http.StatusFound)
	})

	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/log\"")
		err = tmpl.ExecuteTemplate(w, "log", logBuffer.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusFound)
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

	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		log.Println("UI Transition: \"/kill\"")

		err := tmpl.ExecuteTemplate(w, "kill", nil)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/log", http.StatusFound)
			return
		}

		d, _ := time.ParseDuration("1s")
		time.AfterFunc(d, func() { os.Exit(0) })
	})

	http.HandleFunc("/sheet", func(w http.ResponseWriter, r *http.Request) {
		fgStr := r.FormValue("fg")
		bgStr := r.FormValue("bg")
		
		tset, err := axis.Read(state.FS, "df:data/art/"+tileSheet)
		if err != nil {
			log.Println("AXIS error:", err)
			http.Redirect(w, r, "/log", http.StatusFound)
			return
		}

		tileset, _, err := image.Decode(tset)
		tset.Close()
		if err != nil {
			log.Println("Image decode error:", err)
			http.Redirect(w, r, "/log", http.StatusFound)
			return
		}
		
		// TODO: Colorize
		img := tileset
		
		err = png.Encode(w, img)
		if err != nil {
			log.Println("Image encode error:", err)
			//http.Redirect(w, r, "/log", http.StatusFound)
			return
		}
	}
	
	http.HandleFunc("/wshop", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Rendering workshop...")

		

		tset, err := axis.Read(state.FS, "df:data/art/"+tileSheet)
		if err != nil {
			log.Println("AXIS error:", err)
			http.Redirect(w, r, "/log", http.StatusFound)
			return
		}

		tileset, _, err := image.Decode(tset)
		tset.Close()
		if err != nil {
			log.Println("Image decode error:", err)
			http.Redirect(w, r, "/log", http.StatusFound)
			return
		}

		img := brender.Render(wshop, int(stg), tileset)

		err = png.Encode(w, img)
		if err != nil {
			log.Println("Image encode error:", err)
			//http.Redirect(w, r, "/log", http.StatusFound)
			return
		}
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
