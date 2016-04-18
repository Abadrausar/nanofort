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

import "dctech/axis"
import "dctech/axis/axiszip"

import "rubble7/rblutil"
import "rubble7/rblutil/errors"

import "sort"
import "encoding/json"
import "html"
import "html/template"
import "strings"
import "io/ioutil"
import "net/http"

import "regexp"

// Load loads addons from subdirectories of the following locations in the given DataSource:
//	addons:dir:
//	addons:zip:
// In addition global files are loaded from the following location:
//	addons:dir:
func Load(fs axis.Collection, log *rblutil.Logger) (addons []*Addon, globals *FileList) {
	log.Separator()
	log.Println("Loading Addons...")

	globals = NewFileList()
	
	log.Println("  Mounting Local Zip Files...")
	aFS, err := axis.GetChild(fs, "addons:")
	if err != nil {
		errors.RaiseWrappedError("Preparing FS for mounts:", err)
	}
	addonsFS := aFS.(axis.Collection)
	err = axis.Delete(addonsFS, "zip:")
	if err != nil {
		errors.RaiseWrappedError("Preparing FS for mounts:", err)
	}

	zipFS := axis.NewLogicalDir()

	for _, filename := range axis.ListFile(fs, "addons:dir:") {
		if strings.HasSuffix(filename, ".zip") {
			log.Println("    " + filename)

			content, err := axis.ReadAll(fs, "addons:dir:"+filename)
			if err != nil {
				errors.RaiseWrappedError("While mounting zip file:", err)
			}

			fs, err := axiszip.NewRaw(content)
			if err != nil {
				errors.RaiseWrappedError("While mounting zip file:", err)
			}
			zipFS.Mount(rblutil.StripExt(filename), fs)
		}
	}
	addonsFS.Mount("zip", zipFS)
	
	log.Println("  Downloading and Mounting Zip Files from .webload Files...")
	client := new(http.Client)

	for _, filename := range axis.ListFile(fs, "addons:dir:") {
		if strings.HasSuffix(filename, ".webload") {
			name := rblutil.StripExt(filename)
			
			u, err := axis.ReadAll(fs, "addons:dir:"+filename)
			if err != nil {
				log.Println("    Error reading .webload file:", err, "File:", name)
				continue
			}
			url := strings.TrimSpace(string(u))
	
			r, err := client.Head(url)
			if err != nil {
				log.Println("    Error making HEAD request:", err, "For file:", name)
				continue
			}
	
			content, err := axis.ReadAll(fs, "addons:dir:"+name+".zip")
			if err == nil {
				// Too bad HTTP does not seem to send a checksum
				if r.ContentLength == int64(len(content)) {
					log.Println("    " + name + ": Our copy is up to date (I think).")
					continue
				}
				
				err = axis.Delete(fs, "addons:dir:"+name+".zip")
				if err != nil {
					log.Println("      Download Error:", err)
					continue
				}
			}
			log.Println("    " + name + ": Downloading.")
	
			r, err = client.Get(url)
			if err != nil {
				log.Println("      Download Error:", err)
				continue
			}
	
			content, err = ioutil.ReadAll(r.Body)
			r.Body.Close()
			if err != nil {
				log.Println("      Download Error:", err)
				continue
			}
			
			if !axis.Exists(fs, "addons:dir:"+name+".zip") {
				err := axis.Create(fs, "addons:dir:"+name+".zip")
				if err != nil {
					log.Println("      Error saving downloaded file:", err)
					continue
				}
			}
			
			err = axis.WriteAll(fs, "addons:dir:"+name+".zip", content)
			if err != nil {
				log.Println("      Error saving downloaded file:", err)
				continue
			}
			
			zfs, err := axiszip.NewRaw(content)
			if err != nil {
				log.Println("      Error mounting zip file:", err)
				continue
			}
			zipFS.Mount(name, zfs)
		}
	}

	log.Println("  Loading Addons from Mounted Resources...")
	
	loadGlobals(fs, globals, "addons:dir:")
	pmeta := loadMeta(fs, "", "addons:dir:", nil)
	
	for _, dir := range axis.ListDir(fs, "addons:dir:") {
		addons = loadDir(fs, addons, globals, dir, "addons:dir:"+dir, pmeta)
	}

	for _, dir := range axis.ListDir(fs, "addons:zip:") {
		addons = loadDir(fs, addons, globals, dir, "addons:zip:"+dir, nil)
	}

	log.Println("  Downloading and Mounting Zip Files from Meta-Data...")
	
	packs := map[string]string{}
	for _, addon := range addons {
		for name, url := range addon.Meta.Download {
			packs[name] = url
		}
	}
	
	existing := map[string]bool{}
	for _, filename := range axis.ListFile(fs, "addons:dir:") {
		if strings.HasSuffix(filename, ".zip") {
			existing[rblutil.StripExt(filename)] = true
		}
	}
	
	for name, url := range packs {
		r, err := client.Head(url)
		if err != nil {
			log.Println("    Error making HEAD request:", err, "For file:", name)
			continue
		}

		if existing[name] {
			content, err := axis.ReadAll(fs, "addons:dir:"+name+".zip")
			if err == nil {
				// Too bad HTTP does not seem to send a checksum
				if r.ContentLength == int64(len(content)) {
					log.Println("    " + name + ": Our copy is up to date (I think).")
					continue
				}
				
				err = axis.Delete(fs, "addons:dir:"+name+".zip")
				if err != nil {
					log.Println("      Download Error:", err)
					continue
				}
			}
		}
		log.Println("    " + name + ": Downloading.")

		r, err = client.Get(url)
		if err != nil {
			log.Println("      Download Error:", err)
			continue
		}

		content, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			log.Println("      Download Error:", err)
			continue
		}

		if !axis.Exists(fs, "addons:dir:"+name+".zip") {
			err := axis.Create(fs, "addons:dir:"+name+".zip")
			if err != nil {
				log.Println("      Error saving downloaded file:", err)
				continue
			}
		}
		
		err = axis.WriteAll(fs, "addons:dir:"+name+".zip", content)
		if err != nil {
			log.Println("      Error saving downloaded file:", err)
			continue
		}
		
		zfs, err := axiszip.NewRaw(content)
		if err != nil {
			log.Println("      Error mounting zip file:", err)
			continue
		}
		zipFS.Mount(name, zfs)
		
		addons = loadDir(fs, addons, globals, name, "addons:zip:"+name, pmeta)
	}
	
	log.Println("  Checking Loaded Data...")
	
	// Sort the addon list
	sort.Sort(addonSorter(addons))

	// Add global files from addons
	globals.UpdateGlobals(addons)

	// Find any duplicates.
	dups := map[string]bool{}
	for i := range addons {
		if dups[addons[i].Meta.Name] {
			errors.RaiseError("    Duplicate addon found: \"" + addons[i].Meta.Name + "\"")
		}
		dups[addons[i].Meta.Name] = true
	}
	return addons, globals
}

// LoadRemote downloads the given zip files (if they are not already available), then loads any addons from them.
// This is meant to be run immediately after normal loading is finished.
func LoadRemote(fs axis.DataSource, packs map[string]string, addons *[]*Addon, globals *FileList) {
	
}

// fs, addons, the default name of this addon, AXIS path, meta-data of parent
func loadDir(fs axis.DataSource, addons []*Addon, globals *FileList, addonname, path string, pmeta *Meta) []*Addon {
	dirpath := path
	if path != "" {
		path += "/"
	}

	if containsParseable(fs, dirpath) {
		pmeta, addons = loadAddon(fs, addons, addonname, dirpath, pmeta)
	} else {
		loadGlobals(fs, globals, dirpath)
		pmeta = loadMeta(fs, addonname, dirpath, pmeta)
	}

	for _, dir := range axis.ListDir(fs, dirpath) {
		name := ""
		if pmeta.Name == "-" {
			name = dir
		} else {
			name = pmeta.Name + "/" + dir
		}

		addons = loadDir(fs, addons, globals, name, path+dir, pmeta)
	}

	return addons
}

// The JSON parser does not support comments. This is intolerable.
var killMetaComments = regexp.MustCompile("[\t ]*#[^\n]*\n")

// fs, the default name of this addon, AXIS path, meta-data of parent
func loadMeta(fs axis.DataSource, addonname, path string, pmeta *Meta) *Meta {
	if path != "" && path[len(path)-1] != ':' && path[len(path)-1] != '/' {
		path += "/"
	}

	var meta *Meta

	content, err := axis.ReadAll(fs, path+"addon.meta")
	if err == nil {
		meta = NewMeta()
		meta.Header = template.HTML("")

		err := json.Unmarshal(killMetaComments.ReplaceAllLiteral(content, []byte("\n")), meta)
		if err != nil {
			errors.RaiseWrappedError("While loading \""+path+"addon.meta\"", err)
		}

		// Make sure the header and description are ready to inject HTML.
		meta.Header = rblutil.MarkdownLineToHTML(string(meta.Header))
		if meta.DescFile != "" {
			content, err := axis.ReadAll(fs, path+meta.DescFile)
			if err != nil {
				errors.RaiseWrappedError("While loading description file \""+path+meta.DescFile+"\"", err)
			}

			ext := rblutil.GetExt(meta.DescFile)
			switch ext {
			case ".md", ".text":
				meta.Description = rblutil.MarkdownToHTML(string(content))
			case ".htm", ".html":
				meta.Description = template.HTML(content)
			default:
				meta.Description = template.HTML(`<pre>` + html.EscapeString(string(content)) + `</pre>`)
			}
		} else {
			meta.Description = rblutil.MarkdownToHTML(string(meta.Description))
		}

		// Handle the addon name.
		if meta.Name == "" {
			meta.Name = addonname
		} else if meta.Name != "-" {
			if len(meta.Name) > 2 && meta.Name[0] == '$' {
				if pmeta == nil || pmeta.Name == "-" {
					meta.Name = meta.Name[2:]
				} else {
					meta.Name = pmeta.Name + meta.Name[1:]
				}
			}
		}

	} else {
		// NewMeta sets any needed defaults already.
		meta = NewMeta()
		meta.Name = addonname
	}

	// Load priority inheritance
	if meta.LoadPriority < 0 {
		if pmeta == nil {
			meta.LoadPriority = -meta.LoadPriority
		} else {
			meta.LoadPriority = pmeta.LoadPriority
		}
	}

	// Author inheritance
	if meta.Author == " " {
		if pmeta == nil {
			meta.Author = ""
		} else {
			meta.Author = pmeta.Author
		}
	}

	// Version inheritance
	if meta.Version == " " {
		if pmeta == nil {
			meta.Version = ""
		} else {
			meta.Version = pmeta.Version
		}
	}

	return meta
}

func loadAddon(fs axis.DataSource, addons []*Addon, addonname, path string, pmeta *Meta) (*Meta, []*Addon) {
	addon := NewAddon(addonname, path)

	dirpath := path
	if path != "" && path[len(path)-1] != ':' {
		path += "/"
	}

	// Load Meta File
	addon.Meta = loadMeta(fs, addonname, dirpath, pmeta)

	// Load Files
	for _, filepath := range axis.ListFile(fs, dirpath) {
		content, err := axis.ReadAll(fs, path+filepath)
		if err != nil {
			panic(err)
		}

		file := NewFile(filepath, dirpath, content)
		tags := GetFileTags(filepath)
		for _, tag := range tags {
			file.Tags[tag] = true
		}

		if file.Tags["TemplateTest"] {
			addon.Meta.Tags["HasTests"] = true
		}
		if file.Tags["TileSet"] {
			addon.Meta.Tags["TileSet"] = true
		}
		if file.Tags["DFHack"] {
			if _, ok := addon.Meta.Tags["DFHack"]; !ok {
				addon.Meta.Tags["DFHack"] = true
			}
		}
		addon.Files[filepath] = file
	}
	return addon.Meta, append(addons, addon)
}

func loadGlobals(fs axis.DataSource, globals *FileList, path string) {
	dirpath := path
	if path != "" {
		path += "/"
	}

	for _, filepath := range axis.ListFile(fs, dirpath) {
		tags := GetFileTags(filepath)
		for _, tag := range tags {
			if tag == "GlobalFile" {
				content, err := axis.ReadAll(fs, path+filepath)
				if err != nil {
					panic(err)
				}

				globals.Data[filepath] = NewFile(filepath, dirpath, content)
				globals.Order = append(globals.Order, filepath)
				for _, t := range tags {
					globals.Data[filepath].Tags[t] = true
				}
				break
			}
		}
	}
	// Don't bother sorting the list here, it will be done later.
	return
}

func containsParseable(source axis.DataSource, path string) bool {
	for _, filename := range axis.ListFile(source, path) {
		tags := GetFileTags(filename)
		if len(tags) > 0 {
			global := false
			for _, tag := range tags {
				if tag == "GlobalFile" {
					global = true
					break
				}
			}
			if !global {
				return true
			}
		}
	}
	return false
}

type addonSorter []*Addon

func (a addonSorter) Len() int {
	return len(a)
}

// This function is what determines the "load order".
func (a addonSorter) Less(i, j int) bool {
	if a[i].Meta.LoadPriority == a[j].Meta.LoadPriority {
		return a[i].Meta.Name < a[j].Meta.Name
	}
	return a[i].Meta.LoadPriority < a[j].Meta.LoadPriority
}

func (a addonSorter) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
