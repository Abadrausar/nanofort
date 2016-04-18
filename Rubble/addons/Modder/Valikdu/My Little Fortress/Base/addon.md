[Addon Change log](/addonfile?addon=Valikdu___My_Little_Fortress&file=addon_changes.md)
![Haiku My Little Fortress](/addonfile?addon=Abadrausar___Haiku_Vanilla&file=Haiku.jpg)
## So you think that you need some help, want to do a sugestion, inform about a bug or give thanks to modders or developpers, but not sure where to go.
1. If you want to patronize Toady One and the developpment of Dwarf Fortress go to: [Support Dwarf Fortress Vanilla](http://www.bay12games.com/support.html)
2. For questions about the rubble framework go to: [Rubble framework](http://www.bay12forums.com/smf/index.php?topic=154304.0)
3. Anything about HDFPS: Humble Dwarf Fortress Publishing System is there [HDFPS](http://www.bay12forums.com/smf/index.php?topic=157300.0)
4. If your question is addressed to the actual original maintainer of this mod then you have [Haiku My Little Fortress](http://www.bay12forums.com/smf/index.php?topic=121116.0)
# Valikdu presents:
 
My Little Fortress - Friendship is Magma
Version 4.0: Endless Nights 
 
This is a My Little Pony-themed modification, influenced heavily by various works of fanfiction, such as Fallout: Equestria and Captain Hook the Biker Gorilla1. The goal has been to expand and enhance the experience of the original MLP mod.
The mod had been started as a set of auxiliary raw files for Nidokoenig’s MLP in November 2011.  

New Features: 
-Tier-based tissue material system for creatures
-Various new minerals and alloys
-New entities with new positions
-New weapons, with each entity having individual sets of weapons
-Improvements for existing large roaming creatures including new castes
-New large roaming creatures, including spouse-converting megabeasts. The Age of Myth doesn’t end!
-New trees and crop plants
-A total of 37 new buildings and 860 reactions for a wide array of new capabilities

[Before playing, check the d_init.txt file and enter the pop cap and child cap parameters that you’d like to use.
Use of advanced world generation is recommended.]  





The New Material/Tissue system 
There are four tiers of tissue material strength. Within each tier, there exist several groups with different value multipliers.
 
-Normal (Animal, Pony)
-Tainted (Animal, Rare, Exotic, Tainted Pony) – parameters x2 from normal; tainted bone or chitin is stronger than iron or bronze
-Morphling (Wild Morphling, Tyrapony, Morphling) – parameters x3 from normal; much stronger than iron or bronze
-Superior (Mythic Beast, Alicorn) – parameters x4 from normal, stronger than steel and felsteel
-Misc groups like Star Beasts and Crystal Ponies, which don’t fit elsewhere 
Roaming animals have the Tainted Animal, Rare or Exotic group, depending on their value multiplier in vanilla. Some are Feral Morphling. Megabeasts are Morphling, Mythic or Alicorn.
 
Items like leather and meat are now named <group name> <item name>, instead of <creature name> <item name> 
[due to a bug with the butchering of creatures that belong to civilizations in fort mode, Pony, Taint Pony, Morphling, Mythic and Alicorn groups have been downgraded to using materials added inside each creature’s space, and the items are named traditionally]

New Plants
The vanilla pre-0.40 plants are mostly there. Most growdur values have been increased.

Surface:
	 -Special high-value non-farmable plants: alicorn's tears (alcohol & extract), poison joke (extract), death thistle (extract)
	 -Geometric moss (growdur 2500, shaped growths are used in special reactions to make alcohol, magically receive seeds of underground plants, or create gems)
	 -Redsnow moss (growdur 200, found in winter in non-freezing areas, V1 thread)
	 -Redfrost moss (same, but found only on freezing biomes all year long)
	 -Amethyst moss (growdur 1000, V3 thread; it’s possible to spin V2 cloth from 1 amethyst and 3 red moss plants)	 
	 -Royal moss (growdur 1000, V5 drinks and leaves)
	 -Earth wheat (growdur 600; all wheat-type plants give “seedhead” growths, which then are used to get flour, drinks or seeds (threshing); V2 drink, V20 flour)
	 -Sky wheat (growdur 750; V3 drink, V30 flour)
	 -Royal wheat (growdur 1000; V4 drink, V40 flour)
   	 -Crystal berry (grows for a whole season; V10 alcohol) - special use at the Crystal Forge	 
   	 -Chromatic vine (grows for 3 seasons; V10 thread with high resilience, V10 leaf, V10 alcohol, V50 extract, fiber paper)
   	 -Moon stalk (growdur 700, V3 leaf with a 2-month sterility syndrome, V3 alcohol with a 4-month sterility syndrome
   	 -Cloud berry (growdur 400, V1 raw, V1 alcohol)
   	 -Twilight berry (growdur 600, V2 raw, V2 alcohol)
   	 -Sky bulb (growdur 750, V2 raw, V20 dye, V3 alcohol, V25 extract)
   	 -Midnight berry (growdur 850, V3 raw, V25 dye, V4 alcohol)
   	 -Weld (growdur 750, V25 dye, V20 extract)
   	 -Dirt root (growdur 300, V10 dye)
   	 -Blue madder (growdur 700, V25 dye)
   	 -Dyer's friend onion (growdur 700, V2 cooked, V25 dye)
   	 -Black hand leaf (growdur 700, V25 dye)
   	 -Wisdom bush (growdur 700, V25 dye, V8 leaf)
   	 -Sparkle root (growdur 700, V25 dye)
   	 -Tank scale (growdur 700, V25 dye)
	 -Paper sedge (papyrus)

Underground:
   	 -Special mushrooms: green cap (GD 1250, V3 alcohol), silver cap (3500, V5 alcohol), obsidian cap (5000, V10 alcohol) – used in biometal production, see Organic Refinery
   	 -Catalyst cap (growdur 600) - used in many of the new reactions 
  	 -Changeling fern (lvl 1-3, growdur 4000) - special use at the Crystal Forge
   	 -Solar cap (lvl 1-3, growdur 800, special, V3 cooked, special: V5 oil/soap, V5 leaf)
   	 -Brown cap (lvl 1-2, growdur 400, V1 raw)
   	 -Blue lichen (lvl 1-2, growdur 550, V2 cooked, V2 alcohol, V10 dye)
   	 -White lichen (lvl 1-2, growdur 500, V2 cooked, V2 thread, fiber paper)
   	 -Angel hair & blood worms (lvl 1-3, growdur 1000, any season, V2 leaf)
   	 -Cave strawberry (lvl 2-3, growdur 850, V3 raw, V5 alcohol, V35 syrup)
   	 -Fire lichen (lvl3, growdur 3000, V5 cooked, V10 alcohol, V100 powder)
   	 -Bone bush (lvl 3, growdur 1680, V4 alcohol, V50 powder)
   	 -Corpse tongue (lvl 3, growdur 1050, V3 raw, V6 alcohol)
   	 -Skull frond (lvl 3, growdur 1750, V5 alcohol, V65 extract)

New Minerals 
-Compacted Peat (cluster and small cluster in the peat and clay soil layers, 2-6 fuel per one boulder)
-Anthracite (small clusters in Ign.Int., Ign.Ext. and Met, veins in granite; 12 fuel per one boulder)
-Various precious marble (celestial, black, green, midnight, rose)
-Can now use Galena, Talc, hematite, Pitchblende, Cobaltite for glazing
-Taintstone is a new layer stone for the Ign.Int. layers.
-Befouled Adra [see High Tech] – veins everywhere
-Taint Crystal [see High Tech] – veins and clusters in taintstone, small clusters everywhere
-Mana Crystal [see High Tech] – small clusters everywhere
-Vanadinite is an ore of vanadium and lead, found in small clusters in Sed. and Ign.Ext. layers and in large formations in gabbro layers
-Star Silver Nuggets – smelted into silver, star silver and mithril, found inside silver veins
-Dark Iron Nuggets – smelted into iron and dark iron, found inside iron veins
-Taintferrite is processed for iron and taint crystals at the smelter. It’s found in veins in Ign.Int. and Sed. layers.
-Nether Sparkstone and Blood of Yogg-Saron [see High Tech] – small clusters everywhere
-Volcanic Aurum nuggets are found in Ign.Ext. and taintstone layers as small clusters
-Blood of the Eater is secretly raw adamantine

Weapon/Armor-grade metals, ascending order of effectiveness 
-Copper, bronze, taintwood, iron
-Star Silver - V17
-Noble Bronze (alloy of copper + star silver) – V9
-Volcanic Aurum (not good for cutting and armor, but very heavy) – V45
-Ironbark – V20
-Morphling bone and chitin
-Steel
-Mithril (smelted from native star silver with a low chance, slightly better than steel, very light) – V30
-Steelbark – V35
-Fel Steel (smelted using iron, green pig iron, vanadium and taint crystals at the Spirit Smelter) – V35
-Mythic bone and chitin
-Dark Iron (created at the Rainbow Smelter or smelted from nuggets, very heavy) – V40
-Fungal Steel (created at the Organic Refinery) – V40
-Alicorn bone and chitin
-Twilight Steel (smelted using dark iron, star silver and taint crystals at the Spirit Smelter) – V60
-Force Crystal (created at the Crystal Forge, V75)
-Saronite Steel (smelted from dark iron and energized saronite at the Spirit Smelter, heavy) – V80
-Starsteel (only available to the Crows) – V100
-Silksteel (created at the Organic Refinery, light) – V100
-Elementium (created with alchemy, has the characteristics of adamantine with heavy weight) – V175
-Starmetal (renamed adamantine, can also be created at the Rainbow Enchantment Altar, but now needs two thread units per wafer for normal production, which is now done at the Toxin Refinery) – V200


New Non-weapon Metal Values 
-Fungal Pewter – 3
-Fungal Brass – 5
-Fungal Silver – 7
-Vanadium – 10
-Fungal Gold – 20
-Nether Electrum and Noble Electrum [use taint crystals or star silver respectively for smelting] – 25
-Elder Electrum [uses saronite] – 30
-White Gold [gold and nickel] – 35
-Saronite – 65
-Energized Saronite – 80

Pony Redesign 
-Size of the pony creatures has been decreased, to the average of ~95K for the normal castes (alicorns and royals are bigger). Strength has been increased in order to equip armor normally, as its weight scales up with entity creatures’ median size.
-New body models (with joints)
-Standard factions (except zebra) have a Blade Pony caste.
-Standard factions have Anthro castes. Gauntlets are included into entity item sets for them.
-Zebras have the Zony caste with great technical aptitude.
-Diomedians have taint-enhanced castes for each of the old ones.
-All standard factions (except zebra) have very rare Alicorns.
-Diomedians have an entity position system taken from Drow. 
-Other factions' systems slightly altered. Manager and bookkeeper duties are now managed by the same position. New positions: assistant medic (appointed by the chief medic, can diagnose), scribe (appointed by the administrator, can account & manage) and mystic (appointed by the site leader, raises morale)
-Ethics altered to allow the processing of sapient creature materials.


The New Entities 
-The Unity
An unplayable hostile entity that will attack everyone. Their castes are:
-Several types of IMP-enhanced Changeling creatures that don’t equip anything and will instead fight with natural weapons (including scything claws). 
-Infested ponies, Morphling Controllers and chromatic alicorns: all very dangerous, if rare, combatants. They will use weapons and armor made from strong magical metals.

-The Crows
Another unplayable hostile entity. Uses the following castes:
-Giant mutant crows and ravens as cannon fodder.
-Hippocoraxes (bipedal pony-raven hybrids) and the Morrigi (dragon/bird alien creatures with strange anatomy) use weapons and armor made exclusively from Starsteel. 
 
-Chimera Ponies
Several mutant castes, all can equip gear normally. Playable, if you can handle the constant attacks against your baby-snatching civ.

-The Crystal Empire
Playable, non-evil.
Has 6 gem pony castes with body size equal to dwarves (60K) and special material templates: bones and other hard tissues have crystal materials. They grasp with ‘magic’ body parts attached to both front hooves (instead of with mouth), every caste has a bonus for jewelcrafting.
    -Ruby & Sapphire: strong crystal materials, extra strength, +20% skill progression for close combat, smithing and masonry (ruby) or woodworking (sapphire)
    -Emerald: less strong crystal materials, extra agility, +20% skill progression for ranger skills and engineering
    -Amber: least strong crystal materials, +20% skill progression for farming
    -Topaz: same materials, +20% skill progresssion with crafts
    -Amethyst: same materials, +20% skill progresssion with medicine and organization
-The Demicorn caste: rare, 130K size, grasping with 2 magic parts attached to the horn and with mouth (like normal ponies), alicorn material set. All skills grow 30% faster.
The Rainbow Factory
A set of buildings with the purpose of processing creature materials. A lot of the reactions performed there use Pressing or Alchemy labors. Make sure that you’ve got some ponies with those enabled; Alchemy is also used in many other places in this mod.
The Spectral Extractor buildings process butchered body parts into Spectral Orbs of various colors. The “nervous tissue” parts are processed automatically without additional material input, but other parts (muscle tissue, fat, organs and horns/hooves) require catalyst caps. Most roaming creatures’ tissues are turned into Dim spectra; the regular spectra is refined from ponies (including the lesser ponies), Onyx spectra is from tyraponies and morphlings, and Brilliant spectra is from mythic beasts and alicorns. There’s an Extractor for each type; the regular and Onyx Extractors require dim rainbow blocks for construction, and Brilliant Extractors require regular rainbow blocks.
The Spectral Grinder automatically processes various miscellaneous body parts (such as cartilage and claws) into Rainbow Sludge. Alternatively, it can use chunks of bonemold [see Bonemold Production]. The sludge is stored in jugs.
The Spectral Mixer can liquefy the orbs, storing liquid spectra in jugs. The process can produce “dead” Rogue Stones, which can be “energized” using Rainbow Sludge. Liquid Rainbow can be is mixed from seven colors of Spectra.
The Spectral Kiln may be used to solidify liquid rainbow, using synneplast chunks (see Cloud Factory). Rainbow chunks can then be shaped into furniture items and blocks, which are vital for constructing advanced buildings.
The final building in the set is the Spectral Enchantment Altar. It can make regular Rainbow into Brilliant Rainbow using Ether Orbs [see High Tech]. It can also create Starmetal [Adamantine] wafers using brilliant rainbow and elemental diamonds [see High Tech].









Wood Processing
There aren’t any real trees in the world anymore. Instead, there are giant mushrooms and various weird things.
The Sawmill building processes these things. It is strongly advised to create multiple stockpiles for various kinds of wood: all the cut-down mushroom trees will be in the Fungal Tree category, which you will want to run through the sawmill to turn into Fungal Heartwood before constructing things out of it. The value of wood can be improved further through polishing into fairwood (needs oil or flutterwing jelly [see New bees])
The weird trees – flesh trees, bone spikes, glass spires, gem caps, radiant caps, living adra pillars – are not for use by carpenters and would, ideally, not drop logs at all; the game mechanics require that they drop logs, though, so it’s strongly advised to place them in a different stockpile that is linked only to the sawmill and the adra engineer [see Adra Engineer], but not the carpenter.
Flesh trees are processed into meat, hide and bonemold (bone spikes are just bonemold). Glass spires are broken into raw glass, gem caps may contain various gems, and radiant caps are broken into living crystal shards (see below).
Another function of the Sawmill is the processing of taintwood creature parts. These new creatures are dangerous, but hunting them is rewarding: when a taintwood creature is butchered, multiple “glob” items are produced; “taintwood” globs are turned into logs, and “taintwood bark” globs are turned into “bark” tools for the Organic Refinery. The taintwood logs can be polished into highest-value logs, or processed at the Ironbark Carpenter workshop.
At that workshop, a taintwood log can be processed into Ironbark Planks with obsidian and catalyst caps. These can be used to make furniture, weapons and armor in the same workshop; alternatively, they can be processed into Steelbark Planks if you have some onyx spectra. 

The Crystal Forge
Needs dim rainbow bricks for construction.

Processes crystal pony, crystal beast and golem crystal tissues into “living crystal shard” tools.
Crystal shards can be either shattered into gems, or combined into Force Crystal bars (those are used like metal bars).
The other set of functions uses milled Changeling Ferns and rough gems to grow replicated gem chunks: 1, 2 and 3 ferns for ornamental, semi-precious and precious gems; some special gems cannot be replicated. These chunks are then cut into furniture, leaving some gem fragments that can be reconstituted into special gems using crystal berries. Alternatively, they can be broken into 2-3 rough gems to multiply your gem count.
Gem furniture (as well as the ‘chunk’ and ‘fragment’ tools) is stockpiled as STONE.
Floral quarts gems that are created at the Organic Refinery can be enhanced into a semi-precious (needs a fungal garnet) and then a precious (needs a dim rainbow bar) gem. They can also be used to create an upgrade part for crossbows.
Also makes boulders from golem living stone, and rough gems from golem and star beast crystal parts.

The Magic Loom


Needs rainbow bricks for construction.
Creates high-value cloth from normal cloth and precious metals, using catalyst caps.
Silk and fiber cloth can be combined with silver or gold.
Yarn is also combined with silver or gold to create a stronger material as a “tanned hide”-type item to make leather armor.
Flutterwing silk is combine with silver, gold or platinum to make very high-value silks. Star silk (platinum) also needs crystal berries.
Cloth made from chromatic vines can be enhanced into Gemcloth (~adamantium weave).
Barkweave cloth is combined with silver caps or obsidian caps to make ‘soft’ or ‘hard’ high-value barkweave, producing “tanned hide” items. The ‘hard’ types (Ironbark or Steelbark) are good for leather armor.

Bonemold Production

The Organic Refinery can create chunks of Bonemold out of bones and chitin [the chitin is produced as “globs” at butchering to allow storing in a fat stockpile], which are then used at the Bonecrafter’s Shop to make furniture, weapons and armor. There’s four tiers of bonemold materials, analogous to the main creature material tiers in strength.
The Bonemold Enchantment Altar empowers chunks of bonemold using liquid rainbow of appropriate value. The enchanted bonemold types are equal to, in an ascending order: bronze, steel, fungal steel and saronite steel in terms of cutting and protective power; all of them remain light.
Bonemold items are stockpiled as BONE.

Bio-Metals
Another function of the Organic Refinery is the production of bio-metals. The special mushrooms (Green Cap, Silver Cap and Obsidian Cap), as well as bonemold and catalyst caps, are used there.
The bio-metals are: fungal brass, pewter, silver, gold, steel and Silksteel, with the last two being weapon/armor-grade. 
-Brass can be created from green caps and tainted bonemold, and then made into a fungal garnet gem with a (lower-value) fungal pewter as byproduct. 
-Silver is made from green and silver caps. Then an obsidian cap and tainted bonemold are added to make gold. Morphling and superior bonemold are needed to upgrade the bio-metal into steel and Silksteel, respectively.

Other Organic Refinery functions
-Creating floral quartz gems from crystal berries
-Creating barkweave thread from sawmill-produced bark
-Catalyzing Star Beast flesh into building materials

The Cloud Factory
This building houses a chain of four reactions that ends in the creation of Synneplast: a cheap and light material for furniture that is also essential for making high-tech weapons and solid rainbow. The first reaction doesn’t use any reagents, with probability-based products; the next two reactions each use the product of the previous one, and the last one uses catalyst caps. The resulting synneplast chunks can be shaped into items at the Kiln, using the pottery skill.

The Spirit Smelter
This smelter required dim rainbow blocks to construct. It processes Nether Ferrite ore (iron and taint crystals) and creates Felsteel / Green Pig Iron, Saronite Steel, Nether Electrum and Elder Electrum. There is a regular and magma-using version of this smelter.
High Tech
These functions are centered around the Adra gems.
These can be obtained from the Living Adra Pillar “trees”, found on the surface and underground, or from Befouled Adra veins. The pillar “wood” is processed at the Adra Engineer workshop, and the vein stone is processed at the Toxin Refinery workshop. 
The resulting Adra gems can be turned into Basic Adra Circuits at the Adra Engineer. Most high-tech reactions require a charged Adra Battery, which is made from an adra circuit and 2 taint crystals; a battery-using reaction produces a depleted battery, which can be recharged using biofuel (made at the Biofuel Refinery).
The Adra Engineer also manufactures advanced weapon components (Booster Circuit: 2 circuits + silver, Power Core: 2 circuits + gold, Emitter Assembly: 2 circuits + 2 platinum), ammo for the Star Blaster (1 circuit = 20 shots), alchemy transmutation cores (5 circuits = 5 – 10 cores).
The Toxin Refinery also processes the Blood of Yogg-Saron and Nether Sparkstone minerals, makes Ether Orbs out of taint crystals and mana crystals, deals with saronite thread/wafer energizing reactions (for Saronite Steel), grinds taint crystals for the Felglass Furnace and mixes them with clay for use at the Kiln. [Felglass and fel stoneware items are stockpiled as STONE.]
The Starmetal wafer making is also done there.
The Adra Reaction Chamber is used to compress adra gems into a harmonic diamond, and then to create elemental versions of the diamond using Ether Orbs. A set of those can be combined with rainbow sludge and rainbow liquid in the IMP Vats [enabled for ponies, rainbow ponies, cave ponies and diomedian ponies] to release a cloud of vapor that turns ponies into alicorns (or into berserking tentacled monsters… oops).
The Toxin Refinery, Adra Engineer and IMP Vats all require rainbow bricks for construction.

The Alchemy Lab
Needs a rainbow bar for construction; reactions require transmutation cores and batteries.
Worthless boulders can be made into low-value non-magical metals (V<10). The next tier makes higher-value metals from copper (V<30), the next is silver. Platinum is made from gold.
The reactions that create non-magical metals also have a chance of producing Ether Orbs. Those are required to create magical metals: volcanic aurum (from gold), saronite (from aluminium) and Elementium (from energized saronite, brilliant rainbow, and aurum). Rainbow sludge is also needed for those reactions.
Guns
The Gunsmith’s Forge can assemble weapons from the components created in the Adra Engineer, also using the weapon casings and capacitors made in the Forge itself.
Casings require metal bars, and they determine the material the weapon will be made of.
Capacitors require an adra circuit and either a metal bar or a jug of liquid rainbow.
Flechette launchers (1 power core, 2 booster circuits, 2 silver capacitors), slugthrowers (2 power cores, 2 booster circuits, 2 gold capacitors), heavy slugthrowers (2 power cores, 4 booster circuits, 2 platinum capacitors) and star blasters (1 emitter assembly, 4 power cores, 6 booster circuits, 3 brilliant rainbow capacitors) are assembled.
Pistols and airguns are simply made there from metal bars (and synneplast chunks, in the case of airguns). 
The Gunsmith’s Forge can also upgrade ranged weapons (excluding bows, blowguns and various exotics). The first step (for pistols, crossbows, slugthrowers and flechette launchers) involves putting a bayonet on the weapon, which is also forged there (and has to be made from the same metal as the weapon, excluding crossbows); pistols use the ‘small’ bayonet (this is also the only step for pistols). For the later steps, the weapon is enhanced with various components, such as additional booster circuits and capacitors. Heavy slugthrowers and star blasters cannot be upgraded.
Gunsmith reactions may be using either fuel (if something is forged) or a battery (if something delicate is made); sometimes, none is needed.
The Ammunition Mint is required for making advanced weapon ammo. For pistols, slugthrowers and heavy slugthrowers, the ammo is made with synneplast casings (Kiln), metal bars and gunpowder (made from coal and saltpeter, or coal and potash, at the Rock Grinder). For flechette launchers, the ammo is made from a metal bar and synneplast sabots (Kiln).

Miscellaneous New Buildings
-Food Purifier
Purifies the meat and organs of creatures affected by Taint (which is almost every creature except for the normal ponies and domestics), as they can't be eaten otherwise. Needs a dim rainbow bucket for construction.
-Rock Grinder
Can grind worthless rocks into sand; can make gunpowder [see High Tech and Guns]; can prospect 10 worthless rocks for gems.
-Biofuel Refinery
Can produce biofuel using tallow, low-value alcohol or solar cap oil. The biofuel can be used to recharge batteries at the Toxin Refinery, or be converted into normal fuel. The canisters are stockpiled as a STONE tool.
-Advanced Distillery
Can brew drinks using several types of plant at once (for example, wormbrew uses both kinds of worms), into pots or barrels.
-Decorator's workshop
Has several reactions to decorate weapons and armor with stuff. 
-Crematorium
Burns vermin remains into ash and makes flux from bonemold or tainted bonemold. 
-Fel Glass Furnace
Can make Fel Glass items, using ground taint crystals and pearlash. Needs magma.
-Display Cases
Built from a glass box (clear, crystal or fel) and an item of your choosing. Put your awesome artifact tin low boot in there, so that the ponies can admire the completely sublime display case.
-New functions for the old buildings
The Still can now make alcohol from milk. Requires the milk to be stored in buckets (a stockpile for milk, with 0 barrels allowed).
The Still can also make alcohol from specific plants, instead of generally “brewing drinks”.
Tanning reactions now produce Leather Scraps. These are used in the tanner’s shop to make recycled leather. Another new option for that shop is to boil plant cloth in solar cap oil and get “artificial leather”.
The Quern has a special reaction for milling solar caps into paste. To press oil from it, the default “press oil” reaction is used at the screw press.
The Quern can also mill specific plants into powder instead of generally “milling plants”.



Creature Redesign 
-Most Large Roaming crearues have new castes. Variants include full or partial chitin plating, stabbing tusks for predators, dual-stabbing horns for reptiles, skinless pain-insensitive abominations.
-All animal men have been ponified using a new CVar. Insect ponies are grouped into two creatures (hostile and non-hostile). Apex predator ponies are alicorns, grouped into one creature.
-Humanoid monsters are ponified, as well (for ex. Yeti -> Wendigo, Ogre -> Trampler etc.)

 					New Large Creatures 
-Star Beasts (a creature with a special template group; less sensitive to pain although not insensitive; organs are replaced with crystals; multiple castes; flesh and crystals are processed at the organic refinery)
-Taintwood beasts: Dryads, Timber Wolves, Tainted Crushers, Crawlers, Skitterers and Bogstalkers. Taintwood is as strong as bronze.
-Tyranid ponies (several solitary and pack creatures with the Tyrapony material set)
-Spawns of Hakkar (dangerous underground creatures; several castes with different material sets; their hearts are processed into gems at the organic refinery)
-Silithids (also dangerous underground creatures; contain pearls)
-Metal and gem ponies (former iron and amethyst men)
-Felstalkers (surface pack predator with the Wild Morphling set)
-Giant Centipedes, Drawde’s Drakes
-Domestic Lesser Ponies, also encountered in the wild. The Lesser caste is a wagon-puller, the Breakers are carnivorous war beasts. All lesser pony spinal tissue is Spectral.
-Rainbow Worgs for war and Rainbow Ocelots for pest control. Also, the ocelots don’t choose their owner.
-Feral crystal ponies, including feral ruby, sapphire, emerald and "everyone else" creatures. There's normal, mutated (crystal carapace) and abomination (skinless) castes.
-Crystal golems, with bodies consisting of living rock with a crystal substructure (similiar material as the crystal pony bones). Damaging the crystal can shock the golem (pain), also they have Heart Crystals in the head and chest.
-Crystal Beasts
   	 -"Crystal wolf" creature with several castes of varying size and threat, from a pony-sized stalker to an alicorn-sized Slayer with a crystal shell and scythe-limbs
   	 -"Crystal bug" (also several varying castes)
   	 -Crystal Snatcher - a night creature

New Megabeasts 
-Nido’s Trojan Ponies are now made from tainted wood
-Brutalisks (large chitinous monster with scything claws and tentacles)
-Reaper Lizards are huge scaled beasts with enormous claws (based on the real-life Therizinosaur)
-Night Haunters (aka Smarty Pants) are huge, with three heads, stance-tentacles and giant claws. They have spouse conversion (into “undead” Haunted Abominations).
-Jade Destroyers (an alicorn statue made of felsteel)
 
There's an entire group of alicorn megabeasts. Size isn't very large, but they have high natural combat skills.
-Queen of Blades (has dark iron plating, teeth and scyhing claws, and a material breath attack)
-Mares of Pain (aka Lil' Miss Rarity) and Blood (aka Pinkamena). They, and the next three beasts, have spouse conversion. This one converts into “undead” Tormented Abominations.
-Changeling Queens (conversion into Changeling Behemoths with morphling tissues)
-Crystal Mistresses (conversion into Crystal Slaves with crystal tissues)
-Star Nightmares  have both star beast and alicorn tissues, as well as starmetal horns and teeth. Conversion into Star-Touched Abominations with star beast tissues.

New Bees 
Bees are replaced by Flutterwings. There’s several variations of them that are encountered in different biomes, which produce different items. 
-Forest flutterwings’ jelly is used to polish wood instead of oil for greater value
-Ice combs are powdered into dye
-Swamp and Desert flutterwings produce gems
-Plains have no special abilitity

All Flutterwings also produce silk.