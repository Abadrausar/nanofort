
Each creature has a "description code" that comes after it's description. This code contains some
important information about the creature that may not be present in it's normal description.

A full description code will look something like this:

	(#5/250, 5000cc, 1, 20-30)

This means that the creature is number 5 out of the 250 generated, has an average adult size of ~5000cc,
grows to adult in 1 year, and lives 20-30 years.

If you want details about a creature look it up in the raws (`creature_planets_fauna_random_aquatic.txt`)
using the sequence number in it's description code for fast lookup.

You can set how many creatures you want to generate using the provided configuration variable.
