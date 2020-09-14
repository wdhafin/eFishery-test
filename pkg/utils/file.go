package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kennygrant/sanitize"
)

// IsFileExist returns a boolean indicating whether file is exist or not
func IsFileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// ExtractFileName is
func ExtractFileName(fn string) string {
	return sanitize.BaseName(strings.TrimSuffix(fn, filepath.Ext(fn)))
}

// ExtractFileExt is
func ExtractFileExt(fn string) string {
	return filepath.Ext(fn)
}
