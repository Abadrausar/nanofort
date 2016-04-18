
local me = require "tables_planets_fauna_random_aquatic"

local random = require "libs_random"
local bodygen = require "first_landing_libs_random_bodies"
local namegen = require "first_landing_libs_random_names"

local howmany = tonumber(rubble.configvar("FL_RANDOM_AQUATIC_COUNT"))
if howmany == nil then
	rubble.print("      Config var FL_RANDOM_AQUATIC_COUNT is not a number, using the default of 50.\n")
	howmany = 50
else
	rubble.print("      Generating "..howmany.." random aquatic creatures.\n")
end

local out = "[OBJECT:CREATURE]\n"
for i = 1, howmany, 1 do
	local bodyinfo = bodygen.generate(me.body_table)
	
	local name = namegen.generate_animal({
		["Predatory"] = bodyinfo.meta.custom.predator,
		["Large"] = bodyinfo.meta.custom.size > 300000,
		["Small"] = bodyinfo.meta.custom.size < 20000,
		["Scaly"] = true,
		["Fishy"] = true,
	})
	local pref = random.select(bodyinfo.meta.pref)
	if pref == nil then
		pref = "randomness"
	end

	local ccode =
		"(#"..i.."/"..howmany..", "..
		bodyinfo.meta.custom.size.."cc, "..
		math.floor(bodyinfo.meta.custom.age / 15)..", "..
		"~"..(bodyinfo.meta.custom.age + 5)..")"
	
	out = out..
	"\n{!SHARED_CREATURE;RANDOM_AQUATIC_"..i.."_"..string.upper(name)..";\n"..
	"\t[DESCRIPTION:"..bodyinfo:description().."  "..ccode.."]\n"..
	"\t[NAME:"..name..":"..name.."s:"..name.."]\n"..
	"\t[CASTE_NAME:"..name..":STP:"..name.."]\n"..
	"\t[CREATURE_TILE:'"..string.char(name:byte(1)).."']\n"..
	"\t[COLOR:"..random.range(1, 7)..":0:"..random.int(1).."]\n"..
	"\t\n"..
	"\t[PREFSTRING:"..pref.."]\n"..
	"\t\n"..
	"\t[STANCE_CLIMBER][SWIMS_INNATE]\n"..
	"\t[MUNDANE][NATURAL]\n"..
	"\t\n"..
	"\t[AQUATIC]\n"..
	"\t\n"..
	"\t[CREATURE_CLASS:GENERAL_POISON]\n"..
	"\t\n"..
	"\t[PET_EXOTIC]\n"..
	"\t[LARGE_ROAMING]\n"..
	"\t[FREQUENCY:"..random.range(1, 100).."]\n"..
	"\t\n"..
	"\t"..bodyinfo.meta.custom.animal_type.."\n"..
	"\t\n"..
	"\t"..bodyinfo.meta.custom.population.."\n"..
	"\t\n"..
	"\t"..bodyinfo.meta.custom.active.."\n"..
	"\t\n"..
	"\t"..bodyinfo.meta.custom.biome.."\n"..
	"\t\n"..
	"\t"..bodyinfo:body().."\n"..
	"\t\n"..
	"\t"..bodyinfo:extras().."\n"..
	"\t\n"..
	"\t"..bodyinfo.meta.custom.gait.."\n"..
	"\t\n"..
	"\t"..bodyinfo.meta.custom.size_age.."\n"..
	"\t\n"..
	"\t[HOMEOTHERM:10067]\n"..
	"\t\n"..
	"\t[BODY_APPEARANCE_MODIFIER:LENGTH:90:95:98:100:102:105:110]\n"..
	"\t[BODY_APPEARANCE_MODIFIER:HEIGHT:90:95:98:100:102:105:110]\n"..
	"\t[BODY_APPEARANCE_MODIFIER:BROADNESS:90:95:98:100:102:105:110]\n"..
	"\t\n"..
	"\t[CASTE:MALE]\n"..
	"\t\t[MALE]\n"..
	"\t\t[ORIENTATION:FEMALE:0:0:100]\n"..
	"\t\t[ORIENTATION:MALE:100:0:0]\n"..
	"\t\t[SET_BP_GROUP:BY_TYPE:LOWERBODY][BP_ADD_TYPE:GELDABLE]\n"..
	"\t[CASTE:FEMALE]\n"..
	"\t\t[FEMALE]\n"..
	"\t\t[ORIENTATION:MALE:0:0:100]\n"..
	"\t\t[ORIENTATION:FEMALE:100:0:0]\n"..
	"}\n"
end

rubble.files["creature_planets_fauna_random_aquatic.txt"].Content = out
