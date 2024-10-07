package routers

import (
	"github.com/Hellisham/last-api/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitRouters(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", handlers.GetProductHandler(db)).Methods("GET")
	router.HandleFunc("/product/{id}", handlers.GetProductbByIdHandler(db)).Methods("GET")
	router.HandleFunc("/product/create", handlers.CreateProductHandler(db)).Methods("POST")
	router.HandleFunc("/product/update/{id}", handlers.UpdateProductHandler(db)).Methods("PUT")
	router.HandleFunc("/product/delete/{id}", handlers.DeleteProductHandler(db)).Methods("DELETE")
	return router
}
