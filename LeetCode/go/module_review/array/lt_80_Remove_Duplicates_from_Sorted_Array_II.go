/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-14 08:26 
# @File : lt_80_Remove_Duplicates_from_Sorted_Array_II.go
# @Description : 
# @Attention : 
*/
package array

/*
	删除重复次数大于2次的元素,使得最多重复2次
	注意点: 不可以申请额外的内存空间
	思路: 一个控制下标,一个控制下标对应的次数,因为是有序的,所以我们只需要顺序匹配即可
 */

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	removeCurrentIndex := func(nums *[]int, index int) []int {
		l := len(*nums)
		tNums := *nums
		for i := index; i < len(*nums)-1; i++ {
			tNums[i] = tNums[i+1]
		}
		tNums = tNums[:l-1]
		return tNums
	}
	index := 1
	count := 1
	for ; index < len(nums); {
		if nums[index-1] == nums[index] {
			count++
			if count > 2 {
				// 删除这个元素
				nums = removeCurrentIndex(&nums, index)
				count--
			} else {
				index++
			}
		} else {
			count = 1
			index++
		}
	}
	return len(nums)
}
