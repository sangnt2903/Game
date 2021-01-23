package main

import (
	"MiniGameAPI/Database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	r := gin.Default()

	Database.ConnectDatabase()
	Database.InitializeDI(r)

	r.Run("localhost:8091")
}
