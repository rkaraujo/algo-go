package solution0121

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{8, 8}, 0},
		{[]int{2}, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.prices)
		t.Run(testname, func(t *testing.T) {
			ans := maxProfit(tt.prices)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
