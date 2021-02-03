/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-21 09:20 
# @File : BaseServiceInitImpl.go
# @Description : 
*/
package baseImpl

import "bidchain/fabric/log"

type BaseServiceInitImpl struct {
	ReqID string
	Log   log.Logger
}

func (receiver *BaseServiceInitImpl) GetReqId() string {
	return receiver.ReqID
}
func (receiver *BaseServiceInitImpl) SetReqId(id string) {
	receiver.ReqID = id
}

func (receiver *BaseServiceInitImpl) SetLogger(l log.Logger) {
	receiver.Log = l
}
func (receiver *BaseServiceInitImpl) GetLogger() log.Logger {
	return receiver.Log
}
func NewBaseServiceInitImpl(reqID string) *BaseServiceInitImpl {
	iml := new(BaseServiceInitImpl)
	iml.Log=log.NewCommonBaseLoggerWithLog4go(reqID)
	return iml
}
