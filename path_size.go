package code

import (
	"os"
	"path/filepath"
)

func GetSize(path string, all bool, recursive bool) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	var total int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		name := entry.Name()

		if !all && len(name) > 0 && name[0] == '.' {
			continue
		}

		fullPath := filepath.Join(path, name)

		if entry.IsDir() {
			if recursive {
				size, err := GetSize(fullPath, all, recursive)
				if err != nil {
					return 0, err
				}
				total += size
			}
		} else {
			info, err := entry.Info()
			if err != nil {
				return 0, err
			}
			total += info.Size()
		}
	}

	return total, nil
}