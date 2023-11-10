package main

import (
	"SQ-Assessment/config"
	"SQ-Assessment/route"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := config.InitDB()
	r := gin.Default()
	r.Use(cors.Default())
	// Per route middleware, you can add as many as you desire.
	// middleware.Logmiddleware(e)

	route.NewRoute(r, db)
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
