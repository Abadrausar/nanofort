
-- Placeholder template generator.
function rubble.placeholder(name, silent)
	if silent then
		rubble.template(name, [[return ""]])
	else
		rubble.template(name, [[return  "The addon that contains \"]]..name..[[\" is disabled."]])
	end
end

-- Emulation of the way the old (Rubble 6) raw parser worked.
function rubble.rparse.walk(raws, action)
	local tags = rubble.rparse.parse(raws)
	for _, tag in ipairs(tags) do
		action(tag)
	end
	return rubble.rparse.format(tags)
end

-- Default handling of template arguments.
function rubble.targs(args, defaults, noexpand)
	if type(noexpand) == "table" then
		for i = 1, #args, 1 do
			if not noexpand[i] then
				args[i] = rubble.expandvars(args[i])
			end
		end
	else
		if not noexpand then
			for i = 1, #args, 1 do
				args[i] = rubble.expandvars(args[i])
			end
		end
	end
	
	if defaults == nil then
		return args
	end
	
	if #args >= #defaults then
		return table.unpack(args)
	else
		for i = #args + 1, #defaults, 1 do
			table.insert(args, defaults[i])
		end
		return table.unpack(args)
	end
end

local bool_map = {
	["yes"] = true,
	["y"] = true,
	["true"] = true,
	["t"] = true,
	["-1"] = true,
	["1"] = true,
}

-- Convert a string to a boolean value of some kind. Set t and f to specify what should be returned in true and false cases.
function rubble.tobool(opt, t, f)
	if t == nil then
		t = true
	end
	if f == nil then
		f = false
	end
	
	if bool_map[string.lower(opt)] then
		return t
	end
	return f
end
