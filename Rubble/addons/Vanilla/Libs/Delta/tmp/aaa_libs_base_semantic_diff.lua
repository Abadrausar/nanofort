-----------------------------------------------------------------------------------------------------------------------------------------------
--												The precursors actually deprecated
-- {!TEMPLATE;@SET_VALUE;token;value;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
-- {!TEMPLATE;@ATTACH_BEFORE_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};%{preraws}[%{token}:%{value}]}}
-- {!TEMPLATE;@ATTACH_AFTER_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};[%{token}:%{value}]%{preraws}}}
-- {!TEMPLATE;@REPLACE_TAG;token;preraws;{@SHARED_OBJECT_REPLACE_TAG;$OBJECT_CURSOR;%{token};%{preraws}}}
-------------------------------------------------------------------------------------------------------------------------------------------
--											Semantic Delta operators
--		{@CONTEXT;objectId;insertBeforeOrAfterContext;nestedKeyTokenId}
--	Fixes the objectid (for example ENTITY:MOUNTAIN) where the delta will be applied and also if the appended preraws go before or after (this is the default) the attachment token.
--		{@DELTA;tokenId;value;preraws}
--	Modifies the end value of a token inside an object and  eventually attach before or after it the preraws, if the token id is given but not the value the token is killed.
--	if not tokenId is given the preraws are appended at the end of the object. When no object is found a detailed DELTA warning is generated.
--	When nestedKeyTokenId is used the correct tokenId is selected taking the first tokenId between nestedKeyTokenId and the next nestedKeyToken,
--- if the requested tokenId is not found between those limits, a warning of unapplied delta is generated.
-------------------------------------------------------------------------------------------------------------------------------------------

-- There were only two minor issues that prevented the DELTA template from working properly. Both were easy fixes once
-- found :)
-- 
-- These templates (and maybe a few other supporting ones using the same basic system) would make a good standard library
-- addon at some point in the not-too-distant future. Keep polishing and tweaking!
-- 
-- I spent a few hours going over this code line by line (instead of working on the more boring thing I should have been
-- working on). The code quality is very good, particularly since you are working with an API that is, at this level,
-- mostly undocumented. Most of the things I changed were trivial, not really worth worrying about.
-- 
-- I didn't really realize how OCD (https://en.wikipedia.org/wiki/Obsessive%E2%80%93compulsive_disorder) ;) I am about code until I started going through someone else's...

-- Lets save "rubble" for standard and default stuff shall we? If these become standard they should use "rubble.libs_diff"
-- or something, but not yet.

function ipairs_sparse(t)
	local tmpIndex = {}			-- tmpIndex will hold sorted indices, otherwise
	local index, _ = next(t)	-- this iterator would be no different from pairs iterator
	while index do
		tmpIndex[#tmpIndex+1] = index
		index, _ = next(t, index)
	end
	table.sort(tmpIndex)	-- sort table indices
	local j = 1
	return function()
		local i = tmpIndex[j]	-- get index value
		j = j + 1
		if i then
			return i, t[i]
		end
	end
end
DFobjects = {"BODY_DETAIL_PLAN", "BODY", "BUILDING", "CREATURE_VARIATION", "CREATURE", "DESCRIPTOR_COLOR", "COLOR_PATTERN, SHAPE, ENTITY, INORGANIC, ITEM_AMMO, ITEM_ARMOR, ITEM_FOOD, ITEM_GLOVES, ITEM_HELM, ITEM_PANTS, ITEM_SHIELD, ITEM_SHOES, ITEM_SIEGEAMMO, ITEM_TOOL, ITEM_TOY, ITEM_TRAPCOMP, ITEM_WEAPON, TRANSLATION, SYMBOL, WORD, MATERIAL_TEMPLATE, PLANT, REACTION, TISSUE_TEMPLATE, WORLD_GEN, PROFILE, }
DFkeyTokens = {"BODY_DETAIL_PLAN", "BODY", "BUILDING_WORKSHOP", CREATURE_VARIATION, CREATURE, COLOR, COLOR_PATTERN, SHAPE, ENTITY, INORGANIC, ITEM_AMMO, ITEM_ARMOR, ITEM_FOOD, ITEM_GLOVES, ITEM_HELM, ITEM_PANTS, ITEM_SHIELD, ITEM_SHOES, ITEM_SIEGEAMMO, ITEM_TOOL, ITEM_TOY, ITEM_TRAPCOMP, ITEM_WEAPON, TRANSLATION, SYMBOL, WORD, MATERIAL_TEMPLATE, PLANT, REACTION, TISSUE_TEMPLATE, WORLD_GEN, PROFILE, }
	nestedKeyTokens= {BP, CV_CONVERT_TAG, BODY_DETAIL_PLAN, ATTACK, CASTE, SELECT_CASTE, WEAPON, POSITION, TISSUE_STYLE, NOUN, ADJ, VERB, PREFIX, SYNDROME, USE_MATERIAL_TEMPLATE, GROWTH, REAGENT, PRODUCT,}
local whereAreTheNestedKeys = {[BODY][BP] = true, [CREATURE_VARIATION][CV_CONVERT_TAG] = true, [CREATURE][BODY_DETAIL_PLAN] = true, [CREATURE][ATTACK] = true, [CREATURE][CASTE] = true, [CREATURE][SELECT_CASTE] = true, [CREATURE][CAN_DO_INTERACTION] = true, [CREATURE][BODY_APPEARANCE_MODIFIER] = true, [CREATURE][SET_BP_GROUP] = true, [CREATURE][PLUS_TL_GROUP] = true,  [CREATURE][SELECT_TISSUE_LAYER] = true, [CREATURE][SET_BP_GROUP] = true, [CREATURE][TL_COLOR_MODIFIER] = true, [ENTITY][WEAPON] = true, [ENTITY][POSITION] = true, [ENTITY][TISSUE_STYLE] = true, [ITEM][ATTACK] = true, [LANGUAGE][SYMBOL] = true, [LANGUAGE][NOUN] = true, [LANGUAGE][ADJ] = true, [LANGUAGE][VERB] = true, [LANGUAGE][PREFIX] = true, [PLANT][USE_MATERIAL_TEMPLATE] = true, [PLANT][GROWTH] = true, [REACTION][REAGENT] = true, [REACTION][PRODUCT] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true, [][] = true,}

delta_object_cursor = ""
delta_attach_after = true
delta_nestedKey_cursor = ""

-- Changed to new-style template bodies. Error messages tend to be a tiny (but sometimes very helpful) bit better this way.
-- Remember: These functions may not have any upvalues aside from _ENV (which is automatically manged by the VM).
local socontext = function(objectId, insertBeforeOrAfterContext, nestedKeyTokenId)
	objectId, insertBeforeOrAfterContext, nestedKeyTokenId = rubble.expandargs(objectId, insertBeforeOrAfterContext, nestedKeyTokenId)
	-- Expanding variables here was unnecessary, rubble.targs already handled that (and now rubble.expandargs does instead)
	delta_object_cursor = objectId
	delta_nestedKey_cursor = nestedKeyTokenId
	-- If statements that set a boolean are an abomination :P
	-- In addition to eliminating the if statement of which we will not speak, I also added some minor flexibility.
	delta_attach_after = ({["before"] = true, ["after"] = true})[string.lower(after)] ~= nil
end
rubble.template("!CONTEXT", socontext) -- <- Bad idea. Can lead to evaluation order problems (such as plague !SHARED_OBJECT_DUPLICATE). Not saying this should be removed, but the docs need an extra warning for this version.
rubble.template("@CONTEXT", socontext) -- <- Also a bad idea, but maybe not as bad as the other one...
rubble.template("CONTEXT", socontext)
rubble.template("#CONTEXT", socontext)
rubble.template("#C", socontext) -- <- This abbreviation is probably a bad idea.
-- For one thing "C" is the comment template, for another such abbreviations should probably be reserved for certain
-- very well known standard templates (there is too much chance of collision otherwise). Normally I would also complain
-- about the template names not being namespaced in some way, but as it is very likely that these templates will become
-- standard at some point in the future it is better that they are not.
-- 
-- While namespacing is a bad idea in this case, the current names may be too simple for the standard library. Something
-- more (specifically) descriptive may be in order.

local sodelta = function(tokenId, value, preraws)
	tokenId, value, preraws = rubble.expandargs(tokenId, value, preraws)
	
	-- Added this line to make sure "preraws" is a string (for later concatenation).
	-- This is needed because the new-style template bodies leave optional parameters that are not provided set to nil.
	-- I thought that rubble.expandargs took care of this little detail, but it doesn't.
	-- In v8 rubble.expandargs always returns strings, so this line may be removed then.
	preraws = preraws or ""

	if tokenId == "" then
		-- Use the function, Luke! <- Milo 
		-- Aba -> Star Wars citations are the best ;) 
		rubble.libs_base.sharedobject_add(delta_object_cursor, preraws)
		return
	end
	
	-- Moved the warning message here. Where it was it would print for every single tag that did not match the token,
	-- when clearly the effect you wanted was for it to print only if the tag you wanted could not be found.
	-- I split the warning into three parts, one here that triggers if the object does not exist, one later that triggers
	-- if the tag does not exist, and one that triggers if more than one tag matches. Change to taste :)
	if rubble.registry["Libs/Base:!SHARED_OBJECT"].table[delta_object_cursor] == nil or (delta_nestedKey_cursor ~= "" and rubble.registry["Libs/Base:!SHARED_OBJECT"].table[delta_nestedKey_cursor] == nil) then
		-- Is there a particular reason why this doesn't cause an abort? <- Milo
		-- Aba ->The delta are intended to be applied most of the time over Toady raws, from one version to the next he can decide to deprecate some objets.
		-- From the modder point of view it is way more util to have a list of ALL those rejects indicating the file and line number of each one of them
		-- than being obligated to do 100 independent Rubble generations to correct eventually one by one, these hypothetical 100 delta rejects.
		-- Rubble should never abort, except if the integrity of his internal state is compromised without the shadow of a doubt, and there, we are not in this case.
		-- A delta non applied must be logged and investigated by the modder, but does not compromise the rubble state, the only thing that rubble should do is continue
		-- and offer to the modder, all the disponible info to help the modder to do his work!
		rubble.warning("Object: \""..delta_object_cursor.."\" not found,  {@DELTA;"..tokenId..";"..value.."}: Malformed delta operator call or incorrect order of addon application.\n")
--		re-added	{@DELTA;"..token..";"..value.."} to give to the modder the most complete info about the failed delta application, 
--      Would I know how, I would also added the file and line where the unapplied delta is, and even log it also to a special file of delta rejects.
		delta_nestedKey_cursor = ""
		return
	end

	local found_object, found_nestedKey, warned = false, false, false
	-- |
	-- V Changed this line to create a new variable instead of reusing "token" (since the original version of "token" is needed later)
	local stoken, nested = string.split(tokenId, ":"), string.split(delta_nestedKey_cursor, ":")
	rubble.libs_base.sharedobject_walk(delta_nestedKey_cursor, function(tag)
		if rubble.libs_base.matchtag(tag, nested) then
			if found_nestedKey and not warned then
				rubble.libs_base.sharedobject_walk(delta_object_cursor, function(tag)
					if rubble.libs_base.matchtag(tag, stoken) then
						if found_object and not warned then
--	this should not be a problem for the modder in most ocasions, because he have the possibility to indicate how many of the params of a token form the lvalue.
							rubble.warning("Multiple tags fitting the specifier: \""..tokenId.."\" Found in object: \""..delta_object_cursor.."\". This may indicate an improper use of the delta operator.\n")
							warned = true
						end
						if value == "" then
							local preraws = "-"..tag.ID
							for _, v in ipairs(tag.Params) do
					preraws = preraws..":"..v
				end
								tag.Comments = preraws.."-"..tag.Comments
						else
							local fulltoken = "["..tokenId..":"..value.."]"
							if delta_attach_after then
								tag.Comments = preraws..fulltoken..tag.Comments
							else
								tag.Comments = fulltoken..preraws..tag.Comments
							end
						end
						tag.CommentsOnly = true
						found_object = true
					end
				end)
			end
		end
	end)
	if not found_object then
		-- It may be a good idea to allow modders to suppress this warning for certain calls, I cant decide if its worth it or not...
		rubble.warning("A tag fitting the specifier: \""..tokenID.."\" Could not be found in object: \""..delta_object_cursor.."\". This may indicate an improper use of the delta operator.\n")
	end
	delta_nestedKey_cursor = ""
end
rubble.template("!DELTA", sodelta) -- See earlier comments.
rubble.template("@DELTA", sodelta) -- See earlier comments.
rubble.template("DELTA", sodelta)
rubble.template("#DELTA", sodelta)
rubble.template("@D", sodelta) -- See comment about previous abbreviation.