/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-24 12:39 
# @File : blockchain.go
# @Description : 
# @Attention : 
*/
package parent

type BlockChainSysInfoReqParent struct {
	// 特定的channel,若为空,则查询所有的channel中的信息
	ChannelID string
}

type BlockChainSysInfoRespParent struct {
	Items []ChannelNodeParent `json:"items"`
}

type ChannelNodeParent struct {
	ChannelID        string `json:"channelId"`
	BlockHeight      uint64 `json:"blockHeight"`
	BlockCreatedTime int    `json:"blockCreatedTime"`
	// NormalPeerNodes  []PeerNodeParent `json:"normalPeerNodes"`
	AnchorPeerNodes []PeerNodeParent `json:"peers"`
	// private String channelId;
	// // 该channel的区块高度
	// private long blockHeight;
	// // 该channel的区块产生时间
	// private int blockCreatedTime;
	// // 该channel的交易总额
	// private float transAmount;
	// // 节点信息
	// private List<RPCPeerNode> nodes;
}

type PeerNodeParent struct {
	// 组织名称
	Org    string `json:"org"`
	Domain string `json:"domain"`
	Port   int32  `json:"port"`
}

// type BlockChainSysInfoRespNodeParent struct {
// 	BlockHeight uint64 `json:"blockHeight"`
// 	// 运行时间
// 	RunTime int `json:"run_time"`
// 	// 交易用户数量
// 	// TransUserCounts int `json:"transUserCounts"`
// 	// 初始运行时间
// 	Timestamp int64 `json:"timestamp"`
// 	// 总共上链作品数
// 	// TotalUploadedItems int `json:"total_uploaded_items"`
// 	// 交易总额
// 	TotalTransAmount float64 `json:"-"`
// 	// 2019-09-08 change    写死
// 	// 节点数量
// 	PeerCounts int `json:"peerCounts"`
// 	// 区块产生时间 单位 秒
// 	BlockCreatedTime int `json:"blockCreatedTime"`
// }

type BlockPageReqParent struct {
	ChannelID string `json:"channelId"`
	PageSize  int    `json:"pageSize"`
	PageNum   int    `json:"pageNum"`
}

type BlockPageRespParent struct {
	Items      []LatestBlockRespNode `json:"items"`
	TotalCount int64                 `json:"totalCount"`
	MaxPage    int                   `json:"maxPage"`
}

type LatestBlockRespNode struct {
	CurrentHash string `json:"currentHash"`
	PrevHash    string `json:"prevHash"`
	// 这个块中有多少个交易数
	TotalTransactionCounts int `json:"totalTransactionCounts"`
	// 区块编号number
	BlockNumber uint64 `json:"blockNumber"`
	CreatedTime int64 `json:"createdTime"`
	// 交易总金额
	TotalTransactionAmount float64 `json:"totalTransactionAmount"`
}
