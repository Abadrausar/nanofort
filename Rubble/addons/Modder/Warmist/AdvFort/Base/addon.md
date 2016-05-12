# Adventurer fort reborn! in http://www.bay12forums.com/smf/index.php?topic=123944.0
« on: March 15, 2013, 07:06:15 am »
ReplyQuote
Advfort is a tool in dfhack that allows for FULL* experience of fort mode jobs. That means you can:brew, mine, chop trees, dig, carve ramps, smooth and engrave, make crafts (including weapons, statues, decorating stuff, etc...), farm, gather plants and so on.
This will be a short tutorial how to use new dfhack mode/tool: advfort. 
Setup:
get dfhack.
get this save into hack/scripts/gui/advfort.lua
get this save into hack/scripts/gui/advfort_items.lua
set up keybinding, by adding to dfhack.init something like: 'keybinding add Ctrl-T "gui/advfort"'
(optional) <will write later how to install this: here basically some helpful stuff>
start playing (best with axeman because you get axe that is needed to cut the trees down)
Playing:
Now you should be able to invoke the adv-fort overlay by pressing "Ctrl-T". It should look like this:
Spoiler (click to show/hide)
You can change the job with "R" and "T". And get ingame help with "?" (there is more help there).
Spoiler: Trees (click to show/hide)
Now you have an adventurer with an axe and that strange gui on. What is next? Chopping trees! Select (with capital 'R' and 'T') job FellTree and walk near any tree. When close to the tree "careful move" into the tree (default setting alt+movement). It should indicate that it is working( with a timer how much is left) and the tree should fall down. You will get some experience as a wood cutter.
Spoiler: Fishing (click to show/hide)
This is a fun activity to have. Because dwarves fish either with bare hands or using their beards you will not need any tool. Again select the "Fish" job. After some time it should catch a water animal of sorts. This is not yet edible.
This can fail (as it can in fort mode) by saying that there is nothing to catch in "some area".
Spoiler: Building (click to show/hide)
Remember that tree that we cut down? Now we can use it to build something. Maybe a fishery to process that caught fish? Select "Build" job. This is kind of a special job. By "careful walking" in any direction you will be asked what you want to build. Select workshop->fishery. Now one step (technical limitation, df does not update the view) to see how it planned the shop. To finish building you will need some materials. Either drop the log (or blocks) down by your feet or just have it somewhere nearby. Again "careful walk" into the planned shop. Now it should show "working(-1)". It means that your adventurer is doing something that is not exactly the build step (e.g. carrying the materials). Skip steps (default keybinding ".") till you start building the shop.
<rest of tutorial is todo for now.>

Limitations: you need to have tools and materials for jobs (like in fort mode). Doing jobs sometimes get interupted by thirst, or drowsiness or hunger. In that case the fort mode logic kicks in and your adventurer seeks food/water or just drops to sleep. In the seeking case the working timer is (-1) and you need to press "skip turn" for adventurer to find what he wants (or just eat/drink/sleep and retry).
*- almost full... need to finish some stuff (e.g. traps are a bit complicated

# More detailed guide by Isher in http://www.bay12forums.com/smf/index.php?topic=123944.msg5873164#msg5873164
Re: Adventurer fort reborn!
« Reply #161 on: December 12, 2014, 07:45:37 pm »
ReplyQuote
Guide for Getting Started with AdvFort:

1) Option A: Install DFHack using these install instructions.
    Option B: Download and use the starter pack if you have Windows. There are others for Mac and Linux here but I can't guarantee they'll work with this and anyways this entire guide is Windows-specific, sorry.

2) Save this into hack/scripts/gui/advfort.lua (right click, save link as), overwriting any existing advfort.lua

3) Save this into hack/scripts/gui/advfort_items.lua (right click, save link as), overwriting any existing advfort_items.lua

4) In the main DF folder, find dfhack.init and add the following near the top (for convenience's sake) like above Lethosor's scripts:
Code: [Select]
# keybinding for AdvFort
keybinding add Ctrl-T "gui/advfort"

5) Create file called init.lua (e.g. create empty text file and rename to init.lua) in df/raw/ folder if it does not already exist (it probably won't). There's usually not many other files in this folder, just folders like the objects folder. Don't put it in the objects folder.

6) Follow these instructions. This can't be much simpler, because if you were given raw files to download they could overwrite other existing mods and you would be sad.

7) Create a new world. These things didn't apply to my current worlds, only new worlds, but then maybe I messed something up. To be safe, just do it.

8) <-- (#@*%ing emoticons) If you want immediate dig ability without having to luck into finding some guy's pick: When in game, go to DFHack and type createitem WEAPON:ITEM_WEAPON_PICK INORGANIC:IRON 1     All capital letters is important! It will appear in same space as you.

9) If you did not start as axeman and you want immediate wood cutting ability and don't want to wander around for 10 minutes to find an axe: When in game, go to DFHack and type createitem WEAPON:ITEM_WEAPON_AXE_BATTLE INORGANIC:IRON 1     All capital letters is important! It will appear in same space as you.

10) When in game, press Ctrl-T to open the menu. To get out just hit esc. If you see all black you are probably using TWBT, either you installed it yourself or you're using the starter pack. TWBT makes other levels visible, it's totally worth it when you're not using AdvFort. Anyway, you need to go to the graphics tab in the Lazy Newb Pack window and choose print mode: 2D before starting the game.

11) The very first thing you should do is press ? while in ctrl-T menu - this will explain how to do stuff in up or down directions. Don't go pressing alt-5 on numpad and thinking you'll carve a ramp down - alt 8 or 2 will careful move and do jobs in that direction, but ctrl-5 is down and ctrl-e is up. If you alt-5 it'll be all like "you need a wall!" which does not mean something to either side, it means wall/floor/ceiling. Also, if you have another item equipped as a weapon even with a pick out it won't work. You gotta have only the pick. And for a down staircase you create a down staircase where you are (just hit 5) then an updown staircase below it (ctrl-5), etc.

12) To change jobs you do not hit R or T, you hit shift-R or shift-T. These are the defaults. You'll need the correct tools for the jobs. See ITEM_WEAPON_PICK up there in that thing I said to do? If you want the tools without mining/creating in workshop, go to the page for said tool on DF wiki, scroll down to raws, and find the correct raw. 

13) Warmist has a good rundown on how to do a buncha stuff. First go to DFHack and type createitem FIGURINE INORGANIC:IRON 1   (caps are important), then grab the figurine (make sure it isn't in your backpack) hit x and claim a proper site, then do that stuff. Oh wait, having trouble with, say, workshop? You gotta have the mats (and you'll need the anvil reaction for a smithy), then you gotta make the workshop, that only makes plans so you gotta careful move into it again to actually build it, then press . (period) once to go through the x amount of moves to do the job. Then you gotta press tab to use the workshop to build stuffs.


# This mod is Warmist Advfort, that is a tool in dfhack that allows for FULL* experience of fort mode jobs in adventure mode.
Remark however that the objects can be in unusual files or places for the reasons explained.
[Addon Change log](/addonfile?addon=Adventure___Warmist&file=addon_changes.md)
## So you think that you need some help, want to do a suggestion, inform about a bug or give thanks to modders or developers, but not sure where to go.
0. If your question is addressed to the actual original maintainer of this mod then you have [Dwarf Fortress AdvFort](http://www.bay12forums.com/smf/index.php?topic=123944.0)
1. If you want to patronize Toady One and the development of Dwarf Fortress go to: [Support Dwarf Fortress](http://www.bay12games.com/support.html)
2. For questions about the rubble framework go to: [Rubble framework](http://www.bay12forums.com/smf/index.php?topic=154304.0)
3. Anything about HDFPS: Humble Dwarf Fortress Publishing System is there [HDFPS](http://www.bay12forums.com/smf/index.php?topic=157300.0)
![Dwarf Fortress AdvFort Haiku](https://github.com/Abadrausar/nanofort/blob/master/Rubble/addons/Modder/ToadyOne/Vanilla/Base/Toady_One___Vanilla.jpg)