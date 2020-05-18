/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-19 13:42 
# @File : method.go
# @Description : 
# @Attention : 
*/
package rpc

const (
	// 基础RPC,没有对应方法实现
	RPCIBase = "rpc.base"
	// 含有业务实现的rpc
	RPCIUser="rpc.User"


	// 文件服务器
	RPCIFile="rpc.File"

	// 版权rpc
	RPCICopyright="rpc.Copyright"

	// 区块链管理rpc
	RPCIBlockChain="rpc.BlockChain"
)
