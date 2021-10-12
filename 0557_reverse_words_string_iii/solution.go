package solution0557

func reverseWords(s string) string {
	startWordIdx := 0
	spaceIdx := 0

	result := ""
	for startWordIdx < len(s) {
		// find next space
		for spaceIdx < len(s) && s[spaceIdx] != ' ' {
			spaceIdx++
		}

		// append reversed word
		for i := spaceIdx - 1; i >= startWordIdx; i-- {
			result += string(s[i])
		}
		if spaceIdx != len(s) {
			result += " "
		}

		// next word
		startWordIdx = spaceIdx + 1
		spaceIdx++
	}

	return result
}
