package solution0704

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{9}, 9, 0},
		{[]int{9}, 2, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{[]int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{[]int{-1, 0, 3, 5, 9, 12}, 5, 3},
		{[]int{-1, 0, 3, 9, 12}, 3, 2},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %d", tt.nums, tt.target)
		t.Run(testname, func(t *testing.T) {
			ans := search(tt.nums, tt.target)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
