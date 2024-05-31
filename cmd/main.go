package main

import (
	"log"
	"main.go/internal/config"
	"main.go/internal/controllers"
	"main.go/internal/repositories"
	"main.go/internal/routes"
	"main.go/internal/services"
	"main.go/internal/utils"
)

func main() {
	db := config.GetDB()
	hasher := &utils.Hasher{}

	userRepo := repositories.NewUserRepository(db, hasher)

	if err := userRepo.SeedData(); err != nil {
		log.Fatalf("failed to seed initial data: %v", err)
	}

	userService := services.NewUserService(userRepo)

	authController := controllers.NewAuthController(userService)

	router := routes.SetupRouter(authController)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
