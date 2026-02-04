package y2015

import (
	"fmt"
	"regexp"
	"strings"
)

type Day05 struct{}

// SolvePart1 implements solutions.Solver.
func (d Day05) SolvePart1(input []string) (int, error) {
	vowels := "aeiou"
	forgotedStringRegexp := regexp.MustCompile("(ab|cd|pq|xy)")

	result := 0

	for _, word := range input {
		threeVowels := 0
		twice := false

		for i, letter := range word {
			if strings.ContainsRune(vowels, letter) {
				threeVowels++
			}
			if i > 0 {
				if letter == rune(word[i-1]) {
					twice = true
				}
			}
		}

		if threeVowels >= 3 && twice && !forgotedStringRegexp.MatchString(word) {
			result += 1
		}

	}
	return result, nil
}

// SolvePart2 implements solutions.Solver.
func (d Day05) SolvePart2(input []string) (int, error) {
	result := 0
	for _, word := range input {
		twice := false
		repeats := false
		for i, letter := range word {
			if i+1 >= len(word) {
				continue
			}
			letters := fmt.Sprintf("%v%v", string(letter), string(word[i+1]))
			lettersCount := strings.Count(word, letters)
			if lettersCount >= 2 {
				twice = true
			}
			if i+2 >= len(word) {
				continue
			}
			if letter == rune(word[i+2]) {
				repeats = true
			}
		}

		if twice && repeats {
			result++
		}
	}

	return result, nil
}
