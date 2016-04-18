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
import "rubble7/rblutil/addon"

import "html/template"

import "math/rand"
import "time"
import "strconv"
import "hash/crc64"

// *rand.Rand
func rubbleRandomTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		rng := l.ToUserData(1).Value.(*rand.Rand)
		key := l.OptString(2, "")

		switch key {
		case "int":
			l.Push(lua.LNumber(rng.Uint32()))
			return 1
		case "intn":
			l.Push(l.NewFunction(func(l *lua.LState) int {
				rng := l.ToUserData(1).Value.(*rand.Rand)
				n := l.OptInt(2, 1)

				// You would think that "[0,n]" when describing a range would be inclusive, but apparently not...
				l.Push(lua.LNumber(rng.Int31n(int32(n) + 1)))
				return 1
			}))
			return 1
		case "float":
			l.Push(lua.LNumber(rng.Float64()))
			return 1
		case "seed":
			l.Push(lua.LString(strconv.FormatInt(time.Now().UnixNano(), 16)))
			return 1
		}
		return 0
	}))

	tbl.RawSetString("__newindex", l.NewFunction(func(l *lua.LState) int {
		rng := l.ToUserData(1).Value.(*rand.Rand)
		key := l.OptString(2, "")

		switch key {
		case "seed":
			str := l.OptString(3, "")
			seed, err := strconv.ParseInt(str, 16, 64)
			if err != nil {
				seed = int64(crc64.Checksum([]byte(str), crc64.MakeTable(crc64.ECMA)))
			}
			rng.Seed(seed)
			return 0
		}
		return 0
	}))

	return tbl
}

type addonFileListIter struct {
	i    int
	list *addon.FileList
}

// rubble7.*State.Files
func rubbleStateFilesTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		name := l.OptString(2, "")
		file, ok := state.Files.Data[name]
		if ok {
			l.Push(&lua.LUserData{
				Value:     file,
				Metatable: addonFileTbl(l),
			})
			return 1
		}
		return 0
	}))

	pairs := l.NewFunction(func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		l.Push(l.NewFunction(func(l *lua.LState) int {
			iter := l.ToUserData(1).Value.(*addonFileListIter)

			iter.i++
			if iter.i < len(iter.list.Order) {
				l.Push(lua.LString(iter.list.Order[iter.i]))
				l.Push(&lua.LUserData{
					Value:     iter.list.Data[iter.list.Order[iter.i]],
					Metatable: addonFileTbl(l),
				})
				return 2
			}
			l.Push(lua.LNil)
			return 1
		}))
		l.Push(&lua.LUserData{
			Value: &addonFileListIter{
				i:    -1,
				list: state.Files,
			},
		})
		l.Push(lua.LNil)
		return 3
	})

	tbl.RawSetString("__ipairs", pairs)
	tbl.RawSetString("__pairs", pairs)

	return tbl
}

// rubble7.*State.GlobalFiles
func rubbleStateGlobalFilesTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		name := l.OptString(2, "")
		file, ok := state.GlobalFiles.Data[name]
		if ok {
			l.Push(&lua.LUserData{
				Value:     file,
				Metatable: addonFileTbl(l),
			})
			return 1
		}
		return 0
	}))

	pairs := l.NewFunction(func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		l.Push(l.NewFunction(func(l *lua.LState) int {
			iter := l.ToUserData(1).Value.(*addonFileListIter)

			iter.i++
			if iter.i < len(iter.list.Order) {
				l.Push(lua.LString(iter.list.Order[iter.i]))
				l.Push(&lua.LUserData{
					Value:     iter.list.Data[iter.list.Order[iter.i]],
					Metatable: addonFileTbl(l),
				})
				return 2
			}
			l.Push(lua.LNil)
			return 1
		}))
		l.Push(&lua.LUserData{
			Value: &addonFileListIter{
				i:    -1,
				list: state.GlobalFiles,
			},
		})
		l.Push(lua.LNil)
		return 3
	})

	tbl.RawSetString("__ipairs", pairs)
	tbl.RawSetString("__pairs", pairs)

	return tbl
}

// rubble7.*State.Addons
func rubbleStateAddonsTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 4)

	tbl.RawSetString("__len", l.NewFunction(func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		l.Push(lua.LNumber(len(state.Addons)))
		return 1
	}))

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		i := l.OptInt(2, 0)
		if i < 0 || i >= len(state.Addons) {
			return 0
		}

		l.Push(&lua.LUserData{
			Value:     state.Addons[i],
			Metatable: addonTbl(l),
		})
		return 1
	}))

	pairs := l.NewFunction(func(l *lua.LState) int {
		l.Push(l.NewFunction(func(l *lua.LState) int {
			state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

			i := l.ToInt(2)
			i++
			if i < len(state.Addons) {
				l.Push(lua.LNumber(i))
				l.Push(&lua.LUserData{
					Value:     state.Addons[i],
					Metatable: addonTbl(l),
				})
				return 2
			}
			l.Push(lua.LNil)
			return 1
		}))
		l.Push(l.Get(1))
		l.Push(lua.LNumber(-1))
		return 3
	})

	tbl.RawSetString("__ipairs", pairs)
	tbl.RawSetString("__pairs", pairs)

	return tbl
}

// rubble7.*State.AddonsTbl
func rubbleStateAddonsTblTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 1)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		state := l.GetField(l.CheckTable(lua.RegistryIndex), "RUBBLE_STATE").(*lua.LUserData).Value.(*rubble7.State)

		addon, ok := state.AddonsTbl[l.OptString(2, "")]
		if !ok {
			return 0
		}

		l.Push(&lua.LUserData{
			Value:     addon,
			Metatable: addonTbl(l),
		})
		return 1
	}))

	return tbl
}

// *addon.Addon
func addonTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		addon := l.ToUserData(1).Value.(*addon.Addon)
		key := l.OptString(2, "")

		switch key {
		case "Active":
			if addon.Active {
				l.Push(lua.LTrue)
				return 1
			}
			l.Push(lua.LFalse)
			return 1
		case "Files":
			l.Push(&lua.LUserData{
				Value:     addon.Files,
				Metatable: addonFilemapTbl(l),
			})
			return 1
		case "Meta":
			l.Push(&lua.LUserData{
				Value:     addon.Meta,
				Metatable: addonMetaTbl(l),
			})
			return 1
		}
		return 0
	}))

	tbl.RawSetString("__newindex", l.NewFunction(func(l *lua.LState) int {
		addon := l.ToUserData(1).Value.(*addon.Addon)
		key := l.OptString(2, "")

		switch key {
		case "Active":
			addon.Active = lua.LVAsBool(l.Get(3))
			return 0
		}
		return 0
	}))

	return tbl
}

type addonFilemapIter struct {
	i    int
	keys []string
	list map[string]*addon.File
}

// map[string]*addon.File
func addonFilemapTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		files := l.ToUserData(1).Value.(map[string]*addon.File)
		name := l.OptString(2, "")
		file, ok := files[name]
		if ok {
			l.Push(&lua.LUserData{
				Value:     file,
				Metatable: addonFileTbl(l),
			})
			return 1
		}
		return 0
	}))

	tbl.RawSetString("__pairs", l.NewFunction(func(l *lua.LState) int {
		files := l.ToUserData(1).Value.(map[string]*addon.File)
		il := make([]string, 0, len(files))
		for i := range files {
			il = append(il, i)
		}

		l.Push(l.NewFunction(func(l *lua.LState) int {
			iter := l.ToUserData(1).Value.(*addonFilemapIter)

			iter.i++
			if iter.i < len(iter.keys) {
				l.Push(lua.LString(iter.keys[iter.i]))
				l.Push(&lua.LUserData{
					Value:     iter.list[iter.keys[iter.i]],
					Metatable: addonFileTbl(l),
				})
				return 2
			}
			l.Push(lua.LNil)
			return 1
		}))
		l.Push(&lua.LUserData{
			Value: &addonFilemapIter{
				i:    -1,
				keys: il,
				list: files,
			},
		})
		l.Push(lua.LNil)
		return 3
	}))

	return tbl
}

// *addon.Meta
func addonMetaTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		meta := l.ToUserData(1).Value.(*addon.Meta)
		key := l.OptString(2, "")

		switch key {
		case "Tags":
			l.Push(&lua.LUserData{
				Value:     meta.Tags,
				Metatable: tagsTbl(l),
			})
			return 1
		case "Name":
			l.Push(lua.LString(meta.Name))
			return 1
		case "Header":
			l.Push(lua.LString(meta.Header))
			return 1
		case "Description":
			l.Push(lua.LString(meta.Description))
			return 1
		case "DescFile":
			l.Push(lua.LString(meta.DescFile))
			return 1
		case "Activates":
			l.Push(&lua.LUserData{
				Value:     &meta.Activates,
				Metatable: strsliceTbl(l, false),
			})
			return 1
		case "Incompatible":
			l.Push(&lua.LUserData{
				Value:     &meta.Incompatible,
				Metatable: strsliceTbl(l, false),
			})
			return 1
		case "Vars":
			l.Push(&lua.LUserData{
				Value:     meta.Vars,
				Metatable: addonMetaVarmapTbl(l),
			})
			return 1
		case "LoadPriority":
			l.Push(lua.LNumber(meta.LoadPriority))
			return 1
		case "Author":
			l.Push(lua.LString(meta.Author))
			return 1
		case "Version":
			l.Push(lua.LString(meta.Version))
			return 1
		}
		return 0
	}))

	tbl.RawSetString("__newindex", l.NewFunction(func(l *lua.LState) int {
		meta := l.ToUserData(1).Value.(*addon.Meta)
		key := l.OptString(2, "")

		switch key {
		case "Tags":
			return 0
		case "Name":
			return 0
		case "Header":
			meta.Header = template.HTML(lua.LVAsString(l.Get(3)))
			return 0
		case "Description":
			meta.Description = template.HTML(lua.LVAsString(l.Get(3)))
			return 0
		case "DescFile":
			return 0
		case "Activates":
			return 0
		case "Incompatible":
			return 0
		case "Vars":
			return 0
		case "LoadPriority":
			return 0
		case "Author":
			meta.Author = lua.LVAsString(l.Get(3))
			return 0
		case "Version":
			meta.Version = lua.LVAsString(l.Get(3))
			return 0
		}
		return 0
	}))

	return tbl
}

type addonMetaVarmapIter struct {
	i    int
	keys []string
	list map[string]*addon.MetaVar
}

// map[string]*addon.MetaVar
func addonMetaVarmapTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		vars := l.ToUserData(1).Value.(map[string]*addon.MetaVar)
		name := l.OptString(2, "")
		v, ok := vars[name]
		if ok {
			l.Push(&lua.LUserData{
				Value:     v,
				Metatable: addonMetaVarTbl(l),
			})
			return 1
		}
		return 0
	}))

	tbl.RawSetString("__pairs", l.NewFunction(func(l *lua.LState) int {
		vars := l.ToUserData(1).Value.(map[string]*addon.MetaVar)
		il := make([]string, 0, len(vars))
		for i := range vars {
			il = append(il, i)
		}

		l.Push(l.NewFunction(func(l *lua.LState) int {
			iter := l.ToUserData(1).Value.(*addonMetaVarmapIter)

			iter.i++
			if iter.i < len(iter.keys) {
				l.Push(lua.LString(iter.keys[iter.i]))
				l.Push(&lua.LUserData{
					Value:     iter.list[iter.keys[iter.i]],
					Metatable: addonMetaVarTbl(l),
				})
				return 2
			}
			l.Push(lua.LNil)
			return 1
		}))
		l.Push(&lua.LUserData{
			Value: &addonMetaVarmapIter{
				i:    -1,
				keys: il,
				list: vars,
			},
		})
		l.Push(lua.LNil)
		return 3
	}))

	return tbl
}

func addonMetaVarTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		val := l.ToUserData(1).Value.(*addon.MetaVar)
		key := l.OptString(2, "")

		switch key {
		case "Name":
			l.Push(lua.LString(val.Name))
			return 1
		case "Values":
			l.Push(&lua.LUserData{
				Value:     &val.Values,
				Metatable: strsliceTbl(l, true),
			})
			return 1
		}
		return 0
	}))

	tbl.RawSetString("__newindex", l.NewFunction(func(l *lua.LState) int {
		val := l.ToUserData(1).Value.(*addon.MetaVar)
		key := l.OptString(2, "")

		switch key {
		case "Name":
			val.Name = lua.LVAsString(l.Get(3))
			return 0
		}
		return 0
	}))

	return tbl
}

// []string
func strsliceTbl(l *lua.LState, rw bool) lua.LValue {
	tbl := l.CreateTable(0, 5)

	tbl.RawSetString("__len", l.NewFunction(func(l *lua.LState) int {
		slice := l.ToUserData(1).Value.(*[]string)

		l.Push(lua.LNumber(len(*slice)))
		return 1
	}))

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		slice := l.ToUserData(1).Value.(*[]string)

		idx := l.OptInt(2, 0)
		if idx < 0 || idx >= len(*slice) {
			return 0
		}

		l.Push(lua.LString((*slice)[idx]))
		return 1
	}))

	if rw {
		tbl.RawSetString("__newindex", l.NewFunction(func(l *lua.LState) int {
			slice := l.ToUserData(1).Value.(*[]string)

			idx := l.OptInt(2, -1)
			if idx < 0 || idx > len(*slice) {
				return 0
			}

			if idx == len(*slice) {
				*slice = append(*slice, lua.LVAsString(l.Get(3)))
				return 0
			}
			(*slice)[idx] = lua.LVAsString(l.Get(3))
			return 0
		}))
	}

	pairs := l.NewFunction(func(l *lua.LState) int {
		l.Push(l.NewFunction(func(l *lua.LState) int {
			slice := l.ToUserData(1).Value.(*[]string)

			i := l.ToInt(2)
			i++
			if i < len(*slice) {
				l.Push(lua.LNumber(i))
				l.Push(lua.LString((*slice)[i]))
				return 2
			}
			l.Push(lua.LNil)
			return 1
		}))
		l.Push(l.Get(1))
		l.Push(lua.LNumber(-1))
		return 3
	})

	tbl.RawSetString("__ipairs", pairs)
	tbl.RawSetString("__pairs", pairs)

	return tbl
}

// *addon.File
func addonFileTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 2)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		file := l.ToUserData(1).Value.(*addon.File)
		key := l.OptString(2, "")

		switch key {
		case "Content":
			l.Push(lua.LString(file.Content))
			return 1
		case "Name":
			l.Push(lua.LString(file.Name))
			return 1
		case "Source":
			l.Push(lua.LString(file.Source))
			return 1
		case "Tags":
			l.Push(&lua.LUserData{
				Value:     file.Tags,
				Metatable: tagsTbl(l),
			})
			return 1
		}
		return 0
	}))

	tbl.RawSetString("__newindex", l.NewFunction(func(l *lua.LState) int {
		file := l.ToUserData(1).Value.(*addon.File)
		key := l.OptString(2, "")

		switch key {
		case "Content":
			file.Content = []byte(lua.LVAsString(l.Get(3)))
			return 0
		case "Name":
			file.Name = lua.LVAsString(l.Get(3))
		case "Source":
			file.Source = lua.LVAsString(l.Get(3))
			return 0
			// Cant set Tags
		}
		return 0
	}))

	return tbl
}

type tagsIter struct {
	i    int
	keys []string
	list map[string]bool
}

// map[string]bool
func tagsTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 3)

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		tags := l.ToUserData(1).Value.(map[string]bool)
		name := l.OptString(2, "")
		val, ok := tags[name]
		if ok {
			if val {
				l.Push(lua.LTrue)
				return 1
			}
			l.Push(lua.LFalse)
			return 1
		}
		return 0
	}))

	tbl.RawSetString("__newindex", l.NewFunction(func(l *lua.LState) int {
		tags := l.ToUserData(1).Value.(map[string]bool)
		name := l.OptString(2, "")
		tags[name] = lua.LVAsBool(l.Get(3))
		return 0
	}))

	tbl.RawSetString("__pairs", l.NewFunction(func(l *lua.LState) int {
		tags := l.ToUserData(1).Value.(map[string]bool)
		il := make([]string, 0, len(tags))
		for i := range tags {
			il = append(il, i)
		}

		l.Push(l.NewFunction(func(l *lua.LState) int {
			iter := l.ToUserData(1).Value.(*tagsIter)

			iter.i++
			if iter.i < len(iter.keys) {
				l.Push(lua.LString(iter.keys[iter.i]))
				if iter.list[iter.keys[iter.i]] {
					l.Push(lua.LTrue)
					return 2
				}
				l.Push(lua.LFalse)
				return 2
			}
			l.Push(lua.LNil)
			return 1
		}))
		l.Push(&lua.LUserData{
			Value: &tagsIter{
				i:    -1,
				keys: il,
				list: tags,
			},
		})
		l.Push(lua.LNil)
		return 3
	}))

	return tbl
}

type strmapIter struct {
	i    int
	keys []string
	list map[string]string
}

// map[string]string
func strmapTbl(l *lua.LState) lua.LValue {
	tbl := l.CreateTable(0, 4)

	tbl.RawSetString("__len", l.NewFunction(func(l *lua.LState) int {
		strmap := l.ToUserData(1).Value.(map[string]string)

		l.Push(lua.LNumber(len(strmap)))
		return 1
	}))

	tbl.RawSetString("__index", l.NewFunction(func(l *lua.LState) int {
		strmap := l.ToUserData(1).Value.(map[string]string)
		name := l.OptString(2, "")
		val, ok := strmap[name]
		if ok {
			l.Push(lua.LString(val))
			return 1
		}
		l.Push(lua.LNil)
		return 1
	}))

	tbl.RawSetString("__newindex", l.NewFunction(func(l *lua.LState) int {
		strmap := l.ToUserData(1).Value.(map[string]string)
		name := l.OptString(2, "")
		strmap[name] = lua.LVAsString(l.Get(3))
		return 0
	}))

	tbl.RawSetString("__pairs", l.NewFunction(func(l *lua.LState) int {
		strmap := l.ToUserData(1).Value.(map[string]string)
		il := make([]string, 0, len(strmap))
		for i := range strmap {
			il = append(il, i)
		}

		l.Push(l.NewFunction(func(l *lua.LState) int {
			iter := l.ToUserData(1).Value.(*strmapIter)

			iter.i++
			if iter.i < len(iter.keys) {
				l.Push(lua.LString(iter.keys[iter.i]))
				l.Push(lua.LString(iter.list[iter.keys[iter.i]]))
				return 2
			}
			l.Push(lua.LNil)
			return 1
		}))
		l.Push(&lua.LUserData{
			Value: &strmapIter{
				i:    -1,
				keys: il,
				list: strmap,
			},
		})
		l.Push(lua.LNil)
		return 3
	}))

	return tbl
}
