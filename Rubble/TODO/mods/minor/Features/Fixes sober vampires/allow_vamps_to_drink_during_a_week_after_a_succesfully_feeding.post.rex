# 
var foundCreature = false
var racesThatCanLearn = <map>
var currentCreature = ""

(console:print "Fixing sober vampires")
(foreach [rubble:raws] block name content {
	(if (str:cmp "creature_" (str:left [name] 9)) {
		(console:print "      " [name] "\n")
		
		[rubble:raws [name] = (df:raw:walk [content] block tag {
			(if (str:cmp [tag id] "CREATURE") {
				(if (int:eq (len [tag]) 1) {
					[currentCreature = [tag 0]]
					(if (bool:not (str:cmp [tag 0] "TOAD")){ # Except the TOAD, our most precious natural resource, we must protect them ;)
						[foundCreature = true]
					}{
						[foundCreature = false]
					})
				}{
					(rubble:abort "Error: invalid param count to CREATURE raw tag in last file.")
				})
			})
			
			(if (bool:and (str:cmp [tag id] "CAN_LEARN") [foundCreature]) {
				# all the grazers eat the same as a rabbit
				#[tag replace = "[GRAZER:120000]"]
				{SHARED_OBJECT_ADD;currentCreature;[APPLY_CREATURE_VARIATION:VAMPIRE_FODDER_BLOOD]}
				})
			
			#(break true) #
		})]
	})
	(break true)
})
