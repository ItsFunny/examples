/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-11-13 16:26 
# @File : validater.go
# @Description : 
*/
package services

type IValidater interface {
	Validate() error
}
