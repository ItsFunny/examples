package hot100

// 二叉搜索树的定义:
// left<root<right
// 关键: 背公式
// g(n)=g(0)*g(n-1)+g(1)*g(n-2)+...g(n-1)*g(0)
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
