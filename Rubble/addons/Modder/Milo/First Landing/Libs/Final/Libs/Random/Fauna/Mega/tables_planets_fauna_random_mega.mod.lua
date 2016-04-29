
-- Body tables for basic mega beasts.

-- The body tables define almost everything about the creature and how it is constructed.

-- These tables intentionally omit insectoid parts and other parts that would be hard to handle in a
-- simple manner.

_ENV = mkmodule("tables_planets_fauna_random_mega")

local tables = require "first_landing_libs_random_bodies_tables"
local gaitgen = require "first_landing_libs_random_gaits"
local random = require "libs_random"

local core_prefabs = require "core_first_landing_libs_body_parts"
local head_prefabs = require "head_first_landing_libs_body_parts"
local arms_prefabs = require "arms_first_landing_libs_body_parts"
local legs_prefabs = require "legs_first_landing_libs_body_parts"

body_table = tables.make("body_set", {
	parts_body = {
		tables.make("part_body", {
			weight = 5,
			parts = {
				core_prefabs.body_generic,
				{
					{
						weight = 5,
						parts = core_prefabs.ribs_generic,
					},
					{
						weight = 2,
						parts = core_prefabs.ribs_external,
					},
				},
				core_prefabs.guts_generic,
				{
					{
						weight = 7,
						parts = core_prefabs.bdp_furry,
					},
					{
						weight = 2,
						parts = core_prefabs.bdp_furry_no_head,
					},
					{
						weight = 2,
						parts = core_prefabs.bdp_feathery,
					},
					{
						weight = 1,
						parts = core_prefabs.bdp_feathery_no_head,
					},
					{
						weight = 1,
						parts = core_prefabs.bdp_leathery,
					},
					{
						weight = 5,
						parts = core_prefabs.bdp_scaly,
					},
				},
				core_prefabs.bdp_common,
			},
		}),
		tables.make("part_body", {
			weight = 2,
			parts = {
				core_prefabs.body_generic_single,
				{
					{
						weight = 5,
						parts = core_prefabs.ribs_generic,
					},
					{
						weight = 2,
						parts = core_prefabs.ribs_external,
					},
				},
				core_prefabs.guts_generic,
				{
					{
						weight = 7,
						parts = core_prefabs.bdp_furry,
					},
					{
						weight = 2,
						parts = core_prefabs.bdp_furry_no_head,
					},
					{
						weight = 2,
						parts = core_prefabs.bdp_feathery,
					},
					{
						weight = 1,
						parts = core_prefabs.bdp_feathery_no_head,
					},
					{
						weight = 1,
						parts = core_prefabs.bdp_leathery,
					},
					{
						weight = 5,
						parts = core_prefabs.bdp_scaly,
					},
				},
				core_prefabs.bdp_common,
			},
		}),
	},
	
	parts_head = {
		tables.make("part_head", {
			weight = 1,
			parts = head_prefabs.head_generic,
		}),
	},
	
	parts_face = {
		tables.make("part_face", {
			weight = 1,
			parts = {
				head_prefabs.eyes_all,
				{
					{
						weight = 2,
						parts = head_prefabs.mouth_beak,
					},
					{
						weight = 1,
						parts = head_prefabs.mouth_bill,
					},
					{
						weight = 6,
						parts = {
							head_prefabs.mouth_generic,
							{
								{
									weight = 1,
									parts = "",
								},
								{
									weight = 4,
									parts = {
										head_prefabs.teeth_all,
										{
											{
												weight = 10,
												parts = "",
											},
											{
												weight = 1,
												parts = head_prefabs.teeth_extra_tusks,
											},
										},
										head_prefabs.deco_cheeks,
										head_prefabs.deco_lips,
									},
								},
								{
									weight = 1,
									parts = {
										tables.make("part_meta", {
											desc_desc = {"exposed teeth"},
										}),
										head_prefabs.teeth_all,
										{
											{
												weight = 10,
												parts = "",
											},
											{
												weight = 1,
												parts = head_prefabs.teeth_extra_tusks,
											},
										},
									},
								},
							},
						},
					},
				},
				{
					{
						weight = 2,
						parts = head_prefabs.ears_generic,
					},
					{
						weight = 1,
						parts = head_prefabs.ears_none,
					},
				},
				head_prefabs.tongue_all,
			},
		}),
	},
	
	parts_arms = {
		tables.make("part_arms", {
			weight = 10,
			parts = "",
		}),
		tables.make("part_arms", {
			weight = 1,
			parts = {arms_prefabs.wings_all},
		}),
	},
	
	parts_legs = {
		tables.make("part_legs", {
			weight = 5,
			parts = {
				legs_prefabs.legs_quad,
				legs_prefabs.toes_all,
				legs_prefabs.nails_claws_all,
			},
		}),
		tables.make("part_legs", {
			weight = 2,
			parts = {
				legs_prefabs.legs_one_part_6_to_10,
				{
					{
						weight = 2,
						parts = "",
					},
					{
						weight = 1,
						parts = {
							legs_prefabs.toes_all,
							legs_prefabs.nails_claws_all,
						},
					},
				},
			},
		}),
	},
	
	deco_head = {
		tables.make("deco", {
			weight = 15,
			parts = "",
		}),
		tables.make("deco", {
			weight = 5,
			parts = {
				head_prefabs.deco_horns_all,
			},
		}),
		tables.make("deco", {
			weight = 1,
			parts = head_prefabs.deco_antlers_2,
		}),
		tables.make("deco", {
			weight = 1,
			parts = head_prefabs.deco_antennae,
		}),
	},
	
	deco_body = {
		tables.make("deco", {
			weight = 20,
			parts = "",
		}),
		tables.make("deco", {
			weight = 2,
			parts = core_prefabs.deco_hump,
		}),
		tables.make("deco", {
			weight = 1,
			parts = core_prefabs.deco_hump_2,
		}),
		tables.make("deco", {
			weight = 10,
			parts = core_prefabs.deco_tail,
		}),
		tables.make("deco", {
			weight = 5,
			parts = core_prefabs.deco_shell,
		}),
	},
	
	extras = {
		-- Mega vs. semi-mega
		function(gmeta)
			return tables.make("part_meta", {
				custom = {
					ismega = random.chance(25)
				},
			})
		end,
		
		-- Gait info
		function(gmeta)
			local walk = gaitgen.range(10, 80)
			local crawl = gaitgen.range(1, 10)
			local climb = gaitgen.range(1, 10)
			local swim = gaitgen.range(1, 10)
			
			local meta = tables.make("part_meta")
			
			if walk.speed < 20 then
				if gmeta.custom.winged then
					table.insert(meta.desc_verb, "flies slowly through the air")
					table.insert(meta.desc_adj, "slow moving")
				else
					table.insert(meta.desc_verb, "ambles leisurely when it need to go somewhere")
					table.insert(meta.desc_adj, "slow moving")
				end
			end
			
			if walk.speed > 60 then
				if gmeta.custom.winged then
					table.insert(meta.desc_verb, "can fly at high speeds")
					table.insert(meta.desc_adj, "fast moving")
				else
					table.insert(meta.desc_verb, "can run at high speeds")
					table.insert(meta.desc_adj, "fast moving")
				end
			end
			
			if gmeta.custom.winged then
				meta.custom.gait =
					"\t[APPLY_CREATURE_VARIATION:STANDARD_FLYING_GAITS:"..walk.gait.."\n"..
					"\t[APPLY_CREATURE_VARIATION:STANDARD_WALKING_GAITS:"..crawl.gait.."\n"..
					"\t[APPLY_CREATURE_VARIATION:STANDARD_CRAWLING_GAITS:"..crawl.gait.."\n"..
					"\t[APPLY_CREATURE_VARIATION:STANDARD_CLIMBING_GAITS:"..climb.gait.."\n"..
					"\t[APPLY_CREATURE_VARIATION:STANDARD_SWIMMING_GAITS:"..swim.gait
			else
				meta.custom.gait =
					"\t[APPLY_CREATURE_VARIATION:STANDARD_BIPED_GAITS:"..walk.gait.."\n"..
					"\t[APPLY_CREATURE_VARIATION:STANDARD_CRAWLING_GAITS:"..crawl.gait.."\n"..
					"\t[APPLY_CREATURE_VARIATION:STANDARD_CLIMBING_GAITS:"..climb.gait.."\n"..
					"\t[APPLY_CREATURE_VARIATION:STANDARD_SWIMMING_GAITS:"..swim.gait
			end
			return meta
		end,
		
		-- Size, age, and value
		function(gmeta)
			local meta = tables.make("part_meta")
			
			-- / 100 for child
			-- 100 0000
			-- 200 0000
			-- 400 0000
			-- 800 0000
			local size = random.range(100, 400) * 10000
			if gmeta.custom.ismega then
				size = size * 2
			end
			meta.custom.size = size
			
			meta.custom.size_age = 
				"[BODY_SIZE:0:0:"..math.floor(size / 100).."]\n"..
				"\t[BODY_SIZE:10:0:"..size.."]\n"..
				"\t\n"..
				"\t[CHILD:10]\n"..
				"\t[MAXAGE:9000:10000]\n"..
				"\t[PETVALUE:"..math.floor(size / 75).."]"
			
			return meta
		end,
		
		-- Egglayer
		function(gmeta)
			local meta = tables.make("part_meta")
			if random.chance(10) then
				local clutch = random.range(1, 5)
				
				table.insert(meta.desc_verb, "lays eggs")
				
				meta.custom.egglayer =
					"\t\t[USE_MATERIAL_TEMPLATE:EGGS:EGG_TEMPLATE]\n"..
					"\t\t[LAYS_EGGS]\n"..
					"\t\t\t[EGG_MATERIAL:LOCAL_CREATURE_MAT:EGGS:SOLID]\n"..
					"\n\t\t\t[EGG_SIZE:"..math.floor(gmeta.custom.size / 100).."]\n"..
					"\n\t\t\t[CLUTCH_SIZE:"..clutch..":"..(clutch + random.range(1, 5)).."]\n"
			else
				meta.custom.egglayer = ""
			end
			return meta
		end,
	},
})

return _ENV
