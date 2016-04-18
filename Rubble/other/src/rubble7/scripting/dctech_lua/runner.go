/*
Copyright 2015-2016 by Milo Christiansen

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

// Lua scripting for Rubble.
//
// This uses "dctech/lua" as the runtime.
//
// To use Lua with Rubble simply use `import _ "rubble7/scripting/dctech_lua"`.
// Lua scripting will then be available to all Rubble States by default.
package lua

import "dctech/axis"

import "dctech/lua"
import "dctech/lua/lmodbase"
import "dctech/lua/lmodpackage"
import "dctech/lua/lmodstring"
import "dctech/lua/lmodtable"
import "dctech/lua/lmodmath"

import "rubble7"
import "rubble7/rblutil"
import "rubble7/rblutil/actions"

import "bytes"
import "strconv"

// LuaRunner is a Runner for Lua scripts.
type LuaRunner struct {
	l *lua.State
}

// New creates a new LuaRunner.
// You do not normally need to call this, as rubble7.NewScriptCore will automatically do so.
func New(state *rubble7.State) rubble7.Runner {
	runner := &LuaRunner{
		l: lua.NewState(),
	}
	l := runner.l
	l.Output = state.Log
	
	l.NativeTrace = true
	
	// Load standard modules
	l.Push(lmodbase.Open)
	l.Call(0, 0)
	l.Push(lmodpackage.Open)
	l.Call(0, 0)
	l.Push(lmodstring.Open)
	l.Call(0, 0)
	l.Push(lmodtable.Open)
	l.Call(0, 0)
	l.Push(lmodmath.Open)
	l.Call(0, 0)

	// Save a state reference to the registry
	l.Push("RUBBLE_STATE")
	l.Push(state)
	l.SetTableRaw(lua.RegistryIndex)
	
	// And finally preload the Rubble API
	l.Preload("rubble", rubbleLuaAPI)
	l.Preload("axis", axisLuaAPI)
	l.Preload("rubble.rparse", rparseLuaAPI)

	// Isn't it nice to not have to work around third party bugs?
	return runner
}

// Get the Rubble state from the registry.
func getState(l *lua.State) *rubble7.State {
	l.Push("RUBBLE_STATE")
	l.GetTableRaw(lua.RegistryIndex)
	state := l.ToUser(-1).(*rubble7.State)
	l.Pop(1)
	return state
}

func init() {
	// We don't need to add a new language tag as Lua already has a tag (it is used by DFHack).
	// If we did need a new tag we would have to import "rubble7/rblutil/addon", then do the following:
	//	addon.DefaultLast[".lua"] = []string{"LangLua"}
	rubble7.AddDefaultRunner([]string{"LangLua", "LangLuaC", "LangLuaBin"}, New)
	
	// This enables a compile script action for the universal interface.
	actions.RegisterFunc("luac", func(log *rblutil.Logger, rblDir, dfDir, outDir string, addonDirs []string, options []interface{}) bool {
		inFile, outFile := *options[0].(*string), *options[1].(*string)
		
		log.Println("Compiling Script:", inFile)
		log.Println("  Initializing AXIS VFS...")
		fs := rubble7.InitAXIS(rblDir, dfDir, outDir, addonDirs)
		
		log.Println("  Loading Script File...")
		script, err := axis.ReadAll(fs, inFile)
		if err != nil {
			log.Println(err)
			return false
		}
		
		l := lua.NewState()
		
		log.Println("  Compiling...")
		err = l.LoadText(bytes.NewBuffer(script), inFile, 0)
		if err != nil {
			log.Println(err)
			return false
		}
		
		log.Println("  Writing Binary...")
		script = l.DumpFunction(-1, false)
		
		if !axis.Exists(fs, outFile) {
			err = axis.Create(fs, outFile)
			if err != nil {
				log.Println(err)
				return false
			}
		}
		err = axis.WriteAll(fs, outFile, script)
		if err != nil {
			log.Println(err)
			return false
		}
		
		return true
	}, "Compile a Lua script with the same compiler Rubble uses when generating.\nAll paths are AXIS paths. The AXIS filesystem has the same structure as normal.", []actions.Option{
		{
			Name: "script",
			Help: "The input `file`.",
			Flag: false,
			Multiple: false,
		},
		{
			Name: "out",
			Help: "The output `file`.",
			Flag: false,
			Multiple: false,
		},
	})
	actions.RegisterFunc("cluac", func(log *rblutil.Logger, rblDir, dfDir, outDir string, addonDirs []string, options []interface{}) bool {
		inFile, outFile := *options[0].(*string), *options[1].(*string)
		
		log.Println("Compiling Script:", inFile)
		log.Println("  Initializing AXIS VFS...")
		fs := rubble7.InitAXIS(rblDir, dfDir, outDir, addonDirs)
		
		log.Println("  Loading Script File...")
		script, err := axis.ReadAll(fs, inFile)
		if err != nil {
			log.Println(err)
			return false
		}
		
		l := lua.NewState()
		
		log.Println("  Compiling...")
		err = l.LoadTextC(bytes.NewBuffer(script), inFile, 0)
		if err != nil {
			log.Println(err)
			return false
		}
		
		log.Println("  Writing Binary...")
		script = l.DumpFunction(-1, false)
		
		if !axis.Exists(fs, outFile) {
			err = axis.Create(fs, outFile)
			if err != nil {
				log.Println(err)
				return false
			}
		}
		err = axis.WriteAll(fs, outFile, script)
		if err != nil {
			log.Println(err)
			return false
		}
		
		return true
	}, "Compile a CLua script with the same compiler Rubble uses when generating.\nAll paths are AXIS paths. The AXIS filesystem has the same structure as normal.", []actions.Option{
		{
			Name: "script",
			Help: "The input `file`.",
			Flag: false,
			Multiple: false,
		},
		{
			Name: "out",
			Help: "The output `file`.",
			Flag: false,
			Multiple: false,
		},
	})
}

func (runner *LuaRunner) Run(script []byte, tag string, name string, line int, args []string) (string, error) {
	l := runner.l
	if line > 1 {
		name += "@" + strconv.Itoa(line)
	}

	switch tag {
	case "LangLuaBin":
		err := l.LoadBinary(bytes.NewBuffer(script), name, 0)
		if err != nil {
			return "", err
		}
	case "LangLuaC":
		err := l.LoadTextC(bytes.NewBuffer(script), name, 0)
		if err != nil {
			return "", err
		}
	default:
		err := l.LoadText(bytes.NewBuffer(script), name, 0)
		if err != nil {
			return "", err
		}
	}
	
	for _, v := range args {
		l.Push(v)
	}

	err := l.PCall(len(args), 1)
	if err != nil {
		return "", err
	}
	
	if l.IsNil(-1) {
		l.Pop(1)
		return "", nil
	}
	rtn := l.ToString(-1)
	l.Pop(1)
	return rtn, nil
}
