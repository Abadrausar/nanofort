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
import "rubble7"

var listAppend = func(l *lua.LState) int {
	regdata := l.ToUserData(1).Value.(rubble7.ScrRegData)
	
	*regdata.List = append(*regdata.List, l.ToString(2))
	return 0
}

func rubbleRegistryTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 1)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		key := l.OptString(2, "")
		if key == "exists" {
			l.Push(&lua.LUserData{
				Metatable: rubbleRegistryExistsTbl(l),
			})
			return 1
		}

		child, ok := state.ScrRegistry[key]
		if !ok {
			child = rubble7.ScrRegData{
				List:  &[]string{},
				Table: map[string]string{},
			}
			state.ScrRegistry[key] = child
		}

		l.Push(&lua.LUserData{
			Value:     child,
			Metatable: rubbleRegistryChildTbl(l),
		})
		return 1
	}))

	return tbl
}

func rubbleRegistryExistsTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 1)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		key := l.OptString(2, "")

		_, ok := state.ScrRegistry[key]
		if ok {
			l.Push(lua.LTrue)
			return 1
		}
		l.Push(lua.LFalse)
		return 1
	}))

	return tbl
}

func rubbleRegistryChildTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		regdata := l.ToUserData(1).Value.(rubble7.ScrRegData)
		key := l.OptString(2, "")

		switch key {
		case "listappend":
			l.Push(l.NewFunction(listAppend))
			return 1
		case "list":
			l.Push(&lua.LUserData{
				Value:     regdata.List,
				Metatable: strsliceTbl(l, true),
			})
			return 1
		case "table":
			l.Push(&lua.LUserData{
				Value:     regdata.Table,
				Metatable: strmapTbl(l),
			})
			return 1
		}
		return 0
	}))

	return tbl
}
