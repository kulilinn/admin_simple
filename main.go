package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	
	"admin_simple/config"
	"admin_simple/handlers"
)

func main() {
	fmt.Println("🚀 启动 admin_simple 服务器...")
	
	// 加载配置
	cfg := config.LoadConfig()
	
	// 设置路由
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/api/health", handlers.HealthHandler)
	http.HandleFunc("/api/info", handlers.InfoHandler)
	http.HandleFunc("/api/users", handlers.UsersHandler)
	
	// 启动服务器
	port := ":" + cfg.Server.Port
	fmt.Printf("📡 服务器运行在 http://%s%s\n", cfg.Server.Host, port)
	fmt.Println("⏰ 启动时间:", time.Now().Format("2006-01-02 15:04:05"))
	
	log.Fatal(http.ListenAndServe(port, nil))
}
