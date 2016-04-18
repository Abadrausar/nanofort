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

package main

import "rubble7/rblutil"
import "dctech/axis"

func getCSS(fs axis.DataSource) string {
	// Shades of brown and orange. Surprisingly nice looking, and easy on the eyes.
	return rblutil.LoadOrString(fs, "rubble:other/webUI/rbl.css", `
body {
	color: #BDAE9D;
	background-color: #2A211C;
	
	font-family: sans-serif;
	font-size: 12pt;
}

table {
	color: #BDAE9D;
	background-color: #2A211C;
	
	border-radius: 3px;
	border-color: #DE8D47;
}

a {
	color: #BDAE9D;
	text-decoration: underline;
}

a:hover {
	color: #2A211C;
	background: #BDAE9D;
	text-decoration: none;
}

span.mono {
	font-family: monospace;
}

code {
	font-family: monospace;
}

pre {
	border-style: none;
	padding: 5px 5px 5px 5px;
	margin: 5px 5px 5px 5px;
	
	font-family: monospace;
}

pre.border {
	border-radius: 3px;
	border: 2px dashed #DE8D47;
	padding: 5px 5px 5px 5px;
	margin: 5px 5px 5px 5px;
	
	font-family: monospace;
}

h1 {
	color: #2A211C;
	background: gray;
	
	border-top: 4px solid #DE8D47;
	border-bottom: 4px solid #DE8D47;
	text-indent: 10px;
	
	font-size: 20pt;
}

h2 {
	color: #2A211C;
	background: gray;
	
	border-top: 3px solid #DE8D47;
	border-bottom: 3px solid #DE8D47;
	text-indent: 10px;
	
	font-size: 16pt;
}

h3 {
	color: #2A211C;
	background: gray;
	
	border-top: 2px solid #DE8D47;
	border-bottom: 2px solid #DE8D47;
	text-indent: 10px;
	
	font-size: 14pt;
}

h4 {
	color: #2A211C;
	background: gray;
	
	border-top: 1px solid #DE8D47;
	border-bottom: 1px solid #DE8D47;
	text-indent: 10px;
	
	font-size: 12pt;
}

input {
	background-color: #2A211C;
	color: #BDAE9D;
	
	border-radius: 3px;
	border: 1px solid #DE8D47;
	padding: 5px 5px 5px 5px;
}

select {
	background-color: #2A211C;
	color: #BDAE9D;
	
	min-width: 100px;
	
	border-radius: 3px;
	border: 1px solid #DE8D47;
	padding: 5px 5px 5px 5px;
}

select.small {
	min-width: 0px;
}

div.button {
	color: gray;
	background-color: #212121;
	
	cursor: pointer;
	text-align: center;
	margin-bottom:10px;
	
	border-radius: 3px;
	border: 1px solid #DE8D47;
	width: 300px;
	
	display: block;
	position: relative;
	
	padding: 5px 5px 5px 5px;
}

div.small {
	width: auto;
	min-width: 20px;
	
	padding: 2px 2px 2px 2px;
}

div.button .nose {
	border-left: 1px solid #DE8D47;
	padding: 5px !important;
	margin: -5px 0;
	min-width: 20px;
	text-decoration: none;
	
	padding-left: 2px;
	float: right;
}

div.hr {
	display: block;
	
	width: 300px;
	height: 1px;
	
	background-color: gray;
	border: 1px solid #DE8D47;
	
	padding-left: 5px;
}

hr {
	display: block;
	
	height: 1px;
	
	background-color: gray;
	border: 1px solid #DE8D47;
}

`)
}
