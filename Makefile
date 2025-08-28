# Makefile for admin_simple

# 变量定义
BINARY_NAME=admin_simple
BUILD_DIR=build
MAIN_FILE=main.go

# 默认目标
.PHONY: all
all: clean build

# 构建项目
.PHONY: build
build:
	@echo "🔨 构建项目..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ 构建完成: $(BUILD_DIR)/$(BINARY_NAME)"

# 运行项目
.PHONY: run
run:
	@echo "🚀 启动服务器..."
	go run $(MAIN_FILE)

# 测试项目
.PHONY: test
test:
	@echo "🧪 运行测试..."
	go test ./...

# 清理构建文件
.PHONY: clean
clean:
	@echo "🧹 清理构建文件..."
	rm -rf $(BUILD_DIR)
	@echo "✅ 清理完成"

# 格式化代码
.PHONY: fmt
fmt:
	@echo "🎨 格式化代码..."
	go fmt ./...

# 检查代码
.PHONY: vet
vet:
	@echo "🔍 检查代码..."
	go vet ./...

# 安装依赖
.PHONY: deps
deps:
	@echo "📦 安装依赖..."
	go mod tidy
	go mod download

# 开发模式（热重载）
.PHONY: dev
dev:
	@echo "🔥 开发模式启动..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "⚠️  Air未安装，使用普通模式运行"; \
		go run $(MAIN_FILE); \
	fi

# 帮助信息
.PHONY: help
help:
	@echo "📋 可用命令:"
	@echo "  build  - 构建项目"
	@echo "  run    - 运行项目"
	@echo "  test   - 运行测试"
	@echo "  clean  - 清理构建文件"
	@echo "  fmt    - 格式化代码"
	@echo "  vet    - 检查代码"
	@echo "  deps   - 安装依赖"
	@echo "  dev    - 开发模式"
	@echo "  help   - 显示帮助信息"

