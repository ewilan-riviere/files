package mediainfo

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ewilan-riviere/scanner/pkg/mediainfo/convert"
	"github.com/ewilan-riviere/scanner/pkg/mediainfo/output"
	"github.com/ewilan-riviere/scanner/pkg/mediainfo/types"
)

func Make(path string) types.MediainfoItem {
	op := execMediainfo(path)
	list := output.Make(op)
	mediainfoItem := convertList(&list)

	return mediainfoItem
}

func convertList(list *[]types.MediainfoOutput) types.MediainfoItem {
	mediainfoItem := types.MediainfoItem{}
	for _, item := range *list {
		if item.Category == "General" {
			convert.ConvertGeneral(&mediainfoItem, item.Items)
		}
		if strings.Contains(item.Category, "Video") {
			v := types.MediainfoItemVideo{}
			convert.ConvertVideo(&v, item.Items)
			mediainfoItem.Videos = append(mediainfoItem.Videos, v)
		}
		if strings.Contains(item.Category, "Audio") {
			a := types.MediainfoItemAudio{}
			convert.ConvertAudio(&a, item.Items)
			mediainfoItem.Audios = append(mediainfoItem.Audios, a)
		}
		if strings.Contains(item.Category, "Text") {
			t := types.MediainfoItemText{}
			convert.ConvertText(&t, item.Items)
			mediainfoItem.Texts = append(mediainfoItem.Texts, t)
		}
		if strings.Contains(item.Category, "Image") {
			i := types.MediainfoItemImage{}
			convert.ConvertImage(&i, item.Items)
			mediainfoItem.Images = append(mediainfoItem.Images, i)
		}
		if strings.Contains(item.Category, "Menu") {
			if mediainfoItem.Menu == nil {
				mediainfoItem.Menu = make(map[string]string)
			}
			for _, item := range item.Items {
				mediainfoItem.Menu[item.Key] = item.Value
			}
		}
	}

	return mediainfoItem
}

func execMediainfo(path string) string {
	if !checkIfMediainfoExists("mediainfo") {
		panic("mediainfo command not found")
	}

	command := "mediainfo"
	args := []string{path}
	cmd := exec.Command(command, args...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(stdout)
}

func checkIfMediainfoExists(command string) bool {
	cmd := exec.Command("which", command)
	err := cmd.Run()

	return err == nil
}
