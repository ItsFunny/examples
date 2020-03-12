/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-05-04 08:49 
# @File : validation.go
# @Description : 
*/
package utils

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// FIXME
// 手机号校验
func ValidatePhone(phone string) bool {
	// 手机号
	regExp := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	// 固话
	// regExp2 := "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\d{8}$|^0\\d{2,3}-?\\d{7,8}$"
	regExp2 := "^(0\\d{2,3}-?)d?\\d{7,8}$|\\d{7,8}"
	regExp3 := "(\\d{2,5}-\\d{7,8}(-\\d{1,})?)|(13\\d{9})|(159\\d{8})"
	m1, _ := regexp.MatchString(regExp, phone)
	m2, _ := regexp.MatchString(regExp2, phone)
	m3, _ := regexp.MatchString(regExp3, phone)
	return m1 || m2 || m3
}

// 生成随机短信数字验证码

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb string
	for i := 0; i < width; i++ {
		sb += strconv.Itoa(int(numeric[ rand.Intn(r) ]))
	}
	return sb
}
