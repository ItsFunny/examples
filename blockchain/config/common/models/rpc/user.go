/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-19 14:11 
# @File : user.go
# @Description : 
# @Attention : 
*/
package rpc

import (
	"examples/blockchain/config/common/models"
	"examples/blockchain/config/common/parent"
)

type RPCUserUploadReq struct {
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
	Birthday int64 `json:"birthday,string"`
	// 地址
	Address string `json:"address"`
}

type RPCUserUploadResp struct {
	*models.RPCLogicBaseResp
	// 身份唯一码
	DNA string `json:"dna"`
	// 私钥: hex 编码过后的私钥
	PrvKey string `json:"prvKey"`
}

// 更新
type RPCUserUpdateReq struct {
	UserDNA string `json:"userDna"`
	UserID  int    `json:"userId,string"`
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
type RPCUserUpdateResp struct {
}

type RPCUserGetInfo struct {
	parent.UserGetInfoParent
}

type RPCUserGetInfoResp struct {
	models.RPCLogicBaseResp
	parent.UserGetInfoRespParent

	// 区块编号, 既区块高度
	BlockNumber uint64 `json:"blockNumber"`
	// 交易地址,既数据保存在哪个block中
	TransactionBlockAddress string `json:"transactionBlockAddress"`

	Amount float64 `json:"amount"`
}

type RPCUserGetCoinAmountReq struct {
	parent.UserGetCoinAmoutParent
}

type RPCUserGetCoinAmountResp struct {
	models.RPCLogicBaseResp
	parent.CoinGetUserCoinAmountRespParent
}

type RPCUserUpdatePwdReq struct {
	parent.UserUpdatePwdReqParent
}

type RPCUserUpdatePwdResp struct {
	parent.UserUpdatePwdRespParent
}
