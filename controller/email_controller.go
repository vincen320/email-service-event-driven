package controller

import "github.com/gin-gonic/gin"

type EmailController interface {
	Ping(c *gin.Context)
}
