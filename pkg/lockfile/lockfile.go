package lockfile

import (
	"encoding/json"
	"os"
)

// PackageLock represents a single locked dependency
type PackageLock struct {
	Version string `json:"version"`
	URL     string `json:"url"`
	Commit  string `json:"commit"`
}

// Lockfile represents the structure of skill.lock
type Lockfile struct {
	Version      int                    `json:"lockfileVersion"`
	Dependencies map[string]PackageLock `json:"dependencies"`
}

// Load reads the lockfile from path
func Load(path string) (*Lockfile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Lockfile{
				Version:      1,
				Dependencies: make(map[string]PackageLock),
			}, nil
		}
		return nil, err
	}
	var l Lockfile
	err = json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}
	if l.Dependencies == nil {
		l.Dependencies = make(map[string]PackageLock)
	}
	return &l, nil
}

// Save writes the lockfile to path
func (l *Lockfile) Save(path string) error {
	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
