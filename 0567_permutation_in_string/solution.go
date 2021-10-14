package solution0567

func checkInclusion(s1 string, s2 string) bool {
	lenS1 := len(s1)

	for i := 0; i+lenS1 <= len(s2); i++ {
		if checkPermutation(s1, s2[i:i+lenS1]) {
			return true
		}
	}

	return false
}

func checkPermutation(s1 string, subStringS2 string) bool {
	mapS1 := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		mapS1[s1[i]]++
	}

	for i := 0; i < len(subStringS2); i++ {
		if mapS1[subStringS2[i]] == 0 {
			return false
		}
		mapS1[subStringS2[i]]--
	}
	return true
}
