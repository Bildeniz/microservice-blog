package routes

import (
	"User/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ctrl *controller.UserController) *gin.Engine {
	r := gin.Default()

	SetupUserRoutes(r, ctrl)

	return r
}
