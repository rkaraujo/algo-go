package solution

func twoSum(nums []int, target int) []int {
	map_num_index := make(map[int]int)
	for index, num := range nums {
		map_num_index[num] = index
	}

	for index, num := range nums {
		dif := target - num
		other_index, exists := map_num_index[dif]
		if exists && index != other_index {
			if index < other_index {
				return []int{index, other_index}
			} else {
				return []int{other_index, index}
			}
		}
	}

	return nil
}
