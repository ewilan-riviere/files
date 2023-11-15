package scan

// import (
// 	"fmt"
// 	"os/exec"
// 	"reflect"
// 	"regexp"
// 	"strings"
// )

// func Scan(path string) {
// 	output := execMediainfo(path)
// 	list := []MediainfoOutput{}

// 	lines := strings.Split(output, "\n")
// 	item := MediainfoOutput{}
// 	for _, line := range lines {
// 		regex := regexp.MustCompile(`\s{2,}`)
// 		line := regex.ReplaceAllString(line, " ")
// 		line = strings.TrimSpace(line)

// 		if line == "" {
// 			if item.Category != "" {
// 				list = append(list, item)
// 			}
// 			item = MediainfoOutput{}
// 		} else {
// 			parseMediainfoOutput(&item, line)
// 		}
// 	}

// 	mediainfoItem := MediainfoItem{}
// 	for _, item := range list {
// 		if item.Category == "General" {
// 			convertGeneral(&mediainfoItem, item.Items)
// 		}
// 		if strings.Contains(item.Category, "Video") {
// 			v := MediainfoItemVideo{}
// 			convertVideo(&v, item.Items)
// 			mediainfoItem.Videos = append(mediainfoItem.Videos, v)
// 		}
// 		if strings.Contains(item.Category, "Audio") {
// 			a := MediainfoItemAudio{}
// 			convertAudio(&a, item.Items)
// 			mediainfoItem.Audios = append(mediainfoItem.Audios, a)
// 		}
// 		if strings.Contains(item.Category, "Text") {
// 			t := MediainfoItemText{}
// 			convertText(&t, item.Items)
// 			mediainfoItem.Texts = append(mediainfoItem.Texts, t)
// 		}
// 		if strings.Contains(item.Category, "Menu") {
// 			if mediainfoItem.Menu == nil {
// 				mediainfoItem.Menu = make(map[string]string)
// 			}
// 			for _, item := range item.Items {
// 				mediainfoItem.Menu[item.Key] = item.Value
// 			}
// 		}
// 	}

// 	fmt.Println()
// 	printStruct(mediainfoItem)
// 	fmt.Println()
// }

// func convertGeneral(item *MediainfoItem, items []MediainfoOutputItem) {
// 	item.UniqueID = filterItems("Unique ID", items)
// 	item.CompleteName = filterItems("Complete name", items)
// 	item.Format = filterItems("Format", items)
// 	item.FormatVersion = filterItems("Format version", items)
// 	item.FileSize = filterItems("File size", items)
// 	item.Duration = filterItems("Duration", items)
// 	item.OverallBitRate = filterItems("Overall bit rate", items)
// 	item.FrameRate = filterItems("Frame rate", items)
// 	item.MovieName = filterItems("Movie name", items)
// 	item.EncodedDate = filterItems("Encoded date", items)
// 	item.WritingApplication = filterItems("Writing application", items)
// 	item.WritingLibrary = filterItems("Writing library", items)
// 	item.Cover = filterItems("Cover", items)
// 	item.Attachments = filterItems("Attachments", items)
// }
