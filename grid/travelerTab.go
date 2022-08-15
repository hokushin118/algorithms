package grid

/*
Say you are a traveller on a 2D grid. You begin in the top-left corner and your goal is to travel to
the bottom-right corner. You may only move down or right.

In how many ways can you travel to the goal on a grid with dimensions m * n?

Write a function that calculates this.

n number of rows
m number of columns

time:  O(n * m)
space: O(n + m)

Tabulation is used for solving the problem.
*/
func TravelerTab(n int, m int) int {
	// 1. visualize the problem as a table
	// 2. size the table based on the inputs
	// 3. initialize the table with defaults values
	table := make([][]int, n+2)
	for i := range table {
		table[i] = make([]int, m+2)
	}

	// 4. seed the trivial answer into the table
	table[1][1] = 1

	// 5. iterate through the table
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			// 6. fill further positions based on the current position
			currentPosition := table[i][j]
			nextColumn := j + 1
			if nextColumn <= m {
				table[i][nextColumn] += currentPosition
			}
			nextRow := i + 1
			if nextRow <= n {
				table[nextRow][j] += currentPosition
			}
		}
	}

	return table[n][m]
}
