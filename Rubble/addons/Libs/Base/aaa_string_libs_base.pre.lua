
-- String handling

rubble.template("@STR_LOWER", [[
	return string.lower(rubble.targs({...}, {""}))
]])

rubble.template("@STR_UPPER", [[
	return string.upper(rubble.targs({...}, {""}))
]])

rubble.template("@STR_TITLE", [[
	return string.title(rubble.targs({...}, {""}))
]])

rubble.template("@STR_REPLACE", [[
	local str, old, new, n = rubble.targs({...}, {"", "", "", -1})
	return string.replace(str, old, new, n)
]])

rubble.template("@STR_TO_ID", [[
	return string.upper(string.replace(rubble.targs({...}, {""}), " ", "_", -1))
]])
