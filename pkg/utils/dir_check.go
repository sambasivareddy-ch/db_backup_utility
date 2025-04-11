package utils

import (
	"errors"
	"os"
	"path/filepath"
)

// Checks whether the given Dir path exists or not.
// If not throws error
func IsDirExists(path interface{}) error {
	dir := path.(string)

	if dir == "" {
		return errors.New("backup directory required")
	}

	if _, err := os.Stat(filepath.Clean(dir)); os.IsNotExist(err) {
		return errors.New("invalid backup directory")
	}

	return nil
}
