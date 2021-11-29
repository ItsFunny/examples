/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/9/7 8:50 上午
# @File : lt_75_颜色分类.go
# @Description :
# @Attention :
*/
package offer

// 题目关键: 最终结果是最后的值是从小到大
// 解题关键: 双指针: 右指针将后面的2的可能情况全部移到前面
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := 0; i <= right; i++ {
		// 使得如果右指针也是2,则后面的2全部往前移动
		for ; i <= right && nums[i] == 2; right-- {
			nums[i], nums[right] = nums[right], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}
