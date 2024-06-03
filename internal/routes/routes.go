package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/controllers"
	"main.go/internal/middlewares"
)

func SetupRouter(
	authController *controllers.UserController,
	roomController *controllers.RoomController,
	bookingController *controllers.BookingController,
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
		// User routes
		private.PUT("/users/profile", authController.EditProfile)
		private.GET("/users/profile/:id", authController.GetProfile)
		private.GET("/users", authController.GetAllUsers)
		private.DELETE("/users/:id", authController.DeleteUser)

		// Room routes
		private.GET("/rooms/:id", roomController.GetRoomByID)
		private.POST("/rooms", roomController.CreateRoom)
		private.PUT("/rooms/:id", roomController.UpdateRoom)
		private.DELETE("/rooms/:id", roomController.DeleteRoom)
		private.GET("/rooms", roomController.GetRoomsByFilter)

		// Booking routes
		private.POST("/bookings", bookingController.CreateBooking)
		private.PUT("/bookings/:id", bookingController.UpdateBooking)
		private.DELETE("/bookings/:id", bookingController.DeleteBooking)
		private.GET("/bookings/:id", bookingController.GetBookingByID)
		private.GET("/bookings", bookingController.GetAllBookings)
		private.GET("/bookings/user/:user_id", bookingController.GetBookingsByUserID)
	}

	return router
}
