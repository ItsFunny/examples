/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/26 9:16 上午
# @File : lt_208_实现前缀树.go
# @Description :
# @Attention :
*/
package v2

// 关键:
// 1. 字典树的定义: 是根据字典序所排列的数据结构
// 2. 数据结构定义: 必须要有children
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func TrieConstructor() Trie {
	ret := Trie{}
	return ret
}

func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		index := v - 'a'
		if temp.children[index] == nil {
			temp.children[index] = &Trie{}
		}
		temp = temp.children[index]
	}
	temp.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.searchByPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) searchByPrefix(prefix string) *Trie {
	node := this
	for _, v := range prefix {
		index := v - 'a'
		if node.children[index] == nil {
			return nil
		}
		node = node.children[index]
	}
	return node
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchByPrefix(prefix) != nil
}
