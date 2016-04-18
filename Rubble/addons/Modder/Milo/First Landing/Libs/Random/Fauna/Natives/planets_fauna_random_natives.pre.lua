
local me = require "tables_planets_fauna_random_natives"

local random = require "libs_random"
local bodygen = require "first_landing_libs_random_bodies"
local namegen = require "first_landing_libs_random_names"

local howmany = tonumber(rubble.configvar("FL_RANDOM_NATIVE_COUNT"))
if howmany == nil then
	rubble.print("      Config var FL_RANDOM_NATIVE_COUNT is not a number, using the default of 5.\n")
	howmany = 5
else
	rubble.print("      Generating "..howmany.." random natives.\n")
end

local out = "[OBJECT:CREATURE]\n"
for i = 1, howmany, 1 do
	local bodyinfo = bodygen.generate(me.body_table)
	
	local name = namegen.generate(random.range(4, 10), namegen.table_greek)
	local pref = random.select(bodyinfo.meta.pref)
	if pref == nil then
		pref = "randomness"
	end
	
	out = out..
	"{!FL_REGISTER_NATIVE;RANDOM_NATIVE_"..i..";NATIVES}{!FL_REGISTER_NATIVE;RANDOM_NATIVE_"..i..";NATIVES_LL}\n"..
	"{!SHARED_CREATURE;RANDOM_NATIVE_"..i..";\n"..
	"\t[DESCRIPTION:"..bodyinfo:description().."]\n"..
	"\t[NAME:"..name..":"..name.."s:"..name.."]\n"..
	"\t[CASTE_NAME:"..name..":"..name.."s:"..name.."]\n"..
	"\t[CREATURE_TILE:'"..string.char(name:byte(1)).."']\n"..
	"\t[COLOR:"..random.range(1, 7)..":0:"..random.int(1).."]\n"..
	"\t\n"..
	"\t[LOCAL_POPS_CONTROLLABLE][LOCAL_POPS_PRODUCE_HEROES]\n"..
	"\t\n"..
	"\t[PREFSTRING:"..pref.."]\n"..
	"\t\n"..
	"\t[CAN_LEARN][CAN_SPEAK]\n"..
	"\t[CANOPENDOORS][EQUIPS]\n"..
	"\t[ALL_ACTIVE]\n"..
	"\t\n"..
	"\t[CREATURE_CLASS:GENERAL_POISON]\n"..
	"\t\n"..
	"\t"..bodyinfo:body().."\n"..
	"\t\n"..
	"\t"..bodyinfo:extras().."\n"..
	"\t\n"..
	"\t"..bodyinfo.meta.custom.gait.."\n"..
	"\t\n"..
	"\t"..bodyinfo.meta.custom.swimming.."\n"..
	"\t\n"..
	"\t"..bodyinfo.meta.custom.size_age.."\n"..
	"\t\n"..
	"\t[HOMEOTHERM:10040]\n"..
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

rubble.files["creature_planets_fauna_random_natives.txt"].Content = out
