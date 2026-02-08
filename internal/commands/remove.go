package commands

import (
	"asm/pkg/lockfile"
	"asm/pkg/manifest"
	"fmt"
	"os"
	"path/filepath"
)

// RunRemove removes a dependency
func RunRemove(pkgName string) error {
	if pkgName == "" {
		return fmt.Errorf("package name required")
	}

	fmt.Printf("Removing %s...\n", pkgName)

	// 1. Remove from skill.json
	m, err := manifest.Load("skill.json")
	if err == nil {
		if _, ok := m.Dependencies[pkgName]; ok {
			delete(m.Dependencies, pkgName)
			if err := m.Save("skill.json"); err != nil {
				return fmt.Errorf("failed to update skill.json: %v", err)
			}
			fmt.Println("  - Removed from skill.json")
		} else {
			fmt.Printf("  - Warning: %s not found in skill.json\n", pkgName)
		}
	}

	// 2. Remove from skill.lock
	l, err := lockfile.Load("skill.lock")
	if err == nil {
		if _, ok := l.Dependencies[pkgName]; ok {
			delete(l.Dependencies, pkgName)
			if err := l.Save("skill.lock"); err != nil {
				return fmt.Errorf("failed to update skill.lock: %v", err)
			}
			fmt.Println("  - Removed from skill.lock")
		}
	}

	// 3. Remove directory
	targetDir := filepath.Join(modulesDir, pkgName)
	if _, err := os.Stat(targetDir); err == nil {
		if err := os.RemoveAll(targetDir); err != nil {
			return fmt.Errorf("failed to remove directory %s: %v", targetDir, err)
		}
		fmt.Println("  - Deleted directory")
	} else {
		fmt.Printf("  - Warning: Directory %s does not exist\n", targetDir)
	}

	fmt.Printf("Package %s removed successfully.\n", pkgName)
	return nil
}
