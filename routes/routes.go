package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/jwt-twitter-msg/controller"
)

func SetRoutes(userController *controller.UserController, postController *controller.PostController, router *gin.Engine) *gin.Engine {
	r := gin.Default()
	user := r.Group("users")
	{
		user.POST("/", userController.CreateUser)
		user.GET("/:id", userController.GetUser)
		user.GET("/", userController.GetAllUsers)
		user.PUT("/", userController.UpdateUser)
		user.DELETE("/:id", userController.DeleteUser)
	}

	post := r.Group("posts")
	{
		post.POST("/", postController.CreatePost)
		post.GET("/:id", postController.GetPost)
		post.GET("/", postController.GetAllPosts)
		post.PUT("/", postController.UpdatePost)
		post.DELETE("/:id", postController.DeletePost)
		post.GET("/users/:userID", postController.GetPostsByUserId)
	}
	return r
}
