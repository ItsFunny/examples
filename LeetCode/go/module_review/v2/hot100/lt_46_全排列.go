package hot100

// 关键:
// 看到这种题,直接回溯(dfs+裁剪)
// 需要注意点的是,需要有一个bool数组标识是否已经使用
// 同时也需要注意的是,append 数据的时候要使用拷贝
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	used := make([]bool, len(nums))
	var dfs func(index int)
	cur := make([]int, 0)
	dfs = func(count int) {
		if count == len(nums) {
			ret = append(ret, append([]int{},cur...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				// 追加数据
				used[i] = true
				cur = append(cur, nums[i])
				dfs(count + 1)
				// 裁剪
				used[i] = false
				cur = cur[:len(cur)-1]
			}
		}
	}

	dfs(0)
	return ret
}
