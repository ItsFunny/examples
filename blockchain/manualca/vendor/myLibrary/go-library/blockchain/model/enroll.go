/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-27 07:53 
# @File : enroll.go
# @Description : 
# @Attention : 
*/
package model

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"myLibrary/go-library/common/blockchain/base"
)

type UserEnrollReq struct {
	Oid           base.OrganizationID
	UserUniqueKey string
	UserPassword  string
	Profile       string
	Type          string
}

func (req UserEnrollReq) BuildEnrollOptions() []msp.EnrollmentOption {
	result := make([]msp.EnrollmentOption, 0)
	if req.UserPassword != "" {
		result = append(result, msp.WithSecret(req.UserPassword))
	}
	if req.Profile != "" {
		result = append(result, msp.WithProfile(req.Profile))
	}
	if req.Type != "" {
		result = append(result, msp.WithType(req.Type))
	}

	return result
}
