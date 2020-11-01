package handlers

import (
	"log"
	"net/http"

	"github.com/vaibhav/CoffeeOnDemand/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// getting data from package - data
	productList := data.GetProducts()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall productList object into its JSON ", http.StatusInternalServerError)
	}
}
