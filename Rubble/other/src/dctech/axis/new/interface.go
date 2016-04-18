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

// AXIS VFS, a simple virtual file system API.
// 
// AXIS is based on three simple interfaces and a set of API functions that operate on these interfaces.
// Client use the provided implementations of these interfaces (or provide their own custom implementations)
// to create "file systems" that may be used for OS-independent file IO.
// 
// AXIS was originally written to allow files inside of archives to be handled with exactly the same API as
// used for files inside of directories, but it has since grown to allow "logical" files and directories as
// well as "multiplexing" multiple items to the same location (to, for example, make two directories look
// and act like one). These properties make AXIS perfect for handling data and configuration files for any
// program where flexibility is important, the program does not need to know where its files are actually
// located, it simply needs them to be at a certain place in it's AXIS file system. Changing where a program
// loads it's files from is then as simple as changing the code that initializes the file system,
// 
// AXIS uses special two part paths, with the first part being a series of zero or more "location IDs"
// separated by colons and the second part being a traditional slash separated path.
// 
//	locid1:locid2:dir1/dir2/dir3/file.ext
// 
// Whether paths have location ID at all is determined by how the file system is assembled. By convention
// location IDs are used to refer to different "segments" of the file system. Think of locations as Windows
// drive letters. Of course if you dislike the "locations" idea you can use directory names for everything...
// 
// AXIS VFS "officially" stands for Absurdly eXtremely Incredibly Simple Virtual File System (adjectives are
// good for making cool acronyms!). If you think the name is stupid (it is) you can just call it AXIS and
// forget what it is supposed to mean, after all the "official" name is more of a joke than anything...
package axis

import "io"
import "io/ioutil"
import "os"

// Mode is used for flags to the write method of File and Gateway.
type Mode int

// Modes for the Write methods of the File and Gateway interfaces.
// Multiple modes may be used by ORing them together.
const (
	Default Mode = 0x0 // Open the item for appending.
	Create Mode = 0x1 // Create the item if it does not already exist.
	Truncate Mode = 0x2 // Truncate the item instead of appending.
)

// DataSource represents an item in an AXIS filesystem.
// DataSources MUST implement one of File, Dir, or Gateway.
// Accessing data stored in DataSources directly is strongly discouraged, use the provided access API.
type DataSource interface {}

// File is the interface a DataSource that exposes data for reading and/or writing must implement.
type File interface {
	// Read opens an AXIS file for reading and returns the result and any error that may have happened.
	Read() (io.ReadCloser, error)
	
	// Write opens an AXIS file for writing and returns the result and any error that may have happened.
	Write(m Mode) (io.WriteCloser, error)
	
	Size() int64
}

// Dir represents a grouping of DataSources.
// Dir is used for logical directories and such like.
type Dir interface {
	// List lists the children of the Dir.
	// Implementations should return a zero length slice (or nil) if there are no children.
	List() []string
	
	// Mount a DataSource at the specified location, the name will be used as a location ID or
	// file/directory name, depending on context.
	// 
	// Mounting a new DataSource onto an existing location should replace the current DataSource.
	// 
	// Passing nil for the DataSource is OK, in this case the existing DataSource (if any) should
	// simply be unmounted.
	Mount(name string, ds DataSource)
	
	// Get attempts to lookup a child DataSource by path.
	Get(name string) (DataSource, error)
	
	// Location returns true if the items in this Dir should be accessed by location IDs instead of directory names.
	Location() bool
}

// Gateway is the interface a DataSource that acts as an intermediary to a non-AXIS file system must implement.
// Paths passed to a Gateway will be slash separated on all OSs.
type Gateway interface {
	// Info returns information about the item at the given path.
	// If the item does not exist this should return nil.
	Info(path string) os.FileInfo
	
	// List lists all the items contained in the item at the given path.
	// This should return a zero length slice (or nil) for files and empty directories.
	List(path string) []string
	
	// Delete tries to delete a file or directory through the Gateway, this should remove the backing
	// file/directory from the OS's file system (or from whatever the Gateway is exposing).
	// 
	// If for some reason the Gateway does not allow deleting the item then this should return an error
	// with the type ErrReadOnly or ErrBadAction (whichever one make more sense).
	Delete(path string) error
	
	// Read opens an item for reading and returns the result and any error that may have happened.
	// Does not have to handle the Create mode (in fact this *should not* handle the Create mode).
	Read(path string) (io.ReadCloser, error)
	
	// Write opens an item for writing and returns the result and any error that may have happened.
	Write(path string, m Mode) (io.WriteCloser, error)
}

// ================================================================
// Public API
// ================================================================

// Helper
func getByPath(ds DataSource, p *path) (DataSource, string, error) {
	var err error
	
	dir, dok := ds.(Dir)
	_, gok := ds.(Gateway)
	
	// Location IDs must refer to items in Dirs.
	loc := p.NextLoc()
	for loc != "" {
		if !dok || !dir.Location() {
			return nil, "", &Error{Path: p.String(), Typ: ErrNotFound}
		}
		
		ds, err = dir.Get(loc)
		if err != nil {
			return nil, "", wrapError(err, p.String())
		}
		dir, dok = ds.(Dir)
		_, gok = ds.(Gateway)
		loc = p.NextLoc()
	}

	loc = p.NextDir()
	for loc != "" {
		if gok {
			return ds, p.Remainder(), nil
		}
		if !dok || dir.Location() {
			return nil, "", &Error{Path: p.String(), Typ: ErrNotFound}
		}
		
		ds, err = dir.Get(loc)
		if err != nil {
			return nil, "", wrapError(err, p.String())
		}
		dir, dok = ds.(Dir)
		_, gok = ds.(Gateway)
		loc = p.NextDir()
	}
	
	return ds, "", nil
}

// IsDir returns true if the path points to a directory (either an AXIS Dir or a directory exposed by a Gateway)
func IsDir(ds DataSource, path string) bool {
	nds, rem, err := getByPath(ds, newPath(path))
	if err != nil {
		return false
	}
	if _, ok := nds.(Gateway); rem != "" || ok {
		gate := nds.(Gateway)
		return gate.Info(rem).IsDir()
	}
	_, dir := nds.(Dir)
	_, gate := nds.(Gateway)
	return dir || gate
}

// Exists returns true if an AXIS path point to a valid item.
func Exists(ds DataSource, path string) bool {
	nds, rem, err := getByPath(ds, newPath(path))
	if err != nil {
		return false
	}
	if _, ok := nds.(Gateway); rem != "" || ok {
		gate := nds.(Gateway)
		return gate.Info(rem) != nil
	}
	return true
}

// Delete tries to delete an item, this removes the backing file/directory from
// the OS's file system if the item is so backed, otherwise it unmounts the item.
// 
// Some items might not be deletable.
func Delete(ds DataSource, path string) error {
	p := newPath(path)
	ok := p.HoldLast()
	if !ok {
		// Empty path
		return &Error{Path: path, Typ: ErrBadAction}
	}
	
	nds, rem, err := getByPath(ds, p)
	if err != nil {
		return err // Already wrapped
	}
	if _, ok := nds.(Gateway); rem != "" || ok {
		gate := nds.(Gateway)
		err := gate.Delete(rem)
		return wrapError(err, path)
	}
	dir, ok := nds.(Dir)
	if !ok {
		return &Error{Path: p.String(), Typ: ErrNotFound}
	}
	dir.Mount(p.Held, nil)
	return nil
}

// ListDir lists the directories at a path.
// Returns nil if there are no directories or the given item does not exist.
func ListDir(ds DataSource, path string) []string {
	nds, rem, err := getByPath(ds, newPath(path))
	if err != nil {
		return nil
	}
	var rtn []string
	if _, ok := nds.(Gateway); rem != "" || ok {
		gate := nds.(Gateway)
		items := gate.List(rem)
		for _, item := range items {
			if gate.Info(rem+"/"+item).IsDir() {
				rtn = append(rtn, item)
			}
		}
	} else if dir, ok := nds.(Dir); ok {
		items := dir.List()
		for _, item := range items {
			tmp, err := dir.Get(item)
			if err != nil {
				continue
			}
			_, d := tmp.(Dir)
			_, g := tmp.(Gateway)
			if d || g {
				rtn = append(rtn, item)
			}
		}
	}
	return rtn
}

// ListFile lists the files in a AXIS directory at a path.
// Returns nil if there are no files or the given item does not exist.
func ListFile(ds DataSource, path string) []string {
	nds, rem, err := getByPath(ds, newPath(path))
	if err != nil {
		return nil
	}
	var rtn []string
	if _, ok := nds.(Gateway); rem != "" || ok {
		gate := nds.(Gateway)
		items := gate.List(rem)
		for _, item := range items {
			if !gate.Info(rem+"/"+item).IsDir() {
				rtn = append(rtn, item)
			}
		}
	} else if dir, ok := nds.(Dir); ok {
		items := dir.List()
		for _, item := range items {
			tmp, err := dir.Get(item)
			if err != nil {
				continue
			}
			_, d := tmp.(Dir)
			_, g := tmp.(Gateway)
			if d || g {
				rtn = append(rtn, item)
			}
		}
	}
	return rtn
}

// Size attempts to get the size of the item at the given path.
// Returns -1 if the size is not available.
func Size(ds DataSource, path string) int64 {
	nds, rem, err := getByPath(ds, newPath(path))
	if err != nil {
		return -1
	}
	if _, ok := nds.(Gateway); rem != "" || ok {
		gate := nds.(Gateway)
		info := gate.Info(rem)
		if info == nil {
			return -1
		}
		return info.Size()
	}
	
	f, ok := nds.(File)
	if ok {
		return f.Size()
	}
	return -1
}

// Read opens a file for reading and returns the result and any error that may have happened.
func Read(ds DataSource, path string) (io.ReadCloser, error) {
	nds, rem, err := getByPath(ds, newPath(path))
	if err != nil {
		return nil, err // Already wrapped
	}
	if _, ok := nds.(Gateway); rem != "" || ok {
		gate := nds.(Gateway)
		rc, err := gate.Read(rem)
		return rc, wrapError(err, path)
	}
	
	f, ok := nds.(File)
	if !ok {
		return nil, wrapError(err, path)
	}
	rc, err := f.Read()
	return rc, wrapError(err, path)
}

// ReadAll attempts to read a file into memory and returns the result plus any error that may
// have happened.
// This is a simple convenience function wrapping Read.
func ReadAll(ds DataSource, path string) ([]byte, error) {
	reader, err := Read(ds, path)
	if err != nil {
		return nil, wrapError(err, path)
	}
	defer reader.Close()
	
	content, err := ioutil.ReadAll(reader)
	return content, wrapError(err, path)
}

// Write opens a file for writing and returns the result and any error that may have happened.
func Write(ds DataSource, path string, m Mode) (io.WriteCloser, error) {
	nds, rem, err := getByPath(ds, newPath(path))
	if err != nil {
		return nil, wrapError(err, path)
	}
	if _, ok := nds.(Gateway); rem != "" || ok {
		gate := nds.(Gateway)
		wc, err := gate.Write(rem, m)
		return wc, wrapError(err, path)
	}
	
	f, ok := nds.(File)
	if !ok {
		return nil, &Error{Path: path, Typ: ErrBadAction}
	}
	wc, err := f.Write(m)
	return wc, wrapError(err, path)
}

// WriteAll attempts to replace the contents of a file with the data given, any error
// is returned.
// This is a simple convenience function wrapping Write.
func WriteAll(ds DataSource, path string, content []byte) error {
	writer, err := Write(ds, path, Truncate)
	if err != nil {
		return err // Already wrapped
	}
	defer writer.Close()
	
	_, err = writer.Write(content)
	return wrapError(err, path)
}
