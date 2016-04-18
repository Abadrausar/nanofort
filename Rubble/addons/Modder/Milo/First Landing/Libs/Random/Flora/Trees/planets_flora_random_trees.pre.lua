
local namegen = require "first_landing_libs_random_names"
local random = require "libs_random"

local howmany = tonumber(rubble.configvar("FL_RANDOM_TREE_COUNT"))
if howmany == nil then
	rubble.print("      Config var FL_RANDOM_TREE_COUNT is not a number, using the default of 50.\n")
	howmany = 50
else
	rubble.print("      Generating "..howmany.." random trees.\n")
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

local woodcolors = {
	--["BLACK"] = 3,
	["BLUE"] = 3,
	["GREEN"] = 3,
	["CYAN"] = 3,
	["BROWN"] = 3,
	["PURPLE"] = 3,
	["ORANGE"] = 3,
	["LIGHT_GRAY"] = 1,
	["SLATE_GRAY"] = 3,
	["LIGHT_BLUE"] = 1,
	["LIGHT_GREEN"] = 1,
	["LIGHT_CYAN"] = 1,
	["RED"] = 1,
	["PINK"] = 1,
	["YELLOW"] = 1,
	["WHITE"] = 1,
}

local biomes = {
	{
		weight = 2,
		biome = "[BIOME:ANY_WETLAND]",
		archetypes = {"Normal", "Large"},
	},
	{
		weight = 2,
		biome = "[BIOME:FOREST_TAIGA]",
		archetypes = {"Small"},
	},
	{
		weight = 3,
		biome = "[BIOME:ANY_FOREST]",
		archetypes = {"Small", "Normal", "Large"},
	},
	{
		weight = 2,
		biome = "[BIOME:ANY_SHRUBLAND]",
		archetypes = {"Normal", "Small"},
	},
	{
		weight = 1,
		biome = "[BIOME:ANY_GRASSLAND][BIOME:ANY_SAVANNA]",
		archetypes = {"Normal", "Small"},
	},
	{
		weight = 1,
		biome = "[BIOME:ANY_DESERT]",
		archetypes = {"Stunted"},
	},
}

local archetypes = {
	-- Nice random trees, a little smaller than the max
	["Normal"] = function()
		return "\t[TRUNK_PERIOD:"..random.range(5, 15).."]\n"..
		"\t[HEAVY_BRANCH_DENSITY:"..random.range(10, 30).."]\n"..
		"\t[BRANCH_DENSITY:"..random.range(20, 50).."]\n"..
		"\t[MAX_TRUNK_HEIGHT:"..random.range(1, 5).."]\n"..
		"\t[HEAVY_BRANCH_RADIUS:"..random.select_weighted({["0"] = 1, ["1"] = 5}, true).."]\n"..
		"\t[BRANCH_RADIUS:"..random.range(1, 3).."]\n"..
		"\t[TRUNK_BRANCHING:"..random.range(0, 2).."]\n"..
		"\t[MAX_TRUNK_DIAMETER:"..random.range(1, 2).."]\n"..
		"\t[TRUNK_WIDTH_PERIOD:"..random.range(10, 20).."]\n"..
		"\t[ROOT_DENSITY:"..random.range(1, 10).."]\n"..
		"\t[ROOT_RADIUS:"..random.range(1, 3).."]\n"
	end,
	-- Larger than normal
	["Large"] = function()
		return "\t[TRUNK_PERIOD:"..random.range(5, 15).."]\n"..
		"\t[HEAVY_BRANCH_DENSITY:"..random.range(10, 30).."]\n"..
		"\t[BRANCH_DENSITY:"..random.range(20, 50).."]\n"..
		"\t[MAX_TRUNK_HEIGHT:"..random.range(4, 8).."]\n"..
		"\t[HEAVY_BRANCH_RADIUS:"..random.select_weighted({["0"] = 1, ["1"] = 2, ["2"] = 5}, true).."]\n"..
		"\t[BRANCH_RADIUS:"..random.range(1, 4).."]\n"..
		"\t[TRUNK_BRANCHING:"..random.range(0, 2).."]\n"..
		"\t[MAX_TRUNK_DIAMETER:"..random.range(1, 3).."]\n"..
		"\t[TRUNK_WIDTH_PERIOD:"..random.range(5, 15).."]\n"..
		"\t[ROOT_DENSITY:"..random.range(1, 10).."]\n"..
		"\t[ROOT_RADIUS:"..random.range(1, 3).."]\n"
	end,
	-- Smaller than normal
	["Small"] = function()
		return "\t[TRUNK_PERIOD:"..random.range(5, 15).."]\n"..
		"\t[HEAVY_BRANCH_DENSITY:"..random.range(10, 30).."]\n"..
		"\t[BRANCH_DENSITY:"..random.range(20, 50).."]\n"..
		"\t[MAX_TRUNK_HEIGHT:"..random.range(1, 2).."]\n"..
		"\t[HEAVY_BRANCH_RADIUS:"..random.select_weighted({["0"] = 1, ["1"] = 5}, true).."]\n"..
		"\t[BRANCH_RADIUS:"..random.range(1, 3).."]\n"..
		"\t[TRUNK_BRANCHING:"..random.range(0, 2).."]\n"..
		"\t[MAX_TRUNK_DIAMETER:1]\n"..
		"\t[TRUNK_WIDTH_PERIOD:"..random.range(10, 20).."]\n"..
		"\t[ROOT_DENSITY:"..random.range(1, 10).."]\n"..
		"\t[ROOT_RADIUS:"..random.range(1, 3).."]\n"
	end,
	-- Much smaller than normal and slow growing to boot
	["Stunted"] = function()
		return "\t[TRUNK_PERIOD:"..random.range(10, 20).."]\n"..
		"\t[HEAVY_BRANCH_DENSITY:10]\n"..
		"\t[BRANCH_DENSITY:20]\n"..
		"\t[MAX_TRUNK_HEIGHT:1]\n"..
		"\t[HEAVY_BRANCH_RADIUS:"..random.select_weighted({["0"] = 1, ["1"] = 5}, true).."]\n"..
		"\t[BRANCH_RADIUS:"..random.range(1, 2).."]\n"..
		"\t[TRUNK_BRANCHING:"..random.range(0, 2).."]\n"..
		"\t[MAX_TRUNK_DIAMETER:1]\n"..
		"\t[TRUNK_WIDTH_PERIOD:"..random.range(20, 40).."]\n"..
		"\t[ROOT_DENSITY:"..random.range(1, 10).."]\n"..
		"\t[ROOT_RADIUS:"..random.range(1, 3).."]\n"
	end,
}

local out = "[OBJECT:PLANT]\n"
for i = 1, howmany, 1 do
	local color = random.select_weighted(woodcolors, true)
	
	local biome = random.select_weighted(biomes)
	local archetype = archetypes[random.select(biome.archetypes)]
	
	local leaves = random.select(colors)
	local flowers = random.select(colors)
	local fruit = random.select(colors)
	
	local flowers_start = random.range(40, 60) * 1000
	local fruit_start = flowers_start * 2
	local flowers_end = fruit_start - 1
	local fruit_end = "200000"
	
	local name = namegen.generate(random.range(4, 10), namegen.table_greek)
	
	local hasfruit = random.chance(50)
	
	out = out..
	"\n{!SHARED_PLANT;RANDOM_TREE_"..i.."_"..string.upper(name)..";\n"..
	"\t[NAME:"..name.." tree][NAME_PLURAL:"..name.." trees][ADJ:"..name.." tree]\n"..
	"\t\n"..
	"\t[PREFSTRING:randomness]\n"..
	"\t\n"..
	"\t[USE_MATERIAL_TEMPLATE:STRUCTURAL:STRUCTURAL_PLANT_TEMPLATE]\n"..
	"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
	"\t[BASIC_MAT:LOCAL_PLANT_MAT:STRUCTURAL]\n"..
	"\t\n"..
	"\t[USE_MATERIAL_TEMPLATE:SEED:FAKE_SEED_TEMPLATE]\n"..
	"\t[SEED:"..name.." tree seed:"..name.." tree seeds:0:0:1:LOCAL_PLANT_MAT:SEED]\n"..
	"\t\n"..
	"\t[USE_MATERIAL_TEMPLATE:WOOD:WOOD_TEMPLATE]\n"..
	"\t\t[STATE_NAME_ADJ:ALL_SOLID:"..name.." wood]\n"..
	"\t\t[PREFIX:NONE]\n"..
	"\t\t{@COLOR_MATTAGS;"..color..";BLACK}\n"..
	"\t\t[SOLID_DENSITY:"..random.range(100, 2000).."]\n"..
	"\t[TREE:LOCAL_PLANT_MAT:WOOD][TREE_TILE:"..random.select{"226", "5", "6", "20", "23", "24"}.."]\n"..
	"\t\n"..
	"\t[USE_MATERIAL_TEMPLATE:LEAF:LEAF_TEMPLATE]\n"..
	"\t\t{@COLOR_MATTAGS;"..leaves..";BLACK}\n"..
	"\t\n"..
	"\t[USE_MATERIAL_TEMPLATE:FLOWER:FLOWER_TEMPLATE]\n"..
	"\t\t{@COLOR_MATTAGS;"..flowers..";BLACK}\n"..
	"\t\n"..
	"\t"..random.select_weighted{
		{
			weight = 2,
			biome = "[BIOME:ANY_WETLAND]",
		},
		{
			weight = 2,
			biome = "[BIOME:FOREST_TAIGA]",
		},
		{
			weight = 3,
			biome = "[BIOME:ANY_FOREST]",
		},
		{
			weight = 2,
			biome = "[BIOME:ANY_SHRUBLAND]",
		},
		{
			weight = 1,
			biome = "[BIOME:ANY_GRASSLAND][BIOME:ANY_SAVANNA]",
		},
	}.biome.."\n"..
	"\t"..random.select{"[WET]\n", "[WET][DRY]\n", "[DRY]\n"}..
	"\t\n"..
	"\t[SAPLING][STANDARD_TILE_NAMES]\n"..
	"\t\n"
	
	if hasfruit then
		local typ = random.select{"wine", "vodka", "liquor"}
		
		out = out..
		"\t[USE_MATERIAL_TEMPLATE:DRINK:PLANT_ALCOHOL_TEMPLATE]\n"..
		"\t\t[STATE_NAME_ADJ:ALL_SOLID:frozen "..name.." "..typ.."]\n"..
		"\t\t[STATE_NAME_ADJ:LIQUID:"..name.." "..typ.."]\n"..
		"\t\t[STATE_NAME_ADJ:GAS:boiling "..name.." "..typ.."]\n"..
		"\t\t[MATERIAL_VALUE:2]\n"..
		"\t\t{@COLOR_MATTAGS;"..fruit..";BLACK}\n"..
		"\t\t[EDIBLE_RAW][EDIBLE_COOKED]\n"..
		"\t\t[PREFIX:NONE]\n"..
		"\t[DRINK:LOCAL_PLANT_MAT:DRINK]\n"..
		"\t\n"..
		"\t[USE_MATERIAL_TEMPLATE:FRUIT:FRUIT_TEMPLATE]\n"..
		"\t\t{@COLOR_MATTAGS;"..fruit..";BLACK}\n"..
		"\t\t[EDIBLE_VERMIN][EDIBLE_RAW][EDIBLE_COOKED]\n"..
		"\t\t[STOCKPILE_PLANT_GROWTH]\n"..
		"\t\t[MATERIAL_REACTION_PRODUCT:DRINK_MAT:LOCAL_PLANT_MAT:DRINK]\n"..
		"\t\t[MATERIAL_REACTION_PRODUCT:SEED_MAT:LOCAL_PLANT_MAT:SEED]\n"..
		"\t\n"
	end
	
	out = out..
	archetype()..
	"\t\n"..
	"\t[TWIGS_SIDE_BRANCHES:"..random.int(1).."]\n"..
	"\t[TWIGS_ABOVE_BRANCHES:"..random.int(1).."]\n"..
	"\t[TWIGS_BELOW_BRANCHES:"..random.int(1).."]\n"..
	"\t[TWIGS_SIDE_HEAVY_BRANCHES:"..random.int(1).."]\n"..
	"\t[TWIGS_ABOVE_HEAVY_BRANCHES:"..random.int(1).."]\n"..
	"\t[TWIGS_BELOW_HEAVY_BRANCHES:"..random.int(1).."]\n"..
	"\t[TWIGS_SIDE_TRUNK:"..random.int(1).."]\n"..
	"\t[TWIGS_ABOVE_TRUNK:"..random.int(1).."]\n"..
	"\t[TWIGS_BELOW_TRUNK:"..random.int(1).."]\n"..
	"\t\n"..
	"\t[GROWTH:LEAVES]\n"..
	"\t\t[GROWTH_NAME:"..name.." leaf:"..name.." leaves]\n"..
	"\t\t[GROWTH_ITEM:PLANT_GROWTH:NONE:LOCAL_PLANT_MAT:LEAF]\n"..
	"\t\t[GROWTH_DENSITY:1000]\n"..
	"\t\t[GROWTH_HOST_TILE:BRANCHES_AND_TWIGS]\n"..
	"\t\t[GROWTH_HOST_TILE:SAPLING]\n"..
	"\t\t[GROWTH_PRINT:0:6:{@COLOR;"..leaves..";BLACK}:ALL:1]\n"..
	"\t\n"..
	"\t[GROWTH:FLOWERS]\n"..
	"\t\t[GROWTH_NAME:"..name.." flower:STP]\n"..
	"\t\t[GROWTH_ITEM:PLANT_GROWTH:NONE:LOCAL_PLANT_MAT:FLOWER]\n"..
	"\t\t[GROWTH_DENSITY:1000]\n"..
	"\t\t[GROWTH_HOST_TILE:BRANCHES_AND_TWIGS]\n"..
	"\t\t[GROWTH_TIMING:"..flowers_start..":"..flowers_end.."]\n"..
	"\t\t[GROWTH_PRINT:5:5:{@COLOR;"..flowers..";BLACK}:"..flowers_start..":"..flowers_end..":2]\n"
	
	if hasfruit then
		local where = "BRANCHES_AND_TWIGS"
		if random.chance(25) then
			where = "HEAVY_BRANCHES_AND_TRUNK"
		end
		
		out = out..
		"\t\n"..
		"\t[GROWTH:FRUIT]\n"..
		"\t\t[GROWTH_NAME:"..name.." fruit:STP]\n"..
		"\t\t[GROWTH_ITEM:PLANT_GROWTH:NONE:LOCAL_PLANT_MAT:FRUIT]\n"..
		"\t\t[GROWTH_DENSITY:1000]\n"..
		"\t\t[GROWTH_HOST_TILE:"..where.."]\n"..
		"\t\t[GROWTH_TIMING:"..fruit_start..":"..fruit_end.."]\n"..
		"\t\t[GROWTH_DROPS_OFF_NO_CLOUD]\n"..
		"\t\t[GROWTH_PRINT:'%':7:{@COLOR;"..fruit..";BLACK}:"..fruit_start..":"..fruit_end..":3]\n"..
		"\t\t[GROWTH_HAS_SEED]\n"
	end
	out = out.."}\n"
end

rubble.files["plant_planets_flora_random_trees.txt"].Content = out
