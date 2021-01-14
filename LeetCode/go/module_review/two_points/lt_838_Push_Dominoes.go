/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-12 09:17 
# @File : lt_838_Push_Dominoes.go
# @Description : 
# @Attention : 
*/
package two_points

/*
	多米诺骨牌
	受力法分析
 */
func pushDominoes(dominoes string) string {
	if len(dominoes) == 0 {
		return ""
	}
	forces := make([]int, len(dominoes))
	force := 0
	N := len(dominoes)
	for index, v := range dominoes {
		if v == 'R' {
			force = N
		} else if v == 'L' {
			force = 0
		} else {
			force =pushDominoesMax()
		}
	}
}

func pushDominoesMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
