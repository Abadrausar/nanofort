
local utils = require 'utils'
local timeout = rubble.require 'libs_timeout'
local items = rubble.require "libs_items"

validArgs = validArgs or utils.invert({
 'id',
 'mode',
 'help'
})

local args = utils.processArgs({...}, validArgs)

if next(args) == nil or args.help then
	print([[Usage:

Arguments:
    -help
        print this help message
    -id itemid
        The item to transform
	-mode modeid
		Which of several possible actions to carry out
]])
	return
end

if args.id and tonumber(args.id) then
	if args.mode == "transform" then
		local item = df.item.find(args.id)
		if item == nil then
			return
		end
		
		local mat = items.GetMRP("TRANSFORM_MAT", dfhack.matinfo.decode(item))
		if mat == nil then
			return
		end
		
		item:setMaterial(mat.type)
		item:setMaterialIndex(mat.index)
		return
	elseif args.mode == "pottery_dry" then
		local item = df.item.find(args.id)
		if item == nil then
			--print("Could not find drying pottery item: "..args.id)
			return
		end
		
		local mat = items.GetMRP("DRY_MAT", dfhack.matinfo.decode(item))
		if mat == nil then
			--print("Pottery item is dry: "..args.id)
			return
		end
		
		item:setMaterial(mat.type)
		item:setMaterialIndex(mat.index)
		
		-- For some reason the art images dry by themselves here
		-- but they won't do that when firing, weird.
		-- May be caused by the item being improved when firing,
		-- I need to test firing unglazed items to see if that is the case.
		
		--print("Pottery item advancing to next dry stage: "..args.id)
		timeout.add("hooks_flbase|"..args.id, 3000, "hooks_flbase -mode pottery_dry -id "..args.id)
		return
	end
	qerror("Missing or invalid mode!")
end

qerror("Bad argument: Use -help for usage instructions.")
