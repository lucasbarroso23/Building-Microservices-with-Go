package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("hello world")

	// read the body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {

		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}

	log.Printf("Data %s\n", d)

	// write the response
	fmt.Fprintf(rw, "Hello again %s", d)
}
