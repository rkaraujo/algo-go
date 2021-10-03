package solution0387

func firstUniqChar(s string) int {
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	for j := 0; j < len(s); j++ {
		if freq[s[j]] == 1 {
			return j
		}
	}

	return -1
}
