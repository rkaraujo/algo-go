package solution0020

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"{[]}", true},
		{"([)]", false},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := isValid(tt.s)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
