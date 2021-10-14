package solution0567

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	// store s1 count
	var s1count [26]int
	for i := 0; i < len(s1); i++ {
		s1count[s1[i]-'a']++
	}

	// store s2 window count
	var s2windowcount [26]int
	for i := 0; i < len(s1); i++ {
		s2windowcount[s2[i]-'a']++
	}

	// check every window
	for i := 0; i < len(s2)-len(s1); i++ {
		if isSameCount(s1count, s2windowcount) {
			return true
		}

		// reduce last char count
		s2windowcount[s2[i]-'a']--
		// add next char count
		s2windowcount[s2[i+len(s1)]-'a']++
	}

	return isSameCount(s1count, s2windowcount)
}

func isSameCount(s1count, s2count [26]int) bool {
	for i := 0; i < 26; i++ {
		if s1count[i] != s2count[i] {
			return false
		}
	}
	return true
}
