
_ENV = mkmodule("aquatic_first_landing_libs_body_parts")

local tables = require "first_landing_libs_random_bodies_tables"

-- These are part "prefabs", used to save time and avoid repetition when making body tables.
	
-- "Legs"
-------------------------------------------------

tail = tables.make("part_meta", {
	desc_desc = {"a finned tail"},
	
	parts = "FLR_DECO_TAIL",
})

flippers_tail = tables.make("part_meta", {
	desc_desc = {"two flippers with a powerful tail"},
	
	parts = {
		"FLR_FLIPPERS_SIDE",
		"FLR_DECO_TAIL",
	},
})

flippers = tables.make("part_meta", {
	desc_desc = {"four flippers"},
	
	custom = {
		amphibious = true,
	},
	
	parts = {
		"FLR_FLIPPERS_FRONT",
		"FLR_FLIPPERS_REAR",
	},
})

legs = {
	{
		weight = 1,
		parts = tail,
	},
	{
		weight = 1,
		parts = flippers_tail,
	},
	{
		weight = 1,
		parts = flippers,
	},
}

-- Body Deco
-------------------------------------------------

fins_dorsal = tables.make("part_meta", {
	desc_desc = {"a dorsal fin"},
	
	parts = "FLR_FINS_DORSAL",
})

fins_side = tables.make("part_meta", {
	desc_desc = { "a pair of side fins"},
	
	parts = "FLR_FINS_SIDE",
})

fins_both = tables.make("part_meta", {
	desc_desc = {"side and dorsal fins"},
	
	parts = {
		"FLR_FINS_DORSAL",
		"FLR_FINS_SIDE",
	},
})

return _ENV
