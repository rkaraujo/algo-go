package solution

import "testing"
import "fmt"

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{-1, 0, 3, 0, 0, 0}, 3, []int{-2, 5, 6}, 3, []int{-2, -1, 0, 3, 5, 6}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %d, %v, %d", tt.nums1, tt.m, tt.nums2, tt.n)
		t.Run(testname, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !IntArrayEquals(tt.nums1, tt.want) {
				t.Errorf("got %v, want %v", tt.nums1, tt.want)
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
