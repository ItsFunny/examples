/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-06 15:25 
# @File : user.go
# @Description : 
# @Attention : 
*/
package models

import (
	"fmt"
	"vlink.com/v2/vlink-common/parent"
)

type BCUserCoinUpdateUserInfoReq struct {
	// 用户DNA
	UserDNA string `json:"userDna"`

	UserID int `json:"userId,string"`
	// 用户密码,用于生成钱包
	UserPwd  string `json:"userPwd"`
	IDCard   string `json:"idCard"`
	RealName string `json:"realName"`
	Phone    string `json:"phone"`
	WeChatID string `json:"weChatId"`
	// 主密钥是否中心化存储,若true则加密存储到链上,否则返回给用户不存储,支持后续更新
	ServerStorageKey bool `json:"serverStorageKey"`
	// 上链时间
	UploadTime int64 `json:"uploadTime,string"`
	// 2020-01-02 add
	// 性别
	Gender int `json:"gender"`
	// 出生日期
	Birthday int64 `json:"birthday"`
	// 地址
	Address string `json:"address"`
}

func (r BCUserCoinUpdateUserInfoReq) Validate() error {
	fmt.Println("执行 [BCUserCoinUpdateUserInfoReq]的参数校验")
	return nil
}

type BCUserCoinUpdateUserInfoResp struct {
	// models.LogicBaseResp
}

type BCUserCoinUserGetInfoReq struct {
	parent.UserGetInfoParent
}
type BCUserCoinUserGetInfoResp struct {
	parent.UserGetInfoRespParent
}
