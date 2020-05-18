/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-17 13:56 
# @File : resp.go
# @Description : 
# @Attention : 
*/
package models

import "vlink.com/v2/vlink-common/base/fabric"

type BaseFacadedResp struct {
	base.TXRecordInfo
	Data interface{}

	Code int    `json:"code"`
	Msg  string `json:"msg"`

	TxRecords []base.TxRecordInfoNode `json:"records"`
}

/*
	正常的单链的逻辑为: 内部只会执行一条tx,并不会对其他的链码发生调用,此时这种情况是适合的
	但是,当如果是facaded类型的cc发起调用的时候,此时数据会出现多条txRecord的记录
	因此弥补措施为,当发起调用的时候,初始化下标为0的交易,并且后续的set都是针对于下标为0的元素
 */

// FIXME  这里遗留的问题,就是当内部存在多条tx的时候,如何设置tx
// FIXME 或者说是append的形式,而非set的形式

func (r *BaseFacadedResp) SetTxBaseType(baseType base.TransBaseTypeV2) {
	r.BaseType = baseType
	r.TxRecords[0].BaseType = baseType
}

func (r *BaseFacadedResp) SetChannelID(c base.ChannelID) {
	r.TxRecords[0].ChannelID = c
}

func (r *BaseFacadedResp) SetTransactionID(tid string) {
	r.TxRecords[0].TransactionID = tid
}

func (r *BaseFacadedResp) SetTxDescription(d string) {
	r.TxDescription = d
	r.TxRecords[0].TxDescription = d
}

func (r *BaseFacadedResp) GetCode() int {
	return r.Code
}

func (r *BaseFacadedResp) GetMsg() string {
	return r.Msg
}

func (r *BaseFacadedResp) GetTXRecordInfoList() []base.TxRecordInfoNode {
	return r.TxRecords
}

func (r *BaseFacadedResp) GetReturnData() interface{} {
	return r.Data
}

// 用于保存使用到的key,并且判断key是否是加密的
type FindBaseKeyNode struct {
	Key base.Key
	// 是否需要加密或者解密
	Crypt bool
}

//
type FindBaseKeyModel struct {
	// 使用到的key
	KeyMap map[base.ObjectType]FindBaseKeyNode
}

func (m *FindBaseKeyModel) Put(objectType base.ObjectType, key base.Key, crypt bool) {
	if m.KeyMap == nil {
		m.KeyMap = make(map[base.ObjectType]FindBaseKeyNode)
	}
	m.KeyMap[objectType] = FindBaseKeyNode{
		Key:   key,
		Crypt: crypt,
	}
}
func (m *FindBaseKeyModel) Append(m2 FindBaseKeyModel) {
	if m.KeyMap == nil {
		m.KeyMap = make(map[base.ObjectType]FindBaseKeyNode)
	}
	for k, v := range m2.KeyMap {
		m.KeyMap[k] = v
	}
}

func (m *FindBaseKeyModel) Get(ot base.ObjectType) FindBaseKeyNode {
	if m.KeyMap == nil {
		return FindBaseKeyNode{}
	} else {
		return m.KeyMap[ot]
	}
}
