package solution350

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.nums1, tt.nums2)
		t.Run(testname, func(t *testing.T) {
			ans := intersect(tt.nums1, tt.nums2)
			if !IntArrayEquals(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
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
