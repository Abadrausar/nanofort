
-- Body tables for basic intelligent creatures.

-- The body tables define almost everything about the creature and how it is constructed.

-- These tables intentionally omit insectoid parts and other parts that would be hard to handle in a
-- simple manner.

_ENV = mkmodule("tables_planets_fauna_random_natives")

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
										head_prefabs.teeth_no_venom,
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
										head_prefabs.teeth_no_venom,
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
			weight = 5,
			parts = {
				arms_prefabs.arms_generic,
				arms_prefabs.fingers_all,
				arms_prefabs.nails_claws_all,
			},
		}),
		tables.make("part_arms", {
			weight = 2,
			parts = {
				arms_prefabs.arms_generic_4,
				arms_prefabs.fingers_all,
				arms_prefabs.nails_claws_all,
			},
		}),
	},
	
	parts_legs = {
		tables.make("part_legs", {
			weight = 5,
			parts = {
				legs_prefabs.legs_generic,
				legs_prefabs.toes_all,
				legs_prefabs.nails_claws_all,
			},
		}),
		tables.make("part_legs", {
			weight = 3,
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
		tables.make("deco", {
			weight = 15,
			parts = "",
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
		-- Gait info
		function(gmeta)
			local walk = gaitgen.range(10, 80)
			local crawl = gaitgen.range(1, 10)
			local climb = gaitgen.range(1, 10)
			
			local meta = tables.make("part_meta")
			
			if walk.speed < 20 then
				table.insert(meta.desc_verb, "ambles leisurely when it need to go somewhere")
				table.insert(meta.desc_adj, "slow moving")
			end
			
			if walk.speed > 60 then
				table.insert(meta.desc_verb, "can run at high speeds")
				table.insert(meta.desc_adj, "fast moving")
			end
			
			meta.custom.gait =
				"[APPLY_CREATURE_VARIATION:STANDARD_BIPED_GAITS:"..walk.gait.."\n"..
				"\t[APPLY_CREATURE_VARIATION:STANDARD_CRAWLING_GAITS:"..crawl.gait.."\n"..
				"\t[APPLY_CREATURE_VARIATION:STANDARD_CLIMBING_GAITS:"..climb.gait
			
			return meta
		end,
		
		-- Size and age
		function(gmeta)
			local meta = tables.make("part_meta")
			
			-- Size: 15000 - 90000
			-- / 100 for child
			local size = random.range(15, 90) * 1000
			meta.custom.size = size
			
			if size < 20000 then
				table.insert(meta.desc_adj, "small")
			elseif size > 70000 then
				table.insert(meta.desc_adj, "large")
			end
			
			-- Age: 20-100 (+ 10 for max)
			-- child is / 15
			age = random.range(3, 10) * 10
			meta.custom.age = age
			
			if age < 40 then
				table.insert(meta.desc_adj, "short-lived")
			elseif age > 60 then
				table.insert(meta.desc_adj, "long-lived")
			end
			
			meta.custom.size_age = 
				"[BODY_SIZE:0:0:"..math.floor(size / 100).."]\n"..
				"\t[BODY_SIZE:1:0:"..math.floor(size / 10).."]\n"..
				"\t[BODY_SIZE:"..math.floor(age / 10 + 5)..":0:"..size.."]\n"..
				"\t\n"..
				"\t[BABY:1]\n"..
				"\t[CHILD:"..math.floor(age / 10 + 5).."]\n"..
				"\t[MAXAGE:"..age..":"..(age + 10).."]"
			
			return meta
		end,
		
		-- Swimming
		function(gmeta)
			if random.chance(85) then
				return tables.make("part_meta", {
					custom = {
						swimming = "[SWIMS_LEARNED]\n\t[APPLY_CREATURE_VARIATION:STANDARD_SWIMMING_GAITS:"..gaitgen.range(1, 10).gait,
					},
				})
			else
				return tables.make("part_meta", {
					desc_verb = {"is a natural swimmer"},
					
					custom = {
						swimming = "[SWIMS_INNATE]\n\t[APPLY_CREATURE_VARIATION:STANDARD_SWIMMING_GAITS:"..gaitgen.range(5, 20).gait,
					},
				})
			end
		end,
	},
})

return _ENV
