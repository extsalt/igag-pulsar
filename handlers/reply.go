package handlers

import (
	"github.com/gin-gonic/gin"
	"pulsar/config"
	"pulsar/handlers/requests"
	"pulsar/models"
)

type ReplyQueryBind struct {
	ID string `uri:"commentID" binding:"required"`
}

// AddReply Create a new comment on post
func AddReply(c *gin.Context) {
	var replyQueryBind ReplyQueryBind
	if err := c.ShouldBindUri(&replyQueryBind); err != nil {
		c.JSON(404, gin.H{"message": err})
		return
	}
	var createReplyRequest requests.CreateReplyRequest
	if err := c.ShouldBindJSON(&createReplyRequest); err != nil {
		c.JSON(422, gin.H{"message": err})
		return
	}
	reply := &models.Reply{
		CommentID: 1,
		Body:      createReplyRequest.Body,
		UserID:    1,
	}
	err := config.PulsarConfig.DB.Create(reply).Error
	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	c.JSON(201, gin.H{})
}
