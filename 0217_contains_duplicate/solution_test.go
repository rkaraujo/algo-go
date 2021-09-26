package solution

import "testing"
import "fmt"

func TestContainsDuplicate(t *testing.T) {
	var tests = []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 3}, true},
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{8}, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := ContainsDuplicate(tt.nums)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
