package utils

import (
	"os"
)

// IsCli checks if app is running in cli mode
func IsCli() bool {
	if len(os.Args) > 1 {
		commandLine := os.Args[1]
		if commandLine != "" {
			return true
		}
	}
	return false
}
