package hot100

// 关键: dfs
// 1. 双重遍历,当发现了与第一个字符匹配之后, 就围绕这个i,j 进行dfs深搜匹配
func exist(board [][]byte, word string) bool {
	var dfs func(i, j, k int) bool
	// 标记号,防止重复计算
	flags := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		flags[i] = make([]bool, len(board[i]))
	}
	dfs = func(i, j, k int) bool {
		// 退出条件:
		if k == len(word) {
			// 说明之前的都符合条件了,则直接true
			return true
		}
		// 越界了
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || flags[i][j] {
			return false
		}
		if board[i][j] != word[k] {
			return false
		}
		// 标记,防止重复计算
		flags[i][j] = true
		// 最后要裁剪,将数据还原
		defer func() {
			flags[i][j] = false
		}()
		// 然后继续dfs,并且要防止重复计算
		return dfs(i, j-1, k+1) || dfs(i, j+1, k+1) || dfs(i-1, j, k+1) || dfs(i+1, j, k+1)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
