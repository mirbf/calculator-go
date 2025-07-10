# Makefile for Calculator Go Project

# 变量定义
APP_NAME=calculator
MAIN_FILE=main.go
BUILD_DIR=build
COVERAGE_FILE=coverage.out

# 默认目标
.PHONY: all
all: clean test build

# 清理构建文件
.PHONY: clean
clean:
	@echo "清理构建文件..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(COVERAGE_FILE)
	@rm -f $(APP_NAME) $(APP_NAME).exe $(APP_NAME)-mac $(APP_NAME)-linux

# 安装依赖
.PHONY: deps
deps:
	@echo "安装依赖..."
	@go mod tidy
	@go mod download

# 运行测试
.PHONY: test
test:
	@echo "运行测试..."
	@go test -v ./...

# 运行测试并生成覆盖率报告
.PHONY: test-coverage
test-coverage:
	@echo "运行测试并生成覆盖率报告..."
	@go test -coverprofile=$(COVERAGE_FILE) ./...
	@go tool cover -func=$(COVERAGE_FILE)

# 生成HTML覆盖率报告
.PHONY: coverage-html
coverage-html: test-coverage
	@echo "生成HTML覆盖率报告..."
	@go tool cover -html=$(COVERAGE_FILE)

# 运行应用
.PHONY: run
run:
	@echo "运行应用..."
	@go run $(MAIN_FILE)

# 构建当前平台
.PHONY: build
build:
	@echo "构建应用 (当前平台)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# 构建所有平台
.PHONY: build-all
build-all: build-windows build-mac build-linux

# 构建 Windows 版本
.PHONY: build-windows
build-windows:
	@echo "构建 Windows 版本..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-windows.exe $(MAIN_FILE)

# 构建 macOS 版本
.PHONY: build-mac
build-mac:
	@echo "构建 macOS 版本..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-mac $(MAIN_FILE)

# 构建 Linux 版本
.PHONY: build-linux
build-linux:
	@echo "构建 Linux 版本..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux $(MAIN_FILE)

# 代码格式化
.PHONY: fmt
fmt:
	@echo "格式化代码..."
	@go fmt ./...

# 代码检查
.PHONY: vet
vet:
	@echo "代码检查..."
	@go vet ./...

# 代码质量检查（需要安装 golangci-lint）
.PHONY: lint
lint:
	@echo "代码质量检查..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint 未安装，跳过代码质量检查"; \
	fi

# 完整的代码检查
.PHONY: check
check: fmt vet lint test

# 开发模式（监听文件变化并重新运行）
.PHONY: dev
dev:
	@echo "开发模式 (需要安装 air)..."
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "air 未安装，使用普通运行模式"; \
		go run $(MAIN_FILE); \
	fi

# 安装开发工具
.PHONY: install-tools
install-tools:
	@echo "安装开发工具..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/cosmtrek/air@latest

# 显示帮助信息
.PHONY: help
help:
	@echo "可用的命令:"
	@echo "  all           - 清理、测试、构建"
	@echo "  clean         - 清理构建文件"
	@echo "  deps          - 安装依赖"
	@echo "  test          - 运行测试"
	@echo "  test-coverage - 运行测试并生成覆盖率报告"
	@echo "  coverage-html - 生成HTML覆盖率报告"
	@echo "  run           - 运行应用"
	@echo "  build         - 构建当前平台"
	@echo "  build-all     - 构建所有平台"
	@echo "  build-windows - 构建 Windows 版本"
	@echo "  build-mac     - 构建 macOS 版本"
	@echo "  build-linux   - 构建 Linux 版本"
	@echo "  fmt           - 格式化代码"
	@echo "  vet           - 代码检查"
	@echo "  lint          - 代码质量检查"
	@echo "  check         - 完整的代码检查"
	@echo "  dev           - 开发模式"
	@echo "  install-tools - 安装开发工具"
	@echo "  help          - 显示此帮助信息"