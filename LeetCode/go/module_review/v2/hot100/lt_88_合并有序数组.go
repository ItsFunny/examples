package hot100

// 关键:
// 逆向双指针,将数更大的放在后面
func merge88(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, index := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0;index-- {
		var cur int
		// 说明p1到了末尾
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}

		nums1[index] = cur
	}
}
