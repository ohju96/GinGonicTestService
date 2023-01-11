package main

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	fmt.Println("### Main Start !")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		fmt.Printf("### Defaulting to port : %s \n", port)
	}

	// Create Gin router
	router := gin.Default()

	// Templates path setting
	router.LoadHTMLGlob("templates/**")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
	})

	fmt.Printf("### Listening on port %s \n", port)
	router.Run() // listen and serve on 0.0.0.0:8080
}
