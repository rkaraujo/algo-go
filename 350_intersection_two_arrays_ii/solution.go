package solution350

func intersect(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)
	for _, num1 := range nums1 {
		map1[num1]++
	}

	result := []int{}
	for _, num2 := range nums2 {
		count := map1[num2]
		if count > 0 {
			result = append(result, num2)
			map1[num2]--
		}
	}

	return result
}
