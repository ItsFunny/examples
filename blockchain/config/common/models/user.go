/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-16 09:43 
# @File : user.go
# @Description : 
# @Attention : 
*/
package models

import (
	"errors"
	"fmt"
	"vlink.com/v2/vlink-common/parent"
)

// 用户上链并且初始化积分信息
type BCUploadUserAndCoinReq struct {
	BCUploadUserReq
}

type BCUploadUserAndCoinResp struct {
	// 身份唯一码
	DNA string `json:"dna"`
	// 私钥: hex 编码过后的私钥
	PrvKey string `json:"prvKey"`
}

type BCUploadUserReq struct {
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

type BCUploadUserResp struct {
	// 身份唯一码
	DNA string `json:"dna"`
	// 私钥: hex 编码过后的私钥
	PrvKey string `json:"prvKey"`
	// 2020-01-03 add
	// 用于存储积分的钱包地址
	CoinAddress string `json:"coinAddress"`
}

func (r BCUploadUserReq) Validate() error {
	if r.UserID <= 0 {
		return errors.New("参数userId不可小于等于0")
	}
	if r.UserPwd == "" {
		return errors.New("密码不可为空")
	}
	if r.UploadTime <= 0 {
		return errors.New(fmt.Sprintf("上链时间错误,[%d]", r.UploadTime))
	}

	return nil
}

// 更新用户信息
type BCUpdateUserReq struct {
	// 用户DNA
	UserDNA string `json:"userDna"`
	BCUploadUserReq
}

type BCUpdateUserResp struct {
	// 当这个值不为0的时候,意味着需要更新积分,
	// 优化存储空间,将布尔变量移除
	UpdateCoin float64 `json:"updateCoin"`

	// 更新积分的地址,当用户创建的时候,这个值是必然存在的
	CoinWalletAddress string `json:"coinWalletAddress"`
}

// 获取用户信息
type BCGetUserInfoReq struct {
	UserId int    `json:"userId"`
	Dna    string `json:"dna"`
}

type BCUserGetCoinAmoutReq struct {
	parent.UserGetCoinAmoutParent
}

func (r BCUserGetCoinAmoutReq) Validate() error {
	return nil
}

type BCUserGetCoinAmountResp struct {
	parent.CoinGetUserCoinAmountRespParent
}

type BCGetUserInfoResp struct {
	parent.UserGetInfoRespParent

	// 积分总额
	Amount float64 `json:"amount"`
}

type BCWalletInfoReq struct {
	UserId int
}

func (r BCWalletInfoReq) Validate() error {
	if r.UserId <= 0 {
		return errors.New(fmt.Sprintf("用户id不可小于或者等于0,{%d}", r.UserId))
	}
	return nil
}

type BCWalletInfoResp struct {
	UserWalletParent
}

type UserWalletParent struct {
	UserDNA string `json:"userDna"`
	// 主私钥
	// PrvBytes []byte `json:"prvBytes"`
	// 主公钥
	// PubBytes []byte `json:"pubBytes"`
	// 是否server中心化存储
	ServerStorageKey bool `json:"serverStorageKey"`
	// 2020-01-06 update ,上述缺失子钥信息
	MainWallet *VlinkWallet
}

type BCUserUpdatePwdReqBO struct {
	parent.UserUpdatePwdReqParent
}

// TODO
func (this BCUserUpdatePwdReqBO) Validate() error {
	return nil
}

// TODO
func (this BCUserUpdatePwdReqBO) Encrypt(data ...interface{}) (interface{}, error) {
	return this, nil
}

func (this BCUserUpdatePwdReqBO) Decrypt(data ...interface{}) (interface{}, error) {

	return this, nil
}

type BCUserUpdatePwdRespBO struct {
	parent.UserUpdatePwdRespParent
}
