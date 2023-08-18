package routes

import (
	"github.com/gin-gonic/gin"
	"pulsar/handlers"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RegisterRoutes() *gin.Engine {
	engine := gin.Default()
	engine.Use(CORSMiddleware())
	engine.POST("/users/login", handlers.LoginUser)
	engine.POST("/users/register", handlers.RegisterUser)
	engine.POST("/users/logout", handlers.Logout)
	engine.GET("/posts", handlers.GetPosts)
	engine.POST("/posts", handlers.AddPost)
	engine.POST("/posts/:postID/like", handlers.AddLikeToPost)
	engine.POST("/posts/:postID/dislike", handlers.AddDislikeToPosts)
	engine.POST("/posts/:postID/comments", handlers.AddComment)
	engine.POST("/comments/:commentID/replies", handlers.AddReply)
	engine.POST("/comments/:commentID/like", handlers.AddLikeToComment)
	engine.POST("/comments/:commentID/dislike", handlers.AddDislikeToComment)
	engine.GET("/cloudinary/signature", handlers.GetSignature)
	return engine
}
