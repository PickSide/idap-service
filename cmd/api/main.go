package main

import (
	"auth-service/db"
	"auth-service/handlers"
	"auth-service/middlewares"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	g := gin.Default()

	if err := db.OpenConnectionToDB(); err != nil {
		log.Fatal(err)
	}

	g.Use(cors.Default())

	v1 := g.Group("/api/v1", middlewares.FromValidDomain())

	v1.GET("/api/v1/verify", handlers.Verify)

}
