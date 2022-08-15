package word

import (
	"strings"
)

/*
The function takes in a target string and an array of strings as arguments and returns a boolean indicating
whether or not the target can be constructed by concatenating elements of the array.

The element of the array can be used as many times as needed.

m target.length number of characters in the target word
n words.length number of words in the words array

time: O(n * m^2)
space: O(m^2)

Recursion is used for solving the problem.
*/
func CanConstruct(target string, words []string, cache map[string]bool) (bool, error) {
	if val, ok := cache[target]; ok {
		return val, nil
	}
	if target == "" {
		return true, nil
	}

	for _, word := range words {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			isCanConstruct, _ := CanConstruct(suffix, words, cache)
			if isCanConstruct {
				cache[target] = true
				return true, nil
			}
		}
	}

	cache[target] = false
	return false, nil
}
