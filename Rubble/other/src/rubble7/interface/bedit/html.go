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

package main

import "html/template"
import "rubble7/rblutil"
import "dctech/axis"

func LoadHTMLTemplates(fs axis.DataSource) (*template.Template, error) {
	tmpl := template.New("beditUI")

	// Main Menu and Common Pages

	_, err := rblutil.LoadOr(fs, tmpl, "menu", ".html", `
<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/css"/>
	<title>Rubble Building Editor</title>

	<script type="text/javascript">
		window.onload = function() {
			
		}
	</script>
</head>
<body>
	
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
	<title>Rubble Building Editor: Log</title>

	<script type="text/javascript">
		window.onload = function() {
			function listener(evnt) {
				evnt.stopPropagation()

				window.location.replace("http://" + window.location.host + "/kill")
			}

			document.getElementById("Quit1").addEventListener("click", listener, false)

			document.getElementById("Quit2").addEventListener("click", listener, false)
		}
	</script>
</head>
<body>
	<p><div id="Quit1" class="button">Exit Rubble</div></p>

	<p>An error has occurred, scroll to the end of the log for details.</p>

	<pre>{{.Log}}</pre>

	<p><div id="Quit1" class="button">Exit Rubble</div></p>
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
	<title>Rubble Building Editor: Exit</title>
</head>
<body>
	<h2>Thank you for using the Rubble Building Editor!</h2>
	<p>The server has shutdown.
	<p>Please post any suggestions or bugs to the Rubble thread on the forum!
</body>
</html>
`)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
