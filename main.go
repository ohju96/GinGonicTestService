package main

import (
	"github.com/gin-gonic/gin"
	"sample/app/router"
	"sample/config"
)

func main() {

	// Init app
	app := gin.Default()

	// Init toml
	tomlConfig := config.InitTomlConfig("./config/config.toml")

	// Init mysql
	config.InitDB(tomlConfig)

	// Init router
	router.InitRouter(app)

	// app run !!
	app.Run()
}
