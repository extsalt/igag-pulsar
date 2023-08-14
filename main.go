package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pulsar/config"
	"pulsar/routes"
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
	config.PulsarConfig = &config.Config{DB: db}
	engine := routes.RegisterRoutes()
	err = engine.Run()
	if err != nil {
		return
	}
}
