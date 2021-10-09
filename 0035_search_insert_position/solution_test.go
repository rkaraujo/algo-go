package solution0035

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 0, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %d", tt.nums, tt.target)
		t.Run(testname, func(t *testing.T) {
			ans := searchInsert(tt.nums, tt.target)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
