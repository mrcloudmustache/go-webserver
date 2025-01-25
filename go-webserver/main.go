package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := flag.Int("p", 8082, "Listening port of the web server.")
	flag.Parse()
	p := ":" + strconv.Itoa(*port)

	log.Printf("INFO: Listening on port %v", *port)
	log.Fatal(http.ListenAndServe(p, nil))

}
