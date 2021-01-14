/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-15 12:58
# @File : lt_153_Find_Minimum_in_Rotated_Sorted_Array.go
# @Description :
假设按照升序排序的数组在预先未知的某个点上进行了旋转
( 例如，数组  [0,1,2,4,5,6,7] 可能变为  [4,5,6,7,0,1,2] )。 请找出其中最小的元素。
# @Attention :
	因为发生了反转,所以肯定不是升序的
*/
package half

import (
	"fmt"
	"testing"
)

func Test_findMin(t *testing.T) {
	// fmt.Println(findMin([]int{4,5,6,7,0,1,2}))
	fmt.Println(findMin([]int{3,1,2}))
}
