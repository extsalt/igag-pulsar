package routes

import (
	"github.com/gin-gonic/gin"
	"pulsar/handlers"
)

func RegisterRoutes() *gin.Engine {
	engine := gin.Default()
	engine.POST("/users/login", handlers.LoginUser)
	engine.POST("/users/register", handlers.RegisterUser)
	engine.POST("/users/logout", handlers.Logout)
	engine.GET("/posts", handlers.GetPosts)
	engine.POST("/posts", handlers.AddPost)
	return engine
}
