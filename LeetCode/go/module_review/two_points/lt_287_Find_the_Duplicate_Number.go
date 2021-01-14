/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-17 15:41 
# @File : lt_287_Find_the_Duplicate_Number.go
# @Description : 
# @Attention : 
*/
package two_points

/*
	找到重复的数字
	可以直接用位运算获取
 */
func findDuplicate(nums []int) int {
	flag := 0
	for _, n := range nums {
		if flag&n >= n {
			return n
		}
		flag |= n
	}
	return -1
}
