/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 15:46 
# @File : IVlinkBaseService.go
# @Description : 
# @Attention : 
*/
package service

import (
	"myLibrary/go-library/go/base/models"
	"myLibrary/go-library/go/base/service"
)

type IVlinkBaseService interface {
	service.IBaseService
}




type IVlinkBaseRespService interface {
	models.IBaseResponse
}