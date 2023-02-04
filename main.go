package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"user_service/config"
	"user_service/router"
)

func main() {

	// Init app
	app := gin.Default()

	// Init config toml
	tomlConfig := config.IntiTomlConfig("./config/config.toml")

	// Init database
	config.InitDB(tomlConfig)

	// Init Router
	router.InitRouter(app, tomlConfig)

	// app run !!
	app.Run()
}
