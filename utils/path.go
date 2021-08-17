package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

// AbsolutePath get execute binary path
func AbsolutePath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}

	return filepath.Dir(path), nil
}

// Exists whether file or directory exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// IsDir return bool
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile return tool
func IsFile(path string) bool {
	return !IsDir(path)
}
