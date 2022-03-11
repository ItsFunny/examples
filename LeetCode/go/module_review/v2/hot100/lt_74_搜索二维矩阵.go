package hot100

// 关键:
// 1. 从题目出发,所以,要么是从左下开始,要么是从右上开始
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	// 从左下开始

	for i, j := len(matrix)-1, 0; i >= 0 && j < len(matrix[i]); {
		// 说明需要往小的移
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] == target {
			return true
		} else {
			// 说明需要往大的移
			j++
		}
	}

	return false
}
