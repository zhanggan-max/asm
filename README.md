# ASM - Agent Skill Manager

[中文版本 (Chinese Version)](README_CN.md)

**asm** is a lightweight, zero-dependency package manager designed specifically for AI Agent Skills. It manages dependencies, versioning, and installations, similar to `npm` or `go mod`, but for agentic workflows.

## Features

- **Zero Dependency**: Written in Go, distributed as a single binary.
- **Git-based Registry**: Install skills directly from GitHub or any Git repository.
- **Dependency Resolution**: Automatically resolves and installs transitive dependencies (A -> B -> C).
- **Lockfile**: Ensures reproducible builds with `skill.lock`.
- **Versioning**: Supports Git Tags and Branches (e.g., `@v1.0.0`).

## Installation

### From Source

```bash
git clone https://github.com/YOUR_USERNAME/asm.git
cd asm
go build -o asm cmd/asm/main.go
# Move binary to path
mv asm /usr/local/bin/
```

## Usage

### Initialize a Project

```bash
asm init
```
Creates a `skill.json` in your current directory.

### Install a Skill

From GitHub shorthand:
```bash
asm install user/repo
```

From specific URL:
```bash
asm install https://github.com/user/repo.git
```

With specific version:
```bash
asm install user/repo@v1.0.0
```

### Install Dependencies

If you have a `skill.json` or `skill.lock`:
```bash
asm install
```

### List Installed Skills

```bash
asm list
```

## Project Structure

- **skill.json**: Manifest file declaring direct dependencies.
- **skill.lock**: Auto-generated file locking exact commit hashes.
- **.asm_modules/**: Directory where skills are installed (flat structure).

## Development

If you want to contribute to **asm**, please read our [Contributing Guide](CONTRIBUTING.md).

## License

MIT
