package routes

import (
	bugsnaggin "github.com/bugsnag/bugsnag-go-gin"
	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/gin-gonic/gin"
	"pulsar/handlers"
	"pulsar/pkg/sanctum"
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
	engine.Use(bugsnaggin.AutoNotify(bugsnag.Configuration{
		APIKey:          "7b409ddb936f99d820a65cf006adf0a4",
		ProjectPackages: []string{"main", "github.com/extsalt/igag-pulsar"},
	}))
	engine.Use(CORSMiddleware())
	auth := engine.Group("/")
	auth.Use(sanctum.Auth())
	{
		auth.GET("/me", handlers.GetMe)
		auth.POST("/posts", handlers.AddPost)
		auth.POST("/posts/:postID/like", handlers.AddLikeToPost)
		auth.POST("/posts/:postID/dislike", handlers.AddDislikeToPosts)
		auth.POST("/posts/:postID/comments", handlers.AddComment)
		auth.POST("/comments/:commentID/replies", handlers.AddReply)
		auth.POST("/comments/:commentID/like", handlers.AddLikeToComment)
		auth.POST("/comments/:commentID/dislike", handlers.AddDislikeToComment)
		auth.GET("/cloudinary/signature", handlers.GetSignature)
	}
	engine.POST("/users/login", handlers.LoginUser)
	engine.POST("/users/register", handlers.RegisterUser)
	engine.POST("/users/logout", handlers.Logout)
	engine.GET("/posts", handlers.GetPosts)
	engine.GET("/oauth/google", handlers.LoginWithGoogle)
	return engine
}
