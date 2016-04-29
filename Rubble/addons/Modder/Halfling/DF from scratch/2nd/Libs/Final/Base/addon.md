[Addon Change log](/Halfling___2nd_DF_from_Scratch_Final___Drow&file=addon_changes.md)
![Halfling DF from scratch](http://scout.es/wp-content/uploads/haiku2b.jpg)
## So you think that you need some help, want to do a sugestion, inform about a bug or give thanks to modders or developpers, but not sure where to go.
1. If you want to patronize Toady One and the developpment of Dwarf Fortress go to: [Support Dwarf Fortress Vanilla](http://www.bay12games.com/support.html)
2. For questions about the rubble framework go to: [Rubble framework](http://www.bay12forums.com/smf/index.php?topic=154304.0)
3. Anything about HDFPS: Humble Dwarf Fortress Publishing System is there [HDFPS](http://www.bay12forums.com/smf/index.php?topic=157300.0)
4. If your question is addressed to the actual original maintainer of this mod then you have [2nd DF from scratch](http://www.bay12forums.com/smf/index.php?topic=140715.0)
# Halfling DF from scratch: The entirely player-made universe succession.
Dwarf Fortress from Scratch v .40.x


Current raws: 264 kb vs Vanilla’s 3.08 mb


Hello!  I bet you are probably wondering what this is.

This is DF "from scratch". We start with a vanilla version of DF, delete all raws with their creatures and bodies and templates and materials, and build an entire new player-made universe by hand.

One full of player-made creatures and all-new mechanics, hand-built from contributor to contributor from the body templates and language symbols up to entire new dwarf civilizations and untold realms of possibility. 

>>Click Here to View DFFS: Sci-fi<< 

>>Click Here to View DFFS: Fantasy<< 

DFFS Sci-fi ia a totally fresh creation for V .40, placed in the grim darkness of the far future, there is only war placed on a distant mining world called 'THX 1138' by the interstellar corporations who covet it's rich and bizarre resources, while simply called Terra Nihalgai by the local tribal inhabitants. This version is unrestricted to just 14th century high fantasy, and much else other than a balanced experience.

DFFS Fantasy is a port of the work done by Halfling and contributors of the original DFFS thread. It is a world governed by the Greater Creators; competitive gods gone mad with power, constantly creating form to the world and it's inhabitants to better understand the mechanics of the cosmos, and sometimes solely for their pure amusement.    
 
For Modders

For those returning, this DFFS will be a bit different from version .34. This game will NOT be turn based. You do not need to upload anything individually anymore. The whole project is now hosted on Google Drive to keep things easy to access and edit. The privileges are set up so anyone can contribute to the game, at any time, with anything, right down to a Chromebook.  

To access the folder and start contributing, first



then

 

Since we are editing text files and not google doc files I cannot recommend this highly enough:

(click here for a handy Chrome native text editor)

It’s a simple text editor, but for our needs it will be more than fine. Now you can open text files and save them to the project on the cloud with a ctrl-s as if you were using a local version of Windows Notepad. And if you still want to use Windows Notepad, that's fine too. Just upload the entire .txt file into the 'objects' folder from your local computer files. Easy!


Spoiler: Rules for Modders (click to show/hide)
Rules:

You are not allowed to copy and paste or include DF or other mods' raws directly (I had to a little to get it to not crash and burn, but that should be the end of it). Getting caught means your turn will be voided. Copying and pasting parts that you were going to rewrite exactly anyway is fine, doing this to an entire creature, object or set of them is not
You are, however, explicitly allowed to link to previous content, such as use a previous player's templates or interactions. In fact this is probably wise when possible so the raw size doesn't balloon to insanity.
You are allowed to copy values (tensile strength of iron) from anywhere but a new one would be preferable. On the other hand coming up with new values just for the sake of it is not very sane in the long run, so use judgment.
You are allowed to edit content according to the editing rules below. If you want to implement a mechanic that must be modded into all creatures (such as something like Halfling’s contagious disease mod), ask first please.
To make the world eventually more varied, you are, in violation of the above, allowed to reduce the frequency of old creatures or materials or limit them to biomes when new creatures are added, but not remove them entirely - nor assign them to an impossible biome (badgers to the ocean). The new frequency must not be less than 10% of the old frequency, or 1, whichever is greater, and you must mark this edit with a comment to prevent repeating it.
When defining materials, try to add a [STATE_COLOR:ALL:COLOR] where color is a color of your choice. This is so stonesense can figure out what color a sprite should be.
Whenever you add anything to the mod, it is REQUIRED that you add a small note in the CHANGELOG so things can be accounted for easily. Then, please add one to the Version number.
Failure to do so will get your content rejected. Players sending in reports need reliable Versioning! Help them help you and don’t get lazy, guys.



IMPORTANT:

To force compatibility with both vanilla DF and other players' contributions, so that raws never clash, and novelty so that copypasting simply does not work, you must add _XYZ to the end of every object that you newly add/define, where XYZ is an at least 3 letter handle that you chose. Mine was HLG, so that I would add a "BODY:REALLY_SIMPLE_BODY_HLG" instead of "BODY:REALLY_SIMPLE_BODY". The exception to this is language words. Not doing this will eventually destroy everything. Referring to other players raws of course does not need this, if XYZ came before me I could use BODY_XYZ in my creature, but if I make a new body... as above.
Similarly name your files by contributor handle, eg. "creature_domestic_halfling", so we can keep track of things.

About editing other people's contributions:

After your final comment in the CHANGELOG, your work is considered relatively stable and open to scrutiny. Your content may be edited for: rules violations, balance, language/spelling errors, and bugfixes.
1. This will be implemented by any current modder/bugfixer at their discretion, but may be suggested by anyone.
- Rules violations: such as incorrectly naming your creatures without _XXX
- Language/spelling errors, without concern for whether the misspelling/bad grammar was intentional if it wasn't stated to be
- Balance issues, highlighted by play experience
- Bugfixes and seemingly unintentional behavior, such as a material that wasn't supposed to explode exploding spontaneously
The edit may be applied immediately and removed later if you disagree with it.

2. The fix must respect and preserve the original intention of the addition when it is made clear. If it was meant to be an incredibly powerful megabeast, 
it should be preserved as such, even if it's balanced. If specifically intended behavior seems to be the problem, this restriction is lifted, but this should preferably 
not be done and the thing should not have been added in the first place.

3. The author may veto any fix, but a general consensus reached after discussion will override that veto.



For Players
Just want to play the mod? Alright! Just click here:

>>Click Here to Download<< 

While you are playing, if you could you write some reports on your experience and post them in this thread? That would help out immensely. We modders, well, sometimes get lost in the code, sometimes and we need all the eyes we can to find gameplay positives, quirks, and problems. All you have to do? Please open the changelog and put the version number somewhere in your post. This is crucial to crushing bugs.

attn modders: The Google Drive will always be the most up to date project version.

The mod works with all versions of DF v .40 and works with all platforms: Windows, Mac, and linux. (At least until the newest version of DFHack comes out, but we'll cross that bridge when we get there.) 

Spoiler: Fanart (click to show/hide)
Fanart
1 2 3 4 5 6 7

If you just want to chat about the project, why not come down the the dedicated IRC? This for sure is the best way to get in touch with me and I am guaranteed to be on from  7:00 PM to 11:PM Pacific Time.

Now I’m going to be honest, I’m entrusting a lot a power and a lot of tools to this community, and a lot of freedom. I believe you can handle it, because you are the most awesome, smart, driven community i’ve ever seen. 

 Be assured, I have a lot of tools at my disposal to monitor, revert, and restore everything and I pledge my full support to moderate this project to the best of my ability, your hard work isn't going to be destroyed or lost in the void any time soon. If you think something is amiss or could be done better, please, don't hesitate to drop by the IRC so I can rectify it.
 
Last but not least, Credits!

This mod was made possible by Community members like:
DrTaco
Lost in Nowhere
Cerapter
Gnorm
lord_lemonpie
cyberTripping

Remember; it's more productive to throw things in than to ask for permission around here. Be bold and have fun. 

((This top post can be edited by the community, as with everything else))