
_ENV = mkmodule("first_landing_libs_random_bodies")

local bodygen_tables = require "first_landing_libs_random_bodies_tables"
local random = require "libs_random"

local function mergearray(a, b)
	for _, v in ipairs(b) do
		table.insert(a, v)
	end
end

local handle_part

local function addparts(parts, info)
	-- This allows you to add prefabs with random variations by delaying the creation of the
	-- parts array to the last second.
	if type(parts) == "function" then
		parts = parts()
	end
	
	-- "parts" is a single part.
	if type(parts) == "string" then
		if parts ~= "" then
			table.insert(info.body_data, parts)
		end
		return
	end
	
	-- "parts" is not a table.
	-- Probably caused by calling random:select_weighted on an empty table.
	if type(parts) ~= "table" then
		return
	end
	
	-- "parts" is a single part_meta value.
	if parts._magic == "part_meta" then
		handle_part(parts, info)
		return
	end
	
	-- "parts" is a list of parts.
	for _, part in ipairs(parts) do
		handle_part(part, info)
	end
end

handle_part = function(part, info)
	if type(part) == "table" then
		if part._magic == "part_meta" then
			-- It's a part_meta value or prefab.
			for k, v in pairs(part.custom) do
				info.meta.custom[k] = v
			end
			
			mergearray(info.meta.desc_desc, part.desc_desc)
			mergearray(info.meta.desc_verb, part.desc_verb)
			mergearray(info.meta.desc_noun, part.desc_noun)
			mergearray(info.meta.desc_adj, part.desc_adj)
			mergearray(info.meta.pref, part.pref)
			
			mergearray(info.meta.extras, part.extras)
			
			addparts(part.parts, info)
		else
			-- It's a weighted list of part objects.
			addparts(random.select_weighted(part).parts, info)
		end
		return
	end
	
	part = tostring(part)
	if part ~= "" then
		table.insert(info.body_data, part)
	end
end

-- Ideal Form:
--	This [adj], [adj] [noun] has [desc] and [desc]. It [verb] and [verb].
-- If no description elements are provided:
--	This creature lacks a proper description, disgraceful.
local function descriptionstr(info)
	local out = "This "
	
	if #info.meta.desc_adj > 0 then
		local adj1 = random.select(info.meta.desc_adj)
		local adj2 = random.select(info.meta.desc_adj)
		if adj1 == adj2 then
			out = out..adj1.." "
		else
			out = out..adj1..", "..adj2.." "
		end
	end
	
	if #info.meta.desc_noun > 0 then
		out = out..random.select(info.meta.desc_noun)
	else
		out = out.."creature"
	end
	
	if #info.meta.desc_desc > 0 then
		local desc1 = random.select(info.meta.desc_desc)
		local desc2 = random.select(info.meta.desc_desc)
		if desc1 == desc2 then
			out = out.." has "..desc1.."."
		else
			out = out.." has "..desc1.." and "..desc2.."."
		end
	else
		out = out.." lacks a proper description, disgraceful."
	end
	
	if #info.meta.desc_verb > 0 then
		local verb1 = random.select(info.meta.desc_verb)
		local verb2 = random.select(info.meta.desc_verb)
		if verb1 == verb2 then
			out = out.."  It "..verb1.."."
		else
			out = out.."  It "..verb1.." and "..verb2.."."
		end
	end
	
	return out
end

local function extrasstr(info)
	local out = ""
	
	for _, line in ipairs(info.meta.extras) do
		out = out.."\t"..line.."\n"
	end
	return string.trimspace(out)
end

local function bodystr(info)
	local out = "[BODY"
	for _, part in ipairs(info.body_data) do
		out = out..":"..part
	end
	return out.."]"
end

-- Generate a random body based on the body tables.
-- "tables" should be a table as returned by `bodygen_tables.make("body_set")`
-- Most of the magic happens in the tables, this is just a long chain of "pick an item" calls.
-- 
-- The returned table has two data keys:
--	"body_data": A list of the body parts that the creature should have.
--	"meta": A "part_meta" object constructed from those used to generate the returned body.
-- 
-- The returned table also has several methods:
--	"description": Returns a string creature description.
--	"extras": Returns extra raw text that should be inserted into the creature.
--	"body": Returns the creature's "BODY" tag, ready to insert.
function generate(tables)
	local info = {body_data = {}, meta = bodygen_tables.make("part_meta"), description = descriptionstr, extras = extrasstr, body = bodystr}
	
	if tables._magic ~= "body_set" then
		rubble.error "Invalid body set passed to generate."
	end
	
	--  Add the core and a head if needed.
	local core = random.select_weighted(tables.parts_body)
	addparts(core.parts, info)
	if not core.flags.head then
		addparts(random.select_weighted(tables.parts_head).parts, info)
	end
	
	-- Add the face.
	addparts(random.select_weighted(tables.parts_face).parts, info)
	
	-- Add head deco part(s).
	local hdeco = {}
	for _ = 1, 2, 1 do
		local idx = random.select_weighted(tables.deco_head, true)
		if not hdeco[idx] then
			addparts(tables.deco_head[idx].parts, info)
			hdeco[idx] = true
		end
	end
	
	-- Add arms (or wings) if possible.
	if not core.flags.arms then
		addparts(random.select_weighted(tables.parts_arms).parts, info)
	end
	
	-- Add legs.
	if not core.flags.legs then
		addparts(random.select_weighted(tables.parts_legs).parts, info)
	end
	
	-- Add a body deco part.
	addparts(random.select_weighted(tables.deco_body).parts, info)
	
	-- Finally, run all the extras blocks.
	for _, extra in ipairs(tables.extras) do
		handle_part(extra(info.meta), info)
	end
	
	return info
end

return _ENV
