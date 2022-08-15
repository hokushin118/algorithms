package sum

import (
	"errors"
)

/*
The function takes in a target number and an array of numbers as arguments and returns an array
containing any combination of elements that add up to exactly the target.

If there is no combination that adds up to exactly the target, then return null.

If there are multiple combinations possible, any single one is returned.

m target sum
n numbers.length

time: O(n * m ^ 2)
space: O(m ^ 2)

Recursion is used for solving the problem.
*/
func HowSum(target int, numbers []int, cache map[int][]int) ([]int, error) {
	if val, ok := cache[target]; ok {
		return val, nil
	}
	if target == 0 {
		return []int{}, nil
	}
	if target < 0 {
		return nil, errors.New("target can't be less than zero")
	}

	for _, num := range numbers {
		remainder := target - num
		remainderCombination, _ := HowSum(remainder, numbers, cache)
		if remainderCombination != nil {
			cache[target] = append([]int{num}, remainderCombination...)
			return cache[target], nil
		}
	}

	cache[target] = nil
	return nil, nil
}
