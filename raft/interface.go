/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-10 16:31 
# @File : interface.go
# @Description : 
# @Attention : 
*/
package raft

// 内存文件
type LogArray struct {
}

type SnapShotReq struct {
}
type SnapShotResp struct {
}

// 保存本地文件
type CommonInterface interface {
	GetStatus() byte
	SnapShort(req SnapShotReq) (resp SnapShotReq, err error)
}
