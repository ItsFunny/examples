/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-10 16:48 
# @File : follower.go
# @Description : 
# @Attention : 
*/
package raft

type ResponseForVoteReq struct {
}
type ResponseForVoteResp struct {
}

type ResponseForNetworkChangeReq struct {
}
type ResponseForNetworkChangeResp struct {
}
type Follower interface {
	GetStatus() byte
	// 响应选举
	ResponseForVote(req ResponseForVoteReq) (resp ResponseForVoteResp, err error)
	// 响应成员变更接口
	ResponseForNetworkChange(req ResponseForNetworkChangeReq) (resp ResponseForNetworkChangeResp, err error)
	// 响应添加entry 块
	ResponseAppendEntry(req ResponseAppendEntryReq) (resp ResponseAppendEntryResp, err error)
}

type FollowerImpl struct {
}

func (this *FollowerImpl) GetStatus() byte {
	return FOLLOWER
}

func (this *FollowerImpl) ResponseForVote(req ResponseForVoteReq) (resp ResponseForVoteResp, err error) {
	panic("implement me")
}

func (this *FollowerImpl) ResponseForNetworkChange(req ResponseForNetworkChangeReq) (resp ResponseForNetworkChangeResp, err error) {
	panic("implement me")
}

func (this *FollowerImpl) ResponseAppendEntry(req ResponseAppendEntryReq) (resp ResponseAppendEntryResp, err error) {
	panic("implement me")
}
