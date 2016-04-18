Hi Milo.
Fighting uphill is always more hard than strictly necessary,
I am porting some DF mods to the rubble system, you can see it there: https://github.com/Abadrausar/nanofort.
By modder the mods where I am actually involved  are:
Toady -> Vanilla (Complete, I have converted this to a fine grained library able of fine grained imports)
Button -> Modest: Almost complete.
Dirst -> The Earth Strikes back!: no changes.>
GM-X -> Dark AgesII: Complete.
IndigoFenix -> Mostly Mythical Mosters: Complete.
                -> Rise of the mushroom kingdom: Complete.
Lovechild -> All races playable: Complete.
Milo -> First Landing: Beginning to decompose into independents modules.
Putnam -> Fantastic Mini mods:Initiated.
           -> Sparking: Complete.
Wannabehero -> Deep Dwarven Domestication: Complete.
Abadrausar -> MinDF: Complete, This is DF with the minimal number of active objets to have all the basics of a functional vanilla, 6 creatures, less of 15 plants, two civs, 20 inorganics...
                -> Singularity: Only exist at conceptual level now, at release time it will integrate all the other mods.
                -> Humble: FPS aware vanilla collection of minimods.

In order to integrate all of those moods into the rubble system I follow some simple directives:
a) Rubble need all the files of an addon in the directory of the addon or mod, (even if an addon can activate others addons ;) when the original modder has used a hierarchy of dirs for different reasons these are preserved in the filenames.
b) If not already done by the mod original autor, one uppercase mod namespace acronym of between 3 and 7 characters is added to each non Vanilla file of the mod. This is done even  for the PNG and BMP files, in consequence the internal references to those files are mass replaced with notepad++ into the graphics txt files
c) The files whose directory destination is not the raw/object are given the double extension that signal that to the rubble system.
d) For each to be integrated mod I look with Gitextensions GUI the vanilla files that have been modified and capture the diffs in one or more rubble files using the following semantic delta operators that use a global cursor to signal the object  receiver of the changes:
[spoiler]{!TEMPLATE;@SET_VALUE;token;value;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
{!TEMPLATE;@ATTACH_BEFORE_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};%{preraws}[%{token}]}}
{!TEMPLATE;@ATTACH_AFTER_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}]%{preraws}}}
{!TEMPLATE;@REPLACE_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};%{preraws}}}[/spoiler]

But... What is all this system good for? Lets see some facts about it:
1) The Toady vanilla raws for DF 42.06 weight 1.55 Mb on disk.
2) Each other mod reuses some or all of the vanilla base files, even the total mods, in the next spoiler you can see the concrete parts of vanilla that are reused By the Rise of the Mushroom Kingdom mod
[spoiler]{
	"Header": "Rise of the Mushroom Kingdom 0.5.2 for DF42.06+",
	
	"Author": "IndigoFenix, Rubble porting by Abadrausar",
	"Version": "0.5.2",
	"DescFile": "addon.md",
	
	"Activates": [
		"Libs/Base", 
		"Libs/Castes",
		"Libs/Base/Vanilla/BODY",
		"Libs/Base/Vanilla/BODY_DETAIL_PLAN",
		"Libs/Base/Vanilla/BOOK",
		"Libs/Base/Vanilla/BUILDING",
		"Libs/Base/Vanilla/CREATURE_VARIATION",
		"Libs/Base/Vanilla/DESCRIPTOR_COLOR",
		"Libs/Base/Vanilla/DESCRIPTOR_PATTERN",
		"Libs/Base/Vanilla/DESCRIPTOR_PATTERN/pupil",
		"Libs/Base/Vanilla/DESCRIPTOR_SHAPE",
		"Libs/Base/Vanilla/INORGANIC",
		"Libs/Base/Vanilla/INORGANIC/Minimal",
		"Libs/Base/Vanilla/INORGANIC/Minimal/Candy",
		"Libs/Base/Vanilla/INORGANIC/Minimal/Steel",
		"Libs/Base/Vanilla/INORGANIC/Minimal/Bronze",
		"Libs/Base/Vanilla/INORGANIC/Minimal/Bronze/Adventure_Coins",
		"Libs/Base/Vanilla/INTERACTION",
		"Libs/Base/Vanilla/ITEM",
		"Libs/Base/Vanilla/LANGUAGE",
		"Libs/Base/Vanilla/MATERIAL_TEMPLATE",
		"Libs/Base/Vanilla/REACTION",
		"Libs/Base/Vanilla/SECRET",
		"Libs/Base/Vanilla/TISSUE_TEMPLATE"
		]
}[/spoiler]  
3) The addon that enables us to play Toady vanilla version of DF without ANY modif is composed by two files that mather, one dummy.rbl is empty the other addon.meta have the following content: 
[spoiler]
{	"Header": "Contains Rubblized versions of the Vanilla raw files that are used as an include or import base library able of fine grained imports",
	"Author": "Tarn Adams, Rubble porting by Milo Christiansen, Ababrausar make this addon a rubble library able of fine grained imports",
	"DescFile": "addon.md",
	"Activates": [
		"Libs/Base", 
		"Libs/Castes",
		"Libs/Base/Vanilla/BODY",
		"Libs/Base/Vanilla/BODY_DETAIL_PLAN",
		"Libs/Base/Vanilla/BOOK",
		"Libs/Base/Vanilla/BUILDING",
		"Libs/Base/Vanilla/CREATURE",
		"Libs/Base/Vanilla/CREATURE/equipment",
		"Libs/Base/Vanilla/CREATURE/leopard",
		"Libs/Base/Vanilla/CREATURE/Minimal",
		"Libs/Base/Vanilla/CREATURE_VARIATION",
		"Libs/Base/Vanilla/DESCRIPTOR_COLOR",
		"Libs/Base/Vanilla/DESCRIPTOR_PATTERN",
		"Libs/Base/Vanilla/DESCRIPTOR_PATTERN/pupil",
		"Libs/Base/Vanilla/DESCRIPTOR_SHAPE",
		"Libs/Base/Vanilla/INORGANIC",
		"Libs/Base/Vanilla/INORGANIC/Minimal",
		"Libs/Base/Vanilla/INORGANIC/Minimal/Candy",
		"Libs/Base/Vanilla/INORGANIC/Minimal/Steel",
		"Libs/Base/Vanilla/INORGANIC/Minimal/Bronze",
		"Libs/Base/Vanilla/INORGANIC/Minimal/Bronze/Adventure_Coins",
		"Libs/Base/Vanilla/INTERACTION",
		"Libs/Base/Vanilla/ITEM",
		"Libs/Base/Vanilla/LANGUAGE",
		"Libs/Base/Vanilla/LANGUAGE/dwarf",
		"Libs/Base/Vanilla/LANGUAGE/elf",
		"Libs/Base/Vanilla/LANGUAGE/goblin",
		"Libs/Base/Vanilla/LANGUAGE/human",
		"Libs/Base/Vanilla/MATERIAL_TEMPLATE",
		"Libs/Base/Vanilla/PLANT",
		"Libs/Base/Vanilla/PLANT/Minimal",
		"Libs/Base/Vanilla/REACTION",
		"Libs/Base/Vanilla/SECRET",
		"Libs/Base/Vanilla/TISSUE_TEMPLATE",
		"Libs/Base/Vanilla/Faction/Cavern/AnimalPeople",
		"Libs/Base/Vanilla/Faction/Civilized/Dwarves",
		"Libs/Base/Vanilla/Faction/Civilized/Elves",
		"Libs/Base/Vanilla/Faction/Civilized/Humans",
		"Libs/Base/Vanilla/Faction/Savage/Kobolds",
		"Libs/Base/Vanilla/Faction/Slavers/Goblins"
		]
}
[/spoiler]
4) So the complete rubble addon directory that enables us to play the original Toady  vanilla DF uses less of 2000 characters.
5) From now on each mod integrated into the rubble system is made compatible with all the others.

I have some quick questions :

When some rubble var is set

{@SET;OBJECT_CURSOR;REACTION:STEEL_MAKING}

and then I want to acces to the $OBJECT_CURSOR rubble global value from inside a lua function, Could I do something like next code?

rubble.template("!SET_VALUE", [[
	local target,value, repl = rubble.targs({...}, {"", "", ""})
	target = string.split(target, ":")
	local id = rubble.configvar(OBJECT_CURSOR)
	rubble.libs_base.sharedobject_walk(id, function(tag)
		if rubble.libs_base.matchtag(tag, target) then
			tag.Comments = "["..target..":"..value.."]"..repl..tag.Comments
			tag.CommentsOnly = true
		end
	end)
]])

It is possible to call a defined rubble template from normal lua code?
If yes could you give me an example with any rubble template from the library?