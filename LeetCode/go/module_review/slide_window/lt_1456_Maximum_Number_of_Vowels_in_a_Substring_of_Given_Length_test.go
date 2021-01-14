/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-24 09:00
# @File : lt_1456_Maximum_Number_of_Vowels_in_a_Substring_of_Given_Length.go
# @Description :
# @Attention :
*/
package slide_window

import (
	"fmt"
	"testing"
)

func Test_maxVowels(t *testing.T) {
	s := "leetcode"
	k := 3
	vowels := maxVowels(s, k)
	fmt.Println(vowels)
}

