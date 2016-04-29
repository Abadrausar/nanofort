[Addon Change log](/addonfile?addon=Milo___First_Landing_Final&file=addon_changes.md)
[Readme](/addonfile?addon=Milo___First_Landing&file=Readme.md)
![first landing Haiku](http://orig12.deviantart.net/41a4/f/2011/106/8/d/haiku_nature_by_bohleng-d3e3yb8.jpg)
## So you think that you need some help, want to do a sugestion, inform about a bug or give thanks to modders or developpers, but not sure where to go.
1. If you want to patronize Toady One and the developpment of Dwarf Fortress go to: [Support Dwarf Fortress Vanilla](http://www.bay12games.com/support.html)
2. For questions about the rubble framework go to: [Rubble framework](http://www.bay12forums.com/smf/index.php?topic=154304.0)
3. Anything about HDFPS: Humble Dwarf Fortress Publishing System is there [HDFPS](http://www.bay12forums.com/smf/index.php?topic=157300.0)
4. If your question is addressed to the actual original maintainer of this mod then you have [First Landing](http://www.bay12forums.com/smf/index.php?topic=151180.0)
# Welcome to first landing! 

[Changelog and install procedure](addons:First Landing/Readme.md)

In year one the first settlers landed on this planet and have been sending
out small group to colonize the surface ever since. You will guide one of these groups, fending off
rivals, raiders, natives, and hostile wildlife to found a city that will stand the test of time and
the elements.

Will you found a bustling surface metropolis, or will you have a underground city dug into a mountain?
Will your settlement be a military powerhouse or a center of industry? Will you choose to land on a
planet full of unfriendly natives, or will it have legions of hostile wildlife and suffer constant
pirate raids? Lead your settlement to a future as a glorious gem in the colonial administration crown
or to an ignoble end!


Notes:
----------------------------------------------------------------------------------------------------

Artifacts are automatically disabled when loading a First Landing save. Your settings will be restored
when you quit (so non-First Landing saves will not be effected).

You **must** have DFHack installed to use this addon! If you do not have DFHack installed this addon
will not work correctly, leading to all sorts of weirdness!

**Do not use "Play Now!" when embarking!** If you use "Play Now!" you will be missing certain critical
items and will have a much harder time getting established! See the section on workshops farther down
for more details. In the same vein it is a bad idea to to use the "Create New Wold!" option, use
"Design New World with Advanced Parameters" option with one of the First Landing parameter sets. These
parameter sets are included in the file "world_gen.parameters" in the First Landing archive (installation
instructions are in that file).

I suggest you disable all vanilla random creatures/syndromes/werebeasts/etc, as they are all geared
for a fantasy game, and this is scifi. I also suggest you do not generate worlds with a "bottom layer",
as the adamantine replacement cannot be mined, and if you somehow manage anyway, it cannot be used for
anything (plus clowns do not belong in a scifi game).

For the first few fortresses it is a very good idea to set the pop cap very low, as you will have to
pay a great deal of attention to the basic industries, and the less "distractions" you have eating
your food and booze, the more time you can devote to learning the new industries. `addon:Util/Pop Cap`
is very good for this.

It is a good idea to turn mineral scarcity way down, as you will want a mix of many kinds of minerals,
and in my experience worlds with the default settings (scarcity of 2500) end up with large amounts of
only a few mineral types. Some minerals and metals are critical to certain industries, so ensuring
your supply is critical. While it is possible to trade for crates of resources, this can be slow.

Trade is critical to your success, as you will need a broad range of raw materials, not all of which
will be available everywhere. Trading for resource crates is critical to military production and high
end buildings.

It is a very good idea to bind DFHack's "gui/workshop-job" script to a hotkey in your dfhack.init file!
This script will allow you to set the exact material to use with a specific job, which can save a
great deal of time when you want to make one or two items from a specific material, and don't want
to mess with stockpile settings. The default dfhack.init file binds this command to "Alt+A"

This base is not compatible with many of the Rubble default addons, either because it incorporates a
modified version of the addon directly or because the addon depends on objects which are not available.
Not all incompatible addon are marked as such, so use some common sense and don't use anything that
you are not 100% sure will work! Obviously this warning does not apply to preinstalled versions, as
they only include addons that are confirmed compatible.

This addon description is best read in the Rubble web UI, that way the workshop previews, documentation
links, etc all work, and things look much nicer. To do this start the web UI (after installing this
pack) and go to the addon list page, then select the `addon:First Landing/Base` entry.

Currently all plants and (non-domestic) animals are loaded from external addons. Most of these plants
and animals are randomly generated, making each world unique. This leads to a more interesting and
challenging game (as you never know exactly what you will find until you embark). This addon pack is
**not** for beginners!

Speaking of "not for beginners", this addon's version of goblins are much tougher, with mostly steel
equipment and only the best armor and weapons. Unfortunately DF is kinda dumb about equipping invaders
with armor, so they will likely not have full sets. The kobold replacement is numerous and will happily
send early ambushes, but they have crappy weapons and armor. An early militia can be invaluable!
Keep in mind that all your enemies are (supposed to be) allied (sometimes they fight each other), so
you will likely see more invaders than you normally would (as they are not wasting their efforts on
each other). In particular if you settle a long way from your civilization's center of power you may
be the only settlement in the area that is within reach to attack (making invasions *much* more
common than normal).

DF will generate lots of error log spam about missing material templates preventing random creature
generation during worldgen. This is a good thing, as it stops forgotten beasts in their tracks :P If
you find any other errors I want to know about them, but this particular error category should be
ignored. These errors should only happen during world generation, so it's not like your error log
will fill every time you load an FL world.


Embarking:
----------------------------------------------------------------------------------------------------

Proper embark preparation in First Landing is critical. There are a few items you *must* bring if you
want your settlement to get off to a good start, and site selection is just as important.

When selecting a site you will generally look for all the things that you would when embarking normally.
A river is particularly useful, as you will need mechanical power for several of the critical basic
industries. Clay is spectacularly useful, as there are many more things that can be made, and it
is possible to use some of the items (like bricks) without firing them, making pottery the cheapest
and easiest source of large amounts of early game building material. Sand is also nice to have, as
it makes another fine material for furniture and the like (although it does cost fuel).

As far as items go, mostly you will want the same things as you normally would, a pickaxe, hand axe,
alcohol, seeds, food, etc. You will also want to take some crates of hand tools with you, if you do
not you will have a hard time building basic workshops (you will need to use primitive tools, which
degrade over time, until you can get some metal and make proper metal tools). A crate (or two) of
machine tools will make it much easier to establish a metal or plastics industry. A saw blade will
allow you to build a sawmill early on, which is critical if you want to do much with wood. If you
want to build an early plastics factory you can get crates of plastic very cheaply, I always take a
few crates simply for early building materials and for making mechanisms. If you sacrifice most of
your non-essentials it is possible to take a crated electronics assembler, but it is usually better
to buy these from caravans later (as they are very expensive and of limited utility early on).

For skills you will probably want a mechanic as well as several settlers skilled in some way to make
food (fisherman and fish cleaner or grower are probably your best bets). If you want a plastics industry
you will want a manufacturer (wax worker), and a blacksmith is good if you want to work metal. A potter
or glass maker may also be useful. Weapon and armor smiths are probably best not taken, as you will be
hard pressed to get the materials together for good quality arms and armor early on (good weapons
require chemicals and lots of finicky machining, good armor requires either chemicals or electronics).
If you want an early militia you will probably have to equip them poorly anyway, so skilled crafting
is not terribly useful. If you brought crates of expensive finished goods it may be worthwhile to
bring a skilled pump operator to open them (particularly if they contain weapons or armor).

I strongly recommend reading the workshops section (farther down) before you embark, as it has more
information about what to bring and what they are good for. In particular pay attention to the early
industry guide at the top of the workshops section, as it details what to do after arrival.


Nobles:
----------------------------------------------------------------------------------------------------

Most of the nobles should be easy to figure out, as they are not much different from what you are
used to with dwarves. The primary changes lay with the military.

Early on you can only appoint a security chief and security team leaders. Both of these positions
lead teams of 5 security guards, and are your only military in the early game. The catch? They are
also law enforcement personnel, so no avoiding the justice system unless you can do without a military.

Once the town administrator shows up you can appoint a militia captain and militia lieutenants, both
of who lead squads of 10.

The following is a list of all the (non-military) nobles that have different names from their dwarf counterparts.

	Dwarf Name:          Settler Name:
	chief medical dwarf  medic
	bookkeeper           accountant
	manager              foreman
	broker               trader
	hammerer             executioner
	baron (a barony)     town administrator (a town)
	count (a county)     city administrator (a city)
	duke (a duchy)       <none>
	king                 colony administrator
	outpost liaison      merchant


Weapons and Armor:
----------------------------------------------------------------------------------------------------

For weapons there are a variety of melee weapons available, ranging from simple hunks of metal (crowbar)
to powerful bits of technology (chainsword).

* Hand axe: Chop down trees with it, not good for much else.
* Chainsaw: A better way to cut down trees (or foes).
* Pickaxe: Delve into the earth with this simple tool, also delve into the bodies of your foes.
* Mining drill: The bigger, better, cousin of the pickaxe.
* Bayonet: A large dagger fit for perforating enemies or mounting on your assault rifle.
* Crowbar: A curved bar of metal, perfect for opening crates or craniums.
* Sledgehammer: A big hammer. Simple, but potentially deadly.
* Chainsword: Deadly. Also messy.
* Power gauntlet: For those who just can't resist punching things.
* Chaindagger: The chainsword's little brother.
* Power sledge: This is what happens when you leave a bored weapons engineer alone with a sledgehammer.

Melee weapons are useful, but more useful are the ranged weapons, which consist of an assortment of
guns. There are guns for every occasion, from the mighty shotgun (mighty as long as the target has
no armor anyway), the pathetic but cheap air rifle, and the godly anti-material rifle.

* Air rifle: Mostly for training and early hunting, weak and cheap.
* Pistol: Cheap, but only marginally more powerful than the air rifle. This is the only gun that
  can be used one handed.
* Shotgun: A great hunting weapon, but not so good against armor.
* Hunting Rifle: Good basic militia weapon, some anti-armor ability. No real weaknesses, but no
  strengths either.
* Assault Rifle: Exactly the same as the hunting rifle, but with a better melee attack.
* Anti-material rifle: What you need for taking out armor, should stop all but the toughest opponents.

Ammo for the guns comes in lead and silver versions.

* * *

Armor made from leather is by far the worst armor you can get, but it is better than nothing.

Metal armor provides good protection, but you can only make a few pieces, namely a "trauma plate"
(that can be worn over flexceram or leather armor) and a helmet. You can also make metal shields.

Flexceram armor is roughly equivalent to copper, and you can make full suits. Flexceram armor is
enough to make you more-or-less immune to the weaker guns such as the shotgun and pistol. Flexceram's
primary drawback is that you need a chemistry industry to manufacture it, so it is beyond the means
of most early settlements.

Power armor is the best protection you can get, plus it makes the wearer faster, stronger, and tougher.
Of course power armor is expensive and hard to craft, requiring several electronic components (which
tend to be expensive) as well as some bars of metal. As much as it costs, making power armor with
anything but steel and a skilled armor smith is pretty stupid, but if you want to be stupid the
option exists :)


Workshops:
----------------------------------------------------------------------------------------------------

The following vanilla workshops have been removed (they won't show in the build list):

* Metalsmiths forge (use machine shop, armorer, or gunsmith)
* Smelter (use fuel furnace or arc furnace)
* Glass furnace (use fuel furnace or arc furnace and glassblower)
* Bowyer (use gunsmith)
* Mason (use furniture assembler and sawmill)
* Carpenter (use furniture assembler and sawmill)
* Craftsman (use bonecarver, furniture assembler, and bookbinder)
* Mechanic (replaced by custom workshop with the same name)
* Clothier (use tailor)
* Leatherworks (use tailor and armorer)
* Soap maker (use chemist)
* Wood furnace (use fuel furnace)
* Kiln (use fuel furnace or arc furnace and potters wheel)

In general most production chains have at least one extra step or complication, for example most
furniture is made from blocks or planks instead of raw boulders or logs.

Many of the workshops need special items, such as hand or machine tools, to construct. The days of
being able to fashion complicated tools from a random hunk of rock are past!

Some of the workshops require mechanical power to operate. Generally there will be tiles on the workshop
that look like a gear assembly, connect your axle or whatever here. The power supply must be connected
*after* the workshop is constructed, and it must be connected in the same game session. If you save
and quit while the workshop is being built the powered workshop plugin seems to be unable to connect
the workshop to the power grid. If you connect a workshop and it does not start animating (assuming
you have enough power) deconstruct the connecting axle or gear assembly and try again, if it still
won't work deconstruct both the connector and the workshop and build both again (workshop first).
Workshops can be powered from above (through a hole in the ceiling), but not from below (as you cannot
build a workshop over a hole in the floor).

If you want to embark on a location where it will not be possible to have wind or water power you will
want to take a crate of "power plant parts". This will allow you to build a power plant early on so
you can do metal working or plastic manufacturing. If you do not have power you can't work metal, and
if you can't work metal you can't make a power plant. It should always be possible to build a water
reactor, but on maps with little water this can be very difficult.

For more-or-less complete early industry you will need some hand tools (which can be brought on embark
in crates) and a way to get raw materials such as a pick or hand axe. If you will not be able to get
weapons grade metal soon after you embark (due to an aquifer or lack of minerals) you will want to
bring some bars or ore so that you can make the hand and machine tools your settlement will require.
Almost all the workshops require hand tools, so ensuring your supply is critical!

If you forgot to bring hand tools (it happens to us all, we get to the embark site and start unloading
the wagon, and someone goes "hey, did you bring the tools?") it is possible to make temporary
ones at the primitive workbench. These "primitive" tools wear out over time, so they are only good for
emergencies. It is also possible to make primitive hand axes and pickaxes, but once again they wear
out, making them of limited use.

The following is a plan of action for getting basic industry up and running. This plan focuses on
getting a sawmill and hand tool production on-line as soon as possible. If you bring machine tools
and a saw blade on embark much of this can be skipped, but for completeness I assume you embark with
a single hand tool crate, an axe, a pickaxe, and some weapons grade metal or metal ore (copper works).

Immediately after embarking you will want to build an unloading bay and unpack your hand tools, as
they will be needed to build almost all other workshops. Ideally you have more than one crate of hand
tools, in which case you can build several workshops right away, but here I will assume that you only
brought one.

The first thing to do is get some wood or stone, then build a splitting block and a mechanic. Make
some blocks or planks at the splitting block, then make some mechanisms at the mechanic. As soon as
you have ~6 mechanisms use some metal and mechanisms to build a set of machine tools.

Next build a machine shop with your machine tools, also build a windmill or water wheel near the
machine shop, as it needs power. Once the machine shop is finished connect it to your power source.
Now use the machine shop to make a saw blade and any more hand tools you need. If you do not have
much metal it is also possible to build a fuel furnace and a glass blower to make a glass saw blade
instead. Charcoal can be made at the fuel furnace, but without a sawmill this is fairly inefficient.

Now use your glass or metal saw blade, some mechanisms, and a few blocks or planks to build a sawmill.
The sawmill is *very* important for your early industry, as it is twice as efficient as the splitting
block (on that note, feel free to deconstruct your splitting block, everything it can do the sawmill
does better). Keep in mind that like the machine shop the sawmill requires power, so it is probably
a good idea to keep those two close together, at least early on.

Now you have all the basics you need to expand your industry in any way you want, a good next step
would be to build a furniture assembler.

High tech items (such as power armor) require electronic components of various types, these need an
electronics assembler to be constructed. Electronics assemblers are too complicated to produce on-planet
so they will need to be imported (at great expense). It is possible to bring one on embark, but it is
unlikely you will want to do so in most cases (due to their cost).

Many high tech items also require an RTG for power, these need to be made at an electronics workshop,
and among other things they need uranium fuel rods. Uranium requires special methods to refine, and
as such you will need a special workshop. This workshop not only requires lots of mechanical power,
it also needs a centrifuge, which (like the electronics assembler) is an expensive import only item.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_furniture.txt&id=FURNITURE_ASSEMBLER)

The furniture assembler creates basic furniture (and some tools) from wood and stone blocks (made at
the sawmill).

This is the replacement for the carpenter, mason, *and* craftsman.

Items are made by your carpenter or mason (depending on material).

Like many of the basic buildings this requires hand tools, so remember to bring some on embark.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_mechanic.txt&id=CUSTOM_MECHANIC)

The mechanic's workshop is where your mechanic can fashion mechanisms from wood, stone, metal or
plastic. Wood and stone mechanisms are made from planks or blocks, so a sawmill is advised.

This is also where traction benches are assembled.

To help jump-start early industry the mechanic can also make machine tools and hand tools, but at a
higher cost than if you made them at the machine shop. In the same vein you can make primitive radio
parts here from wood planks, carbon (coke or charcoal), and metal.

Like many of the basic buildings this requires hand tools, so remember to bring some on embark.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_tailor.txt&id=TAILOR)

The tailor's workshop is where clothing and other cloth and leather items (such as backpacks and
quivers) are made.

Like many of the basic buildings this requires hand tools, so remember to bring some on embark.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_crates.txt&id=UNLOADING_BAY)

The unloading bay allows you to unpack any crates you buy from caravans. Crates are items that can be
"unpacked" into other items. Crates cannot be stockpiled, but they will be automatically unpacked by
your pump operators if you have an unloading bay. It is generally best to build the unloading bay
near your depot.

Keep in mind that caravans will not bring crates unless you order them!

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_power.txt&id=DFHACK_POWERED_POWER_PLANT)

The power plant (as it's name suggests) produces mechanical power.

To build a power plant you will need to produce two "piston and cylinder" parts and one "steam boiler"
at the machine shop (from "hard" metal).

Once the power plant is built place a machine input (found in the machines category of the build menu)
next to it and dump coke, charcoal, plant oil (in jugs or barrels), wooden logs, or wooden planks onto
it. Every now and then the power plant will look for fuel and consume one fuel item (whichever one it
finds first), if it finds fuel it will produce mechanical power until the next scheduled refueling.
If it should ever run out of fuel it will simply stop producing power and wait until fuel is added.

When burning plant oil place a machine output next to the power plant to mark where you want empty
containers to be placed.

Here is how long each type of fuel lasts (expressed in unit-less measures good only for comparison):

* 1 unit of plant oil: 2 pu
* 1 bar of coke/charcoal: 2 pu
* Wood block: 1 pu
* Wood log: 4 pu

As you can see with a sawmill you can double your power output from wood by making charcoal, and plant
oil is a great source of power if you have skilled farmers (or not many trees).

Keeping a power plant in continuous operation requires a **lot** of fuel! I suggest you only fuel
the plant when you actually want to use the connected machines. Later on when your settlement is
larger and continuous mechanical power is required you will either want a better source of power or
a large oil plant farm!

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_common.txt&id=PRIMITIVE_ASSEMBLY)

The primitive workbench fills a critical role, if you ever find yourself without hand tools or in
need of an axe or pick in an emergency, you can make "primitive" versions of these items from wood
(I assume your settlers pick up some stones to use for tips and such). These items wear out over time,
so they are only good for temporary use, but they are perfect for times where you forgot to bring a
tool crate...

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_sawmill.txt&id=SAWMILL)
![](/wshop?addon=First%20Landing/Base&file=building_flbase_sawmill.txt&id=SPLITTER)

The sawmill allows you to cut logs and boulders into planks or blocks.

All reactions require the carpentry or masonry skill.

Running a sawmill requires mechanical power, connect the axle or gear to the saw blade in the middle of the
left side of the workshop. Power inputs must be constructed after the sawmill has been completed.

You will need a saw blade to build a sawmill, so either bring one on embark or bring the supplies to
make one (needs to be made from a weapons grade metal or glass).

You can also cut certain minerals (coal, sulfur, saltpeter, and any flux) into blocks here with a
special reaction. These minerals are used in block form at the chemist or furnace.

Since the sawmill is where you make blocks and planks, but you need blocks and planks to make
furniture (and the sawmill itself) there is a single tile "splitting block" workshop that you can
make that also allows you to create blocks and/or planks, but only half as many for each input.
In general you will want to replace the splitting block with a sawmill as soon as possible.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_smelter.txt&id=FUEL_SMELTER)
![](/wshop?addon=First%20Landing/Base&file=building_flbase_smelter.txt&id=ARC_SMELTER)

The fuel furnace allows you to process some basic ores and melt metal items into bars using fuel.

The fuel furnace is also where you make charcoal and ash from wood planks. In emergencies it is also
possible to turn coal directly into coke without proper fuel, but this is quite inefficient

Advanced ores (like bauxite) will require the arc furnace. The arc furnace requires some special
parts and a bunch of mechanical power, but no fuel.

Both furnaces can also be used to make raw glass and for firing pottery.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_smelter.txt&id=ROCK_BREAKER_MANUAL&stage=1)

The rock breaker allows your laborers (pump operators) to use a sledge hammer (required
to build the workshop) to break up metal ores, improving smelting efficiency.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_forge.txt&id=MACHINE_SHOP_BASIC)

The basic machine shop allows you to make many items from metal. Unlike the vanilla forge the machine
shop does not require fuel, but uses mechanical power instead. Power use is quite light, so a single
windmill should be enough in most cases.

All reactions at the machine shop are done by your machinists (blacksmiths).

Some items need to be made from certain classes of metals, here are the metal classes:

* Weapon metals (used for melee weapons and such like)
	* iron
	* copper
	* cobalt
	* titanium
	* steel
	* bronze

* Hard metals (used for things like armor plates and gun barrels)
	* iron
	* titanium
	* steel

Most things can be made from any metal, including a bunch of metals that are not listed above.

Some of the fancier melee weapons also require circuits and batteries made at the electronics shop.

Flexceram, although a "metal" as far as the game is concerned, requires special techniques to shape.
This material is limited to making armor and is not suitable for making items at the machine shop.

Power inputs need to be connected to the top corners of the workshop, and inputs must be built after
the workshop has been completed.

The machine shop requires machine tools to build, so make sure you bring some on embark or have
materials to make some.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_armorer.txt&id=ARMORER)

The armorer is where all the armor needed for your military is assembled from metal, leather, or
flexceram.

Flexceram is a ceramic impregnated, ultra-strong, anti-ballistic cloth produced at the chemist.
Flexceram armor is not as good as metal power armor, but it is a lot cheaper.

Metal can only be used to make a few specific armor pieces, most pieces can only be made from leather
or flexceram.

Power armor is also assembled here from armor grade metals, RTG modules, processing units, and
electric motors.

All armor assembly uses the armor smith skill.

Like many of the basic buildings this requires hand tools, so remember to bring some on embark.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_gunsmith.txt&id=GUNSMITH)

The gunsmith's workshop is where your dedicated gunsmiths (also known as weapon smiths) create much
needed firearms and ammunition.

Most of the metal parts for making a gun must be made at the machine shop, the gunsmith only does
final assembly and creation of certain special parts (for example rifling smooth bore barrels).

When assembling a gun the material used for the final weapon is the material of the receiver or frame
except in the case of the assault rifle, where it is the material of the bayonet. Material does not
effect performance as a ranged weapon, but it does effect value and melee performance.

Most guns require proper ammunition made from metal, gunpowder, and primers, but weaker air rifles
can be made for training or settlements that have a gunpowder/primer shortage.

Generally the more powerful a weapon is the more it's ammunition costs. Ammo can be made from lead
or silver and (for all but the air rifle) gunpowder and primers. Most gunpowder and primers are
purchased by the crate from caravans, but established settlements may create their own at the chemist.

Like many of the basic buildings this requires hand tools, so remember to bring some on embark.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_common.txt&id=RESOURCE_COLLECTOR&stage=1)

The resource collection area is a small (free) workshop that allows you to issue "collect sand" and
"collect clay" orders. Build near your clay and sand zones.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_glassblower.txt&id=GLASS_BLOWER)

Glass making has been completely redone. Raw glass is now produced at the fuel furnace or arc furnace,
then formed into useful items at the glass blower.

This takes twice as much fuel per item (unless you are using the arc furnace).

Although glass items cost more to make, they are still very useful and quite cheap for the result.
Sand is almost everywhere, so a glass blower is a fine investment for most settlements.

Crystal glass may be made from rock crystal or sand and lead (plus pearlash of course).

Fill sand bag orders may be issued at the resource collection area workshop.

Like many of the basic buildings this requires hand tools, so remember to bring some on embark.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_pottery.txt&id=POTTERS_WHEEL)

The potter's wheel is the center of the much expanded ceramics industry. Clay is formed into various
"wet" items here. These items then air dry over time until they become "dry". Once the items are dry
they may be glazed and then fired at the kiln.

There are three grades of ceramic, "poor ceramic" which is made from clay, "ceramic" which is made
from fine clay, and "fine ceramic" which is made from orthoclase, microcline, or kaolinite. The only
differences between the different grades are value and the fact that poor ceramic containers must be
glazed before they can hold liquids.

Note that glazing should be done before firing, you can glaze an item after it has been fired, but
this is relatively inefficient (you need to fire the item again to set the glaze). Stone items may
be glazed this way for decoration, but aside from boosting item value this has no real benefit.

To glaze an item you must grind certain boulders (mostly metal ores) or ash bars into glaze powder.
This powder is then used to glaze items with several different patterns.

Glaze powders are broadly classified as red, yellow, blue, and white. Glazing patterns require each
powder to be a different color, and when glazing with more than one color one of the colors must be
white. Unfortunately First Landing has less "useless" minerals so certain colors are hard to find,
in particular the only sources of "yellow" colors are sulfur, uranium and gold, all valuable resources.

Of course colored glazing is just for looks and value, if you need cheap glaze just use ash.

In some cases it may make sense to use the dry (but unfired) items. They are worth a lot less than
the fired items, but they also don't need any fuel and take less labor. Dry bricks make a good early
building material for a settlement that has clay!

Gather clay orders may be issued at a resource collection area workshop.

The potters wheel is one of the (very) few workshops that does not require hand tools.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_chemist.txt&id=CHEMIST)

The chemist is where soap, plastic, gunpowder, etc are made.

Making raw chemicals (for example sulfuric acid or mercury) requires a pack of glass vials (made at
the glass blower) as well as some raw materials (and possibly other chemicals).

Some of the reactions require water. To run these reactions you need some water below the workshop
accessible through an adjacent opening in the floor. 1/7 water is used up for each reaction that
needs it.

Sulfur and saltpeter are *very* important for production of sulfuric acid and nitric acid, both of
which are critical to several industrial processes. If you plan to do a lot of chemical production
it may be a good idea to order boulders of both minerals from caravans. Sulfur in particular is
critical. In the absence of saltpeter you can use potash, but there is no substitute for sulfur.

Gunpowder is made from 1 nitric and 1 sulfuric acid, 2 plant thread, and some water to wash the
resulting nitrocellulose.

Primers for ammunition manufacture are made from 1 nitric acid, 2 mercury, and 4 ethyl alcohol.

Flexceram requires 2 plant cloth, 2 ceramic blocks (either ceramic or fine ceramic), 2 ethyl alcohol,
and 1 sulfuric acid. Flexceram is the perfect material for clothing and under-armor, an ultra-strong
ceramic impregnated cloth suitable for making flexible anti-ballistic armor.

Plastic requires 1 block of coal, 2 ethyl alcohol, 1 nitric acid, and water. After you have formed
some plastic bars you can use them to make synthetic cloth, perfect for cheap clothing when you don't
want to farm.

It is possible to take some plant oil and a little water and process them into solid fuel (coke) that
can be used to fuel furnaces. This is a great way to save your coal, as coal is critical for plastic
production.

Soap is made exactly like it is normally, with the additional option to make alcohol based soaps
instead of lye based.

Building the chemist requires four packs of glass vials to use as laboratory equipment.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_electronics.txt&id=ELECTRONICS)

The electronics workshop is where your hard working technicians (strand extractors) manufacture
complicated electronic components used in the creation of advanced items.

The following items may be created:

* Solder
* Wire
* Circuits
* Processing units
* RTG modules (RTG stands for Radioisotope Thermoelectric Generator)
* Mechanical arms
* Electric motors and generators
* Batteries

This workshop requires an electronics assembler and hand tools to build. Electronics assemblers cannot
be built by your settlement, you will need to import them.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_uranium.txt&id=URANIUM_REFINERY)

Here is where you turn pitchblende into uranium fuel rods for RTGs.

This workshop requires a centrifuge to build, unfortunately centrifuges are import only, and they cost
even more than electronics assemblers, so do not expect to use this workshop until your settlement is
well established.

Refining uranium requires sulfuric acid and ready access to water (which needs to be accessible
through an adjacent opening in the floor).

The centrifuge requires quite a bit of power to run, connect your power supply to any of the workshop's
corners after it has been built.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_plastics.txt&id=PLASTICS_FACTORY)

The plastics factory is the source of most of your furniture, tools, and other common items. Here you
will produce large amounts of cheap items to furnish your settlement.

Plastic does need to be bought from caravans (you can get crates of it very cheaply) or manufactured
at the chemist, but both sources will yield large amounts of it with only a little effort. The only
drawback of plastic is that it tends to melt easy (keep away from magma!) and it isn't worth much
(about the same as wood or stone, but a lot less than metal or glass).

Hint: Plastic crates are cheap, bring a bunch on embark and you will be set for early building
material for walls and such like!

The plastics factory does require a small amount of power to run, connect your power source to any
of the workshop corners after the factory has been constructed.

Plastics are formed by your manufacturers (wax workers).

The plastics factory requires machine tools and hand tools to build, make sure you bring some on embark
or have the material to make them.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_furniture.txt&id=SLAB_ENGRAVER)

The memorial engraver allows you to engrave memorial slabs for your fallen settlers.

Sadly due to technical issues you cannot see if a listed creature has already been memorialized. I
may or may not be able to change this in the future.

Settlement inhabitants are always listed before others, but once you get a bunch of dead units it can
become very tedious to find the unit you want. Whenever possible use a coffin, it's a lot easier.

All work here is done by your engravers.

Like many of the basic buildings this requires hand tools, so remember to bring some on embark.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_bonecarver.txt&id=BONECARVER)

The bonecarver is where all bone and shell item production takes place.

Blocks can be made from a fake "bone" or "shell" stone, since bone and shell blocks cannot be stockpiled.
If you wish you can still make them from the bone or shell material directly, but don't be surprised
they stay in the workshop!

Although it is called a "bonecarver" you can also use shell to make all of the items. If you have
lots of shell (which should be the case if you do much fishing) a bonecarver can be a great way to
get cheap furniture!

Like many of the basic buildings this requires hand tools, so remember to bring some on embark.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_radio.txt&id=RADIO_STATION)

The radio station is an expensive building, both to build and to run, but it allows you to order
special items directly from core world traders in orbit.

To build a radio station you will need radio parts and large gems. Radio parts are made at the
electronics workshop or the mechanic. They are much cheaper at the mechanic, but the radio parts made
there only last for a little while before wearing out, where as those made at the electronics shop
last forever.

Running the radio station requires quite a bit of mechanical power, which can be connected to any of
the four corners of the workshop.

Core world traders handle all transactions in off-world credits. The only way to gain credits is selling
precious metals or gems. In general trading this way is expensive and only good for critical needs.

In general the only items for sale are rare tools (for example electronics assemblers), special machine
parts (crated power plants), military supplies (armor and weapons), and critical raw materials. Food,
drink, clothing, etc are not available.

All items available for sale are packed in crates, unpack these at the unloading bay.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_paper.txt&id=BOOKBINDER)

The bookbinder's workshop is where all your paper making, book binding, scroll making, etc takes place.

Paper may be made here from any fiber plant via a simple two step process.

For those who want to be low tech, scroll rollers may be made from just about anything at the workshop
associated with the material you want to make them from (so you would use the glass blower for glass rollers).

Book making works very much like it does in vanilla DF.

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_common.txt&id=BEETLE_HIVE&stage=1)

The beetle hive raises genetically engineered "resin beetles". These beetles produce small amounts of "beetle resin"
which may then be processed at the chemist to make plastic.

Once a beetle hive has been constructed it will periodically queue up "tend hive" reactions (you cannot queue these
reactions manually). When a hive tender (beekeeper) does this job they will collect any waiting resin. Not tending your
hives will waste resin, as it does not build up over time. Gather it or lose it.

Beetle hives are constructed from the workshop page with a hive and a beetle colony, vanilla hives cannot be constructed.
Beetle colonies are import only, but one you have one it will slowly produce more colonies that may be moved to new hives.

You need five resin globs to make one plastic bar, so it can take a while to get much plastic stockpiled. That said if
you start production early, in a few years you will have more plastic than you know what to do with. Exporting plastic
crafts could be a fine source of income...

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_common.txt&id=ALGAE_TANK&stage=1)

Algae tanks produce a steady stream of low quality food, good for emergencies and such.

Algae tanks are self-contained, import-only items that may be constructed anywhere you like. Once setup your hive tenders
(beekeepers) will periodically gather "algae" from the tank. Algae is pretty tasteless, and no one likes it much, but it
will keep you from starving (eating algae gives a minor bad though).

If you like you can mill the cakes into algae flour and use it for cooking, this will disguise the taste (or rather, lack
of taste).

* * *

![](/wshop?addon=First%20Landing/Base&file=building_flbase_common.txt&id=MOISTURE_CONDENSER&stage=1)

Moisture condensers are very simple, they gather moisture from the air and condense it to fill barrels for drinking.

Moisture condensers are self-contained, import-only items that may be constructed anywhere you like. Once setup your
hive tenders (beekeepers) will periodically fill barrels with water from the output.


Plants and Animals:
----------------------------------------------------------------------------------------------------

You may bring any of the following domestic animals:

* Cats - Standard house cats, brought along to keep the vermin down and as pets.
* Dogs - Brought as pets mostly, but they may be useful as guards against smaller hostiles.
* Leopards - Genetically modified and brought along for guard duty and hunting, leopards serve many
  of the same roles as dogs would on less hazardous planets.
* Donkeys - Every colony needs beasts of burden, and donkeys are small but hardy.

Any other animals will have to be provided by other addons, no wild animals are included in the base.

First Landing comes with several addons that provide random wildlife, see their addon descriptions
for details.

Plants are all provided by other addons, default grass as well as random trees and crops are all
provided. See the individual addons for details.


Fiction Primer:
----------------------------------------------------------------------------------------------------

First Landing focuses on colonization parties on alien planets, how the colonists got to the planet
and where they came from is more-or-less ignored. That said there are some guidelines I use.

The following information is for use by those who want to write fan-fiction or succession game turns,
and should be consulted to help keep things consistent. Feel free to ignore this if you want, but I
would prefer that you pay attention :)

Star travel is assumed to be FTL, and FTL enough to prevent voyages from taking more than a year or
so max, but a week to a month is common for most voyages.

It is generally assumed that all the colonies are in the same galaxy, just different planets/systems.

The more hospitable planets are generally directly colonized with a full tech base by the core worlds,
only planets that are in some way undesirable are put under colonial administration authority and
opened for general colonization. This is mostly seen as a way to get rid of undesirables and restless
portions of the population. This means that while officially all colonists are volunteers (and the
majority are), a few have actually been "volunteered" by their planetary government.  

Seed colonies are generally landed with a minimal tech base (this is officially to prevent colony
failure due to inability to maintain critical industry). Seed colonies then send out colonization
parties to found new settlements. Sometimes these settlement parties are landed directly from off
world, but most of the time they are colonists already on planet. Player settlements are satellite
colonies founded by these colonization parties.

In addition to a place to dump dissidents the core worlds mostly see the colonies as a way to make a
quick credit. These worlds may send trading expeditions to buy exotic items, native crafts, etc and
sell technology needed by the colony. This trade generally takes place at a central location (core
world traders do not want to waste time stopping at every little outpost), distribution is left to
the colony government. While some technology (like power armor and guns) can be produced on-site,
generally the tools required to produce such items are produced off-world.

Colony administrators are not hereditary, generally they are appointed or elected. The highest levels
of colonial administration planetary governments are appointed by the off-planet authorities.
Settlement administrators are generally elected, with the highest levels serving life terms.

Some planets have intelligent natives. These natives are almost universally hostile, and can be
expected to raid settlements. Groups of human raiders also like to set up shop on these planets, as
they are generally great places to hide out. These raiders will steal anything they think can be sold
for a profit, and will often siege settlements in an attempt to drive the settlers away. The natives
tend to leave the raiders alone, as settlers are seen as a greater threat (it's an enemy-of-mine-enemy
thing mostly).

The problem of incompatible biology (not being able to eat the native plants and animals) is ignored,
it is assumed that any native life on a colony planet can be eaten.

While you can eat the animals, the animals can also eat you! Most colony planets have a broad range
of carnivores that think settlers would make a tasty meal. Natives tend to be able to tame these
beasts, and often use them as war animals. Of course you can do the same, some animals can make valuable
additions to any settlement.


TODO:
----------------------------------------------------------------------------------------------------

* Test raider sieges (do they carry the correct equipment?).
* Some workshops may need to be recolored, I think I got all the major issues, but it's hard to be sure.
* The following item types need a method of manufacture:
	* INSTRUMENT (there is a promising looking RL instruments mod being made, but it doesn't have enough content yet).
	* TOTEM (how do you target skulls in a reaction?)
	* ANIMALTRAP (useless, there are no vermin worth catching)
* Change the way items are registered for production so that the item only has to be added in one place.
	* Use follower templates to assign materials to an item:
		* `{!FL_FACTORY_ITEM;<NAME>;<ITEM_TOKENS>;<CATEGORY>;<EXTRAS>=""}{!FL_WOOD;<RAW_MATERIAL_COUNT>;<PRODUCTION_COUNT>}`
* Add more words to the new creature name generator word lists so it can be used as the default.
* Something *needs* to be done about the poor implementation of the vanilla "details" menu.
	* This menu has an absolutely retarded implementation. Instead of working based on the reaction output item type
	  and the input item material specifier as would be logical, it appears to work based on job type and input item
	  *type*, mostly ignoring requested material (definitely not honoring reaction classes, possibly not even material
	  reaction products). AFAIK there is no way to make a reaction that allows specification of an art image *at all*.
	* The way this menu works totally destroys limiting materials based on reaction classes (and maybe MRPs, I need to
	  test that yet), which is a major problem for FL.
	* This menu is so totally broken as far as mods are concerned that it is tempting to figure out a way to disable it.
	* Would it be possible to override this menu?
		* Would it be possible to show a "hybrid" menu? Basically trigger a custom menu when the game tries to show a the
		  vanilla details menu and use that to trigger the vanilla menus as desired (to, for example, force-allow the art
		  image selection page for reactions that create statues, or to show the material selection menu with a list of
		  materials that respects reaction classes)
		* If a hybrid is not possible how hard would it be to completely replace the whole system with something that
		  actually works?
	* Does Toady plan to fix his newest piece of crap? If so it may be better to just wait...
Stuff to do at some point:

* Paper making from wood pulp.
* What should I do about siege engines?
	* My first inclination would be to remove them entirely, maybe replacing them with something.
* Small RTG mechanical power plants should be able to be built from an RTG module and an electric motor.
* A full set of random creature scripts.
	* Insectoids need to be done yet, maybe some day...
* Use the "hive" system with more products:
	* Fish tanks: Provides small amounts of raw fish or shell fish.
	* More ideas?
* Raider drop pods:
	1. Pick a random tile near a settler on the surface, ideally pick several over a period of time and average them.
	2. Create downward flying item projectiles ("drop pod" tools) at the top of the map flying towards the selected area.
	3. Delete the items and spawn hostile raiders (with full top-tier equipment and good skills) where each projectile hits.
	4. Let the fun begin :P
	* This will make a military more important, and also make walled compounds less safe (as raiders would tend to drop
	  inside the walls).
	* To keep this from being *too* annoying only two or three raiders will drop at once.
	* This should be quite rare (once or twice a year at most, usually less).
* I should use the "digging invaders" plugin for both raiders and natives...
* Add a "pneumatic spike thrower", basically my auto-crossbow but make it throw only special "spike" ammo.
* Animals need to have appearance (color, size, etc) modifiers applied, things like stripes, spots, etc would also be cool.
* Would it be possible to dynamically replace the in-game help via DFHack? If so this would be a good place to put detailed
  information about the random plants/animals...
	* Obviously this would only be useful if the in-game help could be replaced on a per-world basis (by restoring the
	  original data on unload).
* Instead of (or maybe in addition to) a better name generator provide a way to permanently change animal, plant, or tree
  names in game.
	* This could be facilitated by a "naturalists workshop" that allows you to run a reaction with any item made from a
	  plant or animal material. The would then trigger a DFHack script that did the actual renaming.
	* The script would have to remember any custom names so that they can be re-applied every time the world loads. This
	  is fairly simple.
	* Obviously you couldn't change any IDs without a prohibitive amount of work, but display names are (if I remember
	  correctly) easy to change on the fly.
	* For creatures it would be really cool to combine the old and new name generators so that a small percentage of
	  creatures receive "new style" names and the others get random labels (to simulate naming during worldgen).
		* Obviously any domesticated creatures should receive new style names...
* Convert the powered workshops to use the new 5x5 powered workshop generator support.
