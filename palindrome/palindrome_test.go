package palindrome

import "testing"

var testCases = []struct {
	input        string
	expectedStep int
}{
	{"abc", 2},
	{"abcde", 6},
	{"aaa", 0},
	{"0930aaa", -1},
	{"AAA", -1},
}

func TestMakePalindrome(t *testing.T) {
	for _, tc := range testCases {
		actualStep, _ := MakePalindrome(tc.input)
		if actualStep != tc.expectedStep {
			t.Errorf("%q: actual (%q), expected (%q)", tc.input, actualStep, tc.expectedStep)
		}
	}
}
