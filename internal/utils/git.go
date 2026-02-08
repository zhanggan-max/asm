package utils

import (
	"os"
	"os/exec"
	"strings"
)

// CloneRepo clones a git repository to a destination with a specific reference (tag/branch)
// If ref is empty, it clones the default branch.
func CloneRepo(url, dest, ref string) error {
	args := []string{"clone", "--depth", "1"}
	if ref != "" {
		args = append(args, "--branch", ref)
	}
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
