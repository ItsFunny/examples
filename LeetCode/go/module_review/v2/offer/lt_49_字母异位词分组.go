/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/31 8:49 上午
# @File : lt_49_字母异位词分组.go
# @Description :
# @Attention :
*/
package offer

// 参考:https://leetcode-cn.com/problems/group-anagrams/solution/zi-mu-yi-wei-ci-fen-zu-by-leetcode-solut-gyoc/
// 解题关键: 异位代表着字符串出现的个数肯定相同
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		node := [26]int{}
		for _, v := range str {
			node[v-'a']++
		}
		m[node] = append(m[node], str)
	}
	ret := make([][]string, 0)
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}
