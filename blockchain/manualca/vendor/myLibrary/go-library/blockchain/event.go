/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 13:00 
# @File : event.go
# @Description : 
# @Attention : 
*/
package config

import (
	"encoding/hex"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"go.uber.org/atomic"
	"myLibrary/go-library/blockchain/handler"
	"myLibrary/go-library/common/blockchain/base"
	"myLibrary/go-library/go/base/service"
)

var (
	executor *TaskExecutor
)

func init() {
	executor = newTaskExecutor()
}

func newTaskExecutor() *TaskExecutor {
	v := new(TaskExecutor)
	v.IBaseService = service.NewBaseServiceImplWithLog4goLogger()
	v.BlockHandlerMap = make(map[base.ChannelID]chan *EventBlockInfo)
	return v
}

type TaskExecutor struct {
	service.IBaseService
	BlockEventListeners []*BlockEventExecutor

	// FIXME 用于监听普通事件
	BlockHandlerMap map[base.ChannelID]chan *EventBlockInfo
}

type EventBlockWrapper struct {
	BlockEvents chan *EventBlockInfo
}

type EventBlockInfo struct {
	Block *cb.Block
}

type TemporaryBlockEventExecutor struct {
	service.IBaseService
	ChannelID          base.ChannelID
	ChainCodes         []string
	InterestChainCodes []string
	Events             <-chan *fab.BlockEvent
	Registration       fab.Registration
	BlockHandler       handler.IBlockHandler
}
type BlockEventExecutor struct {
	service.IBaseService
	ChannelID base.ChannelID

	ChainCodes         []string
	InterestChainCodes []string
	// 期望的block 编号,用于以防数据丢失
	ExpectedBlockIndex *atomic.Int64
	events             <-chan *fab.BlockEvent

	stop chan struct{}

	BlockHandler handler.IBlockHandler
}

type ClientTemporaryBlockEventInfo struct {
	Events             *fab.BlockEvent
	Registration       fab.Registration
	ChannelID          base.ChannelID
	ChainCodes         []string
	InterestChainCodes []string
}

// 注册临时的区块事件
// func RegisterTemporaryBlockEvent(cid base.ChannelID,)
func RegisterBlockEvent(cid base.ChannelID, setup SetupBlockEventExecutor, interestChainCodes []string, events <-chan *fab.BlockEvent, stop chan struct{}) {
	b := new(BlockEventExecutor)
	b.IBaseService = service.NewBaseServiceImplWithLog4goLogger()
	b.ChannelID = cid
	b.stop = stop
	b.InterestChainCodes = interestChainCodes
	b.ExpectedBlockIndex = atomic.NewInt64(1)
	b.events = events
	if setup != nil {
		setup(nil, b)
	}
	if b.BlockHandler == nil {
		b.BlockHandler = handler.NewLogBlockHandler()
	}

	if executor.BlockEventListeners == nil {
		executor.BlockEventListeners = make([]*BlockEventExecutor, 0)
	}
	executor.BlockEventListeners = append(executor.BlockEventListeners, b)
}

type SetupBlockEventExecutor func(interface{}, *BlockEventExecutor) error

func RunTasks() {
	for _, listner := range executor.BlockEventListeners {
		go listner.handleEventWithDetailBlock()
	}
	// go func() {
	// 	for {
	// 		select {
	// 		case event := <-executor.TemporaryEvents:
	// 			// FIXME 暂时先直接处理
	// 			go func() {
	// 				for {
	// 					select {
	// 					case b := <-event.Events:
	// 						defer event.Registratio
	// 					}
	// 				}
	// 			}()
	// 		}
	// 	}
	//
	// }()
}

func (this *BlockEventExecutor) handleEventWithDetailBlock() {
	for {
		select {
		case <-this.stop:
			break
		case e := <-this.events:
			if this.ExpectedBlockIndex.Load() > int64(e.Block.Header.Number) {
				this.Debug("该块编号为[%d],currentHash为[%s]为重复块,当前expectedBlockNumber为:[%v],因此直接跳过", int64(e.Block.Header.Number), hex.EncodeToString(e.Block.Header.DataHash), this.ExpectedBlockIndex)
				continue
			}
			w := handler.BlockWrapper{
				BlockEvent: e,
				ChainCodes: this.ChainCodes,
				ChannelId:  this.ChannelID,
			}
			if e := this.BlockHandler.Handle(w); nil != e {
				this.Error("handler block 失败:" + e.Error())
			} else {
				this.ExpectedBlockIndex.Inc()
			}
		}
	}
}
