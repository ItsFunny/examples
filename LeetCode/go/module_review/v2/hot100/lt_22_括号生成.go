/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/16 9:15 上午
# @File : lt_22_括号生成.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 回溯算法,深度优先遍历(dfs)
// 参考: https://leetcode-cn.com/problems/generate-parentheses/solution/hui-su-suan-fa-by-liweiwei1419/
// 对于回溯算法,都可以使用二叉树来想象代入
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	ret := make([]string, 0)
	generateParenthesisDfs("", n, 0, 0, &ret)
	return ret
}

func generateParenthesisDfs(str string, limit int, leftCount, rightCount int, ret *[]string) {
	if leftCount == limit && rightCount == limit {
		*ret = append(*ret, str)
		return
	}
	// 因为 ( ) 是成对出现的, ( 必然是要比 ) 个数要多的
	if leftCount < rightCount {
		return
	}
	if leftCount < limit {
		generateParenthesisDfs(str+"(", limit, leftCount+1, rightCount, ret)
	}
	if rightCount < limit {
		generateParenthesisDfs(str+")", limit, leftCount, rightCount+1, ret)
	}
}
