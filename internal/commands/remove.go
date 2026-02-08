package commands

import (
	"asm/pkg/lockfile"
	"asm/pkg/manifest"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RunRemove removes a dependency
func RunRemove(pkgName string) error {
	if pkgName == "" {
		return fmt.Errorf("package name required")
	}

	// 1. Load manifest
	m, err := manifest.Load("skill.json")
	if err != nil {
		return fmt.Errorf("failed to load skill.json: %v", err)
	}

	// 2. Smart Match
	targetKey := ""
	if _, ok := m.Dependencies[pkgName]; ok {
		targetKey = pkgName
	} else {
		// Try fuzzy match: if user types "repo" but we have "user/repo"
		matches := []string{}
		for k := range m.Dependencies {
			// Check if the key ends with /pkgName or is exactly pkgName
			if k == pkgName || (strings.Contains(k, "/") && strings.HasSuffix(k, "/"+pkgName)) {
				matches = append(matches, k)
			}
		}

		if len(matches) == 1 {
			targetKey = matches[0]
			fmt.Printf("Smart match found: using '%s' instead of '%s'\n", targetKey, pkgName)
		} else if len(matches) > 1 {
			return fmt.Errorf("ambiguous name '%s': matches multiple dependencies: %v. Please use full name", pkgName, matches)
		}
	}

	if targetKey == "" {
		return fmt.Errorf("package '%s' not found in dependencies", pkgName)
	}

	fmt.Printf("Removing %s...\n", targetKey)

	// 3. Remove from skill.json
	delete(m.Dependencies, targetKey)
	if err := m.Save("skill.json"); err != nil {
		return fmt.Errorf("failed to update skill.json: %v", err)
	}
	fmt.Println("  - Removed from skill.json")

	// 4. Remove from skill.lock
	l, err := lockfile.Load("skill.lock")
	if err == nil {
		if _, ok := l.Dependencies[targetKey]; ok {
			delete(l.Dependencies, targetKey)
			if err := l.Save("skill.lock"); err != nil {
				return fmt.Errorf("failed to update skill.lock: %v", err)
			}
			fmt.Println("  - Removed from skill.lock")
		}
	}

	// 5. Remove directory
	// In our system, the directory name is either the alias (the key) 
	// or the repo's base name (if the key is user/repo).
	dirName := targetKey
	if strings.Contains(targetKey, "/") {
		parts := strings.Split(targetKey, "/")
		dirName = parts[len(parts)-1]
	}

	targetDir := filepath.Join(modulesDir, dirName)
	if _, err := os.Stat(targetDir); err == nil {
		if err := os.RemoveAll(targetDir); err != nil {
			return fmt.Errorf("failed to remove directory %s: %v", targetDir, err)
		}
		fmt.Println("  - Deleted directory")
	}

	fmt.Printf("Package %s removed successfully.\n", targetKey)
	return nil
}
