package hot100

import "sort"

// 关键:
// 1. 排序: 因为涉及到去重(去重必然有排序,排序可以将相同的放在一起)
// 2. dfs+剪枝
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	used := make([]bool, len(nums))
	var dfs func(index int)
	cur := make([]int, 0)
	dfs = func(index int) {
		if index == len(nums) {
			ret = append(ret, append([]int{}, cur...))
			return
		}
		// 因为是要构造成长度相同的值,所以需要
		for i, v := range nums {
			if used[i] || i > 0 && !used[i-1] && nums[i-1] == nums[i] {
				// !used[i-1]的原因在于,当剪枝的时候, used[i-1] 会被设置为false(因为for 循环i-1 肯定是在i之前遍历到的)
				// 而当遍历到当前下标i的时候,i-1肯定为false
				continue
			}
			cur = append(cur, v)
			used[i] = true
			dfs(index + 1)
			// 剪枝
			used[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return ret
}
