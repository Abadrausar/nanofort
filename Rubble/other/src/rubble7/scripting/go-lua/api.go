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

package lua

import "github.com/Shopify/go-lua"

import "strings"

import "dctech/axis"

import "rubble7"
import "rubble7/rblutil"
import "rubble7/rblutil/addon"
import "rubble7/rblutil/errors"
import "rubble7/rblutil/parse"

func rubbleLuaAPI(l *lua.State) int {
	lua.NewLibrary(l, rubbleLibrary)
	
	// The value does not matter, I just need something to attach the metatable to.
	l.PushUserData(0)
	rubbleFilesTable(l) // From metatables.go
	l.SetMetaTable(-2)
	l.SetField(-2, "files")
	
	l.PushUserData(0)
	rubbleAddonsTable(l) // From metatables.go
	l.SetMetaTable(-2)
	l.SetField(-2, "addons")
	
	l.PushString(rubble7.Versions[0])
	l.SetField(-2, "version")
	
	return 1
}

// TODO: I need to expose more of the internal data structures, in particular the file and addon lists,
// but also things like the version number.
var rubbleLibrary = []lua.RegistryFunction{
	// Writes its arguments to the log. No spaces or new lines are added.
	{"print", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		n := l.Top()
		l.Global("tostring")
		for i := 1; i <= n; i++ {
			l.PushValue(-1) // function to be called
			l.PushValue(i)  // value to print
			l.Call(1, 1)
			s, ok := l.ToString(-1)
			if !ok {
				lua.Errorf(l, "'tostring' must return a string to 'rubble.print'")
				panic("unreachable")
			}
			state.Log.Print(s)
			l.Pop(1) // pop result
		}
		return 0
	}},
	
	{"abort", func(l *lua.State) int {
		errors.RaiseAbort(lua.CheckString(l, 1))
		return 0
	}},
	
	{"currentfile", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		l.PushString(state.CurrentFile)
		return 1
	}},
	
	{"configvar", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		name := lua.OptString(l, 1, "")
		value := lua.OptString(l, 2, "")
		if l.IsNone(2) {
			l.PushString(state.VariableData[name])
			return 1
		}
		state.VariableData[name] = value
		return 0
	}},
	
	// This was going to support custom dispatchers, but that proved to be too hard.
	// It is possible, but would be a real pain in the neck. Maybe later.
	{"parse", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		raws := lua.OptString(l, 1, "")
		stage := parse.Stage(lua.OptNumber(l, 2, -1))
		if stage == -1 {
			stage = state.Stage
		}
		
		raws = string(parse.Parse([]byte(raws), "lua.rubble.parse", 1, stage, state.Dispatcher))
		l.PushString(raws)
		return 1
	}},
	
	{"execscript", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		code := lua.OptString(l, 1, "")
		tag := lua.OptString(l, 2, "LangLua")
		rtn, err := state.Script.RunScript([]byte(code), "rubble.execscript", 1, []string{}, tag)
		if err != nil {
			l.PushBoolean(false)
			l.PushString(err.Error())
			return 2
		}
		l.PushBoolean(true)
		l.PushString(rtn)
		return 2
	}},
	
	{"calltemplate", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		n := l.Top() // Number of arguments.
		if n == 0 {
			l.PushString("No arguments")
			l.Error()
		}
		
		args := make([]string, 0, n)
		for i := 1; i <= n; i++ {
			a, ok := l.ToString(i)
			if !ok {
				l.PushString("Bad argument: Not a string")
				l.Error()
			}
			args = append(args, a)
		}
		
		stage := state.Stage
		switch args[0][0] {
		case '!':
			stage = parse.StgPreParse
		case '#':
			stage = parse.StgPreParse
		case '@':
			// Not sure what to do here...
		default:
			stage = parse.StgParse
		}
		
		l.PushString(state.Dispatcher(args, stage, nil, nil))
		return 1
	}},
	
	{"expandvars", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		raws := lua.OptString(l, 1, "")
		openchar := lua.OptString(l, 2, "")
		if len(openchar) == 0 {
			openchar = "$"
		}
		nest := l.ToBoolean(3)
		
		data := state.VariableData
		if l.IsTable(4) {
			data = make(map[string]string)
			l.PushNil()
			for l.Next(4) {
				key := lua.CheckString(l, -2)
				val := lua.CheckString(l, -1)
				data[key] = val
				l.Pop(1)
			}
		}
		
		l.PushString(rblutil.Expand(raws, openchar[0], nest, data))
		return 1
	}},
	
	{"template", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		name := lua.OptString(l, 1, "")
		code := lua.OptString(l, 2, "")
		
		// Make sure the body compiles.
		err := lua.LoadString(l, code)
		if err != nil {
			errors.RaiseWrappedError(name, err)
		}
		l.Pop(1)
		
		state.Templates[name] = &rubble7.Template{
			Code: code,
			Tag: "LangLua",
		}
		return 0
	}},
	
	{"scripttemplate", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		name := lua.OptString(l, 1, "")
		tag := lua.OptString(l, 2, "")
		code := lua.OptString(l, 3, "")
		
		state.Templates[name] = &rubble7.Template{
			Code: code,
			Tag: tag,
		}
		return 0
	}},
	
	{"usertemplate", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		name := lua.OptString(l, 1, "")
		code := lua.OptString(l, 3, "")
		
		names, values := []string{}, []string{}
		if l.IsTable(2) {
			l.PushNil()
			for l.Next(2) {
				l.PushUnsigned(1)
				l.Table(-2)
				l.PushUnsigned(2)
				l.Table(-3)
				
				names = append(names, strings.TrimSpace(lua.CheckString(l, -1)))
				values = append(values, strings.TrimSpace(lua.CheckString(l, -2)))
				l.Pop(3)
			}
		}
		
		state.Templates[name] = &rubble7.Template{
			Code: code,
			Tag: "LangTemplate",
			
			ArgNames: names,
			ArgDefaults: values,
		}
		return 0
	}},
	
	{"fileaction", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		if !(l.IsTable(1) && l.IsFunction(2)) {
			l.PushString("Bad arguments: Needs a table and a function")
			l.Error()
		}
		
		filter := make(map[string]bool)
		l.PushNil()
		for l.Next(1) {
			key := lua.CheckString(l, -2)
			val := l.ToBoolean(-1)
			filter[key] = val
			l.Pop(1)
		}
		
		state.Files.RunAction(filter, func(file *addon.File) error {
			l.PushValue(2)
			l.PushUserData(file)
			fileTable(l) // From metatables.go
			l.SetMetaTable(-2)
			l.Call(1, 0)
			return nil
		})
		return 0
	}},
	
	//{"sortlists", func(l *lua.State) int {
	//	l.Field(lua.RegistryIndex, "RUBBLE_STATE")
	//	state := l.ToUserData(-1).(*rubble7.State)
	//	l.Pop(1)
	//	
	//	state.Files.Sort()
	//	state.GlobalFiles.Sort()
	//	return 0
	//}},
}


func axisLuaAPI(l *lua.State) int {
	lua.NewLibrary(l, axisLibrary)
	return 1
}

var axisLibrary = []lua.RegistryFunction{
	{"read", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		path := lua.OptString(l, 1, "")
		
		file, err := axis.ReadAll(state.FS, path)
		if err != nil {
			l.PushBoolean(false)
			l.PushString(err.Error())
			return 2
		}
		l.PushBoolean(true)
		l.PushString(string(file))
		return 2
	}},
	
	{"write", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		path := lua.OptString(l, 1, "")
		data := lua.OptString(l, 2, "")
		
		err := axis.Create(state.FS, path)
		if err != nil {
			l.PushBoolean(false)
			l.PushString(err.Error())
			return 2
		}
		err = axis.WriteAll(state.FS, path, []byte(data))
		if err != nil {
			l.PushBoolean(false)
			l.PushString(err.Error())
			return 2
		}
		l.PushBoolean(true)
		return 1
	}},
	
	{"exists", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		path := lua.OptString(l, 1, "")
		
		l.PushBoolean(axis.Exists(state.FS, path))
		return 1
	}},
	
	{"isdir", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		path := lua.OptString(l, 1, "")
		
		l.PushBoolean(axis.IsDir(state.FS, path))
		return 1
	}},
	
	{"del", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		path := lua.OptString(l, 1, "")
		
		err := axis.Delete(state.FS, path)
		if err != nil {
			l.PushBoolean(false)
			l.PushString(err.Error())
			return 2
		}
		l.PushBoolean(true)
		return 1
	}},
	
	{"listdirs", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		path := lua.OptString(l, 1, "")
		
		files := axis.ListDir(state.FS, path)
		l.CreateTable(len(files), 0)
		for i := range files {
			l.PushString(files[i])
			l.RawSetInt(-2, i + 1)
		}
		return 1
	}},
	
	{"listfiles", func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		path := lua.OptString(l, 1, "")
		
		files := axis.ListFile(state.FS, path)
		l.CreateTable(len(files), 0)
		for i := range files {
			l.PushString(files[i])
			l.RawSetInt(-2, i + 1)
		}
		return 1
	}},
}

var stringsLibrary = []lua.RegistryFunction{
	{"count", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		sep := lua.OptString(l, 2, "")
		l.PushInteger(strings.Count(str, sep))
		return 1
	}},
	
	{"hasprefix", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		a := lua.OptString(l, 2, "")
		l.PushBoolean(strings.HasPrefix(str, a))
		return 1
	}},
	
	{"hassuffix", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		a := lua.OptString(l, 2, "")
		l.PushBoolean(strings.HasSuffix(str, a))
		return 1
	}},
	
	{"join", func(l *lua.State) int {
		if !l.IsTable(1) {
			l.PushString("Bad arguments: Needs a table")
			l.Error()
		}
		
		set := make([]string, 0)
		l.PushNil()
		for l.Next(1) {
			val := lua.CheckString(l, -1)
			set = append(set, val)
			l.Pop(1)
		}
		sep := lua.OptString(l, 2, ", ")
		
		l.PushString(strings.Join(set, sep))
		return 1
	}},
	
	{"replace", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		o := lua.OptString(l, 2, "")
		n := lua.OptString(l, 3, "")
		i := lua.OptInteger(l, 4, -1)
		l.PushString(strings.Replace(str, o, n, i))
		return 1
	}},
	
	{"split", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		sep := lua.OptString(l, 2, "")
		n := lua.OptInteger(l, 3, -1)
		
		result := strings.SplitN(str, sep, n)
		l.CreateTable(len(result), 0)
		for i := range result {
			l.PushString(result[i])
			l.RawSetInt(-2, i + 1)
		}
		return 1
	}},
	
	{"splitafter", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		sep := lua.OptString(l, 2, "")
		n := lua.OptInteger(l, 3, -1)
		
		result := strings.SplitAfterN(str, sep, n)
		l.CreateTable(len(result), 0)
		for i := range result {
			l.PushString(result[i])
			l.RawSetInt(-2, i + 1)
		}
		return 1
	}},
	
	{"title", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		l.PushString(strings.Title(str))
		return 1
	}},
	
	{"trim", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		l.PushString(strings.TrimSpace(str))
		return 1
	}},
	
	{"trimprefix", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		a := lua.OptString(l, 2, "")
		l.PushString(strings.TrimPrefix(str, a))
		return 1
	}},
	
	{"trimspace", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		l.PushString(strings.TrimSpace(str))
		return 1
	}},
	
	{"trimsuffix", func(l *lua.State) int {
		str := lua.OptString(l, 1, "")
		a := lua.OptString(l, 2, "")
		l.PushString(strings.TrimSuffix(str, a))
		return 1
	}},
}
