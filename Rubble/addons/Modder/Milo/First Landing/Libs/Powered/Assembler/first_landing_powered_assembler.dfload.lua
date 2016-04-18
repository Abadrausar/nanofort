
local eventful = require "plugins.eventful"
local script = require 'gui.script'

local pwshops = rubble.require "workshops_libs_powered"
local pitems = rubble.require "items_libs_powered"
local pfilter = rubble.require "filter_libs_powered"
local ppersist = rubble.require "persist_libs_powered"
local pdisaster = rubble.require "disaster_libs_powered"
local preact = rubble.require "reaction_libs_powered"

local timeout = rubble.require 'libs_timeout'

alreadyAdjusting = false
function assemblerAdjust(reaction, reaction_product, unit, in_items, in_reag, out_items, call_native)
	call_native.value = false
	
	if alreadyAdjusting == false then
		alreadyAdjusting = true
		
		local opt_types = {}
		local opt_names = {}
		
		for i = 0, #reaction.products-1, 1 do
			option = reaction.products[i]
			
			local glassonly = "false"
			if option.flags.FORCE_EDGE then
				glassonly = "true"
			end
			
			table.insert(opt_types, "itype = "..option.item_type..", isubtype = "..option.item_subtype..", inputcount = "..option.count..", outcount = "..option.product_dimension..", glassonly = "..glassonly)
			table.insert(opt_names, pitems.GetItemCaption(option.item_type, option.item_subtype))
		end
		
		script.start(function()
			local choiceok, choice = script.showListPrompt('Auto-Assembler', 'Select item to produce:', COLOR_LIGHTGREEN, opt_names)
			
			local product = ""
			if choiceok then
				product = opt_types[choice]
			else
				product = "NONE"
			end
			
			local wshop = pwshops.MakeFake(unit.pos.x, unit.pos.y, unit.pos.z, 1)
			ppersist.SetOutputType(wshop, "return {"..product.."}")
			
			alreadyAdjusting = false
		end)
	end
end

local makeItem = function(blocks, wshop, extras)
	local output = ppersist.GetOutputTypeAsCode(wshop)
	if output == nil then
		return
	end
	
	local mat = dfhack.matinfo.decode(blocks[1])
	
	for i = 1, output.outcount, 1 do
		local item = pitems.CreateItemNumeric(mat, output.itype, output.isubtype, nil, 0)
		pitems.SetAutoItemQuality(wshop, item)
		pitems.Eject(wshop, item)
	end
end

local validRaw = pfilter.Or{
	pfilter.Item("BLOCKS:NONE", pfilter.RC("STONE_MAT")),
	pfilter.Item("BLOCKS:NONE", pfilter.RC("WOOD_MAT")),
	pfilter.Item("BAR:NONE", pfilter.RC("METAL_MAT")),
	pfilter.Item("BAR:NONE", pfilter.RC("PLASTIC_MAT")),
}

preact.AddRecipe("DFHACK_POWERED_FL_AUTOASSEMBLER", {
	custom = {count = 1},
	validate = validRaw,
	amount = 1,
	input_item = preact.DestroyInputs,
	output_item = makeItem,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOASSEMBLER", {
	custom = {count = 2},
	validate = validRaw,
	amount = 2,
	input_item = preact.DestroyInputs,
	output_item = makeItem,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOASSEMBLER", {
	custom = {count = 3},
	validate = validRaw,
	amount = 3,
	input_item = preact.DestroyInputs,
	output_item = makeItem,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOASSEMBLER", {
	custom = {count = 4},
	validate = validRaw,
	amount = 4,
	input_item = preact.DestroyInputs,
	output_item = makeItem,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOASSEMBLER", {
	custom = {count = 5},
	validate = validRaw,
	amount = 5,
	input_item = preact.DestroyInputs,
	output_item = makeItem,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOASSEMBLER", {
	custom = {count = 0, glass = true},
	validate = pfilter.Item("ROUGH:NONE", pfilter.MatFlag("IS_GLASS")),
	amount = 1,
	input_item = preact.DestroyInputs,
	output_item = makeItem,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOASSEMBLER", {
	custom = {count = 0},
	validate = pfilter.Item("BOULDER:NONE", pfilter.MRP("FIRED_MAT")),
	amount = 1,
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local output = ppersist.GetOutputTypeAsCode(wshop)
		if output == nil then
			return
		end
		
		local mat = pitems.GetMRP("FIRED_MAT", dfhack.matinfo.decode(items[1]))
		
		for i = 1, output.outcount, 1 do
			local item = pitems.CreateItemNumeric(mat, output.itype, output.isubtype, nil, 0)
			pitems.SetAutoItemQuality(wshop, item)
			timeout.add("hooks_default|"..item.id, 3000, "hooks_default.com -mode pottery_dry -id "..item.id)
			pitems.Eject(wshop, item)
		end
	end,
})

pwshops.Register("DFHACK_POWERED_FL_AUTOASSEMBLER", nil, 30, 0, 500, preact.MakeHandler("DFHACK_POWERED_FL_AUTOASSEMBLER", {
	-- This disables all the recipes except the one that matches the needed input count for the current product.
	validate = function(wshop, extra)
		local output = ppersist.GetOutputTypeAsCode(wshop)
		if output == nil then
			return false
		end
		
		if extra.count == 0 then
			if output.glassonly and not extra.glass then
				return false
			end
			return true
		end
		
		return extra.count == output.inputcount
	end,
	
	mangle_item = pdisaster.Switch({
		-- If the item is metal 20% chance to damage the factory and pass the item.
		{pfilter.MatFlag("IS_METAL"), pdisaster.Damage(20, pdisaster.PassItem)},
		
		-- If the item is not hard then 5% chance of damage and rot it if possible, else pass it.
		{pfilter.Not(pfilter.MatFlag("ITEMS_HARD")), pdisaster.Damage(5, pdisaster.RotItem(pdisaster.PassItem))},
		
		-- Any other (invalid) item gives a 10% chance of damage, and the item is made into a base quality product.
		{pfilter.Dummy, pdisaster.Damage(10, function(wshop, item)
			local output = ppersist.GetOutputTypeAsCode(wshop)
			if output == nil then
				return nil
			end
			
			local mat = dfhack.matinfo.decode(item)
			dfhack.items.remove(item)
			local item = pitems.CreateItemNumeric(mat, output.itype, output.isubtype, nil, 0)
			pitems.Eject(wshop, item)
			return nil
		end)},
	}),
	mangle_unit = pdisaster.MangleCreature,
}), 5, 5)

eventful.registerReaction("ADJUST_POWERED_FL_AUTOASSEMBLER", assemblerAdjust)
