package handlers

import (
	"ch3/data"
	"encoding/json"
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
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "Unable to marshall data", http.StatusInternalServerError)
	}
	w.Write(d)
}
