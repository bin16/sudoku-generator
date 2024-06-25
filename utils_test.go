package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test_number_explode(t *testing.T) {
	var helper = func(src int, results []int) {
		r1 := number_explode(src)
		sort.Ints(r1)
		sort.Ints(results)
		if fmt.Sprintf("%v", r1) != fmt.Sprintf("%v", results) {
			t.Errorf("number_explode(%d) -> %v, want %v", src, r1, results)
		}
	}

	helper(256, []int{9})
	helper(256+128, []int{9, 8})
	helper(256+128+64, []int{9, 8, 7})
	helper(256+128+64+4+2+1, []int{9, 8, 7, 3, 2, 1})
	helper(256+128+64+1, []int{9, 8, 7, 1})
}

func Test_other_numbers(t *testing.T) {
	var helper = func(s, r0 []int) {
		r1 := other_numbers(s)
		sort.Ints(r0)
		sort.Ints(r1)
		if fmt.Sprintf("%v", r1) != fmt.Sprintf("%v", r0) {
			t.Errorf("other_numbers(%v) -> %v, want %v", s, r1, r0)
		}
	}

	helper([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9})
	helper([]int{1, 2, 3, 4}, []int{5, 6, 7, 8, 9})
	helper([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9})
	helper([]int{1}, []int{2, 3, 4, 5, 6, 7, 8, 9})
	helper([]int{5}, []int{1, 2, 3, 4, 6, 7, 8, 9})
	helper([]int{}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestGetRelatedCells(t *testing.T) {
	var helper = func(index int, r0 []int) {
		r1 := getRelatedCells(index)
		if fmt.Sprintf("%v", r1) != fmt.Sprintf("%v", r0) {
			t.Errorf("getRelatedCells(%d) => %v, want: %v\n", index, r1, r0)
		}
	}

	helper(0, []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11,
		18, 19, 20,
		27,
		36,
		45,
		54,
		63,
		72,
	})

	helper(21, []int{
		0 + 3, 1 + 3, 2 + 3,
		9 + 3, 10 + 3, 11 + 3,
		18, 19, 20, 21, 22, 23, 24, 25, 26,
		27 + 3,
		36 + 3,
		45 + 3,
		54 + 3,
		63 + 3,
		72 + 3,
	})

	helper(1, []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11,
		18, 19, 20,
		27 + 1,
		36 + 1,
		45 + 1,
		54 + 1,
		63 + 1,
		72 + 1,
	})

	helper(2, []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11,
		18, 19, 20,
		27 + 2,
		36 + 2,
		45 + 2,
		54 + 2,
		63 + 2,
		72 + 2,
	})

	helper(3, []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
		9 + 3, 10 + 3, 11 + 3,
		18 + 3, 19 + 3, 20 + 3,
		27 + 3,
		36 + 3,
		45 + 3,
		54 + 3,
		63 + 3,
		72 + 3,
	})

	helper(8, []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
		9 + 6, 10 + 6, 11 + 6,
		18 + 6, 19 + 6, 20 + 6,
		27 + 8,
		36 + 8,
		45 + 8,
		54 + 8,
		63 + 8,
		72 + 8,
	})

	helper(40, []int{
		0 + 4,
		9 + 4,
		18 + 4,
		27 + 3, 28 + 3, 29 + 3,
		36, 37, 38, 39, 40, 41, 42, 43, 44,
		45 + 3, 46 + 3, 47 + 3,
		54 + 4,
		63 + 4,
		72 + 4,
	})

	helper(41, []int{
		0 + 5,
		9 + 5,
		18 + 5,
		27 + 3, 28 + 3, 29 + 3,
		36, 37, 38, 39, 40, 41, 42, 43, 44,
		45 + 3, 46 + 3, 47 + 3,
		54 + 5,
		63 + 5,
		72 + 5,
	})

}
