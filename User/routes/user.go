package routes

import (
	"User/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userController *controller.UserController) *gin.RouterGroup {
	user_routes := r.Group("/users")
	{
		user_routes.POST("", userController.CreateUser)
		user_routes.GET("", userController.GetAllUser)
	}

	return user_routes
}
