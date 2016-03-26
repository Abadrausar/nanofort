
-- Body tables for basic vertebrate creatures.

-- The body tables define almost everything about the creature and how it is constructed.

-- These tables intentionally omit insectoid parts and other parts that would be hard to handle in a
-- simple manner.

_ENV = mkmodule("tables_planets_fauna_random_vertebrates")

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
			parts = {
				arms_prefabs.wings_all
			},
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
			
			-- Size: 1000 - 600000 (elephants are 500000)
			-- / 100 for child
			local size = random.range(1, 600) * 1000
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
					"\t\t\t[EGG_SIZE:"..math.floor(gmeta.custom.size / 100).."]\n"..
					"\t\t\t[CLUTCH_SIZE:"..clutch..":"..(clutch + random.range(1, 5)).."]\n"
			else
				meta.custom.egglayer = ""
			end
			return meta
		end,
		
		-- Milkable
		function(gmeta)
			local meta = tables.make("part_meta")
			if gmeta.custom.egglayer ~= "" then
				return meta
			end
			
			if random.chance(15) then
				meta.custom.milkable =
					"\t\t[USE_MATERIAL_TEMPLATE:MILK:MILK_TEMPLATE]\n"..
					"\t\t\t[STATE_NAME_ADJ:ALL_SOLID:frozen %name milk]\n"..
					"\t\t\t[STATE_NAME_ADJ:LIQUID:%name milk]\n"..
					"\t\t\t[STATE_NAME_ADJ:GAS:boiling %name milk]\n"..
					"\t\t\t[PREFIX:NONE]\n"..
					"\t\t[MILKABLE:LOCAL_CREATURE_MAT:MILK:"..(random.range(10, 40) * 1000).."]\n"..
					"\t\t[USE_MATERIAL_TEMPLATE:CHEESE:CREATURE_CHEESE_TEMPLATE]\n"..
					"\t\t\t[STATE_NAME_ADJ:SOLID:%name cheese]\n"..
					"\t\t\t[STATE_NAME_ADJ:SOLID_POWDER:%name cheese powder]\n"..
					"\t\t\t[STATE_NAME_ADJ:LIQUID:melted %name cheese]\n"..
					"\t\t\t[STATE_NAME_ADJ:GAS:boiling %name cheese]\n"..
					"\t\t\t[PREFIX:NONE]\n"
			else
				meta.custom.milkable = ""
			end
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
				"[CLUSTER_NUMBER:"..cluster_min..":"..cluster_max.."]\n\t"
			
			return meta
		end,
		
		-- Biome
		function(gmeta)
			info = random.select_weighted{
				{
					weight = 2,
					biome = "[BIOME:ALL_MAIN]",
					name = "is found nearly everywhere",
				},
				{
					weight = 9,
					biome = "[BIOME:ANY_LAND]",
					name = "is found nearly everywhere",
				},
				{
					weight = 28,
					biome = "[BIOME:NOT_FREEZING]",
					name = "is found in most areas",
				},
				{
					weight = 53,
					biome = "[BIOME:ANY_TEMPERATE]",
					name = "likes temperate regions",
				},
				{
					weight = 31,
					biome = "[BIOME:ANY_TROPICAL]",
					name = "likes tropical regions",
				},
				{
					weight = 6,
					biome = "[BIOME:ANY_FOREST]",
					name = "lives in forests",
				},
				{
					weight = 9,
					biome = "[BIOME:ANY_SHRUBLAND]",
					name = "roams the shrub lands",
				},
				{
					weight = 11,
					biome = "[BIOME:ANY_GRASSLAND]",
					name = "roams the grasslands",
				},
				{
					weight = 9,
					biome = "[BIOME:ANY_SAVANNA]",
					name = "roams the savanna",
				},
				{
					weight = 38,
					biome = "[BIOME:ANY_TEMPERATE_FOREST]",
					name = "is found in temperate forests",
				},
				{
					weight = 22,
					biome = "[BIOME:ANY_TROPICAL_FOREST]",
					name = "is found in tropical forests",
				},
				{
					weight = 9,
					biome = "[BIOME:ANY_WETLAND]",
					name = "makes it's home in wetlands",
				},
				{
					weight = 4,
					biome = "[BIOME:ANY_TEMPERATE_WETLAND]",
					name = "enjoys temperate wetlands",
				},
				{
					weight = 3,
					biome = "[BIOME:ANY_TROPICAL_WETLAND]",
					name = "enjoys tropical wetlands",
				},
				{
					weight = 2,
					biome = "[BIOME:ANY_TEMPERATE_MARSH]",
					name = "is found in temperate marshes",
				},
				{
					weight = 1,
					biome = "[BIOME:ANY_TROPICAL_SWAMP]",
					name = "is found in tropical swamps",
				},
				{
					weight = 2,
					biome = "[BIOME:ANY_TEMPERATE_SWAMP]",
					name = "is found in temperate swamps",
				},
				{
					weight = 27,
					biome = "[BIOME:ANY_DESERT]",
					name = "makes it's home in the desert",
				},
				{
					weight = 17,
					biome = "[BIOME:MOUNTAIN]",
					name = "is commonly found roaming the mountains",
				},
				{
					weight = 4,
					biome = "[BIOME:GLACIER][BIOME:TUNDRA]",
					name = "likes cold climes, such as glaciers and the tundra",
				},
				{
					weight = 15,
					biome = "[BIOME:SUBTERRANEAN_WATER][UNDERGROUND_DEPTH:1:2]\n\t[LOW_LIGHT_VISION:10000]",
					name = "lives underground",
				},
				{
					weight = 15,
					biome = "[BIOME:SUBTERRANEAN_WATER][UNDERGROUND_DEPTH:2:3]\n\t[LOW_LIGHT_VISION:10000]",
					name = "lives deep underground",
				},
				{
					weight = 15,
					biome = "[BIOME:SUBTERRANEAN_WATER][UNDERGROUND_DEPTH:3:3]\n\t[LOW_LIGHT_VISION:10000]",
					name = "makes it's home far below the surface",
				},
			}
			
			return tables.make("part_meta", {
				desc_verb = {info.name},
				custom = {
					biome = info.biome,
				},
			})
		end,
		
		-- Animal Type (predator, grazer, etc)
		function(gmeta)
			local meta = tables.make("part_meta")
			
			if not gmeta.custom.predator and not gmeta.custom.prey then
				if random.chance(50) then
					gmeta.custom.predator = true
				else
					gmeta.custom.prey = true
				end
			end
			
			if gmeta.custom.predator then
				table.insert(meta.desc_verb, "will eat anything it can catch")
				meta.custom.animal_type = "[LARGE_PREDATOR][TRAINABLE]"
				
				if random.chance(25) then
					table.insert(meta.desc_verb, "tries to ambush its prey")
					meta.custom.animal_type = meta.custom.animal_type.."\n\t[AMBUSHPREDATOR]"
				end
			else
				if random.chance(50) then
					table.insert(meta.desc_verb, "peacefully grazes")
					meta.custom.animal_type = "[MEANDERER][STANDARD_GRAZER]"
					meta.custom.grazer = true
				else
					table.insert(meta.desc_verb, "wanders around looking for food")
					meta.custom.animal_type = "[MEANDERER]"
				end
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
