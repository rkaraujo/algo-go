package solution

func ContainsDuplicate(nums []int) bool {
	num_exists := make(map[int]bool)

	for _, num := range nums {
		if num_exists[num] {
			return true
		}
		num_exists[num] = true
	}

	return false
}
