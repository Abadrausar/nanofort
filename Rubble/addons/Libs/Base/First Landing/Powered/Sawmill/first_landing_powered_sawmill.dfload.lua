
local pwshops = rubble.require "workshops_libs_powered"
local pitems = rubble.require "items_libs_powered"
local pfilter = rubble.require "filter_libs_powered"
local preact = rubble.require "reaction_libs_powered"
local pdisaster = rubble.require "disaster_libs_powered"

preact.AddRecipe("DFHACK_POWERED_FL_AUTOSAWMILL", {
	validate = pfilter.Item("BOULDER:NONE", pfilter.RC("STONE_MAT")),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras, output)
		local mat = dfhack.matinfo.decode(items[1])
		
		pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
		pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
		pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
		pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
	end,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOSAWMILL", {
	validate = pfilter.RC("STONE_MAT"),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local mat = dfhack.matinfo.decode(items[1])
		pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
	end,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOSAWMILL", {
	validate = pfilter.Item("WOOD:NONE", pfilter.RC("WOOD_MAT")),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local mat = dfhack.matinfo.decode(items[1])
		
		pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
		pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
		pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
		pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
	end,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOSAWMILL", {
	validate = pfilter.RC("WOOD_MAT"),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local mat = dfhack.matinfo.decode(items[1])
		pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
	end,
})

pwshops.Register("DFHACK_POWERED_FL_AUTOSAWMILL", nil, 30, 0, 1000, preact.MakeHandler("DFHACK_POWERED_FL_AUTOSAWMILL", {
	mangle_item = pdisaster.Switch({
		-- If the item is metal 20% chance to damage the factory and pass the item.
		{pfilter.MatFlag("IS_METAL"), pdisaster.Damage(20, pdisaster.PassItem)},
		
		-- If the item is not hard then 5% chance of damage and rot it if possible, else pass it.
		{pfilter.Not(pfilter.MatFlag("ITEMS_HARD")), pdisaster.Damage(5, pdisaster.RotItem(pdisaster.PassItem))},
		
		-- Any other (invalid) item gives a 10% chance of damage, and the item is made into a single block.
		{pfilter.Dummy, pdisaster.Damage(10, function(wshop, item)
			local mat = dfhack.matinfo.decode(item)
			dfhack.items.remove(item)
			pitems.Eject(wshop, pitems.CreateItem(mat, 'item_blocksst', nil, 0))
			return nil
		end)},
	}),
	mangle_unit = pdisaster.MangleCreature,
}), 5, 5)
