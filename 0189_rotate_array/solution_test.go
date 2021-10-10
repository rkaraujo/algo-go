package solution0189

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{1, 2}, 0, []int{1, 2}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 2}, 4, []int{1, 2}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %d", tt.nums, tt.k)
		t.Run(testname, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			if !IntArrayEquals(tt.nums, tt.want) {
				t.Errorf("got %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

func TestRotateInline(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{1, 2}, 0, []int{1, 2}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 2}, 4, []int{1, 2}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %d", tt.nums, tt.k)
		t.Run(testname, func(t *testing.T) {
			rotateInline(tt.nums, tt.k)
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
