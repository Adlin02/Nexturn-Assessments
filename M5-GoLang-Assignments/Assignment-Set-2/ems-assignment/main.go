package main

import (
	"ecommerce-inventory/config"
	"ecommerce-inventory/controller"
	"ecommerce-inventory/middleware"
	"ecommerce-inventory/repository"
	"ecommerce-inventory/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db, err := config.InitializeDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Set up repositories, services, and controllers
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Set up router
	router := gin.Default()

	// Middleware for logging requests
	router.Use(middleware.LoggingMiddleware())

	// User routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// Product routes (authentication required)
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware()) // Middleware for authentication
	{
		// Routes for managing products
		authorized.POST("/product", middleware.ValidationMiddleware(), productController.AddProduct)
		authorized.GET("/product/:id", productController.GetProduct)
		authorized.PUT("/product/:id", productController.UpdateProduct)
		authorized.DELETE("/product/:id", productController.DeleteProduct)
		authorized.GET("/products", productController.GetAllProducts)
	}

	// Start the server on port 8080
	router.Run(":8080")
}
