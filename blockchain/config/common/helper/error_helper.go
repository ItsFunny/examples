/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-30 22:35 
# @File : error_helper.go
# @Description : 
# @Attention : 
*/
package helper

import (
	"examples/blockchain/config/common/constants"
	error2 "examples/blockchain/config/common/error"
	"examples/blockchain/config/common/models"
)

func GetErrorDesc(err error2.IVlinkError) string {
	code := err.GetCode()
	str := "系统错误"
	if code&constants.OUT_PUT_ERROR_CODE > 0 {
		str = err.GetMsg()
		// FIXME 直接将结果持久化到内存中的map,一个for循环直接获取即可
	} else if code&constants.DB_ERROR_CODE > 0 {
		str = "数据库错误"
	} else if code&constants.ARGUMENT_ERROR_CODE > 0 {
		str = "参数错误"
	} else if code&constants.CONFIG_ERROR_CODE > 0 {
		str = "配置错误"
	} else {
		str = "未知错误"
	}
	return str
}

func HandlerError(service models.ILogicBaseService, e error2.IVlinkError) {
	service.SetLogicCode(int(e.GetCode()))
	service.SetLogcMsg(GetErrorDesc(e))
}


func NewErr(str string,args ...interface{}){


}
