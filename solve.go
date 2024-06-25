package main

func solveSudoku(seq [81]int) ([81]int, bool) {
	var solve func(seq *[81]int, index int) bool
	solve = func(seq *[81]int, index int) bool {
		// done
		if index > 80 && seq[80] != 0 {
			return true
		}

		// a. already a number
		if seq[index] != 0 {
			return solve(seq, index+1)
		}

		// b. no number yet
		numbers := getAvailableNumbers(*seq, index)
		for _, num := range numbers {
			seq[index] = num
			if solve(seq, index+1) {
				return true
			}
			seq[index] = 0
		}

		return false
	}

	ok := solve(&seq, 0)

	return seq, ok
}
