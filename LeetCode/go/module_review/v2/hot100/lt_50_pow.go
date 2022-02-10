package hot100

// 实现一个2的n次幂
// 关键: 快速幂, 结果进行平方,而不是通过 原先的值进行平方
// 2^64=2->2^2->2^4->2^8->2^16->2^32 => 2^64
func myPow(x float64, n int) float64 {
	if n<0{
		return 1/quickMyPow(x,n)
	}
	return quickMyPow(x,n)
}
func quickMyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMyPow(x, n/2)
	if n%2 == 0 {
		// 直接通过结果进行平方
		return y * y
	}
	return y * y * x
}
