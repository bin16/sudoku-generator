package main

import (
	"fmt"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	var helper = func(seq, r0 [81]int, ok0 bool) {
		r1, ok1 := solveSudoku(seq)
		if ok1 != ok0 {
			t.Errorf("solveSudoku() => ..., %v; want %v \n; %v\n", ok1, ok0, seq)
		} else if fmt.Sprintf("%v", r1) != fmt.Sprintf("%v", r0) {
			t.Errorf("solveSudoku() => ..., true;\n* got %v; \n* want %v\n", r1, r0)
		}
	}

	helper(VERY_HARD, VERY_HARD_SOLVED, true)
	helper(EASY_GAME, EASY_GAME_SOLVED, true)
}
