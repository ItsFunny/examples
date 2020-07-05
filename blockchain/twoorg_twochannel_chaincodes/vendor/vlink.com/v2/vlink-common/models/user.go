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
)

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
	*LogicBaseResp
	// 身份唯一码
	DNA string `json:"dna"`
	// 私钥: hex 编码过后的私钥
	PrvKey string `json:"prvKey"`
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

func (r BCUpdateUserReq) Validate() error {
	if r.UserDNA == "" && r.UserID <= 0 {
		return errors.New(fmt.Sprintf("userDNA为空并且userID<=0,无法查询用户"))
	}
	// TODO 参数正确性校验

	return nil
}

type BCUpdateUserResp struct {
}
