package sanctum

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	*gin.Context
}

type Authenticate interface {
	AuthUser() any
}

func (r *Request) AuthUser() any {
	fmt.Println("hello")
	return nil
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("iam")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "unauthorized",
			})
			return
		}
		c.Next()
	}
}
