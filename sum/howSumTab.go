package sum

/*
The function takes in a target number and an array of numbers as arguments and returns an array
containing any combination of elements that add up to exactly the target.

If there is no combination that adds up to exactly the target, then return null.

If there are multiple combinations possible, any single one is returned.

m target sum
n numbers.length

time: O(n * m ^ 2)
space: O(m ^ 2)

Tabulation is used for solving the problem.
*/
func HowSumTab(target int, numbers []int) ([]int, error) {
	// 1. visualize the problem as a table
	// 2. size the table based on the inputs
	// 3. initialize the table with defaults values
	table := make([][]int, target+1)
	for i := range table {
		table[i] = nil
	}

	// 4. seed the trivial answer into the table
	table[0] = []int{}

	// 5. iterate through the table
	for i := 0; i < target; i++ {
		// 6. fill further positions based on the current position
		currentPosition := table[i]
		if currentPosition != nil {
			for _, num := range numbers {
				nextPosition := i + num
				if nextPosition <= target {
					table[nextPosition] = append([]int{num}, currentPosition...)
				}
			}
		}
	}

	return table[target], nil
}
