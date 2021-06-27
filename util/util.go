package util

import (
	"regexp"
	"strings"
)

func TrimHtml(src string) string {
	re, _ := regexp.Compile("<[^>]+>")
	src = re.ReplaceAllString(src, "")
	re, _ = regexp.Compile(`\[!\[(.*?)\]\((.*?)\)\]`)
	src = re.ReplaceAllString(src, "")
	return src
}

func Between(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}
