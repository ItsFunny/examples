/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-17 12:08 
# @File : _22_Generate_Parentheses.go
# @Description :
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

# @Attention :
*/
package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrack(n, n, &res, "")
	return res
}

func backtrack(left, right int, res *[]string, cur string) {
	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}

	if right < left {
		return
	}

	if left > 0 {
		backtrack(left-1, right, res, cur+"(")
	}

	if right > 0 {
		backtrack(left, right-1, res, cur+")")
	}
}