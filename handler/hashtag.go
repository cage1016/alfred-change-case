package handler

import (
	"regexp"
	"strings"
)

func HashTag(s string) string {
	array := regexp.MustCompile(" +").Split(s, -1)
	for i := 0; i < len(array); i++ {
		if len(array[i]) > 0 && array[i][0:1] != "#" {
			array[i] = "#" + strings.ToLower(array[i])
		}
	}
	return strings.Join(array, " ")
}
