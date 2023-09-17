package handlers

import (
	"github.com/gin-gonic/gin"
	"pulsar/pkg/sanctum"
)

func GetMe(c *gin.Context) {
	c.JSON(200, gin.H{
		"username": sanctum.AuthUser.Email,
	})
}
