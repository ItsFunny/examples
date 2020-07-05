/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-13 15:20 
# @File : base_controller.go
# @Description : 
# @Attention : 
*/
package controller

import (
	"myLibrary/go-library/go/base/controller"
	"net/http"
	"examples/blockchain/config/common/constants"
	"examples/blockchain/config/common/models"
)

type VlinkBaseController struct {
	controllers.WebBaseController
}

func (this *VlinkBaseController)ReturnSuccess(resp models.InvokeResp)(err error){
	resp.LogicCode=constants.SUCCESS
	resp.LogicMsg="SUCCESS"
	this.Ctx.Output.Status = http.StatusOK
	this.Data["json"] = resp
	this.ServeJSON()

	return
}
