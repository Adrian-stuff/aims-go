package db

import (
	"aims/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost port=1234 user=root password=password dbname=authServer  sslmode=disable"
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
