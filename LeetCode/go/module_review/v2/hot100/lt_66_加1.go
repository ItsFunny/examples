package hot100

// 关键:
// 判断是否需要额外加一即可
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i]!=0{
			return digits
		}
	}
	// 当运行到这里的时候,表明之前的数字都是 99999
	ret:=make([]int, len(digits)+1)
	ret[0]=1
	return ret
}
