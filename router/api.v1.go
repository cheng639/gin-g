package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	RegisterHandler(http.MethodGet, "/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
