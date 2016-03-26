-- tesb-specific-tribute.lua
--[[

Includes elements of Putnam's building/subtype-change script

]]

-- Identify custom_type for Tribute
local tribute_type
for _,x in ipairs(df.global.world.raws.buildings.all) do
	if x.code == "TESB_TRIBUTE" then tribute_type = x.id end
end

function specificWorkshop(workshop_id,stone)
	local workshop = df.building.find(tonumber(workshop_id))
	if workshop.custom_type ~= tribute_type then
		print('Specified worshop is not a Tribute')
	else
		local tribute = "TESB_TRIBUTE_"..string.upper(stone)
		for _,x in ipairs(df.global.world.raws.buildings.all) do
			if x.code == tribute then ctype = x.id end
		end
		workshop.custom_type = ctype
	end
end

-- List of materials associated with specific Tributes
stone_list = { "ANDESITE", "BASALT", "CHALK", "CHERT", "CLAYSTONE",
	"CONGLOMERATE", "DACITE", "DIORITE", "DOLOMITE", "GABBRO", "GNEISS", 
	"GRANITE", "LIMESTONE", "MARBLE", "MUDSTONE", "PHYLLITE", "QUARTZITE",
	"RHYOLITE", "ROCK_SALT", "SANDSTONE", "SCHIST", "SHALE", "SILTSTONE",
	"SLATE" }

workshop_list = {}
	
-- Lists indexed by material ID numbers
for mat = 1, #stone_list do
	local index = dfhack.matinfo.find(stone_list[mat]).index
	workshop_list[index] = stone_list[mat]
	if workshop_list[index] == "ROCK_SALT" then
		workshop_list[index] = "ROCKSALT"
	end
end

buildCheck = require('plugins.eventful')
buildCheck.onJobCompleted.tesbTribute=function(job)
	if job.job_type == 68 then -- ConstructBuilding
		-- check that it's a generic Tribute
		local workshop = df.building.find(job.general_refs[0].building_id)
		if workshop.construction_stage == 3 and workshop.mat_type ~= -1 and not workshop.design then
			local workshop_mat = workshop_list[workshop.mat_index]  -- Will be nil if not on the list
			if workshop.custom_type == tribute_type and workshop_mat then
				workshop_mat_index = workshop.mat_index
				if workshop.contained_items[0].item.mat_index == workshop_mat_index and 
						workshop.contained_items[1].item.mat_index == workshop_mat_index and 
						workshop.contained_items[2].item.mat_index == workshop_mat_index then
					specificWorkshop(workshop.id,workshop_mat)
				end
			end
		end
	end
end
