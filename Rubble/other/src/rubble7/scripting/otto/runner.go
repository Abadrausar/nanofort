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

// ECMA-262 (JavaScript) scripting for Rubble.
//
// This uses "github.com/robertkrimen/otto" as the runtime.
//
// To use Lua with Rubble simply use `import _ "rubble7/scripting/otto"`.
// Lua scripting will then be available to all Rubble States by default.
package lua

import "github.com/robertkrimen/otto"

import "rubble7"
import "rubble7/rblutil/addon"
import "rubble7/rblutil/errors"

import "strconv"

// JSRunner is a Runner for JavaScript scripts.
type JSRunner struct {
	js *otto.Otto
}

// New creates a new JSRunner.
// You do not normally need to call this, as rubble7.NewScriptCore will automatically do so.
func New(state *rubble7.State) rubble7.Runner {
	runner := &JSRunner{
		js: otto.New(),
	}
	js := runner.js

	//runner.js.Set("__RUBBLE_STATE__", state)

	obj, _ := js.Object("rubble = {}")
	
	obj.Set("version", rubble7.Version)
	obj.Set("vmajor", rubble7.VMajor)
	obj.Set("vminor", rubble7.VMinor)
	obj.Set("vpatch", rubble7.VPatch)
	
	obj.Set("string", func(call otto.FunctionCall) otto.Value {
		a, _ := call.Argument(0).Export()
		b, ok := a.([]byte)
		if !ok {
			errors.RaiseError("Value passed to 'rubble.string' is not a byte slice")
		}
		s, _ := call.Otto.ToValue(string(b))
		return s
	})
	
	obj.Set("bytes", func(call otto.FunctionCall) otto.Value {
		s, _ := call.Argument(0).ToString()
		b, _ := call.Otto.ToValue([]byte(s))
		return b
	})
	
	obj.Set("files", state.Files)
	obj.Set("gfiles", state.GlobalFiles)
	obj.Set("addons", state.Addons)
	obj.Set("addonstbl", state.AddonsTbl)
	
	addRegistry(js, state)
	
	return runner
}

func init() {
	addon.DefaultLast[".js"] = []string{"LangJS"}
	rubble7.AddDefaultRunner([]string{"LangJS"}, New)
}

func (runner *JSRunner) Run(script []byte, tag string, name string, line int, args []string) (string, error) {
	if line > 1 {
		name += "@" + strconv.Itoa(line)
	}
	js := runner.js

	//s1, _ := runner.js.Get("__RUBBLE_STATE__")
	//s2, _ := s1.Export()
	//state := s2.(*rubble7.State)

	js.Set("params", args)
	defer runner.js.Set("params", otto.UndefinedValue())

	scr, err := js.Compile(name, string(script))
	if err != nil {
		return "", err
	}

	rtn, err := js.Run(scr)
	if err != nil {
		return "", err
	}

	return rtn.String(), nil
}
