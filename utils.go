package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func getRelatedCells(index int) []int {
	var rMap = map[int]bool{}

	col := index % 9
	row := index / 9
	for i := 0; i < 9; i++ {
		k := row*9 + i
		rMap[k] = true
	}
	for i := 0; i < 9; i++ {
		k := col + i*9
		rMap[k] = true
	}
	hCol := col / 3
	hRow := row / 3
	for i := 0; i < 3; i++ {
		for r := 0; r < 3; r++ {
			k := hRow*3*9 + i*9 + hCol*3 + r
			rMap[k] = true
		}
	}

	var results = []int{}
	for k := range rMap {
		results = append(results, k)
	}
	sort.Ints(results)

	return results
}

func getAvailableNumbers(seq [81]int, index int) []int {
	var m int = 0
	var results []int = []int{}

	col := index % 9
	row := index / 9
	for i := 0; i < 9; i++ {
		k := col + 9*i
		num := seq[k]
		if num > 0 {
			m |= 1 << (num - 1)
		}
	}

	for i := 0; i < 9; i++ {
		k := row*9 + i
		num := seq[k]
		if num > 0 {
			m |= 1 << (num - 1)
		}

	}

	hCol := col / 3
	hRow := row / 3
	hIndex := hCol*3 + hRow*3*9
	for i := 0; i < 3; i++ {
		for r := 0; r < 3; r++ {
			k := hIndex + i + 9*r
			num := seq[k]
			if num > 0 {
				m |= 1 << (num - 1)
			}
		}
	}

	for i := 0; i < 9; i++ {
		if (1<<i)&m == 0 {
			results = append(results, i+1)
		}
	}

	return results
}

func printSudoku(seq [81]int) {
	count := 0
	for i := 0; i < 81; i++ {
		if seq[i] != 0 {
			count += 1
		}
	}
	fmt.Println()
	fmt.Println("+-------+-------+-------+")
	fmt.Printf("|      Sudoku [%2d]      |\n", count)
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		if row >= 3 && row%3 == 0 {
			fmt.Printf("+-------+-------+-------+\n")
		}
		for col := 0; col < 9; col++ {
			if col == 0 {
				fmt.Printf("| ")
			}
			if col >= 3 && col%3 == 0 {
				fmt.Printf("| ")
			}
			k := col + 9*row
			num := seq[k]
			if num == 0 {
				fmt.Printf("_ ")
			} else {
				fmt.Printf("%d ", num)
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Println("+-------+-------+-------+")
	fmt.Printf("data = [\n    ")
	for i, num := range seq {
		fmt.Printf("%d, ", num)
		if i%9 == 8 {
			fmt.Printf("\n    ")
		}
	}
	fmt.Printf("]\n")
}

func shuffle(s []int) {
	for i := range s {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func other_numbers(numbers []int) []int {
	u := numbers_join(numbers)
	u2 := (^u) & (512 - 1)
	return number_explode(u2)
}

func numbers_join(a []int) int {
	u := 0
	for _, num := range a {
		u |= 1 << (num - 1)
	}
	return u
}

func number_explode(u int) []int {
	numbers := []int{}
	for i := 0; i < 9; i++ {
		if 1<<i&u == 0 {
			continue
		}

		numbers = append(numbers, i+1)
	}

	return numbers
}
