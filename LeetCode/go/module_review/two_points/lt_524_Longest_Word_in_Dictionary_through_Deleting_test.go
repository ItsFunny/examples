/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-10 08:34
# @File : lt_524_Longest_Word_in_Dictionary_through_Deleting.go
# @Description :
# @Attention :
*/
package two_points

import (
	"fmt"
	"testing"
)

func Test_findLongestWord(t *testing.T) {
	s := "abpcplea"
	d := []string{"ale", "apple", "monkey", "plea"}
	word := findLongestWord(s, d)
	fmt.Println(word)

}
