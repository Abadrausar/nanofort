
local pwshops = rubble.require "workshops_libs_powered"
local pitems = rubble.require "items_libs_powered"
local pfilter = rubble.require "filter_libs_powered"
local pdisaster = rubble.require "disaster_libs_powered"
local preact = rubble.require "reaction_libs_powered"

local melt = rubble.require "libs_melt_item"

-- Make coke/charcoal

preact.AddRecipe("DFHACK_POWERED_FL_AUTOFURNACE", {
	validate = pfilter.Item("BLOCKS:NONE", pfilter.RC("WOOD_MAT")),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local item = pitems.CreateItem(dfhack.matinfo.find("COAL:CHARCOAL"), 'item_barst', nil, 0)
		item:setDimension(150)
		pitems.Eject(wshop, item)
	end,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOFURNACE", {
	validate = pfilter.Item("BOULDER:NONE", pfilter.Mat("INORGANIC:COAL")),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local mat = dfhack.matinfo.find("COAL:COKE")
		
		local makebar = function()
			local item = pitems.CreateItem(mat, 'item_barst', nil, 0)
			item:setDimension(150)
			pitems.Eject(wshop, item)
		end
		
		makebar()
		if math.random(1, 100) <= 75 then makebar() end
		if math.random(1, 100) <= 50 then makebar() end
		if math.random(1, 100) <= 50 then makebar() end
		if math.random(1, 100) <= 50 then makebar() end
		if math.random(1, 100) <= 50 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
	end,
})

-- Smelt ores

preact.AddRecipe("DFHACK_POWERED_FL_AUTOFURNACE", {
	validate = pfilter.Item("BOULDER:NONE", pfilter.MRP("ORE_HQ_MAT")),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local mat = pitems.GetMRP("ORE_HQ_MAT", dfhack.matinfo.decode(items[1]))
		
		local makebar = function()
			local item = pitems.CreateItem(mat, 'item_barst', nil, 0)
			item:setDimension(150)
			pitems.Eject(wshop, item)
		end
		
		makebar()
		if math.random(1, 100) <= 50 then makebar() end
		if math.random(1, 100) <= 50 then makebar() end
		if math.random(1, 100) <= 50 then makebar() end
		if math.random(1, 100) <= 50 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
	end,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOFURNACE", {
	validate = pfilter.Item("BOULDER:NONE", pfilter.MRP("ORE_LQ_MAT")),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local mat = pitems.GetMRP("ORE_LQ_MAT", dfhack.matinfo.decode(items[1]))
		
		local makebar = function()
			local item = pitems.CreateItem(mat, 'item_barst', nil, 0)
			item:setDimension(150)
			pitems.Eject(wshop, item)
		end
		
		if math.random(1, 100) <= 50 then makebar() end
		if math.random(1, 100) <= 50 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
		if math.random(1, 100) <= 25 then makebar() end
	end,
})

-- Make green glass

local sanditem = function(item)
	return item:isSand()
end

preact.AddRecipe("DFHACK_POWERED_FL_AUTOFURNACE", {
	validate = pfilter.Contains(sanditem),
	input_item = function(items, wshop, extras)
		local sand = pitems.FindItemIn(items[1], sanditem)
		dfhack.items.remove(sand)
		pitems.Eject(wshop, items[1])
	end,
	output_item = function(items, wshop, extras)
		local glass = dfhack.matinfo.find("GLASS_GREEN:NONE")
		local item = pitems.CreateItem(glass, 'item_roughst', nil, 0)
		pitems.Eject(wshop, item)
	end,
})

-- Fire pottery

preact.AddRecipe("DFHACK_POWERED_FL_AUTOFURNACE", {
	validate = pfilter.MRP("KILN_FIRED_MAT"),
	input_item = preact.DoNothing,
	output_item = function(items, wshop, extras)
		local item = items[1]
		
		local mat = pitems.GetMRP("KILN_FIRED_MAT", dfhack.matinfo.decode(item))
		
		-- This *should* take care of the art image and the like
		-- (or at least it does for the earlier drying steps)
		-- but if the item is improved that appears to not work.
		item:setMaterial(mat.type)
		item:setMaterialIndex(mat.index)
		
		-- Fire any fire-able improvements
		if item.improvements ~= nil then
			for _, v in pairs(item.improvements) do
				local imat = pitems.GetMRP("KILN_FIRED_MAT", dfhack.matinfo.decode(v.mat_type, v.mat_index))
				if imat ~= nil then
					v.mat_type = imat.type
					v.mat_index = imat.index
				end
			end
		end
		
		-- Statues and the like need to have their image fired too
		-- I do lots of nil checks here, mostly because if I don't the stupid thing crashes.
		if item.image ~= nil and item.image.id ~= -1 then
			local chunk = df.art_image_chunk.find(item.image.id)
			if chunk ~= nil then
				local img = chunk.images[item.image.subid]
				if img ~= nil then
					local imat = pitems.GetMRP("KILN_FIRED_MAT", dfhack.matinfo.decode(img.mat_type, img.mat_index))
					if imat ~= nil then
						img.mat_type = imat.type
						img.mat_index = imat.index
					end
				end
			end
		end
		pitems.Eject(wshop, item)
	end,
})

-- Recycle metal and glass

preact.AddRecipe("DFHACK_POWERED_FL_AUTOFURNACE", {
	validate = pfilter.MatFlag("IS_METAL"),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local bars = melt.meltMetalItem(items[0])
		for _, bar in ipairs(bars) do
			pitems.Eject(wshop, bar)
		end
	end,
})

preact.AddRecipe("DFHACK_POWERED_FL_AUTOFURNACE", {
	validate = pfilter.MatFlag("IS_GLASS"),
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local glass = dfhack.matinfo.decode(items[1])
		local item = pitems.CreateItem(glass, 'item_roughst', nil, 0)
		pitems.Eject(wshop, item)
	end,
})

-- Organic items to ash

preact.AddRecipe("DFHACK_POWERED_FL_AUTOFURNACE", {
	validate = function(item)
		local mat = dfhack.matinfo.decode(item)
		if mat == nil then
			return false
		end
		
		return mat.mode == "creature" or mat.mode == "plant"
	end,
	input_item = preact.DestroyInputs,
	output_item = function(items, wshop, extras)
		local item = pitems.CreateItem(dfhack.matinfo.find("ASH:NONE"), 'item_barst', nil, 0)
		item:setDimension(150)
		pitems.Eject(wshop, item)
	end,
})

pwshops.Register("DFHACK_POWERED_FL_AUTOFURNACE", nil, 40, 0, 500, preact.MakeHandler("DFHACK_POWERED_FL_AUTOFURNACE", {
	mangle_item = pdisaster.PassItem,
	
	mangle_unit = pdisaster.MangleCreature,
}), 5, 5)
