package commands

import (
	"asm/pkg/manifest"
	"fmt"
	"os"
)

// RunList shows installed packages
func RunList() error {
	// 1. Check manifest deps
	m, err := manifest.Load("skill.json")
	if err != nil {
		fmt.Println("No skill.json found or invalid.")
	} else {
		fmt.Println("Dependencies declared in skill.json:")
		if len(m.Dependencies) == 0 {
			fmt.Println("  (none)")
		} else {
			for name, ver := range m.Dependencies {
				fmt.Printf("  - %s: %s\n", name, ver)
			}
		}
	}

	// 2. Check physical directory
	fmt.Println("\nInstalled modules in .asm_modules/:")
	entries, err := os.ReadDir(modulesDir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("  (directory not found)")
			return nil
		}
		return err
	}

	if len(entries) == 0 {
		fmt.Println("  (empty)")
		return nil
	}

	for _, e := range entries {
		if e.IsDir() {
			fmt.Printf("  - %s\n", e.Name())
		}
	}
	return nil
}
