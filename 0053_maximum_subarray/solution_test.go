package solution

import "testing"
import "fmt"

func TestMaxSubArray(t *testing.T) {
	var tests = []struct {
		nums []int
		want int
	}{
		{[]int{-1}, -1},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := maxSubArray(tt.nums)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
