package api

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	product := Product{
		Id:    1,
		Name:  "ลำโพง",
		Price: 100.00,
	}
	productJson, _ := json.Marshal(product)
	w.Write(productJson)
}
