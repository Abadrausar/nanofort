
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

8.0.0 (For DF 42.6) Addon loading rewrite!
------------------------------------------------------------------------------------------------------------------------

This version brings a much needed update to AXIS VFS and a powerful new addon loader to go with it. Rubble 7 was all about
long term maintainability, this version is all about addons, loading them, updating them, and in general making the addon
pack system into the powerful tool it should be.

The new addon loader is much more flexible. Addons are loaded once, and reused over and over as needed, so Rubble should
run faster (no more annoying pauses while the web UI waits for the loader!). Addon packs may now provide a "pack.meta"
file with information about the addon pack, in particular an auto-update URL.

I ran into some file order problems caused by the convention that mod ID elements in file names should be uppercase.
This convention was causing certain files to sort earlier than they should. As I rather like this convention, I fixed
the problem by changing the way Rubble sorts files and addons. Any items sorted by name are now sorted in a case
insensitive manner (basically everything is temporarily converted to lowercase while sorting).

Most of the changes that are not directly related to the new AXIS and loader are the result of suggestions by my loyal
crew of users. More suggestions are always welcome! There is nothing more valuable than vocal users :)

Engine Changes:

* New AXIS VFS!
	* This required *every single file path* to be changed (as AXIS 2 uses a different path syntax from AXIS 1).
	* The new AXIS is much more flexible, and is far easier to extend.
	* The AXIS script API remains functionally unchanged.
* New addon loader!
	* Addons are now loaded once, and stored together with certain other bits of (mostly) immutable data. This makes State
	  creation much faster, eliminating the the annoying pause when going to the main menu in the web UI.
	* .webload files and the `Downloads` addon.meta key are gone. They made implementing certain parts of the new loader
	  far too hard. Replacements based on the "pack.meta" system are provided.
* Addon packs now are allowed a special meta file that covers all addons in the pack. This "pack.meta" file is a JSON
  file, just like addon.meta, but with a different set of supported keys.
	* Addon packs may request automatic updates via a pack.meta key.
		* It is possible to forcibly prevent addon packs from updating by listing them in "addons/update_blacklist.txt".
	* Take the time to read [the documentation on pack.meta][AddonPacks], it (pack.meta) is critical to many new features.
* Changed the way file output is handled:
	* Any addon pack may now add more tag-based "writers" via a pack.meta key.
	* All the hardcoded write actions have been moved to "addons/global.meta".
		* This file is global, and not part of any addon, basically it is a pack.meta file for the standard addons, but
		  with some special handling.
	* File extension compaction/transformation is now more flexible. Nothing uses this flexibility, but it is there if
	  needed.
* The file tagger is now more flexible.
	* There are no hardcoded file tagging rules anymore, the default rules are loaded from "addons/global.meta".
	* Addon packs may also specify their own tagging rules via a pack.meta key, but these rules then only apply to addons
	  that belong to that pack.
* Dropped JavaScript support. It has never been feature complete, and due to limitations in the API of the VM I was
  using it probably never will be.
* Addon information exposed to scripts is now read only (it has to be this way to prevent scripts from clobbering the
  stored addons).
* Activation state is no longer a property of the addon, it is now stored separately.
	* Lua scripts can use `rubble.addonactive` to query addon activation state.
* Other changes all over the place as required by the new loader or the new AXIS or both.
* Added a new first part extension (".speech") and file tag ("SpeechText") to support speech files. Correctly tagged files
  will be written to "df/data/speech" after generation.
* Added a new "file bank" system:
	* Addons tagged "FileBank" will not be loaded, instead they will have their path added to a list of file banks.
		* Child directories of these addons will not be considered by the loader!
		* The name field of the addon.meta file still works (except it sets the bank name instead of the addon name).
	* You can request that a file bank be written to a specific AXIS path, if you do so all files and directories will
	  be recursively copied to the path you specified.
	* Since Rubble does not actually load files from a file bank they will not be available for editing via the normal
	  methods!
	* See [Rubble Basics][RB2] for more information.
* Added new script function: `rubble.copyfilebank`, requests a file bank copy.
* Files and addons now sort in strict alphabetical order (case is ignored). This prevents certain odd sorting issues.
	* For items that have the same name with different case (which is a very bad idea BTW) order is undefined.

Addon and Standard Library Changes:

* Tracked down and fixed every AXIS path I could find, luckily there aren't very many. (I hope I got them all...)
* Shared object templates added by `addon:Dev/SO Insert` now use the class `ADDON_HOOK_PLABLE` instead of `NULL` (if
  applicable). Items with rarity are given the rarity `COMMON`.
* `rubble.expandargs` will now convert any nil parameters into empty strings (since this is more likely to be what the
  user wants).
* Added new standard template: `SHARED_OBJECT_DELETE`, makes a good effort to eradicate all mention of a specific shared
  object.
* Added new standard template: `@COPY_FILE_BANK`, requests a file bank copy.
* Added new library addon: `addon:Libs/Edit Init` provides templates for adding new embark profiles or worldgen parameters.

Other Changes:

* Lots. This version broke documentation all over the place (mostly in minor ways, but some major).
* [Rubble Basics][RB] now has a table of contents!

[RB]: /doc/Rubble%20Basics.md
[RB1]: /doc/Rubble%20Basics.md#AddonPacks
[RB2]: /doc/Rubble%20Basics.md#FileBank
