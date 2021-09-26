package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	result := make([]int, m+n)

	i := 0
	j := 0
	var n1, n2 int

	for i+j < m+n {
		if i < m {
			n1 = nums1[i]
		}
		if j < n {
			n2 = nums2[j]
		}

		if j >= n || (n1 < n2 && i < m) {
			result[i+j] = n1
			i++
		} else {
			result[i+j] = n2
			j++
		}
	}

	copy(nums1, result)
}
