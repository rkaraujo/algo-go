package solution0118

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.numRows)
		t.Run(testname, func(t *testing.T) {
			ans := generate(tt.numRows)
			if !check(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func check(ans [][]int, want [][]int) bool {
	for i := 0; i < len(want); i++ {
		for j := 0; j < len(want[0]); j++ {
			if ans[i][j] != want[i][j] {
				return false
			}
		}
	}
	return true
}
