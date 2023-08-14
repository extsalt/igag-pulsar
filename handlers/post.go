package handlers

import (
	"github.com/gin-gonic/gin"
	"pulsar/config"
	"pulsar/handlers/requests"
	"pulsar/models"
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
		UserID: 1,
		Title:  postRequest.Title,
		Body:   postRequest.Body,
	}
	config.PulsarConfig.DB.Create(post)
	c.JSON(201, gin.H{
		"message": "Post created",
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.PulsarConfig.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
