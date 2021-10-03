package solution0383

func canConstruct(ransomNote string, magazine string) bool {
	freqmagazine := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		freqmagazine[magazine[i]]++
	}

	for j := 0; j < len(ransomNote); j++ {
		count := freqmagazine[ransomNote[j]]
		if count == 0 {
			return false
		}
		freqmagazine[ransomNote[j]]--
	}

	return true
}
