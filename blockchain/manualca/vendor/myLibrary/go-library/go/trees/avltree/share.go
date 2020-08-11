/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-05 14:29 
# @File : share.go
# @Description : 
# @Attention : 
*/
package avltree

import "github.com/emirpasic/gods/utils"

type ShareUserTree struct {
	*Tree
}

func (t *ShareUserTree) GetBy(key string) (interface{}, bool) {
	return t.Get(utils.StringComparator, key)
}
func (t *ShareUserTree) AddSharedUser(idCard string) {
	t.Put(utils.StringComparator, idCard, struct{}{})
}

func NewShareUserTree() *ShareUserTree {
	t := new(ShareUserTree)
	t.Tree = new(Tree)
	return t
}
