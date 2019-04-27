package palindrome

import (
	"errors"
	"regexp"
	"strings"
)

const alphabets string = "abcdefghijklmnopqrstuvwxyz"

func intToAlphabet(i int) string {
	return string('a' + i)
}

func alphabetToInt(alphabet string) int {
	return strings.Index(alphabets, alphabet)
}

func MakePalindrome(input string) (int, error) {
	match, _ := regexp.MatchString("^[a-z]*$", input)
	if !match {
		return -1, errors.New("can only work with lowercase alphabet")
	}
	steps := 0
	for i := 0; i < len(input)/2; i++ {
		firstLetter, lastLetter := string(input[i]), string(input[len(input)-i-1])
		for firstLetter != lastLetter {
			steps++
			if strings.Compare(firstLetter, lastLetter) == -1 {
				lastLetter = intToAlphabet(alphabetToInt(lastLetter) - 1)
			} else {
				lastLetter = intToAlphabet(alphabetToInt(lastLetter) + 1)
			}
		}
	}
	return steps, nil
}
