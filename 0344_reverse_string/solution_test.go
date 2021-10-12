package solution0344

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		{[]byte{'u', 'v', 'a'}, []byte{'a', 'v', 'u'}},
		{[]byte{'x'}, []byte{'x'}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.s)
		t.Run(testname, func(t *testing.T) {
			reverseString(tt.s)
			if !ByteArrayEquals(tt.s, tt.want) {
				t.Errorf("got %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func ByteArrayEquals(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
