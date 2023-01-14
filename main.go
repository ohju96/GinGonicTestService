package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"sample/config"
	"sample/router"
)

func main() {
	engine := gin.Default()

	config.EntConfig()
	router.InitRouters(engine)

	engine.Run()
}
