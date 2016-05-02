-------------------------------------------------------------------------------------------------------------------------------------------
--                                 Semantic Delta operators
--      {@CONTEXT;objectid;wheretoappend}
--   Fixes the objectid (for example ENTITY:MOUNTAIN) where the delta will be applied and also if the appended preraws go before or after (this is the default) the attachment token.
--      {@DELTA;token;value;preraws}
--   Modifies the end value of a token inside an object and  eventually attach before or after it the preraws, if the token id is given but not the value the token is killed.
--   if not token is given the preraws are appended at the end of the object. When no object is found a detailed DELTA warning is generated.
-------------------------------------------------------------------------------------------------------------------------------------------
___rubble_object_cursor = ""
___rubble_attach_after = true
local socontext = [[
   local value, after = rubble.targs({...}, {"", "AFTER"})
   ___rubble_object_cursor = rubble.expandvars(value, "$", true)
   if after == "BEFORE" then
      ___rubble_attach_after = false
   else
      ___rubble_attach_after = true
   end
]]
rubble.template("!CONTEXT", socontext)
rubble.template("@CONTEXT", socontext)
rubble.template("CONTEXT", socontext)
rubble.template("#CONTEXT", socontext)
rubble.template("#C", socontext)
local sodelta = [[
   local token, value, preraws = rubble.targs({...}, {"", "", ""})
   if token == "" then
      local data = rubble.registry["Libs/Base:SHARED_OBJECT_ADD"].table
      data[___rubble_object_cursor] = (data[___rubble_object_cursor] or "").."\n\t"..rubble.parse(preraws)
      return
   end
   token = string.split(token, ":")
   rubble.libs_base.sharedobject_walk(___rubble_object_cursor, function(tag)
      if rubble.libs_base.matchtag(tag, token) then
         if value == "" then
            local preraws = "-"..tag.ID
            for _, v in ipairs(tag.Params) do
               preraws = preraws..":"..v
            end
            tag.Comments = preraws.."-"..tag.Comments
         else
            local fulltoken = "["..token..":"..value.."]"
            if ___rubble_attach_after == true then
               tag.Comments = preraws..fulltoken..tag.Comments
            else
               tag.Comments = fulltoken..preraws..tag.Comments
            end
         end
         tag.CommentsOnly = true
      else
      rubble.warning("OBJECT: "..___rubble_object_cursor.." NOT found, {@DELTA;"..token..";"..value.."} mal formed or incorrect order of application of mods")
      end
   end)
]]
rubble.template("!DELTA", sodelta)
rubble.template("@DELTA", sodelta)
rubble.template("DELTA", sodelta)
rubble.template("#DELTA", sodelta)
rubble.template("@D", sodelta)