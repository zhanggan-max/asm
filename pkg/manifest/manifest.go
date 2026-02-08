package manifest

import (
	"encoding/json"
	"os"
)

// Skill represents the structure of skill.json
type Skill struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Description  string            `json:"description"`
	Main         string            `json:"main,omitempty"` // Entry point file
	Dependencies map[string]string `json:"dependencies,omitempty"`
}

// Save writes the skill manifest to a specific path
func (s *Skill) Save(path string) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Load reads the skill manifest from a specific path
func Load(path string) (*Skill, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s Skill
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
