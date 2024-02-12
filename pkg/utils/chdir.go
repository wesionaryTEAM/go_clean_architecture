package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func ChDir(paths ...string) error {
	newPath := "../../.."

	if len(paths) > 0 {
		newPath = filepath.Join(paths...)
	}

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Compute the new directory path
	newDir := filepath.Join(currentDir, newPath)

	// Change to the new directory
	if err := os.Chdir(newDir); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", newDir, err)
	}

	return nil
}
