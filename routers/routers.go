package routers

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) {
	router := mux.NewRouter()
	router.HandleFunc("products/")

}
