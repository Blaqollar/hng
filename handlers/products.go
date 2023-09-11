package handlers

import (
	"log"
	"net/http"

	"github.com/Blaqollar/hng/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	r.URL.Query().Get("slack_name")
	r.URL.Query().Get("track")
	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}
