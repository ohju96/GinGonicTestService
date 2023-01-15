package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"user-service/config"
	"user-service/router"
)

func main() {

	engine := gin.Default()

	config.EntConfig()
	router.InitRouter(engine)

	engine.Run()
}
