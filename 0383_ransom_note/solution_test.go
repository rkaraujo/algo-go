package solution0383

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s, %s", tt.ransomNote, tt.magazine)
		t.Run(testname, func(t *testing.T) {
			ans := canConstruct(tt.ransomNote, tt.magazine)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
