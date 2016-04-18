/*
Copyright 2016 by Milo Christiansen

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

// ECMA-262 (JavaScript) scripting for Rubble.
//
// This uses "github.com/robertkrimen/otto" as the runtime.
//
// To use Lua with Rubble simply use `import _ "rubble7/scripting/otto"`.
// Lua scripting will then be available to all Rubble States by default.
package lua

import "github.com/robertkrimen/otto"

import "rubble7"
import "rubble7/rblutil/errors"

func addRegistry(js *otto.Otto, state *rubble7.State) {
	reg, _ := js.Object("rubble.registry = {}")
	
	reg.Set("exists", func(call otto.FunctionCall) otto.Value {
		master, err := call.Argument(0).ToString()
		if err != nil {
			// I hope this works, Otto has absolutely no documentation about how errors inside of
			// functions are supposed to be handled. It's like the author deliberately ignored the
			// issue, the examples all ignore errors and there is no mention of how errors are handled.
			panic(err)
		}
		_, ok := state.ScrRegistry[master]
		rtn, _ := call.Otto.ToValue(ok)
		return rtn
	})
	
	// List
	
	reg.Set("listGet", func(call otto.FunctionCall) otto.Value {
		master, err := call.Argument(0).ToString()
		if err != nil {
			panic(err)
		}
		index, err := call.Argument(1).ToInteger()
		if err != nil {
			panic(err)
		}
		i := int(index)
		
		child, ok := state.ScrRegistry[master]
		if !ok {
			child = rubble7.ScrRegData{
				List:  &[]string{},
				Table: map[string]string{},
			}
			state.ScrRegistry[master] = child
		}
		
		if i < 0 || i >= len(*child.List) {
			return otto.UndefinedValue()
		}
		
		rtn, _ := call.Otto.ToValue((*child.List)[i])
		return rtn
	})
	reg.Set("listSet", func(call otto.FunctionCall) otto.Value {
		master, err := call.Argument(0).ToString()
		if err != nil {
			panic(err)
		}
		index, err := call.Argument(1).ToInteger()
		if err != nil {
			panic(err)
		}
		i := int(index)
		value, err := call.Argument(2).ToString()
		if err != nil {
			panic(err)
		}
		
		child, ok := state.ScrRegistry[master]
		if !ok {
			child = rubble7.ScrRegData{
				List:  &[]string{},
				Table: map[string]string{},
			}
			state.ScrRegistry[master] = child
		}
		
		if i < 0 || i > len(*child.List) {
			errors.RaiseError("Index out of range.")
		}
		
		if i == len(*child.List) {
			*child.List = append(*child.List, value)
			return otto.UndefinedValue()
		}
		(*child.List)[i] = value
		return otto.UndefinedValue()
	})
	
	reg.Set("listLen", func(call otto.FunctionCall) otto.Value {
		master, err := call.Argument(0).ToString()
		if err != nil {
			panic(err)
		}
		
		child, ok := state.ScrRegistry[master]
		if !ok {
			rtn, _ := call.Otto.ToValue(0)
			return rtn
		}
		
		rtn, _ := call.Otto.ToValue(len(*child.List))
		return rtn
	})
	
	// Table
	
	reg.Set("tableGet", func(call otto.FunctionCall) otto.Value {
		master, err := call.Argument(0).ToString()
		if err != nil {
			panic(err)
		}
		index, err := call.Argument(1).ToString()
		if err != nil {
			panic(err)
		}
		
		child, ok := state.ScrRegistry[master]
		if !ok {
			child = rubble7.ScrRegData{
				List:  &[]string{},
				Table: map[string]string{},
			}
			state.ScrRegistry[master] = child
		}
		
		v, ok := child.Table[index]
		if !ok {
			return otto.UndefinedValue()
		}
		rtn, _ := call.Otto.ToValue(v)
		return rtn
	})
	reg.Set("tableSet", func(call otto.FunctionCall) otto.Value {
		master, err := call.Argument(0).ToString()
		if err != nil {
			panic(err)
		}
		index, err := call.Argument(1).ToString()
		if err != nil {
			panic(err)
		}
		value, err := call.Argument(2).ToString()
		if err != nil {
			panic(err)
		}
		
		child, ok := state.ScrRegistry[master]
		if !ok {
			child = rubble7.ScrRegData{
				List:  &[]string{},
				Table: map[string]string{},
			}
			state.ScrRegistry[master] = child
		}
		
		child.Table[index] = value
		return otto.UndefinedValue()
	})
	reg.Set("tableKeys", func(call otto.FunctionCall) otto.Value {
		panic("TODO")
		
		//master, err := call.Argument(0).ToString()
		//if err != nil {
		//	panic(err)
		//}
		//
		//rtn, _ := js.Object("[]")
		//
		//child, ok := state.ScrRegistry[master]
		//if !ok {
		//	return rtn.Value()
		//}
		//
		//for k := range child.Table {
		//}
		//
		//return rtn.Value()
	})
}
