package handlers

import (
	"github.com/gin-gonic/gin"
	"pulsar/config"
	"pulsar/handlers/requests"
	"pulsar/models"
)

type CommentQueryBind struct {
	ID string `uri:"postID" binding:"required"`
}

// AddComment Create a new comment on post
func AddComment(c *gin.Context) {
	var commentQueryBind CommentQueryBind
	if err := c.ShouldBindUri(&commentQueryBind); err != nil {
		c.JSON(404, gin.H{"message": err})
		return
	}
	var createCommentRequest requests.CreateCommentRequest
	if err := c.ShouldBindJSON(&createCommentRequest); err != nil {
		c.JSON(422, gin.H{"message": err})
		return
	}
	comment := &models.Comment{
		PostID: 1,
		Body:   createCommentRequest.Body,
		UserID: 1,
	}
	err := config.PulsarConfig.DB.Create(comment).Error
	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	c.JSON(201, gin.H{})
}
