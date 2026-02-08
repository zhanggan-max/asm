package commands

import (
	"asm/pkg/manifest"
	"fmt"
	"os"
	"path/filepath"
)

// RunInit initializes a new skill project
func RunInit(name string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}

	// Default name to directory name if not provided
	if name == "" {
		name = filepath.Base(cwd)
	}

	s := manifest.Skill{
		Name:        name,
		Version:     "0.1.0",
		Description: "A new agent skill",
		Main:        "index.py", // Assuming python skill by default, can be anything
		Dependencies: make(map[string]string),
	}

	filename := "skill.json"
	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("%s already exists", filename)
	}

	if err := s.Save(filename); err != nil {
		return fmt.Errorf("failed to save manifest: %v", err)
	}

	fmt.Printf("Initialized empty skill repository in %s\n", filepath.Join(cwd, filename))
	return nil
}
