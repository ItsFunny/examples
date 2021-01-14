/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-04 14:52 
# @File : url.go
# @Description : 
*/
package utils

import (
	"errors"
	"net/url"
	"strings"
)

// 获取文件后缀
// func GetFileSuffix(str string) (string, error) {
// 	index := strings.LastIndex(str, ".")
// 	if index == -1 {
// 		return "", errors.New("格式不正确")
// 	}
// 	l := len(str)
// 	if index == l-1 {
// 		return "", nil
// 	}
// 	suffix := SubStringBetween(str, index, l-2)
// 	suffix = strings.ToLower(suffix)
//
// 	return suffix, nil
// }

func GetLowerSuffixFromUrl(str string) (string, error) {
	parse, e := url.Parse(str)
	if nil != e {
		return "", e
	}
	str = parse.Path
	index := strings.LastIndex(str, ".")
	if index == -1 {
		return "", errors.New("格式不正确")
	}
	l := len(str)
	if index == l-1 {
		return "", nil
	}
	suffix := SubStringBetween(str, index, l)
	suffix = strings.ToLower(suffix)
	return suffix, nil
}
