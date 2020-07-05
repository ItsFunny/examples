/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-15 14:11 
# @File : item.go
# @Description : 
# @Attention : 
*/
package parent

type ItemUploadReqParent struct {
	// 唯一的hash码
	Hash string `json:"hash"`
	// 作品分类: 是文字还是视频还是音乐等
	ItemType int `json:"itemType"`
	ItemID   int    `json:"itemId"`
	ItemName string `json:"itemName"`
	// 描述
	Description string `json:"description"`
	// 用户ID
	UserID      int    `json:"userId"`
	UserName    string `json:"userName"`
	UserPhone   string `json:"userPhone"`
	UserEmail   string `json:"userEmail"`
	UserAddress string `json:"userAddress"`
	Remark      string `json:"remark"`
	// 不同分类的作品传递过去的数据也是不同的,因而采用interface
	Data interface{} `json:"data"`
}
