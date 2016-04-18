
This addon generates a random mix of mega beasts and semi-mega beasts.

Each creature has a "description code" that comes after it's description. This code contains some
important information about the creature that may not be present in it's normal description.

A full description code will look something like this:

	(#5/250, 5000cc, FE)

This means that the creature is number 5 out of the 250 generated, has an average adult size of ~5000cc,
and is a flying egg-layer.

The string at the end of the description code signals the presence of certain traits in the creature:

* M: Mega beast (else it is a semi-mega beast)
* F: Flier
* V: Venomous
* E: Lays eggs

If you want details about a creature look it up in the raws (`creature_planets_fauna_random_mega.txt`)
using the sequence number in it's description code for fast lookup.

You can set how many creatures you want to generate using the provided configuration variable.
About 25% of the creatures will be a mega beast, the others will be semi-mega beasts.
