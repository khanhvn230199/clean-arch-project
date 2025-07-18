package main

import (
	"clean-arch-project/api/handler"
	"clean-arch-project/api/route"
	"clean-arch-project/config"
	"clean-arch-project/infrastructure/database"
	"clean-arch-project/infrastructure/repository"
	"clean-arch-project/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database connection
	db, err := database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize use cases
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUseCase)

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	route.SetupRoutes(router, userHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
