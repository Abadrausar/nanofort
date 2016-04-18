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

import "dctech/lua"

import "strings"

import "math/rand"

import "dctech/axis"

import "rubble7"
import "rubble7/rblutil"
import "rubble7/rblutil/addon"
import "rubble7/rblutil/errors"
import "rubble7/rblutil/merge"
import "rubble7/rblutil/parse"
import "rubble7/rblutil/rparse"

func rubbleLuaAPI(l *lua.State) int {
	l.NewTable(0, 32) // 16 functions + 9 variables (the Rubble standard library then adds a few)
	tidx := l.AbsIndex(-1)
	
	l.SetTableFunctions(tidx, rubbleLibrary)
	
	l.Push("version")
	l.Push(rubble7.Version)
	l.SetTableRaw(tidx)
	
	l.Push("vmajor")
	l.Push(int64(rubble7.VMajor))
	l.SetTableRaw(tidx)
	
	l.Push("vminor")
	l.Push(int64(rubble7.VMinor))
	l.SetTableRaw(tidx)
	
	l.Push("vpatch")
	l.Push(int64(rubble7.VPatch))
	l.SetTableRaw(tidx)

	l.Push("files")
	l.Push(&struct{}{})
	rubbleStateFilesTbl(l)
	l.SetMetaTable(-2)
	l.SetTableRaw(tidx)

	l.Push("gfiles")
	l.Push(&struct{}{})
	rubbleStateGlobalFilesTbl(l)
	l.SetMetaTable(-2)
	l.SetTableRaw(tidx)

	l.Push("addons")
	l.Push(&struct{}{})
	rubbleStateAddonsTbl(l)
	l.SetMetaTable(-2)
	l.SetTableRaw(tidx)

	l.Push("addonstbl")
	l.Push(&struct{}{})
	rubbleStateAddonsTblTbl(l)
	l.SetMetaTable(-2)
	l.SetTableRaw(tidx)

	l.Push("registry")
	l.Push(&struct{}{})
	rubbleRegistryTbl(l)
	l.SetMetaTable(-2)
	l.SetTableRaw(tidx)

	return 1
}

var rubbleLibrary = map[string]lua.NativeFunction{
	// Writes its arguments to the log. No spaces or new lines are added.
	"print": func(l *lua.State) int {
		state := getState(l)

		top := l.AbsIndex(-1)
		for i := 1; i <= top; i++ {
			state.Log.Print(l.ToString(i))
		}
		return 0
	},

	"random": func(l *lua.State) int {
		l.Push(rand.New(rand.NewSource(1)))
		rubbleRandomTbl(l)
		l.SetMetaTable(-2)
		return 1
	},

	"warning": func(l *lua.State) int {
		state := getState(l)

		top := l.AbsIndex(-1)
		msg := ""
		for i := 1; i <= top; i++ {
			msg += l.ToString(i)
		}
		state.Log.Warn(msg)
		return 0
	},

	"abort": func(l *lua.State) int {
		errors.RaiseAbort(l.ToString(1))
		return 0
	},

	"error": func(l *lua.State) int {
		errors.RaiseError(l.ToString(1))
		return 0
	},

	"currentfile": func(l *lua.State) int {
		state := getState(l)

		l.Push(state.CurrentFile)
		return 1
	},

	"configvar": func(l *lua.State) int {
		state := getState(l)

		name := l.ToString(1)
		if l.AbsIndex(-1) < 2 {
			l.Push(state.VariableData[name])
			return 1
		}
		state.VariableData[name] = l.ToString(2)
		return 0
	},

	// TODO: See about supporting custom dispatchers.
	"parse": func(l *lua.State) int {
		state := getState(l)

		raws := l.ToString(1)
		stage := parse.Stage(l.OptInt(2, -1))
		if stage == -1 {
			stage = state.Stage
		}

		raws = string(parse.Parse([]byte(raws), "lua.rubble.parse", 1, stage, state.Dispatcher))
		l.Push(string(raws))
		return 1
	},

	"execscript": func(l *lua.State) int {
		state := getState(l)

		code := l.ToString(1)
		tag := l.OptString(2, "LangLua")
		rtn, err := state.Script.RunScript([]byte(code), "rubble.execscript", 1, []string{}, tag)
		if err != nil {
			l.Push(false)
			l.Push(err.Error())
			return 2
		}
		l.Push(true)
		l.Push(string(rtn))
		return 2
	},

	"calltemplate": func(l *lua.State) int {
		state := getState(l)

		n := l.AbsIndex(-1)
		if n == 0 {
			l.Push("No arguments in call to calltemplate")
			l.Error()
		}

		args := make([]string, 0, n)
		for i := 1; i <= n; i++ {
			args = append(args, l.ToString(i))
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

		l.Push(string(state.Dispatcher(args, stage, nil, nil)))
		return 1
	},

	"expandvars": func(l *lua.State) int {
		state := getState(l)

		raws := l.ToString(1)
		openchar := l.OptString(2, "")
		if len(openchar) == 0 {
			openchar = "$"
		}
		nest := l.ToBool(3)

		data := state.VariableData
		if l.AbsIndex(-1) >= 4 {
			data = make(map[string]string)
			
			l.ForEachInTable(4, func() {
				data[l.ToString(-2)] = l.ToString(-1)
			})
		}

		l.Push(string(rblutil.Expand(raws, openchar[0], nest, data)))
		return 1
	},

	"template": func(l *lua.State) int {
		state := getState(l)

		name := l.ToString(1)

		f := 2
		if l.TypeOf(2) != lua.TypFunction {
			err := l.LoadText(strings.NewReader(l.ToString(2)), name, 0)
			if err != nil {
				errors.RaiseWrappedError(name, err)
			}
			f = -1
		}
		
		state.Templates[name] = &rubble7.Template{
			Code: string(l.DumpFunction(f, false)),
			Tag:  "LangLuaBin",
		}
		
		return 0
	},

	"scripttemplate": func(l *lua.State) int {
		state := getState(l)

		state.Templates[l.ToString(1)] = &rubble7.Template{
			Code: l.ToString(3),
			Tag:  l.ToString(2),
		}
		return 0
	},

	"usertemplate": func(l *lua.State) int {
		state := getState(l)

		names, values := []string{}, []string{}
		l.ForEachInTable(2, func() {
			l.Push(int64(1))
			l.GetTable(-2)
			names = append(names, strings.TrimSpace(l.ToString(-1)))
			l.Pop(1)
			
			l.Push(int64(2))
			l.GetTable(-2)
			values = append(values, strings.TrimSpace(l.ToString(-1)))
			l.Pop(1)
		})

		state.Templates[l.ToString(1)] = &rubble7.Template{
			Code: l.ToString(3),
			Tag:  "LangTemplate",

			ArgNames:    names,
			ArgDefaults: values,
		}
		return 0
	},

	"fileaction": func(l *lua.State) int {
		state := getState(l)

		filter := make(map[string]bool)
		l.ForEachInTable(1, func() {
			filter[l.ToString(-2)] = l.ToBool(-1)
		})

		state.Files.RunAction(filter, func(file *addon.File) error {
			l.PushIndex(2)
			l.Push(file)
			addonFileTbl(l)
			l.SetMetaTable(-2)
			l.Call(1, 0)
			return nil
		})
		return 0
	},
	
	"rawmerge": func(l *lua.State) int {
		source := rparse.ParseRaws([]byte(l.ToString(2)))
		dest := rparse.ParseRaws([]byte(l.ToString(3)))
		
		rules := new(merge.RuleNode)
		merge.ParseRules([]byte(l.ToString(1)), rules)
		
		stree := merge.MakeTree(source, rules)
		
		merge.Apply(dest, stree)
		
		l.Push(string(rparse.FormatFile(dest)))
		l.Push(stree.String())
		return 2
	},
	
	"auxfile": func(l *lua.State) int {
		state := getState(l)

		set := false
		if l.AbsIndex(-1) >= 2 {
			set = true
		}
		
		switch l.ToString(1) {
		case "init.txt":
			if set {
				state.Init = []byte(l.ToString(2))
				return 0
			}
			l.Push(string(state.Init))
			return 1
		case "d_init.txt":
			if set {
				state.D_Init = []byte(l.ToString(2))
				return 0
			}
			l.Push(string(state.D_Init))
			return 1
		}
		return 0
	},
}

func axisLuaAPI(l *lua.State) int {
	l.NewTable(0, 8) // 7 functions
	l.SetTableFunctions(-1, axisLibrary)
	return 1
}

var axisLibrary = map[string]lua.NativeFunction{
	"read": func(l *lua.State) int {
		state := getState(l)

		file, err := axis.ReadAll(state.FS, l.ToString(1))
		if err != nil {
			l.Push(false)
			l.Push(err.Error())
			return 2
		}
		l.Push(true)
		l.Push(string(file))
		return 2
	},

	"write": func(l *lua.State) int {
		state := getState(l)

		path := l.ToString(1)
		err := axis.Create(state.FS, path)
		if err != nil {
			l.Push(false)
			l.Push(err.Error())
			return 2
		}
		err = axis.WriteAll(state.FS, path, []byte(l.ToString(2)))
		if err != nil {
			l.Push(false)
			l.Push(err.Error())
			return 2
		}
		l.Push(true)
		return 1
	},

	"exists": func(l *lua.State) int {
		state := getState(l)
		l.Push(axis.Exists(state.FS, l.ToString(1)))
		return 1
	},

	"isdir": func(l *lua.State) int {
		state := getState(l)
		l.Push(axis.IsDir(state.FS, l.ToString(1)))
		return 1
	},

	"del": func(l *lua.State) int {
		state := getState(l)

		err := axis.Delete(state.FS, l.ToString(1))
		if err != nil {
			panic(err)
			l.Push(false)
			l.Push(err.Error())
			return 2
		}
		l.Push(true)
		return 1
	},

	"listdirs": func(l *lua.State) int {
		state := getState(l)

		files := axis.ListDir(state.FS, l.ToString(1))
		l.NewTable(len(files), 0)
		for i := range files {
			l.Push(int64(i+1))
			l.Push(files[i])
			l.SetTableRaw(-3)
		}
		return 1
	},

	"listfiles": func(l *lua.State) int {
		state := getState(l)

		files := axis.ListFile(state.FS, l.ToString(1))
		l.NewTable(len(files), 0)
		for i := range files {
			l.Push(int64(i+1))
			l.Push(files[i])
			l.SetTableRaw(-3)
		}
		return 1
	},
}
