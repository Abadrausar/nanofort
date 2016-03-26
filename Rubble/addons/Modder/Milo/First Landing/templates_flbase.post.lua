
local react = ""
local data = rubble.registry["First Landing/Base:!FL_POTTERY_REACTION"].list

for _, id in ipairs(data) do
	react = react.."\n\t[\""..id.."\"] = true,"
end

local file = rubble.files["hooks_flbase.dfload.lua"]
file.Content = string.replace(file.Content, "--INSERT_POTTERY_REACTIONS_HERE", react, -1)
