package main

import (
	"gin_test/auto"
	"gin_test/config"
	"gin_test/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.CollectRoute(r)
	conf := config.GetConfig()
	r.Run(":" + conf.Httpport)
}

func init() {
	config.LoadConfig()
	config.InitMySql()
	go auto.AutoMigrate()
}
