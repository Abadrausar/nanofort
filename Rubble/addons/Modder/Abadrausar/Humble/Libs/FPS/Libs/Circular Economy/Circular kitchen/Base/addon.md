I need to make already prepared_food (using it as a reagent in a reaction) into dehydrated flour(ingredient) in the quern using some fuel by stack of dried prepared_food.
Losing value, dorf time and fuel but gaining more dorf-exp in milling and cooking and biggers stacks by lavish cooking again and with that, using some QS stockpiles; the possibility of better FPS.
My problem:
 I does not find the identifier of prepared_food, I have even generated a custom stockpile with only prepared_food enabled and looked at the generated file "prepared food.dfstock" in directory stocksettings to no avail, what can be found is:
Spoiler (click to show/hide)
Instead of something like
Spoiler (click to show/hide)
I know that some special identifiers such as those corresponding to 3 kinds of glass have been known by means of binary string dumpers or Reverse engineering. 
Finally I found it in http://dwarffortresswiki.org/index.php/DF2014:Item_token
71   FOOD    item_food.txt   Prepared meals.