package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreratedOn  string  `json:"-"` //wont appear for JSON
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// we'll using this function to get all of the products
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Latte coffee",
		Price:       2.45,
		SKU:         "lat123",
		CreratedOn:  time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Espresso coffee",
		Price:       1.99,
		SKU:         "exp123",
		CreratedOn:  time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
