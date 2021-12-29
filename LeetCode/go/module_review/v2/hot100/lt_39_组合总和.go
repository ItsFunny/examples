/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/29 9:24 上午
# @File : lt_39_组合总和.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 回溯算法,想到回溯,直接无脑 dfs
// 当target ==0 的时候,代表是符合条件的其中一个解
// 回溯算法的基本模板是: 1. 先将推出条件全都写上 2. 直接进入下一个循环 3. 开始剪枝
func combinationSum(candidates []int, target int) [][]int {
	var single []int
	var ret [][]int
	var dfs func(left, index int)

	dfs = func(left int, index int) {
		//1. 先将推出条件全都写上
		if index == len(candidates) {
			return
		}
		if left == 0 {
			// 表明符合条件,可以作为其中一个结果
			ret = append(ret, append([]int(nil), single...))
			return
		}
		//2. 直接进入下一个循环
		// 然后直接跳到下一个
		dfs(left, index+1)

		// 3. 开始剪枝
		if left-candidates[index] >= 0 {
			single = append(single, candidates[index])
			dfs(left-candidates[index], index)
			// 剪枝
			single = single[:len(single)-1]
		}
	}
	dfs(target, 0)
	return ret
}
