# Contributing to ASM

Thank you for your interest in improving **asm**! This document provides a guide for setting up your development environment and understanding the project's architecture.

## 🛠 Development Setup

1. **Prerequisites**: Install [Go 1.21+](https://go.dev/doc/install).
2. **Clone the repo**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/asm.git
   cd asm
   ```
3. **Build from source**:
   ```bash
   go build -o asm cmd/asm/main.go
   ```
4. **Run tests**:
   ```bash
   go test ./...
   ```

## 🏗 Project Architecture

The project follows a modular structure:

- **`cmd/asm/`**: CLI entry point. Handles argument parsing and routing.
- **`internal/commands/`**: Implementation of CLI commands (`init`, `install`, `list`).
- **`pkg/manifest/`**: Logic for reading/writing `skill.json`.
- **`pkg/lockfile/`**: Logic for `skill.lock` management.
- **`pkg/registry/`**: Package name resolution (e.g., shorthand to GitHub URL).
- **`internal/utils/`**: Shared helpers (Git operations, file system utilities).

## 🚀 How to Add a New Command

1. Create a new file in `internal/commands/your_command.go`.
2. Implement your logic in a function like `RunYourCommand(...)`.
3. Register the command in `cmd/asm/main.go` within the switch-case block.
4. Update the `usage()` function in `main.go`.

## 📝 Coding Standards

- **Zero Dependency**: Avoid adding external libraries unless absolutely necessary. We prefer using the Go standard library.
- **Error Handling**: Always return errors and provide context in error messages.
- **Formatting**: Run `go fmt ./...` before committing.

## 🐛 Reporting Bugs

Please use GitHub Issues to report bugs. Provide as much detail as possible, including your OS, Go version, and steps to reproduce.

## 📬 Pull Request Process

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/amazing-feature`).
3. Commit your changes with clear messages.
4. Push to the branch.
5. Open a Pull Request.

---

Happy Coding! 🤖
