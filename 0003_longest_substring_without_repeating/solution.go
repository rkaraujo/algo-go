package solution0003

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	max := 1

	mapIndexes := make(map[byte]int)
	index := 0
	curLength := 0
	for index < len(s) {
		existingIndex, exists := mapIndexes[s[index]]
		if exists {
			mapIndexes = make(map[byte]int)
			index = existingIndex + 1
			curLength = 0
			continue
		}

		mapIndexes[s[index]] = index

		curLength++
		if curLength > max {
			max = curLength
		}

		index++
	}

	return max
}
