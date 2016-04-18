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

import "io"
import "os"
import "sort"

import "dctech/axis/new"

type dirMultiplexer []axis.Dir

// NewDirMultiplexer returns a simple Dir implementation that multiplexes several other Dirs together.
// In general the first Dir (in the list) that can satisfy a request "wins".
// All the Dirs passed in should return the same value when Location is called. This is not checked,
// the value returned by the first Dir in this list is always the one used.
func NewDirMultiplexer(sources ...axis.Dir) axis.Dir {
	return dirMultiplexer(sources)
}

func (mp dirMultiplexer) List() []string {
	tmp := make(map[string]bool)
	for x := range mp {
		for _, y := range mp[x].List() {
			tmp[y] = true
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

func (mp dirMultiplexer) Mount(name string, ds axis.DataSource) {
	if len(mp) < 1 {
		return // Just in case
	}
	mp[0].Mount(name, ds)
}

func (mp dirMultiplexer) Get(name string) (axis.DataSource, error) {
	var ds axis.DataSource
	err := axis.NewError(axis.ErrNotFound)
	for x := range mp {
		if ds, err = mp[x].Get(name); err == nil {
			return ds, nil
		}
	}
	return nil, err
}

func (mp dirMultiplexer) Location() bool {
	if len(mp) < 1 {
		return false // Just in case
	}
	return mp[0].Location()
}

type gateMultiplexer []axis.Gateway

// NewGateMultiplexer returns a simple Gateway implementation that multiplexes several other Gateways together.
// In general the first Gateway (in the list) that can satisfy a request "wins".
func NewGateMultiplexer(sources ...axis.Gateway) axis.Gateway {
	return gateMultiplexer(sources)
}

func (mp gateMultiplexer) Info(path string) os.FileInfo {
	var info os.FileInfo
	for x := range mp {
		if info = mp[x].Info(path); info != nil {
			return info
		}
	}
	return nil
}

func (mp gateMultiplexer) List(path string) []string {
	tmp := make(map[string]bool)
	for x := range mp {
		for _, y := range mp[x].List(path) {
			tmp[y] = true
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

// Return first error found. Try to delete every occurrence of the item.
func (mp gateMultiplexer) Delete(path string) error {
	found := false
	for x := range mp {
		if info := mp[x].Info(path); info != nil {
			err := mp[x].Delete(path)
			if err != nil {
				return err
			}
			found = true
		}
	}
	if !found {
		return axis.NewError(axis.ErrNotFound)
	}
	return nil
}

func (mp gateMultiplexer) Read(path string) (io.ReadCloser, error) {
	for x := range mp {
		if info := mp[x].Info(path); info != nil {
			return mp[x].Read(path)
		}
	}
	return nil, axis.NewError(axis.ErrNotFound)
}

func (mp gateMultiplexer) Write(path string, m axis.Mode) (io.WriteCloser, error) {
	if len(mp) < 1 {
		return nil, axis.NewError(axis.ErrBadAction) // Just in case
	}
	return mp[0].Write(path, m)
}
