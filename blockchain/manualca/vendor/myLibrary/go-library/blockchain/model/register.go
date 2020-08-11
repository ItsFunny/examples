/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-05 13:02 
# @File : register.go
# @Description : 
# @Attention : 
*/
package model

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"myLibrary/go-library/common/blockchain/base"
)

type UserRegisterReq struct {
	Oid base.OrganizationID
	Name string
	Secret string
	Type string

	Attributes []msp.Attribute
}

type UserRegistrationResp struct {
	EnrollSecret string
}
