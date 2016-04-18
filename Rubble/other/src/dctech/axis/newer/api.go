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
// Clients use the provided implementations of these interfaces (or provide their own custom implementations)
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

// DataSource is any item that implements either File or Dir (or, more rarely, both).
type DataSource interface {}

// Dir is the interface directories (or items that act like directories) must implement.
type Dir interface {
	// Child returns a handle to the requested child item.
	// If create is true then a new item should be created to satisfy
	// the request if needed. If this is not possible then just return nil.
	// If create is false and the item does not exist then return nil.
	Child(id string, create bool) DataSource
	
	// Delete the given child item.
	Delete(id string) error
	
	// List all the children of this Dir.
	List() []string
}

// File is the interface files (or items that act like files) must implement.
type File interface {
	// Read opens an AXIS file for reading and returns the result and any error that may have happened.
	Read() (io.ReadCloser, error)
	
	// Write opens an AXIS file for writing and returns the result and any error that may have happened.
	// Any existing contents the file may have are truncated.
	Write() (io.WriteCloser, error)
	
	// Append is exactly like Write, except the file is not truncated before writing. 
	Append() (io.WriteCloser, error)
	
	// Size returns the size of the file or -1 if the value could not be retrieved.
	Size() int64
}

type source struct {
	mp []string
	ds DataSource
}


// FileSystem is the center of an AXIS setup.
// 
// FileSystems have two halves: A "read" half and a "write" half. Any action that changes something (Write, Delete, etc)
// is carried out on the "write" half and any action that involves reading existing information is carried out on the
// "read" half. For this reason most DataSources are mounted on both halves or just the read half, never on just the
// write half.
// 
// Each half also has "locations", separate sections of the file system used for organization. You can think of locations
// as drive letters, but more flexible. Each path must begin with a location ID separated from the path with a single
// colon. If a path has multiple colons then last one is the location separator.
// 
// After the location ID an AXIS path is a simple slash separated path. AXIS does not attach any special meaning to any
// characters except slash and colon, but "." and ".." are illegal as directory and file names. Any path elements that
// will eventually be given to the OS may have more limitations.
// 
// DataSources may be mounted wherever you want, you do not need to mount them at an existing location or path.
type FileSystem struct {
	r map[string][]*source
	w map[string][]*source
}

// NewFileSystem returns a new AXIS FileSystem, initialized and ready to use.
func NewFileSystem() *FileSystem {
	return &FileSystem{
		r: map[string][]*source{},
		w: map[string][]*source{},
	}
}

// Mount a given DataSource onto the FileSystem at the given location/path.
// If rw is true the DataSource is mounted for writing as well as reading.
// The mount path should NOT have a trailing slash! If the mount path is just
// a location ID with no path part it DOES need a trailing colon.
func (fs *FileSystem) Mount(path string, ds DataSource, rw bool) error {
	ok, loc, dirs := validatePath(path)
	if !ok {
		return &Error{Path: path, Typ: ErrPathRelative}
	}
	
	src := &source{
		mp: dirs,
		ds: ds,
	}
	fs.r[loc] = append(fs.r[loc], src)
	if rw {
		fs.w[loc] = append(fs.w[loc], src)
	}
	return nil
}

// Returns a list of mount point parts that begin with the given path.
// The names returned act like valid directory names for most purposes.
func (fs *FileSystem) mountSubset(path string, r bool) []string {
	ok, loc, dirs := validatePath(path)
	if !ok {
		return nil
	}
	
	var sources []*source
	if r {
		sources, ok = fs.r[loc]
	} else {
		sources, ok = fs.w[loc]
	}
	if !ok {
		return nil
	}
	
	var rtn []string
	next:
	for _, src := range sources {
		if len(src.mp) <= len(dirs) {
			continue
		}
		
		for i := range dirs {
			if src.mp[i] != dirs[i] {
				continue next
			}
		}
		rtn = append(rtn, src.mp[len(dirs)])
	}
	return rtn
}

// Internal.
func (fs *FileSystem) getDSAt(path string, create, r bool) (DataSource, error) {
	ok, loc, dirs := validatePath(path)
	if !ok {
		return nil, &Error{Path: path, Typ: ErrPathRelative}
	}
	
	var sources []*source
	if r {
		sources, ok = fs.r[loc]
	} else {
		sources, ok = fs.w[loc]
	}
	if !ok {
		return nil, &Error{Path: path, Typ: ErrNotFound}
	}
	
	next:
	for _, src := range sources {
		// First make sure that the requested path is a superset of the current source's mount point.
		i := 0
		for ; i < len(src.mp); i++ {
			if i >= len(dirs) || src.mp[i] != dirs[i] {
				continue next
			}
		}
		
		// Then try to get a child item from the source that matches the remainder of the path.
		ds := src.ds
		var pds DataSource
		for ; i < len(dirs); i++ {
			pds = ds
			pdir, ok := pds.(Dir)
			if !ok {
				continue next
			}
			ds = pdir.Child(dirs[i], create)
			if ds == nil {
				continue next
			}
		}
		
		return ds, nil
	}
	
	return nil, &Error{Path: path, Typ: ErrNotFound}
}

// Exists returns true if the path points to a valid DataSource.
func (fs *FileSystem) Exists(path string) bool {
	_, err := fs.getDSAt(path, false, true)
	if err != nil {
		// Report that the path exists (even if it doesn't really) when it is a subset of one or more mount points.
		return len(fs.mountSubset(path, true)) != 0
	}
	return true
}

// IsDir returns true if the path points to a valid Dir.
func (fs *FileSystem) IsDir(path string) bool {
	ds, err := fs.getDSAt(path, false, true)
	if err != nil {
		// Report that the path is a directory (even if it isn't really) when it is a subset of one or more mount points.
		return len(fs.mountSubset(path, true)) != 0
	}
	
	_, ok := ds.(Dir)
	return ok
}

// Size returns the size of the File at the given path. If the object is not a File or the path is invalid -1 is returned.
func (fs *FileSystem) Size(path string) int64 {
	ds, err := fs.getDSAt(path, false, true)
	if err != nil {
		return -1
	}
	
	f, ok := ds.(File)
	if !ok {
		return -1
	}
	return f.Size()
}

// Delete attempts to delete the item at the given path. This may or may not work. Deleting is always carried out on the
// write portion of the FileSystem, objects on the read portion will not be effected unless they are also mounted for
// writing. Only the first item found is deleted.
func (fs *FileSystem) Delete(path string) error {
	npath, last := trimLastPath(path)
	
	ds, err := fs.getDSAt(npath, false, false)
	if err != nil {
		return err
	}
	
	d, ok := ds.(Dir)
	if !ok {
		return &Error{Path: path, Typ: ErrNotFound}
	}
	return wrapError(d.Delete(last), path)
}

// List returns a slice of the names of all the items available in the given Dir. If the item at the path is not a Dir
// this returns nil.
func (fs *FileSystem) List(path string) []string {
	ds, err := fs.getDSAt(path, false, true)
	if err != nil {
		// Treat the path like a directory of directories if it is a mount point subset.
		return fs.mountSubset(path, true)
	}
	
	d, ok := ds.(Dir)
	if !ok {
		return fs.mountSubset(path, true)
	}
	// ??? I think? This should only matter in cases where mount points overlap data sources.
	return append(d.List(), fs.mountSubset(path, true)...)
}

// ListDirs returns a slice of all the items that are Dirs in the given Dir. If the item at the path is not a Dir
// this returns nil.
func (fs *FileSystem) ListDirs(path string) []string {
	ds, err := fs.getDSAt(path, false, true)
	if err != nil {
		// Treat the path like a directory of directories if it is a mount point subset.
		return fs.mountSubset(path, true)
	}
	
	d, ok := ds.(Dir)
	if !ok {
		return fs.mountSubset(path, true)
	}
	children := d.List()
	
	var rtn []string
	for _, child := range children {
		cds := d.Child(child, false)
		if _, ok := cds.(Dir); ok {
			rtn = append(rtn, child)
		}
	}
	// ??? I think? This should only matter in cases where mount points overlap data sources.
	return append(rtn, fs.mountSubset(path, true)...)
}

// ListFiles returns a slice of all the items that are Files in the given Dir. If the item at the path is not a Dir
// this returns nil.
func (fs *FileSystem) ListFiles(path string) []string {
	ds, err := fs.getDSAt(path, false, true)
	if err != nil {
		return nil
	}
	
	d, ok := ds.(Dir)
	if !ok {
		return nil
	}
	children := d.List()
	
	var rtn []string
	for _, child := range children {
		cds := d.Child(child, false)
		if _, ok := cds.(File); ok {
			rtn = append(rtn, child)
		}
	}
	return rtn
}

// Read opens the File at the given path for reading.
func (fs *FileSystem) Read(path string) (io.ReadCloser, error) {
	ds, err := fs.getDSAt(path, false, true)
	if err != nil {
		return nil, err
	}
	
	f, ok := ds.(File)
	if !ok {
		return nil, &Error{Path: path, Typ: ErrBadAction}
	}
	
	rc, err := f.Read()
	return rc, wrapError(err, path)
}

// ReadAll reads the File at the given path and returns it's contents.
func (fs *FileSystem) ReadAll(path string) ([]byte, error) {
	reader, err := fs.Read(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	
	content, err := ioutil.ReadAll(reader)
	return content, wrapError(err, path)
}

// Write opens the the File at the given path for writing. Any existing file contents are truncated.
func (fs *FileSystem) Write(path string) (io.WriteCloser, error) {
	ds, err := fs.getDSAt(path, true, false)
	if err != nil {
		return nil, err
	}
	
	f, ok := ds.(File)
	if !ok {
		return nil, &Error{Path: path, Typ: ErrBadAction}
	}
	
	wc, err := f.Write()
	return wc, wrapError(err, path)
}

// Append opens the file at the given path for writing. The write cursor is set beyond any existing file contents.
func (fs *FileSystem) Append(path string) (io.WriteCloser, error) {
	ds, err := fs.getDSAt(path, true, false)
	if err != nil {
		return nil, err
	}
	
	f, ok := ds.(File)
	if !ok {
		return nil, &Error{Path: path, Typ: ErrBadAction}
	}
	
	wc, err := f.Append()
	return wc, wrapError(err, path)
}

// WriteAll replace the contents of the File at the given path with the contents given.
func (fs *FileSystem) WriteAll(path string, content []byte) error {
	writer, err := fs.Write(path)
	if err != nil {
		return err
	}
	defer writer.Close()
	
	_, err = writer.Write(content)
	return wrapError(err, path)
}
