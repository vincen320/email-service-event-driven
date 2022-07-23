package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type EmailControllerImpl struct {
}

func NewEmailController() EmailController {
	return &EmailControllerImpl{}
}
func (ec *EmailControllerImpl) Ping(c *gin.Context) {
	fmt.Fprintf(c.Writer, "pong")
}
