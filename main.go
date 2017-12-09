package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Uint("port", 8080, "specifies the port on which the server will run.")
	dir := flag.String("dir", "./", "specifies the directory that will be served")
	flag.Parse()

	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", fs)

	fmt.Printf("Now serving on: localhost:%d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
