package handlers

import (
	"encoding/json"
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

	// encoding so to use data in JSON format
	productListJSON, err := json.Marshal(productList)
	if err != nil {
		http.Error(rw, "Unable to marshall productList object into its JSON ", http.StatusInternalServerError)
	}

	rw.Write(productListJSON)
}
