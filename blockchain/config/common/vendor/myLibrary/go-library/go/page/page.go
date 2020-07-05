/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-28 15:40 
# @File : page.go
# @Description : 
*/
package page

type BasePageReq struct {
	PageSize int    `form:"page_size" json:"page_size"`
	PageNum  int    `form:"CurrentPage" json:"CurrentPage"`
	OrderBy  string `form:"order_by" json:"order_by"`
}

func (this *BasePageReq)Cc() {
	this.PageNum = -99
}

type BasePageResp struct {
	PageNum    int `json:"page_num"`
	PageSize   int `json:"page_size"`
	TotalCount int `json:"total_count"`
}
