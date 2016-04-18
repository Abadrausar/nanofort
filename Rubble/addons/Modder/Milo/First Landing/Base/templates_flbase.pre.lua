
-- Metal item melting

local melt = require "libs_melt_item"

local melt_changes = {
	["BARREL"] = 2,
	["BIN"] = 2,
	["AMMO"] = 0.02,
	["PIPE_SECTION"] = 2,
}

local melt_subtypes = {
	TOOL = {
		["ITEM_TOOL_HIVE"] = {bars = 2, wafers = 6},
		["ITEM_TOOL_NEST_BOX"] = {bars = 2, wafers = 6},
		
		["ITEM_TOOL_HAND_TOOLS"] = {bars = 1, wafers = 3},
		["ITEM_TOOL_MACHINE_TOOLS"] = {bars = 2, wafers = 6},
		["ITEM_TOOL_PISTON_CYLINDER"] = {bars = 2, wafers = 6},
		["ITEM_TOOL_STEAM_BOILER"] = {bars = 4, wafers = 12},
	},
	TRAPCOMP = {
		-- This is actually from an external addon, but it is easier to put
		-- production support for it in the base so that it looks seamless to the user.
		-- If I had several such items I would make an API for them, but it's not worth
		-- it for just one.
		["DFHACK_POWERED_DRILLING_RIG_BIT"] = {bars = 6, wafers = 12},
	},
}

for _, item in ipairs(melt.data) do
	if melt_changes[item.id] ~= nil then
		item.bars = melt_changes[item.id]
	end
	
	if melt_subtypes[item.id] ~= nil then
		for id, styp in pairs(melt_subtypes[item.id]) do
			if item.subtypes == nil then
				item.subtypes = {}
			end
			item.subtypes[id] = styp
		end
	end
end

table.insert(melt.reactions, "SMELTER_MELT_ITEM_FUEL")
table.insert(melt.reactions, "SMELTER_MELT_ITEM_ARC")

-- Pottery

rubble.template("!FL_POTTERY_REACTION", [[
	local id, industry, raws = rubble.targs({...}, {"", "", ""})
	
	rubble.registry["First Landing/Base:!FL_POTTERY_REACTION"]:listappend(id)
	
	return rubble.parse("{!SHARED_INDUSTRY_REACTION;"..id..";"..industry..";"..raws.."}")
]])

-- Natives

rubble.template("!FL_NATIVE_ENTITY", [[
	local id, raws = rubble.targs({...}, {"", ""})
	
	rubble.registry["First Landing/Base:FL_NATIVE_ENTITY"].table[id] = raws
	
	return "{_FL_NATIVE_ENTITY;"..id.."}"
]])

-- This is the template you would use in a planet pack to register natives with an entity.
rubble.template("!FL_REGISTER_NATIVE", [[
	local creature, id = rubble.targs({...}, {"", ""})
	
	rubble.registry["First Landing/Base:FL_NATIVE_ENTITY:"..id]:listappend(creature)
]])

rubble.template("_FL_NATIVE_ENTITY", [[
	local id = ...
	
	local raws = rubble.registry["First Landing/Base:FL_NATIVE_ENTITY"].table[id]
	local creatures = rubble.registry["First Landing/Base:FL_NATIVE_ENTITY:"..id].list
	
	if #creatures == 0 then
		return "# No valid natives."
	end
	
	raws = rubble.parse(raws)
	
	local out =  "# Native entities: "..id
	for count, creature in ipairs(creatures) do
		out = out.."\n[ENTITY:"..id.."_"..count.."]\n\t[CREATURE:"..creature.."]\n\t\n\t"..raws.."\n"
	end
	return out
]])
