
_ENV = mkmodule("head_first_landing_libs_body_parts")

local tables = require "first_landing_libs_random_bodies_tables"

-- These are part "prefabs", used to save time and avoid repetition when making body tables.

-- The head itself
-------------------------------------------------

head_generic = tables.make("part_meta", {
	extras = {"[BODY_DETAIL_PLAN:FLR_POSITIONS_HEAD]"},
	
	parts = "FLR_HEAD_GENERIC",
})

head_no_neck = tables.make("part_meta", {
	desc_desc = {"no neck"},
	desc_adj = {"neckless"},
	
	extras = {"[BODY_DETAIL_PLAN:FLR_POSITIONS_HEAD]"},
	
	parts = "FLR_HEAD_NO_NECK",
})

-- Use only with an exoskeleton.
head_insect = tables.make("part_meta", {
	extras = {"[BODY_DETAIL_PLAN:FLR_POSITIONS_HEAD]"},
	
	parts = "FLR_HEAD_INSECT",
})

-- Mouths
-------------------------------------------------

mouth_beak = tables.make("part_meta", {
	desc_desc = {"a sharp beak"},
	desc_adj = {"beaked"},
	pref = {"sharp beak"},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_BITE_BEAK]",
	},
	
	parts = "FLR_MOUTH_BEAK",
})

mouth_bill = tables.make("part_meta", {
	desc_desc = {"a large bill"},
	desc_adj = {"billed"},
	pref = {"large bill"},
	
	custom = {
		prey = true,
	},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_BITE_BEAK]",
	},
	
	parts = "FLR_MOUTH_BILL",
})

mouth_generic = tables.make("part_meta", {
	parts = "FLR_MOUTH_GENERIC",
})

-- Tongues
-------------------------------------------------

tongue_generic = tables.make("part_meta", {
	desc_desc = {"a long tongue"},
	
	parts = "FLR_TONGUE_GENERIC",
})

tongue_forked = tables.make("part_meta", {
	desc_desc = {"a forked tongue"},
	desc_adj = {"fork-tongued"},
	
	parts = "FLR_TONGUE_FORKED",
})

tongue_all = {
	{
		weight = 4,
		parts = tongue_generic,
	},
	{
		weight = 1,
		parts = tongue_forked,
	},
}

-- Mouth Deco
-------------------------------------------------

deco_cheeks = tables.make("part_meta", {
	desc_desc = {"fat cheeks"},
	
	parts = "FLR_DECO_CHEEKS",
})

deco_lips = tables.make("part_meta", {
	desc_desc = {"blubbery lips"},
	
	parts = "FLR_DECO_LIPS",
})

-- Teeth
-------------------------------------------------

teeth_generic = tables.make("part_meta", {
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_BITE_TOOTH:NO_VENOM]",
	},
	
	parts = "FLR_TEETH_GENERIC",
})

teeth_large_eye_teeth = tables.make("part_meta", {
	desc_desc = {"prominent eye teeth"},
	
	custom = {
		predator = true,
	},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_BITE_TOOTH:NO_VENOM]",
	},
	
	parts = "FLR_TEETH_LARGE_EYE_TEETH",
})

teeth_large_eye_teeth_venom = tables.make("part_meta", {
	desc_desc = {"prominent eye teeth, dripping with venom"},
	desc_adj = {"venomous"},
	
	custom = {
		predator = true,
		venomous = true,
	},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_BITE_TOOTH:VENOM]",
	},
	
	parts = "FLR_TEETH_LARGE_EYE_TEETH",
})

teeth_fangs = tables.make("part_meta", {
	desc_desc = {"large fangs"},
	desc_adj = {"fanged"},
	
	custom = {
		predator = true,
	},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_BITE_TOOTH:NO_VENOM]",
	},
	
	parts = "FLR_TEETH_FANGS",
})

teeth_fangs_venom = tables.make("part_meta", {
	desc_desc = {"venomous fangs"},
	desc_adj = {"venomous"},
	
	custom = {
		predator = true,
		venomous = true,
	},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_BITE_TOOTH:VENOM]",
	},
	
	parts = "FLR_TEETH_FANGS",
})

-- These should only be used with herbivores.
teeth_rodent = tables.make("part_meta", {
	desc_desc = {"buck teeth"},
	desc_adj = {"buck-toothed"},
	desc_verb = {"constantly grinds it's teeth to sharpen them"},
	
	custom = {
		prey = true,
	},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_BITE_TOOTH:NO_VENOM]",
	},
	
	parts = "FLR_TEETH_RODENT",
})

teeth_extra_tusks = tables.make("part_meta", {
	desc_desc = {"two tusks"},
	desc_adj = {"tusked"},
	desc_verb = {"roots around with it's tusks"},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_TUSK]",
	},
	
	parts = "FLR_TEETH_EXTRA_TUSKS",
})

teeth_none = tables.make("part_meta", {
	desc_desc = {"no teeth"},
	desc_adj = {"toothless"},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_BITE_MOUTH]",
	},
})

teeth_none_sharp = tables.make("part_meta", {
	desc_desc = {"a sharp bony ridge in the place of teeth"},
	
	extras = {
		"[APPLY_CREATURE_VARIATION:FLR_ATTACK_BITE_MOUTH_EDGE]",
	},
})

teeth_all = {
	{
		weight = 1,
		parts = teeth_none,
	},
	{
		weight = 1,
		parts = teeth_none_sharp,
	},
	{
		weight = 10,
		parts = teeth_generic,
	},
	{
		weight = 3,
		parts = teeth_large_eye_teeth,
	},
	{
		weight = 1,
		parts = teeth_large_eye_teeth_venom,
	},
	{
		weight = 4,
		parts = teeth_fangs,
	},
	{
		weight = 1,
		parts = teeth_fangs_venom,
	},
	{
		weight = 2,
		parts = teeth_rodent,
	},
}

teeth_no_venom = {
	{
		weight = 1,
		parts = teeth_none,
	},
	{
		weight = 1,
		parts = teeth_none_sharp,
	},
	{
		weight = 10,
		parts = teeth_generic,
	},
	{
		weight = 3,
		parts = teeth_large_eye_teeth,
	},
	{
		weight = 4,
		parts = teeth_fangs,
	},
	{
		weight = 2,
		parts = teeth_rodent,
	},
}

-- Eyes and Eyelids
-------------------------------------------------

eyes_1 = tables.make("part_meta", {
	desc_desc = {"one eye", "no eyelids"},
	desc_adj = {"one-eyed"},
	
	parts = "FLR_EYES_1",
})

eyes_2 = tables.make("part_meta", {
	desc_desc = {"no eyelids"},
	
	parts = "FLR_EYES_2",
})

eyes_3 = tables.make("part_meta", {
	desc_desc = {"three eyes", "no eyelids"},
	desc_adj = {"three-eyed"},
	pref = {"multiple eyes"},
	
	parts = "FLR_EYES_3",
})

eyes_4 = tables.make("part_meta", {
	desc_desc = {"four eyes", "no eyelids"},
	desc_adj = {"four-eyed"},
	pref = {"many eyes"},
	
	parts = "FLR_EYES_4",
})

eyes_eyelids_1 = tables.make("part_meta", {
	desc_desc = {"one eye"},
	desc_adj = {"one-eyed"},
	
	part = {"FLR_EYES_1", "FLR_DECO_EYELIDS_1"},
})

eyes_eyelids_2 = tables.make("part_meta", {
	part = {"FLR_EYES_2", "FLR_DECO_EYELIDS_2"},
})

eyes_eyelids_3 = tables.make("part_meta", {
	desc_desc = {"three eyes"},
	desc_adj = {"three-eyed"},
	pref = {"multiple eyes"},
	
	part = {"FLR_EYES_3", "FLR_DECO_EYELIDS_3"},
})

eyes_eyelids_4 = tables.make("part_meta", {
	desc_desc = {"four eyes"},
	desc_adj = {"four-eyed"},
	pref = {"many eyes"},
	
	part = {"FLR_EYES_4", "FLR_DECO_EYELIDS_4"},
})

-- All the eye variants, weighted so they gravitate towards two eyes (no eyelids are rarer).
eyes_all = {
	{
		weight = 2,
		parts = eyes_1,
	},
	{
		weight = 3,
		parts = eyes_2,
	},
	{
		weight = 2,
		parts = eyes_3,
	},
	{
		weight = 1,
		parts = eyes_4,
	},
	{
		weight = 4,
		parts = eyes_eyelids_1,
	},
	{
		weight = 6,
		parts = eyes_eyelids_2,
	},
	{
		weight = 4,
		parts = eyes_eyelids_3,
	},
	{
		weight = 2,
		parts = eyes_eyelids_4,
	},
}

-- Ears
-------------------------------------------------

ears_generic = tables.make("part_meta", {
	parts = "FLR_EARS_GENERIC",
})

ears_none = tables.make("part_meta", {
	desc_desc = {"no ears"},
})

-- Noses
-------------------------------------------------

nose_generic = tables.make("part_meta", {
	parts = "FLR_NOSE_GENERIC",
})

nose_none = tables.make("part_meta", {
	desc_desc = {"no nose"},
	desc_adj = {"flat-faced"},
})

nose_snout = tables.make("part_meta", {
	desc_desc = {"a flat snout"},
	desc_verb = {"roots around with it's snout"},
	
	parts = "FLR_NOSE_SNOUT",
})

nose_trunk = tables.make("part_meta", {
	desc_desc = {"a prehensile trunk"},
	desc_verb = {"waves its trunk in the air"},
	
	parts = "FLR_NOSE_TRUNK",
})

nose_all = {
	{
		weight = 5,
		parts = nose_generic,
	},
	{
		weight = 1,
		parts = nose_none,
	},
	{
		weight = 2,
		parts = nose_snout,
	},
	{
		weight = 1,
		parts = nose_trunk,
	},
}

-- General Deco Parts
-------------------------------------------------

deco_antennae = tables.make("part_meta", {
	desc_desc = {"a pair of antennae"},
	desc_verb = {"rubs it's antennae against everything it passes"},
	
	parts = "FLR_DECO_ANTENNAE",
})

deco_horns_1 = tables.make("part_meta", {
	desc_desc = {"a single horn"},
	
	extras = {
		"[USE_MATERIAL_TEMPLATE:HORN:HORN_TEMPLATE]",
		"[USE_TISSUE_TEMPLATE:HORN:HORN_TEMPLATE]",
	},
	
	parts = "FLR_DECO_HORNS_1",
})

deco_horns_2 = tables.make("part_meta", {
	desc_desc = {"a pair of horns"},
	
	extras = {
		"[USE_MATERIAL_TEMPLATE:HORN:HORN_TEMPLATE]",
		"[USE_TISSUE_TEMPLATE:HORN:HORN_TEMPLATE]",
	},
	
	parts = "FLR_DECO_HORNS_2",
})

deco_horns_3 = tables.make("part_meta", {
	desc_desc = {"three horns"},
	
	extras = {
		"[USE_MATERIAL_TEMPLATE:HORN:HORN_TEMPLATE]",
		"[USE_TISSUE_TEMPLATE:HORN:HORN_TEMPLATE]",
	},
	
	parts = "FLR_DECO_HORNS_3",
})

deco_horns_4 = tables.make("part_meta", {
	desc_desc = {"four horns"},
	
	extras = {
		"[USE_MATERIAL_TEMPLATE:HORN:HORN_TEMPLATE]",
		"[USE_TISSUE_TEMPLATE:HORN:HORN_TEMPLATE]",
	},
	
	parts = "FLR_DECO_HORNS_4",
})

deco_antlers_2 = tables.make("part_meta", {
	desc_desc = {"a pair of antlers"},
	
	extras = {
		"[USE_MATERIAL_TEMPLATE:HORN:HORN_TEMPLATE]",
		"[USE_TISSUE_TEMPLATE:HORN:HORN_TEMPLATE]",
	},
	
	parts = "FLR_DECO_ANTLERS_2",
})

deco_horns_all = {
	{
		weight = 1,
		parts = deco_horns_1,
	},
	{
		weight = 4,
		parts = deco_horns_2,
	},
	{
		weight = 3,
		parts = deco_horns_3,
	},
	{
		weight = 2,
		parts = deco_horns_4,
	},
}

return _ENV
