/*
Copyright 2014-2015 by Milo Christiansen

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

// Rhino AXIS VFS script commands.
package axisrhino

import "dctech/rhino"
import "dctech/axis"

// Module returns a map containing the Rhino AXIS VFS commands.
// The AXIS VFS commands are:
//	axis:newdir
//	axis:newfile
//	axis:mount
//	axis:getchild
//	axis:read
//	axis:write
//	axis:exists
//	axis:isdir
//	axis:del
//	axis:walkdirs
//	axis:walkfiles
func Module() *rhino.Value {
	return rhino.NewValueIndex(rhino.NewStaticMap(map[string]*rhino.Value {
		"newdir": rhino.NewValueCommand(rhino.NewNativeCommand(Command_NewDir)),
		"newfile": rhino.NewValueCommand(rhino.NewNativeCommand(Command_NewFile)),
		"mount": rhino.NewValueCommand(rhino.NewNativeCommand(Command_Mount)),
		"getchild": rhino.NewValueCommand(rhino.NewNativeCommand(Command_GetChild)),
		
		"read": rhino.NewValueCommand(rhino.NewNativeCommand(Command_Read)),
		"write": rhino.NewValueCommand(rhino.NewNativeCommand(Command_Write)),
		
		"exists": rhino.NewValueCommand(rhino.NewNativeCommand(Command_Exists)),
		"isdir": rhino.NewValueCommand(rhino.NewNativeCommand(Command_IsDir)),
		
		"del": rhino.NewValueCommand(rhino.NewNativeCommand(Command_Del)),
		
		"walkdirs": rhino.NewValueCommand(rhino.NewNativeCommand(Command_WalkDirs)),
		"walkfiles": rhino.NewValueCommand(rhino.NewNativeCommand(Command_WalkFiles)),
	}))
}

// Create a new empty AXIS logical directory Collection.
// 	axis:newdir
// Returns a reference to the new directory.
func Command_NewDir(rtm *rhino.Runtime) {
	rtm.RetVal.SetUser(axis.NewLogicalDir())
}

// Create a new read/write AXIS logical file DataSource.
// 	axis:newfile content
// Returns a reference to the new file.
func Command_NewFile(rtm *rhino.Runtime) {
	pcount := rtm.Locals.BlockLen()
	if pcount != 1 {
		rhino.ErrorParamCount("axis:newfile", "1")
	}
	
	rtm.RetVal.SetUser(axis.NewLogicalFile([]byte(rtm.Locals.Get(0).String()), true))
}

// Mount an AXIS DataSource on an AXIS Collection.
// 	axis:mount collection path ds
// Returns unchanged.
func Command_Mount(rtm *rhino.Runtime) {
	pcount := rtm.Locals.BlockLen()
	if pcount != 3 {
		rhino.ErrorParamCount("axis:mount", "3")
	}
	
	col, ok := rtm.Locals.Get(0).Data.(axis.Collection)
	if !ok {
		rhino.ErrorGeneralCmd("axis:mount", "Param 0 is not an axis.Collection")
	}
	
	ds, ok := rtm.Locals.Get(2).Data.(axis.DataSource)
	if !ok {
		rhino.ErrorGeneralCmd("axis:mount", "Param 2 is not an axis.DataSource")
	}

	axis.Mount(col, rtm.Locals.Get(1).String(), ds)
}

// Get an AXIS DataSource from an AXIS Collection by path.
// 	axis:getchild collection path
// Returns the DataSource or an error message. May set the Error flag.
func Command_GetChild(rtm *rhino.Runtime) {
	pcount := rtm.Locals.BlockLen()
	if pcount != 2 {
		rhino.ErrorParamCount("axis:getchild", "2")
	}
	
	col, ok := rtm.Locals.Get(0).Data.(axis.Collection)
	if !ok {
		rhino.ErrorGeneralCmd("axis:getchild", "Param 0 is not an axis.Collection")
	}
	
	ds, err := axis.GetChild(col, rtm.Locals.Get(1).String())
	if err != nil {
		rtm.Error = true
		rtm.RetVal.SetString(err.Error())
		return
	}
	rtm.Error = false
	rtm.RetVal.SetUser(ds)
}

// Read from an AXIS file DataSource.
// 	axis:read ds path
// Returns file contents or an error message. May set the Error flag.
func Command_Read(rtm *rhino.Runtime) {
	pcount := rtm.Locals.BlockLen()
	if pcount != 2 {
		rhino.ErrorParamCount("axis:read", "2")
	}
	
	ds, ok := rtm.Locals.Get(0).Data.(axis.DataSource)
	if !ok {
		rhino.ErrorGeneralCmd("axis:read", "Param 0 is not an axis.DataSource")
	}

	file, err := axis.ReadAll(ds, rtm.Locals.Get(1).String())
	if err != nil {
		rtm.Error = true
		rtm.RetVal.SetString(err.Error())
		return
	}
	rtm.Error = false
	rtm.RetVal.SetString(string(file))
}

// Write to an AXIS file DataSource.
// 	axis:write ds path content
// If the directories in the path do not exist axis:write tries to create them.
// Returns unchanged or an error message. May set the Error flag.
func Command_Write(rtm *rhino.Runtime) {
	pcount := rtm.Locals.BlockLen()
	if pcount != 3 {
		rhino.ErrorParamCount("axis:write", "3")
	}
	
	ds, ok := rtm.Locals.Get(0).Data.(axis.DataSource)
	if !ok {
		rhino.ErrorGeneralCmd("axis:write", "Param 0 is not an axis.DataSource")
	}

	err := axis.Create(ds, rtm.Locals.Get(1).String())
	if err != nil {
		rtm.Error = true
		rtm.RetVal.SetString(err.Error())
		return
	}
	err = axis.WriteAll(ds, rtm.Locals.Get(1).String(), []byte(rtm.Locals.Get(2).String()))
	if err != nil {
		rtm.Error = true
		rtm.RetVal.SetString(err.Error())
		return
	}
	rtm.Error = false
}

// Does an AXIS DataSource exist at the given path?
// 	axis:exists ds path
// Returns true or false.
func Command_Exists(rtm *rhino.Runtime) {
	pcount := rtm.Locals.BlockLen()
	if pcount != 2 {
		rhino.ErrorParamCount("axis:exists", "2")
	}
	
	ds, ok := rtm.Locals.Get(0).Data.(axis.DataSource)
	if !ok {
		rhino.ErrorGeneralCmd("axis:exists", "Param 0 is not an axis.DataSource")
	}

	rtm.RetVal.SetBool(axis.Exists(ds, rtm.Locals.Get(1).String()))
}

// Is an AXIS DataSource a directory?
// 	axis:isdir ds path
// Returns true or false.
func Command_IsDir(rtm *rhino.Runtime) {
	pcount := rtm.Locals.BlockLen()
	if pcount != 2 {
		rhino.ErrorParamCount("axis:isdir", "2")
	}
	
	ds, ok := rtm.Locals.Get(0).Data.(axis.DataSource)
	if !ok {
		rhino.ErrorGeneralCmd("axis:isdir", "Param 0 is not an axis.DataSource")
	}

	rtm.RetVal.SetBool(axis.IsDir(ds, rtm.Locals.Get(1).String()))
}

// Delete an AXIS DataSource.
// 	axis:del ds path
// Returns true or false.
func Command_Del(rtm *rhino.Runtime) {
	pcount := rtm.Locals.BlockLen()
	if pcount != 2 {
		rhino.ErrorParamCount("axis:del", "2")
	}
	
	ds, ok := rtm.Locals.Get(0).Data.(axis.DataSource)
	if !ok {
		rhino.ErrorGeneralCmd("axis:del", "Param 0 is not an axis.DataSource")
	}

	err := axis.Delete(ds, rtm.Locals.Get(1).String())
	if err != nil {
		rtm.RetVal.SetString(err.Error())
		rtm.Error = true
	}
	rtm.Error = false
}

// Iterate over all the directories in a AXIS DataSource.
// 	axis:walkdirs ds path code
// Runs code for every directory found, params:
//	path
// code MUST be a block created via a block declaration!
// Returns nil.
func Command_WalkDirs(rtm *rhino.Runtime) {
	pcount := rtm.Locals.BlockLen()
	if pcount != 3 {
		rhino.ErrorParamCount("axis:walkdirs", "3")
	}
	
	ds, ok := rtm.Locals.Get(0).Data.(axis.DataSource)
	if !ok {
		rhino.ErrorGeneralCmd("axis:walkdirs", "Param 0 is not an axis.DataSource")
	}

	code := rtm.Locals.Get(2)
	if code.Type != rhino.TypBlock {
		rhino.ErrorGeneralCmd("axis:walkdirs", "Attempt to run non-executable Value.")
	}
	block := code.Data.(*rhino.Block)
	
	files := axis.ListDir(ds, rtm.Locals.Get(1).String())
	
	rtm.Locals.Add(block)
	for _, file := range files {
		rtm.Locals.Get(0).SetString(file)
		rtm.Exec(block)
		rtm.Return = false
	}
	rtm.Locals.Remove()
	rtm.RetVal.Clear()
}

// Iterate over all the files in a AXIS DataSource.
// 	axis:walkfiles ds path code
// Runs code for every directory found, params:
//	path
// code MUST be a block created via a block declaration!
// Returns nil.
func Command_WalkFiles(rtm *rhino.Runtime) {
	pcount := rtm.Locals.BlockLen()
	if pcount != 3 {
		rhino.ErrorParamCount("axis:walkfiles", "3")
	}
	
	ds, ok := rtm.Locals.Get(0).Data.(axis.DataSource)
	if !ok {
		rhino.ErrorGeneralCmd("axis:walkfiles", "Param 0 is not an axis.DataSource")
	}

	code := rtm.Locals.Get(2)
	if code.Type != rhino.TypBlock {
		rhino.ErrorGeneralCmd("axis:walkfiles", "Attempt to run non-executable Value.")
	}
	block := code.Data.(*rhino.Block)
	
	files := axis.ListFile(ds, rtm.Locals.Get(1).String())
	
	rtm.Locals.Add(block)
	for _, file := range files {
		rtm.Locals.Get(0).SetString(file)
		rtm.Exec(block)
		rtm.Return = false
	}
	rtm.Locals.Remove()
	rtm.RetVal.Clear()
}
