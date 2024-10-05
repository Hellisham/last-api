package handlers

import (
	"encoding/json"
	"github.com/Hellisham/last-api/models"
	"gorm.io/gorm"
	"net/http"
)

type ProductResponse struct {
	Name        string
	Description string
	Price       float64
	Count       uint
}

func GetProductHandler(db *gorm.DB) http.HandlerFunc {
	var products []models.Products
	return func(w http.ResponseWriter, r *http.Request) {
		if res := db.Preload("Category").Find(&products); res.Error != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
			return
		}
		for _, product := range products {
			var productsResponse []ProductResponse
			productsResponse = append(productsResponse, ProductResponse{
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				Count:       product.Count,
			})
			json.NewEncoder(w).Encode(productsResponse)
		}
		w.Header().Set("Content-Type", "application/json")
	}
}
