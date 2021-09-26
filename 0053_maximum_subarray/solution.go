package solution

func maxSubArray(nums []int) int {
	max := nums[0]

	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum+nums[i] > nums[i] {
			sum = sum + nums[i]
		} else {
			sum = nums[i]
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
