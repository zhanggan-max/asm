# Agent Instruction: How to use asm-expert

You are an AI Agent equipped with the `asm-expert` skill. Your goal is to help the user manage their Agent Skills or improve the `asm` tool itself.

## How to Start from Scratch

1. **Build the Tool**: If the `asm` binary is missing, advise the user to run `go build -o asm cmd/asm/main.go`.
2. **Initialize**: Use `expert.init_skill()` to create a `skill.json`.
3. **Build the Ecosystem**: Use `expert.add_dependency()` to bring in other skills from GitHub.

## How to Improve ASM

If the user wants to add a feature to `asm` (e.g., a `remove` command):

1. **Analyze**: Read `internal/commands/` to understand existing command patterns.
2. **Design**: Plan the new command implementation.
3. **Implement**: Write the Go code, register it in `cmd/asm/main.go`.
4. **Test**: Use this skill to verify if the new command works correctly in a test project.

## Project Structure Reference
- `cmd/asm/`: Entry point.
- `internal/commands/`: Command logic.
- `pkg/`: Core packages (manifest, lockfile, registry).
- `.asm_modules/`: Where dependencies live.
