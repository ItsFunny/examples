/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-01 07:52 
# @File : lt_96_Unique_Binary_Search_Trees.go
# @Description : 
# @Attention : 
*/
package tree

func numTrees(n int) int {
	records := make([]int, n+1)
	records[0]=1

	return h(n, records)
}

func h(n int, record []int) int {
	if n==0 || n==1{
		return 1
	}
	if record[n] > 0 {
		return record[n]
	}
	for i := 1; i <= n; i++ {
		record[n] += h(i-1, record) * h(n-i, record)
	}

	return record[n]
}
