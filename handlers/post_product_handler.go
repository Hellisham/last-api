package handlers

import (
	"encoding/json"
	"github.com/Hellisham/last-api/db"
	"github.com/Hellisham/last-api/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type PostProductHandler struct {
	Name        string
	Description string
	Price       float64
	Count       float64
	Category    string
}

func CreateProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Products

		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		if result := db.DB.Preload("Category").First(&product, product.ID); result.Error != nil {
			log.Println("Error preloading category", result.Error)
			http.Error(w, "Error retrieving product", http.StatusInternalServerError)
			return
		}

		result := db.DB.Create(&product)
		if result.Error != nil {
			log.Println("Error creating product", result.Error)
		}

		productsResponse := ProductResponse{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Count:       product.Count,
			Category:    product.Category.Name,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(productsResponse)
	}
}

func UpdateProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Products
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		if result := db.DB.First(&product, id); result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			} else {
				http.Error(w, "Error retrieving book", http.StatusInternalServerError)
				return
			}
		}
		var updateproduct models.Products
		if err := json.NewDecoder(r.Body).Decode(&updateproduct); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		product.Name = updateproduct.Name
		product.Description = updateproduct.Description
		product.Price = updateproduct.Price
		product.Count = updateproduct.Count
		product.CategoryID = updateproduct.CategoryID
		db.DB.Save(&product)
		if result := db.DB.Preload("Category").First(&product, product.ID); result.Error != nil {
			http.Error(w, "error preload category", http.StatusInternalServerError)
		}
		productsResponse := ProductResponse{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Count:       product.Count,
			Category:    product.Category.Name,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(productsResponse)
	}

}
