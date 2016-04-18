-- Automatically tames any wild specimens of a specific species
--[=[
    argument
        <CREATURE_TOKEN>    The creature type to be domesticated, mandatory

    Made by Dirst for use in The Earth Strikes Back mod
	
--]=]

args = {...}

local function findRace(name)
    for k,v in pairs(df.global.world.raws.creatures.all) do
        if v.creature_id==name then
            return k
        end
    end
    qerror("Race:"..name.." not found!")
end

critter = findRace(args[1])

function checkForCritter()
    for k,v in ipairs(df.global.world.units.active) do
		if v.race == critter then 
			if v.civ_id == -1 then
				if df.global.gamemode == df.game_mode.ADVENTURE then
					v.civ_id = df.global.world.units.active[0].civ_id
				else    
					v.civ_id = df.global.ui.civ_id
					v.flags2.resident = false;
					v.flags3.body_temp_in_range = true;
					v.population_id = -1
					v.status.current_soul.unit_id = v.id

					v.animal.population.region_x = -1
					v.animal.population.region_y = -1
					v.animal.population.unk_28 = -1
					v.animal.population.population_idx = -1
					v.animal.population.depth = -1

					v.counters.soldier_mood_countdown = -1
					v.counters.death_cause = -1

					v.enemy.anon_4 = -1
					v.enemy.anon_5 = -1
					v.enemy.anon_6 = -1
				end

				v.flags1.tame = true
				v.training_level = 7 
			end
		end
	end
end

require('repeat-util').scheduleEvery('onAction',1,'ticks',checkForCritter)