[Addon Change log](/addonfile?addon=Dirst___Accessibility&file=addon_changes.md)
![Accessibility Haiku](http://conqueringnetherlands.files.wordpress.com/2012/12/image31.jpg)
## So you think that you need some help, want to do a sugestion, inform about a bug or give thanks to modders or developpers, but not sure where to go.
1. If you want to patronize Toady One and the developpment of Dwarf Fortress go to: [Support Dwarf Fortress Vanilla](http://www.bay12games.com/support.html)
2. For questions about the rubble framework go to: [Rubble framework](http://www.bay12forums.com/smf/index.php?topic=154304.0)
3. Anything about HDFPS: Humble Dwarf Fortress Publishing System is there [HDFPS](http://www.bay12forums.com/smf/index.php?topic=157300.0)
4. If your question is addressed to the actual original maintainer of this mod then you have [Accessibility](http://www.bay12forums.com/smf/index.php?topic=157631.msg6941343#msg6941343)

# A set of simple commands for the DFHack console to aid in navigating the map. Comments and suggestions are welcome. The current version includes three DFHack scripts.

Installation is simple. Unzip Accessibility.zip onto your Dwarf Fortress folder.
# bookmark.lua
This script defines, updates and deletes persistent bookmarks on the map.

bookmark help
Prints a short summary of the bookmark comamands.

bookmark list
Lists all current bookmarks, if any.

bookmark drop foo
Deletes a bookmark named foo.

bookmark clear all
Deletes ALL bookmarks. Note the space between "clear" and "all" as a safety feature.

bookmark foo
Creates or updates the bookmark named foo at the current cursor location. If the cursor is not present, the bookmark will be set as the center of the on-screen map.
# goto.lua
This script will move the map (and cursor if it is present) to a specific location. The goto script has two modes: bookmark and coordinates.

goto help
Prints a short summary of the goto commands.

goto foo
Moves the map to the bookmark named foo, if it exists.

goto # # #
Moves the map to the specified X, Y, Z coordinates.
# whereami.lua
This script reports your current location as X, Y, Z coordinates. whereami will also report if the current location is recorded as a bookmark.

Known issues:
Bookmarks saved with version 1.00 (if you somehow acquired it) have incorrect Z coordinates.