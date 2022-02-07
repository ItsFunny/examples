package hot100

// 关键
// 1. 找到当前下标所能跳的最远的下表
func jump(nums []int) int {
	end := 0 // 当前所能跳到的最远的边界
	maxPosition := 0
	steps := 0

	// 从下标0开始
	// 注意退出条件,因为是只需要跳到最后一个,所以临界条件是长度-1
	for i := 0; i < len(nums)-1; i++ {
		// 当处于该下标,获取得到当前下标能跳跃的最大值,
		// 注意: maxPosition是下标,所以需要加上当前 i
		maxPosition = jumpMax(maxPosition, i+nums[i])
		// 关键是这一步: 当达到了边界之后(这个end其实可以认为是之前的end),更新end边界
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}

func jumpMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
