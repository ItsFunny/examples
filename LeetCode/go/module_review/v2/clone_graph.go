/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/2 9:11 上午
# @File : clone_graph.go
# @Description :
# @Attention :
*/
package v2

func cloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}
	m := make(map[*GraphNode]*GraphNode)
	return clone(node, m)
}
func clone(node *GraphNode, m map[*GraphNode]*GraphNode) *GraphNode {
	if node == nil {
		return node
	}
	if v, exist := m[node]; exist {
		return v
	}

	newNode := &GraphNode{
		Val:       node.Val,
		Neighbors: make([]*GraphNode, len(node.Neighbors)),
	}
	m[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], m)
	}
	return newNode
}
