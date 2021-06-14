/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/29 10:56 下午
# @File : palindrome_linked_list_test.go.go
# @Description :
# @Attention :
*/
package v2

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	node := CreateNodeBy(1, 2, 2, 1)
	fmt.Println(isPalindrome(node))
}

func Test_isPalindrome2(t *testing.T) {
	node := CreateNodeBy(1, 0, 1)
	fmt.Println(isPalindrome(node))
}

func TestStupid(t *testing.T) {
	fmt.Println(true || false)
}
