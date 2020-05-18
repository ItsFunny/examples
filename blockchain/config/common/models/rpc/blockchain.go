/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-24 13:23 
# @File : blockchain.go
# @Description : 
# @Attention : 
*/
package rpc

import (
	"vlink.com/v2/vlink-common/parent"
)

type RPCBlockChainSysInfoReq struct {
	parent.BlockChainSysInfoReqParent
}

type RPCBlockChainSysInfoResp struct {
	parent.BlockChainSysInfoRespParent
}

type RPCBlockChainBlocksPageReq struct {
	parent.BlockPageReqParent
}
type RPCBlockChainBlocksPageResp struct {
	parent.BlockPageRespParent
}
