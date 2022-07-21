package db

import (
	"aims/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost port=5432 user=root password=adrian dbname=auth  sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = connection

	err = connection.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}
}
