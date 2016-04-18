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

// For some reason the default versions of pairs and ipairs do not support metatables,
// so I redefine them here (since the __pairs and __ipairs metamethods are important
// to the Rubble Lua API).

func pairsaux(l *lua.LState) int {
	tb := l.CheckTable(1)
	key, value := tb.Next(l.Get(2))
	if key == lua.LNil {
		return 0
	} else {
		l.Pop(1)
		l.Push(key)
		l.Push(key)
		l.Push(value)
		return 2
	}
}

func basePairs(l *lua.LState) int {
	switch v := l.CheckAny(1).(type) {
	case *lua.LTable:
		l.Push(l.Get(lua.UpvalueIndex(1)))
		l.Push(v)
		l.Push(lua.LNil)
		return 3
	default:
		meta := l.GetMetatable(v).(*lua.LTable)
		if meta == nil {
			l.CheckTable(1) // We already know it's not a table, this is just a lazy way to get an error message.
		}

		l.Push(meta.RawGetString("__pairs"))
		l.CheckFunction(-1) // Error check
		l.Push(v)
		l.Call(1, 3)
		return 3
	}
}

func ipairsaux(l *lua.LState) int {
	tb := l.CheckTable(1)
	i := l.CheckInt(2)
	i++
	v := tb.RawGetInt(i)
	if v == lua.LNil {
		return 0
	} else {
		l.Pop(1)
		l.Push(lua.LNumber(i))
		l.Push(lua.LNumber(i))
		l.Push(v)
		return 2
	}
}

func baseIpairs(l *lua.LState) int {
	switch v := l.CheckAny(1).(type) {
	case *lua.LTable:
		l.Push(l.Get(lua.UpvalueIndex(1)))
		l.Push(v)
		l.Push(lua.LNumber(0))
		return 3
	default:
		meta := l.GetMetatable(v).(*lua.LTable)
		if meta == nil {
			l.CheckTable(1) // We already know it's not a table, this is just a lazy way to get an error message.
		}

		l.Push(meta.RawGetString("__ipairs"))
		l.CheckFunction(-1) // Error check
		l.Push(v)
		l.Call(1, 3)
		return 3
	}
}
