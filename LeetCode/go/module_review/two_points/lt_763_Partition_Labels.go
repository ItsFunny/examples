/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-11 08:53 
# @File : lt_763_Partition_Labels.go
# @Description : 
# @Attention : 
*/
package two_points

/*
	双指针
从左到右遍历字符串，遍历的同时维护当前片段的开始下标 \textit{start}start 和结束下标 \textit{end}end，初始时 \textit{start}=\textit{end}=0start=end=0。

对于每个访问到的字母 cc，得到当前字母的最后一次出现的下标位置 \textit{end}_cend
​
 ，则当前片段的结束下标一定不会小于 \textit{end}_cend
​
 ，因此令 \textit{end}=\max(\textit{end},\textit{end}_c)end=max(end,end

当访问到下标 \textit{end}end 时，当前片段访问结束，当前片段的下标范围是 [\textit{start},\textit{end}][start,end]，长度为 \textit{end}-\textit{start}+1end−start+1，将当前片段的长度添加到返回值，然后令 \textit{start}=\textit{end}+1start=end+1，继续寻找下一个片段。

重复上述过程，直到遍历完字符串。

 */

func partitionLabels(S string) []int {
	last := make([]int, 26)
	for index, v := range S {
		last[v-'a'] = index
	}
	start, end := 0, 0
	result := make([]int, 0)
	for i := 0; i < len(S); i++ {
		end = partitionLabelsMax(end, last[S[i]-'a'])
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}
func partitionLabelsMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
