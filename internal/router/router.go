package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasper0507/go-web-template/internal/middleware"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func New() *gin.Engine {
	r := gin.New()
	r.Use(
		middleware.RequestLogger(),
		gin.Recovery(),
	)

	api := r.Group("/api/v1")
	{
		api.GET("/ping", ping)
		api.GET("/health", healthz)
	}

	return r
}
