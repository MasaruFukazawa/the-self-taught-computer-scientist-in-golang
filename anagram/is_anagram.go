package anagram

import (
	"sort"
	"strings"
)

func sortString(s string) string {

	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

func IsAnagram(s1 string, s2 string) bool {

	s1 = sortString(strings.Replace(strings.ToLower(s1), " ", "", -1))
	s2 = sortString(strings.Replace(strings.ToLower(s2), " ", "", -1))

	return s1 == s2
}
