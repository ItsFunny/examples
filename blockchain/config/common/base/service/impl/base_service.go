/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-17 13:38 
# @File : base_service.go
# @Description : 
# @Attention : 
*/
package impl

import "myLibrary/go-library/go/base/service"

type VlinkBaseServiceImpl struct {
	*service.BaseServiceImpl
}

func NewVlinkBaseServiceImpl()*VlinkBaseServiceImpl{
	v:=new(VlinkBaseServiceImpl)
	v.BaseServiceImpl = service.NewBaseServiceImplWithLog4goLogger()
	return v
}