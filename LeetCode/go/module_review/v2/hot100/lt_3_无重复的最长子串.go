/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/7 9:02 下午
# @File : lt_3_无重复的最长子串.go
# @Description :
# @Attention :
*/
package hot100

// 关键: 双指针
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	ret := 1
	// 关键是: 1. 通过一个set保存byte出现的下标
	// 2. 双指针,右指针不停的向右滑动,当发现对应的元素已经在set中出现过的时候,则挪动左指针到出现过的地方
	// 双指针, 右指针不停的移动
	m := make(map[byte]int)
	for ;right < len(s);right++ {
		index, exist := m[s[right]]
		if exist {
			// 更新左指针的最大下标
			left = lengthOfLongestSubstringMax(left, index)
		}
		// 更新最大值,因为最大值是通过 右指针-左指针+1 的
		ret = lengthOfLongestSubstringMax(ret, right-left+1)
		// 这一步是关键,要记得,要更新右指针的下标,因为此时也是相当于更新左指针
		m[s[right]] = right + 1
	}

	return ret
}
func lengthOfLongestSubstringMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
