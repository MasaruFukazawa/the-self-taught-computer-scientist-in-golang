package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	var s1 string = "Hannah"

	if !IsPalindrome(s1) {
		t.Errorf("not palindrome")
	}
}
