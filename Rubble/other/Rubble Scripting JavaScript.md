
Rubble JavaScript API Documentation:
========================================================================================================================

Rubble provides an API for using JavaScript scripts to extend it's functionality. This is the documentation for that API.

The language tag used for Java Script scripts is "LangJS", this tag will be auto applied to any file that has ".js" as a
last part extension.

Supposedly the runtime I use is compliant with the ECMA script specification aside from the following:

* `use strict` will parse, but does nothing.
* The regular expression engine (re2/regexp) is not fully compatible with the ECMA5 specification.

Regular expressions are translated into something that is compatible with the standard Go regular expression parser.
Unfortunately, backtracking is required for some patterns, and backtracking is not supported by the standard Go engine.

Therefore, the following syntax is incompatible:

	(?=)  // Lookahead (positive), currently a parsing error
	(?!)  // Lookahead (backhead), currently a parsing error
	\1    // Backreference (\1, \2, \3, ...), currently a parsing error

In addition to the above, re2 (Go) has a different definition for \s: [\t\n\f\r ]. The JavaScript definition, on the
other hand, also includes \v, Unicode "Separator, Space", etc. 

`rubble` Module
------------------------------------------------------------------------------------------------------------------------

JavaScript scripting can be summed up in one word: Don't (at least not yet).

There is no real API implemented as of yet, so wait until I get that far before you try to use JS for anything.
