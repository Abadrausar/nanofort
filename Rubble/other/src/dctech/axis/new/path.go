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

import "path/filepath"
import "strings"

// isAbs returns true if the path is not a relative path (includes no "." or ".." parts).
func isAbs(path string) bool {
	// AXIS paths may start with an arbitrary number of location IDs that are separated by colons,
	// they have no meaning to this function, so just strip them off.
	locs := strings.Split(path, ":")
	path = locs[len(locs) - 1]
	dirs := strings.Split(path, "/")
	
	for i := range dirs {
		if dirs[i] == ".." || dirs[i] == "." {
			return false
		}
	}
	return true
}

// Path represents a parsed AXIS path.
type path struct {
	locations []string
	usedlocs int
	
	dirs []string
	useddirs int
	
	Held string
}

// NewPath parses a string into an AXIS path.
func newPath(source string) *path {
	source = strings.TrimSpace(filepath.ToSlash(source))
	if !isAbs(source) {
		return nil
	}
	
	p := new(path)
	if len(source) == 0 {
		p.locations = make([]string, 0)
		p.dirs = make([]string, 0)
	} else if source[len(source) - 1:] == ":" {
		p.locations = strings.Split(source, ":")
		p.locations = p.locations[:len(p.locations) - 1]
		p.dirs = make([]string, 0)
	} else {
		locs := strings.Split(source, ":")
		source = locs[len(locs) - 1]
		p.locations = locs[:len(locs) - 1]
		p.dirs = strings.Split(source, "/")
		if p.dirs[len(p.dirs) - 1] == "" {
			p.dirs = p.dirs[:len(p.dirs) - 1]
		}
	}
	return p
}

func (p *path) String() string {
	out := ""
	for i := range p.locations {
		out += p.locations[i] + ":"
	}
	for i := range p.dirs {
		out += p.dirs[i] + "/"
	}
	
	if len(out) != 0 {
		if out[len(out) - 1] == '/' {
			return out[:len(out) - 1]
		}
	}
	return out
}

func (p *path) HoldLast() bool {
	if len(p.dirs) > 0 {
		p.Held = p.dirs[len(p.dirs) - 1]
		p.dirs = p.dirs[:len(p.dirs) - 1]
		return true
	}
	if len(p.locations) > 0 {
		p.Held = p.locations[len(p.locations) - 1]
		p.locations = p.locations[:len(p.locations) - 1]
		return true
	}
	return false
}

// Remainder returns a string constructed from the remaining path elements.
// Location IDs are always ignored (as this is meant to be used when talking to the OS).
// The last used dir name is included int the return value.
// If the returned string would be empty returns ".".
func (p *path) Remainder() string {
	p.useddirs--
	if p.useddirs < 0 {
		p.useddirs = 0
	}
	
	if len(p.dirs) > p.useddirs {
		out := ""
		for _, v := range p.dirs[p.useddirs:] {
			out += v + "/"
		}
		if len(out) != 0 {
			if out[len(out) - 1] == '/' {
				return out[:len(out) - 1]
			}
		}
		return out
	}
	return "."
}

// NextLoc returns the next location ID in the path.
// Returns "" if no location IDs remain unused.
func (p *path) NextLoc() string {
	if len(p.locations) <= p.usedlocs {
		return ""
	}
	
	rtn := p.locations[p.usedlocs]
	p.usedlocs++
	return rtn
}

// NextDir returns the next directory name in the path (the last element may be a file name).
// Returns "" if locations remain unused or there are no directory names left.
func (p *path) NextDir() string {
	if len(p.locations) > p.usedlocs {
		return ""
	}
	
	if len(p.dirs) <= p.useddirs {
		return ""
	}
	
	rtn := p.dirs[p.useddirs]
	p.useddirs++
	return rtn
}
