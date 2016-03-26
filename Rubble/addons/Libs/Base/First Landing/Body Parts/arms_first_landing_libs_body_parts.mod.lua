
_ENV = mkmodule("arms_first_landing_libs_body_parts")

local tables = require "first_landing_libs_random_bodies_tables"

-- These are part "prefabs", used to save time and avoid repetition when making body tables.

-- Arms
-------------------------------------------------

arms_generic = tables.make("part_meta", {
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_PUNCH]\n\t",
	},
	
	parts = "FLR_ARMS_GENERIC",
})

arms_generic_4 = tables.make("part_meta", {
	desc_desc = {"four arms"},
	desc_adj = {"four armed"},
	pref = {"many arms"},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_PUNCH]\n\t",
	},
	
	parts = "FLR_ARMS_GENERIC_4",
})

arms_pincers = tables.make("part_meta", {
	desc_desc = {"a pair of pincers"},
	pref = {"grasping pincers"},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_PINCER]\n\t",
	},
	
	parts = "FLR_ARMS_PINCERS",
})

arms_claws = tables.make("part_meta", {
	desc_desc = {"a pair of claws"},
	pref = {"fearsome claws"},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_PINCER]\n\t",
	},
	
	parts = "FLR_ARMS_CLAWS",
})

-- Fingers
-------------------------------------------------

fingers_5 = tables.make("part_meta", {
	desc_desc = {"five fingers"},
	
	parts = "FLR_FINGERS_5",
})

fingers_4 = tables.make("part_meta", {
	desc_desc = {"four fingers"},
	
	parts = "FLR_FINGERS_4",
})

fingers_3 = tables.make("part_meta", {
	desc_desc = {"three fingers"},
	
	parts = "FLR_FINGERS_3",
})

fingers_2 = tables.make("part_meta", {
	desc_desc = {"two fingers"},
	
	parts = "FLR_FINGERS_2",
})

add_finger_nails = tables.make("part_meta", {
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ADD_NAILS_FINGERS]\n\t",
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_SCRATCH_NAIL_FINGER]\n\t",
	},
})

add_finger_claws = tables.make("part_meta", {
	desc_desc = {"long claws on its fingers"},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ADD_CLAWS_FINGERS]\n\t",
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_SCRATCH_CLAW_FINGER]\n\t",
	}
})

-- All the fingers, weighted according to finger count (so more fingers are more common).
fingers_all = {
	{
		weight = 5,
		parts = fingers_5,
	},
	{
		weight = 4,
		parts = fingers_4,
	},
	{
		weight = 3,
		parts = fingers_3,
	},
	{
		weight = 2,
		parts = fingers_2,
	},
}

nails_claws_all = {
	{
		weight = 5,
		parts = add_finger_nails,
	},
	{
		weight = 3,
		parts = dd_finger_claws,
	},
	{
		weight = 1,
		parts = "",
	},
}

-- Wings
-------------------------------------------------

wings_generic = tables.make("part_meta", {
	desc_desc = {"wings"},
	desc_adj = {"winged"},
	pref = {"ability to fly"},
	
	custom = {
		winged = true,
	},
	
	parts = "FLR_WINGS_GENERIC",
})

wings_4 = tables.make("part_meta", {
	desc_desc = {"four wings"},
	desc_adj = {"winged"},
	pref = {"ability to fly"},
	
	custom = {
		winged = true,
	},
	
	parts = "FLR_WINGS_4",
})

wings_6 = tables.make("part_meta", {
	desc_desc = {"six wings"},
	desc_adj = {"winged"},
	pref = {"many wings"},
	
	custom = {
		winged = true,
	},
	
	parts = "FLR_WINGS_6",
})

wings_all = {
	{
		weight = 5,
		parts = wings_generic,
	},
	{
		weight = 3,
		parts = wings_4,
	},
	{
		weight = 1,
		parts = wings_6,
	},
}

return _ENV
