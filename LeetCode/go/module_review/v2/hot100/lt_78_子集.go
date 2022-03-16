package hot100

import "sort"

// 关键: dfs
// 1. 排序: 排序使得可以去重
// 2. dfs
func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	var dfs func(cur int)
	temp := make([]int, 0)
	dfs = func(cur int) {
		if cur == len(nums) {
			ret = append(ret, append([]int{}, temp...))
			return
		}
		temp = append(temp, nums[cur])
		// 选择当前值
		dfs(cur + 1)
		// 裁剪
		temp = temp[:len(temp)-1]
		// 不选择当前值
		dfs(cur + 1)
	}
	dfs(0)
	return ret
}
