package routers

import (
	"github.com/Hellisham/last-api/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitRouters(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", handlers.GetProductHandler(db)).Methods("GET")
	return router
}
