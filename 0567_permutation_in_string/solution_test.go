package solution0567

import (
	"fmt"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"ab", "ba", true},
		{"a", "a", true},
		{"ab", "a", false},
		{"abb", "abc", false},
		{"hello", "ooolleoooleh", false},
		{"adc", "dcda", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s, %s", tt.s1, tt.s2)
		t.Run(testname, func(t *testing.T) {
			ans := checkInclusion(tt.s1, tt.s2)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
