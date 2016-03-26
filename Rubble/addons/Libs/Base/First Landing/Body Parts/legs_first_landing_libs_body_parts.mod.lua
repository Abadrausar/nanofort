
_ENV = mkmodule("legs_first_landing_libs_body_parts")

local tables = require "first_landing_libs_random_bodies_tables"

-- These are part "prefabs", used to save time and avoid repetition when making body tables.

-- Legs
-------------------------------------------------

legs_generic = tables.make("part_meta", {
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_KICK]",
	},
	
	parts = "FLR_LEGS_GENERIC",
})

legs_quad = tables.make("part_meta", {
	desc_desc = {"four legs"},
	desc_noun = {"quadruped"},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_KICK]",
	},
	
	parts = "FLR_LEGS_QUAD",
})

legs_one_part_2 = tables.make("part_meta", {
	desc_desc = {"two simple legs"},
	
	parts = "FLR_LEGS_ONE_PART_FIRST",
})

legs_one_part_4 = tables.make("part_meta", {
	desc_desc = {"four simple legs"},
	
	parts = {
		"FLR_LEGS_ONE_PART_FIRST",
		"FLR_LEGS_ONE_PART_SECOND",
	},
})

legs_one_part_6 = tables.make("part_meta", {
	desc_desc = {"six legs"},
	
	parts = {
		"FLR_LEGS_ONE_PART_FIRST",
		"FLR_LEGS_ONE_PART_SECOND",
		"FLR_LEGS_ONE_PART_THIRD",
	},
})

legs_one_part_8 = tables.make("part_meta", {
	desc_desc = {"eight legs"},
	pref = {"many legs"},
	
	parts = {
		"FLR_LEGS_ONE_PART_FIRST",
		"FLR_LEGS_ONE_PART_SECOND",
		"FLR_LEGS_ONE_PART_THIRD",
		"FLR_LEGS_ONE_PART_FOURTH",
	},
})

legs_one_part_10 = tables.make("part_meta", {
	desc_desc = {"ten legs"},
	desc_verb = {"scurries about on ten skinny legs"},
	pref = {"many legs"},
	
	parts = {
		"FLR_LEGS_ONE_PART_FIRST",
		"FLR_LEGS_ONE_PART_SECOND",
		"FLR_LEGS_ONE_PART_THIRD",
		"FLR_LEGS_ONE_PART_FOURTH",
		"FLR_LEGS_ONE_PART_FIFTH",
	},
})

-- One parts legs from 6 to 10, weighted so 6 is most common and 10 is least common.
legs_one_part_6_to_10 = {
	{
		weight = 3,
		parts = legs_one_part_6,
	},
	{
		weight = 2,
		parts = legs_one_part_8,
	},
	{
		weight = 1,
		parts = legs_one_part_10,
	},
}

legs_tail = tables.make("part_meta", {
	desc_desc = {"tail"},
	desc_verb = {"slithers along on it's tail"},
	pref = {"muscular tail"},
	
	parts = "FLR_LEGS_TAIL",
})

-- Toes
-------------------------------------------------

toes_5 = tables.make("part_meta", {
	desc_desc = {"five toes"},
	
	parts = "FLR_TOES_5",
})

toes_4 = tables.make("part_meta", {
	desc_desc = {"four toes"},
	
	parts = "FLR_TOES_4",
})

toes_3 = tables.make("part_meta", {
	desc_desc = {"three toes"},
	
	parts = "FLR_TOES_3",
})

toes_2 = tables.make("part_meta", {
	desc_desc = {"two toes"},
	
	parts = "FLR_TOES_2",
})

-- All the toes, weighted according to toe count (so more toes are more common).
toes_all = {
	{
		weight = 5,
		parts = toes_5,
	},
	{
		weight = 4,
		parts = toes_4,
	},
	{
		weight = 3,
		parts = toes_3,
	},
	{
		weight = 2,
		parts = toes_2,
	},
}

add_toe_nails = tables.make("part_meta", {
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ADD_NAILS_TOES]",
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_SCRATCH_NAIL_TOE]",
	},
})

add_toe_claws = tables.make("part_meta", {
	desc_desc = {"long claws on its toes"},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ADD_CLAWS_TOES]",
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_SCRATCH_CLAW_TOE]",
	},
})

nails_claws_all = {
	{
		weight = 5,
		parts = add_toe_nails,
	},
	{
		weight = 3,
		parts = add_toe_claws,
	},
	{
		weight = 1,
		parts = "",
	},
}

return _ENV
