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

// LogicalDir is a simple logical directory Collection that serves it's mounted DataSources as files
// or child directories.
type LogicalDir map[string]DataSource

// NewOSFile creates a new logical AXIS directory.
func NewLogicalDir() Collection {
	return make(LogicalDir)
}

func (dir LogicalDir) Mount(loc string, ds DataSource) {
	dir[loc] = ds
}

func (dir LogicalDir) GetChild(path *Path) (DataSource, error) {
	locID := path.NextDir()
	if locID == "" {
		if path.Done() {
			return dir, nil
		}
		return dir, nil
	}
	
	if _, ok := dir[locID]; ok {
		ds := dir[locID]
		if col, ok := ds.(Collection); ok {
			return col.GetChild(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (dir LogicalDir) IsDir(path *Path) bool {
	locID := path.NextDir()
	if locID == "" {
		if path.Done() {
			return true
		}
		return false
	}
	defer path.RevertDir()
	
	if _, ok := dir[locID]; ok {
		ds := dir[locID]
		if ds.Exists(path) {
			return ds.IsDir(path)
		}
	}
	return false
}

func (dir LogicalDir) Exists(path *Path) bool {
	locID := path.NextDir()
	if locID == "" {
		if path.Done() {
			return true
		}
		return false
	}
	defer path.RevertDir()
	
	if _, ok := dir[locID]; ok {
		ds := dir[locID]
		if ds.Exists(path) {
			return true
		}
	}
	return false
}

func (dir LogicalDir) Delete(path *Path) error {
	locID := path.NextDir()
	if locID == "" {
		if path.Done() {
			return NewError(ErrBadAction, path)
		}
		return NewError(ErrNotFound, path)
	}
	if path.Done() {
		delete(dir, locID)
		return nil
	}
	
	if _, ok := dir[locID]; ok {
		ds := dir[locID]
		if ds.Exists(path) {
			return ds.Delete(path)
		}
	}
	return NewError(ErrNotFound, path)
}

func (dir LogicalDir) Create(path *Path) error {
	locID := path.NextDir()
	if locID == "" {
		if path.Done() {
			return NewError(ErrBadAction, path)
		}
		return NewError(ErrNotFound, path)
	}
	
	if _, ok := dir[locID]; ok {
		return dir[locID].Create(path)
	}
	return NewError(ErrReadOnly, path)
}

func (dir LogicalDir) ListDir(path *Path) []string {
	locID := path.NextDir()
	if locID == "" {
		if !path.Done() {
			return []string{}
		}
		
		rtn := make([]string, 0, len(dir))
		for item := range dir {
			if dir[item].IsDir(path) {
				rtn = append(rtn, item)
			}
		}
		sort.Strings(rtn)
		return rtn
	}

	if _, ok := dir[locID]; ok {
		ds := dir[locID]
		if ds.Exists(path) {
			return ds.ListDir(path)
		}
	}
	return []string{}
}

func (dir LogicalDir) ListFile(path *Path) []string {
	locID := path.NextDir()
	if locID == "" {
		if !path.Done() {
			return []string{}
		}
		
		rtn := make([]string, 0, len(dir))
		for item := range dir {
			if !dir[item].IsDir(path) {
				rtn = append(rtn, item)
			}
		}
		sort.Strings(rtn)
		return rtn
	}
	
	if locID == "" {
		return []string{}
	}
	
	if _, ok := dir[locID]; ok {
		ds := dir[locID]
		if ds.Exists(path) {
			return ds.ListFile(path)
		}
	}
	return []string{}
}

func (dir LogicalDir) Read(path *Path) (io.ReadCloser, error) {
	locID := path.NextDir()
	if locID == "" {
		if path.Done() {
			return nil, NewError(ErrBadAction, path)
		}
		return nil, NewError(ErrNotFound, path)
	}
	
	if _, ok := dir[locID]; ok {
		ds := dir[locID]
		if ds.Exists(path) {
			return ds.Read(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (dir LogicalDir) ReadAll(path *Path) ([]byte, error) {
	locID := path.NextDir()
	if locID == "" {
		if path.Done() {
			return nil, NewError(ErrBadAction, path)
		}
		return nil, NewError(ErrNotFound, path)
	}
	
	if _, ok := dir[locID]; ok {
		ds := dir[locID]
		if ds.Exists(path) {
			return ds.ReadAll(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (dir LogicalDir) Write(path *Path) (io.WriteCloser, error) {
	locID := path.NextDir()
	if locID == "" {
		if path.Done() {
			return nil, NewError(ErrBadAction, path)
		}
		return nil, NewError(ErrNotFound, path)
	}
	
	if _, ok := dir[locID]; ok {
		ds := dir[locID]
		if ds.Exists(path) {
			return ds.Write(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (dir LogicalDir) WriteAll(path *Path, content []byte) error {
	locID := path.NextDir()
	if locID == "" {
		if path.Done() {
			return NewError(ErrBadAction, path)
		}
		return NewError(ErrNotFound, path)
	}
	
	if _, ok := dir[locID]; ok {
		ds := dir[locID]
		if ds.Exists(path) {
			return ds.WriteAll(path, content)
		}
	}
	return NewError(ErrNotFound, path)
}
