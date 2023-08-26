package routes

import (
	"github.com/gin-gonic/gin"
	"pulsar/handlers"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD")
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
	engine.PUT("/posts/:postID/like", handlers.AddLikeToPost)
	engine.PUT("/posts/:postID/dislike", handlers.AddDislikeToPosts)
	engine.POST("/posts/:postID/comments", handlers.AddComment)
	engine.POST("/comments/:commentID/replies", handlers.AddReply)
	engine.POST("/comments/:commentID/like", handlers.AddLikeToComment)
	engine.POST("/comments/:commentID/dislike", handlers.AddDislikeToComment)
	engine.GET("/cloudinary/signature", handlers.GetSignature)
	engine.GET("/oauth/google", handlers.LoginWithGoogle)
	return engine
}
