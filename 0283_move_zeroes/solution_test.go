package solution0283

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{2}, []int{2}},
		{[]int{1, 0}, []int{1, 0}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.nums)
		t.Run(testname, func(t *testing.T) {
			moveZeroes(tt.nums)
			if !IntArrayEquals(tt.nums, tt.want) {
				t.Errorf("got %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
