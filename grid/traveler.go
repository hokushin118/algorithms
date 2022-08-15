package grid

import (
	"strconv"
)

/*
Say you are a traveller on a 2D grid. You begin in the top-left corner and your goal is to travel to
the bottom-right corner. You may only move down or right.

In how many ways can you travel to the goal on a grid with dimensions m * n?

Write a function that calculates this.

n:                  2, 3, 3
m:                  3, 2, 3

Traveler(2, 3)

| | | |
| | | |

1. right, right, down
2. right, down, right
3. down, right, right

n number of rows
m number of columns

time:  O(n * m)
space: O(n + m)

Recursion is used for solving the problem.
*/
func Traveler(n int, m int, cache map[string]int) int {
	key := strconv.Itoa(n) + "-" + strconv.Itoa(m)

	if _, ok := cache[key]; ok {
		return cache[key]
	}
	if n == 1 && m == 1 {
		return 1
	}
	if n == 0 || m == 0 {
		return 0
	}

	cache[key] = Traveler(n-1, m, cache) + Traveler(n, m-1, cache)

	return cache[key]
}
