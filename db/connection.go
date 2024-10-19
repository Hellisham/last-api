package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	var err error
	dns := "host=localhost port=5432 user=admin password=admin dbname=postgres sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Cant Connect To DataBase", err)
	}
	return DB
}
