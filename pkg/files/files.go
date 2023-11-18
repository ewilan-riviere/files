package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func Make(path string) []string {
	var files []string
	var list []string

	realRoot, err := filepath.EvalSymlinks(path)
	if err != nil {
		fmt.Printf("Error resolving symbolic link: %v\n", err)
		return files
	}

	err = filepath.Walk(realRoot, visit(&files))
	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", realRoot, err)
	}

	// Append the collected file paths to the list
	list = append(list, files...)
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
