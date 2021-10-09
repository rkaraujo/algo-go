package solution0704

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		index := (end + start) / 2
		if nums[index] == target {
			return index
		}

		if nums[index] < target {
			start = index + 1
		} else {
			end = index - 1
		}
	}

	return -1
}
