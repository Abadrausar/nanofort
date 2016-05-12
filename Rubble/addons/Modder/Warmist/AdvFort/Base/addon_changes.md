    version: 0.044
    changelog:
        *0.044
        - added output to clear_jobs of number of cleared jobs
        - another failed attempt at gather plants fix
        - added track stop configuration window
        *0.043
        - fixed track carving: up/down was reversed and removed (temp) requirements because they were not working correctly
        - added checks for unsafe conditions (currently quite stupid). Should save few adventurers that are trying to work in dangerous conditions (e.g. fishing)
        - unsafe checks disabled by "-u" ir "--unsafe"
        *0.042
        - fixed (probably for sure now) the crash bug.
        - added --clear_jobs debug option. Will delete ALL JOBS!
        *0.041
        - fixed cooking allowing already cooked meals
        *0.04
        - add (-q)uick mode. Autoselects materials.
        - fixed few(?) crash bugs
        - fixed job errors not being shown in df
        *0.031
        - make forbiding optional (-s)afe mode
        *0.03
        - forbid doing anything in non-sites unless you are (-c)heating
        - a bit more documentation and tidying
        - add a deadlock fix
        *0.021
        - advfort_items now autofills items
        - tried out few things to fix gather plants
        *0.02
        - fixed axles not being able to be placed in other direction (thanks SyrusLD)
        - added lever linking
        - restructured advfort_items, don't forget to update that too!
        - Added clutter view if shop is cluttered.
        *0.013
        - fixed siege weapons and traps (somewhat). Now you can load them with new menu :)
        *0.012
        - fix for some jobs not finding correct building.
        *0.011
        - fixed crash with building jobs (other jobs might have been crashing too!)
        - fixed bug with building asking twice to input items
        *0.01
        - instant job startation
        - item selection screen (!)
        - BUG:custom jobs need stuff on ground to work
        *0.003
        - fixed farms (i think...)
        - added faster time pasing (yay for random deaths from local wildlife)
        - still hasn't fixed gather plants. but you can visit local market, buy a few local fruits/vegetables eat them and use seeds
        - other stuff
        *0.002
        - kind-of fixed the item problem... now they get teleported (if teleport_items=true which should be default for adventurer)
        - gather plants still not working... Other jobs seem to work.
        - added new-and-improved waiting. Interestingly it could be improved to be interuptable.
    todo list:
        - document everything! Maybe somebody would understand what is happening then and help me :<
        - when building trap add to known traps (or known adventurers?) so that it does not trigger on adventurer
    bugs list:
        - items blocking construction stuck the game
        - burning charcoal crashed game
        - gem thingies probably broken
        - custom reactions semibroken
        - gathering plants still broken