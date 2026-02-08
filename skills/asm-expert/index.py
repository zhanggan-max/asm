import subprocess
import json
import os

class ASMExpert:
    def __init__(self, bin_path="./asm"):
        self.bin_path = bin_path

    def run_command(self, args):
        try:
            result = subprocess.run([self.bin_path] + args, capture_output=True, text=True)
            return {
                "stdout": result.stdout,
                "stderr": result.stderr,
                "exit_code": result.returncode
            }
        except FileNotFoundError:
            return {"error": f"ASM binary not found at {self.bin_path}. Please build it first."}

    def init_skill(self, name):
        """Initializes a new skill project."""
        return self.run_command(["init"])

    def add_dependency(self, pkg_spec):
        """Installs a new dependency and updates skill.json/skill.lock."""
        return self.run_command(["install", pkg_spec])

    def list_installed(self):
        """Lists all installed skills in the current project."""
        return self.run_command(["list"])

    def explain_architecture(self):
        """Returns a high-level overview of how ASM works for the agent."""
        return {
            "core_concept": "ASM is a Go-based package manager for Agent Skills.",
            "manifest": "skill.json (defines dependencies)",
            "lockfile": "skill.lock (locks commit hashes)",
            "storage": ".asm_modules/ (flat directory structure)",
            "workflow": "Recursive git-based dependency resolution"
        }

if __name__ == "__main__":
    # Example usage for an agent
    expert = ASMExpert()
    print("ASM Expert Skill Loaded.")
