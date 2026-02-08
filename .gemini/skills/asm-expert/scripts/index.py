import subprocess
import os

class ASMHelper:
    def __init__(self, bin_path="./asm"):
        self.bin_path = bin_path

    def run_cmd(self, args):
        res = subprocess.run([self.bin_path] + args, capture_output=True, text=True)
        return res.stdout if res.returncode == 0 else res.stderr

if __name__ == "__main__":
    # Helper for complex logic if needed
    print("ASM Helper Script Loaded")
