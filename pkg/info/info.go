package info

import (
	"os"
	"time"
)

type Info struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

func Make(path string) Info {
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if info.IsDir() {
		panic("Path is a directory")
	}

	return Info{
		Name:    info.Name(),
		Size:    info.Size(),
		Mode:    info.Mode(),
		ModTime: info.ModTime(),
		IsDir:   info.IsDir(),
	}
}
