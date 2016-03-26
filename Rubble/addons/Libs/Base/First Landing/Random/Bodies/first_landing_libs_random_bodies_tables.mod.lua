
_ENV = mkmodule("first_landing_libs_random_bodies_tables")

-- Body part type definitions:
types = {
	-- Defaults for "part_meta" objects (used to make prefabs and for embedding in part arrays).
	-- One of these may be included in any body part array to provide extra information to the client.
	part_meta = {
		-- This is a place to store data that the client will need later.
		custom = {},
		
		-- An array of extra little bits of raw text to add after the body/materials/tissues are defined.
		-- This is where creature variations (attacks mostly) and body detail plans go.
		extras = {},
		
		-- Creature description stuff.
		-- You should define as many of these as possible, more is better.
		desc_desc = {}, -- mighty wings
		desc_verb = {}, -- runs on four spindly legs
		desc_noun = {}, -- monster
		desc_adj = {}, -- long-necked
		pref = {}, -- sharp stingers
		
		-- This is a way to attach specific body part(s) directly to the part_meta object.
		-- This is used when making "prefab" parts to use later.
		parts = {},
		
		_magic = "part_meta", -- Do not touch!
	},
	
	-- Defaults for "parts_body"
	part_body = {
		weight = 1, -- The relative commonality of this part.
		
		-- An array of body parts to use.
		-- These can get pretty complicated, look at the default tables and prefabs for examples.
		parts = {},
		
		flags = {
			head = false, -- Does this body have an integrated head?
			legs = false, -- Does this body have integrated legs?
			arms = false, -- Does this body have integrated arms?
		},
		
		_magic = "part_body", -- Do not touch!
	},
	
	-- Defaults for "parts_head".
	part_head = {
		weight = 1,
		parts = {},
		
		_magic = "part_head", -- Do not touch!
	},
	
	-- Defaults for "parts_face".
	part_face = {
		weight = 1,
		parts = {},
		
		_magic = "part_face", -- Do not touch!
	},
	
	-- Defaults for "parts_arms".
	part_arms = {
		weight = 1,
		parts = {},
		
		_magic = "part_arms", -- Do not touch!
	},
	
	-- Defaults for "parts_legs".
	part_legs = {
		weight = 1,
		parts = {},
		
		_magic = "part_legs", -- Do not touch!
	},
	
	-- Defaults for "deco_head" and "deco_body".
	deco = {
		weight = 1,
		parts = {},
		
		_magic = "deco", -- Do not touch!
	},
	
	
	-- A set of body parts for assembling creature bodies.
	body_set = {
		-- A list of all possible upper+lower body combinations.
		-- This is always the first part picked, and defines the general body structure.
		-- Any guts, etc the creature need must be included here.
		-- You may include integrated heads, legs and arms if you like, remember to set the flags if you do.
		parts_body = {
			-- Example:
			--	bodygen_tables.make("part_body", {
			--		parts = {
			--			bodygen_tables.make("part_meta", {
			--				desc_noun = {"insectoid"},
			--			})
			--			
			--			"RCP_CEPHALOTHORAX",
			--			"RCP_ABDOMEN",
			--			"RCP_BRAIN",
			--			"RCP_LUNGS",
			--			"RCP_HEART",
			--			"RCP_GUTS",
			--		},
			--		flags = {
			--			head = true,
			--		},
			--	})
		},
		
		-- A list of all possible heads, this should include a brain, skull, and neck (or whatever),
		-- but not facial features, etc.
		parts_head = {},
		
		-- Mouth, nose, eyes, ears, teeth, etc.
		-- Each entry in this list should be a complete set of facial features, use weighted sub-lists
		-- for more randomness.
		parts_face = {},
		
		-- Arms or wings, do not provide any if the creatures should not have arms or wings.
		-- If you want fingers, finger nails, etc attach them to the arms.
		parts_arms = {},
		
		-- Legs (with attached toes and toe nails if desired).
		parts_legs = {},
		
		-- Any head decorative parts that do not fit another category.
		-- Things like horns belong here.
		deco_head = {},
		
		-- Any body decorative parts that do not fit another category.
		-- Tails, stingers, shells, etc go here.
		deco_body = {},
		
		-- Each entry in this list should be a function that takes and returns a part_meta value.
		-- Do not modify the part_meta value that is passed in unless there is no other way!
		-- This is a place to put non-body items that need to effect the description and such like.
		-- These are functions so that you can take action conditional on what parts are already chosen.
		extras = {
			--function(pmeta)
			--	return tables.make("part_meta", {
			--		-- Stuff here...
			--	})
			--end,
		},
		
		_magic = "body_set", -- Do not touch!
	},
}

local metatbl = {
	__index = function(tbl, key)
		local typ = types[rawget(tbl, "_magic")]
		if typ[key] == nil then
			return nil
		end
		
		if type(typ[key]) == "table" then
			rawset(tbl, key, {})
			return rawget(tbl, key)
		else
			rawset(tbl, key, typ[key])
			return rawget(tbl, key)
		end
	end,
	__newindex = function(tbl, key, val)
		local typ = types[rawget(tbl, "_magic")]
		if typ[key] == nil then
			return
		end
		
		rawset(tbl, key, val)
	end
}

function make(typ, part)
	typ = types[typ]
	if typ == nil then
		rubble.error("Invalid part type passed to make.")
	end
	
	-- Make sure that part is valid and that it's magic ID is set.
	if part == nil then
		part = {}
	end
	part._magic = typ._magic
	
	return setmetatable(part, metatbl)
end

return _ENV
