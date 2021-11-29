/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/29 5:37 下午
# @File : lt_41_缺失的最小正整数.go
# @Description :
# @Attention :
*/
package offer

// 参考: https://leetcode-cn.com/problems/first-missing-positive/solution/tong-pai-xu-python-dai-ma-by-liweiwei1419/
// 关键:
// 将1放在下标0的位置,2放在下标1的位置 依次类推
// 然后最终结果,遍历的时候,如果发现当前的下标与i+1不匹配,则返回这个值
func firstMissingPositive(nums []int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		// 因为是正整数,所以要nums[i]>0,并且可能为 7,8,远超过长度的,所以对应的值也要小于长度值
		// FIXME:为什么这里要多出一重判断,判断是否交换后的相等 : 是为了防止死循环
		for nums[i] > 0 && nums[i] <= l && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for index,_ := range nums {
		if nums[index] > 0 && nums[index] == index+1 {
			return index + 1
		}
	}
	return l + 1
}
