package main

import "testing"

func TestCheckSudoku(t *testing.T) {
	var helper = func(seq [81]int, r0 bool) {
		if r1 := checkSudoku(seq); r1 != r0 {
			t.Errorf("checkSudoku() => %v, want: %v; \n %v", r1, r0, seq)
		}
	}

	helper(VERY_HARD, true)
	helper(EASY_GAME, true)
	helper(NOT_UNIQUE, false)
}
