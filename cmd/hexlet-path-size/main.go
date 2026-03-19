package main

import (
	"fmt"
	"log"
	"os"
)

func GetSize(path string) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64
	for _, entry := range dirEntries {
		ei, err := entry.Info()
		if err != nil {
			return 0, err
		}
		total += ei.Size()
	}

	return total, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a path")
		return
	}

	path := os.Args[1]
	size, err := GetSize(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%dB\t%s\n", size, path)
}