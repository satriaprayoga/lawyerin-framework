package utils

import (
	"regexp"
	"strings"
)

func FormatSearch(search string) string {
	str := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(search, " ")
	rstr := strings.ReplaceAll(str, " ", "&")
	return rstr
}
