package main

import (
	"log"

	"dev.longnt1.git/aessment-hotel-booking.git/internal/config"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/controllers"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/repositories"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/routes"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/services"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/utils"
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

	roomRepo := repositories.NewRoomRepository(db)

	roomService := services.NewRoomService(roomRepo)

	roomController := controllers.NewRoomController(roomService)

	bookingRepo := repositories.NewBookingRepository(db)

	bookingService := services.NewBookingService(bookingRepo)

	bookingController := controllers.NewBookingController(bookingService)

	router := routes.SetupRouter(authController, roomController, bookingController)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
