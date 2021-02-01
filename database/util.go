package database

import (
	"strings"
)


func concat(delim string, values []string) string {
	return strings.Join(values, delim)
}

func replace(s string, delim string, times int, values []string) string {
	for _, value := range values {
		s = strings.Replace(s, delim, value, times)
	}
	return s
}

func replicateToStr(letter string, delim string, times int) string {
	letters := replicateToArr(letter, times)
	return strings.Join(letters, delim)
}

func replicateToArr(letter string, times int) []string {
	letters := make([]string, times)

	for i := range letters {
		letters[i] = letter
	}

	return letters
}

func toInterfaceSplice(arr []string) []interface{} {
	interfaces := make([]interface{}, len(arr))

	for i := range interfaces {
		interfaces[i] = arr[i]
	}

	return interfaces
}