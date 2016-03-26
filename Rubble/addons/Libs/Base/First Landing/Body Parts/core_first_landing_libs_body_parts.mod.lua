
_ENV = mkmodule("core_first_landing_libs_body_parts")

local tables = require "first_landing_libs_random_bodies_tables"

-- These are part "prefabs", used to save time and avoid repetition when making body tables.

-- The body itself
-------------------------------------------------

body_generic_single = tables.make("part_meta", {
	desc_desc = {"a blob-like body"},
	desc_adj = {"bulbous"},
	
	parts = "FLR_BODY_GENERIC_SINGLE",
})

body_generic = tables.make("part_meta", {
	parts = "FLR_BODY_GENERIC",
})

body_insect = tables.make("part_meta", {
	desc_noun = {"insectoid"},
	
	parts = "FLR_BODY_INSECT",
})

body_insect_head = tables.make("part_meta", {
	desc_noun = {"insectoid"},
	
	flags = {
		head = true,
	},
	
	parts = "FLR_BODY_INSECT_HEAD",
})

-- Extra body parts
-------------------------------------------------

ribs_generic = tables.make("part_meta", {
	extras = {"[BODY_DETAIL_PLAN:FLR_POSITIONS_RIBCAGE]"},
	
	parts = "FLR_RIBS_GENERIC",
})

ribs_external = tables.make("part_meta", {
	desc_desc = {"external ribs"},
	
	extras = {"[BODY_DETAIL_PLAN:FLR_POSITIONS_RIBCAGE]"},
	
	parts = "FLR_RIBS_EXTERNAL",
})

guts_generic = tables.make("part_meta", {
	parts = "FLR_GUTS_GENERIC",
})

-- Material/tissue sets
-------------------------------------------------

bdp_furry = tables.make("part_meta", {
	desc_desc = {"fur"},
	desc_adj = {"furry"},
	
	custom = {
		bdp = "furry",
	},
	
	extras = {
		"[BODY_DETAIL_PLAN:FLR_MATERIALS_FURRY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUES_FURRY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUE_LAYERS_VERTEBRATE:SKIN:FAT:MUSCLE:BONE:CARTILAGE]",
		"\n\t[SELECT_TISSUE_LAYER:HEART:BY_CATEGORY:HEART]",
		"\t[PLUS_TISSUE_LAYER:SKIN:BY_CATEGORY:THROAT]",
		"\t\t[TL_MAJOR_ARTERIES]",
		"[BODY_DETAIL_PLAN:FLR_COVER_HEAD:HAIR][BODY_DETAIL_PLAN:FLR_COVER_BODY:HAIR]",
		"[BODY_DETAIL_PLAN:FLR_COVER_ARMS:HAIR][BODY_DETAIL_PLAN:FLR_COVER_LEGS:HAIR]",
	},
})

bdp_furry_no_head = tables.make("part_meta", {
	desc_desc = {"fur with a naked head"},
	desc_adj = {"furry"},
	
	custom = {
		bdp = "furry",
	},
	
	extras = {
		"[BODY_DETAIL_PLAN:FLR_MATERIALS_FURRY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUES_FURRY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUE_LAYERS_VERTEBRATE:SKIN:FAT:MUSCLE:BONE:CARTILAGE]",
		"\n\t[SELECT_TISSUE_LAYER:HEART:BY_CATEGORY:HEART]",
		"\t[PLUS_TISSUE_LAYER:SKIN:BY_CATEGORY:THROAT]",
		"\t\t[TL_MAJOR_ARTERIES]",
		"[BODY_DETAIL_PLAN:FLR_COVER_BODY:HAIR]",
		"[BODY_DETAIL_PLAN:FLR_COVER_ARMS:HAIR][BODY_DETAIL_PLAN:FLR_COVER_LEGS:HAIR]",
	},
})

bdp_feathery = tables.make("part_meta", {
	desc_desc = {"feathers"},
	desc_adj = {"feathered"},
	
	custom = {
		bdp = "feathery",
	},
	
	extras = {
		"[BODY_DETAIL_PLAN:FLR_MATERIALS_FEATHERY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUES_FEATHERY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUE_LAYERS_VERTEBRATE:SKIN:FAT:MUSCLE:BONE:CARTILAGE]",
		"\n\t[SELECT_TISSUE_LAYER:HEART:BY_CATEGORY:HEART]",
		"\t[PLUS_TISSUE_LAYER:SKIN:BY_CATEGORY:THROAT]",
		"\t\t[TL_MAJOR_ARTERIES]",
		"[BODY_DETAIL_PLAN:FLR_COVER_HEAD:FEATHER][BODY_DETAIL_PLAN:FLR_COVER_BODY:FEATHER]",
		"[BODY_DETAIL_PLAN:FLR_COVER_ARMS:FEATHER][BODY_DETAIL_PLAN:FLR_COVER_LEGS:FEATHER]",
	},
})

bdp_feathery_no_head = tables.make("part_meta", {
	desc_desc = {"feathers with a naked head"},
	desc_adj = {"feathered"},
	
	custom = {
		bdp = "feathery",
	},
	
	extras = {
		"[BODY_DETAIL_PLAN:FLR_MATERIALS_FEATHERY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUES_FEATHERY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUE_LAYERS_VERTEBRATE:SKIN:FAT:MUSCLE:BONE:CARTILAGE]",
		"\n\t[SELECT_TISSUE_LAYER:HEART:BY_CATEGORY:HEART]",
		"\t[PLUS_TISSUE_LAYER:SKIN:BY_CATEGORY:THROAT]",
		"\t\t[TL_MAJOR_ARTERIES]",
		"[BODY_DETAIL_PLAN:FLR_COVER_BODY:FEATHER]",
		"[BODY_DETAIL_PLAN:FLR_COVER_ARMS:FEATHER][BODY_DETAIL_PLAN:FLR_COVER_LEGS:FEATHER]",
	},
})

bdp_feathery_no_legs = tables.make("part_meta", {
	desc_desc = {"feathers with naked legs"},
	desc_adj = {"feathered"},
	
	custom = {
		bdp = "feathery",
	},
	
	extras = {
		"[BODY_DETAIL_PLAN:FLR_MATERIALS_FEATHERY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUES_FEATHERY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUE_LAYERS_VERTEBRATE:SKIN:FAT:MUSCLE:BONE:CARTILAGE]",
		"\n\t[SELECT_TISSUE_LAYER:HEART:BY_CATEGORY:HEART]",
		"\t[PLUS_TISSUE_LAYER:SKIN:BY_CATEGORY:THROAT]",
		"\t\t[TL_MAJOR_ARTERIES]",
		"[BODY_DETAIL_PLAN:FLR_COVER_HEAD:FEATHER][BODY_DETAIL_PLAN:FLR_COVER_BODY:FEATHER]",
		"[BODY_DETAIL_PLAN:FLR_COVER_ARMS:FEATHER]",
	},
})

bdp_feathery_no_head_legs = tables.make("part_meta", {
	desc_desc = {"feathers with naked head and legs"},
	desc_adj = {"feathered"},
	
	custom = {
		bdp = "feathery",
	},
	
	extras = {
		"[BODY_DETAIL_PLAN:FLR_MATERIALS_FEATHERY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUES_FEATHERY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUE_LAYERS_VERTEBRATE:SKIN:FAT:MUSCLE:BONE:CARTILAGE]",
		"\n\t[SELECT_TISSUE_LAYER:HEART:BY_CATEGORY:HEART]",
		"\t[PLUS_TISSUE_LAYER:SKIN:BY_CATEGORY:THROAT]",
		"\t\t[TL_MAJOR_ARTERIES]",
		"[BODY_DETAIL_PLAN:FLR_COVER_BODY:FEATHER]",
		"[BODY_DETAIL_PLAN:FLR_COVER_ARMS:FEATHER]",
	},
})

bdp_leathery = tables.make("part_meta", {
	desc_desc = {"a leathery hide"},
	
	custom = {
		bdp = "leathery",
	},
	
	extras = {
		"[BODY_DETAIL_PLAN:FLR_MATERIALS_LEATHERY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUES_LEATHERY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUE_LAYERS_VERTEBRATE:SKIN:FAT:MUSCLE:BONE:CARTILAGE]",
		"\n\t[SELECT_TISSUE_LAYER:HEART:BY_CATEGORY:HEART]",
		"\t[PLUS_TISSUE_LAYER:SKIN:BY_CATEGORY:THROAT]",
		"\t\t[TL_MAJOR_ARTERIES]",
	},
})

bdp_scaly = tables.make("part_meta", {
	desc_desc = {"fine scales"},
	pref = {"scaly hide"},
	
	custom = {
		bdp = "scaly",
	},
	
	extras = {
		"[BODY_DETAIL_PLAN:FLR_MATERIALS_SCALY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUES_SCALY]",
		"[BODY_DETAIL_PLAN:FLR_TISSUE_LAYERS_VERTEBRATE:SCALE:FAT:MUSCLE:BONE:CARTILAGE]",
		"\n\t[SELECT_TISSUE_LAYER:HEART:BY_CATEGORY:HEART]",
		"\t[PLUS_TISSUE_LAYER:SCALE:BY_CATEGORY:THROAT]",
		"\t\t[TL_MAJOR_ARTERIES]",
	},
})

-- Use with insectoid body parts.
bdp_chitinous = tables.make("part_meta", {
	desc_desc = {"a hard, chitinous exoskeleton"},
	
	custom = {
		bdp = "chitinous",
	},
	
	extras = {
		"[BODY_DETAIL_PLAN:FLR_MATERIALS_CHITINOUS]",
		"[BODY_DETAIL_PLAN:FLR_TISSUES_CHITINOUS]",
		"[BODY_DETAIL_PLAN:FLR_TISSUE_LAYERS_EXOSKELETON:CHITIN:FAT:MUSCLE]",
		"\n\t[SELECT_TISSUE_LAYER:HEART:BY_CATEGORY:HEART]",
		"\t\t[TL_MAJOR_ARTERIES]",
	},
})

bdp_common = tables.make("part_meta", {
	extras = {
		"\n\t[USE_MATERIAL_TEMPLATE:SINEW:SINEW_TEMPLATE]",
		"[TENDONS:LOCAL_CREATURE_MAT:SINEW:200]",
		"[LIGAMENTS:LOCAL_CREATURE_MAT:SINEW:200]",
		"[HAS_NERVES]",
		"[USE_MATERIAL_TEMPLATE:BLOOD:BLOOD_TEMPLATE]",
		"[BLOOD:LOCAL_CREATURE_MAT:BLOOD:LIQUID]",
		"[CREATURE_CLASS:GENERAL_POISON]",
		"[GETS_WOUND_INFECTIONS]",
		"[GETS_INFECTIONS_FROM_ROT]",
		"[USE_MATERIAL_TEMPLATE:PUS:PUS_TEMPLATE]",
		"[PUS:LOCAL_CREATURE_MAT:PUS:LIQUID]\n\n",
	},
})

-- Body deco parts
-------------------------------------------------

deco_hump = tables.make("part_meta", {
	desc_desc = {"a hump on it's back"},
	
	parts = "FLR_DECO_HUMP",
})

deco_hump_2 = tables.make("part_meta", {
	desc_desc = {"two humps on it's back"},
	
	parts = "FLR_DECO_HUMP_2",
})

deco_tail = tables.make("part_meta", {
	desc_desc = {"a long tail"},
	
	parts = "FLR_DECO_TAIL",
})

deco_shell = tables.make("part_meta", {
	desc_desc = {"a large shell"},
	
	extras = {"[BODY_DETAIL_PLAN:FLR_POSITIONS_SHELL]"},
	
	parts = "FLR_DECO_SHELL",
})

return _ENV
