package hot100

// 关键: 2次翻转
// 1. 先按水平线翻转 (既中间行)
// 2. 按对角线翻转,此时的下标的变化: (2,1) => (1,2) (注意,当对角线翻转的时候,第二个for循环的退出条件不是全长度)
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	l := len(matrix)
	// 1. 先按水平线进行翻转
	mid := l >> 1
	for i := 0; i < mid; i++ {
		matrix[i], matrix[l-i-1] = matrix[l-i-1], matrix[i]
	}
	// 2. 按对角线翻转
	for i := 0; i < l; i++ {
		for j := 0; j <i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
