package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Response 标准响应结构
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Time    string      `json:"time"`
}

// HomeHandler 首页处理器
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Admin Simple</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            margin: 0;
            padding: 20px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            min-height: 100vh;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            background: rgba(255, 255, 255, 0.1);
            padding: 30px;
            border-radius: 15px;
            backdrop-filter: blur(10px);
            box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
        }
        h1 {
            text-align: center;
            margin-bottom: 30px;
            font-size: 2.5em;
            text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
        }
        .status {
            background: rgba(255, 255, 255, 0.2);
            padding: 20px;
            border-radius: 10px;
            margin: 20px 0;
        }
        .time {
            font-size: 1.2em;
            text-align: center;
            margin-top: 20px;
        }
        .api-links {
            margin-top: 30px;
            text-align: center;
        }
        .api-links a {
            color: white;
            text-decoration: none;
            margin: 0 10px;
            padding: 10px 20px;
            background: rgba(255, 255, 255, 0.2);
            border-radius: 5px;
            transition: background 0.3s;
        }
        .api-links a:hover {
            background: rgba(255, 255, 255, 0.3);
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🎉 Admin Simple 服务器</h1>
        <div class="status">
            <h2>✅ 服务器状态: 运行中</h2>
            <p>这是一个简单的Go Web服务器示例。</p>
        </div>
        <div class="time">
            当前时间: <span id="current-time"></span>
        </div>
        <div class="api-links">
            <a href="/api/health">健康检查</a>
            <a href="/api/info">系统信息</a>
            <a href="/api/users">用户列表</a>
        </div>
    </div>
    <script>
        function updateTime() {
            const now = new Date();
            document.getElementById('current-time').textContent = 
                now.toLocaleString('zh-CN');
        }
        updateTime();
        setInterval(updateTime, 1000);
    </script>
</body>
</html>`
	fmt.Fprint(w, html)
}

// HealthHandler 健康检查处理器
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: true,
		Message: "服务运行正常",
		Data: map[string]interface{}{
			"service": "admin_simple",
			"version": "1.0.0",
			"uptime":  "运行中",
		},
		Time: time.Now().Format(time.RFC3339),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// InfoHandler 系统信息处理器
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: true,
		Message: "系统信息",
		Data: map[string]interface{}{
			"service":     "admin_simple",
			"version":     "1.0.0",
			"environment": "development",
			"timestamp":   time.Now().Unix(),
		},
		Time: time.Now().Format(time.RFC3339),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UsersHandler 用户列表处理器
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []map[string]interface{}{
		{"id": 1, "name": "张三", "email": "zhangsan@example.com", "role": "admin"},
		{"id": 2, "name": "李四", "email": "lisi@example.com", "role": "user"},
		{"id": 3, "name": "王五", "email": "wangwu@example.com", "role": "user"},
	}
	
	response := Response{
		Success: true,
		Message: "用户列表获取成功",
		Data:    users,
		Time:    time.Now().Format(time.RFC3339),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
