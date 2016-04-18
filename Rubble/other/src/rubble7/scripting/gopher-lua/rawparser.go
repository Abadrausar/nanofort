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
import "rubble7/rblutil/rparse"

import "bytes"

func rparseLuaAPI(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), rparseLibrary)

	l.Push(mod)
	return 1
}

var rparseLibrary = map[string]lua.LGFunction{
	"parse": func(l *lua.LState) int {
		tags := rparse.ParseRaws([]byte(l.OptString(1, "")))

		tbl := l.CreateTable(len(tags), 0)
		for i := range tags {
			tbl.RawSetInt(i+1, &lua.LUserData{
				Value:     tags[i],
				Metatable: rparseTagTbl(l),
			})
		}
		l.Push(tbl)
		return 1
	},
	"format": func(l *lua.LState) int {
		tbl := l.CheckTable(1)
		out := new(bytes.Buffer)

		l.ForEach(tbl, func(key, val lua.LValue) {
			val.(*lua.LUserData).Value.(*rparse.Tag).Format(out)
		})

		l.Push(lua.LString(out.String()))
		return 1
	},
	"newtag": func(l *lua.LState) int {
		l.Push(&lua.LUserData{
			Value:     new(rparse.Tag),
			Metatable: rparseTagTbl(l),
		})
		return 1
	},
}

func rparseTagTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		tag := l.ToUserData(1).Value.(*rparse.Tag)
		key := l.OptString(2, "")

		switch key {
		case "ID":
			l.Push(lua.LString(tag.ID))
			return 1
		case "Params":
			l.Push(&lua.LUserData{
				Value:     &tag.Params,
				Metatable: strsliceTbl(l, true),
			})
			return 1
		case "Comments":
			l.Push(lua.LString(tag.Comments))
			return 1
		case "CommentsOnly":
			if tag.CommentsOnly {
				l.Push(lua.LTrue)
				return 1
			}
			l.Push(lua.LFalse)
			return 1
		case "Line":
			l.Push(lua.LNumber(tag.Line))
			return 1
		}
		return 0
	}))

	tbl.RawSetString("__newindex", l.NewFunction(func(l *lua.LState) int {
		tag := l.ToUserData(1).Value.(*rparse.Tag)
		key := l.OptString(2, "")

		switch key {
		case "ID":
			tag.ID = l.OptString(3, "")
			return 0
		case "Params":
			tbl := l.CheckTable(3)
			tag.Params = make([]string, 0, tbl.Len())
			l.ForEach(tbl, func(key, val lua.LValue) {
				tag.Params = append(tag.Params, val.String())
			})
			return 0
		case "Comments":
			tag.Comments = l.OptString(3, "")
			return 0
		case "CommentsOnly":
			tag.CommentsOnly = l.OptBool(3, false)
			return 0
		case "Line":
			tag.Line = l.OptInt(3, 1)
			return 0
		}
		return 0
	}))

	return tbl
}
