package solution0020

func isValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		cur := s[i]

		if len(stack) == 0 {
			stack = append(stack, cur)
			continue
		}

		prev := stack[len(stack)-1]
		if (prev == '[' && cur == ']') || (prev == '(' && cur == ')') || (prev == '{' && cur == '}') {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, cur)
		}
	}

	return len(stack) == 0
}
