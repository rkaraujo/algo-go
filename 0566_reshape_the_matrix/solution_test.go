package solution0566

import (
	"fmt"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	tests := []struct {
		mat  [][]int
		r    int
		c    int
		want [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
		{[][]int{{1, 2}, {3, 4}}, 2, 4, [][]int{{1, 2}, {3, 4}}},
		{[][]int{{1, 2, 3, 4}}, 2, 2, [][]int{{1, 2}, {3, 4}}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %d, %d", tt.mat, tt.r, tt.c)
		t.Run(testname, func(t *testing.T) {
			ans := matrixReshape(tt.mat, tt.r, tt.c)
			if !check(ans, tt.want, tt.r, tt.c) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func check(ans [][]int, want [][]int, r int, c int) bool {
	for i := 0; i < len(want); i++ {
		for j := 0; j < len(want[0]); j++ {
			if ans[i][j] != want[i][j] {
				return false
			}
		}
	}
	return true
}
