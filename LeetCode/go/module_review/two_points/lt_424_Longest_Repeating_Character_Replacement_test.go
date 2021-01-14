/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-21 08:31
# @File : lt_424_Longest_Repeating_Character_Replacement.go
# @Description :
# @Attention :
*/
package two_points

import (
	"fmt"
	"testing"
)

func Test_characterReplacement(t *testing.T) {
	s:= "AABABBA"
	k:=1
	replacement := characterReplacement(s, k)
	fmt.Println(replacement)
}
