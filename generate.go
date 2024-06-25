package main

func generateFullSudoku() [81]int {
	sequence := [81]int{}

	var generate func(seq *[81]int, index int) bool
	generate = func(seq *[81]int, index int) bool {
		if seq[80] != 0 {
			return true
		}

		numbers := getAvailableNumbers(*seq, index)
		shuffle(numbers)

		for _, num := range numbers {
			seq[index] = num
			if generate(seq, index+1) {
				return true
			}
			seq[index] = 0
		}

		return false
	}
	generate(&sequence, 0)

	return sequence
}

func generateSudoku() (q, a [81]int) {
	a = generateFullSudoku()
	q = a

	kl := []int{}
	for i := 0; i < 81; i++ {
		kl = append(kl, i)
	}
	shuffle(kl)

	for _, k := range kl {
		q[k] = 0
		if !checkSudoku(q) {
			q[k] = a[k]
			return q, a
		}
	}

	return
}
