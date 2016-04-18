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

import "dctech/axis"

import "rubble7/rblutil"
import "rubble7/rblutil/errors"
import "rubble7/rblutil/addon"

// TSetList returns a list of all non-lib tileset addons installed.
// Use when you do not have a State handy.
func TSetList(rbldir, dfdir string, addonsdir []string, log *rblutil.Logger) (err error, tilesets []string) {
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
		if meta.Tags["TileSet"] && !meta.Tags["Library"] {
			tilesets = append(tilesets, addon)
		}
	}
	return nil, tilesets
}

// TSetModeRun applies a tileset to the specified region.
func TSetModeRun(region, rbldir, dfdir string, addonsdir, addons []string, log *rblutil.Logger) (err error) {
	output := dfdir
	if region == "raw" {
		output += "/raw"
	} else {
		output += "/data/save/" + region + "/raw"
	}

	log.Separator()
	log.Println("Entering Tileset Mode for Region:", region)

	oops, state := NewState(rbldir, dfdir, output, addonsdir, log)
	if oops != nil {
		return oops
	}
	defer errors.TrapError(&err, state.Log)

	err = state.Load(addons, nil)
	if err != nil {
		return err
	}

	state.Files.Update(state.Addons)

	// This should keep any files from addons from messing things up.
	for _, file := range state.Files.Data {
		if !(file.Tags["TileSet"] && (file.Tags["ImagePNG"] || file.Tags["ImageBMP"] || file.Tags["TWBTOverride"])) {
			file.Tags["NoWrite"] = true
		}
		if !file.Tags["TileSet"] {
			file.Tags["RawFile"] = false
		}
	}

	rawFiles := []*addon.File{}
	for _, filepath := range axis.ListFile(state.FS, "out:objects") {
		content, err := axis.ReadAll(state.FS, "out:objects/"+filepath)
		if err != nil {
			return err
		}

		file := addon.NewFile(filepath, "out:objects", content)
		tags := addon.GetFileTags(filepath)
		for _, tag := range tags {
			file.Tags[tag] = true
		}

		// All the files should be raw files, but just in case...
		if file.Tags["RawFile"] && !file.Tags["NoWrite"] && !file.Tags["TileSet"] {
			rawFiles = append(rawFiles, file)
		}
	}
	state.Files.AddFiles(rawFiles...)

	err = state.ApplyTileSet()
	if err != nil {
		return err
	}

	// Simulate the important parts of state.Write, but without the file headers.
	state.Log.Separator()
	state.Log.Println("Writing Files...")
	state.Log.Println("  Writing Raw Files...")
	err = state.writeFiles(state.Files, "out:objects", map[string]bool{
		"Skip":             false,
		"NoWrite":          false,
		"AUXTextFile":      false,
		"CreatureGraphics": false,
		"TileSet":          false,
		"RawFile":          true,
	}, "", "", "", false) // No header, the files should already have one.
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
			file.Source + "/" + file.Name + "\n\n" + string(file.Content))...)
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
