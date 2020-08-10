/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-10 16:48 
# @File : leader.go
# @Description : 
# @Attention : 
*/
package raft

type NotifyAsLeaderReq struct {
}
type NotifyAsLeaderResp struct {
}
type HandlerClientReq struct {
}
type HandlerClientResp struct {
}

type ReplicateLogReq struct {
}
type ReplicateLogResp struct {
}

// 主导者
type Leader interface {
	CommonInterface
	// 通知,代表leader选举成功
	NotifyAsLeader(req NotifyAsLeaderReq) (resp NotifyAsLeaderResp, err error)
	// 响应客户端请求
	HandlerClient(req HandlerClientReq) (resp HandlerClientResp, err error)
	// 日志同步
	ReplicateLog(req ReplicateLogReq) (resp ReplicateLogResp, err error)
}

