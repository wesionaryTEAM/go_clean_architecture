package utils

import (
	"flag"
	"os"
	"strings"
)

// IsTestEnv checks if we are running in test environment
func IsTestEnv() bool {
	return strings.HasSuffix(os.Args[0], ".test") || flag.Lookup("test.v") != nil
}
