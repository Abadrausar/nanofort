
Building Rubble from Source:
==============================================

To build Rubble you need to have Go (golang.org) installed.
To build the web or universal interface on windows you will need a C compiler, mingw works.

Once you have Go installed add <rubbledir>/other to your GOPATH or copy the contents of the src
directory to a location on your GOPATH.

Now all you need to do is fire up a command prompt and run:

	go build rubble7/interface/universal

More interfaces are available, look in `other/src/rubble7/interface/` for others.
