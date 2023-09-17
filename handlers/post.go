package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"net/http"
	"pulsar/config"
	"pulsar/handlers/requests"
	"pulsar/handlers/resources"
	"pulsar/models"
	"pulsar/pkg/sanctum"
)

func AddPost(c *gin.Context) {
	var postRequest requests.CreatePostRequest
	if c.BindJSON(&postRequest) != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}
	post := &models.Post{
		UserID:        sanctum.AuthUser.ID,
		Title:         postRequest.Title,
		Body:          postRequest.Body,
		OriginalImage: postRequest.ImageUrl,
		Slug:          slug.Make(postRequest.Title),
	}
	tx := config.PulsarConfig.DB.Create(post)
	if tx.RowsAffected != 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error" + tx.Error.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "Post created",
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	err := config.PulsarConfig.DB.Preload("Users").Find(&posts).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error",
		})
		return
	}
	c.JSON(200, resources.PostsJsonResource(&posts))
}
