package periscope

import (
	"github.com/ewilan-riviere/scanner/pkg/files"
	"github.com/ewilan-riviere/scanner/pkg/info"
)

type FileMeta struct {
	Path string
	Info info.Info
	// MediainfoItem types.MediainfoItem
}

func Make(path string) []FileMeta {
	list := []FileMeta{}
	items := files.Make(path)

	for _, item := range items {
		file := info.Make(item)
		// metadata := mediainfo.Make(path)

		list = append(list, FileMeta{
			Path: item,
			Info: file,
			// MediainfoItem: metadata,
		})
	}

	return list
}
