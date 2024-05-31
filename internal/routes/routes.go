package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/controllers"
	"main.go/internal/middlewares"
)

func SetupRouter(
	authController *controllers.UserController,
	roomController *controllers.RoomController,
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

		// Room routes
		private.GET("/room/:id", roomController.GetRoomByID)
		private.POST("/room", roomController.CreateRoom)
		private.PUT("/room", roomController.UpdateRoom)
		private.DELETE("/room", roomController.DeleteRoom)
		private.GET("/rooms", roomController.GetRoomsByFilter)
	}

	return router
}
