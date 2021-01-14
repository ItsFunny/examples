/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-10 16:26 
# @File : Server.go
# @Description : 
# @Attention : 
*/
package raft

import (
	"go.uber.org/atomic"
	"sync"
)

// 消息
const (
	FOLLOWER  = 1 << 1
	CANDIDATE = 1 << 2
	LEADER    = 1 << 3
)

// 存储的是监听器. 如 leader是否存活
type ListenerWrapper struct {
}

// 存储节点处于什么状态
type StatusWrapper struct {
	Status byte
}

// 存储的是配置信息
type ConfigWrapper struct {
	Lock *sync.RWMutex
}

// 存储的是节点信息
type NetworkWrapper struct {
	Lock *sync.RWMutex
}

type LogWrapper struct {
	Lock *sync.RWMutex
}

// 存储的是接口相关的wrapper
type FacadedWrapper struct {
	Lock *sync.RWMutex
}

// 存储的是term周期
type TermWrapper struct {
	// 当前term的周期
	TermID atomic.Int64
}

type Server struct {
	Lock            *sync.RWMutex
	LogWrapper      *LogWrapper
	NetworkWrapper  *NetworkWrapper
	ConfigWrapper   *ConfigWrapper
	StatusWrapper   *StatusWrapper
	ListenerWrapper *ListenerWrapper
}


func (this *Server) Start() error {

}
