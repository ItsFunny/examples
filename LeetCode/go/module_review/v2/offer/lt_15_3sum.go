/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/19 8:46 下午
# @File : lt_15_3sum.go
# @Description :
# @Attention :
*/
package offer

// a+b+c =0
// 固定住首位 ,b,c 头尾双指针
// 然后 a+b+c >0  则移到尾巴指针, 若 a+b+c<0 则移动首指针 ,==0 的话,还需要去除重复元素
func threeSum(nums []int) [][]int {
	// sort.Ints(nums)
	threeSumQSort(nums,0, len(nums)-1)
	var (
		a,b,c int
	)
	r := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		a = nums[i]
		// 去除重复位判断
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			b = nums[j]
			c = nums[k]
			if a+b+c > 0 {
				k--
			} else if a+b+c < 0 {
				j++
			} else {
				r = append(r, []int{a, b, c})
				// 然后去除重复解
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for k > j && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			}
		}
	}

	return r
}

func threeSumQSort(nums []int, left, right int) {
	if left < right {
		paration := threeSumQSortParation(nums, left, right)
		threeSumQSort(nums, left, paration)
		threeSumQSort(nums, paration, right)
	}
}

func threeSumQSortParation(nums []int, left, right int) int {
	standard := nums[left]
	for left < right {
		for right > left && nums[right] > standard {
			right--
		}
		nums[right] = nums[left]

		for left < right && nums[right] < standard {
			left++
		}
		nums[left] = nums[right]
	}
	nums[left] = standard
	return left
}
