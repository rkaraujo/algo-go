package solution0977

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	// the biggest squared values are from the extremes
	right_index := 0
	left_index := len(nums) - 1

	// adding biggest values from the end
	result_index := len(nums) - 1
	for right_index <= left_index {
		var cur_num int
		if abs(nums[left_index]) > abs(nums[right_index]) {
			cur_num = nums[left_index]
			left_index--
		} else {
			cur_num = nums[right_index]
			right_index++
		}

		result[result_index] = cur_num * cur_num
		result_index--
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
