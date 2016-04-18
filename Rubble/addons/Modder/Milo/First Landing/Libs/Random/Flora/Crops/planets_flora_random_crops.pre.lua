
local namegen = require "first_landing_libs_random_names"
local random = require "libs_random"

local howmany = tonumber(rubble.configvar("FL_RANDOM_CROP_COUNT"))
if howmany == nil then
	rubble.print("      Config var FL_RANDOM_CROP_COUNT is not a number, using the default of 100.\n")
	howmany = 100
else
	rubble.print("      Generating "..howmany.." random crops.\n")
end

local colors = {
	--"BLACK",
	"BLUE",
	"GREEN",
	"CYAN",
	"BROWN",
	"PURPLE",
	"ORANGE",
	"LIGHT_GRAY",
	"SLATE_GRAY",
	"LIGHT_BLUE",
	"LIGHT_GREEN",
	"LIGHT_CYAN",
	"RED",
	"PINK",
	"YELLOW",
	"WHITE",
}

local biomes = {
	{
		weight = 4,
		biome = "[BIOME:SUBTERRANEAN_WATER][UNDERGROUND_DEPTH:1:1]",
		tag = "U",
	},
	{
		weight = 3,
		biome = "[BIOME:SUBTERRANEAN_WATER][UNDERGROUND_DEPTH:2:2]",
		tag = "U",
	},
	{
		weight = 3,
		biome = "[BIOME:SUBTERRANEAN_WATER][UNDERGROUND_DEPTH:3:3]",
		tag = "U",
	},
	{
		weight = 10,
		biome = "[BIOME:NOT_FREEZING]",
		tag = "A",
	},
	{
		weight = 1,
		biome = "[BIOME:ANY_TEMPERATE]",
		tag = "N",
	},
	{
		weight = 1,
		biome = "[BIOME:ANY_TROPICAL]",
		tag = "T",
	},
	{
		weight = 1,
		biome = "[BIOME:ANY_WETLAND]",
		tag = "W",
	},
	{
		weight = 1,
		biome = "[BIOME:ANY_DESERT]",
		tag = "D",
	},
}

local categories = {
	["Fruit"] = {
		weight = 6, -- Common
		tag = "F",
		descrip = "Edible, Brewable",
		structlines = {
			"\t\t[EDIBLE_COOKED]\n",
			"\t\t[EDIBLE_RAW]\n",
			"\t\t[MATERIAL_REACTION_PRODUCT:DRINK_MAT:LOCAL_PLANT_MAT:DRINK]\n",
		},
		raws = function(name)
			local color = random.select(colors)
			local wine = random.select{"wine", "vodka"}
			return "\t[USE_MATERIAL_TEMPLATE:DRINK:PLANT_ALCOHOL_TEMPLATE]\n"..
				"\t\t[STATE_NAME_ADJ:ALL_SOLID:frozen "..name.." "..wine.."]\n"..
				"\t\t[STATE_NAME_ADJ:LIQUID:"..name.." "..wine.."]\n"..
				"\t\t[STATE_NAME_ADJ:GAS:boiling "..name.." "..wine.."]\n"..
				"\t\t[MATERIAL_VALUE:2]\n"..
				"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
				"\t\t[EDIBLE_RAW]\n"..
				"\t\t[EDIBLE_COOKED]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t[DRINK:LOCAL_PLANT_MAT:DRINK]\n"
		end,
		test = "",
		combining = {["Dummy"] = 10, ["Oil"] = 4, ["Dye"] = 1},
	},
	["Grain"] = {
		weight = 4, -- Normal
		tag = "G",
		descrip = "Millable (flour), Brewable",
		structlines = {
			"\t\t[MATERIAL_REACTION_PRODUCT:FLOUR_MAT:LOCAL_PLANT_MAT:MILL]\n",
			"\t\t[MATERIAL_REACTION_PRODUCT:DRINK_MAT:LOCAL_PLANT_MAT:DRINK]\n",
		},
		raws = function(name)
			local flourcolor = random.select(colors)
			local beercolor = random.select(colors)
			local beer = random.select{"beer", "whiskey"}
			return "\t[USE_MATERIAL_TEMPLATE:MILL:PLANT_POWDER_TEMPLATE]\n"..
				"\t\t[MATERIAL_VALUE:5]\n"..
				"\t\t[STATE_NAME_ADJ:ALL_SOLID:"..name.." flour]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t{@COLOR_MATTAGS;"..flourcolor..";BLACK}\n"..
				"\t\t[EDIBLE_VERMIN]\n"..
				"\t\t[EDIBLE_COOKED]\n"..
				"\t[MILL:LOCAL_PLANT_MAT:MILL]\n"..
				"\t\n"..
				"\t[USE_MATERIAL_TEMPLATE:DRINK:PLANT_ALCOHOL_TEMPLATE]\n"..
				"\t\t[STATE_NAME_ADJ:ALL_SOLID:frozen "..name.." "..beer.."]\n"..
				"\t\t[STATE_NAME_ADJ:LIQUID:"..name.." "..beer.."]\n"..
				"\t\t[STATE_NAME_ADJ:GAS:boiling "..name.." "..beer.."]\n"..
				"\t\t[MATERIAL_VALUE:3]\n"..
				"\t\t{@COLOR_MATTAGS;"..beercolor..";BLACK}\n"..
				"\t\t[EDIBLE_RAW]\n"..
				"\t\t[EDIBLE_COOKED]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t[DRINK:LOCAL_PLANT_MAT:DRINK]\n"
		end,
		combining = {["Dummy"] = 10, ["Thread"] = 2},
	},
	["Sugar"] = {
		weight = 4, -- Normal
		tag = "S",
		descrip = "Millable (sugar), Millable (paste), Press paste to syrup, Brewable",
		structlines = {
			"\t\t[MATERIAL_REACTION_PRODUCT:SUGAR_MAT:LOCAL_PLANT_MAT:MILL]\n",
			"\t\t[MATERIAL_REACTION_PRODUCT:PASTE_MAT:LOCAL_PLANT_MAT:PASTE]\n",
			"\t\t[MATERIAL_REACTION_PRODUCT:DRINK_MAT:LOCAL_PLANT_MAT:DRINK]\n",
		},
		raws = function(name)
			local sugarcolor = random.select(colors)
			local syrupcolor = random.select(colors)
			local rumcolor = random.select(colors)
			return "\t[USE_MATERIAL_TEMPLATE:MILL:PLANT_POWDER_TEMPLATE]\n"..
				"\t\t[MATERIAL_VALUE:5]\n"..
				"\t\t[STATE_NAME_ADJ:ALL_SOLID:"..name.." sugar]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t{@COLOR_MATTAGS;"..sugarcolor..";BLACK}\n"..
				"\t\t[EDIBLE_VERMIN]\n"..
				"\t\t[EDIBLE_COOKED]\n"..
				"\t\t[EDIBLE_RAW]\n"..
				"\t[MILL:LOCAL_PLANT_MAT:MILL]\n"..
				"\t\n"..
				"\t[USE_MATERIAL_TEMPLATE:PASTE:STRUCTURAL_PLANT_TEMPLATE]\n"..
				"\t\t[MATERIAL_VALUE:2]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t[STATE_NAME_ADJ:ALL:"..name.." paste]\n"..
				"\t\t{@COLOR_MATTAGS;"..syrupcolor..";BLACK}\n"..
				"\t\t[STOCKPILE_GLOB_PASTE]\n"..
				"\t\t[EDIBLE_VERMIN]\n"..
				"\t\t[MATERIAL_REACTION_PRODUCT:PRESS_LIQUID_MAT:LOCAL_PLANT_MAT:SYRUP]\n"..
				"\t\t[MATERIAL_REACTION_PRODUCT:PRESS_CAKE_MAT:LOCAL_PLANT_MAT:CAKE]\n"..
				"\t\n"..
				"\t[USE_MATERIAL_TEMPLATE:CAKE:STRUCTURAL_PLANT_TEMPLATE]\n"..
				"\t\t[MATERIAL_VALUE:2]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t[STATE_NAME_ADJ:ALL:"..name.." press cake]\n"..
				"\t\t{@COLOR_MATTAGS;"..syrupcolor..";BLACK}\n"..
				"\t\t[STOCKPILE_GLOB_PRESSED]\n"..
				"\t\t[EDIBLE_VERMIN]\n"..
				"\t\t[EDIBLE_COOKED]\n"..
				"\t\n"..
				"\t[USE_MATERIAL_TEMPLATE:SYRUP:PLANT_EXTRACT_TEMPLATE]\n"..
				"\t\t[STATE_NAME_ADJ:ALL_SOLID:frozen "..name.." syrup]\n"..
				"\t\t[STATE_NAME_ADJ:LIQUID:"..name.." syrup]\n"..
				"\t\t[STATE_NAME_ADJ:GAS:boiling "..name.." syrup]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t{@COLOR_MATTAGS;"..syrupcolor..";BLACK}\n"..
				"\t\t[MATERIAL_VALUE:10]\n"..
				"\t\t[EDIBLE_COOKED]\n"..
				"\t\n"..
				"\t[USE_MATERIAL_TEMPLATE:DRINK:PLANT_ALCOHOL_TEMPLATE]\n"..
				"\t\t[STATE_NAME_ADJ:ALL_SOLID:frozen "..name.." rum]\n"..
				"\t\t[STATE_NAME_ADJ:LIQUID:"..name.." rum]\n"..
				"\t\t[STATE_NAME_ADJ:GAS:boiling "..name.." rum]\n"..
				"\t\t[MATERIAL_VALUE:1]\n"..
				"\t\t{@COLOR_MATTAGS;"..rumcolor..";BLACK}\n"..
				"\t\t[EDIBLE_RAW]\n"..
				"\t\t[EDIBLE_COOKED]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t[DRINK:LOCAL_PLANT_MAT:DRINK]\n"
		end,
		combining = {["Dummy"] = 10, ["Thread"] = 1},
	},
	["Oil"] = {
		weight = 4, -- Normal
		tag = "O",
		descrip = "Millable (paste), Press paste to oil",
		structlines = {
			"\t\t[MATERIAL_REACTION_PRODUCT:PASTE_MAT:LOCAL_PLANT_MAT:PASTE]\n",
		},
		raws = function(name)
			local color = random.select(colors)
			return "\t[USE_MATERIAL_TEMPLATE:PASTE:STRUCTURAL_PLANT_TEMPLATE]\n"..
				"\t\t[MATERIAL_VALUE:2]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t[STATE_NAME_ADJ:ALL:"..name.." paste]\n"..
				"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
				"\t\t[STOCKPILE_GLOB_PASTE]\n"..
				"\t\t[EDIBLE_VERMIN]\n"..
				"\t\t[MATERIAL_REACTION_PRODUCT:PRESS_LIQUID_MAT:LOCAL_PLANT_MAT:OIL]\n"..
				"\t\t[MATERIAL_REACTION_PRODUCT:PRESS_CAKE_MAT:LOCAL_PLANT_MAT:CAKE]\n"..
				"\t\n"..
				"\t[USE_MATERIAL_TEMPLATE:CAKE:STRUCTURAL_PLANT_TEMPLATE]\n"..
				"\t\t[MATERIAL_VALUE:2]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t[STATE_NAME_ADJ:ALL:"..name.." press cake]\n"..
				"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
				"\t\t[STOCKPILE_GLOB_PRESSED]\n"..
				"\t\t[EDIBLE_VERMIN]\n"..
				"\t\t[EDIBLE_COOKED]\n"..
				"\t\n"..
				"\t[USE_MATERIAL_TEMPLATE:OIL:PLANT_OIL_TEMPLATE]\n"..
				"\t\t[STATE_NAME_ADJ:ALL_SOLID:frozen "..name.." oil]\n"..
				"\t\t[STATE_NAME_ADJ:LIQUID:"..name.." oil]\n"..
				"\t\t[STATE_NAME_ADJ:GAS:boiling "..name.." oil]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
				"\t\t[MATERIAL_VALUE:5]\n"..
				"\t\t[EDIBLE_COOKED]\n"..
				"\t\n"..
				"\t[USE_MATERIAL_TEMPLATE:SOAP:PLANT_SOAP_TEMPLATE]\n"..
				"\t\t[STATE_NAME_ADJ:ALL_SOLID:"..name.." oil soap]\n"..
				"\t\t[STATE_NAME_ADJ:LIQUID:melted "..name.." oil soap]\n"..
				"\t\t[STATE_NAME_ADJ:GAS:boiling "..name.." oil soap]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
				"\t\t[MATERIAL_VALUE:5]\n"
		end,
		combining = {["Dummy"] = 10, ["Fruit"] = 4},
	},
	["Dye"] = {
		weight = 4, -- Normal
		tag = "D",
		descrip = "Millable (dye)",
		structlines = {
			"\t\t[MATERIAL_REACTION_PRODUCT:DYE_MAT:LOCAL_PLANT_MAT:MILL]\n",
		},
		raws = function(name)
			local color = random.select(colors)
			return "\t[USE_MATERIAL_TEMPLATE:MILL:PLANT_POWDER_TEMPLATE]\n"..
				"\t\t[MATERIAL_VALUE:15]\n"..
				"\t\t[STATE_NAME_ADJ:ALL_SOLID:"..name.." dye]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
				"\t\t[POWDER_DYE:{@COLOR;"..color.."}]\n"..
				"\t\t[EDIBLE_VERMIN]\n"..
				"\t\t[EDIBLE_COOKED]\n"..
				"\t\t[EDIBLE_RAW]\n"..
				"\t[MILL:LOCAL_PLANT_MAT:MILL]\n"
		end,
		combining = {["Dummy"] = 10, ["Thread"] = 1, ["Fruit"] = 2},
	},
	["Thread"] = {
		weight = 4, -- Normal
		tag = "T",
		descrip = "Spinnable, Paper",
		structlines = {
			"\t\t[MATERIAL_REACTION_PRODUCT:PRESS_PAPER_MAT:LOCAL_PLANT_MAT:THREAD]\n",
		},
		raws = function(name)
			return "\t[USE_MATERIAL_TEMPLATE:THREAD:THREAD_PLANT_TEMPLATE]\n"..
				"\t\t[STATE_NAME_ADJ:SOLID:"..name.."]\n"..
				"\t\t[STATE_NAME_ADJ:SOLID_PASTE:"..name.." pulp]\n"..
				"\t\t[STATE_NAME_ADJ:SOLID_PRESSED:"..name.." paper]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t[REACTION_CLASS:PAPER_SLURRY]\n"..
				"\t\t[MATERIAL_VALUE:2]\n"..
				"\t[THREAD:LOCAL_PLANT_MAT:THREAD]\n"
		end,
		combining = {["Dummy"] = 10, ["Dye"] = 1, ["Grain"] = 2},
	},
	["Bag"] = {
		weight = 1, -- Rare
		tag = "B",
		descrip = "Storage (bag)",
		structlines = {
			"\t\t[ITEM_REACTION_PRODUCT:SPECIAL_ITEM:BOX:NONE:LOCAL_PLANT_MAT:STRUCTURAL_LEAF]\n",
		},
		raws = function(name)
			local color = random.select(colors)
			return "\t[USE_MATERIAL_TEMPLATE:STRUCTURAL_LEAF:THREAD_PLANT_TEMPLATE]\n"..
				"\t\t[STATE_NAME:ALL_SOLID:"..name.." leaf]\n"..
				"\t\t[STATE_ADJ:ALL_SOLID:"..name.." leaf]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
				"\t\t[SOLID_DENSITY:200]\n"
		end,
		combining = {["Dummy"] = 1},
	},
	["Jug"] = {
		weight = 1, -- Rare
		tag = "J",
		descrip = "Storage (jug)",
		structlines = {
			"\t\t[ITEM_REACTION_PRODUCT:SPECIAL_ITEM:TOOL:ITEM_TOOL_JUG:LOCAL_PLANT_MAT:UNFINISHED_HUSK]\n",
		},
		raws = function(name)
			local color = random.select(colors)
			return "\t[USE_MATERIAL_TEMPLATE:HUSK:WOOD_TEMPLATE]\n"..
				"\t\t[STATE_NAME:ALL_SOLID:"..name.." husk]\n"..
				"\t\t[STATE_ADJ:ALL_SOLID:"..name.." husk]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
				"\t\t[SOLID_DENSITY:200]\n"..
				"\t\n"..
				"\t[USE_MATERIAL_TEMPLATE:UNFINISHED_HUSK:WOOD_TEMPLATE]\n"..
				"\t\t[STATE_NAME:ALL_SOLID:curing "..name.." husk]\n"..
				"\t\t[STATE_ADJ:ALL_SOLID:curing "..name.." husk]\n"..
				"\t\t[PREFIX:NONE]\n"..
				"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
				"\t\t[SOLID_DENSITY:200]\n"..
				"\t\t[ABSORPTION:20]\n"..
				"\t\t[MATERIAL_REACTION_PRODUCT:TRANSFORM_MAT:LOCAL_PLANT_MAT:HUSK]\n"
		end,
		combining = {["Dummy"] = 1},
	},
	["Dummy"] = {
		weight = 0,
		tag = "",
		descrip = "",
		structlines = {},
		raws = function(name)
			return ""
		end,
		-- It should be impossible to choose this as a primary, but just in case...
		combining = {["Fruit"] = 1, ["Grain"] = 1, ["Sugar"] = 1, ["Oil"] = 1, ["Dye"] = 1, ["Thread"] = 1},
	},
}

local out = "[OBJECT:PLANT]\n"
local descrip_lines = "" -- For eventual use (maybe)
for i = 1, howmany, 1 do
	local color = random.select(colors)
	local seed_color = random.select(colors)
	
	local name = namegen.generate(random.range(4, 10), namegen.table_greek)
	
	local tile = random.select{"6", "5", "'|'", "'#'", "':'", "'\"'", "176", "232", "231", "21"}
	
	local cata = random.select_weighted(categories)
	local catb = categories[random.select_weighted(cata.combining, true)]
	
	local biome_data = random.select_weighted(biomes)
	
	local id = "RANDOM_CROP_"..i.."_"..string.upper(name).."_"..cata.tag..catb.tag..biome_data.tag
	
	local seed_value
	if cata.tag == "" then
		seed_value = 1
		descrip_lines = descrip_lines.."\t[\""..id.."\"] = [["..catb.descrip.."]],\n"
	elseif catb.tag == "" then
		seed_value = 1
		descrip_lines = descrip_lines.."\t[\""..id.."\"] = [["..cata.descrip.."]],\n"
	else
		seed_value = 5
		descrip_lines = descrip_lines.."\t[\""..id.."\"] = [["..cata.descrip..", "..catb.descrip.."]],\n"
	end
	
	out = out..
	"\n{!SHARED_PLANT;"..id..";\n"..
	"\t[NAME:"..name.."][NAME_PLURAL:"..name.."][ADJ:"..name.."]\n"..
	"\t\n"..
	"\t[PREFSTRING:randomness]\n"..
	"\t\n"..
	"\t[USE_MATERIAL_TEMPLATE:STRUCTURAL:STRUCTURAL_PLANT_TEMPLATE]\n"..
	"\t\t[MATERIAL_REACTION_PRODUCT:SEED_MAT:LOCAL_PLANT_MAT:SEED]\n"..
	"\t\t[MATERIAL_VALUE:2]\n"..
	"\t\t[EDIBLE_VERMIN]\n"
	for _, line in ipairs(cata.structlines) do
		out = out..line
	end
	for _, line in ipairs(catb.structlines) do
		out = out..line
	end
	out = out..
	"\t[BASIC_MAT:LOCAL_PLANT_MAT:STRUCTURAL]\n"..
	"\t\n"..
	cata.raws(name)..
	catb.raws(name)..
	"\t\n"..
	"\t[USE_MATERIAL_TEMPLATE:SEED:SEED_TEMPLATE]\n"..
	"\t\t[MATERIAL_VALUE:"..seed_value.."]\n"..
	"\t\t[EDIBLE_VERMIN]\n"..
	"\t\t[EDIBLE_COOKED]\n"..
	"\t[SEED:"..name.." seed ("..cata.tag..catb.tag..biome_data.tag.."):"..name.." seeds ("..cata.tag..catb.tag..biome_data.tag.."):{@COLOR;"..seed_color..";BLACK}:LOCAL_PLANT_MAT:SEED]\n"..
	"\t\n"..
	"\t[PICKED_TILE:"..tile.."][PICKED_COLOR:{@COLOR;"..color..";BLACK}]\n"..
	"\t[SHRUB_TILE:"..tile.."][SHRUB_COLOR:{@COLOR;"..color..";BLACK}]\n"..
	"\t[DEAD_SHRUB_TILE:"..tile.."][DEAD_SHRUB_COLOR:0:0:1]\n"..
	"\t\n"..
	"\t[GROWDUR:"..random.range(300, 1000).."]\n"..
	"\t"..random.select_weighted({["[SPRING][SUMMER][AUTUMN][WINTER]"] = 2, ["[SPRING][SUMMER]"] = 3, ["[AUTUMN][WINTER]"] = 1, ["[SPRING][WINTER]"] = 1, ["[SUMMER][AUTUMN]"] = 3, ["[SPRING][SUMMER][AUTUMN]"] = 3}, true).."\n"..
	"\t\n"..
	"\t[FREQUENCY:50]\n"..
	"\t[CLUSTERSIZE:5]\n"..
	"\t\n"..
	"\t"..biome_data.biome.."\n"..
	"\t"..random.select{"[WET]\n", "[WET][DRY]\n", "[DRY]\n"}..
	"}\n"
end

rubble.files["plant_planets_flora_random_crops.txt"].Content = out
