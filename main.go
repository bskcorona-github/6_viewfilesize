package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path>")
		return
	}

	path := os.Args[1]

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if fileInfo.IsDir() {
		totalSize, err := getFolderSize(path)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Total size of %s: %d bytes\n", path, totalSize)
	} else {
		fmt.Printf("Size of %s: %d bytes\n", path, fileInfo.Size())
	}
}

func getFolderSize(folderPath string) (int64, error) {
	var totalSize int64

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return totalSize, nil
}
