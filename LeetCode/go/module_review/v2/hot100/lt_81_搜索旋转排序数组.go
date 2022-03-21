package hot100

func search81(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		}
		// 先判断正常情况的内容
		// 正常情况为: mid 处于正常区间, mid>left
		if nums[mid] > nums[left] {
			// 当处于正常区间,表明 left--> mid 这区间是正常的,但是mid --> right是不能保证的
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 表明处于不正常区间,则此时,从mid-right 区间判断
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == target || nums[right] == target {
		return true
	}
	return false
}
