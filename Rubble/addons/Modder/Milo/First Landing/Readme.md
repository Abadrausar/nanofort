
Overview:
====================================================================================================

This file contains a change log and installation instructions for the First Landing addon pack.
Actual documentation for First Landing is included in the `addon.md` file in the `Base` directory.

It is possible to open that file directly in a text editor (which allows you to read the documentation
in markdown format), but a much better idea is to install this addon pack (using the directions below),
then start the Rubble web UI (following the directions in the Rubble documentation). Once the Rubble
web UI is running simply visit the addon documentation page, then select the entry for `First Landing/Base`.
This will allow you to read the documentation complete with workshop preview images and fancy formatting.

First Landing is a Rubble addon pack, as such it requires Rubble to install and use. Rubble is a
powerful tool for making and managing Dwarf Fortress mods. First Landing uses Rubble to, among other
things, generate part of it's content, allow user configuration (you can change how parts of the mod
act), and allow installing extra content with minimal fuss.


Installation:
====================================================================================================

1. Download and install the latest version of Rubble (see the Rubble readme for installation instructions).
2. Copy this archive into the Rubble addons directory (named "addons" by default) **without extracting it**.
3. Run Rubble and generate with the addons in this pack.

You **must** have DFHack installed to play this mod!

Installing Rubble is easy, in simple terms all you need to do is extract Rubble into a subdirectory
of your Dwarf Fortress directory. Running Rubble is also easy, simply run `rubble_web` and it should
start your web browser and display a menu (which should be fairly self explanatory). On Windows your
default web browser should be automatically used, Linux and OSX users will have a few more things
that need to be done, see the Rubble readme.

For more detailed installation instructions see the Rubble readme.

Documentation for this mod is best viewed with the Rubble web UI.


Bugs and Other Issues:
====================================================================================================

No bugs are known.

If you find a bug please report it in the First Landing forum thread immediately!

I can only check the forums once a week or so, so it will probably take me a little while to respond
to bug reports, but I will fix any reported problems as soon as I can.

Suggestions are also welcome, but keep in mind that just because it was suggested does not mean I will
implement it! Your suggestion is much more likely to be well received if it fits the First Landing
lore.


Changes:
====================================================================================================
1.5 (DF 42.6, Rubble 7.5.0)
----------------------------------------------------------------------------------------------------

I got a DFHack for 42.6, so I celebrated by playing First Landing for a while. Unsurprisingly I found a few bugs, but
only minor ones this time. Yay!

Much to my dismay I have discovered that the new "details" menu does not work with many reactions. I may be able to work
around this in time, but for now you can't do half the really cool things that this menu allows. Don't expect a solution
to this issue just right away.

* Fixed an incorrectly colored tile in the glass blower.
* Eggs are edible again, oops.
* Somehow I forgot to provide a way to make splints and crutches out of anything but plastic, well now you can make them
  out of wood and bone too.
1.4 (DF 42.6, Rubble 7.5.0)
----------------------------------------------------------------------------------------------------

This version is a simple compatibility update for Rubble 7.5. The only changes are minor fixes to all the script modules,
nothing important. If you don't see a pressing need to update to Rubble 7.5.0 then there is no need to update to FL 1.4
either.

* Fixed script modules to use the Lua 5.3 module syntax instead of the old 5.1 syntax.
* Fixed a few cases where zero-based indexes where used for things that the new VM uses 1 based indexes for.

1.3 (DF 42.6, Rubble 7.4.2)
----------------------------------------------------------------------------------------------------

I totally redid colors to use a broader palette. Instead of 7 colors with two shades each plus white and black the new
palette has more like 4 colors with two shades each plus 8 other colors (including white and black). This gives more
display color options, always a nice thing to have.

To go with the new display colors I redid all the color definitions. There are now far fewer colors defined, but things
like skin and hair color are "descriptive" instead of absolute (for example skin may be "lightly tanned" or "pale").

The new colors are a big change, I tried to test as much stuff as possible, but some things probably still look "funny".
As time goes by I hope to get everything ironed out and looking the way it should again...

Random trees have had their color tables rebalanced. Wood materials tend to be more "subdued" than before (although
brown is only one of many possible "subdued" colors, so they are still far more varied than vanilla), but leaves are
still a riot of bright colors. Since I was working on trees anyway I went ahead and implemented different size categories,
so trees in different biomes are more varied now.

Incidentally (and purely by accident) I discovered a safe way to disable forgotten beast generation, so the last possible
hardcoded random creatures are gone! Finally. Sadly this comes at the cost of some ignorable error log spam whenever you
generate a new world.

* Replaced all vanilla colors definitions with new ones plus an all new color palette.
	* This has the effect of clobbering any custom color scheme you may have installed, but don't worry it will only
	  effect FL worlds (the new palette is installed when you load the world and removed when the world is unloaded).
	* The new colors effect pretty much everything to one degree or another.
	* The actual color palette used comes from the DF From Scratch Redux (DFFS3) project.
* Simplified creature materials slightly, all organs now share a single "<creature name> organ" material and eggs are
  solid blobs of eggshell (not that there is any way to tell what they are made of in game).
	* This has the side effect of disabling random creature generation completely, huzzah!
* Random trees are more likely to have wood in "subdued" colors (the darker 8 colors). Bright colored woods are still
  possible, but less likely. The same is true of random underground mushrooms, except bright colors are more common there
  than in aboveground trees.
* Random trees can now be one of several "archetypes" which determine how large the tree is. Which archetype is chosen
  depends on the biome chosen for that tree (for example desert trees will always be "stunted" while wetlands will mostly
  have "normal" and "large" trees).
* Added a new name generator for creatures that is designed to create more memorable names.
	* This generator is optional, currently off by default as it does not have enough word list variety yet.
* Made some minor tweaks to the soil materials.

1.2 (DF 42.5, Rubble 7.4.0)
----------------------------------------------------------------------------------------------------

* Fixed a stupid script bug that trashed all pottery reactions (except glazing).
* Added more biome picks to the random tree addon, you should now see trees in shrubland areas.
* Random wood is now described as being colors other than "brown" (you can see this with certain DFHack interface
  improvement scripts/plugins).

1.1 (DF 42.5, Rubble 7.4.0)
----------------------------------------------------------------------------------------------------

There are two new features this version: "Hive" tending and two more (rare) plant types.

The "hive" system consists of a few semi-automated workshops, each workshop queues jobs automatically at intervals. So
long as you have labor to do the jobs you will have a slow, but steady, trickle of whatever the "hive" is supposed to
produce. This was originally intended to be a beekeeping replacement, but I decided that terrestrial bees didn't make
sense in First Landing, so I added some other stuff instead. Only one of the possible industries is a hiveable insect
simulation, the other two are food/water production machines intended as early game and emergency supply sources.

The rare plant types allow you to make storage items directly, but these plant types are much rarer than the other types.

In addition to the new features I made some internal changes to bring First Landing up to date with recent additions to
the Rubble template language, plus a few other little tweaks here and there.

* Inserted various templates, mostly new `!SHARED_OBJECT` variants from Rubble 7.3.0 and 7.3.1.
* Striped out various files from Modest Mod, replaced with vanilla versions (except the material and tissue templates).
* Changed the way the IDs of random creatures and plants are generated so that they provide useful information when you
  run DFHack's "prospect" command.
	* The IDs now contain the display name as well as any usage or description codes the object has.
	* This feature was suggested by Abadrausar.
* Reworked all the reaction and buildings to use the new `addon:Libs/Industry` templates.
	* Many internal templates have been moved to the pre stage to accommodate this, so those interested in making FL-based
	  addons will want to take another look at things.
* Added basic hive keeping. This does not use the vanilla hive system, but works in a similar way.
	* Currently there are several possible products:
		* Resin beetles produce globs of resin that may be processed into plastic.
		* Algae tanks produce mostly tasteless "algae cake", a good emergency food source (eating unprocessed algae
		  gives a minor bad thought).
		* Moisture condensers produce barrels of drinking water (you supply the barrels of course).
	* I hope I have the production rates more-or-less balanced, If anyone thinks they should be slower/faster let me know.
* There are two new types of (rare) random plants, process these plants with the "process special plant" reaction:
	* Gourd analogs may be made into jugs, but they need to be scraped and dried first (after processing they will need
	  to sit for a while before they may be used).
	* Some plants have particularly strong leaves which may be used as bags after minor processing.
* Hardened drill bits (for the drilling rig) must now be made from weapons grade metal instead of any metal.
* Shell items now cost twice as much as bone items.

1.0 (DF 42.4, Rubble 7.2.0)
----------------------------------------------------------------------------------------------------

The big change here is the move to the new DF and Rubble.

Rubble 7 brings better script performance, as well as a nicer web UI, and other much needed improvements.

Switching to DF 42.x allows some very nice improvements to various industries, namely the bone carver now
works more like the other workshops, and most of the workshops use a hierarchical menu to organize all those
reactions.

Not much new (aside from stuff required by the new DF), but I did add random farmable plants, and made some
minor changes to other farm related stuff.

You should be able to make taverns, libraries, and temples just like vanilla. Bookmaking has its own workshop,
toys, goblets, mugs, bookcases, etc can all be made at the various existing production workshops. Currently
musical instruments are import-only (but as soon as I find a RL instruments mod that I can use...).

* Updated to DF 42.x and Rubble 7.
* Added `addon:First Landing/Planets/Flora/Random Crops`, random farmable plants.
* Removed `First Landing/Planets/Flora/Default Crops`.
	* The "process special plant" reaction used for barrel gourds and quarasote is still present, but not currently used
	  for anything. It will be used later.
* Removed everything related to bee keeping as there are no bees, maybe later.
	* I have plans for a really cool replacement for vanilla bee keeping, so make that "definitely later"
* Plant pastes now use different materials for pressed and unpressed states. This prevents your cooks from cooking all
  your unpressed paste.
* There are now reactions for milling specific types of material (flour, sugar, dye, etc).
* Removed a bunch of stuff related to fruit. The only source of fruit is trees, and the only thing you can do with it is
  brew it or eat it, so there is no need to have, for example, a "press liquid from fruit" reaction.
* Obsidian, quartzite, and sandstone boulders can be ground into various kinds of sand at the millstone and auto-mill.
	* If you want to grind the rock to it's normal material for use at the decorator use blocks.
* Certain of the powered logistics buildings work now.
* Added a "Bookbinder's Workshop", a simple place to make paper and books.
* You can now make a few more items at the various production workshops, for example bookcases, cups and toys.
* All reactions that use bones now produce the number of items they say they will, extra bones are left to use for something
  else if you wish (instead of being used to make more items).
* Most workshops now have proper category menus instead of the old category headers.
* Fixed smelting ores that require the arc smelter.
* Approximately 1 in 50 random vertebrates will have been domesticated by previous bands of settlers, and some of the
  non-egglaying vertebrates may be milked.
