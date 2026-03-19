package main

import (
	"fmt"
	"os"
)

// GetPathSize возвращает размер файла или директории первого уровня
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		return fmt.Sprintf("%dB", info.Size()), nil
	}

	var total int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		entryInfo, err := entry.Info()
		if err != nil {
			return "", err
		}
		total += entryInfo.Size()
	}

	return fmt.Sprintf("%dB", total), nil
}