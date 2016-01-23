package gotitle

import (
	"strings"

	"github.com/itang/gotang"
)

const st = "<title>"
const et = "</title>"

func GetTitleFromUrl(url string) (string, error) {
	content, err := gotang.HttpGetAsString(url)
	if err != nil {
		return "", err
	}

	pos := strings.Index(content, st)
	if pos > 0 {
		end := strings.Index(content, et)
		if end > 0 {
			start := pos + len(st)
			return strings.TrimSpace(content[start:end]), nil
		}
	}
	return "", nil
}
