package main

func generateSudokuV3(maxClues int) (q, a [81]int) {
	a = generateFullSudoku()

	for {
		q = a

		// generate the route
		kl := []int{}
		for i := 0; i < 81; i++ {
			kl = append(kl, i)
		}
		shuffle(kl)

		// remove numbers one by one
		for _, k := range kl {
			q[k] = 0

			// unless it has more than one solutions
			if !checkSudoku(q) {
				q[k] = a[k]
				if 81-k < maxClues {
					return
				}
			}
		}

	}
}
