package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Uint("port", 8080, "Specifies the port on which the server will run.")
	dir := flag.String("dir", "./public", "Specifies the directory that will be served")
	cert := flag.String("tls-cert", "", "Location of tls certificate (this with the key flag enables TLS)")
	key := flag.String("tls-key", "", "Location of the tls key (This with the tls-cert flag enables TLS)")
	tlsOnly := flag.Bool("tls-only", false, "Only runs a tls server")
	tlsport := flag.Uint("tls-port", 40443, "Specifies the TLS port when TLS is enabled")
	flag.Parse()

	// cert and key should both be provided or both empty
	if (*cert == "" && *key != "") || (*cert != "" && *key == "") {
		log.Fatal("Please provide both tls-cert and tls-key to enable TLS")
	}

	// check if tls is enabled
	tlsEnabled := *cert != "" && *key != ""
	// check if tls is enabled if tls-only is set
	if *tlsOnly && !tlsEnabled {
		log.Fatal("Can't set tls only if TLS is not enabled (set tls-cert and tls-key flags)")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", fs)

	// serve plain http
	if !*tlsOnly {
		go listenAndServe(cancel, fmt.Sprintf(":%d", *port), fs)
	}

	// server https
	if tlsEnabled {
		go listenAndServeTLS(cancel, fmt.Sprintf(":%d", *tlsport), *cert, *key, fs)
	}

	<-ctx.Done()
}

// listenAndServe serves a plain http webserver
func listenAndServe(cancel func(), addr string, handler http.Handler) {
	defer cancel()
	fmt.Printf("Now serving plain http on: localhost:%s\n", addr)
	log.Print(http.ListenAndServe(addr, handler))
}

// listenAndServeTLS serves a tls webserver
func listenAndServeTLS(cancel func(), addr string, cert string, key string, handler http.Handler) {
	defer cancel()
	fmt.Printf("Now serving tls on: localhost:%s\n", addr)
	log.Print(http.ListenAndServeTLS(addr, cert, key, handler))
}
