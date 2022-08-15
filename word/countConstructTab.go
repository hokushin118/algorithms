package word

/*
The function takes in a target string and an array of strings as arguments and returns the number of ways
that the target can be constructed by concatenating elements of the words array.

The element of the array can be used as many times as needed.

m target.length number of characters in the target word
n words.length number of words in the words array

time: O(n * m ^ 2)
space: O(m)

Tabulation is used for solving the problem.
*/
func CountConstructTab(target string, words []string) (int, error) {
	// 1. visualize the problem as a table
	// 2. size the table based on the inputs
	// 3. initialize the table with defaults values
	wordLength := len(target)
	table := make([]int, wordLength+1)
	for i := range table {
		table[i] = 0
	}

	// 4. seed the trivial answer into the table
	table[0] = 1

	// 5. iterate through the table
	for i := 0; i <= wordLength; i++ {
		// 6. fill further positions based on the current position
		currentPosition := table[i]
		if currentPosition != 0 {
			for _, word := range words {
				nextPosition := i + len(word)
				if nextPosition <= len(target) {
					if target[i:nextPosition] == word {
						table[nextPosition] += currentPosition
					}
				}
			}
		}
	}

	return table[wordLength], nil
}
