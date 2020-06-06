package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/clemenspeters/golang-tutorials/introduction-to-miroservices-series/episode3/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to Marshal json", http.StatusInternalServerError)
	}

	rw.Write(d)
}
