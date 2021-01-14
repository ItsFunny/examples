/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-15 12:50 
# @File : lt_278_First_Bad_Version.go
# @Description :
假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。 你可以通过调用  bool isBadVersion(version)  接口来判断版本号
version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
# @Attention :
	既第一个错误版本之后都是错的,找到第一个正确版本即可
*/
package half

func firstBadVersion(n int) int {
	start := 0
	end := n
	for start <= end {
		mid := start + (end-start)>>1
		if isBadVersion(mid) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}
func isBadVersion(n int) bool {
	return true
}
