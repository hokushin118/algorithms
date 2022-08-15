package sum

/*
The function takes in a target number and an array of numbers as arguments and returns a boolean indicating
whether or not it is possible to generate the target using numbers from the array.

The element of the array can be used as many times as needed. All input numbers are non negative.

m target sum
n numbers.length

time: O(n * m)
space: O(m)

Tabulation is used for solving the problem.
*/
func CanSumTab(target int, numbers []int) (bool, error) {
	// 1. visualize the problem as a table
	// 2. size the table based on the inputs
	// 3. initialize the table with defaults values
	table := make([]bool, target+1)

	// 4. seed the trivial answer into the table
	table[0] = true

	// 5. iterate through the table
	for i := 0; i < target; i++ {
		// 6. fill further positions based on the current position
		currentPosition := table[i]
		if currentPosition {
			for _, num := range numbers {
				nextPosition := i + num
				if nextPosition <= target {
					table[nextPosition] = true
				}
			}
		}
	}

	return table[target], nil
}
