package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pulsar/config"
	"pulsar/models"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/igag_dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(100)
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	config.PulsarConfig = &config.Config{DB: db}
	r.GET("/", func(ctx *gin.Context) {
		_, err := ctx.Cookie("igag")
		if err != nil {
			ctx.SetCookie("igag", "ok", 3600, "/", "localhost", false, true)
		}
		user := &models.User{
			Username: "a",
			Password: "a",
			Email:    "a@gmail.com",
		}
		tx := config.PulsarConfig.DB.Create(user)
		ctx.JSON(200, gin.H{
			"ID":   user.ID,
			"rows": tx.RowsAffected,
		})
	})
	err = r.Run()
	if err != nil {
		return
	}
}
