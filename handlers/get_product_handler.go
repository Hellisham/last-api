package handlers

import (
	"encoding/json"
	"github.com/Hellisham/last-api/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type ProductResponse struct {
	Name        string
	Description string
	Price       float64
	Count       uint
	Category    string
}

func GetProductHandler(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var products []models.Products

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
				Category:    product.Category.Description,
			})
			json.NewEncoder(w).Encode(productsResponse)
		}
		w.Header().Set("Content-Type", "application/json")
	}
}

func GetProductbByIdHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Products
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid Products ID", http.StatusBadRequest)
			return
		}
		if res := db.Preload("Category").First(&product, id); res.Error != nil {
			log.Println("Product Not Found", res.Error)
		}
		productResponse := ProductResponse{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Count:       product.Count,
		}
		json.NewEncoder(w).Encode(productResponse)
		w.Header().Set("Content-Type", "application/json")
	}
}
