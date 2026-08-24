package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasper0507/go-web-template/internal/middleware"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(
		middleware.RequestLogger(),
		gin.Recovery(),
	)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}
