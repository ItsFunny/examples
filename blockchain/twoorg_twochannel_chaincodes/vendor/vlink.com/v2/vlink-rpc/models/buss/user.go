/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-19 09:52 
# @File : user.go
# @Description : 
# @Attention : 
*/
package buss

import "vlink.com/v2/vlink-common/models"

type BSUserUploadReq struct {
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

type BSUserUploadResp struct {
	models.LogicBaseResp
	// 身份唯一码
	DNA string `json:"dna"`
	// 私钥: hex 编码过后的私钥
	PrvKey string `json:"prvKey"`
}

// 更新用户信息
type BSUserUpdateReq struct {
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

type BSUserUpdateResp struct {
	models.LogicBaseResp
}