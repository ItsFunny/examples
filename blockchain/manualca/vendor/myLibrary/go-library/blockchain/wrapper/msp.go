/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-27 07:37 
# @File : msp.go
# @Description : 
# @Attention : 
*/
package wrapper

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"myLibrary/go-library/common/blockchain/base"
)

type MspOrganizationCaInfo struct {
	CaName string
}

func NewMspClientWrapper(c *msp.Client) *MspClientWrapper {
	w := new(MspClientWrapper)
	w.Client = c
	return w
}

type MspClientWrapper struct {
	Client *msp.Client
	CaInfo MspOrganizationCaInfo
}

// 组织msp wrapper
type OrganizationMspWrapper struct {
	Clients map[base.OrganizationID]*MspClientWrapper
}

func NewOrganizationMspWrapper() *OrganizationMspWrapper {
	w := new(OrganizationMspWrapper)
	w.Clients = make(map[base.OrganizationID]*MspClientWrapper)
	return w
}
