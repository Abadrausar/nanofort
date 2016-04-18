-- whereami script v1.01
-- Reports coordinates of the cursor if it is present, or the center of the screen if not.
-- Also lists any bookmarks set for these coordinates.
-- Version 1.01 removes the need for the "calibrate" step and harmonizes reported Z with screen display.

local bookmarks = dfhack.persistent.get_all("BOOKMARK",true)
local pos = {}
pos = copyall(df.global.cursor)
pos.z = pos.z + df.global.world.map.region_z
local name = "Cursor:"
if pos.x == -30000 then
	dm = require('gui.dwarfmode')
	name = "Center:"
	pos["x"] = df.global.window_x + math.floor(dm.getPanelLayout().map.width / 2)
	pos["y"] = df.global.window_y + math.floor(dm.getPanelLayout().map.height / 2)
	pos["z"] = df.global.window_z + df.global.world.map.region_z
end
print(name .. "  " .. pos.x .. "  " .. pos.y .. "  " .. pos.z)
if bookmarks then
	for _, bookmark in ipairs(bookmarks) do
		if bookmark.ints[1] == pos.x and bookmark.ints[2] == pos.y and bookmark.ints[3] == pos.z then
			print("Bookmarked as " .. bookmark.value)
		end
	end
end

