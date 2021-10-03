package solution0242

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"car", "rat", false},
		{"r", "r", true},
		{"r", "a", false},
		{"ab", "a", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s, %s", tt.s, tt.t)
		t.Run(testname, func(t *testing.T) {
			ans := isAnagram(tt.t, tt.s)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
