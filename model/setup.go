package model

import (
	store "example/backend/v1/database"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "user=sameermishra password=sameer dbname=mydatabase port=5430 sslmode=disable"
	database, err1 := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	store.CheckErr(err1)
	err2 := database.AutoMigrate(&Movie{}, &Show{}, &Theatre{}, &Ticket{})
	store.CheckErr(err2)
	DB = database
}
