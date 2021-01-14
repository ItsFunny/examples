/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-11 09:33
# @File : lt_826_Most_Profit_Assigning_Work.go
# @Description :
# @Attention :
*/
package two_points

import (
	"fmt"
	"testing"
)

func Test_maxProfitAssignment(t *testing.T) {
	// [5,50,92,21,24,70,17,63,30,53]
	// [68,100,3,99,56,43,26,93,55,25]
	// [96,3,55,30,11,58,68,36,26,1]
	diff := []int{5,50,92,21,24,70,17,63,30,53}
	profit := []int{68,100,3,99,56,43,26,93,55,25}
	worker := []int{96,3,55,30,11,58,68,36,26,1}
	assignment := maxProfitAssignment(diff, profit, worker)
	fmt.Println(assignment)
}
