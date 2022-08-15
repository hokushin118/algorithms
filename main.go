package main

import (
	"fmt"
	"github.com/hokushin118/algorithms/fibonacci"
	"github.com/hokushin118/algorithms/grid"
	"github.com/hokushin118/algorithms/sum"
	"github.com/hokushin118/algorithms/word"
)

/*
go run main.go
*/
func main() {
	fmt.Println("1. Fibonacci number - recursive")

	fmt.Println(fibonacci.Fib(0, map[int]int{}))   // 1
	fmt.Println(fibonacci.Fib(1, map[int]int{}))   // 1
	fmt.Println(fibonacci.Fib(2, map[int]int{}))   // 1
	fmt.Println(fibonacci.Fib(7, map[int]int{}))   // 13
	fmt.Println(fibonacci.Fib(8, map[int]int{}))   // 21
	fmt.Println(fibonacci.Fib(50, map[int]int{}))  // 12586269025
	fmt.Println(fibonacci.Fib(100, map[int]int{})) // 3736710778780434371

	fmt.Println("2. Fibonacci number - tabulation")

	fmt.Println(fibonacci.FibTab(0))   // 0
	fmt.Println(fibonacci.FibTab(1))   // 1
	fmt.Println(fibonacci.FibTab(2))   // 1
	fmt.Println(fibonacci.FibTab(3))   // 2
	fmt.Println(fibonacci.FibTab(7))   // 13
	fmt.Println(fibonacci.FibTab(8))   // 21
	fmt.Println(fibonacci.FibTab(50))  // 12586269025
	fmt.Println(fibonacci.FibTab(100)) // 3736710778780434371

	fmt.Println("3. Grid traveller - recursive")

	fmt.Println(grid.Traveler(0, 0, map[string]int{}))   // 0
	fmt.Println(grid.Traveler(0, 1, map[string]int{}))   // 0
	fmt.Println(grid.Traveler(1, 0, map[string]int{}))   // 0
	fmt.Println(grid.Traveler(8, 0, map[string]int{}))   // 0
	fmt.Println(grid.Traveler(1, 1, map[string]int{}))   // 1
	fmt.Println(grid.Traveler(2, 3, map[string]int{}))   // 3
	fmt.Println(grid.Traveler(3, 2, map[string]int{}))   // 3
	fmt.Println(grid.Traveler(3, 3, map[string]int{}))   // 6
	fmt.Println(grid.Traveler(5, 3, map[string]int{}))   // 15
	fmt.Println(grid.Traveler(18, 18, map[string]int{})) // 2333606220

	fmt.Println("4. Grid traveller - tabular")

	fmt.Println(grid.TravelerTab(0, 0))   // 0
	fmt.Println(grid.TravelerTab(0, 1))   // 0
	fmt.Println(grid.TravelerTab(1, 0))   // 0
	fmt.Println(grid.TravelerTab(8, 0))   // 0
	fmt.Println(grid.TravelerTab(1, 1))   // 1
	fmt.Println(grid.TravelerTab(2, 3))   // 3
	fmt.Println(grid.TravelerTab(3, 2))   // 3
	fmt.Println(grid.TravelerTab(3, 3))   // 6
	fmt.Println(grid.TravelerTab(5, 3))   // 15
	fmt.Println(grid.TravelerTab(18, 18)) // 2333606220

	fmt.Println("5. Can Sum - recursive")

	fmt.Println(sum.CanSum(7, []int{2, 3}, map[int]bool{}))       // (2 + 2 + 3) = true
	fmt.Println(sum.CanSum(7, []int{5, 3, 4, 7}, map[int]bool{})) // (7, 3 + 4) = true
	fmt.Println(sum.CanSum(7, []int{2, 4}, map[int]bool{}))       // false
	fmt.Println(sum.CanSum(8, []int{2, 3, 5}, map[int]bool{}))    // (2 + 2 + 2 + 2, 5 + 3) = true
	fmt.Println(sum.CanSum(300, []int{7, 14}, map[int]bool{}))    // false

	fmt.Println("6. Can Sum - tabular")

	fmt.Println(sum.CanSumTab(7, []int{2, 3}))       // (2 + 2 + 3) = true
	fmt.Println(sum.CanSumTab(7, []int{5, 3, 4, 7})) // (7, 3 + 4) = true
	fmt.Println(sum.CanSumTab(7, []int{2, 4}))       // false
	fmt.Println(sum.CanSumTab(8, []int{2, 3, 5}))    // (2 + 2 + 2 + 2, 5 + 3) = true
	fmt.Println(sum.CanSumTab(300, []int{7, 14}))    // false

	fmt.Println("7. How Sum - recursive")

	fmt.Println(sum.HowSum(0, []int{1, 2, 3}, map[int][]int{}))    // []
	fmt.Println(sum.HowSum(7, []int{2, 3}, map[int][]int{}))       // [2, 2, 3]
	fmt.Println(sum.HowSum(7, []int{5, 3, 4, 7}, map[int][]int{})) // [3, 4]
	fmt.Println(sum.HowSum(7, []int{2, 4}, map[int][]int{}))       // []
	fmt.Println(sum.HowSum(8, []int{2, 3, 5}, map[int][]int{}))    // [2, 2, 2, 2]
	fmt.Println(sum.HowSum(300, []int{7, 14}, map[int][]int{}))    // null

	fmt.Println("8. How Sum - tabular")

	fmt.Println(sum.HowSumTab(0, []int{1, 2, 3}))    // []
	fmt.Println(sum.HowSumTab(7, []int{2, 3}))       // [2, 2, 3]
	fmt.Println(sum.HowSumTab(7, []int{5, 3, 4, 7})) // [3, 4]
	fmt.Println(sum.HowSumTab(7, []int{2, 4}))       // []
	fmt.Println(sum.HowSumTab(8, []int{2, 3, 5}))    // [2, 2, 2, 2]
	fmt.Println(sum.HowSumTab(300, []int{7, 14}))    // null

	fmt.Println("9. Best Sum - recursive")

	fmt.Println(sum.BestSum(7, []int{5, 3, 4, 7}, map[int][]int{}))    // [7]
	fmt.Println(sum.BestSum(8, []int{2, 3, 5}, map[int][]int{}))       // [3 5]
	fmt.Println(sum.BestSum(8, []int{1, 4, 5}, map[int][]int{}))       // [4 4]
	fmt.Println(sum.BestSum(100, []int{1, 2, 5, 25}, map[int][]int{})) // [25 25 25 25]

	fmt.Println("10. Best Sum - tabular")

	fmt.Println(sum.BestSumTab(7, []int{5, 3, 4, 7}))    // [7]
	fmt.Println(sum.BestSumTab(8, []int{2, 3, 5}))       // [3 5]
	fmt.Println(sum.BestSumTab(8, []int{1, 4, 5}))       // [4 4]
	fmt.Println(sum.BestSumTab(100, []int{1, 2, 5, 25})) // [25 25 25 25]

	fmt.Println("9. Can Construct - recursive")

	fmt.Println(word.CanConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}, map[string]bool{}))     // true
	fmt.Println(word.CanConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]bool{})) // false
	fmt.Println(word.CanConstruct("", []string{"cat", "dog", "mouse"}, map[string]bool{}))                                 // true
	fmt.Println(word.CanConstruct("aaaaaaa", []string{"a", "aa", "aaa"}, map[string]bool{}))                               // true

	fmt.Println("10. Can Construct - tabular")

	fmt.Println(word.CanConstructTab("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))     // true
	fmt.Println(word.CanConstructTab("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"})) // false
	fmt.Println(word.CanConstructTab("", []string{"cat", "dog", "mouse"}))                                 // true
	fmt.Println(word.CanConstructTab("aaaaaaa", []string{"a", "aa", "aaa"}))                               // true

	fmt.Println("11. Count Construct - recursive")

	fmt.Println(word.CountConstruct("", []string{"cat", "dog", "mouse"}, map[string]int{}))                                 // 1
	fmt.Println(word.CountConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}, map[string]int{}))     // 4
	fmt.Println(word.CountConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]int{})) // 0
	fmt.Println(word.CountConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, map[string]int{}))                // 2
	fmt.Println(word.CountConstruct("aaaaaaa", []string{"a", "aa", "aaa"}, map[string]int{}))                               // 44

	fmt.Println("12. Count Construct - tabular")

	fmt.Println(word.CountConstructTab("", []string{"cat", "dog", "mouse"}))                                 // 1
	fmt.Println(word.CountConstructTab("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))     // 4
	fmt.Println(word.CountConstructTab("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"})) // 0
	fmt.Println(word.CountConstructTab("purple", []string{"purp", "p", "ur", "le", "purpl"}))                // 2
	fmt.Println(word.CountConstructTab("aaaaaaa", []string{"a", "aa", "aaa"}))                               // 44 	// ...
}
