/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-05-02 09:38 
# @File : bussiness.go
# @Description :    业务方面的util
*/
package utils

import (
	"fmt"
	"myLibrary/go-library/go/constants"
)


// 用于更新
// make sure is not empty
// 用于对2个model之间进行比较,同时返回的map用于更新(gorm会自动对于key-value的map更新)
func CompareAndSwap(descriptions, colunms []string, previousValues, currentValues []interface{}) (string, map[string]interface{}) {
	m := make(map[string]interface{})
	l := len(colunms)
	diff := ""
	for i := 0; i < l-1; i++ {
		if previousValues[i] != currentValues[i] {
			diff += fmt.Sprintf(constants.BASE_NAME_MODIFY_TEMPLATE+" , ", descriptions[i], previousValues[i], currentValues[i])
			m[colunms[i]] = currentValues[i]
		}
	}
	if previousValues[l-1] != currentValues[l-1] {
		diff += fmt.Sprintf(constants.BASE_NAME_MODIFY_TEMPLATE+" . ", descriptions[l-1], previousValues[l-1], currentValues[l-1])
		m[colunms[l-1]] = currentValues[l-1]
	}

	return diff, m
}
