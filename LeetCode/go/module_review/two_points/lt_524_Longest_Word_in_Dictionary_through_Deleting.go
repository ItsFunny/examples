/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-10 08:34 
# @File : lt_524_Longest_Word_in_Dictionary_through_Deleting.go
# @Description : 
# @Attention : 
*/
package two_points

import "strings"

/*
	判断是否是子序列即可
 */

func findLongestWord(s string, d []string) string {
	if len(s) == 0 {
		return ""
	}
	stringQuickSort(d)

	for _, v := range d {
		if isChildString(s, v) {
			return v
		}
	}
	return ""
}

func isChildString(s string, s2 string) bool {
	i := 0
	for j := 0; j < len(s); j++ {
		if s2[i] == s[j] {
			i++
		}
	}
	return i == len(s)-1
}

func stringQuickSort(d []string) {
	stringQSort(d, 0, len(d)-1)
}
func stringQSort(d []string, start, end int) {
	if start < end {
		paration := stringParation(d, start, end)
		stringQSort(d, start, paration)
		stringQSort(d, paration, end)
	}
}

func stringParation(data []string, start int, end int) int {
	standard := data[start]
	for start < end {
		for end > start && len(data[end]) > len(standard) || (len(data[end]) == len(standard) && strings.Compare(data[end], standard) > 0) {
			end--
		}
		data[start] = data[end]

		for start < end && len(data[start]) < len(standard) || (len(data[start]) == len(standard) && strings.Compare(data[start], standard) > 0) {
			start++
		}
		data[end] = data[start]
	}
	data[start] = standard
	return start
}
