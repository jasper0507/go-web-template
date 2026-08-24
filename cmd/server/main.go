package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/jasper0507/go-web-template/internal/config"
	"github.com/jasper0507/go-web-template/internal/database"
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
		return fmt.Errorf("init logger: %w", err)
	}
	slog.SetDefault(logger)

	// 初始化数据库连接
	db, err := database.Open(&cfg.MySQL)
	if err != nil {
		return fmt.Errorf("init MySQL: %w", err)
	}

	slog.Info(
		"MySQL initialized",
		"max_open_conns", cfg.MySQL.MaxOpenConns,
		"max_idle_conns", cfg.MySQL.MaxIdleConns,
		"max_lifetime", cfg.MySQL.ConnMaxLifetime,
	)

	defer func() {
		if err := database.Close(db); err != nil {
			slog.Warn("close MySQL", "error", err)
		}
	}()

	// 创建路由
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	slog.Info("starting HTTP server", "address", cfg.HTTP.Addr)
	if err := r.Run(cfg.HTTP.Addr); err != nil {
		return fmt.Errorf("run HTTP server: %w", err)
	}
	return nil
}
