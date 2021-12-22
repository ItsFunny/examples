/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/22 9:16 上午
# @File : lt_31_下一个排列.go
# @Description :
# @Attention :
*/
package hot100

import "sort"

// 关键：
// 1. 从后往前遍历,找到 前一个数比当前数小的值 i,代表着是可以有下一个排列的
// 2. 就算找到之后,也是不可以直接 i+1和i 替换位置的,因为需要找到 nums[i+1:]大于nums[i]的最小值
// 3. 在nums[i+1:] 找到大于nums[i]的最小值之后,替换,替换之后,对nums[i+1:]进行升序排列,确保是最小的
/*

nums = [1,2,7,4,3,1],
第一步: 倒序遍历数组, 找出第一组: 前一个数比后一个数小的两个数, 即[2, 7]
2所处的这个位置就是需要找出比它稍微大的数的位置;
我们从[7,4,3,1]中找出比2大的数中的最小值, 也就是3, 找到后跟2交换即可;; 当然了, 如果没找到的话, 直接跳到第5步, 直接升序排列输出.
目前nums=[1,3,7,4,2,1], 不用我说你们也看出来还不算下一个排列
对3后面的数, 升序排列, 即最终结果: nums = [1,3,1,2,4,7]
*/
func nextPermutation(nums []int)  {
	if len(nums)==0{
		return
	}
	firstIndex:=-1
	// 第一步
	for i:= len(nums)-2;i>=0;i--{
		if nums[i]<nums[i+1]{
			firstIndex=i
			break
		}
	}
	if firstIndex==-1{
		lt31reverse(nums)
		return
	}
	// 第二步 ,在nums[i+1:] 找大于nums[i]的最小值
	secondIndex:=firstIndex+1
	min:=nums[firstIndex+1]
	for i:=firstIndex;i< len(nums);i++{
		if nums[i]<min && nums[i]>nums[firstIndex]{
			min=nums[i]
			secondIndex=i
		}
	}

	// 第三步,对nums[i+1:]进行升序排序
	nums[firstIndex],nums[secondIndex]=nums[secondIndex],nums[firstIndex]
	sort.Ints(nums[firstIndex+1:])
}

func lt31reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}