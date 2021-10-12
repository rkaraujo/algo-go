package solution0557

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"God Ding", "doG gniD"},
		{"a", "a"},
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := reverseWords(tt.s)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
