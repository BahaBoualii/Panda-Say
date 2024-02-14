package main

import "unicode/utf8"

func CalculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}
	return w
}
