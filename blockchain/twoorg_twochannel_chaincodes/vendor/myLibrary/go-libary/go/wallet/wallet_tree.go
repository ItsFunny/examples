/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-16 15:12 
# @File : wallet_tree.go
# @Description : 
# @Attention : 
*/
package wallet

import (
	"github.com/emirpasic/gods/utils"
	"myLibrary/go-libary/go/trees/avltree"
)

var (
// PATH_COMPARATOR = func(a, b interface{}) int {
// 	a1, a2 := a.(string), b.(string)
// 	// m/44/60/0/x
// 	s1 := strings.Split(a1, "/")
// 	s2 := strings.Split(a2, "/")
//
// 	s1 = s1[0:]
// 	s2 = s2[0:]
// 	l := len(s1)
// 	if l > len(s2) {
// 		l = len(s2)
// 	}
//
// 	for i := 0; i < l; i++ {
//
// 	}
// }
)

type WalletTree struct {
	*avltree.Tree
}

func (t *WalletTree) Insert(key interface{}, node *ChildWallet) error {
	t.Put(utils.IntComparator, key, node)
	return nil
}

func (t *WalletTree) GetByType(tt int) (*ChildWallet, bool) {
	value, found := t.Get(utils.IntComparator, tt)
	if found {
		return value.(*ChildWallet), found
	}
	return nil, found
}
