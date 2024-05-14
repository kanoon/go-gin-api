package config

import (
	"demo-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=127.0.0.1 user=postgres password=test1234 dbname=postgres port=5444 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic(err)
    }

    db.AutoMigrate(&models.User{})
    DB = db
}
