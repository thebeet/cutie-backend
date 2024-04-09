package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, world!",
	})
}
