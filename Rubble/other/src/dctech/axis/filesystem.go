/*
Copyright 2013-2015 by Milo Christiansen

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

package axis

import "io"
import "sort"

// FileSystem is a simple Collection that serves it's mounted DataSources as locations.
type FileSystem map[string]DataSource

func NewFileSystem() Collection {
	return make(FileSystem)
}

func (fs FileSystem) Mount(loc string, ds DataSource) {
	fs[loc] = ds
}

func (fs FileSystem) GetChild(path *Path) (DataSource, error) {
	locID := path.NextLoc()
	if locID == "" {
		if path.Done() {
			return fs, nil
		}
		return nil, NewError(ErrNotFound, path)
	}
	
	if _, ok := fs[locID]; ok {
		if col, ok := fs[locID].(Collection); ok {
			return col.GetChild(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (fs FileSystem) IsDir(path *Path) bool {
	locID := path.NextLoc()
	if locID == "" {
		if path.Done() {
			return true
		}
		return false
	}
	defer path.RevertLoc()
	
	if _, ok := fs[locID]; ok {
		ds := fs[locID]
		if ds.Exists(path) {
			return ds.IsDir(path)
		}
	}
	return false
}

func (fs FileSystem) Exists(path *Path) bool {
	locID := path.NextLoc()
	if locID == "" {
		if path.Done() {
			return true
		}
		return false
	}
	defer path.RevertLoc()
	
	if _, ok := fs[locID]; ok {
		ds := fs[locID]
		if ds.Exists(path) {
			return true
		}
	}
	return false
}

func (fs FileSystem) Delete(path *Path) error {
	locID := path.NextLoc()
	if locID == "" {
		if path.Done() {
			return NewError(ErrBadAction, path)
		}
		return NewError(ErrNotFound, path)
	}
	if path.Done() {
		delete(fs, locID)
		return nil
	}
	
	if _, ok := fs[locID]; ok {
		ds := fs[locID]
		if ds.Exists(path) {
			return ds.Delete(path)
		}
	}
	return NewError(ErrNotFound, path)
}

func (fs FileSystem) Create(path *Path) error {
	locID := path.NextLoc()
	if locID == "" {
		if path.Done() {
			return NewError(ErrBadAction, path)
		}
		return NewError(ErrNotFound, path)
	}
	
	if _, ok := fs[locID]; ok {
		return fs[locID].Create(path)
	}
	return NewError(ErrReadOnly, path)
}

func (fs FileSystem) ListDir(path *Path) []string {
	locID := path.NextLoc()
	if locID == "" {
		if !path.Done() {
			return []string{}
		}
		
		rtn := make([]string, 0, len(fs))
		for item := range fs {
			if !fs[item].IsDir(path) {
				rtn = append(rtn, item)
			}
		}
		sort.Strings(rtn)
		return rtn
	}

	if _, ok := fs[locID]; ok {
		ds := fs[locID]
		if ds.Exists(path) {
			return ds.ListDir(path)
		}
	}
	return []string{}
}

func (fs FileSystem) ListFile(path *Path) []string {
	locID := path.NextLoc()
	if locID == "" {
		if !path.Done() {
			return []string{}
		}
		
		rtn := make([]string, 0, len(fs))
		for item := range fs {
			if !fs[item].IsDir(path) {
				rtn = append(rtn, item)
			}
		}
		sort.Strings(rtn)
		return rtn
	}
	
	if _, ok := fs[locID]; ok {
		ds := fs[locID]
		if ds.Exists(path) {
			return ds.ListFile(path)
		}
	}
	return []string{}
}

func (fs FileSystem) Read(path *Path) (io.ReadCloser, error) {
	locID := path.NextLoc()
	if locID == "" {
		if path.Done() {
			return nil, NewError(ErrBadAction, path)
		}
		return nil, NewError(ErrNotFound, path)
	}
	
	if _, ok := fs[locID]; ok {
		ds := fs[locID]
		if ds.Exists(path) {
			return ds.Read(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (fs FileSystem) ReadAll(path *Path) ([]byte, error) {
	locID := path.NextLoc()
	if locID == "" {
		if path.Done() {
			return nil, NewError(ErrBadAction, path)
		}
		return nil, NewError(ErrNotFound, path)
	}
	
	if _, ok := fs[locID]; ok {
		ds := fs[locID]
		if ds.Exists(path) {
			return ds.ReadAll(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (fs FileSystem) Write(path *Path) (io.WriteCloser, error) {
	locID := path.NextLoc()
	if locID == "" {
		if path.Done() {
			return nil, NewError(ErrBadAction, path)
		}
		return nil, NewError(ErrNotFound, path)
	}
	
	if _, ok := fs[locID]; ok {
		ds := fs[locID]
		if ds.Exists(path) {
			return ds.Write(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (fs FileSystem) WriteAll(path *Path, content []byte) error {
	locID := path.NextLoc()
	if locID == "" {
		if path.Done() {
			return NewError(ErrBadAction, path)
		}
		return NewError(ErrNotFound, path)
	}
	
	if _, ok := fs[locID]; ok {
		ds := fs[locID]
		if ds.Exists(path) {
			return ds.WriteAll(path, content)
		}
	}
	return NewError(ErrNotFound, path)
}
