/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-21 08:31 
# @File : lt_424_Longest_Repeating_Character_Replacement.go
# @Description : 
# @Attention : 
*/
package two_points

/*
	替换字符串,然后找到最长子串
 */
func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	if k >= len(s) {
		return len(s)
	}
	nums := [26]int{}
	left, right := 0, 0
	historyMaxCount := 0

	for ;right < len(s) ;right++{
		index:=s[right]-'A'
		nums[index]++
		historyMaxCount=max_424(historyMaxCount,nums[index])
		if (right-left+1)>(historyMaxCount+k){
			nums[s[left]-'A']--
			left++
		}
	}
	return len(s)-left
}
func max_424(a,b int)int{
	if a>b{
		return a
	}
	return b
}
