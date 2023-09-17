package handlers

import (
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// GetSignature returns a signature for signing APIs calls.
func GetSignature(c *gin.Context) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	params := map[string][]string{
		"timestamp": {timestamp},
	}
	resp, err := api.SignParameters(params, "V5gXwLsc4NVc6Nw-7kb-6UxUdxI")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Could not generate signature",
		})
		return
	}
	c.JSON(200, gin.H{
		"signature": resp,
		"timestamp": timestamp,
		"apiKey":    "583843547542968",
	})
}
