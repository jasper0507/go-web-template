package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	config "github.com/jasper0507/go-web-template/internal/config"
	appLogger "github.com/jasper0507/go-web-template/internal/logger"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	// 创建日志记录器
	logger, err := appLogger.New(&cfg.Log)
	if err != nil {
		return fmt.Errorf("create logger: %w", err)
	}
	slog.SetDefault(logger)

	// 创建路由
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	logger.Info("starting HTTP server", "address", cfg.HTTP.Addr)
	if err := r.Run(cfg.HTTP.Addr); err != nil {
		return fmt.Errorf("run HTTP server: %w", err)
	}
	return nil
}
