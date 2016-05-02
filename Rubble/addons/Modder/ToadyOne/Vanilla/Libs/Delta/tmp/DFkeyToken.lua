-- Extensive use of functional linearized sparse arrays to describe the principal interelation between the hierarchie of keyTokenIds, not exhaustive, only those necesary to make the Delta system usable
-- non keyTokens are ignored as they are irrelevant for the structure, the incompleteness is due to the creature objects and their variations, thinking about a more complete system for those...
-- 1st array dimension are the identifiers that appears after the token OBJECT at the beginning of DF vanilla files, namespaces, indent level = 0
-- 2nd array dimension are the keyTokenIds that are possible inside each namespace, indent level = 1
-- 3rd array dimension are the nestedKeyTokenIds that are possible inside each type of object, indent level = 2
-- 4th array dimension are the nestedKeyTokenIds that are possible inside root nestedKeyTokenIds, indent level = 3
-- 5th array dimension... and so on, so on, indent level = 4
-- the integer is the variation of indent level that vanilla raws uses for each keyTokenId, being root the base namespace, when level = -1 the tiken is acting as a closing bracket of a nested zone
-- [""] = 2 ,[""] = 2 ,[""] = 2 ,[""] = 2 ,[""] = 2 ,[""] = 2 ,[""] = 2 ,[""] = 2 ,
Key = {} 
Key["BODY_DETAIL_PLAN"] = {["OBJECT"] = 0,["BODY_DETAIL_PLAN"] = 1,["CREATURE"] = 2,}
Key["BODY"] = {["OBJECT"] = 0,["BODY"] = 1,}
Key["BODYGLOSS"] = {["BODY"] = 1,}
Key["BP"] = {["BODY"] = 2,}
Key["BUILDING"] = {["OBJECT"] = 0,}
Key["BUILDING_WORKSHOP"] = {["BUILDING"] = 1,}
Key["BUILDING_SMELTER"] = {["BUILDING"] = 1,}
Key["CREATURE"] = {["OBJECT"] = 0,["CREATURE"] = 1,}
Key["ATTACK"] = {["CREATURE"] = 2,}
Key["BODY_APPEARANCE_MODIFIER"] = {["CREATURE"] = 2,}
Key["CAN_DO_INTERACTION"] = {["CREATURE"] = 2,}
Key["CASTE"] = {["CREATURE"] = 2,}
Key["EXTRA_BUTCHER_OBJECT"] = {["CREATURE"] = 2,}
Key["MENT_ATT_RANGE"] = {["CREATURE"] = 2,}
Key["SELECT_CASTE"] = {["CREATURE"] = 2,}
Key["SELECT_MATERIAL"] = {["CREATURE"] = 2,}
Key["SELECT_TISSUE_LAYER"] = {["CREATURE"] = 2,}
Key["SEMIMEGABEAST"] = {["CREATURE"] = 2,}
Key["SET_BP_GROUP"] = {["CREATURE"] = 2,}
Key["TL_COLOR_MODIFIER"] = {["CREATURE"] = 2,}
Key["USE_MATERIAL_TEMPLATE"] = {["CREATURE"] = 2,}
Key["SET_TL_GROUP"] = {["SELECT_CASTE"] = 3,}
Key["CASTE_NAME"] = {["SELECT_CASTE"] = 3,}
Key["PLUS_TISSUE_LAYER"] = {["SELECT_TISSUE_LAYER"] = 3,}
Key["SYNDROME"] = {["USE_MATERIAL_TEMPLATE"] = 3,}
Key["CREATURE_VARIATION"] = {["OBJECT"] = 0,["CREATURE_VARIATION"] = 1,}
Key["CV_CONVERT_TAG"] = {["CREATURE_VARIATION"] = 2,}
Key["DESCRIPTOR_COLOR"] = {["OBJECT"] = 0,}
Key["COLOR"] = {["DESCRIPTOR_COLOR"] = 1,}
Key["DESCRIPTOR_PATTERN"] = {["OBJECT"] = 0,}
Key["COLOR_PATTERN"] = {["DESCRIPTOR_PATTERN"] = 1,}
Key["DESCRIPTOR_SHAPE"] = {["OBJECT"] = 0,}
Key["SHAPE"] = {["DESCRIPTOR_SHAPE"] = 1,}
Key["ENTITY"] = {["OBJECT"] = 0,["ENTITY"] = 1,}
Key["WEAPON"] = {["ENTITY"] = 2,}
Key["POSITION"] = {["ENTITY"] = 2,}
Key["TISSUE_STYLE"] = {["ENTITY"] = 2,}
Key["INORGANIC"] = {["OBJECT"] = 0,["INORGANIC"] = 1,}
Key["ITEM"] = {["OBJECT"] = 0,}
Key["ITEM_AMMO"] = {["ITEM"] = 1,}
Key["ITEM_ARMOR"] = {["ITEM"] = 1,}
Key["ITEM_FOOD"] = {["ITEM"] = 1,}
Key["ITEM_GLOVES"] = {["ITEM"] = 1,}
Key["ITEM_HELM"] = {["ITEM"] = 1,}
Key["ITEM_PANTS"] = {["ITEM"] = 1,}
Key["ITEM_SIEGEAMMO"] = {["ITEM"] = 1,}
Key["ITEM_SHIELD"] = {["ITEM"] = 1,}
Key["ITEM_SHOES"] = {["ITEM"] = 1,}
Key["ITEM_TOOL"] = {["ITEM"] = 1,}
Key["ITEM_TOY"] = {["ITEM"] = 1,}
Key["ITEM_TRAPCOMP"] = {["ITEM"] = 1,}
Key["ITEM_WEAPON"] = {["ITEM"] = 1,}
Key["ATTACK"] = {["ITEM_AMMO"] = 2,["ITEM_TOOL"] = 2,["ITEM_TRAPCOMP"] = 2,["ITEM_WEAPON"] = 2,}
Key["INTERACTION"] = {["OBJECT"] = 0,["INTERACTION"] = 1,}
Key["I_SOURCE"] = {["INTERACTION"] = 2,}
Key["I_TARGET"] = {["INTERACTION"] = 2,}
Key["I_EFFECT"] = {["INTERACTION"] = 2,}
Key["LANGUAGE"] = {["OBJECT"] = 0,}
Key["SYMBOL"] = {["LANGUAGE"] = 1,}
Key["TRANSLATION"] = {["LANGUAGE"] = 1,}
Key["WORD"] = {["LANGUAGE"] = 1,}
Key["ADJ"] = {["WORD"] = 2,}
Key["NOUN"] = {["WORD"] = 2,}
Key["PREFIX"] = {["WORD"] = 2,}
Key["VERB"] = {["WORD"] = 2,}
Key["MATERIAL_TEMPLATE"] = {["OBJECT"] = 0,["MATERIAL_TEMPLATE"] = 1,}
Key["SYNDROME"] = {["MATERIAL_TEMPLATE"] = 2,}
Key["PLANT"] = {["OBJECT"] = 0,["PLANT"] = 1,}
Key["GROWTH"] = {["PLANT"] = 2,}
Key["USE_MATERIAL_TEMPLATE"] = {["PLANT"] = 2,}
Key["DRINK"] = {["PLANT"] = -1,}
Key["MILL"] = {["PLANT"] = -1,}
Key["SEED"] = {["PLANT"] = -1,}
Key["BASIC_MAT"] = {["PLANT"] = -1,}
Key["THREAD"] = {["PLANT"] = -1,}
Key["EXTRACT_BARREL"] = {["PLANT"] = -1,}
Key["EXTRACT_STILL_VIAL"] = {["PLANT"] = -1,}
Key["EXTRACT_VIAL"] = {["PLANT"] = -1,}
Key["TREE"] = {["PLANT"] = -1,}
Key["REACTION"] = {["OBJECT"] = 0,["REACTION"] = 1,}
Key["PRODUCT"] = {["REACTION"] = 2,}
Key["REAGENT"] = {["REACTION"] = 2,}
Key["TISSUE_TEMPLATE"] = {["OBJECT"] = 0,}
Key["WORLD_GEN"] = {["OBJECT"] = 0,}
Key["PROFILE"] = {["OBJECT"] = 0,}
Key["TISSUE_TEMPLATE"] = {["TISSUE_TEMPLATE"] = 1,}
Key["WORLD_GEN"] = {["WORLD_GEN"] = 1,}
Key["PROFILE"] = {["PROFILE"] = 1,}
Link = {} 
Link["BODY_DETAIL_PLAN"] = {"OBJECT" ,"BODY_DETAIL_PLAN","CREATURE",}
Link["BODY"] = {"OBJECT","BODY",}
Link["BODYGLOSS"] = {"BODY",}
Link["BP"] = {"BODY",}
Link["BUILDING"] = {"OBJECT",}
Link["BUILDING_WORKSHOP"] = {"BUILDING",}
Link["BUILDING_SMELTER"] = {"BUILDING",}
Link["CREATURE"] = {"OBJECT","CREATURE",}
Link["ATTACK"] = {"CREATURE",}
Link["BODY_APPEARANCE_MODIFIER"] = {"CREATURE",}
Link["CAN_DO_INTERACTION"] = {"CREATURE",}
Link["CASTE"] = {"CREATURE",}
Link["EXTRA_BUTCHER_OBJECT"] = {"CREATURE",}
Link["MENT_ATT_RANGE"] = {"CREATURE",}
Link["SELECT_CASTE"] = {"CREATURE",}
Link["SELECT_MATERIAL"] = {"CREATURE",}
Link["SELECT_TISSUE_LAYER"] = {"CREATURE",}
Link["SEMIMEGABEAST"] = {"CREATURE",}
Link["SET_BP_GROUP"] = {"CREATURE",}
Link["TL_COLOR_MODIFIER"] = {"CREATURE",}
Link["USE_MATERIAL_TEMPLATE"] = {"CREATURE",}
Link["SET_TL_GROUP"] = {"SELECT_CASTE",}
Link["CASTE_NAME"] = {"SELECT_CASTE",}
Link["PLUS_TISSUE_LAYER"] = {"SELECT_TISSUE_LAYER",}
Link["SYNDROME"] = {"USE_MATERIAL_TEMPLATE",}
Link["CREATURE_VARIATION"] = {"OBJECT", "CREATURE_VARIATION",}
Link["CV_CONVERT_TAG"] = {"CREATURE_VARIATION",}
Link["DESCRIPTOR_COLOR"] = {"OBJECT",}
Link["COLOR"] = {"DESCRIPTOR_COLOR",}
Link["DESCRIPTOR_PATTERN"] = {"OBJECT",}
Link["COLOR_PATTERN"] = {"DESCRIPTOR_PATTERN",}
Link["DESCRIPTOR_SHAPE"] = {"OBJECT",}
Link["SHAPE"] = {"DESCRIPTOR_SHAPE",}
Link["ENTITY"] = {"OBJECT","ENTITY",}
Link["WEAPON"] = {"ENTITY",}
Link["POSITION"] = {"ENTITY",}
Link["TISSUE_STYLE"] = {"ENTITY",}
Link["INORGANIC"] = {"OBJECT", "INORGANIC",}
Link["ITEM"] = {"OBJECT",}
Link["ITEM_AMMO"] = {"ITEM",}
Link["ITEM_ARMOR"] = {"ITEM",}
Link["ITEM_FOOD"] = {"ITEM",}
Link["ITEM_GLOVES"] = {"ITEM",}
Link["ITEM_HELM"] = {"ITEM",}
Link["ITEM_PANTS"] = {"ITEM",}
Link["ITEM_SIEGEAMMO"] = {"ITEM",}
Link["ITEM_SHIELD"] = {"ITEM",}
Link["ITEM_SHOES"] = {"ITEM",}
Link["ITEM_TOOL"] = {"ITEM",}
Link["ITEM_TOY"] = {"ITEM",}
Link["ITEM_TRAPCOMP"] = {"ITEM",}
Link["ITEM_WEAPON"] = {"ITEM",}
Link["ATTACK"] = {"ITEM_AMMO","ITEM_TOOL","ITEM_TRAPCOMP","ITEM_WEAPON",}
Link["INTERACTION"] = {"OBJECT","INTERACTION",}
Link["I_SOURCE"] = {"INTERACTION"}
Link["I_TARGET"] = {"INTERACTION",}
Link["I_EFFECT"] = {"INTERACTION",}
Link["LANGUAGE"] = {"OBJECT",}
Link["SYMBOL"] = {"LANGUAGE",}
Link["TRANSLATION"] = {"LANGUAGE",}
Link["WORD"] = {"LANGUAGE",}
Link["ADJ"] = {"WORD",}
Link["NOUN"] = {"WORD",}
Link["PREFIX"] = {"WORD",}
Link["VERB"] = {"WORD",}
Link["MATERIAL_TEMPLATE"] = {"OBJECT","MATERIAL_TEMPLATE",}
Link["SYNDROME"] = {"MATERIAL_TEMPLATE",}
Link["PLANT"] = {"OBJECT","PLANT",}
Link["GROWTH"] = {"PLANT",}
Link["USE_MATERIAL_TEMPLATE"] = {"PLANT",}
Link["DRINK"] = {"PLANT",}
Link["MILL"] = {"PLANT",}
Link["SEED"] = {"PLANT",}
Link["BASIC_MAT"] = {"PLANT",}
Link["THREAD"] = {"PLANT",}
Link["EXTRACT_BARREL"] = {"PLANT",}
Link["EXTRACT_STILL_VIAL"] = {"PLANT",}
Link["EXTRACT_VIAL"] = {"PLANT",}
Link["TREE"] = {"PLANT",}
Link["REACTION"] = {"OBJECT","REACTION",}
Link["PRODUCT"] = {"REACTION",}
Link["REAGENT"] = {"REACTION",}
Link["TISSUE_TEMPLATE"] = {"OBJECT",}
Link["WORLD_GEN"] = {"OBJECT",}
Link["PROFILE"] = {"OBJECT",}
Link["TISSUE_TEMPLATE"] = {"TISSUE_TEMPLATE",}
Link["WORLD_GEN"] = {"WORLD_GEN",}
Link["PROFILE"] = {"PROFILE",}

-- sparse array linear iterator
function ipairs_sparse(t)
   local tmpIndex = {}         -- tmpIndex will hold sorted indices, otherwise
   local index, _ = next(t)   -- this iterator would be no different from pairs iterator
   while index do
      tmpIndex[#tmpIndex+1] = index
      index, _ = next(t, index)
   end
   table.sort(tmpIndex)   -- sort table indices
   local j = 1
   return function()
      local i = tmpIndex[j]   -- get index value
      j = j + 1
      if i then
         return i, t
      end
   end
end
-- delta_object_cursor = ""
-- delta_attach_after = true
-- delta_nestedKey_cursor = ""
-- {@CONTEXT;PLANT:SINGLE-GRAIN_WHEAT;;USE_MATERIAL_TEMPLATE:MILL}
-- {@DELTA;MATERIAL_VALUE;25}
delta_object_cursor = "PLANT:SINGLE-GRAIN_WHEAT"
delta_attach_after = true
delta_nestedKey_cursor = "USE_MATERIAL_TEMPLATE:MILL"
print(Key["PLANT"])
print(Key["PLANT"]["PLANT"])
print(Key["PLANT"]["OBJECT"])
print(Key["BODY_DETAIL_PLAN"]["CREATURE"])
print(Key["BODY_DETAIL_PLAN"][1])
print(Link["BODY_DETAIL_PLAN"][1])
print(Link["WORD"][1])
print(Key["WORD"][Link["WORD"][1]])

for i, v in ipairs_sparse(Key["BODY_DETAIL_PLAN"]) do -- this should act like a sentinel to restric operation to the nested subobject tokens identified by the context
  print(i)
  print(v)
--  print(Link[i])
--  print(Key["BODY_DETAIL_PLAN"][Link[v]])
end    