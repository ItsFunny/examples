package hot100

// 关键:
// 动态规划: f(i,j)=f(i−1,j)+f(i,j−1)
// 当走到 i,j 位置的时候,可以是从i-1,j 过来 ,也可以是从 i,j-1 过来
// 注意点: 对于边界的情况(既i=0,j=0),此时只有一种可能(如i=0,只能从 0,j-1过来)
func uniquePaths(m int, n int) int {
	paths := make([][]int, m)
	// 初始化 边界路径都为1
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
		paths[i][0] = 1
	}
	for i := 0; i < n; i++ {
		paths[0][i] = 1
	}

	// 然后下标从1开始进行累加
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	return paths[m-1][n-1]
}
