
-- Body tables for basic aquatic creatures.

-- The body tables define almost everything about the creature and how it is constructed.

-- These tables intentionally omit insectoid parts and other parts that would be hard to handle in a
-- simple manner.

_ENV = mkmodule("tables_planets_fauna_random_aquatic")

local tables = require "first_landing_libs_random_bodies_tables"
local gaitgen = require "first_landing_libs_random_gaits"
local random = require "libs_random"

local core_prefabs = require "core_first_landing_libs_body_parts"
local head_prefabs = require "head_first_landing_libs_body_parts"
local arms_prefabs = require "arms_first_landing_libs_body_parts"
local aquatic_prefabs = require "aquatic_first_landing_libs_body_parts"

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
		tables.make("part_head", {
			weight = 10,
			parts = head_prefabs.head_no_neck,
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
						weight = 10,
						parts = {
							head_prefabs.mouth_generic,
							{
								{
									weight = 5,
									parts = {
										head_prefabs.teeth_none,
										head_prefabs.deco_cheeks,
										head_prefabs.deco_lips,
									},
								},
								{
									weight = 2,
									parts = {
										head_prefabs.teeth_none_sharp,
										head_prefabs.deco_cheeks,
										head_prefabs.deco_lips,
										tables.make("part_meta", {
											custom = {predator = true},
										}),
									},
								},
								{
									weight = 4,
									parts = {
										head_prefabs.teeth_generic,
										head_prefabs.deco_cheeks,
										head_prefabs.deco_lips,
										tables.make("part_meta", {
											custom = {predator = true},
										}),
									},
								},
								{
									weight = 1,
									parts = {
										tables.make("part_meta", {
											desc_desc = {"exposed teeth"},
										}),
										head_prefabs.teeth_generic,
									},
								},
							},
						},
					},
				},
				-- Aquatic creatures normally lack external ears.
				head_prefabs.tongue_all,
			},
		}),
	},
	
	parts_arms = {
		tables.make("part_arms", {
			weight = 1,
			parts = "",
		}),
	},
	
	parts_legs = {
		tables.make("part_legs", {
			weight = 5,
			parts = {aquatic_prefabs.legs},
		}),
	},
	
	deco_head = {
		tables.make("deco", {
			weight = 20,
			parts = "",
		}),
		tables.make("deco", {
			weight = 1,
			parts = {
				head_prefabs.deco_horns_all,
			},
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
			parts = aquatic_prefabs.fins_dorsal,
		}),
		tables.make("deco", {
			weight = 10,
			parts = aquatic_prefabs.fins_side,
		}),
		tables.make("deco", {
			weight = 10,
			parts = aquatic_prefabs.fins_both,
		}),
		tables.make("deco", {
			weight = 5,
			parts = core_prefabs.deco_shell,
		}),
	},
	
	extras = {
		-- Gait info
		function(gmeta)
			local swim = gaitgen.range(10, 80)
			
			local meta = tables.make("part_meta")
			
			if swim.speed < 20 then
				table.insert(meta.desc_verb, "swims leisurely when it need to go somewhere")
				table.insert(meta.desc_adj, "slow moving")
			end
			
			if swim.speed > 60 then
				table.insert(meta.desc_verb, "can swim at high speeds")
				table.insert(meta.desc_adj, "fast moving")
			end
			
			meta.custom.gait = "[APPLY_CREATURE_VARIATION:STANDARD_SWIMMING_GAITS:"..swim.gait
			return meta
		end,
		
		-- Size, age, and value
		function(gmeta)
			local meta = tables.make("part_meta")
			
			-- Size: 1000 - 400000 (elephants are 500000)
			-- / 100 for child
			local size = random.range(1, 400) * 1000
			meta.custom.size = size
			
			if size < 20000 then
				table.insert(meta.desc_adj, "small")
			elseif size > 300000 then
				table.insert(meta.desc_adj, "huge")
			else
				table.insert(meta.desc_adj, "large")
			end
			
			-- Age: 20-100 (+ 10 for max)
			-- child is / 15
			age = random.range(2, 10) * 10
			meta.custom.age = age
			
			if age < 30 then
				table.insert(meta.desc_adj, "short-lived")
			elseif age > 50 then
				table.insert(meta.desc_adj, "long-lived")
			end
			
			meta.custom.size_age = 
				"[BODY_SIZE:0:0:"..math.floor(size / 100).."]\n"..
				"\t[BODY_SIZE:"..math.floor(age / 15)..":0:"..size.."]\n"..
				"\t\n"..
				"\t[CHILD:"..math.floor(age / 15).."]\n"..
				"\t[MAXAGE:"..age..":"..math.floor(age + 10).."]\n"..
				"\t[PETVALUE:"..math.floor(size / 150).."]"
			
			return meta
		end,
		
		-- Population
		function(gmeta)
			local meta = tables.make("part_meta")
			
			local cluster_min = random.range(1, 5)
			local cluster_max = cluster_min + random.range(1, 5)
	
			local pop_min = cluster_min * 10
			local pop_max = cluster_max * 10
			
			if cluster_min == 1 then
				table.insert(meta.desc_verb, "tends to be solitary")
			end
			
			if cluster_min > 4 then
				table.insert(meta.desc_verb, "appears in groups")
			end
			
			meta.custom.population = 
				"[POPULATION_NUMBER:"..pop_min..":"..pop_max.."]\n\t"..
				"[CLUSTER_NUMBER:"..cluster_min..":"..cluster_max.."]"
			
			return meta
		end,
		
		-- Biome
		function(gmeta)
			local info = random.select_weighted{
				{
					weight = 15,
					biome = "[BIOME:ANY_OCEAN][BIOME:ANY_LAKE]",
					name = "is found in large bodies of water",
				},
				{
					weight = 15,
					biome = "[BIOME:ANY_OCEAN]",
					name = "makes it's home in the ocean",
				},
				{
					weight = 5,
					biome = "[BIOME:ANY_LAKE]",
					name = "lives in fresh water lakes",
				},
				{
					weight = 2,
					biome = "[BIOME:ANY_TEMPERATE_LAKE]",
					name = "lives in temperate lakes",
				},
				{
					weight = 5,
					biome = "[BIOME:ANY_RIVER]",
					name = "likes rivers",
				},
				{
					weight = 2,
					biome = "[BIOME:ANY_TEMPERATE_RIVER]",
					name = "appears in temperate rivers",
				},
			}
			
			return tables.make("part_meta", {
				desc_verb = {info.name},
				custom = {
					biome = info.biome,
				},
			})
		end,
		
		-- Animal Type
		function(gmeta)
			local meta = tables.make("part_meta")
			
			if gmeta.custom.predator then
				table.insert(meta.desc_verb, "will eat anything it can catch")
				meta.custom.animal_type = "[LARGE_PREDATOR][TRAINABLE]"
				
				if random.chance(25) then
					table.insert(meta.desc_verb, "tries to ambush its prey")
					meta.custom.animal_type = meta.custom.animal_type.."\n\t[AMBUSHPREDATOR]"
				end
			else
				meta.custom.animal_type = "[MEANDERER]"
			end
			
			if random.chance(50) then
				meta.custom.animal_type = meta.custom.animal_type.."\n\t[LOOSE_CLUSTERS]"
			end
			
			if random.chance(75) then
				meta.custom.active = "[DIURNAL]"
			else
				table.insert(meta.desc_adj, "nocturnal")
				meta.custom.active = "[NOCTURNAL]"
			end
			
			return meta
		end,
	},
})

return _ENV
