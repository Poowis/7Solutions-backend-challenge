package utils

import (
	"strings"
	"unicode"
)

func Summary(s string) map[string]int {
	m := make(map[string]int)
	s = strings.ToLower(s)
	i := 0
	for j, c := range s {
		if isLimiter(c) {
			if i < j {
				m[s[i:j]]++
			}
			i = j + 1
		}
	}
	if i < len(s) {
		m[s[i:]]++
	}
	return m
}

// using strings.Fields version
func Split(s string) map[string]int {
	s = strings.ToLower(s)
	arr := strings.FieldsFunc(s, isLimiter)
	m := make(map[string]int)
	for _, str := range arr {
		m[str]++
	}
	return m
}

func isLimiter(r rune) bool {
	if unicode.IsSpace(r) || r == '.' || r == ',' {
		return true
	}
	return false
}
