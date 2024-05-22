package handlers

import (
	"ch3/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	// Uses custom encoder in the data
	// better composition of marshalling step of the data
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshall data", http.StatusInternalServerError)
	}
	// w.Write(d)
}
