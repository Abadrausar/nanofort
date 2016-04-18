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

// Multiplexer is a simple DataSource that allows mapping several DataSources to a single location.
// Most actions are carried out on the first DataSource possible, for example Read will operate on
// the first DataSource that has the requested resource (Exists returns true).
type Multiplexer []DataSource

func NewMultiplexer(sources ...DataSource) DataSource {
	return Multiplexer(sources)
}

func (mp Multiplexer) IsDir(path *Path) bool {
	for x := range mp {
		if mp[x].Exists(path) {
			return mp[x].IsDir(path)
		}
	}
	return false
}

func (mp Multiplexer) Exists(path *Path) bool {
	for x := range mp {
		if mp[x].Exists(path) {
			return true
		}
	}
	return false
}

func (mp Multiplexer) Delete(path *Path) error {
	for x := range mp {
		if mp[x].Exists(path) {
			return mp[x].Delete(path)
		}
	}
	return NewError(ErrNotFound, path)
}

func (mp Multiplexer) Create(path *Path) error {
	for x := range mp {
		if mp[x].Exists(path) {
			return nil
		}
	}
	for x := range mp {
		err := mp[x].Create(path)
		if perr, ok := err.(*Error); ok && (perr.Typ == ErrReadOnly || perr.Typ == ErrBadAction) {
			continue
		}
	}
	return NewError(ErrReadOnly, path)
}

func (mp Multiplexer) ListDir(path *Path) []string {
	tmp := make(map[string]bool)
	for x := range mp {
		if mp[x].Exists(path) {
			for _, y := range mp[x].ListDir(path) {
				tmp[y] = true
			}
		}
	}
	
	rtn := make([]string, len(tmp))
	i := 0
	for x := range tmp {
		rtn[i] = x
		i++
	}
	
	sort.Strings(rtn)
	return rtn
}

func (mp Multiplexer) ListFile(path *Path) []string {
	tmp := make(map[string]bool)
	for x := range mp {
		if mp[x].Exists(path) {
			for _, y := range mp[x].ListFile(path) {
				tmp[y] = true
			}
		}
	}
	
	rtn := make([]string, len(tmp))
	i := 0
	for x := range tmp {
		rtn[i] = x
		i++
	}
	
	sort.Strings(rtn)
	return rtn
}

func (mp Multiplexer) Read(path *Path) (io.ReadCloser, error) {
	for x := range mp {
		if mp[x].Exists(path) {
			return mp[x].Read(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (mp Multiplexer) ReadAll(path *Path) ([]byte, error) {
	for x := range mp {
		if mp[x].Exists(path) {
			return mp[x].ReadAll(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (mp Multiplexer) Write(path *Path) (io.WriteCloser, error) {
	for x := range mp {
		if mp[x].Exists(path) {
			return mp[x].Write(path)
		}
	}
	return nil, NewError(ErrNotFound, path)
}

func (mp Multiplexer) WriteAll(path *Path, content []byte) error {
	for x := range mp {
		if mp[x].Exists(path) {
			return mp[x].WriteAll(path, content)
		}
	}
	return NewError(ErrNotFound, path)
}
