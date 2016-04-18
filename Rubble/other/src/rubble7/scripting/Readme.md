
This directory contain "script runners" for various scripting languages I have tried at one time or another.

Rubble interfaces may use these runners to provide scripting languages to the Rubble engine. Each of these runners
contains everything needed to use the associated language for Rubble scripts.

* * *

A Rubble script runner wraps everything needed to run a script together with the API required to interact with the
internals of the Rubble engine. Before a runner can be considered fully functional it needs to provide (at a *minimum*):

* Read/write access to the AXIS VFS file system of the current Rubble State (`(*rubble7.State).FS`).
* Read/write access to the following fields in the current Rubble State:
	* `(*rubble7.State).Files`
	* `(*rubble7.State).GlobalFiles`
	* `(*rubble7.State).Addons`
	* `(*rubble7.State).AddonsTbl`
	* `(*rubble7.State).ScrRegistry`
	* `(*rubble7.State).VariableData`
	* `(*rubble7.State).CurrentFile`
	* `(*rubble7.State).Init`
	* `(*rubble7.State).D_Init`
* Within the above fields references to items of type `*addon.Meta` should have the following fields *read-only*, all
  others should be read/write:
	* `(*addon.Meta).Name`
	* `(*addon.Meta).DescFile`
	* `(*addon.Meta).Activates`
	* `(*addon.Meta).Incompatible`
	* `(*addon.Meta).LoadPriority`
* Functions to print to the log.
* Functions to print warnings and raise aborts.
* Access to the various Rubble version constants.
* Functions to define templates, both script templates (in this language and others) and those defined in RTL.

Access to the other Rubble APIs (the raw parser, the match/merge engine, etc) is optional, but recommended.

* * *

Active/Maintained Runners:

* "dctech_lua": Fully functional Lua 5.3 API for the `dctech/lua` VM.
* "otto": Minimally functional JavaScript API for Otto (`github.com/robertkrimen/otto`).

Other Runners:

* "go-lua": Broken, Lua 5.2 API that uses the (buggy) `github.com/Shopify/go-lua` VM.
* "gopher-lua": Fully functional last time I used it, Lua 5.1 API that uses `github.com/yuin/gopher-lua`.
