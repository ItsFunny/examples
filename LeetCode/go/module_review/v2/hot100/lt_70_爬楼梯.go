package hot100

// 关键: 动态规划
// f(n)=f(n-1)+f(n-2)
func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[i] = 1
			continue
		}
		if i == 2 {
			dp[i] = 2
			continue
		}
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
