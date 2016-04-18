
This addon generates a random mix of surface and cavern creatures, most land-bound, but some can fly.

Each creature has a "description code" that comes after it's description. This code contains some
important information about the creature that may not be present in it's normal description.

A full description code will look something like this:

	(#5/250, 5000cc, ~25, FPE)

This means that the creature is number 5 out of the 250 generated, has an average adult size of ~5000cc,
grows to adult in 1 year (1/15 of its base life span), lives 20-30 years and is a flying, predatory, egg-layer.

The string at the end of the description code signals the presence of certain traits in the creature:

* F: Flier
* V: Venomous
* P: Predator
* G: Grazer
* E: Lays eggs
* M: Milkable

If you want details about a creature look it up in the raws (`creature_planets_fauna_random_vertebrates.txt`)
using the sequence number in it's description code for fast lookup.

You can set how many creatures you want to generate using the provided configuration variable.
