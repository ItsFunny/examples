package hot100

// 关键,反向遍历
// 1. 先找到第一个不为空的单词
func lengthOfLastWord(s string) int {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	ret := 0
	for i := index; i >= 0 && s[i] != ' '; i-- {
		ret++
	}
	return ret
}
