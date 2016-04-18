
Rubble Changelog:
========================================================================================================================

Rubble versions follow a somewhat odd `rewrite.major.minor` format. The "rewrite" number changes only rarely, generally
when I make sweeping structural changes and/or make a major compatibility break. "Major" versions usually contain new
engine features, while "minor" version usually contain only addon (standard library) changes and engine bugfixes.

For purposes of this changelog Rubble is divided into three parts, the engine (the native code that runs everything else),
the standard library (addons and scripts), and "other" (which generally refers to the documentation). The engine is by
far the most important part as far as the version number goes, almost any engine changes will make a version worthy of a
new "major" version. The standard library, while more visible to most users than the engine, is usually only worthy of a
new "minor" version.

Please read this changelog carefully, all kinds of interesting/critical tidbits tend to get stuck in here.

7.7.1 (For DF 42.6) Minor bug fix.
------------------------------------------------------------------------------------------------------------------------

This version fixes one minor bug and adds one "new" feature.

The "new" feature isn't really new, since anyone could already do what it does if they were willing to write a little
bit of JavaScript. This simply makes it slightly easier to do in a consistent manner.

Engine Changes:

* Fixed a bug in the web load system that prevented the downloaded files from being saved to disk.
* Added a bit of CSS and JavaScript to the addon information and documentation pages to support adding "spoilers".
	* You need to use HTML to add the required tags, but markdown allows embedded HTML, so use that.
	* The end of the [Rubble Basics web UI links section][RB1] has an example showing how this works.

Addon and Standard Library Changes:

* None

Other Changes:

* None

7.7.0 (For DF 42.6) Automatic addon downloading!
------------------------------------------------------------------------------------------------------------------------

This version brings a *huge* improvement to the old .webload system. The new features make it feasible for an addon to
require external third-party addons.

The old .webload system was intended for users mostly, while the new system is intended for modders. Basically I added a
new addon.meta key that allows you to specify addon packs to download and the URLs to use. These packs will then be
downloaded when the addons are loaded. Sadly there is no way for Rubble to know if you will need the packs or not, so they
will always be downloaded (provided you do not already have them). To support offline users (of which I am one) the
downloads are not forced. A failed download will simply be skipped and Rubble will go on (but Rubble will probably abort
later if an addon needed addons from the requested pack). Offline users can simply go get the required files manually.

Engine Changes:

* Added a "Download" key to addon.meta, see [Rubble Basics][RB5].
* Moved the .webload code around some, the log messages are now a little different, but it should still work more-or-less
  the same.
* Fixed the workshop preview image renderer to properly work with the shared object versions of the building templates.

Addon and Standard Library Changes:

* Fixed a bug in `addon:Libs/Items` that prevented many powered workshops from working properly.

Other Changes:

* Documented the new addon.meta key in Rubble Basics.

7.6.0 (For DF 42.6) Tweaks!
------------------------------------------------------------------------------------------------------------------------

This version contains a minor tweak to allow a single script runner to support multiple file tags. This allows proper
support for precompiled scripts and such like. I also threw in an experimental alternate Lua syntax I have been playing
with lately. This may or may not stick around, as I included it in Rubble simply so it is easier for me to experiment with.

I have a working implementation of a `continue` keyword for my Lua compiler. For now it is disabled, but it may show up
later if I get tired of working around such a glaring missing feature... (in other words: don't use `continue` as a
variable, function, or label name unless you don't mind having to change it later!)

I fixed several raw consistency checker rule errors for this version. While perusing logs created by running the checker
on various third party mods I was amazed how many mods contain invalid raws, usually in subtle ways. The most common
culprits are tags with missing parameters or extra parameters (caused, I assume, by sloppy copy and paste and not looking
unfamiliar tags up on the wiki). Hopefully this tool will eliminate that problem for Rubble users at least (if only
because they get tired of the nagging). As a reminder: If you have a tag that looks fine reported as a rule match failure
first check the wiki to make sure you used it properly, a lot of tags have non-obvious parameters and/or need to come
after a certain other tag.

While working on First Landing I encountered some problems with variables not expanding soon enough. This problem is
caused by the fact that, by default, variables do not expand inside nested templates. Since this behavior is required,
I decided to leave things as-is and simply add an extra step to template body variable expansion. The extra step expands
variables prefixed with `&` using the same mechanics used for expanding template parameters. This was not a problem in
earlier version of Rubble (versions before 7.0.0) because back then variable expansion in parameters was under the control
of the template dispatcher, where as now it is controlled by the templates themselves. The new expansion style has a high
foot-shot potential, and so should only be used where it is really needed.

I wrote a draft style guide with lots of information on recommended naming conventions and such like. I would welcome
suggestions and criticism from any modders, as this guide should detail the consensuses of the community, not just my
preferences. Give it a read and tell me what you think!

I finally got tired of having to use `@SCRIPT` if I wanted to do simple math on template parameters, so I threw in a
simple set of basic math operation templates. These templates are rather slow (compared to scripts), but if you just need
to add, subtract, multiply, or divide a few numbers they are a real time saver. To allow nesting these templates do a
full parse of their parameters (which makes them much slower, but also much more flexible)

Engine Changes:

* Changed the way script runners work slightly to allow a single runner to handle multiple script types.
* Pre-compiled Lua scripts now need to use the `LangLuaBin` file tag for the Lua runner to recognize them. This tag is
  auto-applied to files with the last part extension `.luab`.
* Added basic support for an alternate Lua syntax usable in scripts that have the file tag `LangLuaC`.
	* This is to be considered *highly* experimental and is likely to go away once I get tired of playing with it!
* The documentation link system now supports formats besides markdown, and allows specifying an addon to restrict the file
  search to. If you specify an addon then do not specify the addon name as part of the file path.
	* This change necessitated minor changes to the way the doc link system works, in particular you will want to specify
	  the file extension now.
* Added alternate modes to many of the web UI link types. Generally you may now specify the file name as part of the path
  instead of in a query. The query style still works (and will continue to work), but the path style is generally better.
* Added a new style of variable expansion: Use `&` instead of `$` to expand configuration variables immediately irrespective
  of whether or not they are contained in nested templates (same basic rules as template parameters).
	* This is useful for passing variables as parameters to user templates that then use the variables in later-stage
	  templates. This situation becomes a problem when doing complex interactions with templates and configuration
	  variables, particularly those that use @PARSE_TO, @FOREACH, @STR_SPLIT, and other templates where temporary
	  variables are commonly used.
* Changed every place I could find that expanded variables to allow the new style. I hope I didn't miss any...
* INI files may now have quoted strings as values (this allows embedded newlines in the form of escape sequences).
* The configuration variable dump code now quotes the values.
* Fixed addon auto-links not working on documentation pages in the web UI.
* Made the web UI "back-button killer" more robust.
	* This makes addon auto-links work properly again (before they worked, but not the way they were supposed to).
	* Any link that is not "killed" properly will now fail to do anything, making any problem more visible.
* Changed the way web UI addon information links work slightly so they now allow you to "go back" in the same way doc
  links do. It's not perfect, but it's a lot better than it was.

Addon and Standard Library Changes:

* Fixed several raw consistency checker rule errors.
* Added several new standard templates: `@STR_SPLIT`, `@MUL`, `@DIV`, `@IDIV`, `@MOD`, `@ADD`, and `@SUB`.
* Added several new standard Lua script functions: `rubble.expandargs`, `rubble.getarraypacked`, and `rubble.inverttable`.
* The standard template `@STR_TO_ID` now replaces colons with underscores.
* Fixed a few broken workshop preview links in forgotten corners (`addon:User/Warcrafter` and `addon:User/Quarry` to be exact.)

Other Changes:

* Added a short section to the [Lua API documentation][RLUA] describing the alternate "CLua" syntax.
* Made some minor tweaks to [Rubble Basics][RB].
* Added a draft proposal for [Rubble Conventions](/doc/Rubble%20Conventions.md), please read and post comments/suggestions
  in the Rubble 7 thread.
* Added documentation for the new templates and script functions.
* Added a bit explaining some RTL (Rubble Template Language) basics to the beginning of the standard template documentation.

7.5.2 (For DF 42.6) More bug fixing...
------------------------------------------------------------------------------------------------------------------------

This version fixes the powered workshop addons that I added in 7.5.0 so they actually work. I had forgotten to account
for the removal of "DFHack" from the names of DFHack dependent addons and the files added by such addons, leading to the
scripts not being able to find the modules that they were looking for. Also I had forgotten that I changed the way the
powered workshop registration function worked. Both were stupid mistakes that would have been discovered immediately had
I not been in too much of a hurry to do in-game testing.

Anyway, I ran a dwarf fort for a while, and everything seems to work now... As an interesting change of pace I finally
got to load and use a Rubble addon made by someone other than me! Dirst's `The Earth Strikes Back!` addon made an
interesting (and fun, in all possible interpretations of the word!) addition to the normal dwarven antics. If you want
something to make mining more... "interesting" you should install the Rubble version of TESB and give it a try!

I finally got tired of reselecting my addons when ever I wanted to switch from dwarves to First Landing (or Saurians, or
Gnomes, or Humans... I play a lot of different races), so I threw together a "presets" system for the web UI. This system
is fairly simple, but it works quite well for its intended purpose.

I also ported the old generic animal materials addons from Rubble 6, so those who want some or all of the creature
materials to be generic (generally for FPS reasons) now have a simple way to do that. These addons should work with most
bases, but they may have minor problems with addons that define non-vanilla tissues and/or use non-standard tissues that
refer to materials not listed in the `STANDARD_MATERIALS` body detail plan. AFAIK these addons will work just fine with
First Landing, but I have not tested this...

Engine Changes:

* The web UI now writes its HTML templates and other data files to the `rubble:other/webUI` directory again. Not sure
  when or how this broke.
* The web UI now allow you to specify "presets". These are ini files of the same format as `addons:dir:addonlist.ini`
  that are placed in `rubble:presets`. The file name is used as the preset name. You may not use the name "Default", as
  that is reserved.
	* If you want to use preset you will have to configure them yourself, they are an optional feature intended for those
	  who use Rubble to ply several common configurations and are tired of needing to select their addons again and again.
	* If you do not have any presets the UI elements related to them will not even show up.
* Fixed a Lua API bug that prevented replacing raw parser tag reference parameter fields (`tag.Params`) with Lua tables.

Addon and Standard Library Changes:

* Fixed loads of incorrect pseudo-module names and missing "power produced" arguments to the `Register` function in the
  powered workshop addon DFHack scripts. The only addons in the `addon:User/Powered` group that were not effected were
  the ones that were ports of the versions made for `First Landing`.
* Castes created with `addon:Libs/Castes` that are defined multiple times are now only inserted once.
	* This broke way back in 7.0.0, but I never used any of the caste addons (too busy playing First Landing), so I
	  never noticed.
* Added some missing creature and creature variation tags to the raw consistency checker rules.
	* I am probably still missing a few...
* Ported the old `addon:Generic Animal Mats` addons from Rubble 6. These addons use scripts to rewrite creature tissues
  and some material templates to point to generic materials. The scripts themselves are kinda hackish, but they work for
  the vast majority of creatures (certainly the vanilla ones).
	* These addons may or may not work with non-vanilla bases, it depends on if they use use custom tissues and material
	  templates or not.

Other Changes:

* Added a bit to the web UI customization section of Rubble Basics about presets.
* Added an "Advanced Installation" section to the readme that details some of the optional configuration steps in simple
  terms.

7.5.1 (For DF 42.6) Interface rework and minor bug fixing.
------------------------------------------------------------------------------------------------------------------------

This version is mostly tweaks and bug-fixes.

Well... If you used the CLI interface a lot it could be considered a big change :P Basically I got tired of needing to
write a new interface every time I added something new, so I turned the old CLI interface into a generic host for
externally supplied "actions". I got most of the way finished with unifying all the CLI stuff and realized that there is
no reason why I couldn't build the web UI into it as well, so all the interfaces are now unified into a single program.

Web UI users will notice little difference (just run `rubble` rather than `rubble_web`), but CLI users will have to
manually specify the operation mode now. Of course there is no reason why you can't build the old style single-action
interfaces if you want, but personally I really like the universal interface (and I won't be maintaining the old ones).

Since all the interfaces are now unified it is feasible to provide 64 bit binaries in addition to the 32 bit binaries.
Before it would inflate the download too much, but now that I only need to provide a single binary for each OS/architecture
combination it is much more feasible. Even with the 64 bit binaries the download is now ~30mb smaller (after decompression,
compressed it's only ~1mb).

I got so busy reworking the interface I almost forgot the feature that prompted the whole thing! Luckily I remembered
after a bit, so there is now a standalone raw merger mode. I also threw in a simple Lua compiler mode while I was at it.

Today marks a historic occasion: A Rubble tool has discovered a vanilla DF bug! In creature_standard.txt there is a
`PHYS_ATT_RANGE` tag that has one argument too few (I assume this is a bug, since every other such tag has 8 arguments,
and this one has 7). This bug was discovered due to the latest set of rules for the raw consistency checker. It is quite
likely that kobolds do not have the correct ranges for their agility attribute :) For now every Rubble run will report a
warning related to this issue.

Engine Changes:

* Raw files with two part extensions (creature graphics raws) will now properly have both extensions stripped from the
  generated file IDs.
	* This did not cause any problems, but it was rather sloppy (and the fix was trivial), so I fixed it.
* Made the following changes to the Lua VM:
	* Line numbers are correct now (they previously always started from zero no matter what you passed in as the initial
	  line number).
	* Compiler errors are, in general, much more detailed, plus they at least attempt to give the problem location now.
	* It should be possible to load binaries compiled with `luac` on 32 bit systems now (so long as the number types are
	  64 bit, which is the default).
* Changed the way interfaces work: There is now a single generic interface program that handles all operations. What
  effect you get depends on how the interface was built (different defaults) and what command line options you use.
	* Currently the default action when running the universal interface is to launch the web UI server.
	* The old CLI interface behavior is provided by the following modes: "generate", "tset", and "iapply".
	* The test interface functionality is provided by the "test" mode.
* The universal interface now provides a raw merger mode (mode `merge`, run `rubble help merge` for more info).
* The universal interface now provides a Lua compiler (mode `luac`, run `rubble help luac` for more info).
* I now provide 64 bit binaries by default. 32 bit binaries are still included (and will continue to be included in the
  future), but you will need to use `rubble32` instead of `rubble` on 32 bit systems.
* While working on consistency checker rules I found myself using a certain bit of syntax that was invalid but reasonable,
  so the match/merge engine rule format now supports specifying "match blocks" inside tag patterns.
* The raw consistency checker was (ironically) ignoring errors in its rule files, resulting in **every single tag**
  being reported as a match failure if there were *any* syntax errors. Needless to say this made Rubble produce *gigantic*
  log files if there were syntax errors in the rules!

Addon and Standard Library Changes:

* I added some more rules for the raw consistency checker.
* Deep gnomes (`addon:User/Gnome`) now have the proper resistance to alcohol and may have strange moods.
* Saurians (`addon:User/Saurian`) may now have strange moods.

Other Changes:

* Updated the readme and some of the tutorials to reflect the new universal interface.

7.5.0 (For DF 42.6) Lua 5.3!
------------------------------------------------------------------------------------------------------------------------

This version is the big one! 

The new Lua 5.3 VM is now powering all Rubble Lua scripts. This VM and it's compiler are not exactly paragons of
stability, but I have tested them fairly well, certainly enough to ensure that core Rubble functionality works as
intended. In any case expect bugs, both major and minor, for a few versions yet as I work out the final kinks.

Writing the VM was a fun and valuable learning experience (although I was thoroughly sick of it by the time I finished
bug-fixing). I am glad I took the time to do it, certainly any future languages I design will greatly benefit from what
I learned. If nothing else I learned a lot about writing more efficient Lua code, and I am now one of an elite few who
truly understand Lua assembly (now *there* is a useless skill!).

The old VM was Lua 5.1, and this VM is Lua 5.3, so there are some unavoidable compatibility issues. Most issues are minor,
with only two worth worrying about. The first problem has less to do with the VM version than it does the Rubble API
rewrite. In the API for the old VM any references to internal data structures used 0 based indexing, in the API for the
new VM all indexes are 1 based. This means that any scripts that interact with certain parts of the Rubble API may need
small tweaks to account for the difference. The second problem relates only to script modules, namely module syntax
changed in Lua 5.2. Any script modules will need to be changed to account for the different module syntax. The good news
is that Rubble now uses the same basic system DFHack uses, so it should be more familiar.

The Rubble Lua script API (obviously) required a total rewrite to interface with the new VM, but the new API is almost
100% compatible with the old. Aside from the aforementioned switch to 1 based indexes the only other noticeable changes
are that `rubble.template` can now accept a function as a template body and `rubble.random` may be seeded with a number
directly instead of needing an hexadecimal string.

I ran generation tests for every addon, both in the base Rubble distribution and any other addons I had laying around
(namely First Landing, which has *way* more scripts than all the default Rubble addons put together). All tests pass,
and raw spot checking plus some in-game region generation tests do not reveal any issues. This does not guarantee that
the VM/compiler work in all cases, but everything Rubble and First Landing use works.

I did not get around to optimizing anything yet, so script performance is lower than it was. Combine this with atrocious
error messages and I have plenty to do in the future :P

I was asked about case-insensitive file extensions, and my first thought was "oh, that'll be easy, just convert to
lowercase in the file tagger." Well it turns out to be much more complicated than that, far too complicated to work
without major structural changes to certain parts of Rubble (mostly the function that handles writing files, but other,
harder to change, places as well). It would be possible to have case-insensitive extensions, but it would be **far** more
work than it is worth. I have always thought that case-insensitivity was sloppy anyway, so this feature will not be
implemented. I have added a few lines to the applicable documentation section making the case sensitivity of the Rubble
engine explicit.

I finally got around to porting the powered workshops from Rubble 6 (I should have done this a long time ago). While I
was working on them anyway I took the time to combine some redundant addons together (the block cutter and bone cutter
are now part of the sawmill for example). Additionally the "logistics" addons (belts, sorter, and the cart loader/launcher)
are combined into one addon and the buildings are added to the machines page instead of the workshops page, just like they
are in First Landing.

Engine Changes:

* New Lua VM and API!
	* The new API is largely compatible with the old API, see the [Lua API documentation][RLUA] for details.
	* The compiler is custom, but it produces code that is *mostly* identical to the output of the reference Lua compiler.
		* Most constructs compile to code identical to that produced by the reference compiler, but some stuff is a little
		  less efficient.
	* Errors messages are *horrible*. Currently an attempt is made at a stack trace for runtime errors (it's not a very
	  good attempt), but compile errors are very vague at best...
		* Better (more informative) error messages are a top priority for future improvements.
* The test interface now always processes addons in load order and test files within an addon in lexical order.
	* This prevents possible weird issues when testing. This hasn't caused any problems (yet), but every part of the
	  Rubble engine *must* be *completely* deterministic, otherwise the "undefined behavior" monster will jump out and
	  eat some poor modder/user, and no one wants that...

Addon and Standard Library Changes:

* Some small tweaks to roughly 15-20% of the scripts to make them work with the new API.
	* Most of the scripts required no changes, and those that needed changes only required tiny tweaks.
* Finally ported the powered workshop addons from Rubble 6. You can find these addons in the `addon:User/Powered` addon
  group (see `addon:Libs/Powered` as well).
	* I haven't had time to do any in-game testing for these addons, but since the DFHack side of things didn't receive
	  any changes they should still work fine.
* The workshop generation API in `addon:Libs/Powered` now provides support for 5x5 workshops in addition to the existing
  support for 3x3 and 1x1 workshops.
	* The way you set the center tile(s) has changed, see the comments in the module file.

Other Changes:

* Deleted the old Lua API documentation.
* The "experimental" Lua API documentation has been "promoted" to being the normal version.
	* I rewrote the header section to reflect it's new status.

7.4.2 (For DF 42.6) Colors!
------------------------------------------------------------------------------------------------------------------------

This version is mostly in support of the newest First Landing version, which does some cool stuff with colors. The new
addons make it very easy to deal with colors, but frankly the kind of problems they are intended to solve are only likely
to occur when making a total conversion mod with all new color descriptors and/or a custom palette. Since most users will
never do this (as this kind of radical reworking is very rare) these addons aren't as useful as you may think.

In addition to the color handling stuff I updated to DF 42.6, although the raw changes from 42.5 are so minor (being
limited to some basic adventure mode crafting) that it almost wasn't worth it :P

I have started testing a version of Rubble that uses my custom Lua VM and compiler. Currently the compiler generates
incorrect output far too often, making the whole process somewhat frustrating (on the plus side, I am now very good at
reading Lua assembly listings...). As far as VM performance goes it seems to be a bit slower than the VM I use now, but
I haven't done any optimizing yet, so by the time I actually start using it in the public version of Rubble it should be
fairly close. I hope to have it ready for 7.5, but we will see.

Engine Changes:

* None.

Addon and Standard Library Changes:

* Updated the `addon:Base` addon for DF 42.6
* Added `addon:Libs/Colors/Swap Palette`, a simple addon to allow per-world color palette settings.
* Added `addon:Libs/Colors`, templates for managing colors. This addon basically allows you to tie DF's two methods of
  defining colors into one, greatly simplifying certain operations.

Other Changes:

* None.

7.4.1 (For DF 42.5)
------------------------------------------------------------------------------------------------------------------------

This version mainly adds a bit more flexibility to `!SHARED_OBJECT_DUPLICATE`, it should now be much more user friendly.
Most common use cases can now do everything required to duplicate an object with a single call to the template, before
only the uncommon simple case could get away with so little effort.

Most of my time right now is being spent on writing a Lua compiler for my VM (and other, non-DF, projects). This is going
a little slower than I had hoped, mainly because, while I have written several interpreters, an optimizing compiler is a
bit harder. Never fear, progress is being made! I really should stop getting new games though, that would probably speed
things up a great deal...

Engine Changes:

* None.

Addon and Standard Library Changes:

* Fixed the pillar tile in `addon:Tilesets/MLC/TWBT - 24px`
* Made the `!SHARED_OBJECT_DUPLICATE` template's raw auto-correction more flexible, namely it can now handle standard two
  part IDs (`<type>:<ID>` as opposed to just the ID).
	* Please note that if you pass an ID with more than one colon or a two part ID where the first part does not match
	  the object ID, the auto-corrector will fail with an error. If you need to do something of this nature you will have
	  to disable auto-correction and correct the raws manually.
* Nuked the temporary missing shared object ID warnings.

Other Changes:

* Updated existing tests and documentation to reflect the change to `!SHARED_OBJECT_DUPLICATE`.

7.4.0 (For DF 42.5) Testing, testing, 1, 2, 3...
------------------------------------------------------------------------------------------------------------------------

This version brings support for basic unit testing. Tests are fairly flexible and easy to write, but also very
simple without many features. The framework works well for testing templates, but it is not optimal for much else.

The testing framework joins the raw consistency checker as a powerful tool for finding errors *before* they escape into
the wild... Nothing will ever be able to replace a long session of play-testing as an error finding tool, but that takes
too much time to do often, so I will take every shortcut I can get. I have found that tests are really nice when working
on a template, I can throw together a bit of code and see immediately if it works or not, no need to run a full generation
cycle. Even better than the faster template development times this allows is the fact that these test can simply be left
in the addon, ready to confirm that stuff still works as expected next version (and the next version, and the next
version...). I never really though much of unit tests (lots of work for minor gain), but they really work well for this.
Unit testing is one of those features that never would have happened if not for my users prodding me to add it, vocal
users are awesome!

I finally got inspired to add a minor feature I have been thinking about adding for quite some time: The new `NotNormal`
addon tag allows you to have addons that can be used in, for example, independent apply mode, but not in normal generation
mode. This is a fairly uncommon case, but it's not like addon tags cost anything... Of course like most such special
behavior tags `NotNormal` only works with the web UI.

Engine Changes:

* Added new default file tag: `TemplateTest`, auto-assigned to the last part extension `.test`.
* Added a parser for template test files.
* Added a new automatically applied addon tag: `HasTests`, applied to any addon that has test files.
* Added a new addon tag: `NotNormal`, addons with this tag are not normal addons in one way or another. Currently this
  causes the web UI to not display the addon as an option during standard generation cycles.
* The web UI recognizes the addon tags `HasTests` and `NotNormal` and annotates the addon with a descriptive string just
  like it does for the other hardcoded tags.
* Added a new interface: `rubble_test` will run tests and report the results. Run `rubble_test -h` for a list of options.
* Removed the final global variable expansion pass.
	* This pass caused some weird and wonderful edge cases and escape-scenarios-from-hell (why does `$100.00` become
	  `.00`? And what do you mean "there is no easy way to fix it"?!).
	* This "feature" hasn't even worked properly for quite some time, it's probably the longest lived feature that has
	  never been used for anything worthwhile, seeing as it was added in version 1.3 and I rarely (maybe never) used it!
	* If you need this behavior just wrap the variable in a call to `#` or `#ECHO`, it's not exactly the same, but it's
	  close enough.
* The experimental Lua bindings now compile cleanly, so they can be used if you are crazy enough to want to try them
  this early...
* Fixed a raw merger bug, when merging range wild cards where min != max the wrong parts of the tag were being matched
  for the items above min.

Addon and Standard Library Changes:

* Added a few tests to `addon:Libs/Base`, lots of templates still have no tests at all, but at least it's a start.
* The `addon:Util/Milo's Settings` addon is now tagged `NotNormal`.
* Changed the way `SHARED_OBJECT_KILL_TAG` and `SHARED_OBJECT_REPLACE_TAG` match tags, they are now much more flexible.
	* This change is fully backwards compatible.
	* You can thank prodding from Abadrausar for this improvement :)
* Added a few new templates to `addon:Libs/Industry`. These templates are basically `!SHARED_OBJECT` variants that tie
  in the industry templates rather than the standard reaction/building templates.
* Added `Libs/Auto-Worker` an experimental system for semi-automatic workshops (they operate without player intervention,
  but require dwarf labor).
	* This library is intended to help implement things like bee hives, fish farms, herb gardens, etc.
* Fixed some reversed key bindings in `addon:Util/Milo's Settings`.

Other Changes:

* Added a section about tests to [Rubble Basics][RB4].
* Added the new file tags and addon tags to the appropriate parts of the standard documentation.

7.3.2 (For DF 42.5) I'm warning you!
------------------------------------------------------------------------------------------------------------------------

This version is mostly just to fix a bug and update the `addon:Base` addon to DF 42.5.

I also threw in a new system for warning messages, it should make them more obvious (since stuff like that tends to get
buried under minor status messages).

In a moment of boredom I did some preliminary work to help make the coming Lua VM change painless. The switch is still
some ways into the future, but it is getting nearer every day.

Engine Changes:

* Creature graphics raws are now written to the correct directory (they were being written to the text directory, oops).
	* Thank you Dirst for reporting this bug!
* Some kindly soul (not the original author) spotted the issue I reported with the Lua VM Rubble uses, and as they had
  the same need (and more willingness to fix other people's code) it is now possible to selectively load (or not load)
  standard Lua modules with that VM. There is absolutely no chance of the "os", "io", and "debug" modules being available
  to Rubble scripts now (before I may have missed a hidden reference or some-such).
	* In the near future I will be switching to my own Lua VM, but until then it is nice to have one less annoyance...
* I extended the Rubble logger so it can have "warning messages" with special behavior. If any warnings are printed they
  will also be saved to print again at the end of the log (together with a message saying how many of them there were).
* Added a new Lua function: `rubble.warning` prints a warning message to the log.
* If warnings are encountered during generation a special header will be show on the web UI log page.
* Raw consistency checks now use part of the warnings system to report results. This means that consistency check results
  will be printed in the warnings section, but they will not be repeated where the checks actually happen (a message
  indicating whether any problems were found is printed there instead).
* A set of (as of yet untested) Lua bindings that use my custom VM is included in this version. Also included is source
  code for the VM itself. Anyone crazy enough to want to should be able to cobble together a working Rubble engine using
  the two.
	* The differences between the normal and experimental APIs are slight, but it will still require lots of (very minor)
	  changes to certain kinds of scripts (mostly template implementations that use the script data registry).
* Added a new Lua method: `rubble.registry["key"]:listappend`, a convenience that is not terribly useful now, but will
  be really nice when the time comes to switch to the new Lua VM.

Addon and Standard Library Changes:

* Updated the `addon:Base` addon to DF 42.5
* Added a new template: `WARN` (and it's aliases `!WARN` and `#WARN`), prints a warning message to the log.
* Changed the temporary missing shared object ID warning messages to use the new warnings system.
	* BTW: These messages are going away soon (probably next version), so use them now if you need them!
	* For those who forgot: These messages help find shared object IDs that are missing the required object type prefixes.
* The `rubble.checkversion` Lua function now uses the warnings API for it's non-abort messages.
* Some templates expanded variables in parameters twice. Due to the way variable expansion works it should not have
  effected anything unless you were explicitly delaying expansion for some reason.
* Added the missing `SHEET` item type to the table in `addon:Libs/DFHack/Melt Item`.
* Went through most of the Lua scripts looking for things that will need to be changed when it is time to change to the
  new VM and marking them. Where possible I made changes now (as some things can be done in a way that works for both,
  for example using `data:listappend(item)` instead of `data[#data] = item`).
* Fixed the layer linked entity in the `addon:Base` addon not having any items.

Other Changes:

* Added the new templates and functions to the documentation.
* Added documentation for the experimental Lua API (the one that uses my custom VM).

7.3.1 (For DF 42.4)
------------------------------------------------------------------------------------------------------------------------

This version fixes some major problems with the `!SHARED_OBJECT` system, at the expense of breaking certain template
calls. The shared object system uses a shared name set for all registered objects, but even in vanilla objects of different
types sometimes use the same ID. I fixed this problem by prepending the object ID with the type of object it refers to
when you use a `!SHARED_OBJECT` variant. Sadly this breaks all existing calls to any template that deals with shared
objects. Updating is easy, you just need to add a type prefix to all shared object IDs you refer to in template calls.

I really hate to make breaking changes like this, but it was either break something or remove at least a few of the
new templates from last version (which would just hide the bug, leaving it to jump out and eat some unsuspecting modder
later).

Engine Changes:

* None, amazing.

Addon and Standard Library Changes:

* Fixed certain objects (tissue templates mostly) being clobbered by other objects with the same ID but a different type.
	* This fix effects every template that deals with shared objects, see [the base template documentation][RBT].
	* As a temporary measure to help fix addons that broke, any template that deals with shared objects will print a
	  warning to the log if you try to use an undefined ID (using an undefined ID is perfectly OK, but this will help
	  catch cases where the ID is OK, but you forgot the prefix).
* Fixed lots of calls to shared object templates to include the required object type prefixes.
* Added several new specialized `!SHARED_OBJECT` template variants, most of these have extra arguments beyond what normal
  variants have and combine functionality of several templates.
	* `!SHARED_REACTION`, automatically inserts the `REACTION` template.
	* `!SHARED_BUILDING_WORKSHOP`, automatically inserts the `BUILDING_WORKSHOP` template.
	* `!SHARED_BUILDING_FURNACE`, automatically inserts the `BUILDING_FURNACE` template.
	* `!SHARED_ENTITY`, automatically inserts the `!ENTITY_PLAYABLE` template.
	* `!SHARED_TRANSLATION`
	* `!SHARED_SYMBOL`
	* `!SHARED_WORD`
	* `!SHARED_COLOR`
	* `!SHARED_COLOR_PATTERN`
	* `!SHARED_SHAPE`
* Inserted the new `!SHARED_OBJECT` variants into the `addon:Base` addon and made `addon:Dev/SO Insert` automatically
  add them.
	* None of the new templates have been inserted into the other standard addons, mostly because it would be a lot of
	  work for only minor gain.

Other Changes:

* Updated the template documentation.

7.3.0 (For DF 42.4) Easier object editing!
------------------------------------------------------------------------------------------------------------------------

Abadrausar suggested some useful sounding templates, so here they are! I also added a few related templates that should
come in handy... It's nice to have vocal users, particularly when they have good ideas :)

The old templates I depreciated a few versions ago? All gone now. Any users of those templates should have updated by now.

I finally got around to moving the last of the object registration data out of the entities. The required templates have
been around for a long time, but somehow I never got around to inserting them.

I finally got fed up with manually having to update my init files when moving to a new DF version, so I threw together a
Rubble addon to do it for me. To make this addon possible (without too much work) I needed a few new script functions,
they will probably be nice for other things as well... While trying to use this addon I exposed a pair of major bugs in
the web UI, independent apply mode (and it's close sibling tileset apply mode) should work now.

Since I have a raw merging script function now, I went ahead and added a template for applying rule-based changes to
shared objects, It seemed to fit the theme for this version :) This template should obsolete a lot of uses of
`SHARED_OBJECT_REPLACE_TAG`, largely because you can change a whole swath of tags in one template call and enjoys more
precise tag targeting.

Engine Changes:

* Added a new Lua function: `rubble.rawmerge`, a simple interface to the rule-based raw merger.
* Added a new Lua function: `rubble.auxfile`, provides access to Rubble's internal auxiliary file buffers.
	* This is required for editing the init files (otherwise Rubble will clobber your changes when it flushes the buffers).
* Fixed a bug in the web UI that caused it to be impossible to choose addons in tileset apply or independent apply modes.
* Fixed a bug in the web UI that caused addons to be active in tileset apply or independent apply modes that should not
  have been active.
* The web UI now gives the proper return URL and return-to name to addon information pages reached from the independent
  apply and tileset apply mode addon selection pages.
* Fixed a bug in tileset apply mode that caused tileset raw files to not be loaded (they were accidentally getting their
  file tags clobbered).
* Fixed a bug in the raw merger. It was clobbering tags that it should not have touched, so I rewrote parts of it until
  it stopped.
	* I have no idea when this broke, it may have never worked from day 1, so it's a good thing it seems to have only
	  effected a fairly rare case.
* Fixed several merger bugs that may have been caused by my rewriting, but were in places I don't remember modifying...

Addon and Standard Library Changes:

* Removed the depreciated `SHARED_OBJECT` template and the other depreciated templates related to it.
* Added several new specialized `!SHARED_OBJECT` variants: `!SHARED_CREATURE_VARIATION`, `!SHARED_TISSUE_TEMPLATE`.
  `!SHARED_BODY_DETAIL_PLAN`, `!SHARED_INTERACTION`, and `!SHARED_BODY`.
	* The only object types missing specialized variants are those rarely changed by modders (except to add more), such
	  as color/pattern/shape definitions and language tokens. Also object types that already have other templates that
	  work just as well and/or are too complicated to work well with `!SHARED_OBJECT` (such as reactions, buildings, and
	  entities) do not have variants either.
* Inserted the new `!SHARED_OBJECT` variants into the `addon:Base` addon.
* `addon:Dev/SO Insert` now handles the new `!SHARED_OBJECT` variants.
* Added `!SHARED_OBJECT_CATEGORY`, allows you to set a "category" for a shared object. These categories may be queried
  with scripts for a list of all shared objects in the category or simply to see if a specific object belongs to a category.
	* This template is automatically inserted by the specialized `!SHARED_OBJECT` variants (both old and new).
	* This template works really well with custom categories and `@FOREACH_LIST`.
* Added `!SHARED_OBJECT_DUPLICATE`, duplicate an existing `!SHARED_OBJECT` with a new ID.
	* This template is a little tricky to use, see [the base template documentation][RBT].
* Added `SHARED_OBJECT_MERGE`, apply rule-based changes to a shared object.
* Added `@PARSE_TO`, parse template code and store the result to a configuration variable.
* Fixed an oversight/bug in some of the entity noble templates that caused the template bodies to not be parsed.
* All items in the `addon:Base` addon are now registered via addon hooks instead of being explicitly listed in the entity.
	* I also made similar changes to `addon:User/Saurian` and `addon:User/Gnome`.
* I added more rules for the syntax checker.
* The default merge rules for the tileset applier now includes rules for shape descriptor tiles.
* New addon: `Util/Milo's Settings` applies my favorite init and interface settings.
	* This addon is included as a template/example for others to use as a guide for making their own addons to do
	  similar things.
	* This addon is intended to be applied once to a new copy of DF (using independent apply mode), then ignored.
* Standard Lua function `rubble.targs` is now slightly more flexible, I also fixed it so variables in nested template
  bodies are not prematurely expanded in some cases (this was a particular problem with `@FOREACH`, as it was almost
  always nested in a call to `ECHO` it would have it's body's variables expanded too soon).

Other Changes:

* Updated the documentation to list the new templates and Lua functions.
* Moved documentation for the raw merger rule format into [Rubble Basics][RB3] from the top of the default tileset
  rules file.

7.2.0 (For DF 42.4) Now with raw syntax checking!
------------------------------------------------------------------------------------------------------------------------

I added a basic syntax check for raw output, it's not very useful yet, but once I add more rules it will help detect
invalid output and minor syntax errors without needing to start DF. It turned out to be very easy to turn part of the
tileset merge code into a rule-based syntax checker. Too bad I didn't think of it before...

Interestingly the syntax checker has already helped me find several bugs, even though it is nowhere near finished!

I finally got a (experimental, pre-release) 42.x DFHack, so I started testing my First Landing addon pack. First Landing
exercises parts of the Rubble engine that the default addons don't use, so of course I found a few bugs (surprisingly
few actually).

Engine Changes:

* The rubble core now validates the output raws using a basic raw syntax checker based on the same rule-based system
  tileset merging uses.
	* Any tags that do not match the rules will be printed to the log with a warning (but Rubble does not abort).
	* Rule files for the syntax checker use the same format as tileset rule files (and use the same `MergeRules` tag, but
	  without the `TileSet` tag).
	* Currently most of the rules are generic "stubs" that match anything, only material templates and inorganics have
	  full rules, the others only match the top level tags.
* `rubble7/rblutil/rparse.Tag` now has a `Line` field that should contain the source line on which the tag was found.
	* This field is also exposed to Lua scripts.
* Fixed a bug in the `rubble.configvar` Lua script function that prevented you from setting a configuration variable.
	* This also effected the `@SET` template.
* Fixed a bug that prevented automatic installation of creature graphic and tileset images.
* The `intn` function provided by the random number generator returned by the `rubble.random()` Lua function now returns
  a number between `0` and `n` *inclusive*, previously it returned a value between `0` and `n-1` (which is not what it
  was supposed to do).
* Templates that "miss" their parse stage (because they were nested in a later stage template) now trigger an error.

Addon and Standard Library Changes:

* Updated the `addon:Base` addon to DF 42.4.
* Added a few missing material color tags to the default tileset rules.
* Added a 24x24 version of my personal TWBT tileset (`addon:Tilesets/MLC/TWBT - 24px`)
	* The old tileset is renamed to `addon:Tilesets/MLC/TWBT - 16px`
* Fixed a pair of bugs in `addon:Libs/Random`, one related to setting the seed and the other causing non-deterministic
  behavior in `select` and `select_weighted`.
* Fixed a major bug in `addon:Libs/Castes/Transform`.
* Fixed a missing close brace (']') in one of the `addon:User/Warcrafter` templates.

Other Changes:

* None

7.1.2 (For DF 42.3) df_version++
------------------------------------------------------------------------------------------------------------------------

The big change here is the new DF version, the second biggest change was incrementing the version number :P

Seriously, if it wasn't for the new DF I would have waited.

Engine Changes:

* None

Addon and Standard Library Changes:

* Updated the `addon:Base` addon to DF 42.3.
* Updated the DFHack script in `addon:User/Display Case` with a newer version.

Other Changes:

* None

7.1.1 (For DF 42.2)
------------------------------------------------------------------------------------------------------------------------

This version introduces one minor change, something I should have done a long time ago.

The original `SHARED_OBJECT` template was put in the main parse stage to prevent me from having to manually parse the
body, but this effectively disallowed the use of pre-parse templates in a shared object body. Most other examples of
excessive laziness have long been fixed, but `SHARED_OBJECT` continued to fester. Anyway, I have now moved it to the
pre-parse stage (where it should have been all along) and inserted the necessary code to parse the object body in all
stages. The old templates will continue to work for a bit yet, but after a few versions they will go away for good.

If you use the old templates I do not want to hear any bug reports! If a template defined in `addon:Libs/Base` is not
listed in the docs you should not be using it.

Engine Changes:

* None

Addon and Standard Library Changes:

* Added `#SHARED_OBJECT_EXISTS`, post-parse alias for `SHARED_OBJECT_EXISTS`.
* `SHARED_OBJECT` is depreciated, use `!SHARED_OBJECT`. This also applies to the specialized variants as well.
	* The old templates will continue to work for a few versions yet, but they will write a warning to the log when used.
	* This also covers `ITEM_CLASS` (it needs to be in the same parse stage as it's preceding call to `!SHARED_ITEM`).
	* This change will help prevent evaluation order issues and allows you to use pre-parse templates inside of shared
	  objects.
* Replaced every call to the depreciated templates (that I could find) in the standard addons.

Other Changes:

* Removed mention of the depreciated templates from the docs (replaced with their replacements).

7.1.0 (For DF 42.2) The "I thought I tested that! Guess not..." version.
------------------------------------------------------------------------------------------------------------------------

This version fixes a *bunch* of bugs.

Some of the old addons have been slightly reworked with the new reaction tags 42.x adds, they all work more or less the
same, but nicer in minor ways.

Contrary to how it may look I really did test 7.0.0! I just didn't do nearly enough of it and most of it was done *before*
I had finished making changes. I really need to write some automated tests for Rubble to prevent this kind of thing in
the future.

Anyway, I did bunches of generation tests for this version, but no actual in-game testing. I refuse to subject myself to
the torture that is playing DF without DFHack...

So much stuff is new that bugs are to be expected, but hopefully most of them are squashed now.

Stuff I Forgot to Mention Last Version (AKA Bugs That Retroactively Became Features):

* Variable names in template parameters are no longer explicitly expanded. Since parameter expansion takes place before
  variable expansion in a template body most cases will work as before, but script templates must expand variables in
  parameters explicitly if they want that.

Engine Changes:

* The web UI now works on some browsers it didn't work on before. Curse my outdated JavaScript books, and praise palu
  who not only found the problem, but fixed it as well!
	* I have no idea which browsers worked and which did not, I don't exactly have a broad range of them to test it on...
* Fixed a nasty bug in the variable/parameter expansion code. If you had a single percent sign with no following braces
  and the variable name was blank then you would end up with it being replaced with a percent sign and an open brace,
  which would lead to an "invalid token" crash if it was ever run through the Rubble parser.
	* Any percent sign followed by a non-alphanumeric character other than `{` inside of a template call would trigger
	  this bug.
	* This bug was exposed by Dirst, but was not recognized as a bug immediately due to faulty (outdated) documentation
	  on my part, sorry.
	* See [Rubble Basics][RB2] for more information.
* Removed the "Prerelease A" version tag that I had forgotten to remove last time.
* Fixed the select all button in the web UI so it *really* only works for the addons that are currently shown.
* Raw files were not being prefixed with a file ID, so literally nothing worked.
* Script errors that trigger during a template call are now annotated with the location of the template call (in addition
  to the template name) if possible.
* Added a (possibly horribly broken, absolutely untested) way to convert strings to byte slices and vice versa for the
  JavaScript API. The JS AP is still very much use at your own risk with no documentation.

Addon and Standard Library Changes:

* Two templates in `addon:Base` were missing the closing braces.
	* Thank you palu for reporting this bug *and* providing a fixed version of the base for others to use until I
	  could push out a new version!
* Removed calls to the removed template `!RESEARCH` (and a related template) that I had missed earlier.
	* Thank you Abadrausar for reporting this bug!
* The `SHARED_OBJECT_KILL_TAG` template was broken (it was caused by a bug in `rubble.rparse.format`, stupid `1` based
  indexing).
	* Thank you palu for reporting this bug!
* `addon:Dev/SO Insert` did not work. I also took the opportunity to fix some minor formatting issues in the output.
	* Thank you Abadrausar for reporting this bug!
* Added two new templates to `addon:Libs/Base`: `REACTION_CATEGORY` and `!REACTION_NEW_CATEGORY` allow you to specify
  categories for reactions in a simple way.
* The `addon:User/Glass Forge` addon now uses the new reaction category templates.
* Added `[MAX_MULTIPLIER:1]` tags to every reaction I could find that used bones (effects `addon:User/Bonecarver`
  `addon:User/Bone Flux`, `addon:User/Archery`, `addon:User/Adv Reactions`, `addon:User/Warcrafter`, and
  `addon:User/Warcrafter/Adventure`).
* `addon:User/Pottery/Glaze` was accidentally being tagged as requiring DFHack (it has some optional DFHack content).
* The two child addons of `addon:User/Adv Reactions` were renamed to remove the `DFHack` part.
* Fixed a bug in `@BUILD_KEY` that caused it to loop forever if the requested key was in use.
	* I remember fixing this bug before, I must have changed the wrong copy, that's embarrassing.
* Fixed some bad calls to `@SCRIPT` in `addon:User/Bonecarver`.
* Ported two Rhino scripts I missed in `addon:User/Metallurgy/Smelter` and `addon:User/Metallurgy/Bloomery` to Lua.
* The standard library Lua function `rubble.targs` now does some more stuff, namely you can have it return a table and
  it expands variables in it's arguments unless you tell it not to.
* Fixed a bunch of possible bugs related to variables failing to expand in script template parameters.
	* So long as you never tried to use a variable as a script template parameter this would never be a problem, but if
	  you ever did it would fail to work with some templates and succeed with others.
	* `@SCRIPT` explicitly does *not* expand variables in it's parameters, and `@SET` allows you to specify if it should.
* Added `addon:User/Zap Aquifers`, this addon does the same thing as the Rubble 6 addon of the same name.

Other Changes:

* Most of the tutorials were way out of date, so I rewrote most of them at least a little.
* I also scanned the other docs looking for incorrect stuff and/or stuff that needed clarification, this resulted in lots
  of little changes all over the place.

7.0.0 (For DF 42.1) 
------------------------------------------------------------------------------------------------------------------------

This version has "only" one major change as far as most users are concerned, the new web UI. All of the other changes
focus on making Rubble easier to work with for modders and making it work with the new DF version.

The new web UI looks much nicer, and is much more dynamic. In general I ditched the old HTML forms for JavaScript forms
that communicate changes to the server immediately. This allows the web UI to "remember" what you were doing better. For
example, if you click on an addon name when selecting addons during generation it will open the addon information page
just like it used to, but unlike before it will not forget what addons you had selected when you return, plus the addon
information page will change slightly to reflect how you got to it ("Back to Addon List" becomes "Back to Generation",
and takes you to the proper page).

I removed regeneration mode from the web UI. If you want to do something so inherently dangerous as regenerating a
region's raws you will have to use the CLI interface, at least that way when you machine-gun your foot you will have
had plenty of time to reconsider :P

The driving force behind Rubble 7 is long-term compatibility. Too many parts of Rubble depended on things that I did not
want to support forever, so I made a major effort to replace any fragile bits with more robust replacements that will be
easier to maintain long term.

Support for scripting has been removed from the Rubble core, replaced with a simple system to register "script runners"
from external libraries. This allows me to support more than one language at a time. For now I am adding support for Lua
and JavaScript, but I may add others later. Lua support is more-or-less complete, but JS support is very rough. Stick to
Lua for now, a proper JS API is on it's way.

addon.meta files are no longer scripts. They are now simple JSON files, with a redesigned structure to better support
additions and changes without breaking compatibility with older/newer Rubble versions. Obviously addon.meta files from
the Rubble 6 version series will not work, but with luck this will be the last such compatibility break.

Trying to support multiple scripting languages overstressed the old file tagger, so I wrote a new (much better) one.
Since I was reworking the addon.meta files as well I went ahead and rewrote the whole addon loader, the new one has some
very nice new capabilities. Chief among the abilities is the ability to specify an addon name that has nothing to do with
the addon's location in the addons directory. This allows addon authors to set a canonical name for their addon without
needing to ensure the addon is installed exactly as directed (which is difficult, since people *never* follow directions
exactly).

The new tagging system makes it easy to add new file types, so I went ahead and added automatic installation support
for a bunch of DF files that Rubble previously handled via templates. Creature graphics (both raws and images), tile
sheets, DFHack Ruby command scripts, and auxiliary text files are all now automatically detected and installed.

Engine Changes:

* Rewrote almost everything to some degree, in particular:
	* The scripting subsystem now loads script "runners" from external libraries without needing to know anything about
	  the language.
	* New addon loader that allows much more user control and has a more flexible file tagging system.
		* This required changing the recognized file extensions, many of the rarer file types will need to be renamed.
	* addon.meta files are JSON encoded and are in general **much** more flexible.
		* Most of the addon.meta files were converted automatically with a simple utility I hacked together for the purpose,
		  their formatting leaves a great deal to be desired.
	* Moved a lot of stuff out of the Rubble core and into child packages, the Rubble core is now almost exclusively
	  "glue" code that ties all the little bits together and provides a coherent API for the interfaces.
* The web UI has received a significant face-lift and functionality upgrade, here are some of the highlights:
	* Regeneration mode is gone, use the CLI if you need to regen a region.
	* There is a new type of special link (addon file links), and workshop image links now take an addon/file pair instead
	  of an AXIS path. See [Rubble Basics][RB1]
	* Finally killed the back button! It is no longer possible to get yourself into trouble by going back to a page that
	  you should not be able to get to right then.
	* You can filter out addons with specific tags from the addon list when selecting addons to generate, so far only the
	  "Dev" and "DFHack" tags are supported (addons tagged "Dev" are hidden by default).
	* New select all/none buttons (that respect the current filter settings!) on the select addons page.
	* The addon information page shows a lot more information about the addon.
	* Addon lists are now consistently in load order, instead of sometimes being alphabetical.
* The standard addons have been significantly reorganized, many of them have new names.
	* In particular addons that use DFHack no longer have "DFHack" as part of the addon name, instead they have it as an
	  addon tag.
* Rhino (the Rubble 6 scripting language) is no longer supported.
* All interfaces now support Lua for scripting, see the [Lua API documentation][RLUA].
	* The Lua API is more or less feature complete. There are still a few minor things that it does not allow that were
	  possible with the old Rhino API, but all the stuff I actually used is there.
* All interfaces now support JavaScript for scripting, see the [JS API documentation][RJS].
	* The JS API is **far** from complete, and what *is* there is *temporary* and *will change*.
* Base 64 encoded zip files are no longer supported. They were originally to be used for patches and such like, but I
  soon discovered that it was easier to just package a new version.
* Template parameters that are quoted now have the quotes stripped and any escape sequences expanded.
	* This makes it easier to have leading and trailing whitespace or for shortening certain sequences.
	* The quoting rules and escape sequences supported are the same as [Go](golang.org) uses.
	* Unquoted parameters are unaffected.

Addon and Standard Library Changes:

* All standard templates in `addon:Libs/Base` have been rewritten in Lua.
* Several templates have been removed from `addon:Libs/Base`:
	* `@IF_CODE`, `!DECOMPRESS`, `!PATCH`
* Several templates that are no longer needed have been removed from `addon:Libs/Base`:
	* `@INSTALL_GRAPHICS_FILE`, `@INSTALL_IMAGES_AS_GRAPHICS`, `@GRAPHICS_FILE`, `@TEXT_FILE`
* Added some new standard templates to `addon:Libs/Base`:
	* `@STR_LOWER`, `@STR_UPPER`, `@STR_TITLE`, `@STR_REPLACE`, `@STR_TO_ID`, `@GENERATE_ID`, `#` (alias for `#ECHO`),
	  `!` (alias for `!ECHO`), `@STORE_LIST`, `@READ_LIST`, `@FOREACH_LIST`, `@STORE_TABLE`, `@READ_TABLE`, `@FOREACH_TABLE`,
	  `@FOREACH`, `@VOID`, `@ECHO`, `@` (alias for `@ECHO`), `@GENERATE_COUNT`
* Some templates in `addon:Libs/Base` have slight changes to their usage (these changes are unlikely to effect anything):
	* `@SCRIPT`, `#USES_BUILDINGS`, `#USES_REACTIONS`, `#USES_TECH`
* The entity playability templates (`!ENTITY_PLAYABLE` and such like) have had arguments related to outsider playability
  removed.
* The `SHARED_OBJECT_ADD` template in `addon:Libs/Base` had the undocumented property that you could not add the same
  text twice. This property no longer holds (I'm not sure why I even had it that way in the first place).
* Removed the old `Register` function from `addon:Libs/DFHack/Powered`, the old `RegisterAdv` function is now named `Register`.
	* The old function was only kept around for compatibility reasons anyway, so it is high time to force an upgrade.
* Some of the old addons from Rubble 6 are not available yet. All of the core library addons are ported, but some of the
  `User` addons are not yet ported.
	* Notably none of the powered workshops have been ported yet.
* Nothing DFHack related has been tested (obviously), since there is no 42.x DFHack available yet.
	* Most things were known working in 40.24, so baring major changes they should still work now.
* I refuse to play DF without DFHack, so nothing has been tested much at all, expect bugs.

Other Changes:

* The general documentation has been fully updated.
* Some of the old tutorials are gone (temporarily, I plan to rewrite them to fit the new APIs).

[RB]: /doc/Rubble%20Basics.md
[RB1]: /doc/Rubble%20Basics.md#WebUILink
[RB2]: /doc/Rubble%20Basics.md#ExpandIssues
[RB3]: /doc/Rubble%20Basics.md#MergeRules
[RB4]: /doc/Rubble%20Basics.md#Tests
[RB5]: /doc/Rubble%20Basics.md#WebLoad

[RLUA]: /doc/Rubble%20Scripting%20Lua.md
[RJS]: /doc/Rubble%20Scripting%20JavaScript.md

[RBT]: /doc/Rubble%20Base%20Templates.md
