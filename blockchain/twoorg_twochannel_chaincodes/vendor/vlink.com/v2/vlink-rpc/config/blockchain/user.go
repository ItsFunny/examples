/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-19 10:08 
# @File : user.go
# @Description : 
# @Attention : 
*/
package blockchain

import (
	"encoding/json"
	"myLibrary/go-libary/go/converters"
	"vlink.com/v2/vlink-common/base/fabric/constants/usercc/method"
	error2 "vlink.com/v2/vlink-common/error"
	"vlink.com/v2/vlink-common/models"
	"vlink.com/v2/vlink-rpc/base"
	"vlink.com/v2/vlink-rpc/constants"
	"vlink.com/v2/vlink-rpc/models/buss"
)

func (this *VlinkBlockChainConfiguration) UploadUser2Chain(req buss.BSUserUploadReq) (buss.BSUserUploadResp, error2.IVlinkError) {
	var (
		result buss.BSUserUploadResp
		e      error2.IVlinkError
	)

	var bcReq models.BCUploadUserReq
	bcReq.UploadTime = req.UploadTime
	bcReq.WeChatID = req.WeChatID
	bcReq.Phone = req.Phone
	bcReq.RealName = req.RealName
	bcReq.IDCard = req.IDCard
	bcReq.UserPwd = req.UserPwd
	bcReq.UserID = req.UserID
	bcReq.ServerStorageKey = req.ServerStorageKey
	bcReq.Address = req.Address
	bcReq.Gender = req.Gender
	bcReq.Birthday = req.Birthday

	response, vlinkError := this.defaultExecute(base.ChainBaseReq{
		MethodName:  method.METHOD_UPLOAD_USER_TO_CHAIN,
		ChannelID:   constants.CHANNEL_ID_USER,
		ChainCodeID: constants.CHAINCODE_ID_USERCC,
	}, bcReq)
	if nil != vlinkError {
		return result, vlinkError
	}

	resp, _ := HandleResponse(response)
	if nil != resp.DataBytes && len(resp.DataBytes) > 0 {
		if e := json.Unmarshal(response.Payload, &result); nil != e {
			return result, error2.NewJSONSerializeError(e, "反序列化失败")
		}
	}

	result.LogicCode = int(converter.BigEndianBytes2Int64(resp.CodeBytes))
	result.LogicMsg = string(resp.MsgBytes)

	return result, e
}

func (this *VlinkBlockChainConfiguration) UpdateUserInfo(req buss.BSUserUpdateReq) (result buss.BSUserUpdateResp, vlinkError error2.IVlinkError) {

	var bcReq models.BCUpdateUserReq
	bcReq.Phone = req.Phone
	bcReq.Address = req.Address
	bcReq.IDCard = req.IDCard
	bcReq.RealName = req.RealName
	bcReq.Birthday = req.Birthday
	bcReq.UserDNA = req.UserDNA
	bcReq.UserID = req.UserID
	bcReq.Gender = req.Gender

	response, e := this.defaultExecute(base.ChainBaseReq{
		MethodName:  method.METHOD_UPDATE_USERINFO,
		ChannelID:   constants.CHANNEL_ID_USER,
		ChainCodeID: constants.CHAINCODE_ID_USERCC,
	}, bcReq)
	if nil!=e{
		return result,e
	}

	resp, e := HandleResponse(response)
	result.LogicCode = int(converter.BigEndianBytes2Int64(resp.CodeBytes))
	result.LogicMsg = string(resp.MsgBytes)


	return
}
