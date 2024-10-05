package handlers

import (
	"last-api/models"
	"net/http"
)

type ProductResponse struct {
	Name        string
	Description string
	Price       float64
	Count       uint
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Products []models.Product

}
