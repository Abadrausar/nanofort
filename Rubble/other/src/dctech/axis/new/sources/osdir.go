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

import gopath "path"
import "path/filepath"
import "os"
import "io"
import "io/ioutil"

import "dctech/axis/new"

type osDir struct {
	path string
	rw bool
}

// NewOSDir creates a new Gateway for reading (and possible writing) to an OS directory and its files.
// rw controls if the Gateway allows deleting and writing it's contained files.
func NewOSDir(path string, rw bool) axis.Gateway {
	dir := new(osDir)
	dir.rw = rw
	dir.path = filepath.ToSlash(path)
	return dir
}

func (dir *osDir) Info(path string) os.FileInfo {
	info, err := os.Stat(dir.path + "/" + path)
	if err != nil {
		return nil
	}
	return info
}

func (dir *osDir) List(path string) []string {
	files, err := ioutil.ReadDir(dir.path + "/" + path)
	if err != nil {
		return nil
	}
	
	rtn := make([]string, 0, len(files))
	for _, file := range files {
		rtn = append(rtn, file.Name())
	}
	return rtn
}

func (dir *osDir) Delete(path string) error {
	if !dir.rw {
		return axis.NewError(axis.ErrReadOnly)
	}
	return os.Remove(dir.path + "/" + path)
}

func (dir *osDir) Read(path string) (io.ReadCloser, error) {
	return os.Open(dir.path + "/" + path)
}

func (dir *osDir) Write(path string, m axis.Mode) (io.WriteCloser, error) {
	if !dir.rw {
		return nil, axis.NewError(axis.ErrReadOnly)
	}
	
	_, err := os.Stat(dir.path + "/" + path)
	if m&axis.Create != 0 && err != nil {
		err := os.MkdirAll(gopath.Dir(dir.path + "/" + path), 0666)
		if err != nil {
			return nil, err
		}
		return os.Create(dir.path + "/" + path)
	}
	
	if m&axis.Truncate != 0 {
		return os.Create(dir.path + "/" + path)
	}
	
	return os.OpenFile(dir.path + "/" + path, os.O_WRONLY, 0)
}
