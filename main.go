package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"sample/config"
	"sample/router"
)

func main() {

	engine := gin.Default()

	config.MySQLConnection()
	router.InitRouter(engine)
	engine.Run()
}
