package hot100

// 关键:
//
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	// 与62类似,当处于边界的时候,只能一种走法
	ret := make([][]int, len(obstacleGrid))
	// 初始化元素
	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, len(obstacleGrid[i]))
	}
	// 初始化第一列
	for i := 0; i < len(ret); i++ {
		if obstacleGrid[i][0] == 1 {
			// 当边界为1的时候,则再也走不下去了,所以直接break
			break
		}
		ret[i][0] = 1
	}
	// 初始化第一行
	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			// 同上
			break
		}
		ret[0][i] = 1
	}
	// 然后从第一个开始
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				ret[i][j] = 0
			} else {
				ret[i][j] = ret[i-1][j] + ret[i][j-1]
			}
		}
	}
	return ret[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
