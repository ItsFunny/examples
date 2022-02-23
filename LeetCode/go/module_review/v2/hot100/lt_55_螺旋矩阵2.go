package hot100

// 关键:
// 1. 从左到右赋值,再从上到下,再从右到左,再从下到上
// 2. index必须从1开始
func generateMatrix(n int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, make([]int, n))
	}
	left, right, up, down := 0, n-1, 0, n-1
	index := 1
	for index <= n*n {
		// 开始从左往右赋值
		for i := left; i <= right; i++ {
			ret[up][i] = index
			index++
		}
		up++
		// 开始从上往下赋值
		for i := up; i <= down; i++ {
			ret[i][right] = index
			index++
		}
		right--
		// 开始从右往左赋值
		for i := right; i >= left; i-- {
			ret[down][i] = index
			index++
		}
		// 开始从下往上赋值
		down--
		for i := down; i >= up; i-- {
			ret[i][left] = index
			index++
		}
		// 继续
		left++
	}
	return ret
}
