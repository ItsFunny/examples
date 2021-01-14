/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-23 09:37
# @File : lt_1438_Longest_Continuous_Subarray_With_Absolute_Diff_Less_Than_or_Equal_to_Limit.go
# @Description :
# @Attention :
*/
package slide_window

import (
	"fmt"
	"testing"
)

func Test_longestSubarray(t *testing.T) {
	nums := []int{10,1,2,4,7,2}
	limit := 5
	subarray := longestSubarray(nums, limit)
	fmt.Println(subarray)
}
