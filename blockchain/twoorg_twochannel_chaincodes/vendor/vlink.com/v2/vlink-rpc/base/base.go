/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-19 09:38 
# @File : base.go
# @Description : 
# @Attention : 
*/
package base

import "vlink.com/v2/vlink-common/base/fabric"

type ChannelID string
type OrganizationID string
type ChainCodeID string

type ChainBaseReq struct {
	MethodName  base.MethodName
	ChannelID   ChannelID
	ChainCodeID ChainCodeID
}
