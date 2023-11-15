package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ewilan-riviere/files/pkg/files"
	"github.com/ewilan-riviere/files/pkg/mediainfo"
	"github.com/ewilan-riviere/files/pkg/printer"
)

func main() {
	args := os.Args[1:]
	path := args[0]

	fmt.Println(path)

	startTime := time.Now()
	items := files.Make(path)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
	printer.Print(items)
	printer.ToJSON(printer.Params{
		V:    items,
		Path: "output/files.json",
	})

	pathToScan := "/Volumes/data/music/librairies/podcasts/F.Kermesse/FK.1_Le.Gore.Philippe.Bouvard.de.la.mort.mp3"
	mp3 := mediainfo.Make(pathToScan)
	pathToScan = "/Volumes/data/video/movies/Action/Heroes/Batman/Batman/Batman.1989.MULTi.1080p.x264.mkv"
	mkv := mediainfo.Make(pathToScan)

	printer.ToJSON(printer.Params{
		V:    mp3,
		Path: "output/mp3.json",
	})
	printer.ToJSON(printer.Params{
		V:    mkv,
		Path: "output/mkv.json",
	})
}
