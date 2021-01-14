/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-10 20:38 
# @File : validate.go
# @Description : 
# @Attention : 
*/
package raft

type IValidator interface {
	Valid()error
}
