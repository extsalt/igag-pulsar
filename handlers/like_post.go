package handlers

import (
	"github.com/gin-gonic/gin"
	"pulsar/handlers/responses"
	"pulsar/repositories"
)

type LikeQueryBind struct {
	PostID string `uri:"postID" binding:"required"`
}

func AddLikeToPost(c *gin.Context) {
	var likeQueryBind LikeQueryBind
	if err := c.ShouldBindUri(&likeQueryBind); err != nil {
		c.JSON(404, gin.H{"message": err})
		return
	}
	_, err := repositories.FindPostById(likeQueryBind.PostID)
	if err != nil {
		c.JSON(404, *responses.GetNotFoundResponse())
	}
	//check user has already liked the post
}

func AddDislikeToPosts(c *gin.Context) {

}
