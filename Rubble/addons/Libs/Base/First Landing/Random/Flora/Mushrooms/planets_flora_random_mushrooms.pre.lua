
local namegen = require "first_landing_libs_random_names"
local random = require "libs_random"

local howmany = tonumber(rubble.configvar("FL_RANDOM_MUSHROOM_COUNT"))
if howmany == nil then
	rubble.print("      Config var FL_RANDOM_MUSHROOM_COUNT is not a number, using the default of 20.\n")
	howmany = 20
else
	rubble.print("      Generating "..howmany.." random giant mushrooms.\n")
end

-- Bright colored mushrooms are more common than bright colored trees
local woodcolors = {
	--["BLACK"] = 2,
	["BLUE"] = 2,
	["GREEN"] = 2,
	["CYAN"] = 2,
	["BROWN"] = 2,
	["PURPLE"] = 2,
	["ORANGE"] = 2,
	["LIGHT_GRAY"] = 1,
	["SLATE_GRAY"] = 2,
	["LIGHT_BLUE"] = 1,
	["LIGHT_GREEN"] = 1,
	["LIGHT_CYAN"] = 1,
	["RED"] = 1,
	["PINK"] = 1,
	["YELLOW"] = 1,
	["WHITE"] = 1,
}

local out = "[OBJECT:PLANT]\n"
for i = 1, howmany, 1 do
	local color = random.select_weighted(woodcolors, true)
	
	local hascap = random.chance(80)
	local depth = random.range(1, 3)
	
	
	local name = namegen.generate(random.range(4, 10), namegen.table_greek)
	
	out = out..
	"\n{!SHARED_PLANT;RANDOM_MUSHROOM_"..i.."_"..string.upper(name)..";\n"..
	"\t[NAME:"..name.." cap][NAME_PLURAL:"..name.." caps][ADJ:"..name.." cap]\n"..
	"\t\n"..
	"\t[PREFSTRING:randomness]\n"..
	"\t\n"..
	"\t[USE_MATERIAL_TEMPLATE:STRUCTURAL:STRUCTURAL_PLANT_TEMPLATE]\n"..
	"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
	"\t[BASIC_MAT:LOCAL_PLANT_MAT:STRUCTURAL]\n"..
	"\t\n"..
	"\t[USE_MATERIAL_TEMPLATE:SEED:FAKE_SEED_TEMPLATE]\n"..
	"\t[SEED:"..name.." spores"..name.." spores:0:0:1:LOCAL_PLANT_MAT:SEED]\n"..
	"\t\n"..
	"\t[USE_MATERIAL_TEMPLATE:WOOD:WOOD_TEMPLATE]\n"..
	"\t\t[STATE_NAME_ADJ:ALL_SOLID:"..name.." wood]\n"..
	"\t\t[PREFIX:NONE]\n"..
	"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
	"\t\t[SOLID_DENSITY:"..random.range(100, 2000).."]\n"..
	"\t[TREE:LOCAL_PLANT_MAT:WOOD]\n"..
	"\t\n"..
	"\t[BIOME:SUBTERRANEAN_WATER]\n"..
	"\t[UNDERGROUND_DEPTH:"..depth..":"..depth.."]\n"..
	"\t"..random.select{"[WET]\n", "[WET][DRY]\n", "[DRY]\n"}..
	"\t\n"..
	"\t[SAPLING][STANDARD_TILE_NAMES]\n"..
	"\t\n"
	
	if hascap then
		out = out..
		"\t[TREE_HAS_MUSHROOM_CAP]\n"..
		"\t[CAP_PERIOD:"..random.range(5, 15).."]\n"..
		"\t[CAP_RADIUS:"..random.range(1, 3).."]\n"..
		"\t[TRUNK_PERIOD:"..random.range(5, 15).."]\n"..
		"\t[MAX_TRUNK_HEIGHT:"..random.range(1, 8).."]\n"..
		"\t[MAX_TRUNK_DIAMETER:"..random.range(1, 3).."]\n"..
		"\t[TRUNK_WIDTH_PERIOD:"..random.range(5, 15).."]\n"
	else
		out = out..
		"\t[TRUNK_PERIOD:"..random.range(5, 15).."]\n"..
		"\t[HEAVY_BRANCH_DENSITY:0]\n"..
		"\t[BRANCH_DENSITY:0]\n"..
		"\t[MAX_TRUNK_HEIGHT:"..random.range(1, 8).."]\n"..
		"\t[HEAVY_BRANCH_RADIUS:0]\n"..
		"\t[BRANCH_RADIUS:0]\n"..
		"\t[TRUNK_BRANCHING:"..random.range(0, 5).."]\n"..
		"\t[MAX_TRUNK_DIAMETER:"..random.range(1, 3).."]\n"..
		"\t[TRUNK_WIDTH_PERIOD:"..random.range(5, 15).."]\n"..
		"\t\n"..
		"\t[TWIGS_SIDE_BRANCHES:0]\n"..
		"\t[TWIGS_ABOVE_BRANCHES:0]\n"..
		"\t[TWIGS_BELOW_BRANCHES:0]\n"..
		"\t[TWIGS_SIDE_HEAVY_BRANCHES:0]\n"..
		"\t[TWIGS_ABOVE_HEAVY_BRANCHES:0]\n"..
		"\t[TWIGS_BELOW_HEAVY_BRANCHES:0]\n"..
		"\t[TWIGS_SIDE_TRUNK:0]\n"..
		"\t[TWIGS_ABOVE_TRUNK:0]\n"..
		"\t[TWIGS_BELOW_TRUNK:0]\n"
	end
	
	out = out.."}\n"
end

rubble.files["plant_planets_flora_random_mushrooms.txt"].Content = out
