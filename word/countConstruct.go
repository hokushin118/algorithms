package word

import (
	"strings"
)

/*
The function takes in a target string and an array of strings as arguments and returns the number of ways
that the target can be constructed by concatenating elements of the words array.

The element of the array can be used as many times as needed.

m target.length number of characters in the target word
n words.length number of words in the words array

time: O(n * m ^ 2)
space: O(m ^ 2)

Recursion is used for solving the problem.
*/
func CountConstruct(target string, words []string, cache map[string]int) (int, error) {
	if val, ok := cache[target]; ok {
		return val, nil
	}
	if target == "" {
		return 1, nil
	}

	totalCount := 0

	for _, word := range words {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			numberOfWaysForRest, _ := CountConstruct(suffix, words, cache)
			totalCount += numberOfWaysForRest
		}
	}

	cache[target] = totalCount
	return totalCount, nil
}
