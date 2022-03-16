package hot100

// 关键:
// dfs 裁剪
// 1. 什么时候裁剪,其实和原先的是一模一样的
// 2. 优化: 当当前数组的长度+ 剩余的长度 < k ,就可以退出了
// 3. 最关键的一点是: 还要考虑,当选择的位置不是当前的时候,
func combine(n int, k int) [][]int {
	temp := make([]int, 0)
	ret := make([][]int, 0)
	var dfs func(cur int)
	dfs = func(cur int) {
		// 优化: 当长度达不到k的时候就可以退出了
		if len(temp)+(n-cur+1) < k {
			return
		}
		if len(temp) == k {
			cp := make([]int, len(temp))
			copy(cp, temp)
			ret = append(ret, cp)
			return
		}
		temp = append(temp, cur)
		dfs(cur + 1)
		// 裁剪
		temp = temp[:len(temp)-1]
		// 关键第三步
		dfs(cur + 1)
	}
	dfs(1)
	return ret
}
