
local eventful = require 'plugins.eventful'
local buildings = require 'plugins.building-hacks'
local timeout = rubble.require 'libs_timeout'
local build_list = rubble.require 'libs_change_build_list'
local fluids = rubble.require "libs_fluids"

local pwshops = rubble.require "workshops_libs_powered"
local pitems = rubble.require "items_libs_powered"
local pfilter = rubble.require "filter_libs_powered"
local ppersist = rubble.require "persist_libs_powered"

local auto_worker = rubble.require "libs_auto_worker"

-- Disable artifacts
----------------------------------------------------------------------------------------------------

artifacts_state = artifacts_state or df.global.d_init.flags4.ARTIFACTS

function onUnload()
	df.global.d_init.flags4.ARTIFACTS = artifacts_state
end

df.global.d_init.flags4.ARTIFACTS = false

-- "Hive" tending
----------------------------------------------------------------------------------------------------

build_list.ChangeBuildingAdv(df.building_type.Hive, -1, -1, "MAIN_PAGE", false)

auto_worker.Register("BEETLE_HIVE", {
	["BEETLE_HIVE_TEND"] = function(wshop)
		return math.random(100) <= 20
	end,
	["BEETLE_HIVE_SPLIT"] = function(wshop)
		return math.random(1000) <= 1
	end,
}, 5000)

auto_worker.Register("ALGAE_TANK", {
	["ALGAE_TANK_TEND"] = function(wshop)
		return math.random(100) <= 40
	end,
}, 5000)

auto_worker.Register("MOISTURE_CONDENSER", {
	["MOISTURE_CONDENSER_TEND"] = {
		function(wshop)
			return math.random(100) <= 20
		end,
		function(job)
			local rid = -1
			for i = 0, #df.global.world.raws.reactions - 1, 1 do
				if df.global.world.raws.reactions[i].code == job.reaction_name then
					rid = i
					break
				end
			end
			
			job.job_items:insert('#', {new = true,
				flags1 = {new = true,
					empty = true,
				},
				flags3 = {new = true,
					food_storage = true,
				},
				reaction_id = rid
			})
		end,
	},
}, 5000)

-- Pottery & Plant Processing
----------------------------------------------------------------------------------------------------

pottery_reactions = {--INSERT_POTTERY_REACTIONS_HERE
}

eventful.onReactionComplete.FLHooksDefault = function(reaction, reaction_product, unit, input_items, input_reagents, output_items, call_native)
	if not call_native then
		if reaction.code == "PROCESS_PLANT_SPECIAL" then
			for _, item in ipairs(output_items) do
				local mat = dfhack.matinfo.decode(item)
				if pitems.GetMRP("TRANSFORM_MAT", mat) ~= nil then
					timeout.add("hooks_flbase|"..item.id, 5000, "hooks_flbase -mode transform -id "..item.id)
				end
			end
		elseif pottery_reactions[reaction.code] then
			for _, item in ipairs(output_items) do
				local mat = dfhack.matinfo.decode(item)
				if pitems.GetMRP("DRY_MAT", mat) ~= nil then
					timeout.add("hooks_flbase|"..item.id, 5000, "hooks_flbase -mode pottery_dry -id "..item.id)
				end
			end
		end
	end
end

-- This Lua hook will fire any item that has a reagent id that ends in "_fire" and has a
-- KILN_FIRED_MAT reaction product. Improvements and an art image (if the item has one)
-- are fired as well.
function firePottery(reaction, reaction_product, unit, in_items, in_reag, out_items, call_native)
	for x = 0, #in_items - 1, 1 do
		if string.match(in_reag[x].code, '%_fire$') then
			local item = in_items[x]
			local mat = pitems.GetMRP("KILN_FIRED_MAT", dfhack.matinfo.decode(item))
			
			if mat == nil then
				qerror "Pottery fire reaction could not find the required material, this should be impossible!"
			end
			
			-- Fire the item
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
			
			-- Statues and the like need to have their image (if they have one) fired too
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
		end
	end
end

eventful.registerReaction("SMELTER_FIRE_ITEM_FUEL", fire)
eventful.registerReaction("SMELTER_FIRE_ITEM_ARC", fire)

-- Sawmill
----------------------------------------------------------------------------------------------------

buildings.registerBuilding{
	name = "SAWMILL",
	consume = 10,
	gears = {{x = 0, y = 1}},
	animate = {
		isMechanical = true,
		frames = {
			{
				{x = 0, y = 1, 15, 7,0,0},
			},
			{
				{x = 0, y = 1, 42, 7,0,0},
			},
		}
	},
}

-- Metal working and smelting
----------------------------------------------------------------------------------------------------

build_list.ChangeBuilding("SMELTER", "FURNACES", false)
build_list.ChangeBuilding("MAGMA_SMELTER", "FURNACES", false)

build_list.ChangeBuilding("METALSMITHSFORGE", "WORKSHOPS", false)
build_list.ChangeBuilding("MAGMAFORGE", "WORKSHOPS", false)

if build_list.IsEntityPermitted("FUEL_SMELTER") then
	build_list.ChangeBuilding("FUEL_SMELTER", "FURNACES", true, "CUSTOM_F")
	build_list.ChangeBuilding("FUEL_SMELTER", "WORKSHOPS", false)
end

if build_list.IsEntityPermitted("ARC_SMELTER") then
	build_list.ChangeBuilding("ARC_SMELTER", "FURNACES", true, "CUSTOM_A")
	build_list.ChangeBuilding("ARC_SMELTER", "WORKSHOPS", false)
end

if build_list.IsEntityPermitted("ROCK_BREAKER_MANUAL") then
	build_list.ChangeBuilding("ROCK_BREAKER_MANUAL", "FURNACES", true, "CUSTOM_R")
	build_list.ChangeBuilding("ROCK_BREAKER_MANUAL", "WORKSHOPS", false)
end

buildings.registerBuilding{
	name = "MACHINE_SHOP_BASIC",
	consume = 10,
	gears = {{x = 0, y = 0}, {x = 2, y = 0}},
	animate = {
		isMechanical = true,
		frames = {
			{
				{x = 0, y = 0, 15, 7,0,0},
				{x = 2, y = 0, 15, 7,0,0},
			},
			{
				{x = 0, y = 0, 42, 7,0,0},
				{x = 2, y = 0, 42, 7,0,0},
			},
		},
	},
}

-- Power Plant
----------------------------------------------------------------------------------------------------

-- ;)
local oilfilter = pfilter.Item("LIQUID_MISC:NONE", pfilter.MRP("SOAP_MAT"))

function calcPower()
	return function(wshop)
		local output = ppersist.GetOutputType(wshop)
		output = tonumber(output) or 0
		output = output - 1
		if output > 0 then
			ppersist.SetOutputType(wshop, ""..output)
			return
		end
		
		-- Refuel!
		
		local item = pitems.FindItemAtInput(wshop, pfilter.Or{
			pfilter.Contains(oilfilter),
			pfilter.Item("BAR:NONE", pfilter.Mat("COAL:NO_MAT_GLOSS")),
			pfilter.Item("WOOD:NONE", pfilter.RC("WOOD_MAT")),
			pfilter.Item("BLOCKS:NONE", pfilter.RC("WOOD_MAT")),
		})
		if item == nil then
			-- No fuel
			ppersist.SetOutputType(wshop, "NONE")
			buildings.setPower(wshop, 0, 0)
			return
		end
		
		-- Now we either have some kind of solid fuel or a container of oil.
		local pu = 0
		if pfilter.Contains(oilfilter)(item) then
			local oil = pitems.FindItemIn(item, oilfilter)
			if not pitems.TakeFromStack(oil, 1) then
				pitems.Eject(wshop, item)
			end
			pu = 2
		elseif pfilter.Item("BAR:NONE")(item) then
			dfhack.items.remove(item)
			pu = 2
		elseif pfilter.Item("WOOD:NONE")(item) then
			dfhack.items.remove(item)
			pu = 4
		elseif pfilter.Item("BLOCKS:NONE")(item) then
			dfhack.items.remove(item)
			pu = 1
		end
		
		ppersist.SetOutputType(wshop, ""..pu)
		buildings.setPower(wshop, 30, 0)
	end
end

-- Should update the power output every 10 seconds (at 100 FPS)
pwshops.Register("DFHACK_POWERED_POWER_PLANT", nil, 0, 30, 1000, calcPower)

-- Other
----------------------------------------------------------------------------------------------------

-- The uranium refinery requires 50 power to run reactions.
pwshops.Register("URANIUM_REFINERY", nil, 50, 0, 0, nil, 5, 5, nil, nil, true)

-- The radio station requires 50 power to run reactions.
pwshops.Register("RADIO_STATION", nil, 50, 0, 0, nil, 5, 5, nil, nil, true)

build_list.ChangeBuilding("BOWYERS", "WORKSHOPS", false)

build_list.ChangeBuilding("MECHANICS", "WORKSHOPS", false)

build_list.ChangeBuilding("CRAFTSDWARFS", "WORKSHOPS", false)

build_list.ChangeBuilding("LEATHERWORKS", "WORKSHOPS", false)
build_list.ChangeBuilding("CLOTHIERS", "WORKSHOPS", false)

build_list.ChangeBuilding("MASONS", "WORKSHOPS", false)
build_list.ChangeBuilding("CARPENTERS", "WORKSHOPS", false)

build_list.ChangeBuilding("WOOD_FURNACE", "FURNACES", false)

build_list.ChangeBuilding("GLASS_FURNACE", "FURNACES", false)
build_list.ChangeBuilding("MAGMA_GLASS_FURNACE", "FURNACES", false)

if build_list.IsEntityPermitted("GLASS_BLOWER") then
	build_list.ChangeBuilding("GLASS_BLOWER", "FURNACES", true, "CUSTOM_G")
	build_list.ChangeBuilding("GLASS_BLOWER", "WORKSHOPS", false)
end

build_list.ChangeBuilding("KILN", "FURNACES", false)
build_list.ChangeBuilding("MAGMA_KILN", "FURNACES", false)

function needs1Water(reaction, reaction_product, unit, in_items, in_reag, out_items, call_native)
	call_native.value = false
	
	if fluids.eatFromArea(unit.pos.x-2, unit.pos.y-2, unit.pos.x+2, unit.pos.y+2, unit.pos.z, false, 1, 2) then
		call_native.value = true
	else
		dfhack.gui.showAnnouncement(dfhack.TranslateName(unit.name).." cancels reaction, needs water!" , COLOR_LIGHTBLUE, true)
		return
	end
end

eventful.registerReaction("CHEMIST_MAKE_SULFURIC_ACID", needs1Water)
eventful.registerReaction("CHEMIST_MAKE_GUNPOWDER", needs1Water)
eventful.registerReaction("CHEMIST_MAKE_PLASTIC", needs1Water)
eventful.registerReaction("CHEMIST_MAKE_FUEL_OIL", needs1Water)

eventful.registerReaction("REFINE_URANIUM", needs1Water)

pwshops.Register("PLASTICS_FACTORY", nil, 15, 0, 0, nil, 5, 5, nil, nil, true)

-- Add hard-coded jobs to user defined workshops.
eventful.postWorkshopFillSidebarMenu.FirstLandingBase = function(workshop)
	if workshop:getType() ~= df.building_type.Workshop then
		return
	end
	
	local ctyp = workshop:getCustomType()
	if ctyp == -1 then
		return
	end
	
	local id = df.building_def_workshopst.find(ctyp).code
	
	if id == "RESOURCE_COLLECTOR" then
		-- Add collect sand job.
		local new_button = df.interface_button_building_new_jobst:new()
		new_button.building = workshop
		new_button.job_type = df.job_type.CollectSand
		
		local wjob = df.global.ui_sidebar_menus.workshop_job
		wjob.choices_all:insert("#", new_button)
		wjob.choices_visible:insert("#", new_button)
		
		-- Add collect clay job.
		new_button = df.interface_button_building_new_jobst:new()
		new_button.building = workshop
		new_button.job_type = df.job_type.CollectClay
		
		local wjob = df.global.ui_sidebar_menus.workshop_job
		wjob.choices_all:insert("#", new_button)
		wjob.choices_visible:insert("#", new_button)
	
	elseif id == "SLAB_ENGRAVER" then
		local allUnits = df.global.world.units.active
		local deadCitizens = {}
		local deadOther = {}
		for i = 0, #allUnits - 1, 1 do
			if allUnits[i].flags1.dead == true and allUnits[i].hist_figure_id ~= -1 then
				if allUnits[i].civ_id == df.global.ui.civ_id then
					table.insert(deadCitizens, allUnits[i])
				else
					table.insert(deadOther, allUnits[i])
				end
			end
		end
		
		local mkBtn = function(unit)
			new_button = df.interface_button_building_new_jobst:new()
			new_button.building = workshop
			new_button.job_type = df.job_type.EngraveSlab
			new_button.hist_figure_id = unit.hist_figure_id
			new_button.item_type = 25971 -- ???
			new_button.unk_48 = true -- ???
			
			local wjob = df.global.ui_sidebar_menus.workshop_job
			wjob.choices_all:insert("#", new_button)
			wjob.choices_visible:insert("#", new_button)
		end
		
		for _, u in ipairs(deadCitizens) do
			mkBtn(u)
		end
		
		for _, u in ipairs(deadOther) do
			mkBtn(u)
		end
	end
end
