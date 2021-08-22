/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/22 7:55 上午
# @File : lt_22_括号生成.go
# @Description :
# @Attention :
*/
package offer



// 关键: 回溯算法,深度优先遍历(dfs)
func generateParenthesis(n int) []string {
	var r []string
	if n == 0 {
		return nil
	}
	generateParenthesisDfs("", 0, 0, n,&r)
	return r
}
func generateParenthesisDfs(str string, leftCount int, rightCount int, n int,res *[]string) {
	if leftCount == rightCount && leftCount == n {
		*res = append(*res, str)
		return
	}
	if leftCount < rightCount {
		return
	}
	if leftCount < n {
		generateParenthesisDfs(str+"(", leftCount+1, rightCount, n,res)
	}
	if rightCount < n {
		generateParenthesisDfs(str+")", leftCount, rightCount+1, n,res)
	}
}
