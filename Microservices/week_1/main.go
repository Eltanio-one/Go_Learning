package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// HandleFunc is a convenience method on the http package, and it registers a function to a path on the default servemux
	// everything related to the server in Go is an HTTP handler
	// if we dont specify a handler for the parameter in ListenAndServe, it uses default servemux
	// but using HandleFunc creates a handler and adds it to the default servemux
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello, Bitch")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Sugma, %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye, Bitch")
	})

	http.ListenAndServe(":9090", nil)

}
