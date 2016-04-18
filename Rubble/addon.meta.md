
# addon.meta:
# ----------------------------------------------

# Addons may include a special file named addon.meta that contains extra information about the
# addon. This file is optional but **highly** recommended. So many features depend on the addon.meta
# file that an addon that lacks one will be half crippled.

# Some fields are inheritable by any addon "below" the current addon in the addon tree. For this reason you may place
# addon.meta files in directories that are not addons (so that their values may be inherited by their children).

# Example addon.meta file (with lots of helpful comments):

	{
	# Example addon.meta file (this file is JSON encoded)
	
	# Yes, my JSON files really use `#` as the comment char. '//' and '/* */' are not supported.
	# Just be glad I hacked comment support in at all, the parser does not support them by default.
	# Comments are only supported on a line of their own.
	
	# This example file includes every possible key. Most addons will only need a subset of these keys.
	
	# Addon tags, the ones listed here are those with hardcoded handling (but you can specify anything you want).
	# The hardcoded tags get special handling various places, for example on the web UI addon information page they will
	# be followed by a short description of their meaning.
	"Tags": {
		# Addon is an automatically managed library.
		"Library": false,
		
		# The addon is not actually an addon, but a documentation pack.
		"DocPack": false,
		
		# Automatically set by the addon loader if the addon has tileset information.
		# You should not set this manually.
		"TileSet": false,
		
		# Automatically set by the addon loader if the addon has test files.
		# You should not set this manually.
		"HasTests": false,
		
		# The addon is safe to apply to existing regions.
		"SaveSafe": false,
		
		# Does this addon require DFHack?
		# If you do not set this explicitly then Rubble will guess the value based on if the addon contains any
		# files tagged "DFHack", but if this is explicitly set your value will not be overridden no mater what.
		# This is used by the web UI to allow the user to hide addons that require DFHack.
		"DFHack": false,
		
		# Set for addons that are of interest mostly to developers, this is used by the web UI addon filtering system
		# so it can hide these addons by default.
		"Dev": false,
		
		# Addons with this tag are not standard addons. Currently causes the addon to not be listed as an option
		# during standard generation cycles when using the web UI. Use if you have a special addon that should only
		# be used with independent apply mode or such like.
		"NotNormal": false
	},
	
	# The addon priority. Lower values load first. If negative it will try to use the same value as the parent addon, if
	# there is no parent the the absolute value will be used. By convention `25` is used for standard libraries, `50` for
	# base addons and `75` for other default addons. Most addons should use the default of `-100`
	"LoadPriority": -100,
	
	# Addon name, optional. If the first character is "$" and the name is longer than two characters then the given name
	# is prefixed with the name of the parent addon. The default name is `<parent name>/<directory name>`.
	# 
	# If you set this to "-" it will cause the addon to have no name. This is a useful technique for organizing addon
	# trees without making the addon names too long. Use this only in addon.meta files that are not in valid addons!
	#"Name" = "Example/Addon",
	"Name": "$/Addon",
	
	# The version and author as strings for display in the web UI.
	# If not provided these will default to the same as the parent addon.
	"Version": "1.0",
	"Author": "Your Name",
	
	# One line description of the addon.
	# This must always be markdown (embedded HTML is OK).
	"Header": "This is an example addon.meta file.",
	
	# Optional longer description loaded from a separate file.
	# The given description file must be contained in the same addon as this meta file.
	# If the file extension is ".htm" or ".html" the file is assumed to be HTML, if it is ".md" or ".text" it is assumed to
	# be markdown. Any other extension is assumed to be preformated plain text.
	"DescFile": "example.md",
	"Description": "If you prefer you can have the description inline.\nObviously this can be convenient, but only for short descriptions.\nIf DescFile is set this is ignored (or rather, overwritten).\nInline descriptions are always assumed to be markdown.",
	
	# If you wish you can specify configuration variables here.
	# These variables will automatically be set to their defaults when addons are activated, but users can
	# change their values in various ways.
	"Vars": {
		# The field name is the name of the variable to set.
		"TEST": {
			# Name is the name to present to the user, should be short and clear.
			# Setting Name to "-" will cause the variable to be hidden in most interfaces.
			"Name": "A test variable",
			
			# Values is an array of possible values. If more than one item is present then the variable will
			# be restricted to just those values. If only one value is provided then you may set the variable
			# to any string. In any case the first value will be the default. If no values are provided then
			# the variable will default to the empty string.
			# Note that Rubble itself does not enforce any value restrictions you set here, so you should always
			# validate any variables (with a script or something) before you use them.
			"Values": ["YES", "NO"]
		}
	},
	
	Activates: [
		"This/Is An/Array",
		"Of All/The",
		"Addons/This One/Depends On",
		"Anything/Listed/Here",
		"Will Be/Activated/When",
		"This/Addon Is"
	],
	
	Incompatible: [
		"This/Is An/Array",
		"Of All/The",
		"Addons/This One/Is",
		"Known/To Be",
		"Incompatible/With",
		"Anything/Listed/Here",
		"Will Cause/Rubble To/Abort",
		"If It Is/Active",
		"When This/Addon Is"
	]
	}
