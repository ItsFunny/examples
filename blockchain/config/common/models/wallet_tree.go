/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-18 13:40 
# @File : wallet_tree.go
# @Description : 
# @Attention : 
*/
package models

import (
	"github.com/emirpasic/gods/utils"
	"myLibrary/go-library/go/trees/avltree"
)

type WalletTree struct {
	avltree.Tree
}

type WalletTreeNode struct {
	Type int
	Path string
	WalletDetail *VlinkChildWallet
}

func NewWalletTree() *WalletTree {
	t := new(WalletTree)
	t.Tree = *new(avltree.Tree)
	return t
}
func (t *WalletTree) Insert(key interface{},node *WalletTreeNode) error {
	t.Put(utils.IntComparator, key, node)
	return nil
}

func (t *WalletTree) GetByType(tt int) (*WalletTreeNode, bool) {
	value, found := t.Get(utils.IntComparator, tt)
	if found {
		return value.(*WalletTreeNode), found
	}
	return nil, found
}
