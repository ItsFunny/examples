package hot100

// 关键: 动态规划
// f(i) 为到当前i的最大值(既累加)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 注意: 只有当 当前值更大的时候,才触发更新,因为求的是连续的最大和,而不是整个数组中的最大和
		// 既: f(i)是单独的一个,还是作为整体中的一部分参与比较大小,而这个判断条件为 和
		if nums[i]+nums[i-1]>nums[i]{
			nums[i]=nums[i]+nums[i-1]
		}
		if nums[i]>max{
			max=nums[i]
		}
	}

	return max
}

