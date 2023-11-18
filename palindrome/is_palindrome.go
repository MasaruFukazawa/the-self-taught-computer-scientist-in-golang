package palindrome

import (
	"strings"
)

func reverseString(s1 string) string {

	var runes []rune = []rune(s1)

	for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}


func IsPalindrome(s1 string) bool {

	s1 = strings.Replace(strings.ToLower(s1), " ", "", -1)

	var s2 string = reverseString(s1)

	return s1 == s2
}
