package service

import "github.com/gin-gonic/gin"

type User interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Login(c *gin.Context)
}
