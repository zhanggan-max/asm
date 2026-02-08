package registry

import (
	"fmt"
	"strings"
)

// ResolvePackage converts a package name to a Git URL.
func ResolvePackage(pkg string) (string, error) {
	// 1. Direct URL
	if strings.HasPrefix(pkg, "http://") || strings.HasPrefix(pkg, "https://") || strings.HasPrefix(pkg, "git@") {
		return pkg, nil
	}

	// 2. Local path
	if strings.HasPrefix(pkg, ".") || strings.HasPrefix(pkg, "/") {
		return pkg, nil
	}
	
	// 3. GitHub shorthand "user/repo"
	parts := strings.Split(pkg, "/")
	if len(parts) == 2 {
		return "https://github.com/" + pkg + ".git", nil
	}

	return "", fmt.Errorf("package '%s' not found in registry (use 'user/repo' or full URL)", pkg)
}
