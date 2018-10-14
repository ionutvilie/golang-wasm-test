package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	var addr = flag.String("addr", ":8080", "The addr of the application.")

	mux := http.NewServeMux()

	// root handler will redirect user to www folder and act as an http file server
	mux.Handle("/", http.FileServer(http.Dir("./build")))

	// start the http server
	log.Println("Starting application on", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
