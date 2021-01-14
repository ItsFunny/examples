/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-04 08:30 
# @File : lt_6_ZigZag_Conversion.go
# @Description :
# @Attention : 
*/
package tree

import "strings"

/*
    以z字的形状打印出这个字符串
	通过 list + flag 来辅助即可
	flag的作用在于使得 可以让索引下标 从 0->1->2..->n-1 ==>  n-1->n-2->,,,->0
 */

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	sbs := make([]strings.Builder, 0)
	for i := 0; i < numRows; i++ {
		sbs = append(sbs, strings.Builder{})
	}

	flag := -1
	i := 0
	for _, c := range s {
		sbs[i].WriteByte(byte(c))
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	res := strings.Builder{}
	for i := 0; i < len(sbs); i++ {
		res.WriteString(sbs[i].String())
	}
	return res.String()
}
