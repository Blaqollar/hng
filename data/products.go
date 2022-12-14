package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

type Products Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

var productList = Products{
	SlackUsername: "uchedingba",
	Backend:       true,
	Age:           27,
	Bio:           "I am a graduate of computer Science interested in Software Engineering(Backend)",
}
