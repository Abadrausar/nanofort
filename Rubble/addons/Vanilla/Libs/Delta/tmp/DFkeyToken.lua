-- Extensive use of functional linearized sparse arrays to describe the principal interelation between the hierarchie of keyTokenIds, not exhaustive, only those necesary to make the Delta system usable
-- non keyTokens are ignored as they are irrelevant for the structure, the incompleteness is due to the creature objects and their variations, thinking about a more complete system for those...
-- 1st array dimension are the identifiers that appears after the token OBJECT at the beginning of DF vanilla files, namespaces, indent level = 0
-- 2nd array dimension are the keyTokenIds that are possible inside each namespace, indent level = 1
-- 3rd array dimension are the nestedKeyTokenIds that are possible inside each type of object, indent level = 2
-- 4th array dimension are the nestedKeyTokenIds that are possible inside root nestedKeyTokenIds, indent level = 3
-- 5th array dimension... and so on, so on, indent level = 4
-- the integer is the variation of indent level that vanilla raws uses for each keyTokenId, being root the base namespace, when level = -1 the tiken is acting as a closing bracket of a nested zone
-- [""] = 2 ,[""] = 2 ,[""] = 2 ,[""] = 2 ,[""] = 2 ,[""] = 2 ,[""] = 2 ,[""] = 2 ,
DFkeyTokens = {} 
DFkeyTokens["OBJECT"] = {["BODY_DETAIL_PLAN"] = 0, ["BODY"] = 0, ["BUILDING"] = 0, ["CREATURE"] = 0, ["CREATURE_VARIATION"] = 0, ["DESCRIPTOR_COLOR"] = 0, ["DESCRIPTOR_PATTERN"] = 0, ["DESCRIPTOR_SHAPE"] = 0, ["ENTITY"] = 0, ["INORGANIC"] = 0, ["ITEM"] = 0, ["INTERACTION"] = 0, ["LANGUAGE"] = 0, ["MATERIAL_TEMPLATE"] = 0, ["PLANT"] = 0, ["REACTION"] = 0, ["TISSUE_TEMPLATE"] = 0, ["WORLD_GEN"] = 0, ["PROFILE"] = 0,}
DFkeyTokens["BODY_DETAIL_PLAN"] = {["BODY_DETAIL_PLAN"] = 1,}
DFkeyTokens["BODY_DETAIL_PLAN"]["BODY_DETAIL_PLAN"] = {["BP"] = 2,}
DFkeyTokens["BODY"] = {["BODY"] = 1,["BODYGLOSS"] = 1,}
DFkeyTokens["BODY"]["BODY"] = {["CATEGORY"] = 2,} 
DFkeyTokens["BUILDING"] = {["BUILDING_WORKSHOP"] = 1, ["BUILDING_SMELTER"] = 1,}
DFkeyTokens["CREATURE"] = {["CREATURE"] = 1,}
DFkeyTokens["CREATURE"]["CREATURE"] = {["ATTACK"] = 2 ,["BODY_APPEARANCE_MODIFIER"] = 2 ,["BODY_DETAIL_PLAN"] = 2 ,["CAN_DO_INTERACTION"] = 2 ,["CASTE"] = 2 ,["EXTRA_BUTCHER_OBJECT"] = 2, ["MENT_ATT_RANGE"] = 2, ["SELECT_CASTE"] = 2, ["SELECT_MATERIAL"] = 2, ["SELECT_TISSUE_LAYER"] = 2, ["SEMIMEGABEAST"] = 2 ,["SET_BP_GROUP"] = 2 ,["TL_COLOR_MODIFIER"] = 2 ,["MENT_ATT_RANGE"] = 2 ,} 
DFkeyTokens["CREATURE"]["CREATURE"]["SELECT_CASTE"] = {["SET_TL_GROUP"] = 3,["CASTE_NAME"] = 3,}
DFkeyTokens["CREATURE"]["CREATURE"]["SELECT_TISSUE_LAYER"] = {["PLUS_TISSUE_LAYER"] = 3,}
DFkeyTokens["CREATURE"]["CREATURE"]["USE_MATERIAL_TEMPLATE"] = {["SYNDROME"] = 3,}

DFkeyTokens["CREATURE_VARIATION"] = {["CREATURE_VARIATION"] = 1,}
DFkeyTokens["CREATURE_VARIATION"]["CREATURE_VARIATION"] = {["CV_CONVERT_TAG"] = 2 ,}
DFkeyTokens["DESCRIPTOR_COLOR"] = {["COLOR"] = 1,}
DFkeyTokens["DESCRIPTOR_PATTERN"] = {["COLOR_PATTERN"] = 1,}
DFkeyTokens["DESCRIPTOR_SHAPE"] = {["SHAPE"] = 1,}
DFkeyTokens["ENTITY"] = {["ENTITY"] = 1,}
DFkeyTokens["ENTITY"]["ENTITY"] = {["WEAPON"] = 2 ,["POSITION"] = 2 ,["TISSUE_STYLE"] = 2 ,}
DFkeyTokens["INORGANIC"] = {["INORGANIC"] = 1,}
DFkeyTokens["INTERACTION"] = {["INTERACTION"] = 1,}
DFkeyTokens["INTERACTION"]["INTERACTION"] = {["I_SOURCE"] = 2 ,["I_TARGET"] = 2 ,["I_EFFECT"] = 2 ,}
DFkeyTokens["ITEM"] = {["ITEM_AMMO"] = 1,["ITEM_ARMOR"] = 1,["ITEM_FOOD"] = 1,["ITEM_GLOVES"] = 1,["ITEM_HELM"] = 1,["ITEM_PANTS"] = 1,["ITEM_SIEGEAMMO"] = 1,["ITEM_SHIELD"] = 1,["ITEM_SHOES"] = 1,["ITEM_TOOL"] = 1,["ITEM_TOY"] = 1,["ITEM_TRAPCOMP"] = 1,["ITEM_WEAPON"] = 1,}
DFkeyTokens["ITEM"]["ITEM_AMMO"] =  {["ATTACK"] = 2,}
DFkeyTokens["ITEM"]["ITEM_TOOL"] =  {["ATTACK"] = 2,}
DFkeyTokens["ITEM"]["ITEM_TRAPCOMP"] =  {["ATTACK"] = 2,}
DFkeyTokens["ITEM"]["ITEM_WEAPON"] =  {["ATTACK"] = 2,}
DFkeyTokens["LANGUAGE"] = {["SYMBOL"] = 1,["TRANSLATION"] = 1,["WORD"] = 1,}
DFkeyTokens["LANGUAGE"]["WORD"] = {["ADJ"] = 2,["NOUN"] = 2,["PREFIX"] = 2,["VERB"] = 2,}
DFkeyTokens["MATERIAL_TEMPLATE"] = {["MATERIAL_TEMPLATE"] = 1,}
DFkeyTokens["PLANT"] = {["PLANT"] = 1,}
DFkeyTokens["PLANT"]["PLANT"] = {["USE_MATERIAL_TEMPLATE"] = 2,["GROWTH"] = 2 ,}
DFkeyTokens["PLANT"]["PLANT"]["USE_MATERIAL_TEMPLATE"] = {["SYNDROME"] = 3,}
DFkeyTokens["PLANT"]["PLANT"] = {["DRINK"] = -1,["MILL"] = -1 ,["SEED"] = -1 ,["BASIC_MAT"] = -1 ,["THREAD"] = -1 ,["EXTRACT_BARREL"] = -1 ,["EXTRACT_STILL_VIAL"] = -1 ,["EXTRACT_VIAL"] = -1 ,["TREE"] = -1 ,}
DFkeyTokens["REACTION"] = {["REACTION"] = 1,}
DFkeyTokens["REACTION"]["REACTION"] = {["REAGENT"] = 2 ,["PRODUCT"] = 2 ,}
DFkeyTokens["TISSUE_TEMPLATE"] = {["TISSUE_TEMPLATE"] = 1,}
DFkeyTokens["WORLD_GEN"] = {["WORLD_GEN"] = 1,}
DFkeyTokens["PROFILE"] = {["PROFILE"] = 1,}
