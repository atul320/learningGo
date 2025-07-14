package main

import (
	"log"
	"os"
	"time"

	"url-shortener/controllers"
	"url-shortener/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	err := utils.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize MongoDB connection
	utils.InitDB()
	defer utils.CloseDB()

	// Create a background job to clean expired URLs
	go func() {
		for {
			time.Sleep(24 * time.Hour) // Run once per day
			controllers.CleanExpiredURLs()
		}
	}()

	// Set up router
	router := gin.Default()

	// Auth routes
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// URL routes
	url := router.Group("/url")
	url.Use(utils.AuthMiddleware())
	{
		url.POST("/create", controllers.CreateURL)
		url.GET("/my-urls", controllers.GetUserURLs)
		url.GET("/:id", controllers.GetURL)
		url.DELETE("/:id", controllers.DeleteURL)
	}

	// Public route for redirecting
	router.GET("/:shortCode", controllers.RedirectURL)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(router.Run(":" + port))
}
