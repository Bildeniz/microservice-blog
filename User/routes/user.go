package routes

import (
	controller "User/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userController *controller.UserController) *gin.RouterGroup {
	user_routes := r.Group("/users")
	{
		// user_routes.POST("", userController.CreateUser) // RabbitMQ for user creation
		user_routes.GET("", userController.GetAllUser)
		user_routes.GET(":id", userController.GetUser)
	}

	return user_routes
}
