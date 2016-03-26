--[[
Rubble Color Palette Switcher

Copyright 2016 Milo Christiansen

This software is provided 'as-is', without any express or implied warranty. In
no event will the authors be held liable for any damages arising from the use of
this software.

Permission is granted to anyone to use this software for any purpose, including
commercial applications, and to alter it and redistribute it freely, subject to
the following restrictions:

1. The origin of this software must not be misrepresented; you must not claim
that you wrote the original software. If you use this software in a product, an
acknowledgment in the product documentation would be appreciated but is not
required.

2. Altered source versions must be plainly marked as such, and must not be
misrepresented as being the original software.

3. This notice may not be removed or altered from any source distribution.
]]

local mapingsA = {
	["BLACK"] = 0,
	["BLUE"] = 1,
	["GREEN"] = 2,
	["CYAN"] = 3,
	["RED"] = 4,
	["MAGENTA"] = 5,
	["BROWN"] = 6,
	["LGRAY"] = 7,
	["DGRAY"] = 8,
	["LBLUE"] = 9,
	["LGREEN"] = 10,
	["LCYAN"] = 11,
	["LRED"] = 12,
	["LMAGENTA"] = 13,
	["YELLOW"] = 14,
	["WHITE"] = 15,
}

local mapingsB = {
	["R"] = 0,
	["G"] = 1,
	["B"] = 2,
}

local function loadPalette(path, revert, prefix)
	if revert then
		print(prefix.."Attempting to restore default color palette...")
	else
		print(prefix.."Attempting to load new color palette from: \""..path.."\"")
	end
	
	local file, err = io.open(path, "rb")
	if file == nil then
		print(prefix.."  Load failed: "..err)
		return
	end
	
	local contents = file:read("*a")
	file:close()
	
	-- If only I could use the Rubble raw parser here... Oh well, I suppose regular expressions will do almost as well.
	for a, b, v in string.gmatch(contents, "%[([A-Z]+)_([RGB]):([0-9]+)%]") do
		local ka, kb, v = mapingsA[a], mapingsB[b], tonumber(v)
		if ka == nil or kb == nil or v == nil then
			if not revert then
				print(prefix.."  Load failed: Color file parse error.")
				loadPalette(dfhack.getDFPath().."/data/init/colors.txt", true, "  ")
			else
				print(prefix.."  Load failed: Color file parse error.")
			end
			return
		end
		
		if v == 0 then
			df.global.enabler.ccolor[ka][kb] = 0
		else
			v = v / 255
			if v > 1 then
				print(prefix.."  Warning: The "..b.." component for color "..a.." is out of range! Adjusting value.")
				v = 1
			end
			df.global.enabler.ccolor[ka][kb] = v
		end
	end
	return
end

loadPalette(dfhack.getSavePath().."/raw/colors.txt", false, "")

function OnUnload()
	loadPalette(dfhack.getDFPath().."/data/init/colors.txt", true, "")
end
