# file user\noExotics\noExotics.post.rex
var foundCreature = false

(console:print "Works like old util noexotics.exe\n Offering a whole lot more of animals to buy in the embark screen\n strike today your dreamed ranch!\nThis reduces the limitations of the retired Overseer noble position\n    Replacing tags PET_EXOTIC and MOUNT_EXOTIC with PET and MOUNT for each creature\n")
(foreach [rubble:raws] block name content {
	(if (str:cmp "creature_" (str:left [name] 9)) {
		(console:print "      " [name] "\n")
		
		[rubble:raws [name] = (df:raw:walk [content] block tag {
			(if (str:cmp [tag id] "CREATURE") {
				(if (int:eq (len [tag]) 1) {
					(if (bool:not (str:cmp [tag 0] "TOAD")){ # Except the TOAD, our most precious natural resource, we must protect them ;)
						[foundCreature = true]
					}{
						[foundCreature = false]
					})
				}{
					(rubble:abort "Error: invalid param count to CREATURE raw tag in last file.")
				})
			})
			
			(if (bool:and (str:cmp [tag id] "PET_EXOTIC") [foundCreature]) {
				#[tag replace = (str:add "[PET]")]
				[tag replace = "[PET]"]
				})
			
			(if (bool:and (str:cmp [tag id] "MOUNT_EXOTIC") [foundCreature]) {
				#[tag replace = (str:add "[MOUNT]")]
				[tag replace = "[MOUNT]"]
			})
			
			#(break true) #
		})]
	})
	(break true)
})
