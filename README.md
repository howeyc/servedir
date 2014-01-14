servedir
========

A Go Language program that serves the current directory over HTTP

Installation
-----------

	go get github.com/howeyc/servedir
	
Running
-----------

Run the executable and it will serve the current working directory on 
port 8080.

Command line flags are available to change the port and to listen on localhost only.

Why
-----------

The extremely basic file http.FileServer and http.Dir do not sort the directory listing and there is no file size information displayed in the listing.
