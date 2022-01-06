/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/4 8:48 上午
# @File : lt_40_组合总和2.go
# @Description :
# @Attention :
*/
package hot100

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	var (
		single []int
		ret    [][]int
		dfs    func(left, index int)
	)
	dfs = func(left, index int) {
		if index == len(candidates) {
			return
		}
		if left == 0 {
			// 去重
			ret = append(ret, append([]int{}, single...))
			return
		}
		// 直接进入下一个
		dfs(left,index+1)

		// 开始剪枝
		if left-candidates[index]>=0{
			single=append(single,candidates[index])

		}
	}
	return nil
}
