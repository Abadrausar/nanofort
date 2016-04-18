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

package sources

import "sort"

import "dctech/axis/new"

type fileSystem map[string]axis.DataSource

// NewFileSystem returns a Dir implementation that serves it's mounted DataSources as locations.
func NewFileSystem() axis.Dir {
	return make(fileSystem)
}

func (fs fileSystem) List() []string {
	rtn := make([]string, 0, len(fs))
	
	for loc := range fs {
		rtn = append(rtn, loc)
	}
	
	sort.Strings(rtn)
	return rtn
}

func (fs fileSystem) Mount(name string, ds axis.DataSource) {
	if ds == nil {
		delete(fs, name)
		return
	}
	fs[name] = ds
}

func (fs fileSystem) Get(name string) (axis.DataSource, error) {
	ds, ok := fs[name]
	if !ok {
		return nil, axis.NewError(axis.ErrNotFound)
	}
	return ds, nil
}

func (fs fileSystem) Location() bool {
	return true
}
