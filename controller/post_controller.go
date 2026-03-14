package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sidz111/jwt-twitter-msg/models"
	"github.com/sidz111/jwt-twitter-msg/service"
)

type PostController struct {
	serv service.PostService
}

func NewPostController(serv service.PostService) *PostController {
	return &PostController{serv: serv}
}

func (c *PostController) CreatePost(ctx *gin.Context) {
	var post models.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.serv.CreatePost(ctx.Request.Context(), &post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}
func (c *PostController) GetPost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	post, err := c.serv.GetPost(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}
func (c *PostController) GetAllPosts(ctx *gin.Context) {
	posts, err := c.serv.GetAllPosts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
func (c *PostController) UpdatePost(ctx *gin.Context) {
	var post *models.Post
	if err := c.serv.UpdatePost(ctx, post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "post updated successfully",
	})
}
func (c *PostController) DeletePost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.serv.DeletePost(ctx, uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "post delete successfully",
	})
}
func (c *PostController) GetPostsByUserId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	posts, err := c.serv.GetPostsByUserId(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
