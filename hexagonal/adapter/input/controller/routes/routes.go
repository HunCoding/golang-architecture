package routes

import (
	"github.com/gin-gonic/gin"
	"hexagonal/adapter/input/controller"
	"hexagonal/adapter/input/controller/middlewares"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", middlewares.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", middlewares.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
}
