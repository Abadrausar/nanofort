
local pwshops = rubble.require "workshops_libs_powered"
local pitems = rubble.require "items_libs_powered"
local pfilter = rubble.require "filter_libs_powered"
local pdisaster = rubble.require "disaster_libs_powered"
local preact = rubble.require "reaction_libs_powered"

preact.AddRecipe("DFHACK_POWERED_FL_AUTOROCK_BREAKER", {
	validate = pfilter.Item("BOULDER:NONE", pfilter.MRP("ORE_REFINABLE")),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local mat = pitems.GetMRP("ORE_REFINABLE", dfhack.matinfo.decode(items[1]))
		
		items[1]:setMaterial(mat.type)
		items[1]:setMaterialIndex(mat.index)
	end,
})

pwshops.Register("DFHACK_POWERED_FL_AUTOROCK_BREAKER", nil, 40, 0, 500, preact.MakeHandler("DFHACK_POWERED_FL_AUTOROCK_BREAKER", {
	mangle_item = pdisaster.Switch({
		-- If the item is metal 20% chance to damage the factory and pass the item.
		{pfilter.MatFlag("IS_METAL"), pdisaster.Damage(20, pdisaster.PassItem)},
		
		-- If the item is not hard then 5% chance of damage and rot it if possible, else pass it.
		{pfilter.Not(pfilter.MatFlag("ITEMS_HARD")), pdisaster.Damage(5, pdisaster.RotItem(pdisaster.PassItem))},
		
		-- Any other (invalid) item gives a 10% chance of damage, and the item is passed.
		{pfilter.Dummy, pdisaster.Damage(10, pdisaster.PassItem)},
	}),
	
	mangle_unit = pdisaster.MangleCreature,
}), 5, 5)
