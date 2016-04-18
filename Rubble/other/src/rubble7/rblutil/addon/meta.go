/*
Copyright 2013-2016 by Milo Christiansen

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

package addon

import "html/template"

// Meta is used to store meta-data for an Addon.
type Meta struct {
	// Addon tags. Some have hardcoded handling, but you can have user defined tags as well.
	Tags map[string]bool

	// The addon name.
	Name string

	// A one line addon description.
	// For use by user interfaces (not used directly by Rubble).
	Header template.HTML

	// A longer addon description, may be as long as you like.
	// If Header is an adequate description leave this empty.
	// For use by user interfaces (not used directly by Rubble).
	Description template.HTML
	DescFile    string

	// Addon names for addons that are automatically activated when this addon is active.
	Activates []string

	// Addon names for addons that are incompatible with this addon.
	Incompatible []string

	// Extra addon packs to download. These packs are downloaded regardless of if they contain addons that are needed,
	// (as Rubble has no way to know what addons are needed at load time). The key is the file name to save the pack
	// under (this must be globally unique, as packs with the same name are assumed to be the same pack), the value
	// is the URL to load the pack from.
	// If multiple addon request the same pack which one is used is undefined.
	Download map[string]string

	// Configuration variables used by this addon and their default values.
	Vars map[string]*MetaVar

	// A load priority number, addons are sorted from lowest to highest.
	// Negative numbers mean "use parent", if there is no parent then the
	// absolute value is used.
	// Default is -100.
	LoadPriority int

	// Author and Version are inheritable from the parent if they are set to " " (the default).
	Author  string
	Version string
}

// MetaVar is used to store meta data about a config var.
type MetaVar struct {
	Name string // A user friendly name/description, may be empty.

	// A list of possible values for this variable. Index 0 is the default value.
	Values []string
}

func NewMeta() *Meta {
	return &Meta{
		Tags:         make(map[string]bool),
		Header:       template.HTML("The author of this addon did not provide addon meta-data."),
		Activates:    []string{},
		Incompatible: []string{},
		Vars:         make(map[string]*MetaVar),

		LoadPriority: -100,
		Author:       " ",
		Version:      " ",
	}
}

// Try making a string ok to put in double quotes...
func escapeString(input string) string {
	in := []byte(input)
	out := make([]byte, 0, len(in))
	for i := range in {
		switch in[i] {
		case '\n':
			out = append(out, "\\n"...)
		case '\r':
			out = append(out, "\\r"...)
		case '\t':
			out = append(out, "\\t"...)
		case '"':
			out = append(out, "\\\""...)
		case '\\':
			out = append(out, "\\\\"...)
		default:
			out = append(out, in[i])
		}
	}
	return string(out)
}
