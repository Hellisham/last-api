package main

import (
	"github.com/Hellisham/last-api/db"
	"github.com/Hellisham/last-api/metrics"
	"github.com/Hellisham/last-api/models"
	"github.com/Hellisham/last-api/routers"
	"log"
	"net/http"
)

func main() {
	metrics.InitMetrics()
	datab := db.Connect()
	err := datab.AutoMigrate(&models.Category{}, &models.Products{}, &models.User{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	router := routers.InitRouters(datab)

	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
