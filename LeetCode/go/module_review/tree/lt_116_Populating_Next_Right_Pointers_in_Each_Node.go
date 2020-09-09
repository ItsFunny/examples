/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-09 08:59 
# @File : lt_116_Populating_Next_Right_Pointers_in_Each_Node.go
# @Description : 
# @Attention : 
*/
package tree

/*
	填充每个节点的下一个右侧节点,如果不存在,则直接置空
	1. BFS
	2. 递归加上父节点的next
 */

func connect(root *Node) *Node {
	if nil == root {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)
		var prev *Node
		for i := 0; i < l; i++ {
			tempNode := queue[0]
			queue = queue[1:]
			if nil != prev {
				prev.Next = tempNode
			}
			if nil != tempNode.Left {
				queue = append(queue, tempNode.Left)
			}
			if nil != tempNode.Right {
				queue = append(queue, tempNode.Right)
			}
			prev = tempNode
		}
		prev = nil
	}

	return root
}

func connect2(root *Node) *Node {
	if nil == root {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Right != nil {
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}
	connect2(root.Left)
	connect2(root.Right)

	return root
}
