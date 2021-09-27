package handlers

import (
	"building-microservices-with-go/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()

	// this is the tradicional way to marshal a json
	// d, err := json.Marshal(lp)

	// using enconde you do not alocate memory to marshal the json, goes directly to the io.Writer
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}

}
