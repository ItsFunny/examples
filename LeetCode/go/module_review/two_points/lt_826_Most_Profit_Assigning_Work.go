/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-11 09:33 
# @File : lt_826_Most_Profit_Assigning_Work.go
# @Description : 
# @Attention : 
*/
package two_points

/*
	贪心算法:
	最大利润:
	则每个员工最大利润,则总体最大利润
	对difficulty 排序 ,难度越大,利润越大,
	优化: 对difficulty 也排序,二分查找,这样能从O(n^2) 退化为 O(nlogn)
 */

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	if len(difficulty) == 0 || len(profit) == 0 || len(worker) == 0 {
		return 0
	}
	// maxProfitAssignmentQSort(difficulty, 0, len(difficulty)-1)
	// maxProfitAssignmentQSort(profit, 0, len(profit)-1)
	maxProfitAssignmentQSort(worker, 0, len(worker)-1)

	result := 0
	for i := 0; i < len(worker); i++ {
		//
		perWorkerBest := 0
		for j := 0; j < len(difficulty); j++ {
			if worker[i] >= difficulty[j] {
				perWorkerBest = maxProfitAssignmentMax(perWorkerBest, profit[j])
			}
		}
		result += perWorkerBest
	}
	return result
}

func maxProfitAssignmentMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfitAssignmentQSort(difficulty []int, start, end int) {
	if start < end {
		paration := maxProfitAssignmentQSortParation(difficulty, start, end)
		maxProfitAssignmentQSort(difficulty, start, paration)
		maxProfitAssignmentQSort(difficulty, paration+1, end)
	}
}

func maxProfitAssignmentQSortParation(ints []int, start int, end int) int {
	standard := ints[start]
	for start < end {
		for ; end > start && ints[end] >= standard; end-- {
		}
		ints[start] = ints[end]
		for ; start < end && ints[start] <= standard; start++ {
		}
		ints[end] = ints[start]
	}
	ints[start] = standard
	return start
}
