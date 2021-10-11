package solution0167

func twoSum(numbers []int, target int) []int {
	mapDifferenceIndex := make(map[int]int)

	for idx, num := range numbers {
		mapDifferenceIndex[target-num] = idx
	}

	for idx, num := range numbers {
		otherIdx, exists := mapDifferenceIndex[num]
		if exists {
			// 1-indexed array
			return []int{idx + 1, otherIdx + 1}
		}
	}

	return nil
}
