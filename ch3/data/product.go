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
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"` // - = ommit from output
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// json.NewEncoder is significantly faster than using the marshall method
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProduct() Products {
	return productList
}

var productList = Products{
	{
		ID:          1,
		Name:        "latte",
		Description: "Frothy milky coffee, very nice!",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Esspresso",
		Description: "Dark and flavorful",
		Price:       2.49,
		SKU:         "abc125",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
