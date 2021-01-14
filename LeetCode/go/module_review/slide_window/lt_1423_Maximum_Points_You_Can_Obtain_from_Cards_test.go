/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-19 19:54
# @File : lt_1423_Maximum_Points_You_Can_Obtain_from_Cards.go
# @Description :
# @Attention :
*/
package slide_window

import (
	"fmt"
	"testing"
)

func Test_maxScore(t *testing.T) {
	scors := []int{1, 2, 3, 4, 5, 6, 1}
	max := 3
	score := maxScore(scors, max)
	fmt.Println(score)
}
