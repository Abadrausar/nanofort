-Kazoo´s silk eggs

Short and sweet:
-Pet spiders claim nest boxes,
-Spider egg shells are made of SILK,
-The new SILK_REEL building costs 1 build material + 1 mechanism,
-At the SILK_REEL eggs made of silk are automatically turned into 3 silk thread each,
-This is done by a WEAVER.

I also added a new domestic 'tunnel tarantula', feel free to remove it.

INSTALLATION: Just merge the raw folder.

And if you just want to make a female creature lay compatible eggs, just add the following below the [CASTE:FEMALE] tag.

[LAYS_EGGS]
			[EGG_MATERIAL:LOCAL_CREATURE_MAT:SILK:SOLID]
			[EGG_MATERIAL:LOCAL_CREATURE_MAT:EGG_WHITE:LIQUID]
			[EGG_MATERIAL:LOCAL_CREATURE_MAT:EGG_YOLK:LIQUID]
			[EGG_SIZE:500] -500=size of a kitten
			[CLUTCH_SIZE:2:3] -Will lay 2-3 eggs

The necessary tags to add to entity_default.txt for mod merging are below:

	[PERMITTED_BUILDING:SILK_REEL]
	[PERMITTED_REACTION:PREP_SILK_EGG]
	[PERMITTED_REACTION:EGG_SILK_SPIN]

-Kazoo