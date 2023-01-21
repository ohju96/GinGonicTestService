package main

import (
	"github.com/gin-gonic/gin"
	"sample/config"
	"sample/router"
)

func main() {

	engine := gin.Default()

	config.MySQLConnection()
	router.InitRouter(engine)
	engine.Run()
}
