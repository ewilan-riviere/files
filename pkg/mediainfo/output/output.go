package output

import (
	"regexp"
	"strings"

	"github.com/ewilan-riviere/files/pkg/mediainfo/types"
)

func Make(output string) []types.MediainfoOutput {
	list := []types.MediainfoOutput{}

	lines := strings.Split(output, "\n")
	item := types.MediainfoOutput{}
	for _, line := range lines {
		regex := regexp.MustCompile(`\s{2,}`)
		line := regex.ReplaceAllString(line, " ")
		line = strings.TrimSpace(line)

		if line == "" {
			if item.Category != "" {
				list = append(list, item)
			}
			item = types.MediainfoOutput{}
		} else {
			parseMediainfoOutput(&item, line)
		}
	}

	return list
}

func parseMediainfoOutput(item *types.MediainfoOutput, line string) {
	if !strings.Contains(line, ":") {
		item.Category = line
	} else {
		key := ""
		value := ""

		parts := strings.Split(line, " : ")
		key = parts[0]
		value = parts[1]

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if key == "" && value == "" {
			return
		}

		newItem := types.MediainfoOutputItem{
			Key:   key,
			Value: value,
		}
		item.Items = append(item.Items, newItem)
	}
}
