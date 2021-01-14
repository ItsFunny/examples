/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-18 13:49 
# @File : of_剑指_Offer_13_机器人的运动范围.go
# @Description : 
# @Attention : 
*/
package offer

func movingCount(m int, n int, k int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i]=make([]int,n)
	}

	count := 0
	grid[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ((i > 0 && grid[i-1][j] == 1 )|| (j >= 0 && grid[i][j-1] == 1)) && sum(i)+sum(j) < k {
				count++
				grid[i][j] = 1
			}
		}
	}
	return count
}

func sum(i int) int {
	left := i % 10
	count := i / 10
	for count > 0 {
		left += count % 10
		count /= 10
	}
	return left
}
