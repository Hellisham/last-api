package routers

import (
	"github.com/Hellisham/last-api/handlers"
	"github.com/Hellisham/last-api/middleware"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
)

func InitRouters(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/products", handlers.GetProductHandler()).Methods("GET")
	router.HandleFunc("/product/{id}", middleware.JWTAuthMiddleware(handlers.GetProductbByIdHandler())).Methods("GET")
	router.HandleFunc("/product/create", handlers.CreateProductHandler()).Methods("POST")
	router.HandleFunc("/product/update/{id}", handlers.UpdateProductHandler()).Methods("PUT")
	router.HandleFunc("/product/delete/{id}", handlers.DeleteProductHandler()).Methods("DELETE")
	router.HandleFunc("/user/create", handlers.RegisterHandler()).Methods("POST")
	router.HandleFunc("/user/login", handlers.LoginHandler()).Methods("POST")

	return router
}
