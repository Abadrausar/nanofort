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

// Rubble main package, this contains everything that makes Rubble actually work.
package rubble7

import "strings"
import "bytes"
import "strconv"

import "dctech/axis"

import "rubble7/rblutil"
import "rubble7/rblutil/merge"
import "rubble7/rblutil/parse"
import "rubble7/rblutil/rparse"
import "rubble7/rblutil/addon"
import "rubble7/rblutil/errors"

// Used by the script registry in the state.
type ScrRegData struct {
	List  *[]string // This needs to be a pointer, don't ask.
	Table map[string]string
}

// State is the core of Rubble, everything connects to the state at some level.
type State struct {
	Log *rblutil.Logger

	// The global AXIS filesystem.
	FS axis.Collection

	// Important paths (OS paths, not AXIS!).
	RblDir    string
	DFDir     string
	OutputDir string
	AddonsDir []string

	// The current parse stage, only valid while parsing, obviously.
	Stage parse.Stage

	// The files of the active loaded addons.
	Files *addon.FileList

	// Global scripts (load and init scripts).
	GlobalFiles *addon.FileList

	// All the loaded addons.
	Addons    []*addon.Addon
	AddonsTbl map[string]*addon.Addon

	DocPacks []*addon.Meta

	// Just the meta-data, this is mostly for the web UI.
	AddonMeta map[string]*addon.Meta

	// Script is the script operation arbitrator.
	Script ScriptCore

	// Shared script data registry.
	ScrRegistry map[string]ScrRegData

	// This is where config variables are stored.
	VariableData map[string]string

	Templates map[string]*Template

	// The file being parsed/executed right now or "".
	CurrentFile string

	// Some files that need special handling
	Init   []byte
	D_Init []byte
}

// NewState creates a new Rubble State with the paths provided.
func NewState(rubbledir, dfdir, output string, addonsdir []string, log *rblutil.Logger) (error, *State) {
	// First create the basic state.
	state := new(State)

	state.Log = log
	state.Log.ClearWarnings()

	state.Log.Separator()
	state.Log.Println("Initializing...")
	state.Log.Println("  Creating New State...")

	
	state.Files = addon.NewFileList()
	state.GlobalFiles = addon.NewFileList()
	state.AddonsTbl = make(map[string]*addon.Addon)
	state.AddonMeta = make(map[string]*addon.Meta)
	state.ScrRegistry = make(map[string]ScrRegData)
	state.VariableData = make(map[string]string)
	state.Templates = make(map[string]*Template)

	// Now setup the global stuff.
	
	state.Log.Println("  Initializing AXIS VFS...")
	state.FS = initAXIS(rubbledir, dfdir, output, addonsdir, state)

	state.Log.Println("  Loading Special Global Files...")
	state.Log.Println("    Reading DF Init Files...")

	var err error
	state.Init, err = axis.ReadAll(state.FS, "df:data/init/init.txt")
	if err != nil {
		return err, nil
	}
	state.D_Init, err = axis.ReadAll(state.FS, "df:data/init/d_init.txt")
	if err != nil {
		return err, nil
	}

	state.Script = NewScriptCore(state)

	return nil, state
}

// InitAXIS creates the standard AXIS filesystem used bu Rubble.
// This is exported for cases where you want to do some Rubble-style IO but don't need a whole State.
func InitAXIS(rubbledir, dfdir, output string, addonsdir []string) axis.Collection {
	return initAXIS(rubbledir, dfdir, output, addonsdir, nil)
}

func initAXIS(rubbledir, dfdir, output string, addonsdir []string, state *State) axis.Collection {
	// Massage some of the path variables to allow AXIS paths before AXIS is setup.
	dfdir = strings.Replace(dfdir, "rubble:", rubbledir+"/", 1)

	output = strings.Replace(output, "rubble:", rubbledir+"/", 1)
	output = strings.Replace(output, "df:", dfdir+"/", 1)

	for i := range addonsdir {
		addonsdir[i] = strings.Replace(addonsdir[i], "rubble:", rubbledir+"/", 1)
		addonsdir[i] = strings.Replace(addonsdir[i], "df:", dfdir+"/", 1)
		addonsdir[i] = strings.Replace(addonsdir[i], "out:", output+"/", 1)
	}
	
	if state != nil {
		state.RblDir = rubbledir
		state.DFDir = dfdir
		state.OutputDir = output
		state.AddonsDir = addonsdir
	}
	
	fs := axis.NewFileSystem()
	
	fs.Mount("df", axis.NewOSDir(dfdir, true))
	fs.Mount("rubble", axis.NewOSDir(rubbledir, true))
	fs.Mount("out", axis.NewOSDir(output, true))

	addonsFS := axis.NewFileSystem()
	zipFS := axis.NewLogicalDir()
	addonsFS.Mount("zip", zipFS)
	dirs := []axis.DataSource{}
	for i := range addonsdir {
		dirs = append(dirs, axis.NewOSDir(addonsdir[i], true))
	}
	addonsFS.Mount("dir", axis.NewMultiplexer(dirs...))
	fs.Mount("addons", addonsFS)
	
	return fs
}

// Run runs a full Rubble parse cycle.
// See *State.Activate for parameter descriptions.
func (state *State) Run(addons, config []string) (err error) {
	err = state.Load(addons, config)
	if err != nil {
		return err
	}

	state.Log.Println("  Updating the Default Addon List File...")
	err = state.UpdateAddonList("addons:dir:addonlist.ini")
	if err != nil {
		return err
	}

	state.Log.Separator()
	state.Log.Println("Generating Sorted Active File List...")
	state.Files.Update(state.Addons)

	err = state.Parse()
	if err != nil {
		return err
	}

	err = state.ApplyTileSet()
	if err != nil {
		return err
	}

	err = state.Write()
	if err != nil {
		return err
	}

	err = state.WriteReport()
	if err != nil {
		return err
	}

	return nil
}

// RunPreLoaded runs a full Rubble parse cycle, minus loading addons.
// See *State.Activate for parameter descriptions.
func (state *State) RunPreLoaded(addons, config []string) (err error) {
	err = state.Activate(addons, config)
	if err != nil {
		return err
	}

	err = state.RunActivated()
	if err != nil {
		return err
	}

	return nil
}

// RunActivated runs a full Rubble parse cycle, minus loading and activating addons.
// See *State.Activate for parameter descriptions.
func (state *State) RunActivated() (err error) {
	state.Log.Println("  Updating the Default Addon List File...")
	err = state.UpdateAddonList("addons:dir:addonlist.ini")
	if err != nil {
		return err
	}

	state.Log.Separator()
	state.Log.Println("Generating Sorted Active File List...")
	state.Files.Update(state.Addons)

	err = state.Parse()
	if err != nil {
		return err
	}

	err = state.ApplyTileSet()
	if err != nil {
		return err
	}

	err = state.Write()
	if err != nil {
		return err
	}

	err = state.WriteReport()
	if err != nil {
		return err
	}

	return nil
}

// Load loads all addons and determines which ones should be
// active, then writes the default addon list file.
// This is where most of the configuration magic happens...
// addons contains addon activation information.
// Each entry must be either:
//	A list of addon names delimited by semicolons (each addon given is activated)
//	The name of an INI file that contains addon names and their activation state, paths must use the AXIS syntax.
// If addons is nil the default addons file is used (addons:dir:addonlist.ini).
// config is exactly the same as addons, just for configuration variables.
func (state *State) Load(addons, config []string) (err error) {
	defer errors.TrapError(&err, state.Log)

	state.Addons, state.GlobalFiles = addon.Load(state.FS, state.Log)
	for _, addon := range state.Addons {
		state.AddonsTbl[addon.Meta.Name] = addon
		state.AddonMeta[addon.Meta.Name] = addon.Meta
		if addon.Meta.Tags["DocPack"] {
			state.DocPacks = append(state.DocPacks, addon.Meta)
		}
	}

	return state.Activate(addons, config)
}

// Activate reruns some of the steps Load does.
// addons contains addon activation information.
// Each entry must be either:
//	A list of addon names delimited by semicolons (each addon given is activated)
//	The name of an INI file that contains addon names and their activation state, paths must use the AXIS syntax.
// If addons is empty the default addons file is used (addons:dir:addonlist.ini).
// config is exactly the same as addons, just for configuration variables.
func (state *State) Activate(addons, config []string) (err error) {
	return state.activate(addons, config, false)
}

// ActivateAny is the same as Activate, except it also allows direct activation of library addons.
func (state *State) ActivateAny(addons, config []string) (err error) {
	return state.activate(addons, config, true)
}

func (state *State) activate(addons, config []string, libsOK bool) (err error) {
	defer errors.TrapError(&err, state.Log)

	state.Log.Separator()
	state.Log.Println("Activating...")

	state.Log.Println("  Clearing Leftover Addon State Information...")
	for i := range state.Addons {
		state.Addons[i].Active = false
	}

	state.Log.Println("  Loading Config Variables...")
	if config != nil && len(config) != 0 {
		for _, val := range config {
			file, err := axis.ReadAll(state.FS, val)
			if err == nil {
				state.Log.Println("    Loading Config File: " + val)
				rblutil.ParseINI(string(file), "\n", func(key, value string) {
					state.VariableData[key] = value
				})
				continue
			}
			rblutil.ParseINI(val, ";", func(key, value string) {
				state.VariableData[key] = value
			})
		}
	} else {
		state.Log.Println("    No variables specified.")
	}

	state.Log.Println("  Generating Active Addon List...")
	if len(addons) == 0 {
		state.Log.Println("    No addons specified, using default addon list file.")
		addons = []string{"addons:dir:addonlist.ini"}
	}

	addonNames := make([]string, 0, 10)
	for _, val := range addons {
		file, err := axis.ReadAll(state.FS, val)
		if err == nil {
			state.Log.Println("    Loading List File: " + val)
			rblutil.ParseINI(string(file), "\n", func(key, value string) {
				value = strings.ToLower(value)
				if ok, _ := strconv.ParseBool(value); ok {
					addonNames = append(addonNames, key)
				}
			})
		} else {
			addonNames = append(addonNames, strings.Split(val, ";")...)
		}
	}

	state.Log.Println("  Activating Addons from Generated List...")
	for _, name := range addonNames {
		if _, ok := state.AddonsTbl[name]; ok {
			state.AddonsTbl[name].Active = true
		}
	}

	if !libsOK {
		state.Log.Println("  Pruning Library Addons...")
		for i := range state.Addons {
			if state.Addons[i].Meta.Tags["Library"] {
				state.Addons[i].Active = false
			}
		}
	} else {
		state.Log.Println("  OK to Activate Libraries Directly, Skipping Pruning Step.")
	}
	
	state.Log.Println("  Activating Required Addons from Meta Data...")
	var activate func(string)
	activate = func(me string) {
		for j := range state.AddonsTbl[me].Meta.Activates {
			it := state.AddonsTbl[me].Meta.Activates[j]
			if _, ok := state.AddonsTbl[it]; !ok {
				errors.RaiseAbort("The \"" + state.AddonsTbl[me].Meta.Name + "\" addon requires the \"" + it + "\" addon!\n" +
					"The required addon is not currently installed, please install the required addon and try again.")
			}

			if !state.AddonsTbl[it].Meta.Tags["Library"] && !state.AddonsTbl[it].Active {
				state.Log.Println("    Warning: The \"" + me + "\" Addon is Activating Non-Library Addon: \"" + it + "\"")
			}

			state.AddonsTbl[it].Active = true
			activate(it)
		}
	}
	for me := range state.AddonsTbl {
		if state.AddonsTbl[me].Active {
			activate(me)
		}
	}

	state.Log.Println("  Running Loader Scripts...")
	err = state.GlobalFiles.RunAction(map[string]bool{
		"Skip":       false,
		"DFHack":     false,
		"LoadScript": true,
	}, func(file *addon.File) error {
		state.CurrentFile = file.Name
		state.Log.Println("    " + file.Name)

		_, err := state.Script.RunScriptFile(file)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	state.Log.Println("  Active Addons:")
	for i := range state.Addons {
		if state.Addons[i].Active {
			state.Log.Println("    " + state.Addons[i].Meta.Name)
		}
	}

	state.Log.Println("  Checking for Incompatible Addons from Meta Data...")
	for _, addon := range state.Addons {
		if addon.Active {
			for _, incompat := range addon.Meta.Incompatible {
				if iaddon, ok := state.AddonsTbl[incompat]; ok && iaddon.Active {
					errors.RaiseAbort("The \"" + addon.Meta.Name + "\" addon is incompatible with the \"" + iaddon.Meta.Name + "\" addon!\n" +
						"Please deactivate one of these addons and try again.")
				}
			}
		}
	}

	state.Log.Println("  Setting Unset Config Variables to Their Defaults...")
	state.Log.Println("    Attempting to load rubble:userconfig.ini...")
	file, err := axis.ReadAll(state.FS, "rubble:userconfig.ini")
	if err == nil {
		rblutil.ParseINI(string(file), "\n", func(key, value string) {
			if _, ok := state.VariableData[key]; !ok {
				state.Log.Println("      " + key)
				state.VariableData[key] = value
			}
		})
	} else {
		state.Log.Println("      Could not load rubble:userconfig.ini:", err)
		state.Log.Println("      (This is probably OK.)")
	}
	state.Log.Println("    Setting Remaining Variables from Meta-Data...")
	for _, addon := range state.Addons {
		if addon.Active {
			for name, vdat := range addon.Meta.Vars {
				if _, ok := state.VariableData[name]; !ok && len(vdat.Values) > 0 {
					state.Log.Println("      " + name)
					state.VariableData[name] = vdat.Values[0]
				}
			}
		}
	}
	state.Log.Println("    (Any variables not listed above were explicitly set elsewhere.)")

	return nil
}

// Parse handles everything from running init scripts to running post scripts.
// If an error is returned state.CurrentFile will still be set to its last value.
func (state *State) Parse() (err error) {
	defer errors.TrapError(&err, state.Log)

	runScript := func(file *addon.File) error {
		state.CurrentFile = file.Name
		state.Log.Println("  " + file.Name)

		_, err := state.Script.RunScriptFile(file)
		if err != nil {
			return err
		}
		return nil
	}

	rawFile := map[string]bool{"Skip": false, "AUXTextFile": false, "TileSet": false, "RawFile": true}
	parseRaw := func(stage parse.Stage) func(file *addon.File) error {
		state.Stage = stage
		return func(file *addon.File) error {
			state.CurrentFile = file.Name
			state.Log.Println("  " + file.Name)

			file.Content = parse.Parse(file.Content, file.Name, 1, stage, state.Dispatcher)
			return nil
		}
	}

	state.Log.Separator()
	state.Log.Println("Running Init Scripts...")
	err = state.GlobalFiles.RunAction(map[string]bool{
		"Skip":       false,
		"InitScript": true,
	}, runScript)
	if err != nil {
		return err
	}

	state.Log.Separator()
	state.Log.Println("Running Prescripts...")
	err = state.Files.RunAction(map[string]bool{
		"Skip":      false,
		"PreScript": true,
	}, runScript)
	if err != nil {
		return err
	}

	state.Log.Separator()
	state.Log.Println("Preparsing...")
	state.Files.RunAction(rawFile, parseRaw(parse.StgPreParse))

	state.Log.Separator()
	state.Log.Println("Parsing...")
	state.Files.RunAction(rawFile, parseRaw(parse.StgParse))

	state.Log.Separator()
	state.Log.Println("Postparsing...")
	state.Files.RunAction(rawFile, parseRaw(parse.StgPostParse))

	state.Log.Separator()
	state.Log.Println("Running Postscripts...")
	err = state.Files.RunAction(map[string]bool{
		"Skip":       false,
		"PostScript": true,
	}, runScript)
	if err != nil {
		return err
	}

	state.CurrentFile = ""
	return nil
}

// ApplyTileSet does exactly what it's name suggests.
// Tileset information is loaded from all active addons that have it and is
// applied to all normal raw files.
func (state *State) ApplyTileSet() error {
	state.Log.Separator()
	state.Log.Println("Applying Tileset...")

	state.Log.Println("  Loading Rules Files...")
	rules := new(merge.RuleNode)
	err := state.Files.RunAction(map[string]bool{
		"Skip":       false,
		"TileSet":    true,
		"MergeRules": true,
	}, func(file *addon.File) error {
		state.Log.Println("    " + file.Name)
		return merge.ParseRules(file.Content, rules)
	})
	if err != nil {
		return err
	}

	state.Log.Println("  Loading Tileset Files...")
	set := new(merge.TagNode)
	state.Files.RunAction(map[string]bool{
		"Skip":    false,
		"RawFile": true,
		"TileSet": true,
	}, func(file *addon.File) error {
		state.Log.Println("    " + file.Name)

		merge.PopulateTree(rparse.ParseRaws(file.Content), set, rules)
		return nil
	})

	state.Log.Println("  Applying Tileset Information...")
	state.Files.RunAction(map[string]bool{
		"Skip":             false,
		"NoWrite":          false,
		"AUXTextFile":      false,
		"CreatureGraphics": false,
		"TileSet":          false,
		"RawFile":          true,
	}, func(file *addon.File) error {
		state.Log.Println("    " + file.Name)

		tags := rparse.ParseRaws(file.Content)
		merge.Apply(tags, set)
		file.Content = rparse.FormatFile(tags)
		return nil
	})

	state.Log.Println("  Applying init Patches...")

	tags := rparse.ParseRaws(state.Init)
	// tags[0] is always a dummy tag used to preserve leading comments, so let's
	// use it for something useful.
	tags[0].ID = "AUX_FILE"
	tags[0].Params = []string{"INIT.TXT"}
	tags[0].CommentsOnly = false // Clear the comments flag so the merger can see the tag.
	merge.Apply(tags, set)
	tags[0].CommentsOnly = true // Then set the flag again so the formatter does not include it.
	state.Init = rparse.FormatFile(tags)
	
	tags = rparse.ParseRaws(state.D_Init)
	tags[0].ID = "AUX_FILE"
	tags[0].Params = []string{"D_INIT.TXT"}
	tags[0].CommentsOnly = false
	merge.Apply(tags, set)
	tags[0].CommentsOnly = true
	state.D_Init = rparse.FormatFile(tags)

	state.Log.Println("  Running Tileset Scripts...")
	err = state.Files.RunAction(map[string]bool{
		"Skip":    false,
		"TileSet": true,
	}, func(file *addon.File) error {
		if !state.Script.IsScriptFile(file) {
			return nil
		}

		state.CurrentFile = file.Name
		state.Log.Println("    " + file.Name)

		_, err := state.Script.RunScriptFile(file)
		if err != nil {
			return err
		}
		return nil
	})
	state.CurrentFile = ""
	if err != nil {
		return err
	}
	return nil
}

// Write handles writing the files to their output directories.
func (state *State) Write() error {
	state.Log.Separator()
	
	state.Log.Println("Checking Raw Files for Consistency...")
	state.Log.Println("  Loading Consistency Checker Rules...")
	rules := new(merge.RuleNode)
	err := state.Files.RunAction(map[string]bool{
		"Skip":       false,
		"TileSet":    false,
		"MergeRules": true,
	}, func(file *addon.File) error {
		state.Log.Println("  " + file.Name)
		return merge.ParseRules(file.Content, rules)
	})
	if err != nil {
		return err
	}
	state.Log.Println("  Running Raw Consistency Checker...")
	pw := state.Log.WC
	err = state.Files.RunAction(map[string]bool{
		"Skip":             false,
		"NoWrite":          false,
		"AUXTextFile":      false,
		"TileSet":          false,
		"RawFile":          true,
	}, func(file *addon.File) error {
		a := bytes.Count(file.Content, []byte("["))
		b := bytes.Count(file.Content, []byte("]"))
		if a != b {
			state.Log.WarnOnlyln("    Consistency warnings for: " + file.Name)
			state.Log.WarnOnlyExtraf("      Opening brace count does not match closing brace count (%v/%v).\n", a, b)
			state.Log.WarnOnlyExtraln("        This may indicate bad output, disabled raw tags, or simply literal braces in text.")
		}
		
		errs := merge.Match(rparse.ParseRaws(file.Content), rules)
		if len(errs) != 0 {
			if a == b {
				state.Log.WarnOnlyln("    Consistency warnings for: " + file.Name)
			}
			state.Log.WarnOnlyExtraln("      Rule match failures:")
			for _, err := range errs {
				state.Log.WarnOnlyExtraln("        " + err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	if pw != state.Log.WC {
		state.Log.Printf("  %v Consistency issues found, check the warnings section for details.\n", state.Log.WC - pw)
	} else {
		state.Log.Println("  No consistency issues found. Yay!")
	}
	
	state.Log.Println("Writing Files...")
	state.Log.Println("  Writing Raw Files...")
	err = state.writeFiles(state.Files, "out:objects", map[string]bool{
		"Skip":             false,
		"NoWrite":          false,
		"AUXTextFile":      false,
		"CreatureGraphics": false,
		"TileSet":          false,
		"RawFile":          true,
	}, "#", "", "", true)
	if err != nil {
		return err
	}

	state.Log.Println("  Writing Auxiliary Text Files...")
	err = state.writeFiles(state.Files, "out:objects/text", map[string]bool{
		"Skip":        false,
		"NoWrite":     false,
		"AUXTextFile": true,
	}, "", ".text.txt", ".txt", false)
	if err != nil {
		return err
	}

	state.Log.Println("  Writing Graphics Files...")
	err = state.writeFiles(state.Files, "out:graphics", map[string]bool{
		"Skip":             false,
		"NoWrite":          false,
		"CreatureGraphics": true,
		"RawFile":          true,
	}, "#", ".graphics.txt", ".txt", true)
	if err != nil {
		return err
	}
	err = state.writeFiles(state.Files, "out:graphics", map[string]bool{
		"Skip":             false,
		"NoWrite":          false,
		"CreatureGraphics": true,
		"ImagePNG":         true,
	}, "", ".graphics.png", ".png", false)
	if err != nil {
		return err
	}
	err = state.writeFiles(state.Files, "out:graphics", map[string]bool{
		"Skip":             false,
		"NoWrite":          false,
		"CreatureGraphics": true,
		"ImageBMP":         true,
	}, "", ".graphics.bmp", ".bmp", false)
	if err != nil {
		return err
	}

	state.Log.Println("  Writing Tileset Files...")
	err = state.writeFiles(state.Files, "df:data/art", map[string]bool{
		"Skip":     false,
		"NoWrite":  false,
		"TileSet":  true,
		"ImagePNG": true,
	}, "", ".tset.png", ".png", false)
	if err != nil {
		return err
	}
	err = state.writeFiles(state.Files, "df:data/art", map[string]bool{
		"Skip":     false,
		"NoWrite":  false,
		"TileSet":  true,
		"ImageBMP": true,
	}, "", ".tset.bmp", ".bmp", false)
	if err != nil {
		return err
	}

	state.Log.Println("  Writing DFHack Pseudo Modules...")
	err = state.writeFiles(state.Files, "out:modules", map[string]bool{
		"Skip":         false,
		"NoWrite":      false,
		"DFHack":       true,
		"ModuleScript": true,
		"LangLua":      true,
	}, "--", ".dfmod.lua", ".lua", false)
	if err != nil {
		return err
	}

	state.Log.Println("  Writing DFHack Load Scripts...")
	err = state.writeFiles(state.Files, "out:init.d", map[string]bool{
		"Skip":       false,
		"NoWrite":    false,
		"DFHack":     true,
		"LoadScript": true,
		"LangLua":    true,
	}, "--", ".dfload.lua", ".lua", false)
	if err != nil {
		return err
	}

	state.Log.Println("  Writing DFHack Command Scripts...")
	err = state.writeFiles(state.Files, "out:scripts", map[string]bool{
		"Skip":          false,
		"NoWrite":       false,
		"DFHack":        true,
		"CommandScript": true,
		"LangLua":       true,
	}, "--", ".dfcom.lua", ".lua", false)
	if err != nil {
		return err
	}
	err = state.writeFiles(state.Files, "out:scripts", map[string]bool{
		"Skip":          false,
		"NoWrite":       false,
		"DFHack":        true,
		"CommandScript": true,
		"LangRuby":      true,
	}, "#", ".dfcom.rb", ".rb", false)
	if err != nil {
		return err
	}

	state.Log.Println("  Writing Init Files...")
	state.Log.Println("    init.txt")
	err = state.writeFile("df:data/init/init.txt", state.Init)
	if err != nil {
		return err
	}
	state.Log.Println("    d_init.txt")
	err = state.writeFile("df:data/init/d_init.txt", state.D_Init)
	if err != nil {
		return err
	}

	state.Log.Println("    overrides.txt")
	state.Log.Println("      Assembling TWBT Overrides File...")
	overrides := []byte("\n# Automatically generated, do not edit!\n")
	twbt_count := 0
	state.Files.RunAction(map[string]bool{
		"Skip":         false,
		"NoWrite":      false,
		"TileSet":      true,
		"TWBTOverride": true,
	}, func(file *addon.File) error {
		state.Log.Println("        " + file.Name)
		twbt_count++

		overrides = append(overrides, ("\n# Source: " +
			file.Source + "/" + file.Name + "\n" + string(file.Content))...)
		return nil
	})
	if twbt_count > 0 {
		state.Log.Println("      Writing File.")
		err = state.writeFile("df:data/init/overrides.txt", overrides)
		if err != nil {
			return err
		}
	} else {
		state.Log.Println("      No Overrides Found.")
	}

	return nil
}

// WriteReport adds a "generation report" to the output directory.
// This report includes information about the configuration variables,
// active addons, and tileset.
func (state *State) WriteReport() error {
	state.Log.Separator()
	state.Log.Println("Writing Generation Report...")

	state.Log.Println("  Writing Addon List...")
	state.Log.Println("    addonlist.ini")
	err := state.UpdateAddonList("out:addonlist.ini")
	if err != nil {
		return err
	}

	state.Log.Println("  Writing Config Variables...")
	state.Log.Println("    genconfig.ini")
	err = state.DumpConfig("out:genconfig.ini")
	if err != nil {
		return err
	}

	state.Log.Println("  Writing Tileset Config File...")
	state.Log.Println("    current.tset.rbl")
	err = state.DumpTSet("out:current.tset.rbl")
	if err != nil {
		return err
	}

	state.Log.Printf("  %v Warnings Encountered While Generating", state.Log.WC)
	if state.Log.WC > 0 {
		state.Log.Println(":")
	} else {
		state.Log.Println(".")
	}
	state.Log.Print(state.Log.WB.String())
	
	return nil
}

func (state *State) writeFiles(list *addon.FileList, dir string, filter map[string]bool, comment, exthas, extgive string, addHeader bool) error {
	return list.RunAction(filter, func(file *addon.File) error {
		state.Log.Println("    " + file.Name)

		name := file.Name
		if exthas != "" {
			name = rblutil.ReplaceExt(name, exthas, extgive)
		}

		content := file.Content
		if comment != "" {
			content = []byte("\n" + comment + " Automatically generated, do not edit!\n" + comment + " Source: " +
				file.Source + "/" + file.Name + "\n" + string(file.Content))
		}
		if addHeader {
			// exthas is only set if the file has a two-part extension.
			if exthas != "" {
				content = append([]byte(rblutil.StripExt(rblutil.StripExt(file.Name)) + "\n"), content...)
			} else {
				content = append([]byte(rblutil.StripExt(file.Name) + "\n"), content...)
			}
		}

		return state.writeFile(dir+"/"+name, content)
	})
}

func (state *State) writeFile(path string, content []byte) error {
	if !axis.Exists(state.FS, path) {
		err := axis.Create(state.FS, path)
		if err != nil {
			return err
		}
	}

	return axis.WriteAll(state.FS, path, content)
}

// Data Dumpers

// DumpConfig writes all configuration variables (in INI format) to the indicated file.
func (state *State) DumpConfig(path string) error {
	out := "\n# Rubble config variable dump.\n# Automatically generated, do not edit!\n\n[config]\n"

	for i := range state.VariableData {
		out += i + " = " + strconv.Quote(state.VariableData[i]) + "\n"
	}

	return state.writeFile(path, []byte(out))
}

// DumpTSet writes the current tileset values from the loaded raws to the indicated file.
func (state *State) DumpTSet(path string) error {
	rules := new(merge.RuleNode)
	err := state.Files.RunAction(map[string]bool{
		"Skip":       false,
		"TileSet":    true,
		"MergeRules": true,
	}, func(file *addon.File) error {
		return merge.ParseRules(file.Content, rules)
	})
	if err != nil {
		return err
	}

	out := []byte("\n# Rubble tileset dump.\n# Automatically generated, do not edit!\n")

	set := new(merge.TagNode)
	state.Files.RunAction(map[string]bool{
		"Skip":             false,
		"NoWrite":          false,
		"AUXTextFile":      false,
		"CreatureGraphics": false,
		"TileSet":          false,
		"RawFile":          true,
	}, func(file *addon.File) error {
		merge.PopulateTree(rparse.ParseRaws(file.Content), set, rules)
		return nil
	})
	out = append(out, set.String()...)

	return state.writeFile(path, out)
}

// UpdateAddonList writes a list of all addons and their activation status (in INI format) to the indicated file.
func (state *State) UpdateAddonList(dest string) error {
	out := make([]byte, 0, 2048)
	out = append(out, "\n# Rubble addon list.\n# Version: "+Version+"\n# Automatically generated, do not edit!\n\n[addons]\n"...)
	for i := range state.Addons {
		if !state.Addons[i].Meta.Tags["Library"] && !state.Addons[i].Meta.Tags["DocPack"] {
			out = append(out, state.Addons[i].Meta.Name+"="...)
			if state.Addons[i].Active {
				out = append(out, "true\n"...)
			} else {
				out = append(out, "false\n"...)
			}
		}
	}

	return state.writeFile(dest, out)
}
