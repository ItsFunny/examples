/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 14:06 
# @File : type.go
# @Description : 
# @Attention : 
*/
package base

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"myLibrary/go-library/go/authentication"
	error2 "myLibrary/go-library/common/error"
)

type TransBaseType int
type ChannelID string
type OrganizationID string
type ChainCodeID string
type MethodName string
// 代表区块链上的key
type Key string
// fromWalletAddress 从哪个钱包过来的
type From string
// toWalletAddress 去往哪个钱包的
type To string
// token 交易coin
type Token float64
type Version uint64
type ObjectType string
type KeyGenerater func(stub shim.ChaincodeStubInterface, objectType ObjectType, param ...interface{}) (string,   error2.IBaseError)
type TransBaseTypeV2 authentication.Authority
type TransBaseTypeV2Value authentication.AuthValue



func (this TransBaseTypeV2) Contains(transValue TransBaseTypeV2Value) bool {
	return authentication.Authority(this).CheckAuthentication(authentication.AuthValue(transValue))
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

)