/*
Copyright 2015-2016 by Milo Christiansen

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

import "rubble7/rblutil"

// DefaultFirst is a map of the first part of a two part extension to the associated file tag(s).
var DefaultFirst = map[string][]string{
	// Rubble Scripts
	".pre":  {"PreScript"},
	".post": {"PostScript"},
	".mod":  {"ModuleScript"},
	".load": {"LoadScript", "GlobalFile"},
	".init": {"InitScript", "GlobalFile"},

	// Other Rubble Files
	".aux":      {"MiscAxillary"},
	".tset":     {"TileSet"},
	".text":     {"AUXTextFile"},
	".graphics": {"CreatureGraphics"},

	// DFHack Scripts
	".dfcom":  {"CommandScript", "DFHack"},
	".dfmod":  {"ModuleScript", "DFHack"},
	".dfload": {"LoadScript", "DFHack"},
}

// DefaultLast is a map of the last part of a two part extension to the associated file tag(s).
var DefaultLast = map[string][]string{
	// Language Extensions
	".lua": {"LangLua"},
	".luab": {"LangLuaBin"},
	".cl": {"LangLuaC"},
	".rb":  {"LangRuby"},

	// Other File Types
	".test":  {"TemplateTest"},
	".txt":   {"RawFile"},
	".rbl":   {"RawFile", "NoWrite"},
	".rules": {"MergeRules"},
	".twbt":  {"TWBTOverride"},
	".png":   {"ImagePNG"},
	".bmp":   {"ImageBMP"},
}

// GetFileTagsAdv uses the given tag maps to find the file tags for a file with the given name.
// The returned slice of tags is yours to keep.
func GetFileTagsAdv(first, last map[string][]string, name string) []string {
	f, l := rblutil.GetExtParts(name)
	fl := first[f]
	ll := last[l]

	return append(append(make([]string, 0, len(fl)+len(ll)), fl...), ll...)
}

// GetFileTags uses the default tag maps to find the file tags for a file with the given name.
// The returned slice of tags is yours to keep.
func GetFileTags(name string) []string {
	return GetFileTagsAdv(DefaultFirst, DefaultLast, name)
}
