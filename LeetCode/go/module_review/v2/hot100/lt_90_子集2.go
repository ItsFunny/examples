package hot100

import "sort"

// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// 关键: dfs
// 若当前数x之前的数y ,x==y,并且y 之前没有被选中,则当前可以直接退出
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	var dfs func(chosePre bool, index int)

	temp := make([]int, 0)
	dfs = func(chosePre bool, index int) {
		if index == len(nums) {
			ret = append(ret, append([]int{}, temp...))
			return
		}
		dfs(false, index+1)
		// 如果之前的数没有被选择
		if !chosePre && index > 0 && nums[index-1] == nums[index] {
			// 并且紧邻的两个数相等
			// 则直接return
			return
		}
		temp = append(temp, nums[index])
		dfs(true, index+1)
		// 裁剪
		temp = temp[:len(temp)-1]
	}
	dfs(false, 0)
	return ret
}
