package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/jasper0507/go-web-template/internal/config"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(cfg.HTTP.Addr); err != nil {
		return fmt.Errorf("run HTTP server: %w", err)
	}
	return nil
}
