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

// Rubble Stand-Alone Tileset Switcher
package main

import "os"
import "fmt"
import "runtime"
import "path/filepath"

import "rubble7"
import "rubble7/rblutil"

import _ "rubble7/scripting/gopher-lua"
import _ "rubble7/scripting/otto"

func main() {
	err, log := rblutil.NewLogger()
	if err != nil {
		fmt.Println("Fatal Error:", err)
		os.Exit(1)
	}

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

	path, err := filepath.Abs("./..")
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
	region := filepath.Base(path)
	log.Println("Switching region:", region)

	err, tilesets := rubble7.TSetList(".", "../../../..", []string{"rubble:tilesets"}, log)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}

	log.Separator()
	if len(tilesets) == 0 {
		log.Println("No Tilesets.")
		return
	}

	log.Println("Choose your tileset:")
	for i, tileset := range tilesets {
		log.Printf("    %v) %v\n", i+1, tileset)
	}

	addons := []string{""}
	index := -1
	for {
		log.Println("Type the number of your choice and press enter, type 0 to cancel:")
		_, err := fmt.Scanf("%d\n", &index)
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		if index < 1 {
			log.Println("Canceled.")
			return
		}

		if index > len(tilesets) {
			log.Println("That item does not exist!")
			continue
		}

		addons[0] = tilesets[index-1]
		break
	}

	err = rubble7.TSetModeRun(region, ".", "../../../..", []string{"rubble:tilesets"}, addons, log)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("Done.")
	return
}
