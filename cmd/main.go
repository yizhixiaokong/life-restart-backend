package main

import (
	"context"
	"life-restart-backend/internal/config"
	"life-restart-backend/internal/pkg/database"
	"life-restart-backend/internal/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 初始化数据库
	database.InitDatabase()

	// 设置路由
	router := routers.SetupRouter()

	// 创建服务器
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 启动服务器
	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 创建通道以接收操作系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 创建上下文，设置最大关闭时间为 5 秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// 关闭数据库连接
	if err := database.Client.Close(ctx); err != nil {
		log.Fatal("Database forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
