package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"pulsar/handlers/responses"
	"pulsar/models"
	"pulsar/pkg/sanctum"
	"pulsar/repositories"
	"strconv"
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
	postID, err := strconv.ParseUint(likeQueryBind.PostID, 10, 64)
	if err != nil {
		c.JSON(404, *responses.GetNotFoundResponse())
		return
	}
	post, err := repositories.FindPostById(postID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, *responses.GetNotFoundResponse())
		return
	}
	action, err := repositories.FindActionOnResource(postID, models.PostResource, sanctum.AuthUser.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_, err = repositories.AddActionToResource(postID, models.PostResource, sanctum.AuthUser.ID, models.LikeAction)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Could not add like",
			})
			return
		}
		post.LikesCount = post.LikesCount + 1
		repositories.Repo().Save(post)
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
	_, err = repositories.AddActionToResource(postID, models.PostResource, sanctum.AuthUser.ID, models.LikeAction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Could not add like",
		})
		return
	}
	post.LikesCount = post.LikesCount + 1
	post.DislikesCount = post.DislikesCount - 1
	repositories.Repo().Save(post)
	c.AbortWithStatusJSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "action added",
	})
	return
}

func AddDislikeToPosts(c *gin.Context) {
	var likeQueryBind LikeQueryBind
	if err := c.ShouldBindUri(&likeQueryBind); err != nil {
		c.JSON(404, gin.H{"message": err})
		return
	}
	postID, err := strconv.ParseUint(likeQueryBind.PostID, 10, 64)
	if err != nil {
		c.JSON(404, *responses.GetNotFoundResponse())
		return
	}
	post, err := repositories.FindPostById(postID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, *responses.GetNotFoundResponse())
		return
	}
	action, err := repositories.FindActionOnResource(postID, models.PostResource, sanctum.AuthUser.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		_, err = repositories.AddActionToResource(postID, models.PostResource, sanctum.AuthUser.ID, models.DislikeAction)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Could not add dislike",
			})
			return
		}
		post.DislikesCount = post.DislikesCount + 1
		repositories.Repo().Save(post)
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
	_, err = repositories.AddActionToResource(postID, models.PostResource, sanctum.AuthUser.ID, models.DislikeAction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Could not add like",
		})
		return
	}
	post.LikesCount = post.LikesCount - 1
	post.DislikesCount = post.DislikesCount + 1
	repositories.Repo().Save(post)
	c.AbortWithStatusJSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "action added",
	})
	return
}
