// +build ignore

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
// This uses "github.com/Shopify/go-lua" as the runtime.
// 
// To use Lua with Rubble simply use `import _ "rubble7/scripting/go-lua"`.
// Lua scripting will then be available to all Rubble States by default.
package lua

import "github.com/Shopify/go-lua"
import "rubble7"

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
	
	runner.l.PushString("RUBBLE_STATE")
	runner.l.PushUserData(state) // Since go-lua uses the go heap for GC this should not be a problem.
	runner.l.RawSet(lua.RegistryIndex)
	
	libs := []lua.RegistryFunction{
		{"_G", lua.BaseOpen},
		{"package", lua.PackageOpen},
		// {"coroutine", lua.CoroutineOpen},
		{"table", lua.TableOpen},
		//{"io", lua.IOOpen},
		//{"os", lua.OSOpen},
		{"string", lua.StringOpen},
		{"bit32", lua.Bit32Open},
		{"math", lua.MathOpen},
		{"debug", lua.DebugOpen},
		
		{"rubble", rubbleLuaAPI},
		{"axis", axisLuaAPI},
	}
	for _, lib := range libs {
		lua.Require(runner.l, lib.Name, lib.Function, true)
		runner.l.Pop(1)
	}
	
	// The default Lua string library sucks, lets give it some extra functions that are actually useful.
	runner.l.Global("string")
	for _, r := range stringsLibrary {
		runner.l.PushGoFunction(r.Function)
		runner.l.SetField(-2, r.Name)
	}
	runner.l.Pop(1)
	
	return runner
}

func init() {
	// We don't need to add a new language tag as Lua already has a tag (it is used by DFHack).
	// If we did need a new tag we would have to import "rubble7/rblutil/addon", then do the following:
	//	addon.DefaultLast[".lua"] = []string{"LangLua"}
	rubble7.AddDefaultRunner([]string{"LangLua"}, New)
}

func (runner *LuaRunner) Run(script []byte, tag string, name string, line int, args []string) (string, error) {
	err := lua.LoadString(runner.l, string(script))
	if err != nil {
		return "", err
	}
	
	for _, v := range args {
		runner.l.PushString(v)
	}
	
	err = runner.l.ProtectedCall(len(args), 1, 0)
	if err != nil {
		return "", err
	}
	
	if str, ok := runner.l.ToString(-1); ok {
		runner.l.Pop(1)
		return string(str), nil
	}
	return "", nil
}
