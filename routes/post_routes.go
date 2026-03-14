package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/jwt-twitter-msg/controller"
)

func PostRoutes(postController controller.PostController) *gin.Engine {
	r := gin.Default()
	post := r.Group("posts")
	{
		post.POST("/", postController.CreatePost)
		post.GET("/:id", postController.GetPost)
		post.GET("/", postController.GetAllPosts)
		post.PUT("/", postController.UpdatePost)
		post.DELETE("/:id", postController.DeletePost)
		post.GET("/:userID", postController.GetPostsByUserId)
	}
	return r
}
