package solution0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	for j := 0; j < len(t); j++ {
		count := freq[t[j]]
		if count == 0 {
			return false
		}
		freq[t[j]]--
	}

	return true
}
