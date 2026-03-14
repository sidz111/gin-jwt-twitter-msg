package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/jwt-twitter-msg/controller"
)

func UserRoutes(userController controller.UserController) *gin.Engine {
	r := gin.Default()
	user := r.Group("users")
	{
		user.POST("/", userController.CreateUser)
		user.GET("/:id", userController.GetUser)
		user.GET("/", userController.GetAllUsers)
		user.PUT("/", userController.UpdateUser)
		user.DELETE("/:id", userController.DeleteUser)
	}
	return r
}
