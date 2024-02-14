package main

import (
	"strings"
	"unicode/utf8"
)

func NormalizeStringLength(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}
