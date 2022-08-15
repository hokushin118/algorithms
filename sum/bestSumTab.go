package sum

/*
The function takes in a target number and an array of numbers as arguments and returns an array
containing the shortest combination of numbers that add up to exactly the target.

If there is a tie for the shortest combinations, the function returns any one of the shortest.

m target sum
n numbers.length

time:  O(n * m ^ 2)
space: O(m ^ 2)

Tabulation is used for solving the problem.
*/
func BestSumTab(target int, numbers []int) ([]int, error) {
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
	for i := 0; i <= target; i++ {
		// 6. fill further positions based on the current position
		currentPosition := table[i]
		if currentPosition != nil {
			for _, num := range numbers {
				nextPosition := i + num
				if nextPosition <= target {
					combination := append(currentPosition, num)
					if table[nextPosition] == nil || len(table[nextPosition]) > len(combination) {
						table[nextPosition] = combination
					}
				}
			}
		}
	}

	return table[target], nil
}
