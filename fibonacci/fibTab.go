package fibonacci

/*
The function takes in a number as an argument and returns the n-th number of the Fibonacci sequence.

The 0th number of the sequence is 0.
The 1st number of the sequence is 1.

To generate the next number of the sequence, we sum the previous two.

n:         0, 1, 2, 3, 4, 5, 6,  7,  8,  9, ...
fibTab(n): 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...

n number

time:  O(n)
space: O(n)

Tabulation is used for solving the problem.
*/
func FibTab(n int) int {
	// 1. visualize the problem as a table
	// 2. size the table based on the inputs
	// 3. initialize the table with defaults values
	table := make([]int, n+2)

	// 4. seed the trivial answer into the table
	table[1] = 1

	// 5. iterate through the table
	for i := 0; i < n; i++ {
		// 6. fill further positions based on the current position
		currentPosition := table[i]
		table[i+1] += currentPosition
		table[i+2] += currentPosition
	}

	return table[n]
}
