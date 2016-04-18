
local me = require "tables_planets_fauna_random_mega"

local random = require "libs_random"
local bodygen = require "first_landing_libs_random_bodies"
local namegen = require "first_landing_libs_random_names"

local howmany = tonumber(rubble.configvar("FL_RANDOM_MEGA_COUNT"))
if howmany == nil then
	rubble.print("      Config var FL_RANDOM_MEGA_COUNT is not a number, using the default of 25.\n")
	howmany = 25
else
	rubble.print("      Generating "..howmany.." random mega and semi-mega beasts.\n")
end

local out = "[OBJECT:CREATURE]\n"
for i = 1, howmany, 1 do
	local bodyinfo = bodygen.generate(me.body_table)
	
	local name = namegen.generate_animal({
		["Predatory"] = true,
		["Flier"] = bodyinfo.meta.custom.winged,
		["Large"] = true,
		["Furry"] = bodyinfo.meta.custom.bdp == "furry",
		["Leathery"] = bodyinfo.meta.custom.bdp == "leathery",
		["Scaly"] = bodyinfo.meta.custom.bdp == "scaly",
	})
	local pref = random.select(bodyinfo.meta.pref)
	if pref == nil then
		pref = "randomness"
	end
	
	local codes = ""
	if bodyinfo.meta.custom.ismega then
		codes = codes.."M"
	end
	if bodyinfo.meta.custom.winged then
		codes = codes.."F"
	end
	if bodyinfo.meta.custom.venomous then
		codes = codes.."V"
	end
	if bodyinfo.meta.custom.egglayer ~= "" then
		codes = codes.. "E"
	end
	
	local id = string.upper(name)
	if codes ~= "" then
		id = id.."_"..codes
		codes = ", "..codes
	end
	local ccode = "(#"..i.."/"..howmany..", "..bodyinfo.meta.custom.size.."cc"..codes..")"
	
	local megastr
	if bodyinfo.meta.custom.ismega then
		megastr = "[MEGABEAST][DIFFICULTY:10]\n\t[ATTACK_TRIGGER:80:10000:100000]\n\t[BUILDINGDESTROYER:2]"
	else
		megastr = "[SEMIMEGABEAST][DIFFICULTY:5]\n\t[ATTACK_TRIGGER:50:5000:50000]\n\t[BUILDINGDESTROYER:1]"
	end
	
	local flier = ""
	if bodyinfo.meta.custom.winged then
		flier = "\t[FLIER]\n\t\n"
	end
	
	local venom = ""
	if bodyinfo.meta.custom.venomous then
		venom =
		"\t[USE_MATERIAL_TEMPLATE:VENOM:CREATURE_EXTRACT_TEMPLATE]\n"..
		"\t\t[STATE_NAME_ADJ:ALL_SOLID:frozen "..name.." venom]\n"..
		"\t\t[STATE_NAME_ADJ:LIQUID:"..name.." venom]\n"..
		"\t\t[STATE_NAME_ADJ:GAS:boiling "..name.." venom]\n"..
		"\t\t[ENTERS_BLOOD][PREFIX:NONE]\n"..
		"\t\t[SYNDROME]\n"..
		"\t\t\t[SYN_NAME:"..name.." bite]\n"..
		"\t\t\t[SYN_AFFECTED_CLASS:GENERAL_POISON]\n"..
		"\t\t\t[SYN_IMMUNE_CREATURE:RANDOM_MEGA_"..i..":ALL]\n"..
		"\t\t\t[SYN_INJECTED]\n"..
		random.select{
			"\t\t\t# GCS\n"..
			"\t\t\t[CE_PARALYSIS:SEV:100:PROB:100:RESISTABLE:SIZE_DILUTES:START:5:PEAK:10:END:20]\n",
			
			"\t\t\t# Cave Spider, but with an end\n"..
			"\t\t\t[CE_DIZZINESS:SEV:10:PROB:100:RESISTABLE:START:5:PEAK:100:END:1000]\n",
			
			"\t\t\t# Iron Man gas\n"..
			"\t\t\t[CE_COUGH_BLOOD:SEV:100:PROB:100:RESISTABLE:SIZE_DILUTES:START:5:PEAK:100:END:1000]\n",
			
			"\t\t\t# Helmet Snake\n"..
			"\t\t\t[CE_FEVER:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_NAUSEA:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_DIZZINESS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_PAIN:SEV:75:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_SWELLING:SEV:100:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_OOZING:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_BRUISING:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_BLEEDING:SEV:10:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:75:END:100]\n"..
			"\t\t\t[CE_NECROSIS:SEV:100:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:800:END:3000]\n",
			
			"\t\t\t# Cave Floater gas\n"..
			"\t\t\t[CE_FEVER:SEV:50:PROB:100:RESISTABLE:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_NAUSEA:SEV:35:PROB:100:RESISTABLE:START:50:PEAK:100:END:300]\n"..
			"\t\t\t[CE_DROWSINESS:SEV:75:PROB:100:RESISTABLE:START:1000:PEAK:2000:END:4000]\n"..
			"\t\t\t[CE_DIZZINESS:SEV:75:PROB:100:RESISTABLE:START:1000:PEAK:2000:END:3000]\n",
			
			"\t\t\t# King Cobra\n"..
			"\t\t\t[CE_PAIN:SEV:75:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:10:PEAK:50:END:1200]\n"..
			"\t\t\t[CE_DIZZINESS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:30:PEAK:100:END:1200]\n"..
			"\t\t\t[CE_DROWSINESS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:30:PEAK:100:END:1200]\n"..
			"\t\t\t[CE_PARALYSIS:SEV:100:PROB:100:RESISTABLE:SIZE_DILUTES:START:60:PEAK:100:END:1200]\n",
			
			"\t\t\t# Black Mamba\n"..
			"\t\t\t[CE_PARALYSIS:SEV:100:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_DIZZINESS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:30:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_DROWSINESS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:30:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_UNCONSCIOUSNESS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_FEVER:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_PAIN:SEV:75:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:10:PEAK:500:END:1500]\n",
			
			"\t\t\t# Bushmaster\n"..
			"\t\t\t[CE_PAIN:SEV:75:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:10:PEAK:50:END:1200]\n"..
			"\t\t\t[CE_BLEEDING:SEV:10:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:10:PEAK:30:END:50]\n"..
			"\t\t\t[CE_DIZZINESS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:100:END:1200]\n"..
			"\t\t\t[CE_UNCONSCIOUSNESS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:100:END:200]\n"..
			"\t\t\t[CE_NAUSEA:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:60:PEAK:200:END:300]\n"..
			"\t\t\t[CE_PAIN:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:100:PEAK:200:END:300]\n"..
			"\t\t\t[CE_PARALYSIS:SEV:100:PROB:100:RESISTABLE:SIZE_DILUTES:START:125:PEAK:200:END:1200]\n",
			
			"\t\t\t# Adder\n"..
			"\t\t\t[CE_NAUSEA:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:400:PEAK:500:END:1200]\n"..
			"\t\t\t[CE_PAIN:SEV:75:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:10:PEAK:50:END:2400]\n"..
			"\t\t\t[CE_SWELLING:SEV:25:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_BLISTERS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n",
			
			"\t\t\t# Rattlesnake\n"..
			"\t\t\t[CE_PAIN:SEV:75:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_BLEEDING:SEV:10:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:75:END:100]\n"..
			"\t\t\t[CE_SWELLING:SEV:100:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_BRUISING:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_BLISTERS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_NECROSIS:SEV:100:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:800:END:3000]\n"..
			"\t\t\t[CE_NAUSEA:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_DIZZINESS:SEV:50:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:500:END:1500]\n",
			
			"\t\t\t# Copperhead\n"..
			"\t\t\t[CE_PAIN:SEV:75:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_SWELLING:SEV:10:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:50:PEAK:500:END:1500]\n"..
			"\t\t\t[CE_NAUSEA:SEV:75:PROB:100:RESISTABLE:SIZE_DILUTES:START:50:PEAK:500:END:1500]\n",
			
			"\t\t\t# Gila Monster\n"..
			"\t\t\t[CE_PAIN:SEV:75:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:10:PEAK:50:END:1200]\n"..
			"\t\t\t[CE_SWELLING:SEV:10:PROB:100:RESISTABLE:SIZE_DILUTES:LOCALIZED:VASCULAR_ONLY:START:25:PEAK:50:END:1200]\n",
			
			"\t\t\t# GDS\n"..
			"\t\t\t[CE_NECROSIS:SEV:100:PROB:100:RESISTABLE:BP:BY_TYPE:THOUGHT:ALL:BP:BY_TYPE:NERVOUS:ALL:START:30:PEAK:60:END:1200]\n",
		}.."\t\n"
	end
	
	out = out..
	"\n{!SHARED_CREATURE;RANDOM_MEGA_"..i.."_"..id..";\n"..
	"\t[DESCRIPTION:"..bodyinfo:description().."  "..ccode.."]\n"..
	"\t[NAME:"..name..":"..name.."s:"..name.."]\n"..
	"\t[CASTE_NAME:"..name..":"..name.."s:"..name.."]\n"..
	"\t[CREATURE_TILE:'"..string.upper(string.char(name:byte(1))).."']\n"..
	"\t[COLOR:"..random.range(1, 7)..":0:"..random.int(1).."]\n"..
	"\t\n"..
	"\t[PREFSTRING:"..pref.."]\n"..
	"\t\n"..
	"\t[STANCE_CLIMBER][SWIMS_INNATE]\n"..
	"\t[FANCIFUL]\n"..
	"\t[ALL_ACTIVE]\n"..
	"\t[PET_EXOTIC]\n"..
	"\t\n"..
	"\t[CREATURE_CLASS:GENERAL_POISON]\n"..
	"\t\n"..
	"\t[LARGE_PREDATOR]\n"..
	"\t[NOFEAR]\n"..
	"\t\n"..
	"\t"..megastr.."\n"..
	"\t\n"..
	"\t[LAIR:"..random.select{
		"SIMPLE_BURROW",
		"SIMPLE_MOUND",
		"WILDERNESS_LOCATION",
	}..":100]\n"..
	"\t\n"..
	"\t# Spheres are hardcoded for now...\n"..
	"\t[SPHERE:STRENGTH]\n"..
	"\t[SPHERE:WAR]\n"..
	"\t\n"..
	flier..
	"\t"..bodyinfo:body().."\n"..
	"\t\n"..
	venom..
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
	bodyinfo.meta.custom.egglayer..
	"}\n"
end

rubble.files["creature_planets_fauna_random_mega.txt"].Content = out
