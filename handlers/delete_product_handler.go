package handlers

import (
	"encoding/json"
	"errors"
	"github.com/Hellisham/last-api/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func DeleteProductHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		var product models.Products
		if result := db.First(&product, id); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			} else {
				http.Error(w, "Error retrieving Error", http.StatusInternalServerError)
				return
			}
		}
		if result := db.Delete(&product); result.Error != nil {
			log.Println("Error Deleting Product", result.Error)
			return
		}
		productResponse := ProductResponse{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(productResponse)
	}

}
