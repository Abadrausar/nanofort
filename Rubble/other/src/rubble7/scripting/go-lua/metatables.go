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

// Lua scripting for Rubble.
// 
// This uses "github.com/Shopify/go-lua" as the runtime.
// 
// To use Lua with Rubble simply use `import _ "rubble7/scripting/go-lua"`.
// Lua scripting will then be available to all Rubble States by default.
package lua

import "github.com/Shopify/go-lua"
import "rubble7"
import "rubble7/rblutil/addon"
import "html/template"

func rubbleFilesTable(l *lua.State) {
	l.CreateTable(0, 2)
	
	l.PushGoFunction(func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		name := lua.OptString(l, 2, "")
		file, ok := state.Files.Data[name]
		if ok {
			l.PushUserData(file)
			fileTable(l) // From metatables.go
			l.SetMetaTable(-2)
			return 1
		}
		return 0
	})
	l.SetField(-2, "__index")
	
	l.PushGoFunction(func(l *lua.State) int {
		l.PushGoFunction(func(l *lua.State) int {
			l.Field(lua.RegistryIndex, "RUBBLE_STATE")
			state := l.ToUserData(-1).(*rubble7.State)
			l.Pop(1)
			
			i, _ := l.ToInteger(2)
			i++
			if i < len(state.Files.Order) {
				l.PushInteger(i)
				l.PushString(state.Files.Order[i])
				l.PushUserData(state.Files.Data[state.Files.Order[i]])
				fileTable(l) // From metatables.go
				l.SetMetaTable(-2)
				return 3
			}
			l.PushNil()
			return 1
		})
		l.PushValue(1)
		l.PushInteger(0)
		return 3
	})
	l.SetField(-2, "__ipairs")
}

func rubbleAddonsTable(l *lua.State) {
	l.CreateTable(0, 2)
	
	l.PushGoFunction(func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		l.PushInteger(len(state.Addons))
		return 1
	})
	l.SetField(-2, "__len")
	
	l.PushGoFunction(func(l *lua.State) int {
		l.Field(lua.RegistryIndex, "RUBBLE_STATE")
		state := l.ToUserData(-1).(*rubble7.State)
		l.Pop(1)
		
		idx := lua.OptInteger(l, 2, 0)
		if idx < 0 || idx >= len(state.Addons) {
			return 0
		}
		
		l.PushUserData(state.Addons[idx])
		addonTable(l) // From metatables.go
		l.SetMetaTable(-2)
		return 1
	})
	l.SetField(-2, "__index")
}

func addonTable(l *lua.State) {
	l.CreateTable(0, 2)
	
	l.PushGoFunction(func(l *lua.State) int {
		addon := l.ToUserData(1).(*addon.Addon)
		key := lua.OptString(l, 2, "")
		
		switch key {
		case "Active":
			l.PushBoolean(addon.Active)
			return 1
		case "Files":
			l.PushUserData(addon.Files)
			addonfilesTable(l) // From metatables.go
			l.SetMetaTable(-2)
			return 1
		case "Meta":
			l.PushUserData(addon.Meta)
			addonmetaTable(l) // From metatables.go
			l.SetMetaTable(-2)
			return 1
		}
		return 0
	})
	l.SetField(-2, "__index")
	
	l.PushGoFunction(func(l *lua.State) int {
		addon := l.ToUserData(1).(*addon.Addon)
		key := lua.OptString(l, 2, "")
		
		switch key {
		case "Active":
			addon.Active = l.ToBoolean(3)
			return 0
		}
		return 0
	})
	l.SetField(-2, "__newindex")
}

type addonfilesIter struct {
	keys []string
	list map[string]*addon.File
}

func addonfilesTable(l *lua.State) {
	l.CreateTable(0, 2)
	
	l.PushGoFunction(func(l *lua.State) int {
		files := l.ToUserData(1).(map[string]*addon.File)
		name := lua.OptString(l, 2, "")
		file, ok := files[name]
		if ok {
			l.PushUserData(file)
			fileTable(l) // From metatables.go
			l.SetMetaTable(-2)
			return 1
		}
		return 0
	})
	l.SetField(-2, "__index")
	
	l.PushGoFunction(func(l *lua.State) int {
		files := l.ToUserData(1).(map[string]*addon.File)
		il := make([]string, 0, len(files))
		for i := range files {
			il = append(il, i)
		}
		
		l.PushGoFunction(func(l *lua.State) int {
			iter := l.ToUserData(1).(addonfilesIter)
			
			i, _ := l.ToInteger(2)
			i++
			if i < len(iter.keys) {
				l.PushInteger(i)
				l.PushString(iter.keys[i])
				l.PushUserData(iter.list[iter.keys[i]])
				fileTable(l) // From metatables.go
				l.SetMetaTable(-2)
				return 3
			}
			l.PushNil()
			return 1
		})
		l.PushUserData(addonfilesIter{
			keys: il,
			list: files,
		})
		l.PushInteger(0)
		return 3
	})
	l.SetField(-2, "__pairs")
}

func addonmetaTable(l *lua.State) {
	l.CreateTable(0, 2)
	
	l.PushGoFunction(func(l *lua.State) int {
		meta := l.ToUserData(1).(*addon.Meta)
		key := lua.OptString(l, 2, "")
		
		switch key {
		case "Tags":
			l.PushUserData(meta.Tags)
			tagsTable(l) // From metatables.go
			l.SetMetaTable(-2)
			return 1
		case "Name":
			l.PushString(meta.Name)
			return 1
		case "Header":
			l.PushString(string(meta.Header))
			return 1
		case "Description":
			l.PushString(string(meta.Description))
			return 1
		case "Activates":
			l.PushUserData(meta.Activates)
			strsliceTable(l) // From metatables.go
			l.SetMetaTable(-2)
			return 1
		case "Incompatible":
			l.PushUserData(meta.Incompatible)
			strsliceTable(l) // From metatables.go
			l.SetMetaTable(-2)
			return 1
		case "Vars":
			// TODO:
			//return 1
		case "LoadPriority":
			l.PushInteger(meta.LoadPriority)
			return 1
		case "Author":
			l.PushString(meta.Author)
			return 1
		case "Version":
			l.PushString(meta.Version)
			return 1
		}
		return 0
	})
	l.SetField(-2, "__index")
	
	l.PushGoFunction(func(l *lua.State) int {
		meta := l.ToUserData(1).(*addon.Meta)
		key := lua.OptString(l, 2, "")
		
		switch key {
		case "Tags":
			return 0
		case "Name":
			return 0
		case "Header":
			meta.Header = template.HTML(lua.OptString(l, 3, ""))
			return 1
		case "Description":
			meta.Description = template.HTML(lua.OptString(l, 3, ""))
			return 1
		case "Activates":
			return 0
		case "Incompatible":
			return 0
		case "Vars":
			return 0
		case "LoadPriority":
			return 0
		case "Author":
			meta.Author = lua.OptString(l, 3, "")
			return 1
		case "Version":
			meta.Version = lua.OptString(l, 3, "")
			return 1
		}
		return 0
	})
	l.SetField(-2, "__newindex")
}

func strsliceTable(l *lua.State) {
	l.CreateTable(0, 2)
	
	l.PushGoFunction(func(l *lua.State) int {
		slice := l.ToUserData(1).([]string)
		
		l.PushInteger(len(slice))
		return 1
	})
	l.SetField(-2, "__len")
	
	l.PushGoFunction(func(l *lua.State) int {
		slice := l.ToUserData(1).([]string)
		
		idx := lua.OptInteger(l, 2, 0)
		if idx < 0 || idx >= len(slice) {
			return 0
		}
		
		l.PushString(slice[idx])
		return 1
	})
	l.SetField(-2, "__index")
}

func fileTable(l *lua.State) {
	l.CreateTable(0, 2)
	
	l.PushGoFunction(func(l *lua.State) int {
		file := l.ToUserData(1).(*addon.File)
		key := lua.OptString(l, 2, "")
		
		switch key {
		case "Content":
			l.PushString(string(file.Content))
			return 1
		case "Name":
			l.PushString(file.Name)
			return 1
		case "Source":
			l.PushString(file.Source)
			return 1
		case "Tags":
			l.PushUserData(file.Tags)
			tagsTable(l) // From metatables.go
			l.SetMetaTable(-2)
			return 1
		}
		return 0
	})
	l.SetField(-2, "__index")
	
	l.PushGoFunction(func(l *lua.State) int {
		file := l.ToUserData(1).(*addon.File)
		key := lua.OptString(l, 2, "")
		data := lua.OptString(l, 3, "")
		
		switch key {
		case "Content":
			file.Content = []byte(data)
			return 0
		case "Name":
			file.Name = data
		case "Source":
			file.Source = data
			return 0
		// Cant set Tags
		}
		return 0
	})
	l.SetField(-2, "__newindex")
}

type tagsIter struct {
	keys []string
	list map[string]bool
}

func tagsTable(l *lua.State) {
	l.CreateTable(0, 3)
	
	l.PushGoFunction(func(l *lua.State) int {
		tags := l.ToUserData(1).(map[string]bool)
		name := lua.OptString(l, 2, "")
		val, ok := tags[name]
		if ok {
			l.PushBoolean(val)
			return 1
		}
		return 0
	})
	l.SetField(-2, "__index")
	
	l.PushGoFunction(func(l *lua.State) int {
		tags := l.ToUserData(1).(map[string]bool)
		name := lua.OptString(l, 2, "")
		val := l.ToBoolean(3)
		tags[name] = val
		return 0
	})
	l.SetField(-2, "__newindex")
	
	l.PushGoFunction(func(l *lua.State) int {
		tags := l.ToUserData(1).(map[string]bool)
		il := make([]string, 0, len(tags))
		for i := range tags {
			il = append(il, i)
		}
		
		l.PushGoFunction(func(l *lua.State) int {
			iter := l.ToUserData(1).(tagsIter)
			
			i, _ := l.ToInteger(2)
			i++
			if i < len(iter.keys) {
				l.PushInteger(i)
				l.PushString(iter.keys[i])
				l.PushBoolean(iter.list[iter.keys[i]])
				return 3
			}
			l.PushNil()
			return 1
		})
		l.PushUserData(tagsIter{
			keys: il,
			list: tags,
		})
		l.PushInteger(0)
		return 3
	})
	l.SetField(-2, "__pairs")
}
