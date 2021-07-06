/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/26 1:15 下午
# @File : lt_offer_替换空格.go
# @Description :
# @Attention :
*/
package v2

// 思路: 栈处理
func replaceSpace(s string) string {
	stack := make([]byte, 0)
	for i:=0;i< len(s);i++{
		v:=s[i]
		if v == ' ' {
			stack = append(stack, '%', '2', '0')
		} else {
			stack = append(stack, v)
		}
	}

	return string(stack)
}
