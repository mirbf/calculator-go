# Go Fyne 计算器

一个使用 Go 语言和 Fyne 框架构建的现代化计算器应用程序，具有图形用户界面、计算历史记录和跨平台支持。

## 功能特性

### 🧮 核心计算功能
- **基本算术运算**: 加法(+)、减法(-)、乘法(×)、除法(÷)
- **取模运算**: 支持取模运算(%)
- **小数运算**: 完整的小数点运算支持
- **连续运算**: 支持连续进行多次计算
- **错误处理**: 智能的除零错误检测和提示

### 📱 用户界面
- **现代化设计**: 基于 Fyne 框架的原生 GUI
- **响应式布局**: 自适应窗口大小的按钮布局
- **实时显示**: 即时显示输入和计算结果
- **清晰的视觉反馈**: 直观的按钮交互效果

### 📚 历史记录
- **计算历史**: 自动保存所有计算记录
- **历史查看**: 专门的历史记录标签页
- **持久化存储**: 历史记录在应用重启后保持
- **清除功能**: 支持清除历史记录

### ⚙️ 配置管理
- **用户配置**: 可自定义的应用设置
- **自动保存**: 配置更改自动保存
- **跨平台路径**: 智能的配置文件路径管理

### 🖥️ 跨平台支持
- **Windows**: 完整支持 Windows 10/11
- **macOS**: 原生 macOS 应用体验
- **Linux**: 支持主流 Linux 发行版

## 项目结构

```
calculator-go/
├── main.go                    # 程序入口点
├── internal/                  # 内部模块
│   ├── calculator/           # 计算器核心逻辑
│   │   ├── calculator.go     # 计算器实现
│   │   └── calculator_test.go # 单元测试
│   ├── config/              # 配置管理
│   │   ├── config.go        # 配置实现
│   │   └── config_test.go   # 配置测试
│   └── ui/                  # 用户界面
│       └── ui.go            # UI 实现
├── go.mod                   # Go 模块定义
├── go.sum                   # 依赖校验和
├── Makefile                 # 构建脚本
├── .gitignore              # Git 忽略文件
├── README.md               # 项目说明
└── PROJECT_SUMMARY.md      # 项目总结
```

## 模块设计

### 📊 Calculator 模块
- **核心功能**: 数字输入、运算符处理、计算执行
- **状态管理**: 当前显示值、运算符、操作数管理
- **历史记录**: 计算过程和结果的记录
- **错误处理**: 除零等异常情况的处理

### ⚙️ Config 模块
- **配置结构**: 应用设置的定义和管理
- **文件操作**: 配置的加载、保存和验证
- **路径管理**: 跨平台的配置文件路径处理
- **默认值**: 合理的默认配置设置

### 🎨 UI 模块
- **界面构建**: 计算器界面的创建和布局
- **事件处理**: 按钮点击和用户交互
- **显示更新**: 实时更新计算结果显示
- **菜单系统**: 应用菜单和选项管理

## 安装和运行

### 前置要求
- Go 1.21 或更高版本
- Git（用于克隆仓库）

### 克隆项目
```bash
git clone https://github.com/mirbf/calculator-go.git
cd calculator-go
```

### 安装依赖
```bash
go mod tidy
```

### 运行应用
```bash
# 直接运行
go run main.go

# 或使用 Makefile
make run
```

### 构建应用
```bash
# 构建当前平台
make build

# 构建所有平台
make build-all

# 构建特定平台
make build-windows  # Windows
make build-mac      # macOS
make build-linux    # Linux
```

## 测试

### 运行测试
```bash
# 运行所有测试
make test

# 运行测试并生成覆盖率报告
make test-coverage

# 生成 HTML 覆盖率报告
make coverage-html
```

### 测试覆盖
- **Calculator 模块**: 完整的单元测试覆盖
- **Config 模块**: 配置管理功能测试
- **错误场景**: 异常情况和边界条件测试

## 使用说明

### 基本操作
1. **数字输入**: 点击数字按钮输入数字
2. **运算符**: 点击 +、-、×、÷、% 进行运算
3. **等号**: 点击 = 计算结果
4. **清除**: 点击 C 清除当前输入
5. **退格**: 点击 ← 删除最后一位
6. **小数点**: 点击 . 输入小数

### 高级功能
- **连续运算**: 可以连续进行多次计算而无需清除
- **历史查看**: 切换到"历史"标签页查看计算记录
- **菜单选项**: 使用菜单访问更多功能和设置

## 数据文件

应用程序会在用户目录下创建配置文件夹：
- **Windows**: `%APPDATA%\calculator-go\`
- **macOS**: `~/Library/Application Support/calculator-go/`
- **Linux**: `~/.config/calculator-go/`

包含的文件：
- `config.json`: 应用配置
- `history.json`: 计算历史记录

## 开发指南

### 代码规范
```bash
# 格式化代码
make fmt

# 代码检查
make vet

# 代码质量检查（需要 golangci-lint）
make lint

# 完整检查
make check
```

### 开发模式
```bash
# 安装开发工具
make install-tools

# 开发模式（自动重载）
make dev
```

### 添加新功能
1. 在相应模块中添加功能代码
2. 编写对应的单元测试
3. 更新文档和注释
4. 运行完整测试确保质量

## 技术栈

- **语言**: Go 1.21+
- **GUI 框架**: Fyne v2.6.1
- **构建工具**: Make
- **测试框架**: Go 内置测试
- **版本控制**: Git

## 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。

## 贡献

欢迎贡献代码！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 更新日志

### v1.0.0 (2024-12-19)
- ✨ 初始版本发布
- 🧮 基本算术运算功能
- 📱 现代化 GUI 界面
- 📚 计算历史记录
- ⚙️ 配置管理系统
- 🖥️ 跨平台支持
- 🧪 完整的测试覆盖
- 📖 详细的文档说明

---

**享受计算的乐趣！** 🎉