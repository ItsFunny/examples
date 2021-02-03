/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/2/1 14:31
# @File : result.go
# @Description :
# @Attention :
*/
package utils

import (
	"bidchain/fabric/bserror"
	"bidchain/http_framework/protocol"
)

func ReturnWithError(command protocol.ICommand,err bserror.IBSError){
	command.SetErrCode(err.GetCode())
	command.SetErrDesc(err.Error())
}