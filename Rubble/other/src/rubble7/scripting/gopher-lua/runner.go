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
// This uses "github.com/yuin/gopher-lua" as the runtime.
//
// To use Lua with Rubble simply use `import _ "rubble7/scripting/gopher-lua"`.
// Lua scripting will then be available to all Rubble States by default.
package lua

import "github.com/yuin/gopher-lua"
import "rubble7"
import "bytes"
import "strconv"

func init() {
	lua.LuaPathDefault = ""
	lua.LuaLDir = ""
	lua.LuaOS = "rubble"
}

// LuaRunner is a Runner for Lua scripts.
type LuaRunner struct {
	l *lua.LState
}

// New creates a new LuaRunner.
// You do not normally need to call this, as rubble7.NewScriptCore will automatically do so.
func New(state *rubble7.State) rubble7.Runner {
	runner := &LuaRunner{
		l: lua.NewState(lua.Options{
			SkipOpenLibs: true,
		}),
	}
	l := runner.l

	l.RawSet(l.CheckTable(lua.RegistryIndex), lua.LString("RUBBLE_STATE"), &lua.LUserData{Value: state})

	libs := []struct{
		n lua.LString
		f lua.LGFunction
	}{
		{lua.LoadLibName, lua.OpenPackage},
		{lua.BaseLibName, lua.OpenBase},
		{lua.TabLibName, lua.OpenTable},
		//{lua.IoLibName, lua.OpenIo},
		//{lua.OsLibName, lua.OpenOs},
		{lua.StringLibName, lua.OpenString},
		{lua.MathLibName, lua.OpenMath},
		//{lua.DebugLibName, lua.OpenDebug},
		{lua.ChannelLibName, lua.OpenChannel},
		{lua.CoroutineLibName, lua.OpenCoroutine},
	}
	for _, lib := range libs {
		l.Push(l.NewFunction(lib.f))
		l.Push(lib.n)
		l.Call(1, 0)
	}

	// The default Lua string library sucks, let's give it some extra functions that are actually useful.
	tbl := runner.l.FindTable(runner.l.CheckTable(lua.RegistryIndex), "_LOADED", 1)
	strtbl := l.GetField(tbl, "string").(*lua.LTable)
	l.SetFuncs(strtbl, stringLibrary)

	l.PreloadModule("rubble", rubbleLuaAPI)
	l.PreloadModule("axis", axisLuaAPI)
	l.PreloadModule("rubble.rparse", rparseLuaAPI)

	// Kill the standard module loader.
	l.FindTable(l.CheckTable(lua.RegistryIndex), "_LOADERS", 1).(*lua.LTable).Remove(2)

	globals := l.Get(lua.GlobalsIndex).(*lua.LTable)

	// The default pairs and ipairs do not support meta tables, fix that.
	globals.RawSetString("ipairs", l.NewClosure(baseIpairs, l.NewFunction(ipairsaux)))
	globals.RawSetString("pairs", l.NewClosure(basePairs, l.NewFunction(pairsaux)))

	return runner
}

func init() {
	// We don't need to add a new language tag as Lua already has a tag (it is used by DFHack).
	// If we did need a new tag we would have to import "rubble7/rblutil/addon", then do the following:
	//	addon.DefaultLast[".lua"] = []string{"LangLua"}
	rubble7.AddDefaultRunner([]string{"LangLua"}, New)
}

func (runner *LuaRunner) Run(script []byte, tag string, name string, line int, args []string) (string, error) {
	if line > 1 {
		name += "@" + strconv.Itoa(line)
	}
	l := runner.l

	exe, err := l.Load(bytes.NewBuffer(script), name)
	if err != nil {
		return "", err
	}
	l.Push(exe)

	for _, v := range args {
		l.Push(lua.LString(v))
	}

	err = l.PCall(len(args), 1, nil)
	if err != nil {
		return "", err
	}

	rtn := l.ToString(-1)
	l.Pop(1)
	return rtn, nil
}
