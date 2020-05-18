/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-05 15:11 
# @File : user.go
# @Description : 
# @Attention : 
*/
package parent

type UserGetCoinAmoutParent struct {
	UserId int
	Dna    string
}

type UserGetInfoParent struct {
	UserId int
	Dna    string
}
type UserGetInfoRespParent struct {
	Dna      string `json:"dna"`
	RealName string `json:"realName"`
	IDCard   string `json:"idCard"`
	// 上链时间
	UploadTime int64  `json:"uploadTime"`
	Phone      string `json:"phone"`

	WechatID          string `json:"wechatId"`
	MainWalletAddress string `json:"mainWalletAddress"`
	UserID            int    `json:"userId"`
	// 2020-01-02 add
	// 性别
	Gender int `json:"gender"`
	// 出生日期
	Birthday int64 `json:"birthday"`
	// 地址
	Address string `json:"address"`
	// 2020-01-06 add
	// 信息完善度
	// 姓名,性别,身份证,身份证,地址,手机 :6
	CompletePoint float64 `json:"completePoint"`

	TxID string `json:"txId"`
}

type UserUpdatePwdReqParent struct {
	UserId int    `json:"userId,string"`
	NewPwd string `json:"newPwd"`
}

type UserUpdatePwdRespParent struct {
	NewMainAddress string `json:"newMainAddress"`
}
