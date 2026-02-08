# 参与贡献 ASM

感谢你对改进 **asm** 的兴趣！本指南将帮助你搭建开发环境并理解项目的架构设计。

## 🛠 开发环境搭建

1. **前置条件**: 安装 [Go 1.21+](https://go.dev/doc/install)。
2. **克隆仓库**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/asm.git
   cd asm
   ```
3. **编译**:
   ```bash
   go build -o asm cmd/asm/main.go
   ```
4. **运行测试**:
   ```bash
   go test ./...
   ```

## 🏗 项目架构

项目采用模块化设计：

- **`cmd/asm/`**: CLI 入口，负责参数解析和命令分发。
- **`internal/commands/`**: 各个命令的具体实现 (`init`, `install`, `list`)。
- **`pkg/manifest/`**: `skill.json` 的读写逻辑。
- **`pkg/lockfile/`**: `skill.lock` 的管理逻辑。
- **`pkg/registry/`**: 包名解析逻辑（例如将简写转换为 GitHub URL）。
- **`internal/utils/`**: 公用工具类（Git 操作、文件系统等）。

## 🚀 如何添加新命令

1. 在 `internal/commands/` 目录下创建新文件，例如 `your_command.go`。
2. 在该文件中实现逻辑，如 `RunYourCommand(...)` 函数。
3. 在 `cmd/asm/main.go` 的 switch-case 块中注册新命令。
4. 更新 `main.go` 中的 `usage()` 函数以显示新命令。

## 📝 编码标准

- **零依赖**: 除非绝对必要，否则避免添加外部库。我们优先使用 Go 标准库。
- **错误处理**: 始终返回 error 并提供清晰的错误上下文。
- **格式化**: 提交前运行 `go fmt ./...`。

## 🐛 报告问题

请使用 GitHub Issues 报告 Bug。请提供详细信息，包括操作系统、Go 版本以及复现步骤。

## 📬 提交 PR

1. Fork 本仓库。
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)。
3. 提交更改，并附带清晰的 Commit 信息。
4. 推送到你的分支。
5. 发起 Pull Request。

---

祝你开发愉快! 🤖
