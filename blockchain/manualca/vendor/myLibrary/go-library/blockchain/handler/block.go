/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 13:14 
# @File : block.go
# @Description : 
# @Attention : 
*/
package handler

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"myLibrary/go-library/common/blockchain/base"

	// "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/core/ledger/util"
	"myLibrary/go-library/go/base/service"
)

type BaseLogicServiceImpl struct {
	*service.BaseServiceImpl
}

type BlockWrapper struct {
	*fab.BlockEvent
	ChainCodes []string
	ChannelId base.ChannelID
}

// FIXME 提供一个基础的handler service,而不是logic 的service
type IBlockHandler interface {
	Handle(BlockWrapper) error
}

func NewBaseLogicServiceImpl() *BaseLogicServiceImpl {
	l := new(BaseLogicServiceImpl)
	l.BaseServiceImpl = service.NewBaseServiceImplWithLog4goLogger()
	return l
}

type LogBlockHandler struct {
	*BaseLogicServiceImpl
}

func NewLogBlockHandler() *LogBlockHandler {
	l := new(LogBlockHandler)
	l.BaseLogicServiceImpl = NewBaseLogicServiceImpl()
	return l
}

func (this *LogBlockHandler) Handle(wrapper BlockWrapper) error {
	// TxValidationFlags
	this.Debug("接收到块:%v", wrapper)
	return nil
	// e:=wrapper.BlockEvent
	// block := e.Block
	// sourceUrl := e.SourceURL
	// blockNumber := block.Header.Number
	// currentHash := block.Header.DataHash
	// prevHash := block.Header.PreviousHash
	//
	// l := len(block.Data.Data)
	// this.Debug("检测到生成了新的block的交易总数为:%d", l)
	// txFilter := util.TxValidationFlags(block.Metadata.Metadata[sdkCommon.BlockMetadataIndex_TRANSACTIONS_FILTER])
	// if len(txFilter) == 0 {
	// 	txFilter = util.NewTxValidationFlags(l)
	// 	block.Metadata.Metadata[common.BlockMetadataIndex_TRANSACTIONS_FILTER] = txFilter
	// }
	// totalTransaction := 0
	// detailInfo, err := utils.GetBlockDetailV2(wrapper.ChainCodes, block.Data.Data[0])
	// createdTime := detailInfo.CreateTime
	// amount := detailInfo.Amount
	// signature := ""
	// if detailInfo.Signature != nil && len(detailInfo.Signature) > 0 {
	// 	signature = hex.EncodeToString(detailInfo.Signature)
	// }
	// if nil != err {
	// 	this.Error("获取区块详情失败:%s", err.Error())
	// } else {
	// 	for i := 1; i < l; i++ {
	// 		if !txFilter.IsValid(i) {
	// 			logs.Error("检测到无效交易")
	// 			continue
	// 		}
	// 		isLogic, _, am, err := utils.GetBlockDetail(wrapper.ChainCodes, block.Data.Data[i])
	// 		if nil != err {
	// 			this.Error("获取区块详情失败:%s", err.Error())
	// 		} else if isLogic {
	// 			totalTransaction++
	// 			amount += am
	// 		}
	// 	}
	// 	id, _ := idGenerator.NextId()
	// 	b := models.DBVlinkBlock{
	// 		ID:                          int(id),
	// 		CreatedDate:                 createdTime,
	// 		BlockNumber:                 blockNumber,
	// 		BlockCurrentHash:            hex.EncodeToString(currentHash),
	// 		BlockPrevHash:               hex.EncodeToString(prevHash),
	// 		BlockTotalTransaction:       totalTransaction,
	// 		BlockTotalTransactionAmount: amount,
	// 		BlockSourceUrl:              sourceUrl,
	// 		ChannelID:                   string(this.ChannelID),
	// 		Signature:                   signature,
	// 	}
	// }
}
