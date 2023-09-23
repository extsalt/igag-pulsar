package handlers

import (
	"github.com/gin-gonic/gin"
	"pulsar/config"
	"pulsar/handlers/requests"
	"pulsar/models"
	"pulsar/pkg/sanctum"
	"strconv"
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
	postID, err := strconv.ParseUint(commentQueryBind.ID, 10, 64)
	if err != nil {
		c.JSON(422, gin.H{"message": err})
		return
	}
	var createCommentRequest requests.CreateCommentRequest
	if err := c.ShouldBindJSON(&createCommentRequest); err != nil {
		c.JSON(422, gin.H{"message": err})
		return
	}

	comment := &models.Comment{
		ResourceID:   postID,
		ResourceType: models.PostResource,
		Body:         createCommentRequest.Body,
		UserID:       sanctum.AuthUser.ID,
	}
	err = config.PulsarConfig.DB.Create(comment).Error
	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	c.JSON(201, gin.H{})
}

func AddLikeToComment(c *gin.Context) {

}

func AddDislikeToComment(c *gin.Context) {

}
