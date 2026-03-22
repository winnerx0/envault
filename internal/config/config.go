package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const FileName = "envault.json"

func FindProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		configPath := filepath.Join(dir, FileName)
		if _, err := os.Stat(configPath); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("envault.json not found — run 'envault init' to initialize a project")
		}
		dir = parent
	}
}
