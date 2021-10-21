package handlers

import (
	"net/http"

	"github.com/lucasbarroso23/building-microservices-with-go/data"
)

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}

	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
