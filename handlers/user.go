package handlers

import (
	"fmt"
	"github.com/extsalt/gojwt"
	"github.com/gin-gonic/gin"
	"pulsar/config"
	"pulsar/handlers/requests"
	"pulsar/models"
	"pulsar/pkg/googleoauth2"
	"pulsar/repositories"
	"time"
)

func LoginUser(c *gin.Context) {

}

// RegisterUser with provided credentials
func RegisterUser(c *gin.Context) {
	var registerRequest requests.RegisterUserRequest
	var err error
	err = c.BindJSON(&registerRequest)
	if err != nil {
		c.JSON(422, gin.H{
			"error": err,
		})
	}
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
}

// Logout user
func Logout(c *gin.Context) {
	//Logout user
}

func LoginWithGoogle(c *gin.Context) {
	code := c.Query("code")
	fmt.Println(code)
	if code == "" {
		c.JSON(422, gin.H{
			"error": "No valid code",
		})
		return
	}
	token, err := googleoauth2.GetToken(c, code)
	if err != nil {
		c.JSON(422, gin.H{
			"error": err,
		})
		return
	}
	oauth2Identity, err := googleoauth2.GetUserInfo(c, token)
	if err != nil {
		c.JSON(422, gin.H{
			"error": err,
		})
		return
	}
	user, err := repositories.FindUserByOauth2Identify(oauth2Identity)
	if err != nil {
		c.JSON(422, gin.H{
			"error": err,
		})
		return
	}
	jwt, err := gojwt.Create(user.Username, "web", "secret")
	if err != nil {
		c.JSON(422, gin.H{
			"error": err,
		})
		return
	}
	c.SetCookie("iam", jwt, int(time.Hour*24*400), "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"status": "success",
	})
}
