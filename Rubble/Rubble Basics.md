
Rubble Basics
==============================================

Despite the name this file deals with some very advanced topics! Basically
anything that needed to be documented but didn't deserve it's own file got
a section here. If you do not understand something just skip it, chances are
it is written for advanced users about a feature added to solve a rare problem
you have no need to solve :P

That said you should at least skim over everything so you have some idea
what is possible.

Note that some of the things described here are added by the `addon:Libs/Base` addon
and are not hard-coded parts of Rubble. `addon:Libs/Base` is commonly assumed to be
an integral part of Rubble, it is safe to assume that it's features are always
available.

<div id="OSSpecific"></div>
Rubble OS Specific Information:
----------------------------------------------

On Windows:

On Windows Rubble should "just work", everything was originally written and
tested on a 64 bit Windows 7 system, so if Rubble works anywhere it should
work there :) (Nowadays I run 64 bit Windows 10, and Rubble still works flawlessly)

No specific issues are known.

* * *

On Linux:

To get anything working at all you need to:

* The binary need to be marked "executable" (chmod +x) and moved to the
  same directory as the Windows binary.
* The browser startup script needs to be marked "executable" (chmod +x)

After that things *should* work fine.

If the web UI does not automatically start your browser, you will have to modify
the browser startup script (comments in the default script list some options).

If anyone who knows Linux shell script can write a more universal startup script
let me know!

Keep in mind that I have never actually played DF on Linux, and Rubble testing on
that platform was limited to some simple generation tests a long time ago. I would
love to hear from someone who uses Rubble on Linux.

* * *

On OSX:

It is impossible for me to do any OSX testing (and I have never used OSX for
anything but the most trivial tasks on a borrowed computer). If you want to know
how Rubble works on this platform either try it yourself or donate some hardware :)

At a minimum you will have to move the binary to the same directory as the
Windows binary (do you have to mark it executable like on Linux? OSX and
Linux are both UNIX clones).

There is no default browser startup script for the web UI on OSX. It will try to
use the Linux script, but this is unlikely to work.

If anyone *has* run Rubble on OSX please contact me and tell me how it went!

<div id="AXIS"></div>
AXIS VFS:
----------------------------------------------

For a variety of reasons Rubble uses AXIS VFS for (almost) all file access.

Most places where a path is given the AXIS syntax is used, so you will need to
know at least a little so you can read these paths.

AXIS uses special (OS-independent) multi-part paths. The first part is one or more
colon separated location IDs. You can think of location IDs as drive letters, they
each specify a different location on the file system. The second part is the
traditional slash separated path.

Examples:

	df:data/init/d_init.txt
	out:objects/entity_default.txt

In some cases it may be possible to specify multiple location IDs on a path, in
that case just put them one right after another.

Example:

	addons:dir:addonlist.ini

One thing to keep in mind: Relative paths (paths containing "." or ".." elements)
are utterly illegal in AXIS, so don't do that.

The only paths that are not handled by AXIS are:

* rbldir command-line/config option
* dfdir command-line/config option
* outputdir command-line/config option
* addonsdir command-line/config option

And even these paths may use the AXIS syntax (with a reduced set of location IDs):

* dfdir may use the `rubble:` location
* outputdir may use the `df:` and `rubble:` locations
* addonsdir may use the `out:`, `df:`, and `rubble:` locations

You may still use OS paths (which may be relative) for these three settings, but
you do not have to if using the AXIS syntax would be easier (you can even use AXIS
paths with relative elements if you wish).

Location IDs:

	rubble:        The rubble directory, most of the time this is where the Rubble binary is, defaults to `.`
	df:            The dwarf fortress directory, defaults to `..`
	out:           The output directory, by default `df:raw`
	addons:dir:    The addons directory, by default `rubble:addons`, there may be more than one addons directory
	addons:zip:    A composite, read-only, directory made up of the contents of all the zip files in the addons directories

<div id="FileTags"></div>
File Extensions and Tags:
----------------------------------------------

Rubble keeps track of file type and other attributes via "file tags", simple boolean flags
that can be set at will.

File tags default to false and can be changed at any time, any string can be used as a tag ID.

Rubble generally uses file extensions to determine file tags, these associations
are hardcoded (but may be overridden by clever use of file tags in loader scripts).

To keep complexity under control Rubble uses "two part" extensions for most files. These extensions take the form
`<file name>.<first>.<last>`. The first and last parts are each assigned tags separately, and the two groups of tags
are added together to get the file's tags. If a file's extension only has one part it will only get tags from the "last"
group.

File extensions are case-sensitive!

Filenames should also be assumed to be case-sensitive. On some platforms (Windows) this property may not hold, but if
you try to use an addon that ignores filename case it will fail on platforms where filenames are case sensitive (Linux).

First part extensions:

	Extension   Tags                     Use
	.pre        PreScript                A pre script, runs just after init scripts.
	.post       PostScript               A post script, runs just after generation.
	.mod        ModuleScript             A script module, semantics differ with scripting language.
	.load       LoadScript, GlobalFile   A loader script, run just after addons are activated, but before incompatibilities are checked.
	.init       InitScript, GlobalFile   An init script, run just before generation (even if their addon is not active).
	
	.tset       TileSet                  A tileset related file, various uses depending on the second part tag.
	.text       AUXTextFile              An auxiliary text file (to be installed into `out:objects/text`).
	.graphics   CreatureGraphics         A file related to creature graphics. Generally used with the "Image" or "RawFile" tag.
	
	.dfcom      DFHack, CommandScript    A DFHack command script.
	.dfmod      DFHack, ModuleScript     A DFHack script module.
	.dfload     DFHack, LoadScript       A DFHack initialization script.

Second part extensions:

	Extension   Tags                     Use
	.lua        LangLua                  A Lua script. Use with any of the first part extensions except `.text` and `.graphics`.
	.luab       LangLuaBin               Effectively identical to LangLua, except for precompiled Lua binaries.
	.rb         LangRuby                 A Ruby script. Should work with `.dfcom`.
	
	.test       TemplateTest             A template unit test file. Use by itself.
	.txt        RawFile                  A DF raw file. Use with `.tset`, `.text`, `.graphics`, or by itself.
	.rbl        RawFile, NoWrite         A DF raw file. Use with `.tset`, `.text`, `.graphics`, or by itself.
	.rules      MergeRules               A set of raw merging/matching rules. Use with `.tset` or by itself.
	.twbt       TWBTOverride             A TWBT override file. Use with `.tset`.
	
	.png        ImagePNG                 A PNG image file. Use with `.graphics` or `.tset`.
	.bmp        ImageBMP                 A BMP image file. Use with `.graphics` or `.tset`.

In addition to the tagged file types there are some extensions that have hardcoded handling:

	Extension     Use
	.webload      A webload addon descriptor. May only appear in the root of an addons directory.
	.zip          A compressed addon pack. May only appear in the root of an addons directory.

There are some files that have hardcoded names. You may only have one of each of these for each addon (obviously).

	Name          Use
	addon.meta    Addon meta-data file, contains JSON encoded information about the addon.

Some special tags have hard coded functionality but are not automatically applied:
	
	Tag           Description
	Skip          This file should be skipped in any subsequent actions that would otherwise involve it.

Some extra special tags have hard coded functionality, are strictly internal, and should *never* be used for anything:
	
	Tag           Description
	LangTemplate  Used to mark user template bodies, these templates need special fields set in the template body to function.

The following file tags are applied to raw files based on the beginning of the file name:

	File Name Prefix     Tag
	b_detail_plan_       BodyDetailRaws
	body_                BodyRaws
	building_            BuildingRaws
	c_variation_         CreatureVarRaws
	creature_            CreatureRaws
	descriptor_          DescriptorRaws
	entity_              EntityRaws
	inorganic_           InorganicRaws
	interaction_         InteractionRaws
	item_                ItemRaws
	language_            LanguageRaws
	material_template_   MatTemplateRaws
	plant_               PlantRaws
	reaction_            ReactionRaws
	tissue_template_     TissueRaws

These tags are not perfect, make sure you use `Skip` and `NoWrite` as negatives and `RawFile` as
a positive in your filters to prevent false positives.

Certain file tag combinations cause Rubble to write a file to a certain location after generation.
These combinations are as follows:

	Write To            If Tags                          But Not Tags
	out:objects         RawFile                          TileSet, AUXTextFile, CreatureGraphics, Skip, NoWrite
	out:objects/text    RawFile, AUXTextFile             Skip, NoWrite
	out:graphics        RawFile, CreatureGraphics        Skip, NoWrite
	out:graphics        RawFile, ImagePNG                Skip, NoWrite
	out:graphics        RawFile, ImageBMP                Skip, NoWrite
	df:data/art         TileSet, ImagePNG                Skip, NoWrite
	df:data/art         TileSet, ImageBMP                Skip, NoWrite
	out:modules         DFHack, ModuleScript, LangLua    Skip, NoWrite
	out:init.d          DFHack, LoadScript, LangLua      Skip, NoWrite
	out:scripts         DFHack, CommandScript, LangLua   Skip, NoWrite
	out:scripts         DFHack, CommandScript, LangRuby  Skip, NoWrite

Most of the time when writing files based on tags Rubble will look for the standard two part extension that
would result in those tags and replaces it with a basic one part extension. For example `test.tset.png` would
be written as `test.png`. What follows is a list of all effected extensions:

	Two Part        One Part
	.text.txt       .txt
	.graphics.txt   .txt
	.graphics.png   .png
	.graphics.bmp   .bmp
	.tset.png       .bmp
	.test.bmp       .bmp
	.dfmod.lua      .lua
	.dfload.lua     .lua
	.dfcom.lua      .lua
	.dfcom.rb       .rb

If a file with tags that would normally trigger file extension compaction does not have the expected extension
then it is not modified. Extension compaction on write is possible because Rubble uses extensions to convey data
that DF and DFHack get from file location.

<div id="Prefixes"></div>
Template Prefixes Explained:
----------------------------------------------

When a template is run is determined by a one character prefix (or lack of a prefix).
There are three parse stages and four prefixes.

	Prefix  Meaning
	!       Preparse, a template with this prefix will run in the preparse stage.
	none    Parse, a template with no prefix will run in the parse stage.
	#       Postparse, a template with this prefix will run in the postparse stage.
	@       Any, a template with this prefix will run in the earliest possible parse stage.

@ templates are usually run in preparse but could be run in any stage depending on when it
is parsed (parsing can be delayed by nesting in a later parse stage template for example).
This prefix is used for templates that may be run in any stage, if you need to control when
these templates are parsed use !ECHO, ECHO, or #ECHO.

The template prefix is part of the template's name, you cannot change an existing template's parse
stage by calling it with a different prefix.

<div id="ParseSeq"></div>
User Template Parse Sequence:
----------------------------------------------

The following details how templates created with the !TEMPLATE template (or similar) are parsed.

1. The template's body is loaded and these transforms are run on it:
	1. Text of the forms `%<NAME>` or `%{<NAME>}` is replaced with the value of the param named
	   `<NAME>` or the default value specified for that param.
	2. Text of the forms `$<NAME>` or `${<NAME>}` *that is not in a child template* is replaced
	   with the value of the variable `<NAME>`.
2. The template body is passed to the stage parser and the result is returned

As you can see user templates are very simple in implementation.

See the section on [variable expansion pitfalls](#ExpandIssues) for ways that variable and parameter
expansion can go wrong (and how to fix it if it does).

<div id="ConfigFile"></div>
The Configuration File:
----------------------------------------------

Rubble allows you to change its settings via command line options or a config file.
To see these options and their defaults run `rubble -h` and `rubble_web -h` (each
one has a different list).

*All* command line options may also be specified in the config file.

Rubble tries to read the file "./rubble.ini" (*not* an AXIS path!), if this does not fail Rubble
will load settings from here *before* processing command line options (command line options will
override the config settings).

Example "rubble.ini" (using some of the defaults):

	[rubble]
	rbldir = .
	dfdir = ..
	outputdir = df:raw
	addonsdir = rubble:addons

Duplicate keys are just fine, they will act pretty much exactly like duplicate options
on the command line.

Most users will never need the configuration file, but when you do need it it is invaluable.

<div id="CFGVars"></div>
Configuration Variables:
----------------------------------------------

Rubble has a system to store simple string key/value pairs for extra configuration information.

Internally Rubbles uses these "configuration variables" (or "config vars" for short) to store user
settings that effect operations outside of the Rubble core and for temporary storage of internal
state information in interfaces.

By convention variables are named using all uppercase with no punctuation aside from underscores,
but there is no reason why you cannot use any characters you want (including high Unicode and
unprintable characters).

Variables that start and end with underscores are reserved by Rubble, do not use such names.

The following hardcoded variables are defined:

* `_RUBBLE_NO_CLEAR_`: Stops the clear outputs script from running when set to "true". Setting this effectively
  sets `_RUBBLE_NO_CLEAR_GRAPHICS_` as well.
* `_RUBBLE_NO_CLEAR_GRAPHICS_`: Stops the creature graphics directory from being cleared when set to "true".
* `_RUBBLE_IA_REGION_`: Used internally by the web UI, don't touch.

In general you should not set any of the internal variables unless you fully understand what they do!

* * *

There are several ways to set config vars, which you should use mostly depends on how the variable will be used.

For variables that will be used to set options related to an addon the variables should be listed in the addon's
addon.meta file, as this will automatically set the variable to whatever default value you specify and will allow
easy editing of the variable's value from the web UI. See the [section on addon.meta](#AddonMeta) for details.

For variables that are used to track internal state you should **not** list them in the addon.meta file, and
instead just set them directly using the `@SET` template or in a script with the APIs provided (see the API
documentation for the language you are using for details).

Sometimes an addon's only reason for being is to set a config var a certain way. This is rare (as it is generally
better to use the config var editor). In this case it may be best to list the variable in the addon.meta file
with the display name set to "-". This will suppress the normal entry in the config var editor, but still sets
the variable to the specified default. If you have an addon like this you will probably need to add a dummy file
to the directory to make it a valid addon (as the presence of an addon.meta file does not automatically make a
directory into an addon). By convention an empty file named "dummy.rbl" should be used if you need a dummy file.

* * *

Reading variables is simple.

Rubble provides APIs for reading variables to all scripting languages it supports, see the individual API
documents for details.

To read a variable in template code prefix the name with a dollars sign (`$`) or ampersand (`&`) and optionally surround
the name with curly brackets (`{}`). Using a dollars sign is the usual way, ampersand is a special "immediate expansion"
syntax discussed later.

Examples:

	$VAR
	${VAR}

To read a variable this way it's name must consist of only alphanumeric characters and underscores (`a-zA-Z0-9_`).
If you need to read a variable with non-alphanumeric characters you will have to use a script (so don't use
non-alphanumeric characters in your variable names).

The form with curly brackets is for cases where you need to be exact about what is part of the variable name
and what is not, a common problem when working with templates.

A variable will be expanded at different times depending on where it is:

* In a user template body: Expanded just before the template body is parsed, after the template parameters are expanded
  (so variables in user template parameters are expanded then to).
* In script template parameters: If and when the template decides to expand them, most templates do so as the first thing
  they do.

Variables in template bodies are a special case, the common dollars sign expansion form will only be expanded if they
it not inside a call to another template. For example `$VAR` will not be expanded until the post parse stage, even
though it is part of the body of a pre parse template, because it is contained in the body of a post parse template:

	{!;{#;$VAR}}

This may cause minor problems for thoughtless template authors. Consider the following:

	{!TEMPLATE;EXAMPLE;arg;{#PRINT;%arg}}
	{E;{@SET;TEMP;abc}}{EXAMPLE;$TEMP}

This will *probably* not print "abc", as `TEMP` is not expanded until the post-parse stage, and `TEMP` is (by convention)
the variable used to store temporary values for template language calculations. There is nothing inherently wrong with
using a variable here, so long as you keep in mind that it won't be expanded right away (so you should not use a common
variable name like `TEMP`, use something individual to your addon).

If you cannot use a unique variable (or it would be more trouble than it is worth, consider `@FOREACH`) you may work
around the problem by using the ampersand "immediate expansion" syntax:

	{!TEMPLATE;EXAMPLE;arg;{#PRINT;%arg}}
	{E;{@SET;TEMP;abc}}{EXAMPLE;&TEMP}

This form will expand as soon as it is parsed, irrespective of whether it is nested inside a template or not. This can
lead to over-eager expansion, so take care.

* * *

As a user you can set config vars in several ways.

The easiest way is to use the web UI, as it has a built-in variable editor that will be presented to you during
normal generation. This editor is not available during any of the independent apply modes though, so if you need
to do something weird you may have to fall back to one of the more advanced (and harder to use) CLI modes.

<div id="ExpandIssues"></div>
Variable Expansion Pitfalls:
----------------------------------------------

Variable and template parameter expansion is fairly simple, but in some rare cases template and variable syntaxes can
combine in horrible ways.

If something that looks like a variable name to the expander turns out to not be a valid name after all it uses the
following rules to determine what to return:

* Valid form, but mapper returns "": Returns "".
* Open char followed by open brace, but non-alphanumeric char is found before close brace: Makes no changes.
* There are no alphanumeric chars after the open char or inside any braces it may have: Returns the open char.

Take the following example:

	%{CALC_CHANCE}

If `CALC_CHANCE` is the name of a template parameter all is well, but what if it is a template call? This example is easy
to fix:

	%{}{CALC_CHANCE}

The same technique works for variable expansion, just use `$` instead of `%`.

A possible problem with the "immediate" variable expansion form (using `&` instead of `$`) is over-eager expansion.
Take this simplified example from First Landing:

	{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	{@PARSE_TO;TEMP;{@GENERATE_ID;FURNITURE_ASSEMBLER}}
	{!SHARED_INDUSTRY_REACTION;&{}TEMP;FURNITURE_ASSEMBLER;
		...
	}
	}

As you may note I needed to delay expansion here to prevent the outer `TEMP` (the on passed to `@FOREACH`) from clobbering
the inner one. Frankly this wouldn't be a problem in this case if I didn't use `TEMP` for all my temporaries, but other
times the solution is not so simple.

One other expand issue that may cause problems is if you have nested blocks with template parameters, for example a call
to `@FOREACH` inside a template. The outer template will expand all parameters when it is called, clobbering the inner
block's parameters. This may be solved by replacing `%` signs in the inner block with `%{}`. They will be expanded to `%`
in the first pass, "assembling" the expansion statement(s) for the next pass. If you need to nest such templates deeper
than two levels you either need to do something really tricky or you can do the smart thing and define some extra
"helper" templates.

<div id="UserCFGDef"></div>
User Configuration Variable Defaults Override File:
----------------------------------------------

During generation configuration variables will get default values from the addon.meta files. This
works great except for one thing: While the defaults are relatively sane most advanced users will
quickly get tired of changing them to their favorite settings :)

The solution to this is simple: Create a file named `rubble:userconfig.ini` and put
values to override the defaults there. This file uses the same format as `out:genconfig.ini`,
so you can just copy the values you want from that file.

Any values you put in this file will be loaded regardless of addons, so be careful what you put in
this file (for example a variable named "ADDON_XYZ_ACTIVE" is probably not a good thing to add).

<div id="WebUI"></div>
Rubble Web UI Customization:
----------------------------------------------

Presets:

Presets are special ini files that list a set of addons to activate. These ini files are in the same
format used by `addons:dir:addonlist.ini` and are loaded by the web UI from `rubble:presets`. A preset
takes its name from the name of it's ini file. You may not use the name "Default", as it is reserved for
the default preset.

If you do not have any presets installed (the common case) the preset selection box will not be displayed
in the web UI!

To install a preset simply copy it's ini file into `rubble:presets` (creating the directory if needed).

To create a new preset simply copy the `addonlist.ini` file from `addons:dir:` or from the raw directory
of a Rubble generated region into `rubble:presets`. Remember to rename the ini to something descriptive!
The next time you generate there should be a drop-down selection box with the other selection controls at
the bottom of the page. Simply select your preset and it should activate all the addons that are listed
as active in the preset.

I do not expect that this feature will be commonly used, it is intended for those who use Rubble to generate
the same 2-3 setups over and over (such as I do when testing), but it will also be nice for those who like
to play a common selection of different addon configurations.

* * *

The Browser Startup Script:

This functionality is not needed or used on Windows.

Linux and OSX versions of Rubble use a "browser startup script" to start your web browser when
the web UI starts. To use this functionality create a file named `rubble:other/webUI/browser`
that starts your web browser, the server will try to run this file when it starts. This
file will be passed the URL for the main menu as an argument, the script should pass this
argument through to the browser (so the page opens automatically).

Rubble comes with a default browser script, so things *may* work just fine right out of the box.
The included script should work on most Linux systems (you will have to mark the script executable).

OSX users are more or less on their own (sorry).

* * *

HTML Template Customization:

When the web UI server is run it tries to load HTML templates and CSS files from `rubble:other/webUI/`,
if it does not find the file it is looking for it tries to write the default version of the
file out to that directory.

If you want to customize the look of the UI modify these files, from then on the server will
use your files.

Each page template receives raw data about whatever it is supposed to do, most pages receive a
reference to the entire Rubble state (which is usually well over 7mb of data!) so it is important
to write your templates so that they only send what is really needed to the client.

Each template must act a certain way (as observed by the server), so there is not too much wiggle
room when it comes to changing functionality, but appearance can be changed in any way you like.

For more details on how these templates work see `golang.org/pkg/html/template/` and
`golang.org/pkg/text/template/`. For details about exactly what Rubble passes each page you will
have to look at the source code for Rubble and the web universal interface action.

<div id="WebUILink"></div>
Rubble Web UI Special Links:
----------------------------------------------

The web UI supports the following special link types:

* Raw AXIS file links
* Workshop preview image links
* Addon file links
* Documentation file links
* Addon auto-links

* * *

Raw axis file links take an AXIS path and return whatever is at that location.

For example:

	/axis/df%3adata/art/curses_800x600.png

Feeding this URL into an image tag produces the following:

![800x600 default tilesheet](/axis?path=df%3adata/art/curses_800x600.png)

Older version of Rubble required you to pack the path into a query, this is no longer required
(but it is still supported):

	/axis?path=df:data/art/curses_800x600.png

If you use the new path syntax you will have to remember to escape special characters, particularly
colons! (use `%3a` for them). The primary benefit of the path syntax is that web browsers will give
downloaded files the proper file name.

* * *

Workshop preview images are automatically generated on request and served at the path "/wshop"

	/wshop?addon=Base&file=building_custom.txt&id=SOAP_MAKER&stage=3

This example, when fed into an image tag, produces:

![](/wshop?addon=Base&file=building_custom.txt&id=SOAP_MAKER&stage=3)

"stage" is optional and defaults to 3. This link will return an image of the workshop rendered with
the current full screen graphics tileset. The raws are not run through the Rubble parser, so templated
building definitions may not work. The workshop parser will compensate for `BUILDING_WORKSHOP` and
`BUILDING_FURNACE` templates (and other templates with a similar structure) so the majority of workshops
will work fine. If you use `MAT` instead of a color in your workshop definition it will act like "7:0:0".

If you like you may specify a file name after "/wshop" to trick web browsers into using a sensible name for
the images if you try to save them.

	/wshop/soapmaker.png?addon=Base&file=building_custom.txt&id=SOAP_MAKER

This is recommended, but not required.

These workshop images are PNG format images, with the image resolution determined by the size of the tileset
and the workshop.

* * *

If you need a specific file from a known addon you can use an addon file link:

	/addonfile/addon.md?addon=Base

These are much safer than a raw AXIS link, as it is much easier to be sure that you have the correct path, and you do not
need to worry about all your links breaking simply because someone didn't follow your install instructions exactly (provided
you set a static addon name in the addon.meta file).

As a legacy feature you may also specify the file name in a query:

	/addonfile?addon=Base&file=addon.md

For several reasons this is no longer recommended.

* * *

Documentation file links are special in that they are intended for linking text, HTML, or markdown documents only and
they do more for you.

The following link will display this file wrapped with a standard header and footer (if you are reading this in the
web UI then this link gives the page you are currently viewing):

	/doc/Rubble Basics.md

When written like this a documentation link first tries to find the requested file in `rubble:other`. If that fails it
tries `addons:dir:` and finally `addons:zip:`. This search behavior is very much a legacy feature, nowadays if you want
to link to a file in an addon you would provide an addon name in a query:

	/doc/addon.md?addon=Base

If the file extension is `.md` or `.text` the file is assumed to be markdown, if it is `.html` or `.htm` it is assumed
to be HTML, otherwise the file is assumed to be preformatted plain text.

Documentation links are important for one simple reason: Rubble makes a concerted effort to make your browser back button
inoperable because going "back" to certain pages (or from certain pages) can break Rubble in subtle and annoying ways.
Rather than deal with bogus bug reports I simply kill the cause. This means that unless a page provides a way to get back
to where you were you can get stuck. By default the header and footer of a documentation page contains two buttons, one
to go to the main menu and one to go back to the page you reached the doc page from.

Keep in mind that if you use a HTML document for a doc page it will be injected into the middle of an existing HTML
document. For this reason you should not include anything but the body of the document (and don't include the "body"
tags!).

* * *

In markdown documents addon names may be autolinked with their addon information page by surrounding
the addon name with back ticks (`` ` ``) and prefixed with `addon:`. For example:
`` `addon|Example/Addon` `` (replace the '|' with a ':')

* * *

This isn't really a link type, but on doc pages and in addon description pages Rubble provides support for "spoilers",
sections of the document that are hidden by default and toggled on and off via a button. To enable this support simply
use the following markup:

	<a href="javascript:Toggle('<ID>')">Spoiler: Show/Hide</a><div id="<ID>" style="display: none;" class="spoiler">
	
	Your content ....
	
	</div>

`<ID>` needs to be a unique ID. Avoid anything starting with `FILE:` or anything like `More`, `MoreData`, `Back`, `Menu1`
or `Menu2`, as those IDs are used by Rubble. Pick something descriptive, as that is less likely to clash with existing IDs.

This works in both markdown and HTML documents.

<a href="javascript:Toggle('ExampleSpoiler')">Spoiler: Show/Hide</a><div id="ExampleSpoiler" style="display: none;" class="spoiler">
See? It works just like this.

Don't worry, formating the contents of the block with markdown works fine! (BTW: Apparently this is not compliant with
the markdown spec, but in this case it's a good thing)

All Rubble does to support spoilers is add the `Toggle` function and the `spoiler` CSS class. If you don't like the way
these spoilers look or act it is easy to provide your own implementation and/or tweak the existing one in some way (use
custom CSS maybe?). Remember: If you want to do something in a markdown document that markdown does not support you can
always drop into HTML. Any HTML tags in a markdown document are passed through unchanged (even the markdown specification
has embedded HTML in it!).
</div>

<div id="AddonMeta"></div>
addon.meta:
----------------------------------------------

Addons may include a special file named addon.meta that contains extra information about the
addon. This file is optional but **highly** recommended. So many features depend on the addon.meta
file that an addon that lacks one will be half crippled.

Some fields are inheritable by any addon "below" the current addon in the addon tree. For this reason you may place
addon.meta files in directories that are not addons (so that their values may be inherited by their children).

Example addon.meta file (with lots of helpful comments):

	{
	# Example addon.meta file (this file is JSON encoded)
	
	# Yes, my JSON files really use `#` as the comment char. '//' and '/* */' are not supported.
	# Just be glad I hacked comment support in at all, the parser does not support them by default.
	# Comments are only supported on a line of their own.
	
	# This example file includes every possible key. Most addons will only need a subset of these keys.
	
	# Addon tags, the ones listed here are those with hardcoded handling (but you can specify anything you want).
	# The hardcoded tags get special handling various places, for example on the web UI addon information page they will
	# be followed by a short description of their meaning.
	"Tags": {
		# Addon is an automatically managed library.
		"Library": false,
		
		# The addon is not actually an addon, but a documentation pack.
		"DocPack": false,
		
		# Automatically set by the addon loader if the addon has tileset information.
		# You should not set this manually.
		"TileSet": false,
		
		# Automatically set by the addon loader if the addon has test files.
		# You should not set this manually.
		"HasTests": false,
		
		# The addon is safe to apply to existing regions.
		"SaveSafe": false,
		
		# Does this addon require DFHack?
		# If you do not set this explicitly then Rubble will guess the value based on if the addon contains any
		# files tagged "DFHack", but if this is explicitly set your value will not be overridden no mater what.
		# This is used by the web UI to allow the user to hide addons that require DFHack.
		"DFHack": false,
		
		# Set for addons that are of interest mostly to developers, this is used by the web UI addon filtering system
		# so it can hide these addons by default.
		"Dev": false,
		
		# Addons with this tag are not standard addons. Currently causes the addon to not be listed as an option
		# during standard generation cycles when using the web UI. Use if you have a special addon that should only
		# be used with independent apply mode or such like.
		"NotNormal": false
	},
	
	# The addon priority. Lower values load first. If negative it will try to use the same value as the parent addon, if
	# there is no parent the the absolute value will be used. By convention `25` is used for standard libraries, `50` for
	# base addons and `75` for other default addons. Most addons should use the default of `-100`
	"LoadPriority": -100,
	
	# Addon name, optional. If the first character is "$" and the name is longer than two characters then the given name
	# is prefixed with the name of the parent addon. The default name is `<parent name>/<directory name>`.
	# 
	# If you set this to "-" it will cause the addon to have no name. This is a useful technique for organizing addon
	# trees without making the addon names too long. Use this only in addon.meta files that are not in valid addons!
	#"Name" = "Example/Addon",
	"Name": "$/Addon",
	
	# The version and author as strings for display in the web UI.
	# If not provided these will default to the same as the parent addon.
	"Version": "1.0",
	"Author": "Your Name",
	
	# One line description of the addon.
	# This must always be markdown (embedded HTML is OK).
	"Header": "This is an example addon.meta file.",
	
	# Optional longer description loaded from a separate file.
	# The given description file must be contained in the same addon as this meta file.
	# If the file extension is ".htm" or ".html" the file is assumed to be HTML, if it is ".md" or ".text" it is assumed to
	# be markdown. Any other extension is assumed to be preformated plain text.
	"DescFile": "example.md",
	"Description": "If you prefer you can have the description inline.\nObviously this can be convenient, but only for short descriptions.\nIf DescFile is set this is ignored (or rather, overwritten).\nInline descriptions are always assumed to be markdown.",
	
	# If you wish you can specify configuration variables here.
	# These variables will automatically be set to their defaults when addons are activated, but users can
	# change their values in various ways.
	"Vars": {
		# The field name is the name of the variable to set.
		"TEST": {
			# Name is the name to present to the user, should be short and clear.
			# Setting Name to "-" will cause the variable to be hidden in most interfaces.
			"Name": "A test variable",
			
			# Values is an array of possible values. If more than one item is present then the variable will
			# be restricted to just those values. If only one value is provided then you may set the variable
			# to any string. In any case the first value will be the default. If no values are provided then
			# the variable will default to the empty string.
			# Note that Rubble itself does not enforce any value restrictions you set here (outside of the
			# web UI config var editor), so you should always validate any variables (with a script or
			# something) before you use them.
			"Values": ["YES", "NO"]
		}
	},
	
	"Activates": [
		"This/Is An/Array",
		"Of All/The",
		"Addons/This One/Depends On",
		"Anything/Listed/Here",
		"Will Be/Activated/When",
		"This/Addon Is"
	],
	
	"Incompatible": [
		"This/Is An/Array",
		"Of All/The",
		"Addons/This One/Is",
		"Known/To Be",
		"Incompatible/With",
		"Anything/Listed/Here",
		"Will Cause/Rubble To/Abort",
		"If It Is/Active",
		"When This/Addon Is"
	]
	
	# Complicated, see the section on the WebLoad system (farther down).
	"Download": [
		#"Addon Pack": "http://download.url/Addon%20Pack.zip"
	]
	}

<div id="WebLoad"></div>
Addon WebLoad System:
----------------------------------------------

Rubble supports two ways to automatically download files from the internet: .webload files, and the "Download" addon.meta
key.

URLs used for web loading *must* be complete, eg `www.example.com/somefile.zip` will not
work, use `http://www.example.com/somefile.zip`.

The following facts about the WebLoad system should be pointed out:

1. The WebLoad URL *must* always remain *constant* to enable the automatic update of your addon. This makes it a very
   bad idea to include the addon version in the download URL (by, for example, including it in the addon pack file name)
   when you want automatic updates to be possible.
2. The URL must be to a direct download of the file, this effectively rules out most file upload services.
3. You must be able to upload a new version of the file without changing the URL, this rules out most of the upload
   services that point 2 doesn't rule out.

Luckily DFFD is compatible with all three of these points, so use that. After all DFFD is dedicated entirely to Dwarf
Fortress and has no annoying advertisements (a massive plus considering the state of most other file hosts now-days).

Currently the only valid protocol is HTTP (HTTPS should also work). I have never personally tested this with anything
other than a server running on local host.

* * *

.webload files are intended mostly for users who want to have their favorite addon packs automatically updated. All you
need to do is create a file with the `.webload` extension and put the URL of the addon pack you want Rubble to download
for you as its contents. Every time Rubble is run it will scan all .webload files, downloading the zip file they point
to if there is no local copy of the zip file or if the local copy does not match the remote copy.

.webload files may only be placed in the root of the addons directory.

* * *

The "Download" addon.meta key is intended for addon developers who want to use third-party library addons. This system
is more flexible than .webload files, but it is also possible to cause Rubble to error out during loading if you are not
careful (by accidentally creating duplicate addons).

This key holds a list of key/value pairs, where each key is the name of an addon *pack* (*not* an addon name!), and the
value is the URL to download the pack from.

After all other addons have been loaded (including those from .webload files) Rubble will make a list of all addon packs
listed in "Download" keys. Any duplicates will be ignored, and which URL is used is undefined. Each listed addon pack is
then download if there is not a matching local file or the matching local file does not match the server file.

Addons in packs downloaded this way may not load other addons in the same manner! If your library packs needs other
third-party addons you should instead publish a list of dependencies for your users to add to their "Download" keys.

Make sure your are consistent with your pack names! If you are not consistent Rubble may download several copies of a
particular pack, causing Rubble to abort due to duplicate addons. Authors of packs that are intended to be loaded with
this system should publish a "canonical pack name", the name that should always be used to refer to this addon pack.
To allow simple manual installation this name should match the name of the file you get if you download the file with a
web browser.

Remember: Packs listed here are always be downloaded. As Rubble has no idea what addons will be "active" during loading,
it cannot make any assumptions about what will (or will not) be needed.

<div id="DocPack"></div>
Doc Packs:
----------------------------------------------

Sometimes it is nice to be able to have extra documentation for something beyond what is convenient
to put in an addon description. This functionality is provided by "Doc Packs", addons that have the
"DocPack" tag set to true in their addon.meta files.

By setting this tag you cause several things to happen:

* The addon will not be listed as such, it will not be available in the normal addon lists, and it
  will not be possible to activate it.
* The contents of the `Header` field from the addon.meta file will be displayed on the web UI documentation page.

Thats it.

Obviously this is only useful if the header contains a link to something :) There is no restriction
on what may be displayed, so long as it is valid markdown.

For a doc pack to be loaded it must be a valid addon, you may need to include a "dummy file" to make
Rubble load it properly. By convention dummy files should be empty files named `dummy.rbl`.

<div id="Tests"></div>
Template Test Files:
----------------------------------------------

Rubble comes with a basic framework for running template (and by extension, script) tests. These tests
are short chunks of template code coupled with the expected results.

To run these tests a special operation mode is used. This mode loads all addons and makes a list of all
addons that contain tests, then it loads and activates each test containing addon, runs all init scripts,
runs pre scripts from the active addons, and finally runs each test for the current addon in sequence.
Each addon is tested in a different state, so one addon's tests will not interfere with another addon's
tests.

Each test body is parsed through the standard three stages, just like a normal raw file. Since the templates
are parsed in sequence templates that make heavy use of persistent data such as the script data registry may
need to have data from previous tests in the same addon cleared (tests from other addons will not interfere).
This is unlikely to be much of a problem in practice.

The test file format is very simple. For example the following is the simplest possible test:

	>>>:===<<<

This is a perfectly valid (albeit useless) test. A (slightly) more complicated example:

	>>> Example Test:
	{ECHO;Something...}
	===
	Something...
	<<<

This creates a test with the ID "Example Test". The template code following the opening `>>>` and the ID,
but before the `===` is parsed and the result compared to the text after the `===` but before the closing
`<<<`. Both blocks have leading and trailing whitespace stripped while the tests are being loaded, and the
result of parsing the first block is striped of leading and trailing whitespace before it it compared to
the second block.

Each test file can contain as many tests as you like. Any text outside of a test block is ignored, there
is no need for a special comment marker.

<div id="TilesetAddon"></div>
Tileset Addons:
----------------------------------------------

Tilsets are handled as a special extra generation step. After the raws are generated and just
before they are written to disk Rubble collates all files with both the `TileSet` and `RawFile` tags
and uses the information in them to apply tileset information to the generated raws.

These files are normal raws with some minor differences:

* No "OBJECT" tags are needed, all objects can go in the same file.
* Any tags that do not deal with tile or color data can be left out.

For example a tileset raw file for "creature_amphibians.txt" would be:

	[CREATURE:TOAD]
		[CREATURE_TILE:249]
		[COLOR:2:0:0]

Simple.

Note that while these files use raw syntax they are NOT parsed by Rubble! This is due to the
fact that they cannot be parsed when applying a tileset to existing raws, so to keep things
consistent they cannot be parsed at all, ever.

In general all tileset information will be in a single file (as there is no point in having
loads of little files).

Tileset files are generally automatically generated, to make things easy every time you generate
the raws Rubble will create one in "out:current.tset.rbl" using the information from the current
set of addons. If you already have raws for a tileset simply place them in a temporary addon
and generate this addon by itself to get mappings for the tileset.

If you need to edit the init files you can use the "object names" `AUX_FILE:INIT.TXT` and
`AUX_FILE:D_INIT.TXT`.

If you want to change the tileset of an already generated raw set Rubble provides three ways
of doing this:

* Via the web UI, this is very easy to use, just select your region and addons to load the tileset from.
* Via the command line: `rubble -tileset="<region name>" -addons="<the name of the tileset addon(s)>"`.

Scripts with the `TileSet` tag are run as part of generation AND when a tileset is
added to a pregenerated raw set. These scripts are functionally identical to post scripts.

Files with both the `TileSet` and `TWBTOverride` tags will be concatenated together and written to
"df:data/init/overrides.txt". This allows you to easily make TWBT tilesets without needing to worry
about other addons clobbering your overrides, or better yet, create addons that have overrides for a
few items or some such.

Files with `TileSet` and `ImageBMP` or `ImagePNG` will be installed to "df:data/art/".

Files with both the `TileSet` and `MergeRules` tags contain rules specifying how tileset data
should be merged. In general you should not mess with the rules, but advanced users can do so if
they wish (for example you can extend the rules to allow setting the `PRINT_MODE` for a TWBT
tileset). Documentation of the rule format is [just below](#MergeRules).

If you need (or want) to use a template in a tileset file you will have to parse it by hand.
To parse these files you will need to use `rubble.parse` (or equivalent) from a tileset install
script. Multiple calls to the parser may be required depending on what templates are used.

Please note that the base templates will NOT be loaded when applying a tileset to existing raws!
(and many of these templates depend on thing that would not be present anyway) If you don't want
unhappy users test your tileset addon both during normal generation and in independent application!

It would also be possible to make Rubble parse these files by playing with the file tags some,
but **DO NOT DO THIS!** It will not work when applying a tileset to existing raws!

One interesting little "feature" of the way this system works is that it will work on any raw
set, even ones that were not generated by Rubble!

<div id="MergeRules"></div>
Raw Merger Rule Format:
----------------------------------------------

Rubble has a rule-based raw merger that is, at heart, a reimplementation of a standalone utility called BAMM (Button's
Automated Mod Merger). Rubble's version of this merging engine has a few custom extensions, but it is largely BAMM
compatible.

This merger should be able to load any rules that can be used with BAMM, but most of the time you will want to use the
Rubble extensions, as they can greatly simplify the process of writing a set of merge rules.

The raw merger powers Rubble's tileset installer, the raw syntax checker, and certain script functions and templates.

* * *

You may insert comments by prefixing them with the `#` symbol. Comments and rules on the same line works just fine.

Each line is a rule specifying a tag that is to be merged (or matched). A rule is made up of tag patterns separated by
pipe symbols (`|`). The elements of each tag pattern are separated by colons (`:`) and consist of either text to match
or a wildcard.

There are three basic types of wildcard:

* The "key" wildcard is the dollars sign (`$`). Values in a position with a key wildcard must match values in the same
  position in a merge candidate for it to match the rule.
* The "merge" wildcard is the question mark (`?`). Values in a position with a merge wildcard will be used to replace
  the values in the same position in any merge candidate that matches the rule.
* The "discard" wildcard is the ampersand character (`&`). Values in a position with a discard wildcard will not effect
  matching except by their presence or absence.

A wildcard may be followed by a range statement. A range statement specifies that a wildcard will cover between `min`
and `max` elements. A range statement is written like so: `(min, max)`. The values for `min` and `max` must be valid
integers, if `max` is negative or not a number it is assumed to be infinite. If a range statement only has one argument
`min` and `max` are assumed to be equal. If the max value of a range is less than the min value (but still >= 0) it will
be adjusted to equal the min value. This means that `(2,2)`, `(2)` and `(2,0)` are all equivalent.

For two tags to match when a range is used they not only have to match the rule, they also need to match each other.
This means that they need the same number of elements and they need the same elements for any "key" wildcards.

Each rule may only have one range statement where the min and max values do not match *unless* there is a non-wildcard
element between the range statements. For example the following would not work:

	TEST:&(2,5):&(1,2)

But this will:

	TEST:&(2,5):NONE:&(1,2)

With wildcards there is simply no way to decide which range gets a certain element without hard matches delimiting the
boundary.

It is possible to group rule tag patterns so that they can fill a single slot in a rule. This is done by surrounding the
patterns with curly brackets (`{}`).

Tag patterns and groups of tag patterns may be named and inserted later by name with the following syntax:

	!(NAME)RULE
	%(RULE_TO_INSERT)

Named rules are generally most useful with rule blocks.

Sometimes a certain tag takes a limited number of values for a certain parameter in a tag. If you want to use a limited
number of constants to match it is possible to use a variant of the block syntax:

	TEST:{
		FOO
		BAR
	}

Matches both of the following (but the two tags will not match each other):

	[TEST:FOO]
	[TEST:BAR]

You may not include wildcards inside such "match blocks" (although wild cards may appear before or after). A match block
works as a "hard match" for closing wildcard ranges:

	TEST:&(2,5):{
		A
		B
		C
	}:&(1,2)

Due to ambiguities in the syntax you may not have a match block as the first item in a tag pattern (allowing this would
require the rule parser to support an arbitrary amount of lookahead, aka it would be too much work to be worth it).

* * *

BAMM Compatibility:

If you want your rules to be BAMM compatible then you may not use the following features:

* Range statements with anything other than the simple two value form, this includes any form where `max` is less than `min`.
* Rule blocks
* Named rules
* Comments
* Match blocks

* * *

Examples:

The following two sets of tags will match under the following rule:

Rule:

	EXAMPLE:$|SOMETAG:&:?(2)

Tags A:

	[EXAMPLE:TEST][SOMETAG:ABC:1:2]

Tags B:

	[EXAMPLE:TEST][SOMETAG:XYZ:3:4]

Result of merging A into B:

	[EXAMPLE:TEST][SOMETAG:XYZ:1:2]

<div id="LuaMod"></div>
The Rubble DFHack Lua Pseudo-Module System:
----------------------------------------------

DFHack can load commands from the save directory, but if you need to load a module you still
need to install it globally. It is possible to get around this by using the `dfhack.script_environment`
function, but I find this cumbersome and annoying. "Pseudo-modules" are my solution for this issue.

Pseudo-modules closely parallel the syntax of a standard module for use with the default Lua `require`
function, the only differences are to provide reload support.

The module is mostly the same as with normal modules, the only difference is instead of:

	local _ENV = mkmodule("example")

You have this instead:

	local _ENV = rubble.mkmodule("example")

Simple. (remember to `return _ENV` at the end of the module!)

Pseudo-modules may have two special functions defined. `onUnload` will be called when the world is
unloaded, and `onStateChange` will be called for every state change event.

Clients are mostly the same as with normal modules, the only difference is instead of:

	local example = require "example"

You use:

	local example = rubble.require "example"

Obviously, due to the way they are loaded, pseudo modules will only work when a save is
loaded, but since they are only for use with region-local scripts this is not an issue.

Any file with the `DFHackModuleScript` tag will be installed as a pseudo-module.

<div id="TweakScript"></div>
Tweak Scripts:
----------------------------------------------

Tweak scripts allow you to run arbitrary scripts before or after generation. These scripts can change
the raws, allowing you to "tweak" the generated output.

* Pre scripts have the `PreScript` tag and are run just before parsing starts.
* Post scripts have the `PostScript` tag and are run just after parsing ends.

These scripts may be written in any language Rubble supports. Currently this means Lua or JavaScript.

<div id="InitScript"></div>
Initialization Scripts:
----------------------------------------------

Init scripts are very limited, but also very useful. Basically they are scripts that ALWAYS
run, even when their containing addon is not active.

* Init scripts have the tags `InitScript` and `GlobalFile`.

Use init scripts to do setup that may be shared by multiple addons. The standard addons use init
scripts to install some standard script functions.

Init scripts have a very narrow use-case and should be used with care. Most of the problems that
init scripts were created to solve have been made by obsolete by one improvement or another. The
default addons only contain a few init scripts.

<div id="LoadScript"></div>
Loader Scripts:
----------------------------------------------

Loader scripts are a lot like initialization scripts, just that they run earlier.

* Load scripts have the tags `LoadScript` and `GlobalFile`.

Loader scripts only do two things that init scripts can't do better: they can modify the
activation state of addons and they can modify the files of a specific addon without effecting
other addons. Since loader scripts run just after addons have been activated but before the
active file list has been generated you can examine what addons are active and change this
list if need be.

Like init scripts, loader scripts are global and pay no attention to their containing addon.

During a normal generation cycle the addons can be loaded more than once on the same state,
so the load scripts will be run several times as well. If you need to do something in a load
script that may only be done once it is possible to set the file tags to skip the file in the
future (use the `Skip` tag).

Loader scripts have a *very* narrow use case, almost everything they were designed to do
has been superseded by other, better, systems. That said they are needed for some rare
advanced actions, but in general you do not need to worry about them.

<div id="ScriptMod"></div>
Script Modules:
----------------------------------------------

Some scripting languages support "modules", libraries of script code that may be loaded on demand.

Rubble supports these for languages that use them. Any file with the `ModuleScript` tag will be
flagged for consideration as a script module.

Exactly how these modules are made available to scripts depends on the scripting language.

Lua allows you to load modules with the default `require` function. The modules are made available
under their file name minus the extension (by default the extension is `.mod.lua`).

JavaScript does not currently support modules (it may later).

<div id="AddonLoader"></div>
The Addon Loader:
----------------------------------------------

Rubble has a recursive addon loader that can read addons from zip files or directories.
What this means is that you can group addons in directories to have something like an addon tree.

The following rules determine how addons are loaded:

Addons are loaded in this order:

1. Global scripts (files with the `GlobalFile` tag) are read from `addons:dir:`.
2. All addons are loaded into memory:
	1. Addons are loaded from `addons:dir:` in lexical order.
		* Global scripts in non-addon subdirectories of `addons:dir:` are also loaded at this point.
	2. Addons are loaded from `addons:zip:` in lexical order.
		* Global scripts in non-addon subdirectories of `addons:zip:` are also loaded at this point.
3. Addons are sorted by their addon.meta `LoadPriority` key, lowest to highest, and if priority is equal
   in lexical order by name.
4. Global scripts in addons are added to the global file list in addon load order (priority, then
   lexical) with files from addons overriding files outside of addons.

The following additional rules apply:

* Zip files that are in the root of `addons:dir:` are used as subdirectories for `addons:zip:`.
* For those who don't know, lexical order means that capital letters come before small letters
  ("Zebras" comes before "ants").
* Addons are loaded recursively, eg you can have addons inside addons.
* If an addon does not have a `LoadPriority` key in it's addon.meta file it's priority is inherited from
  it's parent. If it does not have a parent it's priority is 100.
* Global scripts outside of addons are treated like they have a lower priority than the the lowest
  possible priority number (so any file in a proper addon will override them).

If two active addons have a file with the same name whichever addon is last in the list is the one that "wins".

If you have readmes or the like in an otherwise empty addon it is recommended that you do not
give them a file extension. In any case you do not want such an addon appearing in the addon
list, so do not use any of the "parseable" extensions (.txt, etc).

A good alternative extension for plain text is .nfo or maybe .me (as in "read.me"), but I
prefer no extension at all. Obviously if your readme is not plain text you will want to use
whatever extension is normally used for that format (.md for markdown for example).

All addons are loaded into memory, but only the active addons are processed.

An addon is active if:

* It is listed as active in the active addon list (exactly what this is differs based on the operation mode).
* It is listed in another active addon's "Activates" addon.meta key.

So why do you (the modder) need to know all this? To put it simply you need to know what is
permissible for structuring your addon(s) so that you can choose an optimal method that makes
it easy on both you and your users. 

Do you want to use multiple small addons or just one large one? Do you want all of your mod's
addons in one large directory, or do you want to group related things together? These decisions
are important part of deciding how you want to package your mod and how you want the addon list
presented to users.
