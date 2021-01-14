/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-03 09:07 
# @File : stack.go
# @Description : 
# @Attention : 
*/
package stack

import (
	"examples/data-structure-algorithm/src/main/go/v2/tree"
	"strconv"
	"strings"
)

// 给定一个经过编码的字符串，返回它解码后的字符串。 s = "3[a]2[bc]",
// 返回 "aaabcbc". s = "3[a2[c]]", 返回 "accaccacc". s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
func decodeString(s string) string {
	sb := strings.Builder{}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			// 需要弹出直到为[ 前的字符串
			values := make([]byte, 0)
			internalBuilder := strings.Builder{}
			for j := i - 1; j > 0 && s[j] != '['; j-- {
				values = append(values, stack[j])
				stack = stack[:stack[i]]
			}
			internalBuilder.Write(values)
			// 当前末尾元素为 [,弹出
			stack = stack[:len(stack)-1]
			// 当前末尾元素为数字
			count, _ := strconv.Atoi(string(stack[len(stack)-1]))
			// 重复
			for j := 0; j < count; j++ {
				values = append(values, values...)
				// internalBuilder.Write(values);
			}
			// 重新入栈
			stack = append(stack, values...)
		default:
			stack = append(stack, s[i])
		}
	}
	return sb.String()
}

// 通过栈实现DFS
func DFS(node *tree.TreeNode) []interface{} {
	if node == nil {
		return nil
	}
	stack := make([]*tree.TreeNode, 0)
	stack = append(stack, node)
	result := make([]interface{}, 0)

	for walkerNode := stack[0]; nil != walkerNode && len(stack) > 0; walkerNode = stack[len(stack)-1] {
		result = append(result, walkerNode.Data)
		if nil != walkerNode.LeftNode {
			stack = append(stack, walkerNode.LeftNode)
		}
		if nil != walkerNode.RightNode {
			stack = append(stack, walkerNode.RightNode)
		}
		stack = stack[:len(stack)-1]
	}

	return result
}

// 二叉树的中序遍历
func MiddleOrderTree(node *tree.TreeNode) []interface{} {
	stack := make([]*tree.TreeNode, 0)
	result := make([]interface{}, 0)

	for nil != node || len(stack) > 0 {
		for nil != node {
			stack = append(stack, node)
			node = node.LeftNode
		}
		valNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, valNode.Data)
		node = valNode.RightNode
	}

	return result
}

// 给定一个由  '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量
// 一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// dfs:计算周围是否也是岛屿,是的话,将其赋值为0,这样可以避免重复判断,因为连成一片的岛屿都是算为1个
			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
				count++
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i int, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == '1' {
		// 赋值为0的原因在于,可以避免重复探测
		grid[i][j] = '0'
		// +1的原因在于如果4周都是空的,而调用这个方法的前提在上面,表示必然是一个岛屿,所以+1 使得上面的if判断成立
		return dfs(grid, i, j-1) + dfs(grid, i, j+1) + dfs(grid, i-1, j) + dfs(grid, i+1, j) + 1
	}
	return 0
}
