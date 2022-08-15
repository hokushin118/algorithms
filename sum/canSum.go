package sum

import (
	"errors"
)

/*
The function takes in a target number and an array of numbers as arguments and returns a boolean indicating
whether or not it is possible to generate the target using numbers from the array.

The element of the array can be used as many times as needed. All input numbers are non negative.

m target sum
n numbers.length

time: O(n * m)
space: O(m)

Recursion is used for solving the problem.
*/
func CanSum(target int, numbers []int, cache map[int]bool) (bool, error) {
	if val, ok := cache[target]; ok {
		return val, nil
	}
	if target == 0 {
		return true, nil
	}
	if target < 0 {
		return false, errors.New("target can't be less than zero")
	}

	for _, num := range numbers {
		remainder := target - num
		remainderCombination, _ := CanSum(remainder, numbers, cache)
		if remainderCombination {
			cache[target] = true
			return true, nil
		}
	}

	cache[target] = false
	return false, nil
}
