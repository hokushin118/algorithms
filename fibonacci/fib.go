package fibonacci

/*
The function takes in a number as an argument and returns the n-th number of the Fibonacci sequence.

The 1st and 2nd number of the sequence is 1.

To generate the next number of the sequence, we sum the previous two.

n:      1, 2, 3, 4, 5, 6,  7,  8,  9, ...
fib(n): 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

n number

time:  O(n)
space: O(n)

Recursion is used for solving the problem.
*/
func Fib(n int, cache map[int]int) int {
	if i, ok := cache[n]; ok {
		return i
	}
	if n <= 2 {
		return 1
	}

	cache[n] = Fib(n-1, cache) + Fib(n-2, cache)
	return cache[n]
}
