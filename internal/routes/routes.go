package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/controllers"
	"main.go/internal/middlewares"
)

func SetupRouter(
	authController *controllers.UserController,
	// bookingController *controllers.BookingController,
	// searchController *controllers.SearchController,
) *gin.Engine {
	router := gin.Default()

	// Public routes
	public := router.Group("/api/users")
	{
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)
	}

	// Private routes
	private := router.Group("/api")
	private.Use(middlewares.AuthMiddleware())
	{
		private.PUT("/profile", authController.EditProfile)
		private.GET("/profile/:id", authController.GetProfile)
		private.GET("/users", authController.GetAllUsers)
		private.DELETE("/user/:id", authController.DeleteUser)

		//// Booking routes
		//private.POST("/bookings", bookingController.CreateBooking)
		//private.GET("/bookings", bookingController.GetBookings)
		//private.GET("/bookings/:id", bookingController.GetBooking)
		//private.PUT("/bookings/:id", bookingController.UpdateBooking)
		//private.DELETE("/bookings/:id", bookingController.CancelBooking)
		//
		//// Search routes
		//private.GET("/search/hotels", searchController.SearchHotels)
		//private.GET("/search/rooms", searchController.SearchRooms)
	}

	return router
}
