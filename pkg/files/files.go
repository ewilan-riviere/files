package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func Make(path string) []string {
	var list []string
	var files []string
	disallowedExtensions := []string{".DS_Store"}

	err := filepath.Walk(path, visit(&files))
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", path, err)
	}

	// Print the collected file paths
	for _, file := range files {
		extension := filepath.Ext(file)
		valid := true
		for _, disallowedExtension := range disallowedExtensions {
			if extension == disallowedExtension {
				valid = false
			}
		}

		if valid {
			list = append(list, file)
		}
	}

	return list
}

func visit(files *[]string) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err) // can't walk here,
			return nil       // ignore this error.
		}
		if !info.IsDir() {
			*files = append(*files, path)
		}
		return nil
	}
}
