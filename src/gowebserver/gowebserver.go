package main

import (
	"flag"
    "fmt"
	"net/http"
)

var port = flag.String("port", "8080", "TCP port to listen on")
var root = flag.String("root", ".", "Filesystem root path")

func main() {
	flag.Parse()
    fmt.Println("James Spurin's mini go webserver - listening on port:", *port, "serving files from location:", *root)
	panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*root))))
}
