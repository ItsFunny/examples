/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-19 16:06 
# @File : copyright.go
# @Description : 
# @Attention : 
*/
package rpc

import (
	"examples/blockchain/config/common/parent"
)

type RPCItemUplaod2ChainReq struct {
	parent.ItemUpload2BlockChainReqParent
}

type RPCItemUplaod2ChainResp struct {
	parent.ItemUpload2BlockChainRespParent
}

type RPCItemGetCopyrightReq struct {
	parent.GetItemCopyrightReqParent
}

type RPCItemGetCopyrightResp struct {
	parent.GetItemCopyrightRespParent

	// 区块编号, 既区块高度
	BlockNumber uint64 `json:"blockNumber"`
	// 交易地址,既数据保存在哪个block中
	TransactionBlockAddress string `json:"transactionBlockAddress"`
}
