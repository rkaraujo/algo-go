package solution0387

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
		{"x", 0},
		{"xxxxz", 4},
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := firstUniqChar(tt.s)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
