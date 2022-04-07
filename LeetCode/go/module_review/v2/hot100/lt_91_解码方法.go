package hot100

// 关键:
// 一大堆边界条件判断,dp的时候的长度都是len+1,返回值都是返回len
// 状态转移方程:
// f(i)=f(i-1) + f(i-2)
// 当选择一个数的时候,f(i)+=f(i-1)
// 当选择2个数的时候,f(i)=f(i-1)+f(i-2)
// 解码的时候,可以由1个数解码,也可以是2个数合在一起解码
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] > '0' && ((s[i-2]-'0')*10+(s[i-1]-'0')) <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}