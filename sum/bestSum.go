package sum

import (
	"errors"
)

/*
The function takes in a target number and an array of numbers as arguments and returns an array
containing the shortest combination of numbers that add up to exactly the target.

If there is a tie for the shortest combinations, the function returns any one of the shortest.

m target sum
n numbers.length

time:  O(n * m ^ 2)
space: O(m ^ 2)

Recursion is used for solving the problem.
*/
func BestSum(target int, numbers []int, cache map[int][]int) ([]int, error) {
	if val, ok := cache[target]; ok {
		return val, nil
	}
	if target == 0 {
		return []int{}, nil
	}
	if target < 0 {
		return nil, errors.New("target can't be less than zero")
	}

	var shortestCombination []int

	for _, num := range numbers {
		remainder := target - num
		remainderCombination, _ := BestSum(remainder, numbers, cache)
		if remainderCombination != nil {
			combination := append([]int{num}, remainderCombination...)
			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}
	}

	if shortestCombination == nil {
		return nil, errors.New("can't find best sum from number combinations")
	}

	cache[target] = shortestCombination
	return shortestCombination, nil
}
