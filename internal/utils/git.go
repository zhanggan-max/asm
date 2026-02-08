package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// CacheDir returns the global cache directory
func CacheDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".asm", "cache")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

// getCachePath returns the path to the cached bare repo for a given URL
func getCachePath(url string) (string, error) {
	cacheRoot, err := CacheDir()
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256([]byte(url))
	hashStr := hex.EncodeToString(hash[:])
	return filepath.Join(cacheRoot, hashStr), nil
}

// updateCache ensures the cached repo exists and is up-to-date
func updateCache(url, cachePath string) error {
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		// Clone as bare mirror
		fmt.Printf("[Cache] Miss. Cloning %s to cache...\n", url)
		cmd := exec.Command("git", "clone", "--mirror", url, cachePath)
		// Suppress output unless error
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("git clone mirror failed: %s", string(out))
		}
	} else {
		// Update existing cache
		// fmt.Printf("[Cache] Hit. Updating %s...\n", url)
		cmd := exec.Command("git", "remote", "update")
		cmd.Dir = cachePath
		if _, err := cmd.CombinedOutput(); err != nil {
			// Warn but don't fail, maybe offline
			fmt.Printf("Warning: failed to update cache: %v\n", err)
		}
	}
	return nil
}

// CloneRepo clones a git repository to a destination using the global cache
func CloneRepo(url, dest, ref string) error {
	// 1. Try to use cache
	cachePath, err := getCachePath(url)
	if err == nil {
		if err := updateCache(url, cachePath); err == nil {
			// Use cache path as source
			url = cachePath
		}
	}

	// 2. Clone from source (cache or remote)
	args := []string{"clone"}
	if ref != "" {
		args = append(args, "--branch", ref)
	}
	// Note: We can't use --depth 1 easily when cloning from a local mirror if we need a specific commit not at tip
	// But usually it's fine. For now, let's keep it simple.
	// If cloning from local cache, git handles it efficiently.
	args = append(args, url, dest)

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// GetHeadHash returns the current HEAD commit hash of a git repo
func GetHeadHash(repoDir string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = repoDir
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// GetRepoName extracts the repo name from a URL
func GetRepoName(url string) string {
	parts := strings.Split(url, "/")
	last := parts[len(parts)-1]
	return strings.TrimSuffix(last, ".git")
}
