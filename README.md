servedir
========

A Go Language program serve the current directory over HTTP

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

The extremely basic file http.FileServer and http.Dir does not sort the directory listing and there is no file size information displayed like listings.
