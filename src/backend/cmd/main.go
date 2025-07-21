package main

import (
	"log"
	"os"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	
	"vue-gin-backend/internal/handler"
	"vue-gin-backend/pkg/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}
	
	database.Connect()
	database.Migrate()
	database.SeedData()
	
	r := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))
	
	api := r.Group("/")
	{
		api.GET("/users", handler.GetUsers)
		api.GET("/microposts", handler.GetMicroposts)
		api.GET("/microposts/:id", handler.GetMicropostByID)
		api.POST("/microposts", handler.CreateMicropost)
	}
	
	r.Static("/templates/admin", "./templates/admin")
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	
	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}
