# Dwarf Fortress Vanilla
Contains Rubblized versions of the complete Vanilla raw files that are used as an include or import base library able of fine grained imports.
EACH time that a new TOTAL mod is included, for the first time, this lib should be adapted to take into account the import needs of the new TOTAL mod. 


# The following changes have been made from vanilla by Milo:
* Some of the files have been re-indented (tag order is not effected).
* The `DWARF` creature has templates from the `addon:Libs/Castes` addon inserted.
* Specialized `!SHARED_OBJECT` templates have been inserted where appropriate.
* `REACTION` and `BUILDING_WORKSHOP` templates have been inserted where appropriate.
* `@BUILD_KEY` templates have been inserted where appropriate.
* `ADDON_HOOKS` have been added to all entities.
* `!ENTITY_PLAYABLE` templates have been added to all entities.
* The dwarven outpost liaison has been modded to work as a trade rep if the dwarves are unplayable.
* Dwarves have the same progress triggers as humans.
* Dwarves have the alchemist labour permitted in their entity.
* All file IDs have been stripped.
* The files from "objects/text" have been added (with appropriate double internal rubble file extension).
* Added `#ENTITY_NOBLES` to all entities.

# The following changes have been made from vanilla by Abadrausar:
* For each kind of DF object a directory has been created to group them, so the objects of type BODY are inside a directory named BODY, those of type CREATURE inside a directory named BODY CREATURE and so on.
* When a mod needs to import some vanilla objects of a type but not the rest, those objects are moved into a nested dir inside the base dir of the type. 
* Following the anterior principle we can find almost all vanilla creatures inside the base CREATURE directory, but then we remark that the wagon creature has been moved into the  CREATURE/equipment directory, the reason to do that is that even if some total mod does not use any vanilla creature you will always need a wagon to start the game ;)
* Until now I have only found a mod (Rise of the Mushroom Kingdom by IndigoFenix) that modifies the wagon for the rest of the total mods the vanilla creature wagon is the only vanilla creature that they need in order to have a working game
* Then some others creatures are separated to answer the particular needs of the Haiku, First Landing and in general any kind of TOTAL mods.

# This mod is Toady One Vanilla, if even a little token or object lacks it is a bug! Please report it.
Remark however that the objects can be in unusual files or places for the reasons explained.
[Addon Change log](/addonfile?addon=Toady_One___Vanilla&file=addon_changes.md)
## So you think that you need some help, want to do a suggestion, inform about a bug or give thanks to modders or developers, but not sure where to go.
0. If your question is addressed to the actual original maintainer of this mod then you have [Dwarf Fortress Vanilla](http://www.bay12games.com/dwarves/dev.html)
1. If you want to patronize Toady One and the development of Dwarf Fortress go to: [Support Dwarf Fortress Vanilla](http://www.bay12games.com/support.html)
2. For questions about the rubble framework go to: [Rubble framework](http://www.bay12forums.com/smf/index.php?topic=154304.0)
3. Anything about HDFPS: Humble Dwarf Fortress Publishing System is there [HDFPS](http://www.bay12forums.com/smf/index.php?topic=157300.0)
![Dwarf Fortress Vanilla Haiku](https://github.com/Abadrausar/nanofort/blob/master/Rubble/addons/Modder/ToadyOne/Vanilla/Base/Toady_One___Vanilla.jpg)
