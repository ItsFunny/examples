/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-10 16:48 
# @File : candidate.go
# @Description : 
# @Attention : 
*/
package raft

import "sync"

type AskForVoteNodeReq struct {
	Address string
	Port    int
}
type AskForVoteTermInfo struct {
	// 当前term 周期
	Term int64
}
type AskForVoteLogInfo struct {
	LastLogIndex int32
	LastLogTerm  int64
}
type AskForVoteReq struct {
	// 通信的地址,都有一个默认的地址
	Nodes []AskForVoteNodeReq
	AskForVoteTermInfo AskForVoteTermInfo
	AskForVoteLogInfo AskForVoteLogInfo
	// 停止选举的channel
	StopElectionChan <-chan struct{}
}

func (this AskForVoteReq) Valid() error {
	if len(this.Nodes) == 0 {
		return ErrorF("请求投票节点不可为空")
	}

	return nil
}

type AskForVoteResp struct {
	Status byte
}

type ResponseAppendEntryReq struct {
}

type ResponseAppendEntryResp struct {
}

// 候选人
type Candidate interface {
	GetStatus() byte
	// 请求投票
	AskForVote(req AskForVoteReq) (resp AskForVoteResp, err error)
}

type CandidateImpl struct {
}

func (this *CandidateImpl) GetStatus() byte {
	return CANDIDATE
}

// 投票有如下规则:
/*
	1. 需要携带当前的term周期,和上一个日志的term周期和index 下标
	注意点:
		1. 如果成为candidate,则选举自己成为leader
		2. 并行的请求其他节点
		3. 停止选举的情况:
			3.1 有节点成为了leader
			3.2 该节点成为了leader
			3.3 该任期内无leader出现
	It then votes for itself a
 */
func (this *CandidateImpl) AskForVote(req AskForVoteReq) (resp AskForVoteResp, err error) {
	wg := sync.WaitGroup{}
	wg.Add(len(req.Nodes))
	for i := 0; i < len(req.Nodes); i++ {
		go func() {

		}()
	}
	for {
		select {
		case <-req.StopElectionChan:
			log.Println("收到退出选举信号,退化为follower状态")
			resp.Status = FOLLOWER
		}
	}
	return
}
