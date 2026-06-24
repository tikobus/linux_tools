package common

import (
	"path/filepath"
	"strings"
)

// NormalizePath cleans a path and converts slashes to the native separator.
func NormalizePath(p string) string {
	p = filepath.Clean(p)
	return filepath.FromSlash(p)
}

// IsAbs reports whether p is an absolute path on any supported platform.
func IsAbs(p string) bool {
	return filepath.IsAbs(p)
}

// Join joins path elements using the native separator.
func Join(elem ...string) string {
	return filepath.Join(elem...)
}

// Base returns the last element of p.
func Base(p string) string {
	return filepath.Base(p)
}

// Dir returns all but the last element of p.
func Dir(p string) string {
	return filepath.Dir(p)
}

// SplitExt splits p into base and extension.
func SplitExt(p string) (string, string) {
	ext := filepath.Ext(p)
	base := strings.TrimSuffix(p, ext)
	return base, ext
}
