package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pulsar/config"
	"pulsar/handlers/requests"
	"pulsar/handlers/responses"
	"pulsar/models"
	"pulsar/pkg/sanctum"
	"pulsar/repositories"
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
	post, err := repositories.FindPostById(postID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
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
	post.CommentsCount = post.CommentsCount + 1
	repositories.Repo().Save(&post)
	c.JSON(201, gin.H{})
}

type CommentLikeQueryBind struct {
	CommentID string `uri:"commentID" binding:"required"`
}

func AddLikeToComment(c *gin.Context) {
	var likeQueryBind CommentLikeQueryBind
	if err := c.ShouldBindUri(&likeQueryBind); err != nil {
		c.JSON(404, gin.H{"message": err})
		return
	}
	commentID, err := strconv.ParseUint(likeQueryBind.CommentID, 10, 64)
	if err != nil {
		c.JSON(404, *responses.GetNotFoundResponse())
		return
	}
	comment, err := repositories.FindCommentById(commentID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, *responses.GetNotFoundResponse())
		return
	}
	action, err := repositories.FindActionOnResource(commentID, models.CommentResource, sanctum.AuthUser.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_, err = repositories.AddActionToResource(commentID, models.CommentResource, sanctum.AuthUser.ID, models.LikeAction)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Could not add like",
			})
			return
		}
		comment.LikesCount = comment.LikesCount + 1
		repositories.Repo().Save(comment)
		c.AbortWithStatusJSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "action added",
		})
		return
	}
	if action.Action == models.LikeAction {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "user already has this action on the resource",
		})
		return
	}
	err = repositories.RemoveActionFromResource(action.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Could not add like",
		})
		return
	}
	_, err = repositories.AddActionToResource(commentID, models.CommentResource, sanctum.AuthUser.ID, models.LikeAction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Could not add like",
		})
		return
	}
	comment.LikesCount = comment.LikesCount + 1
	comment.DislikesCount = comment.DislikesCount - 1
	repositories.Repo().Save(comment)
	c.AbortWithStatusJSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "action added",
	})
	return
}

func AddDislikeToComment(c *gin.Context) {
	var likeQueryBind CommentLikeQueryBind
	if err := c.ShouldBindUri(&likeQueryBind); err != nil {
		c.JSON(404, gin.H{"message": err})
		return
	}
	commentID, err := strconv.ParseUint(likeQueryBind.CommentID, 10, 64)
	if err != nil {
		c.JSON(404, *responses.GetNotFoundResponse())
		return
	}
	comment, err := repositories.FindCommentById(commentID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, *responses.GetNotFoundResponse())
		return
	}
	action, err := repositories.FindActionOnResource(commentID, models.CommentResource, sanctum.AuthUser.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_, err = repositories.AddActionToResource(commentID, models.CommentResource, sanctum.AuthUser.ID, models.DislikeAction)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Could not add dislike",
			})
			return
		}
		comment.DislikesCount = comment.DislikesCount + 1
		repositories.Repo().Save(comment)
		c.AbortWithStatusJSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "action added",
		})
		return
	}
	if action.Action == models.DislikeAction {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "user already has this action on the resource",
		})
		return
	}
	err = repositories.RemoveActionFromResource(action.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Could not add like",
		})
		return
	}
	_, err = repositories.AddActionToResource(commentID, models.CommentResource, sanctum.AuthUser.ID, models.DislikeAction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Could not add like",
		})
		return
	}
	comment.LikesCount = comment.LikesCount - 1
	comment.DislikesCount = comment.DislikesCount + 1
	repositories.Repo().Save(comment)
	c.AbortWithStatusJSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "action added",
	})
	return
}
