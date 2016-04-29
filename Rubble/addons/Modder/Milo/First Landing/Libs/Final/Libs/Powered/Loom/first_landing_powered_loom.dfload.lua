
local pwshops = rubble.require "workshops_libs_powered"
local pitems = rubble.require "items_libs_powered"
local pfilter = rubble.require "filter_libs_powered"
local pdisaster = rubble.require "disaster_libs_powered"
local preact = rubble.require "reaction_libs_powered"

preact.AddRecipe("DFHACK_POWERED_FL_AUTOLOOM", {
	validate = pfilter.Item("THREAD:NONE"),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local mat = dfhack.matinfo.decode(items[1])
		
		local item = pitems.CreateItem(mat, 'item_clothst', nil, 0)
		item:setDimension(10000)
		pitems.SetAutoItemQuality(wshop, item)
		pitems.Eject(wshop, item)
	end,
})

pwshops.Register("DFHACK_POWERED_FL_AUTOLOOM", nil, 20, 0, 500, preact.MakeHandler("DFHACK_POWERED_FL_AUTOLOOM", {
	mangle_item = pdisaster.Switch({
		-- If the item is not hard then 5% chance of damage and rot it if possible, else pass it.
		{pfilter.Not(pfilter.MatFlag("ITEMS_HARD")), pdisaster.Damage(5, pdisaster.RotItem(pdisaster.PassItem))},
		
		-- Any other (invalid) item gives a 20% chance of damage, and the item is passed.
		{pfilter.Dummy, pdisaster.Damage(20, pdisaster.PassItem)},
	}),
	mangle_unit = pdisaster.MangleCreature,
}), 5, 5)
