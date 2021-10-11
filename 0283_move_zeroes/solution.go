package solution0283

func moveZeroes(nums []int) {
	insertPos := 0

	for _, num := range nums {
		if num != 0 {
			nums[insertPos] = num
			insertPos++
		}
	}

	for i := insertPos; i < len(nums); i++ {
		nums[i] = 0
	}
}
