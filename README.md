Go Webserver
==============

Go static standalone webserver - I needed a simple webserver for a project and found that existing options 
required compilation from source or installation of dependant packages.  The Golang library provides the
excellent net/http modules which provides a rudimentary multi threaded webserver with little code.  As
golang compiles statically, the binaries can be used with no external dependancies.

For convenience, I've compiled the src into all available golang supported targets.

--------------

See the bin directory for precompiled binaries for -

- windows_386
- windows_amd64
- linux_386
- linux_amd64
- linux_arm
- openbsd_386
- openbsd_amd64
- freebsd_386
- freebsd_amd64
- freebsd_arm
- darwin ( MacOS )