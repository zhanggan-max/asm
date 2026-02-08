# ASM - Agent Skill 管理器

**asm** 是一个轻量级、零依赖的包管理工具，专为 AI Agent Skill 设计。它能够管理依赖关系、版本控制和安装流程，类似于针对 Agent 场景的 `npm` 或 `go mod`。

## 核心特性

- **零依赖**: 使用 Go 编写，分发为单个二进制文件，无需预装环境。
- **基于 Git 的注册表**: 直接从 GitHub 或任何 Git 仓库安装 Skill。
- **递归依赖解析**: 自动解析并安装传递依赖 (A -> B -> C)。
- **锁定文件**: 通过 `skill.lock` 确保安装的可重现性。
- **版本控制**: 支持 Git Tag 和分支 (例如 `@v1.0.0`)。

## 安装方式

### 从源码编译

```bash
git clone https://github.com/YOUR_USERNAME/asm.git
cd asm
go build -o asm cmd/asm/main.go
# 移动到系统路径
mv asm /usr/local/bin/
```

## 使用指南

### 初始化项目

```bash
asm init
```
这将在当前目录下创建一个 `skill.json` 文件。

### 安装 Skill

使用 GitHub 简写：
```bash
asm install user/repo
```

使用完整 URL：
```bash
asm install https://github.com/user/repo.git
```

指定版本：
```bash
asm install user/repo@v1.0.0
```

### 安装项目所有依赖

如果你已经有了 `skill.json` 或 `skill.lock`：
```bash
asm install
```

### 查看已安装列表

```bash
asm list
```

### 开发者助手 (Built-in Skill)

本项目内置了一个名为 `asm-expert` 的 Skill（位于 `skills/asm-expert`），它可以帮助 AI Agent：
- 快速理解 `asm` 的架构。
- 自动化执行初始化和依赖管理任务。
- 引导开发者对 `asm` 进行二次开发。

如果你正在使用支持 Skill 的 Agent，可以直接让它加载该目录以获得 `asm` 的操作能力。

## 项目结构

- **skill.json**: 清单文件，声明直接依赖。
- **skill.lock**: 自动生成的文件，锁定精确的 Commit Hash。
- **.asm_modules/**: Skill 的安装目录（采用扁平化结构）。

## 许可证

MIT
