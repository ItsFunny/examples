/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-10 20:38 
# @File : error.go
# @Description : 
# @Attention : 
*/
package raft

import (
	"errors"
	"fmt"
)

func ErrorF(msg string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(msg, args...))
}
