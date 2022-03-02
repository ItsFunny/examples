package hot100

// 关键
// 1. 动态规划
// 2. 当处于第一行的时候,只能从左边过来
// 3. 当处于第一列的时候,只能从上面过来
// 4. 当处于其他地方的时候,结果值为 最小值+ 当前值
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
	}

	// 初始化第一行
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	// 初始化第一列
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 初始化其他地方
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			dp[i][j] = minPathSumMin(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
func minPathSumMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
