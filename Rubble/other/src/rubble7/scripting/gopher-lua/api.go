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

import "github.com/yuin/gopher-lua"

import "strings"
import "strconv"

import "math/rand"

import "dctech/axis"

import "rubble7"
import "rubble7/rblutil"
import "rubble7/rblutil/addon"
import "rubble7/rblutil/errors"
import "rubble7/rblutil/merge"
import "rubble7/rblutil/parse"
import "rubble7/rblutil/rparse"

func rubbleLuaAPI(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), rubbleLibrary)

	l.SetField(mod, "version", lua.LString(rubble7.Version))
	l.SetField(mod, "vmajor", lua.LNumber(rubble7.VMajor))
	l.SetField(mod, "vminor", lua.LNumber(rubble7.VMinor))
	l.SetField(mod, "vpatch", lua.LNumber(rubble7.VPatch))

	l.SetField(mod, "files", &lua.LUserData{
		Metatable: rubbleStateFilesTbl(l),
	})

	l.SetField(mod, "gfiles", &lua.LUserData{
		Metatable: rubbleStateGlobalFilesTbl(l),
	})

	l.SetField(mod, "addons", &lua.LUserData{
		Metatable: rubbleStateAddonsTbl(l),
	})

	l.SetField(mod, "addonstbl", &lua.LUserData{
		Metatable: rubbleStateAddonsTblTbl(l),
	})

	l.SetField(mod, "registry", &lua.LUserData{
		Metatable: rubbleRegistryTbl(l),
	})

	l.Push(mod)
	return 1
}

var rubbleLibrary = map[string]lua.LGFunction{
	// Writes its arguments to the log. No spaces or new lines are added.
	"print": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		top := l.GetTop()
		for i := 1; i <= top; i++ {
			state.Log.Print(l.ToStringMeta(l.Get(i)).String())
		}
		return 0
	},

	"random": func(l *lua.LState) int {
		l.Push(&lua.LUserData{
			Value:     rand.New(rand.NewSource(1)),
			Metatable: rubbleRandomTbl(l),
		})
		return 1
	},

	"warning": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		top := l.GetTop()
		msg := ""
		for i := 1; i <= top; i++ {
			msg += l.ToStringMeta(l.Get(i)).String()
		}
		state.Log.Warn(msg) // Only call once to avoid messing up the warning count.
		return 0
	},
	
	"abort": func(l *lua.LState) int {
		errors.RaiseAbort(l.CheckString(1))
		return 0
	},

	"error": func(l *lua.LState) int {
		errors.RaiseError(l.CheckString(1))
		return 0
	},

	"currentfile": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		l.Push(lua.LString(state.CurrentFile))
		return 1
	},

	"configvar": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		name := l.OptString(1, "")
		if l.GetTop() < 2 {
			l.Push(lua.LString(state.VariableData[name]))
			return 1
		}
		state.VariableData[name] = l.OptString(2, "")
		return 0
	},

	// TODO: See about supporting custom dispatchers.
	"parse": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		raws := l.OptString(1, "")
		stage := parse.Stage(l.OptInt(2, -1))
		if stage == -1 {
			stage = state.Stage
		}

		raws = string(parse.Parse([]byte(raws), "lua.rubble.parse", 1, stage, state.Dispatcher))
		l.Push(lua.LString(raws))
		return 1
	},

	"execscript": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		code := l.OptString(1, "")
		tag := l.OptString(2, "LangLua")
		rtn, err := state.Script.RunScript([]byte(code), "rubble.execscript", 1, []string{}, tag)
		if err != nil {
			l.Push(lua.LFalse)
			l.Push(lua.LString(err.Error()))
			return 2
		}
		l.Push(lua.LTrue)
		l.Push(lua.LString(rtn))
		return 2
	},

	"calltemplate": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		n := l.GetTop()
		if n == 0 {
			l.Error(lua.LString("No arguments"), 1)
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

		l.Push(lua.LString(state.Dispatcher(args, stage, nil, nil)))
		return 1
	},

	"expandvars": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		raws := l.OptString(1, "")
		openchar := l.OptString(2, "")
		if len(openchar) == 0 {
			openchar = "$"
		}
		nest := l.OptBool(3, false)

		data := state.VariableData
		if l.GetTop() >= 4 {
			data = make(map[string]string)
			tbl := l.CheckTable(4)
			l.ForEach(tbl, func(key, val lua.LValue) {
				data[key.String()] = val.String()
			})
		}

		l.Push(lua.LString(rblutil.Expand(raws, openchar[0], nest, data)))
		return 1
	},

	"template": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		name := l.OptString(1, "")
		code := l.OptString(2, "")

		// Make sure the body compiles.
		// TODO: It should be possible to save the compiled form for later.
		_, err := l.LoadString(code)
		if err != nil {
			errors.RaiseWrappedError(name, err)
		}

		state.Templates[name] = &rubble7.Template{
			Code: code,
			Tag:  "LangLua",
		}
		return 0
	},

	"scripttemplate": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		name := l.OptString(1, "")
		tag := l.OptString(2, "")
		code := l.OptString(3, "")

		state.Templates[name] = &rubble7.Template{
			Code: code,
			Tag:  tag,
		}
		return 0
	},

	"usertemplate": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		name := l.OptString(1, "")

		names, values := []string{}, []string{}
		tbl := l.CheckTable(2)
		l.ForEach(tbl, func(key, val lua.LValue) {
			arg, ok := val.(*lua.LTable)
			if !ok {
				l.Error(lua.LString("Non-table item in args table."), 1)
			}

			names = append(names, strings.TrimSpace(l.GetTable(arg, lua.LNumber(1)).String()))
			values = append(values, strings.TrimSpace(l.GetTable(arg, lua.LNumber(2)).String()))
		})

		code := l.OptString(3, "")

		state.Templates[name] = &rubble7.Template{
			Code: code,
			Tag:  "LangTemplate",

			ArgNames:    names,
			ArgDefaults: values,
		}
		return 0
	},

	"fileaction": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		tbl := l.CheckTable(1)
		f := l.CheckFunction(2)

		filter := make(map[string]bool)
		l.ForEach(tbl, func(key, val lua.LValue) {
			filter[key.String()] = lua.LVAsBool(val)
		})

		state.Files.RunAction(filter, func(file *addon.File) error {
			l.Push(f)
			l.Push(&lua.LUserData{
				Value:     file,
				Metatable: addonFileTbl(l),
			})
			l.Call(1, 0)
			return nil
		})
		return 0
	},
	
	"rawmerge": func(l *lua.LState) int {
		source := rparse.ParseRaws([]byte(l.OptString(2, "")))
		dest := rparse.ParseRaws([]byte(l.OptString(3, "")))
		
		rules := new(merge.RuleNode)
		merge.ParseRules([]byte(l.OptString(1, "")), rules)
		
		stree := merge.MakeTree(source, rules)
		
		merge.Apply(dest, stree)
		
		l.Push(lua.LString(rparse.FormatFile(dest)))
		l.Push(lua.LString(stree.String()))
		return 2
	},
	
	"auxfile": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		name := l.OptString(1, "")
		content := l.OptString(2, "")
		
		set := false
		if l.GetTop() >= 2 {
			set = true
		}
		
		switch name {
		case "init.txt":
			if set {
				state.Init = []byte(content)
				return 0
			}
			l.Push(lua.LString(state.Init))
			return 1
		case "d_init.txt":
			if set {
				state.D_Init = []byte(content)
				return 0
			}
			l.Push(lua.LString(state.D_Init))
			return 1
		}
		return 0
	},
}

func axisLuaAPI(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), axisLibrary)

	l.Push(mod)
	return 1
}

var axisLibrary = map[string]lua.LGFunction{
	"read": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		path := l.OptString(1, "")

		file, err := axis.ReadAll(state.FS, path)
		if err != nil {
			l.Push(lua.LFalse)
			l.Push(lua.LString(err.Error()))
			return 2
		}
		l.Push(lua.LTrue)
		l.Push(lua.LString(file))
		return 2
	},

	"write": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		path := l.OptString(1, "")
		data := l.OptString(2, "")

		err := axis.Create(state.FS, path)
		if err != nil {
			l.Push(lua.LFalse)
			l.Push(lua.LString(err.Error()))
			return 2
		}
		err = axis.WriteAll(state.FS, path, []byte(data))
		if err != nil {
			l.Push(lua.LFalse)
			l.Push(lua.LString(err.Error()))
			return 2
		}
		l.Push(lua.LTrue)
		return 1
	},

	"exists": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		path := l.OptString(1, "")

		if axis.Exists(state.FS, path) {
			l.Push(lua.LTrue)
			return 1
		}
		l.Push(lua.LFalse)
		return 1
	},

	"isdir": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		path := l.OptString(1, "")

		if axis.IsDir(state.FS, path) {
			l.Push(lua.LTrue)
			return 1
		}
		l.Push(lua.LFalse)
		return 1
	},

	"del": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		path := l.OptString(1, "")

		err := axis.Delete(state.FS, path)
		if err != nil {
			l.Push(lua.LFalse)
			l.Push(lua.LString(err.Error()))
			return 2
		}
		l.Push(lua.LTrue)
		return 1
	},

	"listdirs": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		path := l.OptString(1, "")

		files := axis.ListDir(state.FS, path)
		tbl := l.CreateTable(len(files), 0)
		for i := range files {
			tbl.RawSetInt(i+1, lua.LString(files[i]))
		}
		l.Push(tbl)
		return 1
	},

	"listfiles": func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		path := l.OptString(1, "")

		files := axis.ListFile(state.FS, path)
		tbl := l.CreateTable(len(files), 0)
		for i := range files {
			tbl.RawSetInt(i+1, lua.LString(files[i]))
		}
		l.Push(tbl)
		return 1
	},
}

var stringLibrary = map[string]lua.LGFunction{
	"count": func(l *lua.LState) int {
		str := l.OptString(1, "")
		sep := l.OptString(2, "")
		l.Push(lua.LNumber(strings.Count(str, sep)))
		return 1
	},

	"hasprefix": func(l *lua.LState) int {
		str := l.OptString(1, "")
		a := l.OptString(2, "")
		if strings.HasPrefix(str, a) {
			l.Push(lua.LTrue)
			return 1
		}
		l.Push(lua.LFalse)
		return 1
	},

	"hassuffix": func(l *lua.LState) int {
		str := l.OptString(1, "")
		a := l.OptString(2, "")
		if strings.HasSuffix(str, a) {
			l.Push(lua.LTrue)
			return 1
		}
		l.Push(lua.LFalse)
		return 1
	},

	"join": func(l *lua.LState) int {
		tbl := l.CheckTable(1)

		set := make([]string, 0)
		l.ForEach(tbl, func(key, val lua.LValue) {
			set = append(set, val.String())
		})

		sep := l.OptString(2, ", ")

		l.Push(lua.LString(strings.Join(set, sep)))
		return 1
	},

	"replace": func(l *lua.LState) int {
		str := l.OptString(1, "")
		o := l.OptString(2, "")
		n := l.OptString(3, "")
		i := l.OptInt(4, -1)
		l.Push(lua.LString(strings.Replace(str, o, n, i)))
		return 1
	},

	"split": func(l *lua.LState) int {
		str := l.OptString(1, "")
		sep := l.OptString(2, "")
		n := l.OptInt(3, -1)

		result := strings.SplitN(str, sep, n)
		tbl := l.CreateTable(len(result), 0)
		for i := range result {
			tbl.RawSetInt(i+1, lua.LString(result[i]))
		}
		l.Push(tbl)
		return 1
	},

	"splitafter": func(l *lua.LState) int {
		str := l.OptString(1, "")
		sep := l.OptString(2, "")
		n := l.OptInt(3, -1)

		result := strings.SplitAfterN(str, sep, n)
		tbl := l.CreateTable(len(result), 0)
		for i := range result {
			tbl.RawSetInt(i+1, lua.LString(result[i]))
		}
		l.Push(tbl)
		return 1
	},

	"title": func(l *lua.LState) int {
		str := l.OptString(1, "")
		l.Push(lua.LString(strings.Title(str)))
		return 1
	},

	"trim": func(l *lua.LState) int {
		str := l.OptString(1, "")
		l.Push(lua.LString(strings.TrimSpace(str)))
		return 1
	},

	"trimprefix": func(l *lua.LState) int {
		str := l.OptString(1, "")
		a := l.OptString(2, "")
		l.Push(lua.LString(strings.TrimPrefix(str, a)))
		return 1
	},

	"trimspace": func(l *lua.LState) int {
		str := l.OptString(1, "")
		l.Push(lua.LString(strings.TrimSpace(str)))
		return 1
	},

	"trimsuffix": func(l *lua.LState) int {
		str := l.OptString(1, "")
		a := l.OptString(2, "")
		l.Push(lua.LString(strings.TrimSuffix(str, a)))
		return 1
	},

	"unquote": func(l *lua.LState) int {
		str := l.OptString(1, "")

		rtn, err := strconv.Unquote(str)
		if err != nil {
			l.Push(lua.LString(str))
			return 1
		}
		l.Push(lua.LString(rtn))
		return 1
	},
}
