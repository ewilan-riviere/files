package main

import (
	"testing"

	"github.com/ewilan-riviere/scanner/pkg/files"
	"github.com/ewilan-riviere/scanner/pkg/info"
	"github.com/ewilan-riviere/scanner/pkg/mediainfo"
	"github.com/ewilan-riviere/scanner/pkg/periscope"
	"github.com/ewilan-riviere/scanner/pkg/printer"
)

func TestWebhook(t *testing.T) {
	items := files.Make("./")
	printer.Print(items)
	printer.ToJSON(printer.Params{
		Data: items,
		Path: "./output/test-files.json",
	})

	metadata := mediainfo.Make("./test/media/test.mp3")
	printer.Print(metadata)
	printer.ToJSON(printer.Params{
		Data: metadata,
		Path: "./output/test-metadata.json",
	})

	info := info.Make("./test/media/test.mp3")
	printer.Print(info)
	printer.ToJSON(printer.Params{
		Data: info,
		Path: "./output/test-info.json",
	})

	periscope := periscope.Make("./")
	printer.ToJSON(printer.Params{
		Data: periscope,
		Path: "./output/test-periscope.json",
	})
}
