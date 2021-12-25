package hot100

// 关键:
// mid 值可能在 反转区间
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[left] {
			// 表明在正常区间
			// 则开始判断target 所处的范围: 是在反转区间内,还是在正常区间
			if target >= nums[left] && target < nums[mid] {
				// 正常区间
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 表面,mid 处于 反转区间
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}
