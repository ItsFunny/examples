package hot100

// 关键: 标记数组的方式
// 1. 先遍历二维数组, 如果有一个为0,则将行和列都标记为0
func setZeroes(matrix [][]int) {
	cols := make([]bool, len(matrix))
	rows := make([]bool, len(matrix[0]))
	for col, v := range matrix {
		for row, vv := range v {
			if vv == 0 {
				cols[col] = true
				rows[row] = true
			}
		}
	}

	// 然后开始遍历
	for col, v := range matrix {
		for row, _ := range v {
			if cols[col] || rows[row] {
				matrix[col][row] = 0
			}
		}
	}

}
