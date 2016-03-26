Procedural generation in Dwarf Fortress (DF) and modding possibilities are both leading in their respective categories, this makes DF deep and immersive.

We should have generic variations for each one of the DF objects (Ammo, Armor, Building, Body, Body_detail_plan, Bodygloss, Creature, Variation, Entity, Item, Interaction, Language, Material, Plant, Position, Syndrome, Tissue, Tool, Trap, Unit, Weapon, World, ...)
Why? better reuse of prototypes in modding, and almost all the necessary code is already there because nothing in the contextual lexical substitutions done by the variations is specifically and inherently vinculated to the creature prototypes.
If we are going this way, a syntax modification for the variations to indicate the kind of object that is being variated is also proposed.

Some examples of what we could do implementing this proposal:
Code: [Select]
[OBJECT:VARIATION]
[VARIATION:PLANT:SPROUTING]
	[REMOVE_TAG:NAME]
	[REMOVE_TAG:NAME_PLURAL]
	[REMOVE_TAG:ADJ]
	[CONVERT_TAG]
		[MASTER:GROWDUR]
		[TARGET:300]
		[REPLACEMENT:60]
	[CONVERT_TAG]
		[MASTER:GROWDUR]
		[TARGET:500]
		[REPLACEMENT:100]
	[CONVERT_TAG]
		[MASTER:CLUSTERSIZE]
		[TARGET:5]
		[REPLACEMENT:1]
This could be used as:
Code: [Select]
[PLANT:SPROUTING_MUSHROOM_HELMET_PLUMP]
	[COPY_TAGS_FROM:PLANT:MUSHROOM_HELMET_PLUMP]
	[APPLY_VARIATION:PLANT:SPROUTING]
	[REMOVE_TAG:SEED]
	[APPLY_CURRENT_VARIATION]
	[GO_TO_END]
	[SELECT_NESTED:ALL]
	[GO_TO_START]
        [SEED:plump helmet sproutings spawn:plump helmet sproutings spawn:4:0:1:LOCAL_PLANT_MAT:SEED]
	[ALL_NAMES:plump helmet sproutings]