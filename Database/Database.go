package Database

import (
	"MiniGameAPI/Logging/Error"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	database := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		username,
		password,
		host,
		port,
		database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if Error.ErrorService.HasError(err) {
		panic(err)
	}
	// Migrate
	if os.Getenv("DEBUGGING") == "true" {
		DB = db.Debug()
	} else {
		DB = db
	}

	if os.Getenv("MIGRATE") == "true" {
		AutoMigrate()
	}
}

func GetDatabase() *gorm.DB {
	return DB
}
