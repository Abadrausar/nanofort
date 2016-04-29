{@SET;<NAME>;<VALUE>;[<EXPAND>=true]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_base_libs_base.test (4 hits)
	Line 11: {@SET;TEST;Fun}
	Line 26: {@SET;TEMP;OK!}{ECHO;$TEMP}
	Line 38: {@SET;A;Y}{@SET;B;N}
	Line 38: {@SET;A;Y}{@SET;B;N}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_material_template_MODEST.rbl (13 hits)
	Line 9: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:SKIN_TEMPLATE}
	Line 20: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:FAT_TEMPLATE}
	Line 23: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:MUSCLE_TEMPLATE}
	Line 29: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:SINEW_TEMPLATE}
	Line 32: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:BONE_TEMPLATE}
	Line 38: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:CARTILAGE_TEMPLATE}
	Line 42: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:HAIR_TEMPLATE}
	Line 44: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:FEATHER_TEMPLATE}
	Line 47: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:SCALE_TEMPLATE}
	Line 62: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:NAIL_TEMPLATE}
	Line 65: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:TOOTH_TEMPLATE}
	Line 68: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:%{key}_TEMPLATE}
	Line 71: {@SET;OBJECT_CURSOR;MATERIAL_TEMPLATE:CHITIN_TEMPLATE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_plant_MODEST.rbl (1 hit)
	Line 1: {@SET;OBJECT_CURSOR;PLANT:MUSHROOM_HELMET_PLUMP}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_reaction_MODEST.rbl (2 hits)
	Line 30: {@SET;OBJECT_CURSOR;REACTION:PIG_IRON_MAKING}
	Line 33: {@SET;OBJECT_CURSOR;REACTION:STEEL_MAKING}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\creature_TESB.txt (1 hit)
	Line 3: {@IF;$TESB_HIDDEN_GEMS;NEVER;;{@SET;TESB_WYRM_EYES;"HIDDEN "}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\reaction_TESB.txt (6 hits)
	Line 8: {@IF;$TESB_LIVING_STONE;COMMON;{@SET;TESB_LIVING_PROB;.002}}
	Line 9: {@IF;$TESB_LIVING_STONE;RARE;{@SET;TESB_LIVING_PROB;.0005}}
	Line 10: {@IF;$TESB_LIVING_STONE;NEVER;{@SET;TESB_LIVING_PROB;0}}
	Line 12: {@IF;$TESB_HIDDEN_GEMS;COMMON;{@SET;TESB_GEM_PROB;.005}}
	Line 13: {@IF;$TESB_HIDDEN_GEMS;RARE;{@SET;TESB_GEM_PROB;.00125}}
	Line 14: {@IF;$TESB_HIDDEN_GEMS;NEVER;{@SET;TESB_GEM_PROB;0}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\GM-X\DarkAgesII\Libs\_DAII_Dark_Ages.rbl (6 hits)
	Line 25: {@SET;OBJECT_CURSOR;ENTITY:MOUNTAIN}
	Line 104: {@SET;OBJECT_CURSOR;ENTITY:FOREST}
	Line 121: {@SET;OBJECT_CURSOR;ENTITY:PLAINS}
	Line 195: {@SET;OBJECT_CURSOR;ENTITY:EVIL}
	Line 215: {@SET;OBJECT_CURSOR;ENTITY:SKULKING}
	Line 258: {@SET;OBJECT_CURSOR;ENTITY:SUBTERRANEAN_ANIMAL_PEOPLES}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Lovechild\All races playable\Libs\_All races playable.rbl (5 hits)
	Line 21: 	{@SET;OBJECT_CURSOR;ITEM_WEAPON:ITEM_WEAPON_%{key}}
	Line 27: {@SET;DWARF_PLAYABLE;true}
	Line 28: {@SET;ELF_PLAYABLE;true}
	Line 29: {@SET;HUMAN_PLAYABLE;true}
	Line 30: {@SET;GOBLIN_PLAYABLE;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\aaa_templates_flbase.rbl (1 hit)
	Line 3: 	{!REACTION_NEW_CATEGORY;%id;%name;%description;%key}{V;{@SET;FL_REACTION_CATEGORY;%id}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_furniture.txt (2 hits)
	Line 28: {V;{@SET;FL_REACTION_CATEGORY;FURNITURE}}
	Line 74: {V;{@SET;FL_REACTION_CATEGORY;TOOLS}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_tailor.txt (2 hits)
	Line 24: {V;{@SET;FL_REACTION_CATEGORY;TAILOR_CLOTHING}}
	Line 44: {V;{@SET;FL_REACTION_CATEGORY;TAILOR_OTHER}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test06\entity_test.txt (1 hit)
	Line 764: {@SET;OBJECT_CURSOR;ENTITY:MOUNTAIN}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test07\entity_test.txt (4 hits)
	Line 11: {@IF;%{token};{@VOID};{@VOID};{@SET;TOKEN_CURSOR;%{token}}}
	Line 12: {@IF;%{object};{@VOID};{@VOID};{@SET;OBJECT_CURSOR;%{object}}}
	Line 13: {@IF;%{nested};{@VOID};{@VOID};{@SET;TOKEN_NESTED;%{nested}}}}
	Line 42: {@SET;OBJECT_CURSOR;ENTITY:MOUNTAIN}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test08\entity_test.rbl (1 hit)
	Line 22: {@SET;OBJECT_CURSOR;ENTITY:MOUNTAIN}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test09\test.rbl (2 hits)
	Line 17: {@SET;OBJECT_CURSOR;REACTION:PIG_IRON_MAKING}
	Line 21: {@SET;OBJECT_CURSOR;REACTION:STEEL_MAKING}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test10\test.rbl (2 hits)
	Line 17: {@SET;OBJECT_CURSOR;REACTION:PIG_IRON_MAKING}
	Line 21: {@SET;OBJECT_CURSOR;REACTION:STEEL_MAKING}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test11\test.rbl (2 hits)
	Line 17: {@SET;OBJECT_CURSOR;REACTION:PIG_IRON_MAKING}
	Line 20: {@SET;OBJECT_CURSOR;REACTION:STEEL_MAKING}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test12\test.lua (2 hits)
	Line 17: {@SET;OBJECT_CURSOR;REACTION:PIG_IRON_MAKING}
	Line 21: {@SET;OBJECT_CURSOR;REACTION:STEEL_MAKING}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test14\test.rbl (2 hits)
	Line 17: {@SET;OBJECT_CURSOR;REACTION:PIG_IRON_MAKING}
	Line 20: {@SET;OBJECT_CURSOR;REACTION:STEEL_MAKING}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Base\Haiku\Extreme\Base\Extreme Haiku.rbl (3 hits)
	Line 87: {@SET;ABOVE_GROUND_BIOMES;
	Line 96: {@SET;SUBTERRANEAN_BIOMES;
	Line 112: {@SET;OBJECT_CURSOR;CREATURE:HONEY_BEE}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (3 hits)
	Line 70: 	{@SET;TEST;?Fun?}
	Line 155: 	{@SET;<NAME>;<VALUE>;[<EXPAND>=true]}
	Line 172: 	{@SET;TEST;Hello!}{ECHO;$TEST}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Basics.md (2 hits)
	Line 407: 	{E;{@SET;TEMP;abc}}{EXAMPLE;$TEMP}
	Line 418: 	{E;{@SET;TEMP;abc}}{EXAMPLE;&TEMP}
{@SCRIPT;<CODE>;[<TAG>="LangLua"]}
 Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Random\Fauna\Vertebrates\creature_planets_fauna_random_vertebrates.txt (1 hit)
	Line 2: {@SCRIPT;(rubble:stageparse (require "planets_fauna_random_vertebrates"))}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_bonecarver.txt (2 hits)
	Line 38: 	[REAGENT:shell:{@SCRIPT;return %mats * 2;LangLua}:NONE:NONE:NONE:NONE]
	Line 51: 	[REAGENT:shell:{@SCRIPT;return %mats * 2;LangLua}:NONE:NONE:NONE:NONE]
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_forge.txt (1 hit)
	Line 7: 	[REAGENT:bar:{@SCRIPT;return %{mats} * 150;LangLua}:BAR:NONE:INORGANIC:NONE]
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_plastics.txt (1 hit)
	Line 7: 	[REAGENT:bar:{@SCRIPT;return %{mats} * 150;LangLua}:BAR:NONE:NONE:NONE]
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_radio.txt (1 hit)
	Line 7: 	[REAGENT:price:{@SCRIPT;return %{cost} * 500;LangLua}:COIN:NONE:INORGANIC:CREDIT_COINS]
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_tailor.txt (1 hit)
	Line 7: 	[REAGENT:cloth:{@SCRIPT;return %{mats} * 10000;LangLua}:CLOTH:NONE:NONE:NONE]
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (1 hit)
	Line 145: 	{@SCRIPT;<CODE>;[<TAG>="LangLua"]}
{@IF;<STRING1>;<STRING2>;<THEN_PRERAWS>;[<ELSE_PRERAWS>=""]}
 Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_base_libs_base.test (5 hits)
	Line 39: {@IF;$A;Y;
	Line 40: 	{@IF;$B;Y;
	Line 61: {@IF;$TEMP;0;Index OK.;Bad Index.}
	Line 63: {@IF;$TEMP;Testing...;Can read key.;Could not read key.}
	Line 73: {@IF;$TEMP;Testing...;Can read key.;Could not read key.}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_sobject_libs_base.test (1 hit)
	Line 29: {#;{@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;INORGANIC:TEST_SOD_C2}}{@IF;$TEMP;t;OK!;ERROR!}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Bonecarver\reaction_user_bonecarver.txt (1 hit)
	Line 11: 	[SKILL:BONECARVE]{@IF;%fake;true;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\building_TESB.txt (1 hit)
	Line 3: {@IF;$TESB_LIVING_STONE $TESB_HIDDEN_GEMS;NEVER NEVER;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\creature_TESB.txt (4 hits)
	Line 3: {@IF;$TESB_HIDDEN_GEMS;NEVER;;{@SET;TESB_WYRM_EYES;"HIDDEN "}}
	Line 4: {@IF;$TESB_LIVING_STONE;NEVER;
	Line 1149: {@IF;$TESB_PET_ROCKS;NO;
	Line 1185: 		{@IF;%{name};rock salt;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\descriptor_pattern_TESB.txt (1 hit)
	Line 3: {@IF;${TESB_LIVING_STONE};NEVER;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\graphics_TESB_creatures.graphics.txt (1 hit)
	Line 1: {@IF;$TESB_GRAPHICS;NO;Creature graphics have been disabled.;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\inorganic_TESB_stone.txt (2 hits)
	Line 3: {@IF;$TESB_HIDDEN_GEMS;NEVER;
	Line 340: {@IF;$TESB_LIVING_STONE;NEVER;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\interaction_TESB.txt (7 hits)
	Line 225: {@IF;$TESB_SECRETS;NO;
	Line 251: 	{@IF;%{sphere3};nil;No third sphere.;
	Line 274:             {@IF;%{type};standard;
	Line 287: 			{@IF;%{type};magma;
	Line 300: 			{@IF;%{type};earthquake;
	Line 314: 			{@IF;%{type};aquifer;
	Line 327: 			{@IF;%{type};flux;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\plant_TESB.txt (1 hit)
	Line 5: {@IF;$TESB_PLANTS;NO;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\reaction_TESB.txt (13 hits)
	Line 8: {@IF;$TESB_LIVING_STONE;COMMON;{@SET;TESB_LIVING_PROB;.002}}
	Line 9: {@IF;$TESB_LIVING_STONE;RARE;{@SET;TESB_LIVING_PROB;.0005}}
	Line 10: {@IF;$TESB_LIVING_STONE;NEVER;{@SET;TESB_LIVING_PROB;0}}
	Line 12: {@IF;$TESB_HIDDEN_GEMS;COMMON;{@SET;TESB_GEM_PROB;.005}}
	Line 13: {@IF;$TESB_HIDDEN_GEMS;RARE;{@SET;TESB_GEM_PROB;.00125}}
	Line 14: {@IF;$TESB_HIDDEN_GEMS;NEVER;{@SET;TESB_GEM_PROB;0}}
	Line 17: {@IF;$TESB_LIVING_STONE $TESB_HIDDEN_GEMS;NEVER NEVER;
	Line 24: {@IF;$TESB_LIVING_STONE $TESB_HIDDEN_GEMS $TESB_PET_ROCKS $TESB_PLANTS $TESB_SECRETS;NEVER NEVER NO NO NO;
	Line 31: {@IF;$TESB_LIVING_STONE $TESB_HIDDEN_GEMS;NEVER NEVER;
	Line 109: {@IF;$TESB_PLANTS;NO;
	Line 111: {@IF;$TESB_HIDDEN_GEMS;NEVER;
	Line 179: {@IF;$TESB_PLANTS;NO;
	Line 181: {@IF;$TESB_HIDDEN_GEMS;NEVER;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Lovechild\All races playable\Libs\_All races playable.rbl (9 hits)
	Line 67: 		{@IF;%{soldier};{@VOID};{@VOID};[SQUAD:10:%{soldier}:%{soldier}s]}}
	Line 71: 		{@IF;%{replaced};{@VOID};{@VOID};[REPLACED_BY:%{replaced}]}}
	Line 74: 		{@IF;%{responsability2};{@VOID};{@VOID};[RESPONSIBILITY:%{responsability2}]}
	Line 75: 		{@IF;%{responsability3};{@VOID};{@VOID};[RESPONSIBILITY:%{responsability3}]}
	Line 76: 		{@IF;%{responsability4};{@VOID};{@VOID};[RESPONSIBILITY:%{responsability4}]}}
	Line 79: 		{@IF;%{appointer2};{@VOID};{@VOID};[RESPONSIBILITY:%{appointer2}]}
	Line 80: 		{@IF;%{appointer3};{@VOID};{@VOID};[RESPONSIBILITY:%{appointer3}]}
	Line 81: 		{@IF;%{appointer4};{@VOID};{@VOID};[RESPONSIBILITY:%{appointer4}]}}
	Line 324: 		{@IF;%{replaced};{@VOID};{@VOID};[REPLACED_BY:%{replaced}]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_bonecarver.txt (2 hits)
	Line 19: {@IF;%{fake};NO;;{V}
	Line 44: }{@IF;%{fake};NO;;{V}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test07\entity_test.txt (4 hits)
	Line 5: 		{@IF;$TEMP;t;"\t%{val}\n"}
	Line 11: {@IF;%{token};{@VOID};{@VOID};{@SET;TOKEN_CURSOR;%{token}}}
	Line 12: {@IF;%{object};{@VOID};{@VOID};{@SET;OBJECT_CURSOR;%{object}}}
	Line 13: {@IF;%{nested};{@VOID};{@VOID};{@SET;TOKEN_NESTED;%{nested}}}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Base\Haiku\Extreme\Base\Extreme Haiku.rbl (12 hits)
	Line 2: 	{@IF;%{oil};true;
	Line 16: 	{@IF;%{flour};true;
	Line 18: 		[STATE_NAME_ADJ:ALL_SOLID:%{fruit} {@IF;%{dye};true;dye;@VOID}{@IF;%{flour};true;flour;@VOID}]
	Line 18: 		[STATE_NAME_ADJ:ALL_SOLID:%{fruit} {@IF;%{dye};true;dye;@VOID}{@IF;%{flour};true;flour;@VOID}]
	Line 23: 		{@IF;%{flour};true;[EDIBLE_COOKED];@VOID}
	Line 24: 		{@IF;%{dye};true;[POWDER_DYE:%{color}];@VOID}
	Line 27: 	{@IF;%{syrup};true;
	Line 39: 	{@IF;%{alcohol};true;
	Line 50: 	{@IF;%{thread};true;
	Line 59: 	{@IF;%{nogrowth};true;
	Line 67: 		{@IF;%{alcohol};true;
	Line 73: 		{@IF;%{oil};true;
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (4 hits)
	Line 180: 	{@IF;<STRING1>;<STRING2>;<THEN_PRERAWS>;[<ELSE_PRERAWS>=""]}
	Line 187: 	{@IF;$TEST_VAR;YES;[FOO];[BAR]}
	Line 268: 		}{@IF;$TEMP;t;"\t%{val}\n"}
	Line 1025: 	}{@IF;$TEMP;t;IRON is defined!;IRON is not defined!}
{@IF_ACTIVE;<ADDON>;<THEN_PRERAWS>;[<ELSE_PRERAWS>=""]}
 Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_base_libs_base.test (1 hit)
	Line 53: {@IF_ACTIVE;Libs/Base;OK!;Oops!}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Archery\reaction_user_archery.txt (1 hit)
	Line 34: {@IF_ACTIVE;User/Gnome;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Glass Forge\reaction_user_glass_forge.txt (2 hits)
	Line 80: {@IF_ACTIVE;User/Saurian;
	Line 92: {@IF_ACTIVE;User/Gnome;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Warcrafter\Adventure\reaction_user_warcrafter_adventure.txt (1 hit)
	Line 8: {@IF_ACTIVE;User/Tanning;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Uncle spam\Base\leathery vegetables\Uncle span.rbl (1 hit)
	Line 1: {@IF_ACTIVE;User/Fix/Usable Scale;{SHARED_OBJECT_ADD;SCALE_TEMPLATE;[MEAT][BUTCHER_SPECIAL:MEAT:NONE][MEAT_NAME:raw:scale:scale][REACTION_CLASS:SKIN]}}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (1 hit)
	Line 191: 	{@IF_ACTIVE;<ADDON>;<THEN_PRERAWS>;[<ELSE_PRERAWS>=""]}
{@IF_SKIP;<STRING1>;<STRING2>}
{!SHARED_OBJECT_CATEGORY;<ID>;<CATEGORY>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\aaa_sitem_libs_base.pre.lua (2 hits)
	Line 9: 		return rubble.parse("{!SHARED_OBJECT;"..id..";\n[ITEM_FOOD:"..id.."]\n\t"..def.."\n}{!SHARED_OBJECT_CATEGORY;"..id..";ITEM_FOOD}")
	Line 42: 		"{!SHARED_OBJECT_CATEGORY;ITEM_"..typ..":"..id..";ITEM_"..typ.."}")
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\aaa_sobject_libs_base.pre.lua (5 hits)
	Line 257: 		"}{!SHARED_OBJECT_CATEGORY;"..v..":%{id};"..v.."}"
	Line 265: 	"}{!SHARED_OBJECT_CATEGORY;REACTION:%{id};REACTION}"
	Line 272: 	"}{!SHARED_OBJECT_CATEGORY;BUILDING_WORKSHOP:%{id};BUILDING_WORKSHOP}"
	Line 279: 	"}{!SHARED_OBJECT_CATEGORY;BUILDING_FURNACE:%{id};BUILDING_FURNACE}"
	Line 287: 	"}{!SHARED_OBJECT_CATEGORY;ENTITY:%{id};ENTITY}"
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Industry\libs_industry.pre.lua (3 hits)
	Line 13: 	"}{!SHARED_OBJECT_CATEGORY;REACTION:%{id};REACTION}"
	Line 20: 	"}{!SHARED_OBJECT_CATEGORY;BUILDING_WORKSHOP:%{id};BUILDING_WORKSHOP}"
	Line 27: 	"}{!SHARED_OBJECT_CATEGORY;BUILDING_FURNACE:%{id};BUILDING_FURNACE}"
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test14\aaa_sobject_libs_base.pre.lua (5 hits)
	Line 273: 		"}{!SHARED_OBJECT_CATEGORY;"..v..":%{id};"..v.."}"
	Line 281: 	"}{!SHARED_OBJECT_CATEGORY;REACTION:%{id};REACTION}"
	Line 288: 	"}{!SHARED_OBJECT_CATEGORY;BUILDING_WORKSHOP:%{id};BUILDING_WORKSHOP}"
	Line 295: 	"}{!SHARED_OBJECT_CATEGORY;BUILDING_FURNACE:%{id};BUILDING_FURNACE}"
	Line 303: 	"}{!SHARED_OBJECT_CATEGORY;ENTITY:%{id};ENTITY}"
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Libs\Delta\aaa_sobject_libs_base.pre.lua (5 hits)
	Line 273: 		"}{!SHARED_OBJECT_CATEGORY;"..v..":%{id};"..v.."}"
	Line 281: 	"}{!SHARED_OBJECT_CATEGORY;REACTION:%{id};REACTION}"
	Line 288: 	"}{!SHARED_OBJECT_CATEGORY;BUILDING_WORKSHOP:%{id};BUILDING_WORKSHOP}"
	Line 295: 	"}{!SHARED_OBJECT_CATEGORY;BUILDING_FURNACE:%{id};BUILDING_FURNACE}"
	Line 303: 	"}{!SHARED_OBJECT_CATEGORY;ENTITY:%{id};ENTITY}"
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (2 hits)
	Line 252: 	{!SHARED_OBJECT_CATEGORY;<ID>;<CATEGORY>}
	Line 565: 	}{!SHARED_OBJECT_CATEGORY;ITEM_<TYPE>:<ID>;ITEM_<TYPE>}
{SHARED_OBJECT_EXISTS;<ID>;<THEN_PRERAWS>;[<ELSE_PRERAWS>=""]}
{#SHARED_OBJECT_EXISTS;<ID>;<THEN_PRERAWS>;[<ELSE_PRERAWS>=""]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\creature_TESB.txt (9 hits)
	Line 43: 	{SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 251: 	{SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 424: 	{SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 638: 	{SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 694: 	{SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 750: 	{SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 807: 	{SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 855: 	{SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 1172: 	{SHARED_OBJECT_EXISTS;INORGANIC:%{id};
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\inorganic_TESB_stone.txt (1 hit)
	Line 353: {SHARED_OBJECT_EXISTS;INORGANIC:%{id};
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\interaction_TESB.txt (6 hits)
	Line 5: {SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 26: {SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 47: {SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 76: {SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 97: {SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 233: {SHARED_OBJECT_EXISTS;INORGANIC:%{id};
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\plant_TESB.txt (1 hit)
	Line 12: {SHARED_OBJECT_EXISTS;INORGANIC:HIDDEN %{id};
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\reaction_TESB.txt (2 hits)
	Line 36: {SHARED_OBJECT_EXISTS;INORGANIC:%{id};
	Line 115: {SHARED_OBJECT_EXISTS;INORGANIC:%{id};
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_forge.txt (1 hit)
	Line 55: {SHARED_OBJECT_EXISTS;ITEM_TRAPCOMP:DFHACK_POWERED_DRILLING_RIG_BIT;
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (2 hits)
	Line 273: 	{SHARED_OBJECT_EXISTS;<ID>;<THEN_PRERAWS>;[<ELSE_PRERAWS>=""]}
	Line 274: 	{#SHARED_OBJECT_EXISTS;<ID>;<THEN_PRERAWS>;[<ELSE_PRERAWS>=""]}
{SHARED_OBJECT_ADD;<ID>;<PRERAWS>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Wood Additions\libs_wood_additions.rbl (1 hit)
	Line 2: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:WOOD_TEMPLATE;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Remove\Animalmen\Libs\ExceptFromLegends.rbl (1 hit)
	Line 2: 	{SHARED_OBJECT_ADD;CREATURE_VARIATION:%{key};[CV_NEW_TAG:DOES_NOT_EXIST]}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Remove\Giants\Libs\ExceptFromLegends.rbl (1 hit)
	Line 1: {SHARED_OBJECT_ADD;CREATURE_VARIATION:GIANT;[CV_NEW_TAG:DOES_NOT_EXIST]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Tanning\uncle span.rbl (4 hits)
	Line 1: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:SHELL_TEMPLATE;[BUTCHER_SPECIAL:MEAT:NONE][MEAT_NAME:raw:shell:shell]} }
	Line 2: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:CHITIN_TEMPLATE;[BUTCHER_SPECIAL:MEAT:NONE][MEAT_NAME:raw:chitin:chitin]} }
	Line 3: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:SCALE_TEMPLATE;[BUTCHER_SPECIAL:MEAT:NONE][MEAT_NAME:raw:scale:scale]} }
	Line 4: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:SKIN_TEMPLATE;[BUTCHER_SPECIAL:MEAT:NONE][MEAT_NAME:raw:hide:hide]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\ThirdParty\Steam Engine\item_user_steam_engine.txt (1 hit)
	Line 11: }{!ITEM_CLASS;ADDON_HOOK_PLAYABLE}{SHARED_OBJECT_ADD;USER_DFHACK_POWERED_METAL_PRESS_ITEMS;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Grind Sand\user_grind_sand.rbl (1 hit)
	Line 2: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:STONE_TEMPLATE;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Pottery\Glaze\user_pottery_glaze.rbl (19 hits)
	Line 10: {SHARED_OBJECT_ADD;INORGANIC:HEMATITE;
	Line 16: {SHARED_OBJECT_ADD;INORGANIC:LIMONITE;
	Line 22: {SHARED_OBJECT_ADD;INORGANIC:GARNIERITE;
	Line 28: {SHARED_OBJECT_ADD;INORGANIC:NATIVE_GOLD;
	Line 34: {SHARED_OBJECT_ADD;INORGANIC:NATIVE_SILVER;
	Line 40: {SHARED_OBJECT_ADD;INORGANIC:NATIVE_COPPER;
	Line 46: {SHARED_OBJECT_ADD;INORGANIC:MALACHITE;
	Line 52: {SHARED_OBJECT_ADD;INORGANIC:GALENA;
	Line 58: {SHARED_OBJECT_ADD;INORGANIC:SPHALERITE;
	Line 64: {SHARED_OBJECT_ADD;INORGANIC:CASSITERITE;
	Line 70: {SHARED_OBJECT_ADD;INORGANIC:CINNABAR;
	Line 76: {SHARED_OBJECT_ADD;INORGANIC:COBALTITE;
	Line 82: {SHARED_OBJECT_ADD;INORGANIC:TETRAHEDRITE;
	Line 88: {SHARED_OBJECT_ADD;INORGANIC:HORN_SILVER;
	Line 94: {SHARED_OBJECT_ADD;INORGANIC:ORPIMENT;
	Line 100: {SHARED_OBJECT_ADD;INORGANIC:STIBNITE;
	Line 106: {SHARED_OBJECT_ADD;INORGANIC:MAGNETITE;
	Line 112: {SHARED_OBJECT_ADD;INORGANIC:PYROLUSITE;
	Line 118: {SHARED_OBJECT_ADD;INORGANIC:PITCHBLENDE;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Pottery\More Porcelain\user_pottery_more_porcelain.rbl (2 hits)
	Line 2: {SHARED_OBJECT_ADD;INORGANIC:ORTHOCLASE;
	Line 6: {SHARED_OBJECT_ADD;INORGANIC:MICROCLINE;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_inorganic_stone_MODEST.rbl (2 hits)
	Line 2: 	{SHARED_OBJECT_ADD;INORGANIC:%{key};
	Line 6: 	{SHARED_OBJECT_ADD;INORGANIC:%{key};
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_item_MODEST.rbl (1 hit)
	Line 1: {SHARED_OBJECT_ADD;ITEM_HELM:ITEM_HELM_HELM;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_material_template_MODEST.rbl (4 hits)
	Line 1: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:STONE_TEMPLATE;
	Line 3: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:METAL_TEMPLATE;
	Line 78: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:FLAME_TEMPLATE;
	Line 80: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:SPONGE_TEMPLATE;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_plant_MODEST.rbl (1 hit)
	Line 12: {SHARED_OBJECT_ADD;PLANT:MUSHROOM_HELMET_PLUMP;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_tissue_template_MODEST.rbl (2 hits)
	Line 3: 	{SHARED_OBJECT_ADD;TISSUE_TEMPLATE:%{key}_TEMPLATE;[HEALING_RATE:%{val}]}}
	Line 6: {SHARED_OBJECT_ADD;TISSUE_TEMPLATE:HOOF_TEMPLATE;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Lovechild\All races playable\Libs\_All races playable.rbl (8 hits)
	Line 7: 	{SHARED_OBJECT_ADD;CREATURE:%{key};
	Line 10: 	{SHARED_OBJECT_ADD;CREATURE:%{key};
	Line 15: 	{SHARED_OBJECT_ADD;CREATURE:%{key};
	Line 18: {SHARED_OBJECT_ADD;BEAK_DOG;[PET_EXOTIC][TRAINABLE][PACK_ANIMAL][WAGON_PULLER][TRADE_CAPACITY:2000]}
	Line 31: {SHARED_OBJECT_ADD;FOREST;
	Line 54: {SHARED_OBJECT_ADD;PLAINS;[SITE_CONTROLLABLE][TOOL:ITEM_TOOL_STEPLADDER]}
	Line 56: {SHARED_OBJECT_ADD;EVIL;[SITE_CONTROLLABLE][ALL_MAIN_POPS_CONTROLLABLE][COMMON_DOMESTIC_PULL][TOOL:ITEM_TOOL_STEPLADDER]}
	Line 57: {SHARED_OBJECT_ADD;SKULKING;[SITE_CONTROLLABLE][TOOL:ITEM_TOOL_STEPLADDER]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Putnam\Fantastic Mini mods\Libs\Materials Plus\Libs\_inorganic_PFMM.rbl (8 hits)
	Line 2: 	{SHARED_OBJECT_ADD;INORGANIC:%{key};
	Line 5: 	{SHARED_OBJECT_ADD;INORGANIC:%{key};
	Line 8: 	{SHARED_OBJECT_ADD;INORGANIC:%{key};
	Line 11: 	{SHARED_OBJECT_ADD;INORGANIC:%{key};
	Line 14: 	{SHARED_OBJECT_ADD;INORGANIC:%{key};
	Line 17: 	{SHARED_OBJECT_ADD;INORGANIC:%{key};
	Line 20: 	{SHARED_OBJECT_ADD;INORGANIC:%{key};
	Line 23: 	{SHARED_OBJECT_ADD;INORGANIC:%{key};
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Uncle spam\Base\leathery vegetables\Uncle span.rbl (2 hits)
	Line 1: {@IF_ACTIVE;User/Fix/Usable Scale;{SHARED_OBJECT_ADD;SCALE_TEMPLATE;[MEAT][BUTCHER_SPECIAL:MEAT:NONE][MEAT_NAME:raw:scale:scale][REACTION_CLASS:SKIN]}}
	Line 2: {SHARED_OBJECT_ADD;SKIN_TEMPLATE;[MEAT][BUTCHER_SPECIAL:MEAT:NONE][MEAT_NAME:raw:hide:hide][REACTION_CLASS:SKIN]
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Wannabehero\Deep dwarven domestication\Base\_Deep dwarven domestication.rbl (3 hits)
	Line 4: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:NAIL_TEMPLATE;[HORN]}
	Line 5: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:CHITIN_TEMPLATE;[SHELL]}
	Line 6: {SHARED_OBJECT_ADD;MATERIAL_TEMPLATE:FEATHER_TEMPLATE;[PEARL]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test01\plant_test.txt (1 hit)
	Line 61: {SHARED_OBJECT_ADD;PLANT:MUSHROOM_HELMET_PLUMP_SPROUTING;[ALL_NAMES:plump helmet sproutings]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test02\plant_test2.txt (1 hit)
	Line 5: {SHARED_OBJECT_ADD;PLANT:MUSHROOM_HELMET_PLUMP_INTENSIVE;[ALL_NAMES:plump helmet intensive]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Base\Haiku\Extreme\Base\Extreme Haiku.rbl (4 hits)
	Line 100: {SHARED_OBJECT_ADD;PLANT:NETHER_CAP;{@TREE_GIVE_FRUIT;DARK_INDIGO;1:0:0;nether-cap;pit;beer;true;true;true;true;true;true;true}$ABOVE_GROUND_BIOMES}
	Line 101: {SHARED_OBJECT_ADD;PLANT:MUSHROOM_HELMET_PLUMP;{@TREE_GIVE_FRUIT;DARK_INDIGO;5:0:0;plump helmet;spawn;wine;true;true;true;true;true;true;false}$ABOVE_GROUND_BIOMES}
	Line 102: {SHARED_OBJECT_ADD;PLANT:MEADOW-GRASS;$SUBTERRANEAN_BIOMES}
	Line 103: {SHARED_OBJECT_ADD;CREATURE:LEOPARD;
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (1 hit)
	Line 284: 	{SHARED_OBJECT_ADD;<ID>;<PRERAWS>}
{REGISTER_REACTION_CLASS;<ID>;<CLASS>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Lamination\material_template_user_lamination.txt (4 hits)
	Line 1: [OBJECT:MATERIAL_TEMPLATE]{REGISTER_REACTION_CLASS;MATERIAL_TEMPLATE:WOOD_TEMPLATE;NON_LAMINATED_WOOD_MAT}
	Line 53: {REGISTER_REACTION_CLASS;MATERIAL_TEMPLATE:LAMINATE_LOG_TEMPLATE;WOOD_MAT}
	Line 54: {REGISTER_REACTION_CLASS;MATERIAL_TEMPLATE:LAMINATE_LOG_TEMPLATE;MILLABLE_WOOD}
	Line 100: }{REGISTER_REACTION_CLASS;MATERIAL_TEMPLATE:GLUE_TEMPLATE;GLUE_MAT}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Sawmill\reaction_basics_user_sawmill.txt (1 hit)
	Line 1: [OBJECT:REACTION]{REGISTER_REACTION_CLASS;MATERIAL_TEMPLATE:WOOD_TEMPLATE;MILLABLE_WOOD}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\ThirdParty\Spatter\reaction_user_dfhack_spatter.txt (2 hits)
	Line 3: {REGISTER_REACTION_CLASS;MATERIAL_TEMPLATE:CREATURE_EXTRACT_TEMPLATE;EXTRACT}
	Line 4: {REGISTER_REACTION_CLASS;MATERIAL_TEMPLATE:PLANT_EXTRACT_TEMPLATE;EXTRACT}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (1 hit)
	Line 290: 	{REGISTER_REACTION_CLASS;<ID>;<CLASS>}
{REGISTER_REACTION_PRODUCT;<ID>;<CLASS>;<PRODUCT>}
{SHARED_OBJECT_KILL_TAG;<ID>;<TAG>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_sobject_libs_base.test (2 hits)
	Line 42: {SHARED_OBJECT_KILL_TAG;TEST_SOKT;BAR}
	Line 43: {SHARED_OBJECT_KILL_TAG;TEST_SOKT;TEST:X:Y}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Remove\From subterranean entity\Libs\AllExceptOne.rbl (1 hit)
	Line 2: 	{SHARED_OBJECT_KILL_TAG;ENTITY:SUBTERRANEAN_ANIMAL_PEOPLES;CREATURE:%{key}}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Remove\Ubiquitous vermin\Libs\ubiquitous.rbl (1 hit)
	Line 2: 	{SHARED_OBJECT_KILL_TAG;CREATURE:%{key};UBIQUITOUS}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Cat Damper\user_cat_damper.rbl (2 hits)
	Line 2: {SHARED_OBJECT_KILL_TAG;CREATURE:CAT;ADOPTS_OWNER}
	Line 3: {SHARED_OBJECT_KILL_TAG;CREATURE:CAT;RETURNS_VERMIN_KILLS_TO_OWNER}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Diggable Slade\user_diggable_slade.rbl (1 hit)
	Line 2: {SHARED_OBJECT_KILL_TAG;INORGANIC:SLADE;UNDIGGABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Lovechild\All races playable\Libs\_All races playable.rbl (2 hits)
	Line 53: {SHARED_OBJECT_KILL_TAG;PLAINS;SITE_VARIABLE_POSITIONS}
	Line 55: {SHARED_OBJECT_KILL_TAG;EVIL;SITE_VARIABLE_POSITIONS}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (2 hits)
	Line 306: 	{SHARED_OBJECT_KILL_TAG;<ID>;<TAG>}
	Line 322: 	{SHARED_OBJECT_KILL_TAG;SLADE;UNDIGGABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\other\tutorial.md (1 hit)
	Line 371: {SHARED_OBJECT_KILL_TAG;<ID>;<TAG>}
{SHARED_OBJECT_REPLACE_TAG;<ID>;<TAG>;<REPLACEMENT>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_sobject_libs_base.test (2 hits)
	Line 50: {SHARED_OBJECT_REPLACE_TAG;TEST_SORT;BAR;[PUB]}
	Line 51: {SHARED_OBJECT_REPLACE_TAG;TEST_SORT;TEST:X:&;[TEST:M:N:O]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Bigger\Containers\Libs\Bigger containers.rbl (4 hits)
	Line 1: {SHARED_OBJECT_REPLACE_TAG;ITEM_TOOL_LARGE_POT;CONTAINER_CAPACITY;[CONTAINER_CAPACITY:600000]}
	Line 2: {SHARED_OBJECT_REPLACE_TAG;ITEM_TOOL_WHEELBARROW;CONTAINER_CAPACITY;[CONTAINER_CAPACITY:1200000]}
	Line 3: {SHARED_OBJECT_REPLACE_TAG;ITEM_TOOL_MINECART;CONTAINER_CAPACITY;[CONTAINER_CAPACITY:6000000]}
	Line 4: {SHARED_OBJECT_REPLACE_TAG;ITEM_TOOL_POUCH;CONTAINER_CAPACITY;[CONTAINER_CAPACITY:10000]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\All year around subterranean crops\Libs\All year around subterranean crops.rbl (4 hits)
	Line 1: {SHARED_OBJECT_REPLACE_TAG;PLANT:GRASS_WHEAT_CAVE;AUTUMN;[AUTUMN][WINTER][SPRING]}
	Line 2: {SHARED_OBJECT_REPLACE_TAG;PLANT:GRASS_TAIL_PIG;AUTUMN;[AUTUMN][WINTER][SPRING]}
	Line 3: {SHARED_OBJECT_REPLACE_TAG;PLANT:BUSH_QUARRY;AUTUMN;[AUTUMN][WINTER]}
	Line 4: {SHARED_OBJECT_REPLACE_TAG;PLANT:POD_SWEET;SUMMER;[SUMMER][AUTUMN][WINTER]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Intensive Sprouting\Libs\plant_zISS.txt (6 hits)
	Line 3: {SHARED_OBJECT_REPLACE_TAG;PLANT:{@STR_TO_ID;%{object}_%{variety}};NAME;[NAME:%{name} %{variety}]}
	Line 4: {SHARED_OBJECT_REPLACE_TAG;PLANT:{@STR_TO_ID;%{object}_%{variety}};NAME_PLURAL;[NAME_PLURAL:%{name}s %{variety}]}
	Line 5: {SHARED_OBJECT_REPLACE_TAG;PLANT:{@STR_TO_ID;%{object}_%{variety}};ADJ:[ADJ:%{name} %{variety}]}
	Line 6: {SHARED_OBJECT_REPLACE_TAG;PLANT:{@STR_TO_ID;%{object}_%{variety}};SEED;[SEED:%{name} %{variety} %{seed}:%{name} %{variety} %{seed}s:%{color}:LOCAL_PLANT_MAT:SEED]}
	Line 7: {SHARED_OBJECT_REPLACE_TAG;PLANT:{@STR_TO_ID;%{object}_%{variety}};GROWDUR;[GROWDUR:%{growdur}]}  
	Line 8: {SHARED_OBJECT_REPLACE_TAG;PLANT:{@STR_TO_ID;%{object}_%{variety}};CLUSTERSIZE;[CLUSTERSIZE:%{clustersize}]}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Intensive Sprouting\Libs\Z\Intensive sprouting2\plant_ZISS.txt (6 hits)
	Line 3: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};NAME;[NAME:%{name} %{variety}]}
	Line 4: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};NAME_PLURAL;[NAME_PLURAL:%{name}s %{variety}]}
	Line 5: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};ADJ:[ADJ:%{name}] %{variety}}
	Line 6: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};SEED;[SEED:%{name} %{variety} %{seed}:%{name} %{variety} %{seed}s:%{color}:LOCAL_PLANT_MAT:SEED]}
	Line 7: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};GROWDUR;[GROWDUR:%{growdur}]}  
	Line 8: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};CLUSTERSIZE;[CLUSTERSIZE:%{clustersize}]}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Intensive Sprouting\Libs\Z\Intensive sprouting3\plant_ZISS.txt (6 hits)
	Line 3: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};NAME;[NAME:%{name} %{variety}]}
	Line 4: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};NAME_PLURAL;[NAME_PLURAL:%{name}s %{variety}]}
	Line 5: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};ADJ:[ADJ:%{name}] %{variety}}
	Line 6: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};SEED;[SEED:%{name} %{variety} %{seed}:%{name} %{variety} %{seed}s:%{color}:LOCAL_PLANT_MAT:SEED]}
	Line 7: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};GROWDUR;[GROWDUR:%{growdur}]}  
	Line 8: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};CLUSTERSIZE;[CLUSTERSIZE:%{clustersize}]}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Tanning\uncle span.rbl (1 hit)
	Line 5: {SHARED_OBJECT_REPLACE_TAG;REACTION:TAN_A_HIDE;REAGENT;[REAGENT:A:600:GLOB:NONE:NONE:NONE][USE_BODY_COMPONENT]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\c_variation_MODEST.txt (3 hits)
	Line 2: {SHARED_OBJECT_REPLACE_TAG;CREATURE_VARIATION:ANIMAL_PERSON;CV_NEW_TAG:LOCAL_POPS_PRODUCE_HEROES;
	Line 55: {SHARED_OBJECT_REPLACE_TAG;CREATURE_VARIATION:ANIMAL_PERSON_LEGLESS;CV_NEW_TAG:LOCAL_POPS_PRODUCE_HEROES;
	Line 258: {SHARED_OBJECT_REPLACE_TAG;CREATURE_VARIATION:GIANT;CV_NEW_TAG:CHANGE_FREQUENCY_PERC;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_material_template_MODEST.rbl (1 hit)
	Line 5: {SHARED_OBJECT_REPLACE_TAG;MATERIAL_TEMPLATE:WOOD_TEMPLATE;BOILING_POINT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_reaction_MODEST.rbl (9 hits)
	Line 1: {!TEMPLATE;!SET_VALUE3;token;value;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 2: {!TEMPLATE;ATTACH_BEFORE_TAG;token;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};
	Line 5: {!TEMPLATE;ATTACH_AFTER_TAG;token;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};
	Line 8: {!TEMPLATE;REPLACE_TAG;token;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};
	Line 11: {SHARED_OBJECT_REPLACE_TAG;REACTION:TAN_A_HIDE;REAGENT:A;
	Line 17: {SHARED_OBJECT_REPLACE_TAG;REACTION:PROCESS_PLANT_TO_BAG;PRODUCT:100:1:SEEDS;
	Line 21: {SHARED_OBJECT_REPLACE_TAG;REACTION:MAKE_PARCHMENT;REAGENT:A:1;
	Line 27: {SHARED_OBJECT_REPLACE_TAG;REACTION:PIG_IRON_MAKING;REAGENT:A;[REAGENT:A:300:BAR:NO_SUBTYPE:METAL:IRON]}
	Line 28: {SHARED_OBJECT_REPLACE_TAG;REACTION:PIG_IRON_MAKING;PRODUCT:100;[PRODUCT:100:2:BAR:NO_SUBTYPE:METAL:PIG_IRON]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_tissue_template_MODEST.rbl (1 hit)
	Line 4: {SHARED_OBJECT_REPLACE_TAG;TISSUE_TEMPLATE:BONE_TEMPLATE;PAIN_RECEPTORS;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Kazoo\Silk egg\Libs\kazoos silk eggs.rbl (2 hits)
	Line 1: {SHARED_OBJECT_REPLACE_TAG;CREATURE:SPIDER_CAVE_GIANT;FEMALE;[FEMALE][LAYS_EGGS][EGG_MATERIAL:LOCAL_CREATURE_MAT:SILK:SOLID][EGG_MATERIAL:LOCAL_CREATURE_MAT:EGG_WHITE:LIQUID][EGG_MATERIAL:LOCAL_CREATURE_MAT:EGG_YOLK:LIQUID][EGG_SIZE:20000][CLUTCH_SIZE:2:3]}
	Line 2: {SHARED_OBJECT_REPLACE_TAG;MATERIAL_TEMPLATE:SILK_TEMPLATE;ITEMS_SOFT;[ITEMS_SOFT][STOCKPILE_GLOB][DO_NOT_CLEAN_GLOB][REACTION_CLASS:SILK_EGG]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\aaa_templates_flbase.rbl (1 hit)
	Line 5: {SHARED_OBJECT_REPLACE_TAG;CREATURE:LEOPARD;PET_EXOTIC;[COMMON_DOMESTIC][PET]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Uncle spam\Base\leathery vegetables\Uncle span.rbl (1 hit)
	Line 5: {SHARED_OBJECT_REPLACE_TAG;REACTION:TAN_A_HIDE;REAGENT:A;[REAGENT:A:1:NONE:NONE:NONE:NONE]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Wannabehero\Deep dwarven domestication\Base\_Deep dwarven domestication.rbl (1 hit)
	Line 2: {SHARED_OBJECT_REPLACE_TAG;REACTION:MAKE_MEAD;REAGENT:honey;[REAGENT:honey:150:LIQUID_MISC:NONE:CREATURE_MAT:NONE:HONEY][HAS_MATERIAL_REACTION_PRODUCT:DRINK_MAT]generic honey}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\IntensiveSprouting\plant_zISS.txt (1 hit)
	Line 5: 	{SHARED_OBJECT_REPLACE_TAG;PLANT:{@STR_TO_ID;%{object}_%{variety}};%{}{key};%{}{val}}}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\ISS\plant_ISS.txt (6 hits)
	Line 3: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};NAME;[NAME:%{name} %{variety}]}
	Line 4: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};NAME_PLURAL;[NAME_PLURAL:%{name}s %{variety}]}
	Line 5: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};ADJ:[ADJ:%{name}] %{variety}}
	Line 6: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};SEED;[SEED:%{name} %{variety} %{seed}:%{name} %{variety} %{seed}s:%{color}:LOCAL_PLANT_MAT:SEED]}
	Line 7: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};GROWDUR;[GROWDUR:%{growdur}]}  
	Line 8: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};CLUSTERSIZE;[CLUSTERSIZE:%{clustersize}]}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\ISS2\plant_ISS.txt (6 hits)
	Line 4: {SHARED_OBJECT_REPLACE_TAG;{@STR_TO_ID;%{object}_%{variety}};NAME;[NAME:%{name} %{variety}]}
	Line 5: {SHARED_OBJECT_REPLACE_TAG;{@STR_TO_ID;%{object}_%{variety}};NAME_PLURAL;[NAME_PLURAL:%{name}s %{variety}]}
	Line 6: {SHARED_OBJECT_REPLACE_TAG;{@STR_TO_ID;%{object}_%{variety}};ADJ:[ADJ:%{name} %{variety}]}
	Line 7: {SHARED_OBJECT_REPLACE_TAG;{@STR_TO_ID;%{object}_%{variety}};SEED;[SEED:%{name} %{variety} %{seed}:%{name} %{variety} %{seed}s:%{color}:LOCAL_PLANT_MAT:SEED]}
	Line 8: {SHARED_OBJECT_REPLACE_TAG;{@STR_TO_ID;%{object}_%{variety}};GROWDUR;[GROWDUR:%{growdur}]}  
	Line 9: {SHARED_OBJECT_REPLACE_TAG;{@STR_TO_ID;%{object}_%{variety}};CLUSTERSIZE;[CLUSTERSIZE:%{clustersize}]}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test01\plant_test.txt (3 hits)
	Line 62: {SHARED_OBJECT_REPLACE_TAG;PLANT:MUSHROOM_HELMET_PLUMP_SPROUTING;SEED;[SEED:plump helmet sproutings spawn:plump helmet sproutings spawn:4:0:1:LOCAL_PLANT_MAT:SEED]}
	Line 63: {SHARED_OBJECT_REPLACE_TAG;PLANT:MUSHROOM_HELMET_PLUMP_SPROUTING;GROWDUR;[GROWDUR:60]} divide by 5 # divide by 5, good for exp
	Line 64: {SHARED_OBJECT_REPLACE_TAG;PLANT:MUSHROOM_HELMET_PLUMP_SPROUTING;CLUSTERSIZE;[CLUSTERSIZE:1]} divide by 5, yearly output more or less maintained, except for the seeds
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test02\plant_test2.txt (3 hits)
	Line 3: {SHARED_OBJECT_REPLACE_TAG;PLANT:MUSHROOM_HELMET_PLUMP_INTENSIVE;GROWDUR;[GROWDUR:3900]}   # multiply by 13, good for FPS (bigger stacks)
	Line 4: {SHARED_OBJECT_REPLACE_TAG;PLANT:MUSHROOM_HELMET_PLUMP_INTENSIVE;CLUSTERSIZE;[CLUSTERSIZE:65]} # m
	Line 6: {SHARED_OBJECT_REPLACE_TAG;PLANT:MUSHROOM_HELMET_PLUMP_SPROUTING;SEED;[SEED:plump helmet intensive spawn:plump helmet intensive spawn:4:0:1:LOCAL_PLANT_MAT:SEED]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test03\plant_test3.txt (6 hits)
	Line 3: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_{@STR_UPPER;%{variety}};NAME;[NAME:%{name} %{variety}]}
	Line 4: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_{@STR_UPPER;%{variety}};NAME_PLURAL;[NAME_PLURAL:%{name}s %{variety}]}
	Line 5: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_{@STR_UPPER;%{variety}};ADJ:[ADJ:%{name}] %{variety}}
	Line 6: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_{@STR_UPPER;%{variety}};SEED;[SEED:%{name} %{variety} %{seed}:%{name} %{variety} %{seed}s:%{color}:LOCAL_PLANT_MAT:SEED]}
	Line 7: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_{@STR_UPPER;%{variety}};GROWDUR;[GROWDUR:%{growdur}]}  
	Line 8: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_{@STR_UPPER;%{variety}};CLUSTERSIZE;[CLUSTERSIZE:%{clustersize}]}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test04\plant_test4.txt (6 hits)
	Line 4: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};NAME;[NAME:%{name}]}
	Line 5: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};NAME_PLURAL;[NAME_PLURAL:%{name}s]}
	Line 6: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};ADJ:[ADJ:%{name}]}
	Line 7: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};SEED;[SEED:%{name} %{variety} %{seed}:%{name} %{variety} %{seed}s:%{color}:LOCAL_PLANT_MAT:SEED]}
	Line 8: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};[GROWDUR:%{growdur}]}  
	Line 9: {SHARED_OBJECT_REPLACE_TAG;PLANT:%{object}_%{variety};CLUSTERSIZE;[CLUSTERSIZE:%{clustersize}]}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test05\plant_test5.txt (6 hits)
	Line 4: {SHARED_OBJECT_REPLACE_TAG;%{object}_{@STR_UPPER;%{variety}};NAME;[NAME:%{name} %{variety}]}
	Line 5: {SHARED_OBJECT_REPLACE_TAG;%{object}_{@STR_UPPER;%{variety}};NAME_PLURAL;[NAME_PLURAL:%{name}s %{variety}]}
	Line 6: {SHARED_OBJECT_REPLACE_TAG;%{object}_{@STR_UPPER;%{variety}};ADJ:[ADJ:%{name} %{variety}]}
	Line 7: {SHARED_OBJECT_REPLACE_TAG;%{object}_{@STR_UPPER;%{variety}};SEED;[SEED:%{name} %{variety} %{seed}:%{name} %{variety} %{seed}s:%{color}:LOCAL_PLANT_MAT:SEED]}
	Line 8: {SHARED_OBJECT_REPLACE_TAG;%{object}_{@STR_UPPER;%{variety}};GROWDUR;[GROWDUR:%{growdur}]}  
	Line 9: {SHARED_OBJECT_REPLACE_TAG;%{object}_{@STR_UPPER;%{variety}};CLUSTERSIZE;[CLUSTERSIZE:%{clustersize}]}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test06\entity_test.txt (2 hits)
	Line 760: {!TEMPLATE;SET_VALUE;token;value;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 761: {!TEMPLATE;REPLACE_TAG;token;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}]%{preraws}}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test07\entity_test.txt (2 hits)
	Line 15: {!TEMPLATE;SET_VALUE;token;value;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 16: {!TEMPLATE;ATTACH_TO_TAG;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;$TOKEN_CURSOR;[$TOKEN_CURSOR]%{preraws}}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test09\test.rbl (3 hits)
	Line 1: {SHARED_OBJECT_REPLACE_TAG;REACTION:TAN_A_HIDE;REAGENT:A;
	Line 7: {SHARED_OBJECT_REPLACE_TAG;REACTION:PROCESS_PLANT_TO_BAG;PRODUCT:100:1:SEEDS;
	Line 11: {SHARED_OBJECT_REPLACE_TAG;REACTION:MAKE_PARCHMENT;REAGENT:A:1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test10\test.rbl (5 hits)
	Line 1: {SHARED_OBJECT_REPLACE_TAG;REACTION:TAN_A_HIDE;REAGENT:A;
	Line 7: {SHARED_OBJECT_REPLACE_TAG;REACTION:PROCESS_PLANT_TO_BAG;PRODUCT:100:1:SEEDS;
	Line 11: {SHARED_OBJECT_REPLACE_TAG;REACTION:MAKE_PARCHMENT;REAGENT:A:1;
	Line 18: {SHARED_OBJECT_REPLACE_TAG;REACTION:PIG_IRON_MAKING;REAGENT:A;[REAGENT:A:300:BAR:NO_SUBTYPE:METAL:IRON]}
	Line 19: {SHARED_OBJECT_REPLACE_TAG;REACTION:PIG_IRON_MAKING;PRODUCT:100;[PRODUCT:100:2:BAR:NO_SUBTYPE:METAL:PIG_IRON]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test11\test.rbl (3 hits)
	Line 1: {SHARED_OBJECT_REPLACE_TAG;REACTION:TAN_A_HIDE;REAGENT:A;
	Line 7: {SHARED_OBJECT_REPLACE_TAG;REACTION:PROCESS_PLANT_TO_BAG;PRODUCT:100:1:SEEDS;
	Line 11: {SHARED_OBJECT_REPLACE_TAG;REACTION:MAKE_PARCHMENT;REAGENT:A:1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test12\test.lua (3 hits)
	Line 1: {SHARED_OBJECT_REPLACE_TAG;REACTION:TAN_A_HIDE;REAGENT:A;
	Line 7: {SHARED_OBJECT_REPLACE_TAG;REACTION:PROCESS_PLANT_TO_BAG;PRODUCT:100:1:SEEDS;
	Line 11: {SHARED_OBJECT_REPLACE_TAG;REACTION:MAKE_PARCHMENT;REAGENT:A:1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test13\test.rbl (3 hits)
	Line 1: {SHARED_OBJECT_REPLACE_TAG;REACTION:TAN_A_HIDE;REAGENT:A;
	Line 7: {SHARED_OBJECT_REPLACE_TAG;REACTION:PROCESS_PLANT_TO_BAG;PRODUCT:100:1:SEEDS;
	Line 11: {SHARED_OBJECT_REPLACE_TAG;REACTION:MAKE_PARCHMENT;REAGENT:A:1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test14\test.rbl (3 hits)
	Line 1: {SHARED_OBJECT_REPLACE_TAG;REACTION:TAN_A_HIDE;REAGENT:A;
	Line 7: {SHARED_OBJECT_REPLACE_TAG;REACTION:PROCESS_PLANT_TO_BAG;PRODUCT:100:1:SEEDS;
	Line 11: {SHARED_OBJECT_REPLACE_TAG;REACTION:MAKE_PARCHMENT;REAGENT:A:1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test15\test.rbl (3 hits)
	Line 1: {SHARED_OBJECT_REPLACE_TAG;REACTION:TAN_A_HIDE;REAGENT:A;
	Line 7: {SHARED_OBJECT_REPLACE_TAG;REACTION:PROCESS_PLANT_TO_BAG;PRODUCT:100:1:SEEDS;
	Line 11: {SHARED_OBJECT_REPLACE_TAG;REACTION:MAKE_PARCHMENT;REAGENT:A:1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Base\Haiku\Extreme\Base\Extreme Haiku.rbl (1 hit)
	Line 117: 	{SHARED_OBJECT_REPLACE_TAG;ENTITY:EVIL;%{key}:GOBLIN;[%{key}:DWARF]}}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (2 hits)
	Line 326: 	{SHARED_OBJECT_REPLACE_TAG;<ID>;<TAG>;<REPLACEMENT>}
	Line 341: 	{SHARED_OBJECT_REPLACE_TAG;SLADE;UNDIGGABLE;[AQUIFER]}
  Z:\Dwarf Fortress\nanofort\Rubble\other\tutorial.md (1 hit)
	Line 392: {SHARED_OBJECT_REPLACE_TAG;<ID>;<TAG>;<REPLACEMENT>}
{!SHARED_OBJECT_DUPLICATE;<OLD_ID>;<NEW_ID>;[<EDIT_RAWS>=true];[<ADD_CATEGORY>=true]}
 Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_sobject_libs_base.test (3 hits)
	Line 5: {!SHARED_OBJECT_DUPLICATE;TEST_SOD_A;TEST_SOD_A2;false}
	Line 15: {!SHARED_OBJECT_DUPLICATE;INORGANIC:TEST_SOD_B;INORGANIC:TEST_SOD_B2}
	Line 27: {!SHARED_OBJECT_DUPLICATE;INORGANIC:TEST_SOD_C;INORGANIC:TEST_SOD_C2}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Intensive Sprouting\Libs\plant_zISS.txt (12 hits)
	Line 11: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:{@STR_TO_ID;MUSHROOM_HELMET_PLUMP_sprouting};true;true}
	Line 12: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:{@STR_TO_ID;MUSHROOM_HELMET_PLUMP_intensive};true;true}
	Line 13: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:{@STR_TO_ID;GRASS_TAIL_PIG_sprouting};true;true}
	Line 14: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:{@STR_TO_ID;GRASS_TAIL_PIG_intensive};true;true}
	Line 15: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:{@STR_TO_ID;GRASS_WHEAT_CAVE_sprouting};true;true}
	Line 16: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:{@STR_TO_ID;GRASS_WHEAT_CAVE_intensive};true;true}
	Line 17: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:{@STR_TO_ID;POD_SWEET_sprouting};true;true}
	Line 18: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:{@STR_TO_ID;POD_SWEET_intensive};true;true}
	Line 19: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:{@STR_TO_ID;BUSH_QUARRY_sprouting};true;true}
	Line 20: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:{@STR_TO_ID;BUSH_QUARRY_intensive};true;true}
	Line 21: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:{@STR_TO_ID;MUSHROOM_CUP_DIMPLE_sprouting};true;true}
	Line 22: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:{@STR_TO_ID;MUSHROOM_CUP_DIMPLE_intensive};true;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Intensive Sprouting\Libs\Z\Intensive sprouting2\plant_ZISS.txt (12 hits)
	Line 10: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_sprouting;true;true}
	Line 11: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_intensive;true;true}
	Line 16: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:GRASS_TAIL_PIG_sprouting;true;true}
	Line 17: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:GRASS_TAIL_PIG_intensive;true;true}
	Line 21: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:GRASS_WHEAT_CAVE_sprouting;true;true}
	Line 22: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:GRASS_WHEAT_CAVE_intensive;true;true}
	Line 26: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:POD_SWEET_sprouting;true;true}
	Line 27: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:POD_SWEET_intensive;true;true}
	Line 31: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:BUSH_QUARRY_sproutingtrue;true}
	Line 32: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:BUSH_QUARRY_intensive;true;true}
	Line 36: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:MUSHROOM_CUP_DIMPLE_sprouting;true;true}
	Line 37: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:MUSHROOM_CUP_DIMPLE_intensive;true;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Intensive Sprouting\Libs\Z\Intensive sprouting3\plant_ZISS.txt (12 hits)
	Line 10: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_sprouting;true;true}
	Line 11: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_intensive;true;true}
	Line 16: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:GRASS_TAIL_PIG_sprouting;true;true}
	Line 17: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:GRASS_TAIL_PIG_intensive;true;true}
	Line 21: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:GRASS_WHEAT_CAVE_sprouting;true;true}
	Line 22: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:GRASS_WHEAT_CAVE_intensive;true;true}
	Line 26: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:POD_SWEET_sprouting;true;true}
	Line 27: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:POD_SWEET_intensive;true;true}
	Line 31: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:BUSH_QUARRY_sproutingtrue;true}
	Line 32: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:BUSH_QUARRY_intensive;true;true}
	Line 36: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:MUSHROOM_CUP_DIMPLE_sprouting;true;true}
	Line 37: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:MUSHROOM_CUP_DIMPLE_intensive;true;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\IntensiveSprouting\plant_zISS.txt (14 hits)
	Line 9: 	{!SHARED_OBJECT_DUPLICATE;PLANT:%{key};PLANT:{@STR_TO_ID;%{object}_%{variety}};true;true}
	Line 13: 	{!SHARED_OBJECT_DUPLICATE;PLANT:%{key};PLANT:{@STR_TO_ID;%{object}_%{variety}};true;true}
	Line 20: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_sprouting;true;true}
	Line 21: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_intensive;true;true}
	Line 22: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:GRASS_TAIL_PIG_sprouting;true;true}
	Line 23: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:GRASS_TAIL_PIG_intensive;true;true}
	Line 24: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:GRASS_WHEAT_CAVE_sprouting;true;true}
	Line 25: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:GRASS_WHEAT_CAVE_intensive;true;true}
	Line 26: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:POD_SWEET_sprouting;true;true}
	Line 27: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:POD_SWEET_intensive;true;true}
	Line 28: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:BUSH_QUARRY_sprouting;true;true}
	Line 29: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:BUSH_QUARRY_intensive;true;true}
	Line 30: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:MUSHROOM_CUP_DIMPLE_sprouting;true;true}
	Line 31: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:MUSHROOM_CUP_DIMPLE_intensive;true;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\ISS\plant_ISS.txt (12 hits)
	Line 10: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_sprouting;true;true}
	Line 11: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_intensive;true;true}
	Line 16: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:GRASS_TAIL_PIG_sprouting;true;true}
	Line 17: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:GRASS_TAIL_PIG_intensive;true;true}
	Line 21: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:GRASS_WHEAT_CAVE_sprouting;true;true}
	Line 22: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:GRASS_WHEAT_CAVE_intensive;true;true}
	Line 26: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:POD_SWEET_sprouting;true;true}
	Line 27: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:POD_SWEET_intensive;true;true}
	Line 31: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:BUSH_QUARRY_sproutingtrue;true}
	Line 32: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:BUSH_QUARRY_intensive;true;true}
	Line 36: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:MUSHROOM_CUP_DIMPLE_sprouting;true;true}
	Line 37: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:MUSHROOM_CUP_DIMPLE_intensive;true;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\ISS2\plant_ISS.txt (1 hit)
	Line 3: {!SHARED_OBJECT_DUPLICATE;PLANT:%{object};PLANT:{@STR_TO_ID;%{object}_%{variety}};true;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test02\plant_test2.txt (1 hit)
	Line 2: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_INTENSIVE;true;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test03\plant_test3.txt (12 hits)
	Line 10: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_SPROUTING;true;true}
	Line 11: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_HELMET_PLUMP;PLANT:MUSHROOM_HELMET_PLUMP_INTENSIVE;true;true}
	Line 16: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:GRASS_TAIL_PIG_SPROUTING;true;true}
	Line 17: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_TAIL_PIG;PLANT:GRASS_TAIL_PIG_INTENSIVE;true;true}
	Line 21: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:GRASS_WHEAT_CAVE_SPROUTING;true;true}
	Line 22: {!SHARED_OBJECT_DUPLICATE;PLANT:GRASS_WHEAT_CAVE;PLANT:GRASS_WHEAT_CAVE_INTENSIVE;true;true}
	Line 26: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:POD_SWEET_SPROUTING;true;true}
	Line 27: {!SHARED_OBJECT_DUPLICATE;PLANT:POD_SWEET;PLANT:POD_SWEET_INTENSIVE;true;true}
	Line 31: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:BUSH_QUARRY_SPROUTING;true;true}
	Line 32: {!SHARED_OBJECT_DUPLICATE;PLANT:BUSH_QUARRY;PLANT:BUSH_QUARRY_INTENSIVE;true;true}
	Line 36: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:MUSHROOM_CUP_DIMPLE_SPROUTING;true;true}
	Line 37: {!SHARED_OBJECT_DUPLICATE;PLANT:MUSHROOM_CUP_DIMPLE;PLANT:MUSHROOM_CUP_DIMPLE_INTENSIVE;true;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test04\plant_test4.txt (1 hit)
	Line 3: {!SHARED_OBJECT_DUPLICATE;PLANT:%{object};PLANT:%{object}_%{variety};true;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test05\plant_test5.txt (1 hit)
	Line 3: {!SHARED_OBJECT_DUPLICATE;PLANT:%{object};PLANT:%{object}_{@STR_UPPER;%{variety}};true;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test16\zzz___test.txt (1 hit)
	Line 3: 	{!SHARED_OBJECT_DUPLICATE;MATERIAL_TEMPLATE:%{key}_TEMPLATE;MATERIAL_TEMPLATE:{@STR_TO_ID;%{key}_GENERIC_TEMPLATE};true;true}}	
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (1 hit)
	Line 345: 	{!SHARED_OBJECT_DUPLICATE;<OLD_ID>;<NEW_ID>;[<EDIT_RAWS>=true];[<ADD_CATEGORY>=true]}
  Z:\Dwarf Fortress\nanofort\Rubble\other\tutorial.md (1 hit)
	Line 544: {!SHARED_OBJECT_DUPLICATE;<OLD_ID>;<NEW_ID>;[<EDIT_RAWS>=true];[<ADD_CATEGORY>=true]}
{SHARED_OBJECT_MERGE;<ID>;<RULES>;<SOURCE>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_sobject_libs_base.test (2 hits)
	Line 58: {SHARED_OBJECT_MERGE;TEST_SOM_A;$:&:?(1,-1);[TEST:M:N:O:P]}
	Line 65: {SHARED_OBJECT_MERGE;TEST_SOM_B;$:&:?(1,-1);[TEST:M:N:O]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test13\test.rbl (2 hits)
	Line 19: {SHARED_OBJECT_MERGE;REACTION:PIG_IRON_MAKING;$:$:?(1,-1);
	Line 22: {SHARED_OBJECT_MERGE;REACTION:STEEL_MAKING;$:$:?(1,-1);
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (2 hits)
	Line 372: 	{SHARED_OBJECT_MERGE;<ID>;<RULES>;<SOURCE>}
	Line 379: 	{SHARED_OBJECT_MERGE;TEST;$:?:&;[FOO:BAZ:BAR]}
  Z:\Dwarf Fortress\nanofort\Rubble\other\tutorial.md (1 hit)
	Line 643: {SHARED_OBJECT_MERGE;<ID>;<RULES>;<SOURCE>}
{BUILDING_WORKSHOP;<ID>;<CLASS>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 270: 	"{BUILDING_WORKSHOP;%{id};%{class}}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Powered\libs_powered.mod.lua (1 hit)
	Line 185: 			out = out.."\n{BUILDING_WORKSHOP;"..id..outputid..";"..hook..[[}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\ThirdParty\Steam Engine\building_user_steam_engine.txt (2 hits)
	Line 3: {BUILDING_WORKSHOP;STEAM_ENGINE;ADDON_HOOK_PLAYABLE}
	Line 47: {BUILDING_WORKSHOP;MAGMA_STEAM_ENGINE;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Powered\Geothermal Plant\building_x_user_powered_geothermal_plant.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_GEOTHERMAL_PLANT;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\building_TESB.txt (1 hit)
	Line 5: {BUILDING_WORKSHOP;TESB_TRIBUTE;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Assembler\building_x_first_landing_powered_assembler.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOASSEMBLER;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Decorator\building_x_first_landing_powered_decorator.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTODECORATOR;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Drill\building_x_first_landing_powered_drill.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_DRILLING_RIG;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Furnace\building_x_first_landing_powered_furnace.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOFURNACE;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Geothermal Plant\building_x_first_landing_powered_geothermal_plant.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_GEOTHERMAL_PLANT;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Loom\building_x_first_landing_powered_loom.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOLOOM;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Mill\building_x_first_landing_powered_mill.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOMILL;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Quarry\building_x_first_landing_powered_quarry.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOQUARRY;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Rock Breaker\building_x_first_landing_powered_rock_breaker.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOROCK_BREAKER;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Sawmill\building_x_first_landing_powered_sawmill.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOSAWMILL;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Screw Press\building_x_first_landing_powered_screw_press.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOSCREWPRESS;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Spinner\building_x_first_landing_powered_spinner.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOSPINNER;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Still\building_x_first_landing_powered_still.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOSTILL;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Tailor\building_x_first_landing_powered_tailor.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOTAILOR;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Unraveler\building_x_first_landing_powered_unraveler.txt (1 hit)
	Line 3: {BUILDING_WORKSHOP;DFHACK_POWERED_FL_AUTOUNRAVELER;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\Toolbox\Dummy\WShops\dev_dummy_wshops.pre.lua (1 hit)
	Line 15: 	out = out.."\n{BUILDING_WORKSHOP;DUMMY_WSHOP_"..count..[[;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test14\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 286: 	"{BUILDING_WORKSHOP;%{id};%{class}}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Libs\Delta\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 286: 	"{BUILDING_WORKSHOP;%{id};%{class}}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (1 hit)
	Line 389: 	{BUILDING_WORKSHOP;<ID>;<CLASS>}
  Z:\Dwarf Fortress\nanofort\Rubble\other\tutorial.md (1 hit)
	Line 655: {BUILDING_WORKSHOP;<ID>;<CLASS>}
{BUILDING_FURNACE;<ID>;<CLASS>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 277: 	"{BUILDING_FURNACE;%{id};%{class}}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Display Case\building_user_decorations_display_case.txt (1 hit)
	Line 3: {BUILDING_FURNACE;DISPLAY_CASE;ADDON_HOOK_PLAYABLE}	
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\Toolbox\Dummy\WShops\dev_dummy_wshops.pre.lua (1 hit)
	Line 30: 	out = out.."\n{BUILDING_FURNACE;DUMMY_FURNACE_"..count..[[;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test14\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 293: 	"{BUILDING_FURNACE;%{id};%{class}}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Libs\Delta\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 293: 	"{BUILDING_FURNACE;%{id};%{class}}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (1 hit)
	Line 399: 	{BUILDING_FURNACE;<ID>;<CLASS>}
  Z:\Dwarf Fortress\nanofort\Rubble\other\tutorial.md (1 hit)
	Line 707: {BUILDING_FURNACE;<ID>;<CLASS>}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Tutorials\Common standard library templates.md (1 hit)
	Line 20: 	{BUILDING_FURNACE;BONFIRE;ADDON_HOOK_PLAYABLE}
{#USES_BUILDINGS;<CLASS>}
{REACTION;<ID>;<CLASS>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 263: 	"{REACTION;%{id};%{class}}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Crates\examples_libs_crates.rbl (1 hit)
	Line 45: {REACTION;EXAMPLE_WORLDGEN;ADDON_HOOK_GENERIC}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Crates\libs_crates.pre.lua (4 hits)
	Line 98: 			"\n{REACTION;UNPACK_"..id..";"..techclass.."}\n"..
	Line 108: 			"\n{REACTION;UNPACK_"..id..";"..techclass.."}\n"..
	Line 143: 			"\n{REACTION;UNPACK_"..id..";"..techclass.."}\n"..
	Line 153: 			"\n{REACTION;UNPACK_"..id..";"..techclass.."}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Circular kitchen\Libs\reaction_dried_meals.txt (1 hit)
	Line 2: {REACTION;DRY_MEALS;ADDON_HOOK_MOUNTAIN}{REACTION_ADD_CLASS;ADDON_HOOK_EVIL}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Block\Boulder\reaction_user_block_boulder.txt (2 hits)
	Line 3: {REACTION;PRESS_BOULDER_%{color};ADDON_HOOK_PLAYABLE}
	Line 12: {REACTION;PRESS_ORNAMENTAL_BOULDER_%{color};ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Block\Cutter\reaction_user_block_cutter.txt (2 hits)
	Line 3: {REACTION;CUT_BLOCK_%{color};ADDON_HOOK_PLAYABLE}
	Line 11: {REACTION;CUT_ORNAMENTAL_BLOCK_%{color};ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Block\Kiln\reaction_user_block_kiln.txt (2 hits)
	Line 3: {REACTION;FIRE_BLOCK_%{color};ADDON_HOOK_PLAYABLE}
	Line 13: {REACTION;FIRE_ORNAMENTAL_BLOCK_%{color};ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Bonecarver\reaction_user_bonecarver.txt (4 hits)
	Line 3: {REACTION;CARVE_BONE_%id;ADDON_HOOK_PLAYABLE}
	Line 13: {REACTION;CARVE_FAKE_BONE_%id;ADDON_HOOK_PLAYABLE}
	Line 29: {REACTION;CARVE_BONE_BLOCK_TREATED;ADDON_HOOK_PLAYABLE}
	Line 49: {REACTION;CARVE_BONE_CRAFTS;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Bonfire\reaction_user_bonfire.txt (2 hits)
	Line 3: {REACTION;BONFIRE_BIG_START;ADDON_HOOK_PLAYABLE}
	Line 12: {REACTION;BONFIRE_SMALL_START;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Edible Vermin\reaction_user_edible_vermin.txt (1 hit)
	Line 3: {REACTION;COOK_VERMIN;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Lamination\reaction_user_lamination.txt (4 hits)
	Line 3: {REACTION;MAKE_HIDE_GLUE;ADDON_HOOK_PLAYABLE}
	Line 11: {REACTION;MAKE_BONE_GLUE;ADDON_HOOK_PLAYABLE}
	Line 21: {REACTION;MAKE_HORN_GLUE;ADDON_HOOK_PLAYABLE}
	Line 31: {REACTION;MAKE_LAMINATED_LOG_PLANKS;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Sawmill\reaction_basics_user_sawmill.txt (3 hits)
	Line 3: {REACTION;CUT_WOOD_SAWMILL;ADDON_HOOK_PLAYABLE}
	Line 11: {REACTION;CHARCOAL_SAWMILL;ADDON_HOOK_PLAYABLE}
	Line 20: {REACTION;ASH_SAWMILL;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Sawmill\reaction_scrap_user_sawmill.txt (1 hit)
	Line 3: {REACTION;SCRAP_WOODEN_%item;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Shovel\reaction_user_shovel.txt (2 hits)
	Line 3: {REACTION;SHOVEL_WOOD_BREAKABLE;ADDON_HOOK_PLAYABLE}
	Line 10: {REACTION;SHOVEL_BONE_BREAKABLE;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Stone Anvil\reaction_user_stone_anvil.txt (1 hit)
	Line 3: {REACTION;ANVIL_STONE_BREAKABLE;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Wooden Mechanisms\reaction_user_wooden_mechanisms.txt (1 hit)
	Line 3: {REACTION;MECHANISMS_WOOD_GENERIC;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\ThirdParty\Spatter\reaction_user_dfhack_spatter.txt (3 hits)
	Line 6: {REACTION;SPATTER_ADD_OBJECT_LIQUID;ADDON_HOOK_PLAYABLE}
	Line 25: {REACTION;SPATTER_ADD_WEAPON_EXTRACT;ADDON_HOOK_PLAYABLE}
	Line 45: {REACTION;SPATTER_ADD_AMMO_EXTRACT;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\ThirdParty\Steam Engine\reaction_user_steam_engine.txt (1 hit)
	Line 3: {REACTION;STOKE_BOILER;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Archery\reaction_user_archery.txt (11 hits)
	Line 3: {REACTION;ARCHERY_WOOD_%{type};ADDON_HOOK_PLAYABLE}
	Line 11: {REACTION;ARCHERY_BONE_%{type};ADDON_HOOK_PLAYABLE}
	Line 21: {REACTION;ARCHERY_%{type}_AMMO;ADDON_HOOK_PLAYABLE}
	Line 40: {REACTION;ARCHERY_WOOD_SHAFTS;ADDON_HOOK_PLAYABLE}
	Line 48: {REACTION;ARCHERY_BONE_SHAFTS;ADDON_HOOK_PLAYABLE}
	Line 58: {REACTION;ARCHERY_COPPER_SHAFTS;ADDON_HOOK_PLAYABLE}
	Line 66: {REACTION;ARCHERY_BRONZE_SHAFTS;ADDON_HOOK_PLAYABLE}
	Line 74: {REACTION;ARCHERY_BISMUTH_BRONZE_SHAFTS;ADDON_HOOK_PLAYABLE}
	Line 82: {REACTION;ARCHERY_SILVER_SHAFTS;ADDON_HOOK_PLAYABLE}
	Line 90: {REACTION;ARCHERY_IRON_SHAFTS;ADDON_HOOK_PLAYABLE}
	Line 98: {REACTION;ARCHERY_STEEL_SHAFTS;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Bone Flux\reaction_user_bone_flux.txt (1 hit)
	Line 3: {REACTION;FLUX_BONE;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Cast Anvil\reaction_user_cast_anvil.txt (1 hit)
	Line 3: {REACTION;SMELTER_IRON_ANVIL;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Glass Forge\reaction_user_glass_forge.txt (9 hits)
	Line 9: {REACTION;GREEN_GLASS_%id;%hook}
	Line 20: {REACTION;CLEAR_GLASS_%id;%hook}
	Line 30: {REACTION;CRYSTAL_GLASS_%id;%hook}
	Line 126: {REACTION;GREEN_GLASS_BAR_TO_ROUGH;ADDON_HOOK_PLAYABLE}
	Line 135: {REACTION;CLEAR_GLASS_BAR_TO_ROUGH;ADDON_HOOK_PLAYABLE}
	Line 144: {REACTION;CRYSTAL_GLASS_BAR_TO_ROUGH;ADDON_HOOK_PLAYABLE}
	Line 153: {REACTION;GREEN_GLASS_ROUGH_TO_BAR;ADDON_HOOK_PLAYABLE}
	Line 163: {REACTION;CLEAR_GLASS_ROUGH_TO_BAR;ADDON_HOOK_PLAYABLE}
	Line 173: {REACTION;CRYSTAL_GLASS_ROUGH_TO_BAR;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Grind Sand\reaction_user_grind_sand.txt (1 hit)
	Line 3: {REACTION;GRIND_STONE_TO_SAND;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Melt Item\reaction_user_melt_item.txt (1 hit)
	Line 3: {REACTION;SMELTER_MELT_METAL_ITEM;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Metallurgy\Aluminum Bronze\reaction_user_metallurgy_aluminum_bronze.txt (1 hit)
	Line 3: {REACTION;ALLOY_ALUMINUM_BRONZE;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Metallurgy\Bloomery\reaction_user_metallurgy_bloomery.txt (9 hits)
	Line 3: {REACTION;BLOOM_WORLDGEN;ADDON_HOOK_PLAYABLE}
	Line 11: {REACTION;BLOOM_IRON_START;ADDON_HOOK_PLAYABLE}
	Line 23: {REACTION;BLOOM_IRON_START_FLUX;ADDON_HOOK_PLAYABLE}
	Line 37: {REACTION;BLOOM_STEEL_START;ADDON_HOOK_PLAYABLE}
	Line 49: {REACTION;BLOOM_STEEL_START_FLUX;ADDON_HOOK_PLAYABLE}
	Line 64: {REACTION;BLOOM_HAMMER_1;ADDON_HOOK_PLAYABLE}
	Line 76: {REACTION;BLOOM_HAMMER_2;ADDON_HOOK_PLAYABLE}
	Line 88: {REACTION;BLOOM_HAMMER_3;ADDON_HOOK_PLAYABLE}
	Line 100: {REACTION;BLOOM_HAMMER_4;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Metallurgy\Mithril\reaction_user_metallurgy_mithril.txt (1 hit)
	Line 3: {REACTION;ALLOY_MITHRIL;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Metallurgy\reaction_aaa_user_metallurgy.txt (14 hits)
	Line 3: {REACTION;ALLOY_PIG_IRON;ADDON_HOOK_PLAYABLE}
	Line 16: {REACTION;ALLOY_STEEL;ADDON_HOOK_PLAYABLE}
	Line 30: {REACTION;ALLOY_BRASS;ADDON_HOOK_PLAYABLE}
	Line 41: {REACTION;ALLOY_BRONZE;ADDON_HOOK_PLAYABLE}
	Line 52: {REACTION;ALLOY_ELECTRUM;ADDON_HOOK_PLAYABLE}
	Line 63: {REACTION;ALLOY_BILLON;ADDON_HOOK_PLAYABLE}
	Line 74: {REACTION;ALLOY_PEWTER_FINE;ADDON_HOOK_PLAYABLE}
	Line 85: {REACTION;ALLOY_PEWTER_TRIFLE;ADDON_HOOK_PLAYABLE}
	Line 96: {REACTION;ALLOY_PEWTER_LAY;ADDON_HOOK_PLAYABLE}
	Line 108: {REACTION;ALLOY_NICKEL_SILVER;ADDON_HOOK_PLAYABLE}
	Line 120: {REACTION;ALLOY_BLACK_BRONZE;ADDON_HOOK_PLAYABLE}
	Line 132: {REACTION;ALLOY_STERLING_SILVER;ADDON_HOOK_PLAYABLE}
	Line 143: {REACTION;ALLOY_ROSE_GOLD;ADDON_HOOK_PLAYABLE}
	Line 154: {REACTION;ALLOY_BISMUTH_BRONZE;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Metallurgy\Smelter\reaction_aaa_user_metallurgy_smelter.txt (2 hits)
	Line 3: {REACTION;SMELT_ORE_%{metal};ADDON_HOOK_PLAYABLE}
	Line 22: {REACTION;SMELT_ORE_SILVER_POOR;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Pottery\Batch\reaction_user_pottery_batch.txt (2 hits)
	Line 3: {REACTION;BATCH_CLAY_%id;ADDON_HOOK_PLAYABLE}
	Line 31: {REACTION;BATCH_CLAY_CRAFTS;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Pottery\Glaze\reaction_user_pottery_glaze.txt (10 hits)
	Line 3: {REACTION;GRIND_GLAZE_ASH;ADDON_HOOK_PLAYABLE}
	Line 18: {REACTION;GRIND_GLAZE_WHITE;ADDON_HOOK_PLAYABLE}
	Line 34: {REACTION;GRIND_GLAZE_RED;ADDON_HOOK_PLAYABLE}
	Line 50: {REACTION;GRIND_GLAZE_YELLOW;ADDON_HOOK_PLAYABLE}
	Line 66: {REACTION;GRIND_GLAZE_BLUE;ADDON_HOOK_PLAYABLE}
	Line 82: {REACTION;GRIND_GLAZE_ORE;ADDON_HOOK_PLAYABLE}
	Line 98: {REACTION;GRIND_GLAZE_MINERAL;ADDON_HOOK_PLAYABLE}
	Line 116: {REACTION;GLAZE_ITEM_POWDER_ONE;ADDON_HOOK_PLAYABLE}
	Line 135: {REACTION;GLAZE_ITEM_POWDER_TWO;ADDON_HOOK_PLAYABLE}
	Line 162: {REACTION;GLAZE_ITEM_POWDER_FOUR;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Pottery\reaction_user_pottery.txt (4 hits)
	Line 3: {REACTION;CLAY_%id;ADDON_HOOK_PLAYABLE}
	Line 31: {REACTION;CLAY_CRAFTS;ADDON_HOOK_PLAYABLE}
	Line 84: {REACTION;GLAZE_ITEM_BAR;ADDON_HOOK_PLAYABLE}
	Line 99: {REACTION;GLAZE_ITEM_BOULDER;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Quarry\reaction_user_quarry.txt (5 hits)
	Line 3: {REACTION;QUARRY_CLAY;ADDON_HOOK_PLAYABLE}
	Line 10: {REACTION;SMALL_QUARRY_CLAY;ADDON_HOOK_PLAYABLE}
	Line 16: {REACTION;QUARRY_SAND;ADDON_HOOK_PLAYABLE}
	Line 29: {REACTION;SMALL_QUARRY_SAND;ADDON_HOOK_PLAYABLE}
	Line 42: {REACTION;QUARRY_STONE;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Saurian\reaction_saurian.txt (4 hits)
	Line 3: {REACTION;RESIZE_PICK;ADDON_HOOK_SWAMP}
	Line 13: {REACTION;UNCUT_GEM_LARGE;ADDON_HOOK_SWAMP}
	Line 20: {REACTION;PRESS_BOULDER;ADDON_HOOK_SWAMP}
	Line 30: {REACTION;WORLDGEN_GLASS;ADDON_HOOK_SWAMP}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Warcrafter\reaction_user_warcrafter.txt (8 hits)
	Line 6: {REACTION;WARCRAFTER_LEATHER_%id;ADDON_HOOK_PLAYABLE}
	Line 27: {REACTION;WARCRAFTER_LEATHER_ARMOR_SET;ADDON_HOOK_PLAYABLE}
	Line 42: {REACTION;WARCRAFTER_SHELL_SCALES;ADDON_HOOK_PLAYABLE}
	Line 50: {REACTION;WARCRAFTER_BONE_SCALES;ADDON_HOOK_PLAYABLE}
	Line 63: {REACTION;WARCRAFTER_SCALE_%id;ADDON_HOOK_PLAYABLE}
	Line 82: {REACTION;WARCRAFTER_SCALE_ARMOR_SET;ADDON_HOOK_PLAYABLE}
	Line 102: {REACTION;CLEAN_SHELL_BARS;ADDON_HOOK_PLAYABLE}
	Line 108: {REACTION;CLEAN_BONE_BARS;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Wicker\reaction_user_wicker.txt (16 hits)
	Line 3: {REACTION;WEAVE_GRASS_LONGLAND;ADDON_HOOK_PLAYABLE}
	Line 12: {REACTION;WEAVE_REED_ROPE;ADDON_HOOK_PLAYABLE}
	Line 21: {REACTION;WEAVE_GRASS_WHEAT_CAVE;ADDON_HOOK_PLAYABLE}
	Line 30: {REACTION;WEAVE_GRASS_TAIL_PIG;ADDON_HOOK_PLAYABLE}
	Line 39: {REACTION;WICKER_ARMORSTAND;ADDON_HOOK_PLAYABLE}
	Line 46: {REACTION;WICKER_BIN;ADDON_HOOK_PLAYABLE}
	Line 53: {REACTION;WICKER_CABINET;ADDON_HOOK_PLAYABLE}
	Line 60: {REACTION;WICKER_CAGE;ADDON_HOOK_PLAYABLE}
	Line 67: {REACTION;WICKER_CHEST;ADDON_HOOK_PLAYABLE}
	Line 74: {REACTION;WICKER_DOOR;ADDON_HOOK_PLAYABLE}
	Line 81: {REACTION;WICKER_GRATE;ADDON_HOOK_PLAYABLE}
	Line 88: {REACTION;WICKER_HATCH_COVER;ADDON_HOOK_PLAYABLE}
	Line 95: {REACTION;WICKER_NESTBOX;ADDON_HOOK_PLAYABLE}
	Line 102: {REACTION;WICKER_TABLE;ADDON_HOOK_PLAYABLE}
	Line 109: {REACTION;WICKER_THRONE;ADDON_HOOK_PLAYABLE}
	Line 116: {REACTION;WICKER_WEAPONRACK;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\reaction_TESB.txt (4 hits)
	Line 38: }{REACTION;TESB_TRIBUTE_TO_%{id};ADDON_HOOK_PLAYABLE}
	Line 47: }{REACTION;TESB_TRIBUTE_TO_%{id}2;ADDON_HOOK_PLAYABLE}
	Line 116: {REACTION;TESB_EXTRACT_SEED_%{id};ADDON_HOOK_PLAYABLE}
	Line 185: {REACTION;PROCESS_GEM_CLUSTER;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Lovechild\All races playable\Libs\reaction_arp.txt (17 hits)
	Line 7: {REACTION;MAKE_WOODEN_SWORD_SHORT_ARP;ADDON_HOOK_FOREST}
	Line 14: {REACTION;MAKE_WOODEN_SPEAR_ARP;ADDON_HOOK_FOREST}
	Line 21: {REACTION;MAKE_WOODEN_BREASTPLATE_ARP;ADDON_HOOK_FOREST}
	Line 28: {REACTION;MAKE_WOODEN_MAIL_SHIRT_ARP;ADDON_HOOK_FOREST}
	Line 35: {REACTION;MAKE_WOODEN_HELM_ARP;ADDON_HOOK_FOREST}
	Line 42: {REACTION;MAKE_WOODEN_CAP_ARP;ADDON_HOOK_FOREST}
	Line 49: {REACTION;MAKE_WOODEN_BOOTS_ARP;ADDON_HOOK_FOREST}
	Line 56: {REACTION;MAKE_WOODEN_BOOTS_LOW_ARP;ADDON_HOOK_FOREST}
	Line 63: {REACTION;MAKE_WOODEN_GAUNTLETS_ARP;ADDON_HOOK_FOREST}
	Line 70: {REACTION;MAKE_WOODEN_GREAVES_ARP;ADDON_HOOK_FOREST}
	Line 77: {REACTION;MAKE_WOODEN_LEGGINGS_ARP;ADDON_HOOK_FOREST}
	Line 84: {REACTION;MAKE_WOODEN_MECHANISMS_ARP;ADDON_HOOK_FOREST}
	Line 91: {REACTION;GROW_WOOD_ARP;ADDON_HOOK_FOREST}
	Line 100: {REACTION;TAN_A_HIDE_ADV_ARP;ADDON_HOOK_FOREST}
	Line 108: {REACTION;MAKE_LEATHER_BACKPACK_ADV_ARP;ADDON_HOOK_FOREST}
	Line 115: {REACTION;MAKE_LEATHER_WATERSKIN_ADV_ARP;ADDON_HOOK_FOREST}
	Line 122: {REACTION;MAKE_LEATHER_QUIVER_ADV_ARP;ADDON_HOOK_FOREST}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Assembler\reaction_first_landing_powered_assembler.txt (1 hit)
	Line 3: {REACTION;ADJUST_POWERED_FL_AUTOASSEMBLER;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Drill\reaction_first_landing_powered_drill.txt (3 hits)
	Line 3: {REACTION;DFHACK_POWERED_DRILLING_RIG_DOWN;ADDON_HOOK_PLAYABLE}
	Line 10: {REACTION;DFHACK_POWERED_DRILLING_RIG_UP;ADDON_HOOK_PLAYABLE}
	Line 16: {REACTION;DFHACK_POWERED_DRILLING_RIG_PUMP;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Item Sensor\reaction_first_landing_powered_item_sensor.txt (1 hit)
	Line 3: {REACTION;ADJUST_POWERED_ITEM_SENSOR;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Logistics\reaction_first_landing_powered_logistics.txt (2 hits)
	Line 9: {REACTION;ADJUST_POWERED_CART_LAUNCHER;ADDON_HOOK_PLAYABLE}
	Line 19: {REACTION;ADJUST_POWERED_SORTER;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\Final\Libs\Powered\Tailor\reaction_first_landing_powered_tailor.txt (1 hit)
	Line 3: {REACTION;ADJUST_POWERED_FL_AUTOTAILOR;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_crates.txt (1 hit)
	Line 3: {REACTION;WORLDGEN_CRATES_PLAYABLE;ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\Toolbox\Cheat\reaction_dev_cheat.txt (1 hit)
	Line 3: {REACTION;{@GENERATE_ID;DEV_CHEAT};ADDON_HOOK_PLAYABLE}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\Toolbox\Dummy\Reactions\dev_dummy_reactions.pre.lua (1 hit)
	Line 10: 	out = out.."\n{REACTION;DUMMY_REACTION_"..count..";ADDON_HOOK_PLAYABLE}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test14\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 279: 	"{REACTION;%{id};%{class}}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Libs\Delta\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 279: 	"{REACTION;%{id};%{class}}\n"..
{DFHACK_REACTION;<ID>;<COMMAND>;<CLASS>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Upgrade Building\libs_upgrade_building.rbl (1 hit)
	Line 3: {DFHACK_REACTION;UPGRADE_BUILDING_%{source}_TO_%{dest};libs_upgrade_building -type "%{dest}" -loc [ \\LOCATION ];%{class}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Adv Reactions\Disinfect\reaction_user_adv_reactions_disinfect.txt (1 hit)
	Line 3: {DFHACK_REACTION;DISINFECT_ADV;user_adv_reactions_disinfect;NULL}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Metallurgy\Volcanic\reaction_user_metallurgy_volcanic.txt (1 hit)
	Line 3: {DFHACK_REACTION;ALLOY_VOLCANIC;libs_fluids -area 5x5 -amount 4 -loc [ \\LOCATION ];ADDON_HOOK_PLAYABLE}
{DFHACK_REACTION_BIND;<COMMAND>;[<ID>=nil]}
{REACTION_ADD_CLASS;<CLASS>;[<ID>=nil]}
Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Circular kitchen\Libs\reaction_dried_meals.txt (2 hits)
	Line 2: {REACTION;DRY_MEALS;ADDON_HOOK_MOUNTAIN}{REACTION_ADD_CLASS;ADDON_HOOK_EVIL}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 2: {REACTION;DRY_MEALS;ADDON_HOOK_MOUNTAIN}{REACTION_ADD_CLASS;ADDON_HOOK_EVIL}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Glass Forge\reaction_user_glass_forge.txt (10 hits)
	Line 40: {GLASS_ITEM;WEAPON;ITEM_WEAPON_PICK;pick;ADDON_HOOK_MOUNTAIN}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 44: {GLASS_ITEM;WEAPON;ITEM_WEAPON_AXE_BATTLE;battle axe;ADDON_HOOK_MOUNTAIN}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 46: {GLASS_ITEM;WEAPON;ITEM_WEAPON_HAMMER_WAR;war hammer;ADDON_HOOK_MOUNTAIN}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 48: {GLASS_ITEM;WEAPON;ITEM_WEAPON_SWORD_SHORT;short sword;ADDON_HOOK_MOUNTAIN}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 50: {GLASS_ITEM;WEAPON;ITEM_WEAPON_SPEAR;spear;ADDON_HOOK_MOUNTAIN}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 52: {GLASS_ITEM;WEAPON;ITEM_WEAPON_MACE;mace;ADDON_HOOK_MOUNTAIN}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 54: {GLASS_ITEM;WEAPON;ITEM_WEAPON_CROSSBOW;crossbow;ADDON_HOOK_MOUNTAIN}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 56: {GLASS_ITEM;AMMO;ITEM_AMMO_BOLTS;bolt;ADDON_HOOK_MOUNTAIN;1;25}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 58: {GLASS_ITEM;WEAPON;ITEM_WEAPON_BOW;bow;ADDON_HOOK_PLAINS}{REACTION_ADD_CLASS;ADDON_HOOK_SWAMP}
	Line 60: {GLASS_ITEM;AMMO;ITEM_AMMO_ARROWS;arrow;ADDON_HOOK_PLAINS;1;25}{REACTION_ADD_CLASS;ADDON_HOOK_SWAMP}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Gnomes\aaa_gnome.rbl (35 hits)
	Line 51: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;BITUMINOUS_COAL_TO_COKE}
	Line 52: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;BITUMINOUS_COAL_TO_COKE}
	Line 54: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;TAN_A_HIDE}
	Line 55: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;RENDER_FAT}
	Line 56: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_SOAP_FROM_TALLOW}
	Line 57: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_SOAP_FROM_OIL}
	Line 58: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_PEARLASH}
	Line 59: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_PLASTER_POWDER}
	Line 60: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MILL_SEEDS_NUTS_TO_PASTE}
	Line 61: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;PRESS_OIL}
	Line 62: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;PRESS_HONEYCOMB}
	Line 63: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_WAX_CRAFTS}
	Line 64: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_MEAD}
	Line 65: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;BREW_DRINK_FROM_PLANT}
	Line 66: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;BREW_DRINK_FROM_PLANT_GROWTH}
	Line 67: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;PROCESS_PLANT_TO_BAG}
	Line 68: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_CLAY_JUG}
	Line 69: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_CLAY_BRICKS}
	Line 70: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_CLAY_STATUE}
	Line 71: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_LARGE_CLAY_POT}
	Line 72: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_CLAY_CRAFTS}
	Line 73: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_CLAY_HIVE}
	Line 74: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;GLAZE_JUG}
	Line 75: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;GLAZE_STATUE}
	Line 76: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;GLAZE_LARGE_POT}
	Line 77: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;GLAZE_CRAFT}
	Line 79: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_QUICKLIME}
	Line 80: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_MILK_OF_LIME}
	Line 81: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_PARCHMENT}
	Line 82: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_SCROLL}
	Line 83: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_QUIRE}
	Line 84: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_SHEET_FROM_PLANT}
	Line 85: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;MAKE_SLURRY_FROM_PLANT}
	Line 86: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;PRESS_PLANT_PAPER}
	Line 87: {REACTION_ADD_CLASS;ADDON_HOOK_MECHANICS;BIND_BOOK}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Human\Playable\user_human_playable.rbl (5 hits)
	Line 8: {REACTION_ADD_CLASS;ADDON_HOOK_PLAINS;BITUMINOUS_COAL_TO_COKE}
	Line 9: {REACTION_ADD_CLASS;ADDON_HOOK_PLAINS;BITUMINOUS_COAL_TO_COKE}
	Line 11: {REACTION_ADD_CLASS;ADDON_HOOK_PLAINS;MAKE_PEARLASH}
	Line 12: {REACTION_ADD_CLASS;ADDON_HOOK_PLAINS;MAKE_PLASTER_POWDER}
	Line 13: {REACTION_ADD_CLASS;ADDON_HOOK_PLAINS;PROCESS_PLANT_TO_BAG}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Saurian\aaa_saurian.rbl (35 hits)
	Line 60: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;BITUMINOUS_COAL_TO_COKE}
	Line 61: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;BITUMINOUS_COAL_TO_COKE}
	Line 63: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;TAN_A_HIDE}
	Line 64: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;RENDER_FAT}
	Line 65: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_SOAP_FROM_TALLOW}
	Line 66: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_SOAP_FROM_OIL}
	Line 67: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_PEARLASH}
	Line 68: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_PLASTER_POWDER}
	Line 69: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MILL_SEEDS_NUTS_TO_PASTE}
	Line 70: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;PRESS_OIL}
	Line 71: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;PRESS_HONEYCOMB}
	Line 72: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_WAX_CRAFTS}
	Line 73: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_MEAD}
	Line 74: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;BREW_DRINK_FROM_PLANT}
	Line 75: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;BREW_DRINK_FROM_PLANT_GROWTH}
	Line 76: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;PROCESS_PLANT_TO_BAG}
	Line 78: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_CLAY_JUG}
	Line 79: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_CLAY_BRICKS}
	Line 80: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_CLAY_STATUE}
	Line 81: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_LARGE_CLAY_POT}
	Line 82: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_CLAY_CRAFTS}
	Line 83: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_CLAY_HIVE}
	Line 84: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;GLAZE_JUG}
	Line 85: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;GLAZE_STATUE}
	Line 86: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;GLAZE_LARGE_POT}
	Line 87: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;GLAZE_CRAFT}
	Line 89: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_QUICKLIME}
	Line 90: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_MILK_OF_LIME}
	Line 91: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_PARCHMENT}
	Line 92: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_SCROLL}
	Line 93: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_QUIRE}
	Line 94: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_SHEET_FROM_PLANT}
	Line 95: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;MAKE_SLURRY_FROM_PLANT}
	Line 96: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;PRESS_PLANT_PAPER}
	Line 97: {REACTION_ADD_CLASS;ADDON_HOOK_SWAMP;BIND_BOOK}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Uncle spam\Base\leathery vegetables\reaction_plant_USM.txt (2 hits)
	Line 12: }{REACTION_ADD_CLASS;ADDON_HOOK_EVIL}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 12: }{REACTION_ADD_CLASS;ADDON_HOOK_EVIL}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\REACTION\reaction_VANILLA_other.txt (33 hits)
	Line 14: }{REACTION_ADD_CLASS;ADDON_HOOK_EVIL}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 14: }{REACTION_ADD_CLASS;ADDON_HOOK_EVIL}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 24: }{REACTION_ADD_CLASS;ADDON_HOOK_EVIL}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 24: }{REACTION_ADD_CLASS;ADDON_HOOK_EVIL}{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 39: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 59: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 97: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 107: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 117: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 127: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 137: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 147: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 167: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 182: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 197: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 212: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 229: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 247: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 257: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 274: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 283: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 301: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 319: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 339: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 366: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 384: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 398: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 409: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 417: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 429: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 442: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 453: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 471: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\REACTION\reaction_VANILLA_smelter.txt (16 hits)
	Line 29: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 39: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 49: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 59: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 69: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 79: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 89: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 99: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 109: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 119: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 129: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 139: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 150: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 186: {!SHARED_REACTION;BLACK_BRONZE_MAKING;ADDON_HOOK_MOUNTAIN;{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 205: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
	Line 215: }{REACTION_ADD_CLASS;ADDON_HOOK_PLAINS}
{REMOVE_REACTION;<ID>;<CLASS>}
{REMOVE_REACTION_FROM_PLAYABLES;<ID>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Metallurgy\Bloomery\user_metallurgy_bloomery.rbl (5 hits)
	Line 2: {REMOVE_REACTION_FROM_PLAYABLES;ALLOY_PIG_IRON}
	Line 3: {REMOVE_REACTION_FROM_PLAYABLES;ALLOY_STEEL}
	Line 4: {REMOVE_REACTION_FROM_PLAYABLES;PIG_IRON_MAKING}
	Line 5: {REMOVE_REACTION_FROM_PLAYABLES;STEEL_MAKING}
	Line 6: {REMOVE_REACTION_FROM_PLAYABLES;SMELT_ORE_IRON}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Metallurgy\Smelter\Remove Alloys\user_metallurgy_smelter_remove_alloys.rbl (20 hits)
	Line 2: {REMOVE_REACTION_FROM_PLAYABLES;BRASS_MAKING}
	Line 3: {REMOVE_REACTION_FROM_PLAYABLES;BRASS_MAKING2}
	Line 4: {REMOVE_REACTION_FROM_PLAYABLES;BRONZE_MAKING}
	Line 5: {REMOVE_REACTION_FROM_PLAYABLES;BRONZE_MAKING2}
	Line 6: {REMOVE_REACTION_FROM_PLAYABLES;ELECTRUM_MAKING}
	Line 7: {REMOVE_REACTION_FROM_PLAYABLES;ELECTRUM_MAKING2}
	Line 8: {REMOVE_REACTION_FROM_PLAYABLES;BILLON_MAKING}
	Line 9: {REMOVE_REACTION_FROM_PLAYABLES;BILLON_MAKING2}
	Line 10: {REMOVE_REACTION_FROM_PLAYABLES;PEWTER_FINE_MAKING}
	Line 11: {REMOVE_REACTION_FROM_PLAYABLES;PEWTER_FINE_MAKING2}
	Line 12: {REMOVE_REACTION_FROM_PLAYABLES;PEWTER_TRIFLE_MAKING}
	Line 13: {REMOVE_REACTION_FROM_PLAYABLES;PEWTER_TRIFLE_MAKING2}
	Line 14: {REMOVE_REACTION_FROM_PLAYABLES;PEWTER_LAY_MAKING}
	Line 15: {REMOVE_REACTION_FROM_PLAYABLES;PIG_IRON_MAKING}
	Line 16: {REMOVE_REACTION_FROM_PLAYABLES;STEEL_MAKING}
	Line 17: {REMOVE_REACTION_FROM_PLAYABLES;NICKEL_SILVER_MAKING}
	Line 18: {REMOVE_REACTION_FROM_PLAYABLES;BLACK_BRONZE_MAKING}
	Line 19: {REMOVE_REACTION_FROM_PLAYABLES;STERLING_SILVER_MAKING}
	Line 20: {REMOVE_REACTION_FROM_PLAYABLES;ROSE_GOLD_MAKING}
	Line 21: {REMOVE_REACTION_FROM_PLAYABLES;BISMUTH_BRONZE_MAKING}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Pottery\Glaze\user_pottery_glaze.rbl (6 hits)
	Line 2: {REMOVE_REACTION_FROM_PLAYABLES;GLAZE_ITEM_BAR}
	Line 3: {REMOVE_REACTION_FROM_PLAYABLES;GLAZE_ITEM_BOULDER}
	Line 5: {REMOVE_REACTION_FROM_PLAYABLES;GLAZE_JUG}
	Line 6: {REMOVE_REACTION_FROM_PLAYABLES;GLAZE_STATUE}
	Line 7: {REMOVE_REACTION_FROM_PLAYABLES;GLAZE_LARGE_POT}
	Line 8: {REMOVE_REACTION_FROM_PLAYABLES;GLAZE_CRAFT}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Pottery\user_pottery.rbl (10 hits)
	Line 2: {REMOVE_REACTION_FROM_PLAYABLES;MAKE_CLAY_JUG}
	Line 3: {REMOVE_REACTION_FROM_PLAYABLES;MAKE_CLAY_BRICKS}
	Line 4: {REMOVE_REACTION_FROM_PLAYABLES;MAKE_CLAY_STATUE}
	Line 5: {REMOVE_REACTION_FROM_PLAYABLES;MAKE_LARGE_CLAY_POT}
	Line 6: {REMOVE_REACTION_FROM_PLAYABLES;MAKE_CLAY_CRAFTS}
	Line 7: {REMOVE_REACTION_FROM_PLAYABLES;MAKE_CLAY_HIVE}
	Line 8: {REMOVE_REACTION_FROM_PLAYABLES;GLAZE_JUG}
	Line 9: {REMOVE_REACTION_FROM_PLAYABLES;GLAZE_STATUE}
	Line 10: {REMOVE_REACTION_FROM_PLAYABLES;GLAZE_LARGE_POT}
	Line 11: {REMOVE_REACTION_FROM_PLAYABLES;GLAZE_CRAFT}
{#USES_REACTIONS;<CLASS>}
{!REACTION_NEW_CATEGORY;<ID>;<NAME>;<DESCRIPTION>;[<KEY>="";[<PARENT>=""]]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Glass Forge\reaction_user_glass_forge.txt (3 hits)
	Line 3: {!REACTION_NEW_CATEGORY;GLASS_FORGE_GREEN;green glass;Forge items from raw green glass;CUSTOM_G}
	Line 5: {!REACTION_NEW_CATEGORY;GLASS_FORGE_CLEAR;clear glass;Forge items from raw clear glass;CUSTOM_C}
	Line 7: {!REACTION_NEW_CATEGORY;GLASS_FORGE_CRYSTAL;crystal glass;Forge items from raw crystal glass;CUSTOM_R}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\aaa_templates_flbase.rbl (1 hit)
	Line 3: 	{!REACTION_NEW_CATEGORY;%id;%name;%description;%key}{V;{@SET;FL_REACTION_CATEGORY;%id}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_armorer.txt (4 hits)
	Line 3: {!REACTION_NEW_CATEGORY;ARMORER_METAL;metal;Forge metal armor pieces.;CUSTOM_M}
	Line 37: {!REACTION_NEW_CATEGORY;ARMORER_FLEXCERAM;flexceram;Shape flexceram armor pieces.;CUSTOM_F}
	Line 86: {!REACTION_NEW_CATEGORY;ARMORER_LEATHER;leather;Sew leather armor pieces.;CUSTOM_L}
	Line 135: {!REACTION_NEW_CATEGORY;ARMORER_POWER;power armor;Assemble the king of protection, power armor.;CUSTOM_L}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_bonecarver.txt (2 hits)
	Line 3: {!REACTION_NEW_CATEGORY;BONECARVER_BONE;bone;Carve items from bone.;CUSTOM_B}
	Line 5: {!REACTION_NEW_CATEGORY;BONECARVER_SHELL;shell;Carve items from shell.;CUSTOM_S}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_furniture.txt (4 hits)
	Line 25: {!REACTION_NEW_CATEGORY;FURNITURE_STONE;stone furniture;Assemble various bits of furniture from stone blocks.;CUSTOM_S}
	Line 27: {!REACTION_NEW_CATEGORY;FURNITURE_WOOD;wooden furniture;Assemble various bits of furniture from wooden planks.;CUSTOM_W}
	Line 71: {!REACTION_NEW_CATEGORY;TOOLS_STONE;stone tools and crafts;Assemble tools, crafts, and other small items from stone blocks.;CUSTOM_SHIFT_S}
	Line 73: {!REACTION_NEW_CATEGORY;TOOLS_WOOD;wooden tools and crafts;Assemble tools, crafts, and other small items from wooden planks.;CUSTOM_SHIFT_W}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_gunsmith.txt (3 hits)
	Line 6: {!REACTION_NEW_CATEGORY;GUNSMITH_PARTS;parts;Make various small gun parts.;CUSTOM_P}
	Line 82: {!REACTION_NEW_CATEGORY;GUNSMITH_ASSEMBLY;assembly;Assemble firearms from prepared parts.;CUSTOM_A}
	Line 161: {!REACTION_NEW_CATEGORY;GUNSMITH_AMMO;ammo;Make ammunition for your firearms.;CUSTOM_R}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_mechanic.txt (2 hits)
	Line 3: {!REACTION_NEW_CATEGORY;MECHANIC_MECHANISMS;mechanisms;Assemble mechanisms from various materials.;CUSTOM_M}
	Line 55: {!REACTION_NEW_CATEGORY;MECHANIC_ITEMS;misc. items;Assemble various mechanical items.;CUSTOM_I}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_tailor.txt (4 hits)
	Line 21: {!REACTION_NEW_CATEGORY;TAILOR_CLOTHING_CLOTH;cloth clothing;Sew clothing from cloth.;CUSTOM_C}
	Line 23: {!REACTION_NEW_CATEGORY;TAILOR_CLOTHING_LEATHER;leather clothing;Sew clothing from leather.;CUSTOM_L}
	Line 41: {!REACTION_NEW_CATEGORY;TAILOR_OTHER_CLOTH;cloth items;Sew items from cloth.;CUSTOM_SHIFT_C}
	Line 43: {!REACTION_NEW_CATEGORY;TAILOR_OTHER_LEATHER;leather items;Sew items from leather.;CUSTOM_SHIFT_L}
{REACTION_CATEGORY;<ID>}
{#USES_TECH;<CLASS>}
{!ITEM_CLASS;<TYPE>;<ITEM>;<CLASS>;[<RARITY>=COMMON]}
{REMOVE_ITEM;<ID>;<CLASS>}
{REMOVE_ITEM_FROM_PLAYABLES;<ID>}
{#USES_ITEMS;<CLASS>}
{ADDON_HOOKS;<ID>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Gnomes\entity__Civilized_MECHANICS.txt (1 hit)
	Line 52: 	{ADDON_HOOKS;MECHANICS}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Saurian\entity__Savage_SWAMP.txt (1 hit)
	Line 45: 	{ADDON_HOOKS;SWAMP}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\entity_flbase_native.txt (2 hits)
	Line 12: 	{ADDON_HOOKS;NATIVES}
	Line 113: 	{ADDON_HOOKS;NATIVES_LL}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\entity_flbase_raider.txt (1 hit)
	Line 14: 	{ADDON_HOOKS;RAIDER}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\entity_flbase_settler.txt (1 hit)
	Line 12: 	{ADDON_HOOKS;SETTLER}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\entity_FL_base_native.txt (2 hits)
	Line 12: 	{ADDON_HOOKS;NATIVES}
	Line 113: 	{ADDON_HOOKS;NATIVES_LL}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\entity_FL_base_raider.txt (1 hit)
	Line 14: 	{ADDON_HOOKS;RAIDER}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\entity_FL_base_settler.txt (1 hit)
	Line 12: 	{ADDON_HOOKS;SETTLER}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Rhenaya\Drow\Libs\entity_RHENAYA_DROW.txt (1 hit)
	Line 18: 	{ADDON_HOOKS;RHENAYAS_DROW_ENTITY}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test06\entity_test.txt (1 hit)
	Line 12: 	{ADDON_HOOKS;MOUNTAIN}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Cavern\AnimalPeople\entity_VANILLA___Cavern_SUBTERRANEAN_ANIMAL_PEOPLES_default.txt (1 hit)
	Line 14: 	{ADDON_HOOKS;SUBTERRANEAN_ANIMAL_PEOPLES}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Civilized\Dwarves\entity_VANILLA___Civilized_MOUNTAIN_default.txt (1 hit)
	Line 12: 	{ADDON_HOOKS;MOUNTAIN}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Civilized\Elves\entity_VANILLA___Civilized_FOREST_default.txt (1 hit)
	Line 7: 	{ADDON_HOOKS;FOREST}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Civilized\Humans\entity_VANILLA___Civilized_PLAINS_default.txt (1 hit)
	Line 7: 	{ADDON_HOOKS;PLAINS}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Savage\Kobolds\entity_VANILLA___Savage_SKULKING_default.txt (1 hit)
	Line 5: 	{ADDON_HOOKS;SKULKING}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Slavers\Goblins\entity_VANILLA___Slavers_EVIL_default.txt (1 hit)
	Line 7: 	{ADDON_HOOKS;EVIL}
{!ENTITY_PLAYABLE;<ID>;<FORT>;<ADV>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\aaa_entity_libs_base.pre.lua (1 hit)
	Line 13: -- {!ENTITY_PLAYABLE;MOUNTAIN;true;true;false}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 285: 	"	{!ENTITY_PLAYABLE;%{id};%{fort};%{adv}}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Gnomes\entity__Civilized_MECHANICS.txt (1 hit)
	Line 4: 	{!ENTITY_PLAYABLE;MECHANICS;$GNOME_PLAYABLE;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Saurian\entity__Savage_SWAMP.txt (1 hit)
	Line 4: 	{!ENTITY_PLAYABLE;SWAMP;$SAURIAN_PLAYABLE;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test14\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 301: 	"	{!ENTITY_PLAYABLE;%{id};%{fort};%{adv}}\n"..
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Libs\Delta\aaa_sobject_libs_base.pre.lua (1 hit)
	Line 301: 	"	{!ENTITY_PLAYABLE;%{id};%{fort};%{adv}}\n"..
{ENTITY_PLAYABLE_EDIT;<ID>;<KEY>;<VALUE>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Gnomes\aaa_gnome.rbl (1 hit)
	Line 3: 	{ENTITY_PLAYABLE_EDIT;MOUNTAIN;FORT;false}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Human\Playable\user_human_playable.rbl (2 hits)
	Line 2: {ENTITY_PLAYABLE_EDIT;PLAINS;FORT;true}
	Line 3: {ENTITY_PLAYABLE_EDIT;MOUNTAIN;FORT;false}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Saurian\aaa_saurian.rbl (1 hit)
	Line 3: 	{ENTITY_PLAYABLE_EDIT;MOUNTAIN;FORT;false}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Lovechild\All races playable\Libs\_All races playable.rbl (3 hits)
	Line 24: {ENTITY_PLAYABLE_EDIT;FOREST;FORT;true}
	Line 25: {ENTITY_PLAYABLE_EDIT;PLAINS;FORT;true}
	Line 26: {ENTITY_PLAYABLE_EDIT;EVIL;FORT;true}
{@IF_ENTITY_PLAYABLE;<ID>;<KEY>;<THEN>;[<ELSE>=""]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\aaa_tech_libs_base.pre.lua (1 hit)
	Line 223: 	{#ECHO;{@IF_ENTITY_PLAYABLE;%{id};FORT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Gnomes\aaa_gnome.rbl (1 hit)
	Line 2: {ECHO;{@IF_ENTITY_PLAYABLE;MECHANICS;FORT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Gnomes\entity__Civilized_MECHANICS.txt (1 hit)
	Line 299: 		{@IF_ENTITY_PLAYABLE;MECHANICS;FORT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Human\Nobles\user_human_nobles.rbl (1 hit)
	Line 218: 		{@IF_ENTITY_PLAYABLE;PLAINS;FORT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Powered\Metal Press\reaction_user_powered_metal_press.txt (4 hits)
	Line 56: 	{#ECHO;{@IF_ENTITY_PLAYABLE;MOUNTAIN;FORT;
	Line 69: 	{#ECHO;{@IF_ENTITY_PLAYABLE;PLAINS;FORT;
	Line 94: 	{#ECHO;{@IF_ENTITY_PLAYABLE;SWAMP;FORT;
	Line 105: 	{#ECHO;{@IF_ENTITY_PLAYABLE;MECHANICS;FORT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Saurian\aaa_saurian.rbl (1 hit)
	Line 2: {ECHO;{@IF_ENTITY_PLAYABLE;SWAMP;FORT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Saurian\entity__Savage_SWAMP.txt (1 hit)
	Line 283: 		{@IF_ENTITY_PLAYABLE;SWAMP;FORT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\entity_flbase_settler.txt (1 hit)
	Line 287: 		{@IF_ENTITY_PLAYABLE;SETTLER;FORT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\entity_FL_base_settler.txt (1 hit)
	Line 287: 		{@IF_ENTITY_PLAYABLE;SETTLER;FORT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test06\entity_test.txt (1 hit)
	Line 503: 		{@IF_ENTITY_PLAYABLE;MOUNTAIN;FORT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Civilized\Dwarves\entity_VANILLA___Civilized_MOUNTAIN_default.txt (1 hit)
	Line 503: 		{@IF_ENTITY_PLAYABLE;MOUNTAIN;FORT;
{#ENTITY_NOBLES;<ID>;<DEFAULT>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Gnomes\entity__Civilized_MECHANICS.txt (1 hit)
	Line 244: 	{#ENTITY_NOBLES;MECHANICS;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Saurian\entity__Savage_SWAMP.txt (1 hit)
	Line 231: 	{#ENTITY_NOBLES;SWAMP;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\entity_flbase_raider.txt (1 hit)
	Line 171: 	{#ENTITY_NOBLES;RAIDER;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\entity_flbase_settler.txt (1 hit)
	Line 232: 	{#ENTITY_NOBLES;SETTLER;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\entity_FL_base_raider.txt (1 hit)
	Line 171: 	{#ENTITY_NOBLES;RAIDER;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\entity_FL_base_settler.txt (1 hit)
	Line 232: 	{#ENTITY_NOBLES;SETTLER;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test06\entity_test.txt (1 hit)
	Line 236: 	{#ENTITY_NOBLES;MOUNTAIN;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Cavern\AnimalPeople\entity_VANILLA___Cavern_SUBTERRANEAN_ANIMAL_PEOPLES_default.txt (1 hit)
	Line 100: 	{#ENTITY_NOBLES;SUBTERRANEAN_ANIMAL_PEOPLES;# This entity has no nobles}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Civilized\Dwarves\entity_VANILLA___Civilized_MOUNTAIN_default.txt (1 hit)
	Line 236: 	{#ENTITY_NOBLES;MOUNTAIN;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Civilized\Elves\entity_VANILLA___Civilized_FOREST_default.txt (1 hit)
	Line 147: 	{#ENTITY_NOBLES;FOREST;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Civilized\Humans\entity_VANILLA___Civilized_PLAINS_default.txt (1 hit)
	Line 180: 	{#ENTITY_NOBLES;PLAINS;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Savage\Kobolds\entity_VANILLA___Savage_SKULKING_default.txt (1 hit)
	Line 108: 	{#ENTITY_NOBLES;SKULKING;# This entity has no nobles}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\ENTITY\Slavers\Goblins\entity_VANILLA___Slavers_EVIL_default.txt (1 hit)
	Line 176: 	{#ENTITY_NOBLES;EVIL;
{ENTITY_ADD_NOBLE;<ID>;<NOBLE>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Succession Nobles\Dwarf\user_succession_nobles_dwarf.rbl (2 hits)
	Line 1: {ENTITY_ADD_NOBLE;MOUNTAIN;
	Line 3: {ENTITY_ADD_NOBLE;MOUNTAIN;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Succession Nobles\Gnome\user_succession_nobles_gnome.rbl (2 hits)
	Line 1: {ENTITY_ADD_NOBLE;MECHANICS;{ADD_POSITION_OVERSEER;Overseer;1;55;OUTPOST_ENGINEER;SETTLEMENT_ENGINEER;CITY_ARCHITECT;TOWN_ARCHITECT}}
	Line 2: {ENTITY_ADD_NOBLE;MECHANICS;{ADD_POSITION_OVERSEER;Former Overseer;AS_NEEDED;56;OUTPOST_ENGINEER;SETTLEMENT_ENGINEER;CITY_ARCHITECT;TOWN_ARCHITECT}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Succession Nobles\Human\user_succession_nobles_human.rbl (2 hits)
	Line 1: {ENTITY_ADD_NOBLE;PLAINS;{ADD_POSITION_OVERSEER;Overseer;1;65;EXPEDITION_LEADER;MAYOR;BARON;DUKE}}
	Line 2: {ENTITY_ADD_NOBLE;PLAINS;{ADD_POSITION_OVERSEER;Former Overseer;AS_NEEDED;66;EXPEDITION_LEADER;MAYOR;BARON;DUKE}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Succession Nobles\Saurian\user_succession_nobles_saurian.rbl (2 hits)
	Line 1: {ENTITY_ADD_NOBLE;SWAMP;{ADD_POSITION_OVERSEER;Overseer;1;65;SHAMAN;HIGH_SHAMAN;CHIEF_SAURIAN;HEAD_SAURIAN}}
	Line 2: {ENTITY_ADD_NOBLE;SWAMP;{ADD_POSITION_OVERSEER;Former Overseer;AS_NEEDED;66;SHAMAN;HIGH_SHAMAN;CHIEF_SAURIAN;HEAD_SAURIAN}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\GM-X\DarkAgesII\Libs\_DAII_Dark_Ages.rbl (2 hits)
	Line 100: {ENTITY_ADD_NOBLE;MOUNTAIN;
	Line 102: {ENTITY_ADD_NOBLE;MOUNTAIN;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Lovechild\All races playable\Libs\_All races playable.rbl (3 hits)
	Line 349: {ENTITY_ADD_NOBLE;FOREST;
	Line 362: {ENTITY_ADD_NOBLE;PLAINS;
	Line 377: {ENTITY_ADD_NOBLE;EVIL;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test07\entity_test.txt (1 hit)
	Line 21: {ENTITY_ADD_NOBLE;MECHANICS;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test08\entity_test.rbl (3 hits)
	Line 1: {ENTITY_ADD_NOBLE;MOUNTAIN;{ADD_POSITION_OVERSEER;Overseer;1;65;EXPEDITION_LEADER;MAYOR;BARON;DUKE}}
	Line 2: {ENTITY_ADD_NOBLE;MOUNTAIN;{ADD_POSITION_OVERSEER;Former Overseer;AS_NEEDED;66;EXPEDITION_LEADER;MAYOR;BARON;DUKE}}
	Line 3: {ENTITY_ADD_NOBLE;MECHANICS;
{ENTITY_REPLACE_NOBLES;<ID>;<NOBLES>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Human\Nobles\user_human_nobles.rbl (1 hit)
	Line 2: {ENTITY_REPLACE_NOBLES;PLAINS;
{DFHACK_RUNCOMMAND;<COMMAND>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\creature_TESB.txt (1 hit)
	Line 1151: {DFHACK_RUNCOMMAND;tesb-tame-all TESB_PET_ROCK}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\reaction_TESB.txt (5 hits)
	Line 19: 	{DFHACK_RUNCOMMAND;tesb-mining -living $TESB_LIVING_PROB -gem $TESB_GEM_PROB}
	Line 20: 	{DFHACK_RUNCOMMAND;tesb-weather -creature TESB_AWAKENED_STORM -weather rain}
	Line 21: 	{DFHACK_RUNCOMMAND;tesb-tribute}
	Line 37: {V;{DFHACK_RUNCOMMAND;modtools/reaction-trigger -reactionName TESB_TRIBUTE_TO_%{id} -syndrome "%{name} favor"}
	Line 46: {V;{DFHACK_RUNCOMMAND;modtools/reaction-trigger -reactionName TESB_TRIBUTE_TO_%{id}2 -syndrome "%{name} favor"}
{@ADV_TIME;<COUNT>;<UNIT>}
{@FORT_TIME;<COUNT>;<UNIT>}
{@GROWDUR;<COUNT>;<UNIT>}
{@BUILD_KEY;<KEY>;[<FURNACE>=false]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_bkey_libs_base.test (14 hits)
	Line 7: {@BUILD_KEY;CTRL_Z}
	Line 8: {@BUILD_KEY;CTRL_Z;true}
	Line 9: {@BUILD_KEY;CTRL_Z}
	Line 10: {@BUILD_KEY;CTRL_Z;true}
	Line 11: {@BUILD_KEY;CTRL_Z}
	Line 12: {@BUILD_KEY;CTRL_Z;true}
	Line 13: {@BUILD_KEY;A}
	Line 14: {@BUILD_KEY;A;true}
	Line 15: {@BUILD_KEY;A}
	Line 16: {@BUILD_KEY;A;true}
	Line 17: {@BUILD_KEY;A}
	Line 18: {@BUILD_KEY;A;true}
	Line 21: {@BUILD_KEY;A}
	Line 22: {@BUILD_KEY;A;true}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\ThirdParty\Steam Engine\building_user_steam_engine.txt (2 hits)
	Line 9: 	{@BUILD_KEY;SHIFT_S}
	Line 53: 	{@BUILD_KEY;ALT_S}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\building_TESB.txt (1 hit)
	Line 9: 	{@BUILD_KEY;ALT_T;false}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_armorer.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_A}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_bonecarver.txt (1 hit)
	Line 6: 	{@BUILD_KEY;SHIFT_B}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_chemist.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_H}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_common.txt (3 hits)
	Line 8: 	{@BUILD_KEY;SHIFT_R}
	Line 71: 	{@BUILD_KEY;SHIFT_Z}
	Line 94: 	{@BUILD_KEY;P}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_crates.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_L}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_electronics.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_E}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_forge.txt (1 hit)
	Line 9: 	{@BUILD_KEY;SHIFT_D}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_furniture.txt (2 hits)
	Line 7: 	{@BUILD_KEY;SHIFT_F}
	Line 48: 	{@BUILD_KEY;SHIFT_N}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_gunsmith.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_G}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_mechanic.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_C}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_paper.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_X}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_plastics.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_P}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_pottery.txt (1 hit)
	Line 10: 	{@BUILD_KEY;SHIFT_W}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_power.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_O}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_radio.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_Q}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_sawmill.txt (2 hits)
	Line 8: 	{@BUILD_KEY;SHIFT_S}
	Line 51: 	{@BUILD_KEY;SHIFT_K}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_tailor.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_T}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\building_flbase_uranium.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_U}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_armorer.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_A}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_bonecarver.txt (1 hit)
	Line 6: 	{@BUILD_KEY;SHIFT_B}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_chemist.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_H}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_common.txt (3 hits)
	Line 8: 	{@BUILD_KEY;SHIFT_R}
	Line 71: 	{@BUILD_KEY;SHIFT_Z}
	Line 94: 	{@BUILD_KEY;P}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_crates.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_L}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_electronics.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_E}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_forge.txt (1 hit)
	Line 9: 	{@BUILD_KEY;SHIFT_D}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_furniture.txt (2 hits)
	Line 7: 	{@BUILD_KEY;SHIFT_F}
	Line 48: 	{@BUILD_KEY;SHIFT_N}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_gunsmith.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_G}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_mechanic.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_C}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_paper.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_X}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_plastics.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_P}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_pottery.txt (1 hit)
	Line 10: 	{@BUILD_KEY;SHIFT_W}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_power.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_O}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_radio.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_Q}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_sawmill.txt (2 hits)
	Line 8: 	{@BUILD_KEY;SHIFT_S}
	Line 51: 	{@BUILD_KEY;SHIFT_K}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_tailor.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_T}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\building_b_FL_base_uranium.txt (1 hit)
	Line 7: 	{@BUILD_KEY;SHIFT_U}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\BUILDING\building_VANILLA_custom.txt (2 hits)
	Line 9: 	{@BUILD_KEY;SHIFT_S}
	Line 47: 	{@BUILD_KEY;P}
{@CLEAR_KEY;<KEY>;[<FURNACE>=false]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_bkey_libs_base.test (3 hits)
	Line 5: {@CLEAR_KEY;A;true}{@CLEAR_KEY;B}
	Line 5: {@CLEAR_KEY;A;true}{@CLEAR_KEY;B}
	Line 20: {@CLEAR_KEY;A;true}
{@STR_LOWER;<STRING>}
{@STR_UPPER;<STRING>}
{@STR_TITLE;<STRING>}
{@STR_REPLACE;<STRING>;<OLD>;<NEW>;[<N>=-1]}
{@STR_TO_ID;<STRING>}
{@STR_SPLIT;<STRING>;<DELIM>;[<MAX>=-1]}
{@GENERATE_ID;<PREFIX>}
{@GENERATE_COUNT;<NUMBER>}
{@MUL;<X>;<Y>}
{@DIV;<X>;<Y>}
{@IDIV;<X>;<Y>}
{@MOD;<X>;<Y>}
{@ADD;<X>;<Y>}
{@SUB;<X>;<Y>}
{@STORE_LIST;<MASTERKEY>;<INDEX>;<ITEM>;[<RTN>=false]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_base_libs_base.test (1 hit)
	Line 60: {@PARSE_TO;TEMP;{@STORE_LIST;Libs/Base:Tests;#;Testing...;true}}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (1 hit)
	Line 909: 	{@STORE_LIST;<MASTERKEY>;<INDEX>;<ITEM>;[<RTN>=false]}
{@READ_LIST;<MASTERKEY>;<INDEX>}

{@FOREACH_LIST;<MASTERKEY>;<RAWS>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test07\entity_test.txt (1 hit)
	Line 3: 	{@FOREACH_LIST;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (2 hits)
	Line 266: 	{E;{@FOREACH_LIST;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;
	Line 930: 	{@FOREACH_LIST;<MASTERKEY>;<RAWS>}
{@STORE_TABLE;<MASTERKEY>;<KEY>;<ITEM>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test07\entity_test.txt (1 hit)
	Line 3: 	{@FOREACH_LIST;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (2 hits)
	Line 266: 	{E;{@FOREACH_LIST;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;
	Line 930: 	{@FOREACH_LIST;<MASTERKEY>;<RAWS>}
{@READ_TABLE;<MASTERKEY>;<KEY>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_base_libs_base.test (1 hit)
	Line 72: {@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:Tests;TestKey}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_sobject_libs_base.test (1 hit)
	Line 29: {#;{@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;INORGANIC:TEST_SOD_C2}}{@IF;$TEMP;t;OK!;ERROR!}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test07\entity_test.txt (1 hit)
	Line 4: 		{@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;%val}}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (3 hits)
	Line 267: 		{@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;%val}
	Line 950: 	{@READ_TABLE;<MASTERKEY>;<KEY>}
	Line 1024: 	{@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;INORGANIC:IRON}
{@FOREACH_TABLE;<MASTERKEY>;<RAWS>}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (1 hit)
	Line 959: 	{@FOREACH_TABLE;<MASTERKEY>;<RAWS>}
{@FOREACH;<ITEMS>;<RAWS>;[<SEPA=|>;[<SEPB>==]]}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_base_libs_base.test (1 hit)
	Line 80: {@FOREACH;a=1|b=2|c=3;"\"%{key}\" = %{val},\n"}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Remove\Animalmen\Libs\ExceptFromLegends.rbl (1 hit)
	Line 1: {@FOREACH;ANIMAL_PERSON=0|ANIMAL_PERSON_LEGLESS=0;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Remove\From subterranean entity\Libs\AllExceptOne.rbl (1 hit)
	Line 1: {@FOREACH;REPTILE_MAN=0|SERPENT_MAN=0|RODENT MAN=0|BAT_MAN=0|ANT_MAN=0|OLM_MAN=0|CAVE_SWALLOW_MAN=0|CAVE_FISH_MAN=0;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Remove\Ubiquitous vermin\Libs\ubiquitous.rbl (1 hit)
	Line 1: {@FOREACH;GRASSHOPPER=0|TICK=0|LOUSE=0|THRIPS=0|MOSQUITO=0|TERMITE=0|FLY=0|BEETLE=0|ANT=0;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_inorganic_stone_MODEST.rbl (2 hits)
	Line 1: {@FOREACH;SANDSTONE=0|DOLOMITE=0|GABBRO=0|BASALT=0|QUARTZITE=0|SLATE=0|TALC=0|PETRIFIED_WOOD=0|PERICLASE=0|ILMENITE=0|RUTILE=0|MAGNETITE=0|PITCHBLENDE=0|BAUXITE=0|OLIVINE=0|ORTHOCLASE=0|MICA=0|ANHYDRITE=0|ALUNITE=0;
	Line 4: {@FOREACH;SILTSTONE=0|MUDSTONE=0|SHALE=0|CLAYSTONE=0|ROCK_SALT=0|LIMESTONE=0|GRANITE=0|DIORITE=0|RHYOLITE=0|ANDESITE=0|DACITE=0|PHYLLITE=0|SCHIST=0|GNEISS=0|CINNABAR=0|COBALTITE=0|JET=0|PUDDINGSTONE=0|GRAPHITE=0|BRIMSTONE=0|
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_material_template_MODEST.rbl (1 hit)
	Line 67: {@FOREACH;EYE=0|NERVE=0|BRAIN=0|LUNG=0|HEART=0|LIVER=0|GUT=0|STOMACH=0|GIZZARD=0|PANCREAS=0|SPLEEN=0|KIDNEY=0|HORN=0|HOOF=0|SHELL=0|CHITIN=0|STRUCTURAL_PLANT=0|=0|=0|=0|=0|=0|=0|=0|=0|=0|=0|=0|=0|=0|=0;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_tissue_template_MODEST.rbl (1 hit)
	Line 1: {@FOREACH;HORN=1000|HOOF=1000|CARTILAGE=1000|HAIR=1000|CHEEK_WHISKERS=1000|CHIN_WHISKERS=1000|MOUSTACHE=1000|SIDEBURNS=1000|
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Lovechild\All races playable\Libs\_All races playable.rbl (4 hits)
	Line 6: {@FOREACH;DWARF=0|ELF=0|GOBLIN=0|KOBOLD=0|CAVE_FISH_MAN=0|OLM_MAN=0|BAT_MAN=0|CAVE_SWALLOW_MAN=0|AMPHIBIAN_MAN=0|REPTILE_MAN=0|SERPENT_MAN=0|ANT_MAN=0|RODENT MAN=0;
	Line 9: {@FOREACH;TROLL=500|OGRE=500|BLIZZARD_MAN=500|GRIMELING=250;
	Line 14: {@FOREACH;WOLF_ICE=0|BLENDEC_FOUL=0|STRANGLER=0|NIGHTWING=0|HARPY=0|SEA_MONSTER=0;
	Line 20: {@FOREACH;PIKE=0|HALBERD=0|SWORD=0|MAUL=0|AXE_GREAT=0;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_bonecarver.txt (4 hits)
	Line 60: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;BONE;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 78: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;SHELL;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 97: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;BONE;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 114: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;SHELL;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_forge.txt (2 hits)
	Line 32: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;METAL;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 47: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;METAL;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_furniture.txt (4 hits)
	Line 6: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;STONE;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 36: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;STONE;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 50: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_glassblower.txt (2 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;GLASS;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 21: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;GLASS;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_plastics.txt (2 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;PLASTIC;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;PLASTIC;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_pottery.txt (2 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;POTTERY;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;POTTERY;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Putnam\Fantastic Mini mods\Libs\Materials Plus\Libs\_inorganic_PFMM.rbl (8 hits)
	Line 1: {@FOREACH;BLACK ZIRCON=0|GREEN ZIRCON=0|RED ZIRCON=0|BROWN ZIRCON=0|YELLOW ZIRCON=0|CLEAR ZIRCON=0;
	Line 4: {@FOREACH;HELIODOR=0|RED BERYL=0|MORGANITE=0|HONEY YELLOW BERYL=0|GOLDEN BERYL=0|GOSHENITE=0|AQUAMARINE=0|EMERALD=0;
	Line 7: {@FOREACH;DOLOMITE=0;
	Line 10: {@FOREACH;ANDESITE=0|OLIVINE=0|HORNBLENDE=0|SERPENTINE=0|ORTHOCLASE=0|MICROCLINE=0|MICA=0;
	Line 13: {@FOREACH;ILMENITE=0|RUTILE=0;
	Line 16: {@FOREACH;CHROMITE=0;
	Line 19: {@FOREACH;PYROLUSITE=0;
	Line 22: {@FOREACH;CHROMITE=0;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\IntensiveSprouting\plant_zISS.txt (3 hits)
	Line 3: {@FOREACH;NAME=[NAME:%{name} %{variety}]|NAME_PLURAL=NAME_PLURAL;[NAME_PLURAL:%{name}s %{variety}]|ADJ=[ADJ:%{name} %{variety}]|GROWDUR=[GROWDUR:%{growdur}]|CLUSTERSIZE=[CLUSTERSIZE:%{clustersize}]
	Line 8: {@FOREACH;MUSHROOM_HELMET_PLUMP=plump helmet|GRASS_TAIL_PIG=pig tail|GRASS_WHEAT_CAVE=cave wheat|POD_SWEET=pod sweet|BUSH_QUARRY=bush quarry|MUSHROOM_CUP_DIMPLE=dimple cup;
	Line 12: {@FOREACH;MUSHROOM_HELMET_PLUMP=plump helmet|GRASS_TAIL_PIG=pig tail|GRASS_WHEAT_CAVE=cave wheat|POD_SWEET=pod sweet|BUSH_QUARRY=bush quarry|MUSHROOM_CUP_DIMPLE=dimple cup;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test16\zzz___test.txt (1 hit)
	Line 2: {@FOREACH;BONE=0|HAIR=0|TOOTH=0|LEATHER=0|HORN=0|HOOF=0|PEARL=0|SILK=0|SHELL=0|SWEAT=0|TEARS=0|SPIT=0|MILK=0|TALLOW=0|SOAP=0|CHITIN=0|SCALE=0|CARTILAGE=0|BLOOD=0|ICHOR=0|GOO=0|SLIME=0|PUS=0|WOOD=0|THREAD=0;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Base\Haiku\Extreme\Base\Extreme Haiku.rbl (1 hit)
	Line 116: {@FOREACH;CREATURE=0|TRANSLATION=0;
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_bonecarver.txt (4 hits)
	Line 60: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;BONE;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 78: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;SHELL;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 97: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;BONE;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 114: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;SHELL;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_forge.txt (2 hits)
	Line 32: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;METAL;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 47: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;METAL;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_furniture.txt (4 hits)
	Line 6: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;STONE;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 36: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;STONE;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 50: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_glassblower.txt (2 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;GLASS;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 21: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;GLASS;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_plastics.txt (2 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;PLASTIC;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;PLASTIC;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_pottery.txt (2 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;POTTERY;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;POTTERY;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (3 hits)
	Line 970: 	{@FOREACH;<ITEMS>;<RAWS>;[<SEPA=|>;[<SEPB>==]]}
	Line 988: 	{@FOREACH;a=1|b=2|c=3;"\"%{key}\" = %{val},\n"}
	Line 999: 		{@FOREACH;a=1|b=2|c=3;"%a: \"%{}{key}\" = %{}{val},\n"}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Basics.md (1 hit)
	Line 459: 	{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
{!TEMPLATE;;;;}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_base_libs_base.test (3 hits)
	Line 3: {!TEMPLATE;FOO;bar}
	Line 6: {!TEMPLATE;GREET;thing;Hello %{thing}!}
	Line 8: {!TEMPLATE;GREET_DWARF;dwarf=Urist;{GREET;%{dwarf}}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Upgrade Building\libs_upgrade_building.rbl (1 hit)
	Line 2: {!TEMPLATE;DFHACK_REACTION_UPGRADE_BUILDING;source;dest;name;class;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Intensive Sprouting\Libs\plant_zISS.txt (1 hit)
	Line 2: {!TEMPLATE;PROTOTYPE;object;name;variety;seed;color;growdur;clustersize;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Intensive Sprouting\Libs\Z\Intensive sprouting2\plant_ZISS.txt (1 hit)
	Line 2: {!TEMPLATE;PROTOTYPE;object;name;variety;seed;color;growdur;clustersize;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\FPS\Libs\Circular Economy\Intensive Sprouting\Libs\Z\Intensive sprouting3\plant_ZISS.txt (1 hit)
	Line 2: {!TEMPLATE;PROTOTYPE;object;name;variety;seed;color;growdur;clustersize;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Block\Boulder\reaction_user_block_boulder.txt (1 hit)
	Line 2: {!TEMPLATE;PRESS_BOULDER;color;name;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Block\Cutter\reaction_user_block_cutter.txt (1 hit)
	Line 2: {!TEMPLATE;CUT_BLOCKS;color;name;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Block\inorganic_user_block.txt (1 hit)
	Line 2: {!TEMPLATE;@INORGANIC_BLOCK;id;color;name;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Block\Kiln\reaction_user_block_kiln.txt (1 hit)
	Line 2: {!TEMPLATE;FIRE_BLOCKS;color;name;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Bonecarver\reaction_user_bonecarver.txt (1 hit)
	Line 2: {!TEMPLATE;BONECARVER_ITEM;id;name;item;mats=1;count=1;fake=false;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Common\Sawmill\reaction_scrap_user_sawmill.txt (1 hit)
	Line 2: {!TEMPLATE;SCRAP_WOODEN_ITEM;item;name;scrap;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Archery\reaction_user_archery.txt (1 hit)
	Line 2: {!TEMPLATE;ARCHERY_ITEM;type;name;ammo;ammoname;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Glass Forge\reaction_user_glass_forge.txt (1 hit)
	Line 8: {!TEMPLATE;GLASS_ITEM;type;id;name;hook;materials=1;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Metallurgy\Smelter\reaction_aaa_user_metallurgy_smelter.txt (1 hit)
	Line 2: {!TEMPLATE;SMELT_ORE;metal;name;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Pottery\Batch\reaction_user_pottery_batch.txt (1 hit)
	Line 2: {!TEMPLATE;BATCH_POTTERY_ITEM;id;name;item;count=5;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Pottery\reaction_user_pottery.txt (1 hit)
	Line 2: {!TEMPLATE;POTTERY_ITEM;id;name;item;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Succession Nobles\user_succession_nobles.rbl (1 hit)
	Line 1: {!TEMPLATE;ADD_POSITION_OVERSEER;noble;number;precedence;position1;position2;position3;position4;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Warcrafter\Adventure\reaction_user_warcrafter_adventure.txt (2 hits)
	Line 109: {!TEMPLATE;LEATHER_ITEM_ADV;type;id;name;materials=1;count=1;
	Line 148: {!TEMPLATE;BONE_ITEM_ADV;type;id;name;materials=1;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Abadrausar\Humble\Libs\Quickfort\Libs\Vanilla\Warcrafter\reaction_user_warcrafter.txt (2 hits)
	Line 5: {!TEMPLATE;LEATHER_ITEM;type;id;name;materials=1;count=1;
	Line 62: {!TEMPLATE;SCALE_ITEM;type;id;name;materials=1;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Button\Modest\Libs\_reaction_MODEST.rbl (4 hits)
	Line 1: {!TEMPLATE;!SET_VALUE3;token;value;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 2: {!TEMPLATE;ATTACH_BEFORE_TAG;token;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};
	Line 5: {!TEMPLATE;ATTACH_AFTER_TAG;token;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};
	Line 8: {!TEMPLATE;REPLACE_TAG;token;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\building_TESB.txt (1 hit)
	Line 52: {!TEMPLATE;TESB_WSHOP;id;name;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\creature_TESB.txt (9 hits)
	Line 42: 	{!TEMPLATE;TESB_AWAKE_CASTE;name;id;color;surface;eye;
	Line 250: 	{!TEMPLATE;TESB_INCANDESCENT_CASTE;id;name;ratio;eye;
	Line 423: 	{!TEMPLATE;TESB_MAGMA_CASTE;id;name;ratio;eye;
	Line 637: 	{!TEMPLATE;TESB_STANDARD_WYRM_CASTE;name;id;color;surface;eye;gem;glowcolor;
	Line 693: 	{!TEMPLATE;TESB_MAGMA_WYRM_CASTE;name;id;color;surface;eye;gem;glowcolor;
	Line 749: 	{!TEMPLATE;TESB_QUAKE_WYRM_CASTE;name;id;color;surface;eye;gem;glowcolor;
	Line 806: 	{!TEMPLATE;TESB_AQUIFER_WYRM_CASTE;name;id;color;surface;eye;gem;glowcolor;
	Line 854: 	{!TEMPLATE;TESB_FLUX_WYRM_CASTE;name;id;color;surface;eye;gem;glowcolor;
	Line 1171: 	{!TEMPLATE;TESB_PET_ROCK_CASTE;id;name;color;surface;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\inorganic_TESB_stone.txt (1 hit)
	Line 352: {!TEMPLATE;TESB_FAVOR_SYNDROME;id;name;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\interaction_TESB.txt (6 hits)
	Line 3: {!TEMPLATE;TESB_SPIT_ROCK_INTERACTION;id;
	Line 24: {!TEMPLATE;TESB_SPIT_MAGMA_INTERACTION;id;
	Line 45: {!TEMPLATE;TESB_EARTHQUAKE_INTERACTION;id;
	Line 74: {!TEMPLATE;TESB_THROW_STEEL_INTERACTION;id;
	Line 95: {!TEMPLATE;TESB_THROW_WATER_INTERACTION;id;
	Line 232: {!TEMPLATE;TESB_FAVOR_INTERACTION;id;name;name2;type;sphere1;sphere2;sphere3=nil;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\plant_TESB.txt (1 hit)
	Line 11: {!TEMPLATE;TESB_VINE;vine;id;name;drinkcolor;shrubcolor;pickedcolor;growthcolor;displaycolor;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Dirst\The Earth Strikes Back!\Libs\reaction_TESB.txt (2 hits)
	Line 35: {!TEMPLATE;TESB_SACRIFICE;id;name;gemid;gemname;
	Line 114: {!TEMPLATE;TESB_EXTRACT;id;gemid;gemname;plant;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\GM-X\DarkAgesII\Libs\_DAII_Dark_Ages.rbl (4 hits)
	Line 20: {!TEMPLATE;@SET_VALUE;token;value;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 21: {!TEMPLATE;@ATTACH_BEFORE_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};%{preraws}[%{token}]}}
	Line 22: {!TEMPLATE;@ATTACH_AFTER_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}]%{preraws}}}
	Line 23: {!TEMPLATE;@REPLACE_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};%{preraws}}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Lovechild\All races playable\Libs\_All races playable.rbl (22 hits)
	Line 1: {!TEMPLATE;!SET_VALUE;token;value;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 2: {!TEMPLATE;@ATTACH_BEFORE_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};%{preraws}[%{token}:%{value}]}}
	Line 3: {!TEMPLATE;@ATTACH_AFTER_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 4: {!TEMPLATE;@REPLACE_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};%{preraws}}}
	Line 59: {!TEMPLATE;ADD_NOBLE_POSITION;name;number;precedence;soldier;
	Line 68: {!TEMPLATE;ADD_LAND_POSITION;landholder;landname;replaced;
	Line 72: {!TEMPLATE;ADD_LAND_RESPONSABILITY;responsability1;responsability2;responsability3;responsability4;
	Line 77: {!TEMPLATE;ADD_LAND_APPOINTER;appointer1;appointer2;appointer3;rappointer4;
	Line 97: {!TEMPLATE;ADD_REQUIRED;demands;mandates;boxes;cabinets;racks;stands;roomvalue;replaced;
	Line 110: {!TEMPLATE;MILITIA_COMMANDER;name;plural;soldier;soldiers;
	Line 123: {!TEMPLATE;MILITIA_CAPTAIN;name;plural;soldier;soldiers;
	Line 134: {!TEMPLATE;SHERIFF;name;plural;
	Line 155: {!TEMPLATE;CAPTAIN_OF_THE_GUARD;name;plural;soldier;soldiers;
	Line 178: {!TEMPLATE;EXPEDITION_LEADER;name;plural;
	Line 192: {!TEMPLATE;MAYOR;name;plural;
	Line 221: {!TEMPLATE;MANAGER;name;plural;
	Line 235: {!TEMPLATE;CHIEF_MEDICAL_DWARF;name;plural;
	Line 248: {!TEMPLATE;BROKER;name;plural;
	Line 261: {!TEMPLATE;BOOKKEEPER;name;plural;
	Line 275: {!TEMPLATE;OUTPOST_LIAISON;name;plural;
	Line 297: {!TEMPLATE;HAMMERER;name;plural;
	Line 311: {!TEMPLATE;ADD_POSITION_NOBLE;noble;female;landholder;landname;precedence;demands;mandates;boxes;cabinets;racks;stands;roomvalue;replaced;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\aaa_templates_flbase.rbl (1 hit)
	Line 2: {!TEMPLATE;!FL_REACTION_NEW_CATEGORY;id;name;description;key;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\descriptor_pattern_iris_eye.txt (1 hit)
	Line 2: {!TEMPLATE;!FL_EYE_PATTERN;color;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\material_template_flbase.txt (2 hits)
	Line 2: {!TEMPLATE;FL_MAT_NAME_COLOR_ADV;name;color;hname;hcolor;
	Line 10: }{!TEMPLATE;FL_MAT_NAME_COLOR;name;color;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_bonecarver.txt (1 hit)
	Line 6: {!TEMPLATE;!BONECARVER_ITEM;id;name;item;mats=1;count=1;fake=NO;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_forge.txt (1 hit)
	Line 2: {!TEMPLATE;!MACHINE_ITEM;id;name;item;mats=1;count=1;class=METAL_MAT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_furniture.txt (1 hit)
	Line 2: {!TEMPLATE;!WOOD_STONE_ITEM;id;name;item;mats=1;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_glassblower.txt (1 hit)
	Line 2: {!TEMPLATE;!BLOW_GLASS_ITEM;id;name;item;mats=1;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_plastics.txt (1 hit)
	Line 2: {!TEMPLATE;!PLASTIC_ITEM;id;name;item;mats=1;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_pottery.txt (1 hit)
	Line 2: {!TEMPLATE;!POTTERY_ITEM;id;name;item;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_radio.txt (2 hits)
	Line 2: {!TEMPLATE;!RADIO_BUY_CRATE;id;name;count;cost;
	Line 84: {!TEMPLATE;!RADIO_SELL_GEM;id;name;count;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_smelter.txt (1 hit)
	Line 2: {!TEMPLATE;!SMELTER_REACTION;id;body;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_tailor.txt (1 hit)
	Line 2: {!TEMPLATE;!TAILOR_ITEM;id;name;item;mats=1;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\aaa_templates_FL_base.rbl (1 hit)
	Line 2: {!TEMPLATE;!FL_REACTION_NEW_CATEGORY;id;name;description;key;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\descriptor_pattern_FL_base.txt (1 hit)
	Line 2: {!TEMPLATE;!FL_EYE_PATTERN;color;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\material_template_FL_base.txt (2 hits)
	Line 2: {!TEMPLATE;FL_MAT_NAME_COLOR_ADV;name;color;hname;hcolor;
	Line 10: }{!TEMPLATE;FL_MAT_NAME_COLOR;name;color;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_forge.txt (1 hit)
	Line 2: {!TEMPLATE;!MACHINE_ITEM;id;name;item;mats=1;count=1;class=METAL_MAT;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_radio.txt (2 hits)
	Line 2: {!TEMPLATE;!RADIO_BUY_CRATE;id;name;count;cost;
	Line 84: {!TEMPLATE;!RADIO_SELL_GEM;id;name;count;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_smelter.txt (1 hit)
	Line 2: {!TEMPLATE;!SMELTER_REACTION;id;body;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_tailor.txt (1 hit)
	Line 2: {!TEMPLATE;!TAILOR_ITEM;id;name;item;mats=1;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\Toolbox\Cheat\reaction_dev_cheat.txt (1 hit)
	Line 2: {!TEMPLATE;DEV_CHEAT_ITEM;name;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\IntensiveSprouting\plant_zISS.txt (1 hit)
	Line 2: {!TEMPLATE;PROTOTYPE;object;name;variety;seed;color;growdur;clustersize;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\ISS\plant_ISS.txt (1 hit)
	Line 2: {!TEMPLATE;PROTOTYPE;object;name;variety;seed;color;growdur;clustersize;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\ISS2\plant_ISS.txt (1 hit)
	Line 2: {!TEMPLATE;!PROTOTYPE;object;name;variety;seed;color;growdur;clustersize;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test03\plant_test3.txt (1 hit)
	Line 2: {!TEMPLATE;PROTOTYPE_HUMBLE;object;name;variety;seed;color;growdur;clustersize;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test04\plant_test4.txt (1 hit)
	Line 2: {!TEMPLATE;PROTOTYPE;object;name;variety;seed;color;growdur;clustersize;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test05\plant_test5.txt (1 hit)
	Line 2: {!TEMPLATE;!PROTOTYPE;object;name;variety;seed;color;growdur;clustersize;
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test06\entity_test.txt (2 hits)
	Line 760: {!TEMPLATE;SET_VALUE;token;value;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 761: {!TEMPLATE;REPLACE_TAG;token;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}]%{preraws}}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test07\entity_test.txt (3 hits)
	Line 10: {!TEMPLATE;CURSOR;token;object;nested;
	Line 15: {!TEMPLATE;SET_VALUE;token;value;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 16: {!TEMPLATE;ATTACH_TO_TAG;preraws;{SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;$TOKEN_CURSOR;[$TOKEN_CURSOR]%{preraws}}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Base\Haiku\Extreme\Base\Extreme Haiku.rbl (1 hit)
	Line 1: {!TEMPLATE;@TREE_GIVE_FRUIT;color;ncolor;fruit;seed;wine;alcohol;oil;flour;thread;dye;syrup;nogrowth
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Vanilla\Libs\Delta\aaa_libs_base_semantic_diff.lua (4 hits)
	Line 3: -- {!TEMPLATE;@SET_VALUE;token;value;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 4: -- {!TEMPLATE;@ATTACH_BEFORE_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};%{preraws}[%{token}:%{value}]}}
	Line 5: -- {!TEMPLATE;@ATTACH_AFTER_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
	Line 6: -- {!TEMPLATE;@REPLACE_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};%{preraws}}}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\aaa_templates_FL_base.rbl (1 hit)
	Line 2: {!TEMPLATE;!FL_REACTION_NEW_CATEGORY;id;name;description;key;
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\descriptor_pattern_FL_base.txt (1 hit)
	Line 2: {!TEMPLATE;!FL_EYE_PATTERN;color;
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\material_template_FL_base.txt (2 hits)
	Line 2: {!TEMPLATE;FL_MAT_NAME_COLOR_ADV;name;color;hname;hcolor;
	Line 10: }{!TEMPLATE;FL_MAT_NAME_COLOR;name;color;
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_forge.txt (1 hit)
	Line 2: {!TEMPLATE;!MACHINE_ITEM;id;name;item;mats=1;count=1;class=METAL_MAT;
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_radio.txt (2 hits)
	Line 2: {!TEMPLATE;!RADIO_BUY_CRATE;id;name;count;cost;
	Line 84: {!TEMPLATE;!RADIO_SELL_GEM;id;name;count;
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_smelter.txt (1 hit)
	Line 2: {!TEMPLATE;!SMELTER_REACTION;id;body;
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_tailor.txt (1 hit)
	Line 2: {!TEMPLATE;!TAILOR_ITEM;id;name;item;mats=1;count=1;
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (5 hits)
	Line 31: 	{!TEMPLATE;<NAME>;[<PARAM>[=<DEFAULT_VALUE>]...];<CODE>}
	Line 62: 	{!TEMPLATE;FOO;bar}
	Line 65: 	{!TEMPLATE;GREET;thing;Hello %{thing}!}
	Line 67: 	{!TEMPLATE;GREET_DWARF;dwarf=Urist;{GREET;%{dwarf}}}
	Line 998: 	{!TEMPLATE;T;a;
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Basics.md (2 hits)
	Line 406: 	{!TEMPLATE;EXAMPLE;arg;{#PRINT;%arg}}
	Line 417: 	{!TEMPLATE;EXAMPLE;arg;{#PRINT;%arg}}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Tutorials\User templates.md (1 hit)
	Line 20: 	{!TEMPLATE;FORGE_ITEM;id;name;type;mat;matname;count=1;
{@PARSE_TO;<ID>;<RAWS>}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_base_libs_base.test (4 hits)
	Line 32: {@PARSE_TO;TEMP;OK!}{ECHO;$TEMP}
	Line 60: {@PARSE_TO;TEMP;{@STORE_LIST;Libs/Base:Tests;#;Testing...;true}}
	Line 62: {@PARSE_TO;TEMP;{@READ_LIST;Libs/Base:Tests;0}}
	Line 72: {@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:Tests;TestKey}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Libs\Base\test_sobject_libs_base.test (1 hit)
	Line 29: {#;{@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;INORGANIC:TEST_SOD_C2}}{@IF;$TEMP;t;OK!;ERROR!}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing\Libs\reaction_flbase_forge.txt (1 hit)
	Line 54: {@PARSE_TO;TEMP;{!MACHINE_ITEM;DRILLING_RIG_BIT;hardened drill bit;TRAPCOMP:DFHACK_POWERED_DRILLING_RIG_BIT;6;1;WEAPON_METAL_MAT}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_bonecarver.txt (8 hits)
	Line 60: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;BONE;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 61: {@PARSE_TO;TEMP;{@GENERATE_ID;BONECARVER}}
	Line 78: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;SHELL;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 79: {@PARSE_TO;TEMP;{@GENERATE_ID;BONECARVER}}
	Line 97: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;BONE;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 98: {@PARSE_TO;TEMP;{@GENERATE_ID;BONECARVER}}
	Line 114: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;SHELL;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 115: {@PARSE_TO;TEMP;{@GENERATE_ID;BONECARVER}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_forge.txt (4 hits)
	Line 32: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;METAL;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 33: {@PARSE_TO;TEMP;{@GENERATE_ID;MACHINE_SHOP}}
	Line 47: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;METAL;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 48: {@PARSE_TO;TEMP;{@GENERATE_ID;MACHINE_SHOP}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_furniture.txt (8 hits)
	Line 6: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;STONE;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 7: {@PARSE_TO;TEMP;{@GENERATE_ID;FURNITURE_ASSEMBLER}}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 21: {@PARSE_TO;TEMP;{@GENERATE_ID;FURNITURE_ASSEMBLER}}
	Line 36: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;STONE;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 37: {@PARSE_TO;TEMP;{@GENERATE_ID;FURNITURE_ASSEMBLER}}
	Line 50: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 51: {@PARSE_TO;TEMP;{@GENERATE_ID;FURNITURE_ASSEMBLER}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_glassblower.txt (4 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;GLASS;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 5: {@PARSE_TO;TEMP;{@GENERATE_ID;GLASS_BLOWER}}
	Line 21: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;GLASS;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 22: {@PARSE_TO;TEMP;{@GENERATE_ID;GLASS_BLOWER}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_plastics.txt (4 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;PLASTIC;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 5: {@PARSE_TO;TEMP;{@GENERATE_ID;PLASTICS}}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;PLASTIC;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 21: {@PARSE_TO;TEMP;{@GENERATE_ID;PLASTICS}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Modder\Milo\First Landing1.7\Base\reaction_FL_base_pottery.txt (4 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;POTTERY;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 5: {@PARSE_TO;TEMP;{@GENERATE_ID;POTTERY}}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;POTTERY;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 21: {@PARSE_TO;TEMP;{@GENERATE_ID;POTTERY}}
  Z:\Dwarf Fortress\nanofort\Rubble\addons\Test\test07\entity_test.txt (1 hit)
	Line 4: 		{@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;%val}}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_bonecarver.txt (8 hits)
	Line 60: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;BONE;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 61: {@PARSE_TO;TEMP;{@GENERATE_ID;BONECARVER}}
	Line 78: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;SHELL;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 79: {@PARSE_TO;TEMP;{@GENERATE_ID;BONECARVER}}
	Line 97: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;BONE;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 98: {@PARSE_TO;TEMP;{@GENERATE_ID;BONECARVER}}
	Line 114: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;SHELL;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 115: {@PARSE_TO;TEMP;{@GENERATE_ID;BONECARVER}}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_forge.txt (4 hits)
	Line 32: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;METAL;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 33: {@PARSE_TO;TEMP;{@GENERATE_ID;MACHINE_SHOP}}
	Line 47: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;METAL;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 48: {@PARSE_TO;TEMP;{@GENERATE_ID;MACHINE_SHOP}}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_furniture.txt (8 hits)
	Line 6: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;STONE;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 7: {@PARSE_TO;TEMP;{@GENERATE_ID;FURNITURE_ASSEMBLER}}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 21: {@PARSE_TO;TEMP;{@GENERATE_ID;FURNITURE_ASSEMBLER}}
	Line 36: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;STONE;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 37: {@PARSE_TO;TEMP;{@GENERATE_ID;FURNITURE_ASSEMBLER}}
	Line 50: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 51: {@PARSE_TO;TEMP;{@GENERATE_ID;FURNITURE_ASSEMBLER}}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_glassblower.txt (4 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;GLASS;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 5: {@PARSE_TO;TEMP;{@GENERATE_ID;GLASS_BLOWER}}
	Line 21: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;GLASS;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 22: {@PARSE_TO;TEMP;{@GENERATE_ID;GLASS_BLOWER}}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_plastics.txt (4 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;PLASTIC;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 5: {@PARSE_TO;TEMP;{@GENERATE_ID;PLASTICS}}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;PLASTIC;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 21: {@PARSE_TO;TEMP;{@GENERATE_ID;PLASTICS}}
  Z:\Dwarf Fortress\nanofort\Rubble\First Landing1.7\Base\reaction_FL_base_pottery.txt (4 hits)
	Line 4: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;POTTERY;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 5: {@PARSE_TO;TEMP;{@GENERATE_ID;POTTERY}}
	Line 20: }{@PARSE_TO;TEMP;{!FL_GET_ITEMS;POTTERY;SMALL}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 21: {@PARSE_TO;TEMP;{@GENERATE_ID;POTTERY}}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Base Templates.md (3 hits)
	Line 267: 		{@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;%val}
	Line 1011: 	{@PARSE_TO;<ID>;<RAWS>}
	Line 1024: 	{@PARSE_TO;TEMP;{@READ_TABLE;Libs/Base:!SHARED_OBJECT_CATEGORY:INORGANIC;INORGANIC:IRON}
  Z:\Dwarf Fortress\nanofort\Rubble\other\Rubble Basics.md (2 hits)
	Line 459: 	{@PARSE_TO;TEMP;{!FL_GET_ITEMS;WOOD;LARGE}}{@FOREACH;$TEMP;{@STR_SPLIT;%val;:;4}
	Line 460: 	{@PARSE_TO;TEMP;{@GENERATE_ID;FURNITURE_ASSEMBLER}}