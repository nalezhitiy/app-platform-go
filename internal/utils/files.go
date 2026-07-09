package utils

import (
	"os"
	"path/filepath"
)

// FindFilesByExtension scans the directory for files with matching extensions
func FindFilesByExtension(dir string, extensions ...string) ([]string, error) {
	var matchedFiles []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := filepath.Ext(file.Name())
		for _, targetExt := range extensions {
			if ext == targetExt {
				fullPath := filepath.Join(dir, file.Name())
				matchedFiles = append(matchedFiles, fullPath)
				break
			}
		}
	}

	return matchedFiles, nil
}
