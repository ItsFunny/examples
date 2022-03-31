package hot100

// 关键
// 无他: 死记硬背
// 1. for 循环, 右移缩小index位
// 2. 异或^ 缩小的值即可
func grayCode(n int) []int {
	ret := make([]int, 1<<n)
	for index := range ret {
		ret[index] = index>>1 ^ index
	}
	return ret
}
