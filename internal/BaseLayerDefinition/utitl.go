package base

import (
	"strings"
	"unicode"
)


func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func removeAndCamelCase(atrName string, chars ...string) string {
	name := atrName
	for _, char := range chars {
		if strings.Contains(atrName, char) {
			index := strings.Index(atrName, char)
			name = strings.Replace(name, char, "", -1)
			name = name[:index] + strings.ToUpper(string(name[index])) + name[index+1:]
		}
	}
	return name
}
