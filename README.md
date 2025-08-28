# Admin Simple

一个简单的Go Web服务器项目，提供基本的HTTP API和Web界面。

## 🚀 功能特性

- ✅ 简洁的Web界面
- ✅ RESTful API接口
- ✅ 健康检查端点
- ✅ 用户管理功能
- ✅ 配置管理
- ✅ 日志记录

## 📁 项目结构

```
admin_simple/
├── main.go              # 主程序入口
├── config/
│   └── config.go        # 配置管理
├── handlers/
│   └── handlers.go      # HTTP处理器
├── models/
│   └── user.go          # 数据模型
├── utils/
│   └── logger.go        # 日志工具
├── go.mod               # Go模块文件
├── .gitignore           # Git忽略文件
└── README.md            # 项目说明
```

## 🛠️ 安装和运行

### 前置要求

- Go 1.19 或更高版本

### 快速开始

1. **克隆项目**
   ```bash
   git clone https://github.com/kulilinn/admin_simple.git
   cd admin_simple
   ```

2. **运行项目**
   ```bash
   go run main.go
   ```

3. **访问应用**
   - Web界面: http://localhost:8080
   - API文档: http://localhost:8080/api/health

## 🌐 API接口

### 健康检查
- **GET** `/api/health` - 服务健康状态

### 系统信息
- **GET** `/api/info` - 系统信息

### 用户管理
- **GET** `/api/users` - 获取用户列表

## ⚙️ 配置

通过环境变量配置应用：

```bash
# 服务器配置
PORT=8080
HOST=localhost

# 数据库配置
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=admin
DB_PASSWORD=password
DB_NAME=admin_simple
```

## 🧪 测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./handlers
```

## 📦 构建

```bash
# 构建可执行文件
go build -o admin_simple main.go

# 运行构建后的程序
./admin_simple
```

## 🤝 贡献

欢迎提交Issue和Pull Request！

## 📄 许可证

MIT License