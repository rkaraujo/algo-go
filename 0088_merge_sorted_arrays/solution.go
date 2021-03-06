package solution

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for k := len(nums1) - 1; k >= 0; k-- {
		if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) {
			nums1[k] = nums1[i]
			i--
		} else if i < 0 || (j >= 0 && nums2[j] > nums1[i]) {
			nums1[k] = nums2[j]
			j--
		}
	}
}
