/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-10 16:41 
# @File : log.go
# @Description : 
# @Attention : 
*/
package raft

import "sync"


// 存储的是日志信息
type LogInfo struct {

}

// 管理日志
type LogManager struct {
	Lock *sync.RWMutex
	Logs []LogInfo
}
