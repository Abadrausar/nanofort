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

import "path"
import "path/filepath"
import "os"
import "io"

import "dctech/axis/new"

type osFile struct {
	path string
	rw bool
}

// NewOSFile creates a new OS AXIS file interface.
// rw controls if Delete, Create, Write, and WriteAll work.
func NewOSFile(path string, rw bool) axis.File {
	file := new(osFile)
	file.rw = rw
	file.path = filepath.ToSlash(path)
	return file
}

func (file *osFile) Size() int64 {
	s, err := os.Lstat(file.path)
	if err != nil {
		return -1
	}
	return s.Size()
}

func (file *osFile) Read() (io.ReadCloser, error) {
	return os.Open(file.path)
}

func (file *osFile) Write(m axis.Mode) (io.WriteCloser, error) {
	if !file.rw {
		return nil, axis.NewError(axis.ErrReadOnly)
	}
	
	_, err := os.Stat(file.path)
	if m&axis.Create != 0 && err != nil {
		err := os.MkdirAll(path.Dir(file.path), 0666)
		if err != nil {
			return nil, err
		}
		return os.Create(file.path)
	}
	
	if m&axis.Truncate != 0 {
		return os.Create(file.path)
	}
	
	return os.OpenFile(file.path, os.O_WRONLY, 0)
}
