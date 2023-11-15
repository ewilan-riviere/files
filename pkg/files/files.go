package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func Make(path string) []string {
	var files []string
	var list []string

	disallowedExtensions := []string{".DS_Store"}

	realRoot, err := filepath.EvalSymlinks(path)
	if err != nil {
		fmt.Printf("Error resolving symbolic link: %v\n", err)
		return files
	}

	err = filepath.Walk(realRoot, visit(&files))
	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", realRoot, err)
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
			fmt.Println(err)
			return nil
		}
		if info.IsDir() {
			//
		} else {
			*files = append(*files, path)
		}
		return nil
	}
}
