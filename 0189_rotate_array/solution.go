package solution0189

func rotate(nums []int, k int) {
	if k >= len(nums) {
		k = k % len(nums)
	}

	nums_copy := make([]int, len(nums))
	copy(nums_copy, nums)

	for i := 0; i < len(nums_copy); i++ {
		new_index := i + k
		if new_index > len(nums_copy)-1 {
			new_index = new_index % len(nums_copy)
		}

		nums[new_index] = nums_copy[i]
	}
}

func rotateInline(nums []int, k int) {
	if k >= len(nums) {
		k = k % len(nums)
	}

	for i := 0; i < k; i++ {
		new_first := nums[len(nums)-1]

		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}

		nums[0] = new_first
	}
}
