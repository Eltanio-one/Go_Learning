package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// this is the handler we are creating
type Hello struct {
	l *log.Logger
}

// new hello handler
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// serveHTTP handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// we replace log.Println here with our new hello handler,
	// which is a log.Logger and has a println method we can call also.
	h.l.Println("Hello, Bitch")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Sugma, %s", d)
}
