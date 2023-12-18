package utils

import "strings"

func EmbedStringToStringSlice(input string) []string {
	return strings.Split(strings.TrimSuffix(input, "\n"), "\n")
}
