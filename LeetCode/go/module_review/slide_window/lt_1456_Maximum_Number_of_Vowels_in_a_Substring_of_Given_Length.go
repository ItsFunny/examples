/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-24 09:00 
# @File : lt_1456_Maximum_Number_of_Vowels_in_a_Substring_of_Given_Length.go
# @Description : 
# @Attention : 
*/
package slide_window

/*
	在s中子串长度为k 的子串获取元音字母个数的最大值
	滑动窗口方式:
 */

func maxVowels(s string, k int) int {
	if len(s) == 0 || k <= 0 {
		return 0
	}
	m := make(map[byte]struct{})
	m['a'] = struct{}{}
	m['e'] = struct{}{}
	m['i'] = struct{}{}
	m['o'] = struct{}{}
	m['u'] = struct{}{}

	left, right := 0, 0
	count := 0
	result := 0
	for ; right < len(s); right++ {
		if _, exist := m[s[right]]; exist {
			count++
		}
		if right > k-1 {
			if _, exist := m[s[left]]; exist {
				count--
			}
			left++
		}
		// 滑动窗口,我其实不确定的是哪里需要进行compare 赋值
		result = maxVowelsMax(result, count)
	}
	return result
}
func maxVowelsMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxVowelsRangeSearch(m map[byte]struct{}, s string, left, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if _, exist := m[s[i]]; exist {
			count++
		}
	}
	return count
}
