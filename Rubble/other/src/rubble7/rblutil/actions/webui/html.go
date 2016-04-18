/*
Copyright 2013-2016 by Milo Christiansen

This software is provided 'as-is', without any express or implied warranty. In
no event will the authors be held liable for any damages arising from the use of
this software.

Permission is granted to anyone to use this software for any purpose, including
commercial applications, and to alter it and redistribute it freely, subject to
the following restrictions:

1. The origin of this software must not be misrepresented you must not claim
that you wrote the original software. If you use this software in a product, an
acknowledgment in the product documentation would be appreciated but is not
required.

2. Altered source versions must be plainly marked as such, and must not be
misrepresented as being the original software.

3. This notice may not be removed or altered from any source distribution.
*/

package webui

import "html/template"
import "rubble7/rblutil"
import "dctech/axis"

func LoadHTMLTemplates(fs axis.DataSource) (*template.Template, error) {
	tmpl := template.New("webUI")

	// Main Menu and Common Pages

	_, err := rblutil.LoadOr(fs, tmpl, "menu", ".html", `
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Menu</title>

	<script type="text/javascript">
		window.onload = function() {
			var btn = document.getElementById("Menu");
			btn.addEventListener("click", function(evnt) {
				evnt.stopPropagation()
				evnt.preventDefault()

				var to = {
					"Generate": "/genaddons",
					"Tileset": "/tsetaddons",
					"IApply": "/iaaddons",
					"Addons": "/addons",
					"Docs": "/doclist",
					"Quit": "/kill",
					"Log": "/log",
					"About": "/doc/Rubble License.md",
				}[evnt.target.id]

				if (to === undefined) {
					return
				}
				window.location.replace("http://" + window.location.host + to)
			}, false)
		}
	</script>
</head>
<body>
	<h1>Welcome to the Rubble Web UI!</h1>
	<div id="Menu">
		<p><div id="Generate" class="button">Generate raws</div></p>
		<p><div id="Tileset" class="button">Switch a region's tileset</div></p>
		<p><div id="IApply" class="button">Apply a save-safe addon to a region</div></p>
		<div class="hr"></div>
		<p><div id="Addons" class="button">Addon documentation</div></p>
		<p><div id="Docs" class="button">General documentation</div></p>
		<div class="hr"></div>
		<p><div id="Log" class="button">Current log file</div></p>
		<p><div id="About" class="button">About Rubble</div></p>
		<div class="hr"></div>
		<p><div id="Quit" class="button">Exit Rubble</div></p>
	</div>
</body>
</html>
`)
	if err != nil {
		return nil, err
	}

	_, err = rblutil.LoadOr(fs, tmpl, "pleasewait", ".html", `
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Running Operation</title>

	<script type="text/javascript">
		window.onload = function() {
			window.location.replace("http://" + window.location.host + "{{.}}")
		}
	</script>
</head>
<body>
	<h2>Rubble is working</h2>
	<p>Please wait while the requested operation completes.
	<p>You will automatically be redirected to the log page to view the results of the operation once Rubble is finished.
</body>
</html>
`)
	if err != nil {
		return nil, err
	}

	_, err = rblutil.LoadOr(fs, tmpl, "log", ".html", `
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Log</title>

	<script type="text/javascript">
		window.onload = function() {
			function listener(evnt) {
				evnt.stopPropagation()

				var to = {
					"Back": "/",
					"Quit": "/kill",
				}[evnt.target.id]

				if (to === undefined) {
					return
				}
				window.location.replace("http://" + window.location.host + to)
			}

			var Menu1 = document.getElementById("Menu1")
			Menu1.addEventListener("click", listener, false)

			var Menu2 = document.getElementById("Menu2")
			Menu2.addEventListener("click", listener, false)
		}
	</script>
</head>
<body>
	<div id="Menu1">
		<p><div id="Back" class="button">Back to Menu</div></p>
		<p><div id="Quit" class="button">Exit Rubble</div></p>
	</div>

	{{if .Header}}{{if eq .State "error"}}<p>An error has occurred or Rubble has aborted, scroll to the end of the log for details.</p>{{else if eq .State "warn"}}<p>Action completed with warnings, scroll to the end of the log for details.</p>{{else}}<p>Action completed without errors.</p>{{end}}{{end}}

	<pre>{{.Log}}</pre>

	<div id="Menu2">
		<p><div id="Back" class="button">Back to Menu</div></p>
		<p><div id="Quit" class="button">Exit Rubble</div></p>
	</div>
</body>
</html>
`)
	if err != nil {
		return nil, err
	}

	_, err = rblutil.LoadOr(fs, tmpl, "kill", ".html", `
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Exit</title>
</head>
<body>
	<h2>Thank you for using Rubble Web UI!</h2>
	<p>The server has shutdown.
	<p>Please post any suggestions or bugs to the Rubble thread on the forum!
</body>
</html>
`)
	if err != nil {
		return nil, err
	}

	_, err = rblutil.LoadOr(fs, tmpl, "doclist", ".html", `
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Documentation</title>

	<script type="text/javascript">
		window.onload = function() {
			// Kill the back button as best I can. There may be a better way to do this, but I don't know it.
			function killback(evnt) {
				evnt.preventDefault()
				if (evnt.target.href === undefined && this.href !== undefined) {
					evnt.stopPropagation()
					window.location.replace(this.href)
				} else if (evnt.target.href !== undefined && evnt.target.href !== "" && evnt.target.href !== "#") {
					evnt.stopPropagation()
					window.location.replace(evnt.target.href)
				}
			}
			var a = document.getElementsByTagName("a")
			for (var i = 0; i < a.length; i++) {
				a[i].addEventListener("click", killback, false)
				
				// Fix addon links
				if (a[i].href.indexOf("/addondata?addon=") != -1) {
					a[i].href = a[i].href + "&from=" + escape(window.location.pathname + window.location.search) + "&fromname=Previous Page"
				}
			}

			function listener(evnt) {
				evnt.stopPropagation()

				window.location.replace("http://" + window.location.host + "/")
			}

			document.getElementById("Back").addEventListener("click", listener, false)
		}
	</script>
</head>
<body>
	<p>Stuff everyone should read:</p>
	<li><a href="/doc/Rubble Readme.md">Rubble Readme</a></li>
	<li><a href="/doc/Rubble Changelog.md">Rubble Changelog</a></li>

	<p><div class="hr"></div></p>

	<p>Extras:</p>
	<li><a href="/doc/Rubble History and Philosophy.md">Rubble History and Philosophy</a></li>

	<p><div class="hr"></div></p>

	<p>For modders:
	<li><a href="/doc/Rubble Basics.md">Rubble Basics</a></li>
	<li><a href="/doc/Rubble Base Templates.md">Rubble Base Templates</a></li>
	<li><a href="/doc/Rubble Conventions.md">Rubble Conventions</a></li>

	<p><div class="hr"></div></p>

	<p>Miscellaneous tutorials:
	<li><a href="/doc/Tutorials/Common standard library templates.md">Common standard library templates</a></li>
	<li><a href="/doc/Tutorials/User templates.md">User templates</a></li>
	<li><a href="/doc/Tutorials/Tileset addons.md">Tileset addons</a></li>

	{{if gt (len .) 0}}<p><div class="hr"></div></p><p>External Documentation:</p>
	{{range .}}<li>{{.Header}}</li>
	{{end}}{{end}}

	<p><div id="Back" class="button">Back to Menu</div></p>
</body>
</html>
`)
	if err != nil {
		return nil, err
	}

	_, err = rblutil.LoadOr(fs, tmpl, "docpage", ".html", `
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Documentation</title>

	<script type="text/javascript">
		function Toggle(id) {
			var el = document.getElementById(id)
			if (el.style.display == "none") {
				el.style.display = ""
			} else {
				el.style.display = "none"
			}
		}
		
		window.onload = function() {
			// Kill the back button as best I can. There may be a better way to do this, but I don't know it.
			function killback(evnt) {
				if (evnt.target.href === undefined && this.href !== undefined) {
					evnt.preventDefault()
					evnt.stopPropagation()
					window.location.replace(this.href)
				}else if (evnt.target.href !== undefined && evnt.target.href !== "" && evnt.target.href !== "#") {
					evnt.preventDefault()
					evnt.stopPropagation()
					window.location.replace(evnt.target.href)
				}
			}
			var a = document.getElementsByTagName("a")
			for (var i = 0; i < a.length; i++) {
				a[i].addEventListener("click", killback, false)
				
				// Fix doc links
				if (a[i].href.indexOf("/doc/") != -1) {
					a[i].href = "http://" + a[i].host + a[i].pathname + "?from=" + escape(window.location.pathname + window.location.search) + a[i].hash
				}
				
				// Fix addon links
				if (a[i].href.indexOf("/addondata?addon=") != -1) {
					a[i].href = a[i].href + "&from=" + escape(window.location.pathname + window.location.search) + "&fromname=Previous Page"
				}
			}

			function listener(evnt) {
				evnt.stopPropagation()

				var to = {
					"Back": {{if eq .From ""}}"/doclist"{{else}}{{.From}}{{end}},
					"ToMenu": "/",
				}[evnt.target.id]

				if (to === undefined) {
					return
				}
				window.location.replace("http://" + window.location.host + to)
			}

			var Menu1 = document.getElementById("Menu1")
			Menu1.addEventListener("click", listener, false)

			var Menu2 = document.getElementById("Menu2")
			Menu2.addEventListener("click", listener, false)
		}
	</script>
</head>
<body>
	<div id="Menu1">
		<p><div id="Back" class="button">Back to {{if eq .From ""}}Documentation{{else}}Previous Page{{end}}</div></p>
		<p><div id="ToMenu" class="button">To Main Menu</div></p>
	</div>

	{{.Body}}

	<div id="Menu2">
		<p><div id="Back" class="button">Back to {{if eq .From ""}}Documentation{{else}}Previous Page{{end}}</div></p>
		<p><div id="ToMenu" class="button">To Main Menu</div></p>
	</div>
</body>
</html>
`)
	if err != nil {
		return nil, err
	}

	_, err = rblutil.LoadOr(fs, tmpl, "addondata", ".html", `{{with .Addon}}
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Addon Information for "{{.Meta.Name}}"{{if ne .Meta.Version ""}} ({{.Meta.Version}}){{end}}</title>

	<script type="text/javascript">
		function Toggle(id) {
			var el = document.getElementById(id)
			if (el.style.display == "none") {
				el.style.display = ""
			} else {
				el.style.display = "none"
			}
		}

		window.onload = function() {
			// Kill the back button as best I can. There may be a better way to do this, but I don't know it.
			function killback(evnt) {
				if (evnt.target.href === undefined && this.href !== undefined) {
					evnt.preventDefault()
					evnt.stopPropagation()
					window.location.replace(this.href)
				}else if (evnt.target.href !== undefined && evnt.target.href !== "" && evnt.target.href !== "#") {
					evnt.preventDefault()
					evnt.stopPropagation()
					window.location.replace(evnt.target.href)
				}
			}
			var a = document.getElementsByTagName("a")
			for (var i = 0; i < a.length; i++) {
				a[i].addEventListener("click", killback, false)

				// Fix addon links
				if (a[i].href.indexOf("/addondata?addon=") != -1) {
					a[i].href = a[i].href + "&from={{urlquery $.FromUrl}}&fromname={{urlquery $.FromName}}"
				}

				// Fix doc links
				if (a[i].href.indexOf("/doc/") != -1) {
					a[i].href = "http://" + a[i].host + a[i].pathname + "?from=" + escape(window.location.pathname + window.location.search) + a[i].hash
				}
			}

			document.getElementById("More").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				data = document.getElementById("MoreData")
				if (data.style.display == "none") {
					data.style.display = "";
					evnt.target.innerHTML = "Hide Details";
				} else {
					data.style.display = "none";
					evnt.target.innerHTML = "Show Details";
				}
			}, false)

			document.getElementById("Back").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				window.location.replace("http://" + window.location.host + {{$.FromUrl}})
			}, false)
		}
	</script>
</head>
<body>
	<h2>{{.Meta.Name}}{{if ne .Meta.Version ""}} ({{.Meta.Version}}){{end}}</h2>

	{{if ne .Meta.Header ""}}<p>{{.Meta.Header}}</p>{{end}}

	{{if ne .Meta.Description ""}}{{.Meta.Description}}{{end}}

	<p><div id="More" class="button">Show Details</div></p>
	<div id="MoreData" style="display:none;">
		<h3>Extra Data</h3>

		{{if ne .Meta.Author ""}}<p>Addon written by: {{.Meta.Author}}</p>{{end}}
		{{if ne .Meta.Version ""}}<p>Version: {{.Meta.Version}}</p>{{end}}

		<p>Addon Path: <code>"{{.Source}}"</code></p>

		<p>Load Priority: <code>"{{.Meta.LoadPriority}}"</code></p>

		<h4>Tags:</h4>{{range $n, $v := .Meta.Tags}}
		<li><code>"{{$n}}" = {{$v}}</code>{{if eq $n "Library"}} (Is this addon is an automatically managed library?){{else if eq $n "DocPack"}} (You should never see this, addons with this tag are not really addons.){{else if eq $n "TileSet"}} (Does this addon contains tileset information?){{else if eq $n "SaveSafe"}} (Can this addon be applied to worlds in progress?){{else if eq $n "DFHack"}} (Does this addon requires DFHack to operate?){{else if eq $n "HasTests"}} (Does this addon have template test files?){{else if eq $n "NotNormal"}} (Is this addon not a normal addon, eg should it be hidden during normal generation cycles?){{else if eq $n "Dev"}} (Is this addon something that will only be interesting to addon developers?){{end}}</li>{{else}}
		<li>This addon has no tags.</li>{{end}}

		<h4>Dependencies (automatically activated):</h4>{{range .Meta.Activates}}
		<li><a href="/addondata?addon={{.}}">{{.}}</a></li>{{else}}
		<li>This addon has no dependencies.</li>{{end}}

		<h4>Incompatibilities:</h4>{{range .Meta.Incompatible}}
		<li><a href="/addondata?addon={{.}}">{{.}}</a></li>{{else}}
		<li>This addon has no incompatibilities.</li>{{end}}

		<h4>Variables (and their defaults):</h4>
		<script language="JavaScript">
			// This is me being lazy, I really should do this with templates like everything else...
			var vars = {{.Meta.Vars}}

			var ok = false
			for (var i in vars) {
				if (vars[i].Name == "-") {
					continue
				}

				ok = true
				if (vars[i].Values.length > 0) {
					document.write("<li>" + vars[i].Name + " (<code>" + i + "</code>): <code>\"" + vars[i].Values[0] + "\"</code></a>")
				} else {
					document.write("<li>" + vars[i].Name + " (<code>" + i + "</code>): No Default Value Specified</a>")
				}
			}
			if (!ok) {
				document.write("<li>This addon has no variables.</li>")
			}
		</script>

		<h4>Addon Files:</h4>
		{{range .Files}}<li><a href="javascript:Toggle('FILE:{{.Name}}');">{{.Name}}</a><pre id="FILE:{{.Name}}" style="display:none;" class="border">{{printf "%s" .Content}}</pre></li>
		{{else}}<li>This addon has no files (how is this even possible?).</li>{{end}}
	</div>

	<p><div id="Back" class="button">Back to {{$.FromName}}</div></p>
</body>
</html>
{{end}}`)
	if err != nil {
		return nil, err
	}

	// Master Addon List

	_, err = rblutil.LoadOr(fs, tmpl, "addons", ".html", `
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Addon List</title>

	<script type="text/javascript">
		window.onload = function() {
			// Kill the back button as best I can. There may be a better way to do this, but I don't know it.
			function killback(evnt) {
				if (evnt.target.href === undefined && this.href !== undefined) {
					evnt.preventDefault()
					evnt.stopPropagation()
					window.location.replace(this.href)
				}else if (evnt.target.href !== undefined && evnt.target.href !== "" && evnt.target.href !== "#") {
					evnt.preventDefault()
					evnt.stopPropagation()
					window.location.replace(evnt.target.href)
				}
			}
			var a = document.getElementsByTagName("a")
			for (var i = 0; i < a.length; i++) {
				a[i].addEventListener("click", killback, false)
				
				// Fix doc links
				if (a[i].href.indexOf("/doc/") != -1) {
					a[i].href = "http://" + a[i].host + a[i].pathname + "?from=" + escape(window.location.pathname + window.location.search) + a[i].hash
				}
				
				// Fix addon links
				if (a[i].href.indexOf("/addondata?addon=") != -1) {
					a[i].href = a[i].href + "&from=" + escape(window.location.pathname + window.location.search) + "&fromname=Previous Page"
				}
			}

			function listener(evnt) {
				evnt.stopPropagation()

				window.location.replace("http://" + window.location.host + "/")
			}

			document.getElementById("Back1").addEventListener("click", listener, false)
			document.getElementById("Back2").addEventListener("click", listener, false)
		}
	</script>
</head>
<body>
	<p><div id="Back1" class="button">Back to Menu</div></p>
	<table id="Addons">
		{{range .Addons}}{{if not .Meta.Tags.DocPack}}<tr><td><a href="/addondata?addon={{.Meta.Name}}">{{.Meta.Name}}</a></td><td>{{.Meta.Header}}</td></tr>
		{{end}}{{end}}
	</table>
	<p><div id="Back2" class="button">Back to Menu</div></p>
</body>
</html>
`)
	if err != nil {
		return nil, err
	}

	// Normal Generation

	_, err = rblutil.LoadOr(fs, tmpl, "genaddons", ".html", `{{with .State}}
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Select Addons</title>
	<script type="text/javascript">
		function Toggle(addon, state) {
			var req = new XMLHttpRequest()
			req.open("POST", "/toggle", false)
			req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
			req.send("addon=" + encodeURIComponent(addon) + "&state=" + encodeURIComponent(state))
		}

		var addontags = { {{range .Addons}}
			"{{.Meta.Name}}": {{.Meta.Tags}},{{end}}
		}

		window.onload = function() {
			// Kill the back button as best I can. There may be a better way to do this, but I don't know it.
			function killback(evnt) {
				if (evnt.target.href === undefined && this.href !== undefined) {
					evnt.preventDefault()
					evnt.stopPropagation()
					window.location.replace(this.href)
				}else if (evnt.target.href !== undefined && evnt.target.href !== "" && evnt.target.href !== "#") {
					evnt.preventDefault()
					evnt.stopPropagation()
					window.location.replace(evnt.target.href)
				}
			}
			var a = document.getElementsByTagName("a")
			for (var i = 0; i < a.length; i++) {
				a[i].addEventListener("click", killback, false)

				// Fix addon links
				if (a[i].href.indexOf("/addondata?addon=") != -1) {
					a[i].href = a[i].href + "&from=%2fgenaddons&fromname=Generation"
				}
				
				// Fix doc links
				if (a[i].href.indexOf("/doc/") != -1) {
					a[i].href = "http://" + a[i].host + a[i].pathname + "?from=" + escape(window.location.pathname + window.location.search) + a[i].hash
				}
			}

			document.getElementById("AddonList").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				if (evnt.target.name != "") {
					Toggle(evnt.target.name, evnt.target.checked == true)
				}
			}, false)

			document.getElementById("Select").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				var state
				if (evnt.target.id == "All") {
					state = true
				} else if (evnt.target.id == "None") {
					state = false
				}
				if (state === undefined) {
					return
				}

				var addons = document.getElementById("AddonList").getElementsByTagName("tr")
				for (var i = 0; i < addons.length; i++) {
					var item = addons[i].getElementsByTagName("input")[0]
					
					if (addons[i].style.display != "none") {
						Toggle(item.name, state)
						item.checked = state
					}
				}
			}, false)

			function ToggleTag(tag) {
				var addons = document.getElementById("AddonList").getElementsByTagName("tr")
				for (var i = 0; i < addons.length; i++) {
					if (addontags[addons[i].id][tag] == true) {

						if (addons[i].style.display == "none") {
							addons[i].style.display = ""
						} else {
							addons[i].style.display = "none"
						}

						//if (addons[i].checked) {
						//	Toggle(addons[i].name, false)
						//	addons[i].checked = false
						//}
					}
				}
			}
			ToggleTag("Dev")

			document.getElementById("Hide/Show").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				if (evnt.target.innerHTML.indexOf("+") != -1) {
					evnt.target.innerHTML = evnt.target.id + " -"
				} else {
					evnt.target.innerHTML = evnt.target.id + " +"
				}

				ToggleTag(evnt.target.id)
			}, false)

			document.getElementById("EditConfig").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				var to = {
					"EditConfig": "/genvars",
					"Generate": "/genvars?hidden=true",
				}[evnt.target.id]

				if (to === undefined) {
					return
				}
				window.location.replace("http://" + window.location.host + to)
			}, false)

			{{if gt (len $.Presets) 1}}
			var presets = { {{range $.Presets}}
				"{{.Name}}": { {{range .Addons}}
					"{{.}}": true,{{end}}
				},{{end}}
			}
			
			document.getElementById("Preset").addEventListener("change", function(evnt) {
				evnt.stopPropagation()
				
				var addons = document.getElementById("AddonList").getElementsByTagName("tr")
				for (var i = 0; i < addons.length; i++) {
					var item = addons[i].getElementsByTagName("input")[0]
					
					if (presets[evnt.target.value][item.name] == true) {
						if (!item.checked) {
							Toggle(item.name, true)
							item.checked = true
						}
					} else {
						if (item.checked) {
							Toggle(item.name, false)
							item.checked = false
						}
					}
				}
			}, false)
			{{end}}
			
			document.getElementById("Back").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				window.location.replace("http://" + window.location.host + "/")
			}, false)
		}
	</script>
</head>
<body>
	<table id="AddonList">{{range .Addons}}{{if and (not .Meta.Tags.Library) (not .Meta.Tags.DocPack) (not .Meta.Tags.NotNormal)}}
		<tr id="{{.Meta.Name}}"><td nowrap>{{if .Active}}<input type="checkbox" checked name="{{.Meta.Name}}">{{else}}<input type="checkbox" name="{{.Meta.Name}}">{{end}}<a href="/addondata?addon={{.Meta.Name}}">{{.Meta.Name}}</a></td><td>{{.Meta.Header}}</td></tr>{{end}}{{end}}
	</table>
	{{if gt (len $.Presets) 1}}
	<p>Preset: <select id="Preset">{{range $.Presets}}
		<option{{if eq .Name "Default"}} selected{{end}}>{{.Name}}</option>{{end}}
	</select></p>{{end}}
	<p><div id="Select" class="button">
		Select
		<div id="All" class="nose">All</div>
		<div id="None" class="nose">None</div>
	</div></p>
	<p><div id="Hide/Show" class="button">
		Hide/Show
		<div id="DFHack" class="nose">DFHack -</div>
		<div id="Dev" class="nose">Dev +</div>
	</div></p>
	<p><div id="EditConfig" class="button">
		Edit configuration variables
		<div id="Generate" class="nose">Skip</div>
	</div></p>
	<p><div id="Back" class="button">Back to Menu</div></p>
</body>
</html>
{{end}}`)
	if err != nil {
		return nil, err
	}

	_, err = rblutil.LoadOr(fs, tmpl, "genvars", ".html", `
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Edit Configuration Variables</title>

	<script type="text/javascript">
		function Set(key, val) {
			var req = new XMLHttpRequest()
			req.open("POST", "/setvar", false)
			req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
			req.send("key=" + encodeURIComponent(key) + "&val=" + encodeURIComponent(val))
		}

		window.onload = function() {
			document.getElementById("VarList").addEventListener("change", function(evnt) {
				evnt.stopPropagation()

				if (evnt.target.name != "") {
					Set(evnt.target.name, evnt.target.value)
				}
			}, false)

			document.getElementById("Menu").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				var to = {
					"Generate": "/pleasewait?to=/genrun",
					"Back": "/",
				}[evnt.target.id]

				if (to === undefined) {
					return
				}
				window.location.replace("http://" + window.location.host + to)
			}, false)
		}
	</script>
</head>
<body>
	<table id="VarList">
		<script language="JavaScript">
			var items = {{.Vars}}
			var values = {{.Vals}}
			var noitems = true

			for (var i in items) {
				if (items[i].Name == "-") {
					continue
				}

				noitems = false

				var name = i
				if (items[i].Name != "") {
					name = items[i].Name
				}

				document.write("<tr><td>" + name + "</td>")
				document.write("<td>")
				if (items[i].Values.length > 1) {
					document.write("<select name=\"" + i + "\">")
					for (var j in items[i].Values) {
						if (items[i].Values[j] == values[i]) {
							document.write("<option selected>" + items[i].Values[j])
						} else {
							document.write("<option>" + items[i].Values[j])
						}
					}
					document.write("</select>")
				} else {
					document.write("<input type=\"text\" value=\"" + values[i] + "\" name=\"" + i + "\">")
				}
				document.write("</td></tr>")
			}
			if (noitems) {
				document.write("<tr><td>No Variables listed in Meta Data.</td></tr>")
			}
		</script>
	</table>

	<div id="Menu">
		<p><div id="Generate" class="button">Generate!</div></p>
		<p><div id="Back" class="button">Back to Menu</div></p>
	</div>
</body>
</html>
`)
	if err != nil {
		return nil, err
	}

	// Tilesets and Independent Apply

	_, err = rblutil.LoadOr(fs, tmpl, "iaaddons", ".html", `{{with .State}}
<!DOCTYPE html>
<html><!-- This is also used for applying tilesets -->
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Web UI: Select Addons</title>
	<script type="text/javascript">
		function Toggle(addon, state) {
			var req = new XMLHttpRequest()
			req.open("POST", "/toggle", false)
			req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
			req.send("addon=" + encodeURIComponent(addon) + "&state=" + encodeURIComponent(state))
		}

		function Set(key, val) {
			var req = new XMLHttpRequest()
			req.open("POST", "/setvar", false)
			req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
			req.send("key=" + encodeURIComponent(key) + "&val=" + encodeURIComponent(val))
		}

		window.onload = function() {
			// Kill the back button as best I can. There may be a better way to do this, but I don't know it.
			function killback(evnt) {
				if (evnt.target.href === undefined && this.href !== undefined) {
					evnt.preventDefault()
					evnt.stopPropagation()
					window.location.replace(this.href)
				}else if (evnt.target.href !== undefined && evnt.target.href !== "" && evnt.target.href !== "#") {
					evnt.preventDefault()
					evnt.stopPropagation()
					window.location.replace(evnt.target.href)
				}
			}
			var a = document.getElementsByTagName("a")
			for (var i = 0; i < a.length; i++) {
				a[i].addEventListener("click", killback, false)

				// Fix addon links
				if (a[i].href.indexOf("/addondata?addon=") != -1) {
					a[i].href = a[i].href + "&from={{urlquery $.Back_URL}}&fromname={{urlquery $.Name}}"
				}
				
				// Fix doc links
				if (a[i].href.indexOf("/doc/") != -1) {
					a[i].href = "http://" + a[i].host + a[i].pathname + "?from=" + escape(window.location.pathname + window.location.search) + a[i].hash
				}
			}

			document.getElementById("Region").addEventListener("change", function(evnt) {
				evnt.stopPropagation()

				Set("_RUBBLE_IA_REGION_", evnt.target.value)
			}, false)

			document.getElementById("AddonList").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				if (evnt.target.name != "") {
					Toggle(evnt.target.name, evnt.target.checked == true)
				}
			}, false)

			document.getElementById("Select").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				var state
				if (evnt.target.id == "All") {
					state = "true"
				} else if (evnt.target.id == "None") {
					state = "false"
				}
				if (state === undefined) {
					return
				}

				var addons = document.getElementById("AddonList").getElementsByTagName("input")
				for (var i = 0; i < addons.length; i++) {
					Toggle(addons[i].name, state)
					addons[i].value = state
					addons[i].checked = state === "true"
				}
			}, false)

			document.getElementById("Menu").addEventListener("click", function(evnt) {
				evnt.stopPropagation()

				var to = {
					"Apply": "/pleasewait?to={{urlquery $.URL}}",
					"Back": "/",
				}[evnt.target.id]

				if (to === undefined) {
					return
				}
				window.location.replace("http://" + window.location.host + to)
			}, false)
		}
	</script>
</head>
<body>
	<p>Region: <select id="Region">
		<option selected>raw</option>{{range $.Regions}}
		<option>{{.}}</option>{{end}}
	</select></p>

	<table id="AddonList">{{range .Addons}}{{if and (and (not .Meta.Tags.Library) (not .Meta.Tags.DocPack)) (index .Meta.Tags $.Tag)}}
		<tr><td nowrap>{{if .Active}}<input type="checkbox" value="true" checked name="{{.Meta.Name}}">{{else}}<input type="checkbox" value="false" name="{{.Meta.Name}}">{{end}}<a href="/addondata?addon={{.Meta.Name}}">{{.Meta.Name}}</a></td><td>{{.Meta.Header}}</td></tr>{{end}}{{end}}
	</table>

	<p><div id="Select" class="button">
		Select
		<div id="All" class="nose">All</div>
		<div id="None" class="nose">None</div>
	</div></p>
	<div id="Menu">
		<p><div id="Apply" class="button">Apply Selected</div></p>
		<p><div id="Back" class="button">Back to Menu</div></p>
	</div>
</body>
</html>
{{end}}`)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
