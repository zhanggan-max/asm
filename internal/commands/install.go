package commands

import (
	"asm/internal/utils"
	"asm/pkg/lockfile"
	"asm/pkg/manifest"
	"asm/pkg/registry"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const modulesDir = ".asm_modules"

// RunInstall handles the install command
func RunInstall(pkgSpec string) error {
	if pkgSpec == "" {
		return installAll()
	}
	return installWithDependencies(pkgSpec, true)
}

func installAll() error {
	lock, err := lockfile.Load("skill.lock")
	if err == nil && len(lock.Dependencies) > 0 {
		fmt.Println("Installing from skill.lock...")
		return installFromLock(lock)
	}

	m, err := manifest.Load("skill.json")
	if err != nil {
		return fmt.Errorf("failed to load skill.json: %v", err)
	}

	fmt.Println("Installing from skill.json...")
	for name, spec := range m.Dependencies {
		pkgSpec := name
		if spec != "*" && spec != "" && !strings.Contains(spec, "/") {
			pkgSpec = name + "@" + spec
		} else if strings.Contains(spec, "/") {
			pkgSpec = spec
		}
		
		if err := installWithDependencies(pkgSpec, false); err != nil {
			return err
		}
	}
	return nil
}

func installWithDependencies(pkgSpec string, updateManifest bool) error {
	// Parse alias: pkg:alias
	alias := ""
	actualSpec := pkgSpec
	if idx := strings.LastIndex(pkgSpec, ":"); idx > 0 && idx < len(pkgSpec)-1 {
		// Verify it's not a git URL with colon (like git@github.com:user/repo)
		if !strings.HasPrefix(pkgSpec, "git@") {
			actualSpec = pkgSpec[:idx]
			alias = pkgSpec[idx+1:]
		}
	}

	installName, targetDir, err := installSingleRaw(actualSpec, alias)
	if err != nil {
		return err
	}

	if updateManifest {
		m, _ := manifest.Load("skill.json")
		if m != nil {
			parts := strings.Split(actualSpec, "@")
			val := "*"
			if len(parts) > 1 {
				val = parts[1]
			} else if strings.Contains(actualSpec, "://") {
				val = actualSpec
			}
			
			if m.Dependencies == nil {
				m.Dependencies = make(map[string]string)
			}
			m.Dependencies[installName] = val
			m.Save("skill.json")
		}
	}

	// Recursion
	subManifestPath := filepath.Join(targetDir, "skill.json")
	if _, err := os.Stat(subManifestPath); err == nil {
		subM, err := manifest.Load(subManifestPath)
		if err == nil && len(subM.Dependencies) > 0 {
			fmt.Printf("  -> %s has %d dependencies, resolving...\n", installName, len(subM.Dependencies))
			for subName, subSpec := range subM.Dependencies {
				subPkgSpec := subName
				if subSpec != "*" && subSpec != "" && !strings.Contains(subSpec, "/") {
					subPkgSpec = subName + "@" + subSpec
				} else if strings.Contains(subSpec, "/") {
					subPkgSpec = subSpec
				}

				if _, err := os.Stat(filepath.Join(modulesDir, subName)); err == nil {
					continue
				}

				if err := installWithDependencies(subPkgSpec, false); err != nil {
					return fmt.Errorf("failed to install dependency %s of %s: %v", subName, installName, err)
				}
			}
		}
	}

	return nil
}

func installSingleRaw(pkgSpec, alias string) (string, string, error) {
	parts := strings.Split(pkgSpec, "@")
	pkgName := parts[0]
	version := ""
	if len(parts) > 1 {
		version = parts[1]
	}

	url, err := registry.ResolvePackage(pkgName)
	if err != nil {
		return "", "", err
	}

	installName := alias
	if installName == "" {
		installName = utils.GetRepoName(url)
	}
	targetDir := filepath.Join(modulesDir, installName)

	if _, err := os.Stat(targetDir); err == nil {
		return installName, targetDir, nil
	}

	fmt.Printf("Fetching %s...\n", url)
	if err := os.MkdirAll(modulesDir, 0755); err != nil {
		return "", "", err
	}

	if err := utils.CloneRepo(url, targetDir, version); err != nil {
		return "", "", err
	}

	l, _ := lockfile.Load("skill.lock")
	hash, _ := utils.GetHeadHash(targetDir)
	l.Dependencies[installName] = lockfile.PackageLock{
		Version: version,
		URL:     url,
		Commit:  hash,
	}
	l.Save("skill.lock")

	return installName, targetDir, nil
}

func installFromLock(lock *lockfile.Lockfile) error {
	if err := os.MkdirAll(modulesDir, 0755); err != nil {
		return err
	}
	for name, pkg := range lock.Dependencies {
		targetDir := filepath.Join(modulesDir, name)
		if _, err := os.Stat(targetDir); err == nil {
			continue
		}
		fmt.Printf("Restoring %s@%s...\n", name, pkg.Commit[:7])
		if err := utils.CloneRepo(pkg.URL, targetDir, ""); err != nil {
			return err
		}
	}
	return nil
}