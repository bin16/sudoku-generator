package main

import (
	"fmt"
	"testing"
)

func TestGenerateSudoku(t *testing.T) {
	var helper = func() {
		q, a := generateSudoku()
		a1, ok := solveSudoku(q)
		if !ok {
			t.Errorf("sudoku should be solved: \n%v\n", q)
		} else if fmt.Sprintf("%v", a1) != fmt.Sprintf("%v", a) {
			t.Errorf("sudoku not match: \n Q: %v\nA0: %v \nA1: %v\n", q, a, a1)
		}
	}

	for i := 0; i < 100; i++ {
		helper()
	}
}

func BenchmarkGenerateSudoku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q, _ := generateSudoku()
		solveSudoku(q)
	}
}
