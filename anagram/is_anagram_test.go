package anagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {

	var s1 string = "Emperor Octavian"
	var s2 string = "Captain over Rome"

	if !IsAnagram(s1, s2) {
		t.Errorf("not anagram")
	}
}
