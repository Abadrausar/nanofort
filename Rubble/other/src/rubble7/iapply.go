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

package rubble7

import "rubble7/rblutil"
import "rubble7/rblutil/errors"
import "rubble7/rblutil/addon"

// IAList returns a list of all non-lib save safe addons installed.
// Use when you do not have a State handy.
func IAList(rbldir, dfdir string, addonsdir []string, log *rblutil.Logger) (err error, addons []string) {
	oops, state := NewState(rbldir, dfdir, "df:raw", addonsdir, log)
	if oops != nil {
		return oops, nil
	}
	defer errors.TrapError(&err, state.Log)

	err = state.Load(nil, nil)
	if err != nil {
		return err, nil
	}

	for addon, meta := range state.AddonMeta {
		if meta.Tags["SaveSafe"] && !meta.Tags["Library"] {
			addons = append(addons, addon)
		}
	}
	return nil, addons
}

// IAModeRun applies one or more save safe addons to the specified region.
// The addon's scripts are run and any DFHack scripts are installed. Raw files
// and the like are not parsed and tileset information is NOT applied!
func IAModeRun(region, rbldir, dfdir string, addonsdir, addons []string, log *rblutil.Logger) (err error) {
	output := dfdir
	if region == "raw" {
		output += "/raw"
	} else {
		output += "/data/save/" + region + "/raw"
	}

	log.Separator()
	log.Println("Entering Independent Application Mode for Region:", region)

	oops, state := NewState(rbldir, dfdir, output, addonsdir, log)
	if oops != nil {
		return oops
	}
	defer errors.TrapError(&err, state.Log)

	state.VariableData["_RUBBLE_NO_CLEAR_"] = "true"

	err = state.Load(addons, nil)
	if err != nil {
		return err
	}

	state.Files.Update(state.Addons)

	runScript := func(file *addon.File) error {
		state.CurrentFile = file.Name
		state.Log.Println("  " + file.Name)

		_, err := state.Script.RunScriptFile(file)
		if err != nil {
			return err
		}
		return nil
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
	state.Log.Println("Running Postscripts...")
	err = state.Files.RunAction(map[string]bool{
		"Skip":       false,
		"PostScript": true,
	}, runScript)
	if err != nil {
		return err
	}

	state.CurrentFile = ""

	// Simulate the important parts of state.Write.
	state.Log.Separator()
	state.Log.Println("Writing Files...")
	state.Log.Println("  Writing DFHack Pseudo Modules...")
	err = state.writeFiles(state.Files, "out:modules", map[string]bool{
		"Skip":               false,
		"NoWrite":            false,
		"DFHackModuleScript": true,
		"LangLua":            true,
	}, "--", ".dfmod.lua", ".lua", false)
	if err != nil {
		return err
	}

	state.Log.Println("  Writing DFHack Load Scripts...")
	err = state.writeFiles(state.Files, "out:init.d", map[string]bool{
		"Skip":             false,
		"NoWrite":          false,
		"DFHackLoadScript": true,
		"LangLua":          true,
	}, "--", ".dfload.lua", ".lua", false)
	if err != nil {
		return err
	}

	state.Log.Println("  Writing DFHack Command Scripts...")
	err = state.writeFiles(state.Files, "out:scripts", map[string]bool{
		"Skip":                false,
		"NoWrite":             false,
		"DFHackCommandScript": true,
		"LangLua":             true,
	}, "--", ".dfcom.lua", ".lua", false)
	if err != nil {
		return err
	}
	err = state.writeFiles(state.Files, "out:scripts", map[string]bool{
		"Skip":                false,
		"NoWrite":             false,
		"DFHackCommandScript": true,
		"LangRuby":            true,
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
	return nil
}
