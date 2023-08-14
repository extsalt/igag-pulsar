package routes

import (
	"github.com/gin-gonic/gin"
	"pulsar/handlers"
)

func RegisterRoutes() *gin.Engine {
	engine := gin.Default()
	engine.GET("/posts", handlers.GetPosts)
	engine.POST("/posts", handlers.AddPost)
	return engine
}
