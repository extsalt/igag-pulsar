package handlers

import (
	"github.com/gin-gonic/gin"
	"pulsar/config"
	"pulsar/handlers/requests"
	"pulsar/models"
)

func LoginUser(c *gin.Context) {

}

func RegisterUser(c *gin.Context) {
	var registerRequest requests.RegisterUserRequest
	var err error
	if err = c.BindJSON(&registerRequest); err == nil {
		user := &models.User{
			Username: registerRequest.Username,
			Password: registerRequest.Password,
			Email:    registerRequest.Email,
			Active:   true,
		}
		config.PulsarConfig.DB.Create(&user)
		c.JSON(201, gin.H{
			"user": user,
		})
		return
	}
	c.JSON(422, gin.H{
		"error": err,
	})
}
