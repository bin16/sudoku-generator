package main

// WARN: game must be solvable
func checkSudoku(seq [81]int) bool {
	// seq1 := seq
	var count = 0
	var solve func(seq *[81]int, index int)
	solve = func(seq *[81]int, index int) {
		// done, count+1
		if index >= 80 && seq[80] != 0 {
			count += 1
			// fmt.Printf("Solution.%d", count)
			// printSudoku(*seq)
			return
		}

		// exit when count > 1
		if count > 1 {
			return
		}

		// a. already a number
		if seq[index] != 0 {
			solve(seq, index+1)
			return
		}

		// b. no number yet
		numbers := getAvailableNumbers(*seq, index)
		for _, num := range numbers {
			seq[index] = num
			solve(seq, index+1)
			seq[index] = 0
		}
	}

	solve(&seq, 0)

	return count == 1
}
