Contains versions of the Vanilla raw files that are used in Rubble as include or import base library able of fine grained imports to serve the needs of all other DF mods.
EACH time that a new TOTAL mod is included, for the first time, this lib should be adapted to take into account the import needs of the new TOTAL mod. 

The following changes have been made from vanilla by Milo:

* Some of the files have been re-indented (tag order is not effected).
* The `DWARF` creature has templates from the `addon:Libs/Castes` addon inserted.
* Specialized `!SHARED_OBJECT` templates have been inserted where appropriate.
* `REACTION` and `BUILDING_WORKSHOP` templates have been inserted where appropriate.
* `@BUILD_KEY` templates have been inserted where appropriate.
* `ADDON_HOOKS` have been added to all entities.
* `!ENTITY_PLAYABLE` templates have been added to all entities.
* The dwarven outpost liaison has been modded to work as a trade rep if the dwarves are unplayable.
* Dwarves have the same progress triggers as humans.
* Dwarves have the alchemist labor permitted in their entity.
* All file IDs have been stripped.
* The files from "objects/text" have been added (with appropriate file extension).
* Added `#ENTITY_NOBLES` to all entities.

The following changes have been made from vanilla by Abadrausar:

* For each kind of DF object a directory has been created to group them, so the objects of type BODY are inside a directory named BODY, those of type CREATURE inside a directory named BODY CREATURE and so on.
* When a mod needs to import some vanilla objects of a type but not the rest, those objects are moved into a nested dir inside the base dir of the type. 
* Following the anterior principle we can find almost all vanilla creatures inside the base CREATURE directory, but then we remark that the wagon creature has been moved into the  CREATURE/equipment directory, the reason to do that is that even if some total mod does not use any vanilla creature you will always need a wagon to start the game ;)
* Until now I have only found a mod (Rise of the Mushroom Kingdom by IndigoFenix) that modifies the wagon for the rest of the total mods the vanilla creature wagon is the only vanilla creature that they need in order to have a working game
* Then some others creatures are separated to answer the particular needs of the MinDF, First Landing and in general any kind of TOTAL mods.

#  Haiku Vanilla -> https://s-media-cache-ak0.pinimg.com/736x/65/65/99/656599ba6db9c96db4c26ad1ce6c1c48.jpg
This is DF with the minimal number of active objects to have *all* the basics of a functional vanilla *without modifying any of them* only pruning all the non redundant object into a separate minimal directory nested inside the base of their type.
# What we have until now is:
1) 10 Creatures: 1 wagon,3 vermin,1 colonies (bees), 3 domestics (dog, cat and donkey) 2 civilized (dwarves and goblins) and nothing more. For the variety you have the titans, FB semiFB, night creatures and rest of random.
2) 34 Inorganics: 5 to enable adventure coins, 6 for the bronze line, 5 for the steel line, 3 for the deep specials, and 15 for the rest: stones and ores.
3) 21 Plants: 2 grass one subterranean (cave moss) another above ground (meadow-grass), 8 crops: 6 subterranean (Cave wheat, Dimple cup, Pig tail, Plump helmet, Quarry bush and Sweet pod) and 2 above ground (Sliver barb and Whip vine), 11 trees 7 subterranean (Black-cap, Blood thorn, Fungiwood, Goblin-cap, Nether-cap, Spore tree, Tower-cap and Tunnel tube)  and 4 above ground (glumprong, feather tree, highwood and willow) 
4) That makes a great total of 65 active objects to play with, nothing more (this is not strictly true) nothing less; AND HEY!, before you begin to protest, remember that this is a minimal mod done in the interest of SCIENCE, it could be FUN to play or not (it is up to you to see and decide), but this is not certainly the primary objective of this mod.

The number of active objects have been already dramatically reduced (good news for your FPS) but I think that it could be reduced even further, some more SCIENCE required... When you experiment trying to remove even more objects you will find some of the hard limits of the game.

# Some of those hard limits already more or less identified:
1) It seems that you must have at least two entities defined to elude world generation rejects.
2) A wagon creature must be define in order to be able to start the game in fort mode.
3) You need at least the vanilla base language files and then the specifics language files of the creature of your civilization if those creatures are able to speak.
4) BODY, BODY_DETAIL_PLAN, BOOK, BUILDING, DESCRIPTOR_COLOR, DESCRIPTOR_PATTERN, DESCRIPTOR_SHAPE, INTERACTION, ITEM, MATERIAL_TEMPLATE, REACTION, SECRET and TISSUE_TEMPLATE have not been investigated for pruning.

# Haiku Extreme mod: Arghhh Abomination!!!"
This is DF with the minimal number of active objects to have *all* the basics of a functional game.
# What we have until now is:
1) 4 versatile Creatures: 1 wagon,1 vermin colonies (amphibious fishable ubiquitous), 1 domestics (a kill vermin wagon puller) 1 civilized (dwarves into 2 separated entities) and nothing more. For the variety you have the titans, FB semiFB, night creatures and rest of random.
2) 34 Inorganics: 5 to enable adventure coins, 6 for the bronze line, 5 for the steel line, 3 for the deep specials, and 15 for the rest: stones and ores.(This one could be object of further simplification)
3) 3 multiuse Plants: 1 grass (subterranean and above ground) , 1 crop (subterranean and above ground), 1 tree (subterranean and above ground). 
4) That makes a great total of 41 active objects to play with, nothing more (this is not strictly true) nothing less; AND HEY!, before you begin to protest, remember that this is a minimal mod done in the interest of SCIENCE, it could be FUN to play or not (it is up to you to see and decide), but this is not certainly the primary objective of this mod.
