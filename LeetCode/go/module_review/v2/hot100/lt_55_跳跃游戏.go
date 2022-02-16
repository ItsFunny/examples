package hot100

// 关键:
// 1个变量,维护当前能跳到的最大值
// for 循环,如果下标比最大值都要大,表明肯定跳不到该下标
func canJump(nums []int) bool {
	max := 0

	for index, v := range nums {
		if index > max {
			return false
		}

		// 计算当处于当前index的时候,能跳到的最远距离
		if v+index >= max {
			max = v + index
		}
	}

	return true
}
