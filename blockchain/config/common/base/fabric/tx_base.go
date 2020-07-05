/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-25 13:49 
# @File : tx_base.go
# @Description : 
# @Attention : 注意 这里添加常量只能尾部添加,不可在中部添加
*/
package base

import (
	"encoding/hex"
	"myLibrary/go-library/go/authentication"
)

type TransBaseTypeV2 authentication.Authority

type TransBaseTypeV2Value authentication.AuthValue

func (this TransBaseTypeV2) Contains(transValue TransBaseTypeV2Value) bool {
	return authentication.Authority(this).CheckAuthentication(authentication.AuthValue(transValue))
}
func (this TransBaseTypeV2) String() string {
	return hex.EncodeToString(this.BigEndianConvtBytes())
}

func (this TransBaseTypeV2) BigEndianConvtBytes() []byte {
	return authentication.Authority(this).BigEndianConvt2Bytes()
}
func NewTransBaseTypeV2WithValue(value TransBaseTypeV2Value) TransBaseTypeV2 {
	authority := authentication.NewAuthority()
	authority.AddAuthentication(authentication.AuthValue(value))
	return TransBaseTypeV2(authority)
}

func NewTransBaseTypeV2() TransBaseTypeV2 {
	return TransBaseTypeV2(authentication.NewAuthority())
}
func AddBaseType(v2 TransBaseTypeV2, baseType TransBaseTypeV2Value) TransBaseTypeV2 {
	return TransBaseTypeV2(authentication.Authority(v2).AddAuthentication(authentication.AuthValue(baseType)))
	// authority := authentication.NewAuthority()
	// authority.AddAuthentication(authentication.AuthValue(baseType))
	// return TransBaseTypeV2(authority)
}

func CreateNeedRecordBaseType(baseType TransBaseTypeV2Value) TransBaseTypeV2 {
	authority := authentication.NewAuthority()
	authority = authority.AddAuthentication(authentication.AuthValue(TX_BASE_NEED_RECORD))
	authority = authority.AddAuthentication(authentication.AuthValue(baseType))
	return TransBaseTypeV2(authority)
}

func CreateUnNeedRecordBaseType(baseType TransBaseTypeV2Value) TransBaseTypeV2 {
	authority := authentication.NewAuthority()
	authority = authority.AddAuthentication(authentication.AuthValue(TX_BASE_UNNEED_RECORD))
	authority = authority.AddAuthentication(authentication.AuthValue(baseType))
	return TransBaseTypeV2(authority)
}

const (
	// 代表需要记录交易详情
	// 1
	TX_BASE_NEED_RECORD = TransBaseTypeV2Value(iota + 1)
	// 2
	TX_BASE_UNNEED_RECORD

	/*
		COINCC
	 */
	// 更新用户积分
	// 3
	COINCC_TX_BASE_USER_COIN_UPDATE
	// 获取用户积分,不需要记录
	// 4
	COINCC_TX_BASE_USER_COIN_GET
	/*
		COPYRIGHTCC
	 */
	// 5
	COPYRIGHTCC_TX_BASE_UPLOAD_TO_CHAIN
	/*
		USERCC
	 */
	// 6
	USERCC_TX_BASE_USER_UPLOAD
	USERCC_TX_BASE_UPDATE_USERINFO
	USERCC_TX_BASE_GET_USERINFO
	USERCC_TX_BASE_GET_USERCOIN
	USERCC_TX_BASE_GET_USERWALLET
	/*
		USERCOINCC
	 */
	USERCOIN_TX_BASE_UPLOAD_USER_WITH_INIT_COIN
	USERCOIN_TX_BASE_UPDATE_USER_INFO


	// 2020-03-20 10:05 add
	COPYRIGHTCC_TX_BASE_GET_COPYRIGHTINFO


	USERCC_TX_BASE_UPDATE_USERPWD
)
