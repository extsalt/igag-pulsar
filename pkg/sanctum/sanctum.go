package sanctum

import (
	"github.com/extsalt/gojwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pulsar/config"
	"pulsar/models"
)

var AuthUser *models.User = nil

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, err := c.Cookie("iam")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "unauthorized",
			})
			return
		}
		if jwt == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "unauthorized",
			})
			return
		}
		payload, err := gojwt.VerifySignature(jwt, "secret")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "unauthorized",
			})
			return
		}
		var user models.User
		err = config.PulsarConfig.DB.Where(&models.User{Username: payload.Subject}).First(&user).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "unauthorized",
			})
			return
		}
		AuthUser = &user
		c.Next()
	}
}
