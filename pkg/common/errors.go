package common

import (
	"fmt"
	"os"
)

// ExitCode convention.
const (
	ExitSuccess = 0
	ExitFailure = 1
)

// Fatal prints msg to stderr and exits with code 1.
func Fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(ExitFailure)
}

// Warn prints msg to stderr.
func Warn(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}
