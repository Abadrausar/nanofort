# file user\noExotics\noExotics.post.rex
var foundCreature = false

(console:print "Grazers have less feeding needs")
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
			
			(if (bool:and (str:cmp [tag id] "STANDARD_GRAZER") [foundCreature]) {
				# all the grazers eat the same as a rabbit
				[tag replace = "[GRAZER:120000]"]
				})
			
			#(break true) #
		})]
	})
	(break true)
})
