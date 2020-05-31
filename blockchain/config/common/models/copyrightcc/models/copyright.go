/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-19 13:54 
# @File : copyright.go
# @Description : 
# @Attention : 
*/
package models

import (
	"errors"
	"fmt"
	"examples/blockchain/config/common/parent"
)

type ItemUpload2ChainReqBO struct {
	parent.ItemUpload2BlockChainReqParent
}

type ItemGetCopyrightInfoReq struct {
	parent.GetItemCopyrightReqParent
}

func (this ItemGetCopyrightInfoReq) Validate() error {
	return nil
}

func (this ItemUpload2ChainReqBO) Validate() error {
	if this.TimeStamp == "" {
		return errors.New("参数时间戳不可为空")
	}
	if this.SerialNo == "" {
		return errors.New("参数时间戳的序列号不可为空")
	}
	if this.ItemCode == "" {
		return errors.New("参数作品码不可为空")
	}
	if this.ItemId <= 0 {
		return errors.New("参数itemId必须为有效值")
	}
	if this.ItemName == "" {
		return errors.New("参数作品名称不可为空")
	}
	if this.ItemCategory <= 0 {
		return errors.New(fmt.Sprintf("参数作品类型必须为有效值,{%d}", this.ItemCategory))
	}
	if this.ItemDescription == "" {
		return errors.New("参数作品描述不可为空")
	}
	if this.ItemType <= 0 {
		return errors.New(fmt.Sprintf("参数创作类别必须为有效值:{%d}", this.ItemType))
	}
	if this.ItemSpecificCategory <= 0 {
		return errors.New(fmt.Sprintf("参数作品具体风格必须为有效值,{%d}", this.ItemSpecificCategory))
	}
	if this.ItemStyle <= 0 {
		return errors.New(fmt.Sprintf("作品风格必须为有效值:{%d}", this.ItemStyle))
	}
	if this.ItemStatus <= 0 {
		return errors.New("作品上链时的状态必须为有效值")
	}

	if this.ItemAuthorName == "" {
		return errors.New("作品名称不可为空")
	}
	if this.CreateDate <= 0 {
		return errors.New("上传日期必须为有效值")
	}
	if this.CreateUser == "" {
		return errors.New("上传用户不可为空")
	}

	// // // 上传用户ID
	// // private Long createUserId;
	// CreateUserId int `json:"createUserId,string"`
	// // // 最近一次更新日期
	// // private long lastUpdateDate;
	// LastUpdateDate int64 `json:"lastUpdateDate,string"`
	// // // 最近一次更新用户
	// // private String lastUpdateUser;
	// LastUpdateUser string `json:"lastUpdateUser"`
	// // // 最近一次更新的用户ID
	// // private Long lastUpdateUserId;
	// LastUpdateUserId int `json:"lastUpdateUserId,string"`
	//

	return nil
}

type ItemUpload2ChainRespBO struct {
	parent.ItemUpload2BlockChainRespParent
}


type ItemGetCopyrightInfoResp struct {
	parent.GetItemCopyrightRespParent

}