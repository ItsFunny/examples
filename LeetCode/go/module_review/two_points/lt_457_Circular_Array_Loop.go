/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-09 09:50 
# @File : lt_457_Circular_Array_Loop.go
# @Description : 
# @Attention : 
*/
package two_points

/*
	循环数组判断是否有环
	快慢指针实现
	有环:
		环中的元素,所有符号都一致
		快/慢指针对应的上一个元素
		循环长度为1:既[8,2],移动一直都是原来的位置,所以也是无环的
 */

func circularArrayLoop(nums []int) bool {
	if len(nums) <= 1 || (len(nums) == 2 && nums[0]*nums[1] < 0) {
		return false
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]==0{
			continue
		}
		lastJ, lastK, j, k := 0, 0, i, i
		for {
			lastJ = j
			j = (j + nums[j] + 5000*len(nums)) % len(nums)
			if nums[lastJ]*nums[j] < 0 || nums[j] == 0 || lastJ == j {
				setZero(nums, j)
				break
			}
			lastK = k
			k = (k + nums[k] + 5000*len(nums)) % len(nums)
			if nums[k]*nums[lastK] < 0 || nums[k] == 0 || lastK == k {
				setZero(nums, k)
				break
			}
			lastK = k
			k = (k + nums[k] + 5000*len(nums)) % len(nums)
			if nums[k]*nums[lastK] < 0 || nums[k] == 0 || lastK == k {
				setZero(nums, k)
				break
			}
			if j == k {
				return true
			}
		}
	}
	return false
}

func setZero(nums []int, i int) {
	for {
		j := (i + nums[i] + 5000*len(nums)) % len(nums)
		if nums[j] == 0 || nums[i]*nums[j] < 0 {
			nums[i] = 0
			break
		}
		nums[i] = 0
		i = j
	}
}
