/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-08-20 09:44 
# @File : base.go
# @Description : 
# @Attention : 
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTree(node *TreeNode) []int {
	// 根左右
	if node == nil {
		return nil
	}
	result := make([]int, 0)
	left := preOrderTree(node.Left)
	right := preOrderTree(node.Right)
	result = append(result, node.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// 前序非递归
func preOrderWithStack(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for nil != root || len(stack) > 0 {
		for nil != root {
			result = append(result, root.Val)
			stack = append(stack, root.Left)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// 中序遍历
func inOrderTree(root *TreeNode) []int {
	// 左根右
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	left := inOrderTree(root.Left)
	result = append(result, root.Val)
	right := inOrderTree(root.Right)
	result = append(result, left...)
	result = append(result, right...)

	return result
}

func inOrderWithStack(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for nil != root || len(stack) > 0 {
		for nil != root {
			stack = append(stack, root.Left)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}

	return result
}

// 后序遍历
func afterOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	left := afterOrder(root.Left)
	right := afterOrder(root.Right)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)

	return result
}

// 后序遍历不同在于 根节点必须在右节点之后弹出
func afterOrderWithStack(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	var lastNode *TreeNode
	for nil != root || len(stack) > 0 {
		for nil != root {
			stack = append(stack, root.Left)
			root = root.Left
		}
		node := stack[len(stack)-1]
		// 校验 右节点是否被弹出过了
		if node.Right == nil || node.Right == lastNode {
			result = append(result, node.Val)
			stack = stack[:len(stack)-1]
			lastNode = node
		} else {
			root = node.Right
		}
	}

	return result
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	result := make([]int, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		if nil != node.Right {
			stack = append(stack, node.Right)
		}
		if nil != node.Left {
			stack = append(stack, node.Left)
		}
		stack = stack [1:]
	}

	return result
}

func bfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([]int, 0)
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		result = append(result, node.Val)
		if nil != node.Left {
			queue = append(queue, node.Left)
		}
		if nil != node.Right {
			queue = append(queue, node.Right)
		}
		return result
	}
	return result
}

// 通过分治法遍历二叉树
func loopTreeByDivide(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	left := loopTreeByDivide(root.Left)
	right := loopTreeByDivide(root.Right)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)

	return result
}

// 归并排序
func mergeSort(nums []int) {

}
func divide(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	d := len(nums) >> 1
	left := divide(nums[:d])
	right := divide(nums[d:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	result := make([]int, 0)
	i := 0
	j := 0
	for ; i < len(left) && j < len(right); {
		i++
		j++
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[i])
		i++
	}

	return result
}

// 快速排序
func QSort(data []int) {
	qSort(data, 0, len(data)-1)
}
func qSort(data []int, start int, end int) {
	if start < end {
		index := paration(data, start, end)
		qSort(data, start, index)
		qSort(data, index+1, end)
	}
}
func paration(data []int, start, end int) int {
	standard := data[start]
	for start < end {
		for end > start && data[end] >= standard {
			end--
		}
		data[start] = data[end]

		for start < end && data[start] <= standard {
			start++
		}
		data[end] = data[start]
	}
	data[start] = standard
	return start
}

func maxDepth(root *TreeNode) int {
	return dep(root)
}
func dep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dep(root.Left)
	right := dep(root.Right)
	if left < right {
		return right + 1
	}
	return left + 1
}

// 判断是否是高度平衡的二叉树
// 1. 则子树也要平衡 ,判断深度即可

func isBalanced(root *TreeNode) bool {

	return balanced(root) > -1
}

func balanced(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := balanced(node.Left)
	right := balanced(node.Right)

	if left == -1 || right == -1 || mabs(left, right) <= -1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func mabs(a, b int) int {
	if a < b {
		return a - b
	}
	return b - a
}

// 树的最大路径和
func MaxSumOfTree(root *TreeNode) int {
	return maxSumOfTree(root)
}

func maxSumOfTree(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := maxSumOfTree(node.Left)
	right := maxSumOfTree(node.Right)
	leftRightMax := max(left, right)
	sumMax := max(leftRightMax+node.Val, node.Val)
	allMax := max(sumMax, left+right+node.Val)
	return allMax
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// lowest-common-ancestor-of-a-binary-tree
// 最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

func levelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		result := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if nil != node.Left {
				queue = append(queue, node.Left)
			}
			if nil != node.Right {
				queue = append(queue, node.Right)
			}
			result = append(result, node.Val)
		}
	}
	return result
}

// binary-tree-level-order-traversal-ii
// 打印输出的时候反转结果

func levelOrderBottom(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		single := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if nil != node.Left {
				queue = append(queue, node.Left)
			}
			if nil != node.Right {
				queue = append(queue, node.Right)
			}
			single = append(single, node.Val)
		}
		result = append(result, single)
	}
	reverse(result)
	return result
}

func reverse(ints [][]int) {
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	result := make([][]int, 0)
	toogle := false
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		single := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if nil != node.Left {
				queue = append(queue, node.Left)
			}
			if nil != node.Right {
				queue = append(queue, node.Right)
			}
			if toogle {
				rr(single)
			}
			toogle = !toogle
			result = append(result, single)
		}
	}
	return result
}
func rr(ints []int) {

}

// validate-binary-search-tree
// 中序遍历判断是否有序
func isValidBST(root *TreeNode) bool {
	data := inOrder(root)
	fmt.Println(data)
	return true
}
func inOrder(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}
	result := make([]int, 0)
	left := inOrder(root.Left)
	result = append(result, root.Val)
	right := inOrder(root.Right)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

// DFS的方式确定: 左子树< 根<右子树
func isVBST(root *TreeNode) bool {
	if nil == root {
		return true
	}
	if (nil != root.Left && root.Left.Val > root.Val) || (nil != root.Right && root.Right.Val < root.Val) {
		return false
	}
	left := isValidBST(root.Left)
	right := isValidBST(root.Right)
	return left && right
}

// insert-into-a-binary-search-tree
// 二叉搜索树中插入
// DFS查找插入位置
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if nil == root {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// / 链表
/*
	谨记:
		1. 快慢指针找中点
		2. 如果涉及到头结点,使用dummy node
 */

// 删除重复的元素
// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
func deleteDuplicates(head *ListNode) *ListNode {
	walkerNode := head
	for nil != walkerNode {
		for nil != walkerNode && walkerNode.Next.Val == walkerNode.Val {
			walkerNode.Next = walkerNode.Next.Next
		}
		walkerNode = walkerNode.Next
	}
	return head
}

// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中   没有重复出现的数字。
// 判断条件为 下一个的下一个不为空
func deleteDuplicates2(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	head = dummy
	for nil != head.Next && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal := head.Next.Val
			for nil != head.Next && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}

// reverse-linked-list
// 反转单链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for nil != cur {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}
	return prev
}

// reverse-linked-list-ii
// 反转从位置  m  到  n  的链表。请使用一趟扫描完成反转。
// 思路就是遍历到m 出,然后 [m-n] 处反转即可
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if nil == head {
		return nil
	}
	dummyNode := &ListNode{
		Next: head,
	}
	walkerNode := dummyNode.Next
	var oldPrev *ListNode
	i := 0
	for nil != walkerNode && i < m {
		oldPrev = walkerNode
		walkerNode = walkerNode.Next
		i++
	}
	// walkerNode为 m所处的node
	// oldPrev为之前的节点
	// 反转[m-n]节点
	var prev *ListNode
	curNode := walkerNode
	var oldNext *ListNode
	for nil != curNode && i < n {
		oldNext = curNode.Next
		curNode.Next = prev

		prev = curNode
		curNode = oldNext
		i++
	}

	oldPrev.Next = prev
	walkerNode.Next = oldNext

	return dummyNode.Next
}

// partition-list
// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于  x  的节点都在大于或等于  x  的节点之前
// 既,遍历链表, 小于的匹配之后删除该节点即可

// 链表排序
// sort-list
// 在  O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
// 解决方法: 归并排序 ,并且通过找中点的思路,快慢指针找中点
// 链表排序: 归并排序->找中点
func sortList(head *ListNode) *ListNode {
	return mergeSortListNode(head)
}
func mergeSortListNode(head *ListNode) *ListNode {
	if nil == head || head.Next == nil {
		return head
	}
	middle := findMiddle(head)
	tail := middle.Next
	middle.Next = nil
	left := mergeSortListNode(head)
	right := mergeSortListNode(tail)
	return merge2ListNode(left, right)
}
func merge2ListNode(left, right *ListNode) *ListNode {
	return &ListNode{}
}
func findMiddle(head *ListNode) *ListNode {
	fast := head.Next
	slow := head
	for nil != fast.Next && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// reorder-list
// 给定一个单链表  L：L→L→…→L__n→L 将其重新排列后变为： L→L__n→L→L__n→L→L__n→…
// reorder: => 找到中间节点=>后续反转即可
func reorderList(head *ListNode) {
	if nil == head {
		return
	}
	slow := head
	fast := head.Next
	for nil != fast.Next && fast.Next.Next != nil {
		fast = head.Next.Next
		slow = slow.Next
	}

	// 中间节点链表反转
	mid := slow
	afterMid := mid.Next
	var prev *ListNode
	cur := afterMid
	for nil != cur {
		next := cur.Next
		cur.Next = prev

		cur = next
		prev = cur
	}

	// 合并2个链表
	first := head
	second := prev
	toogle := false
	dummyNode := &ListNode{
		Next: first,
	}
	walkerNode := dummyNode
	for nil != first && nil != second {
		if toogle {
			walkerNode.Next = first
			first = first.Next
		} else {
			walkerNode.Next = second
			second = second.Next
		}
		toogle = !toogle
		walkerNode = walkerNode.Next
	}

	for nil != first {
		walkerNode.Next = first
		first = first.Next
	}
	for nil != second {
		walkerNode.Next = second
		second = second.Next
	}

	head = dummyNode.Next
}

// linked-list-cycle
// 判断是否有环 => 快慢节点
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head
	for nil != fast.Next {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

// linked-list-cycle-ii
// 给定一个链表，返回链表开始入环的第一个节点。  如果链表无环，则返回  null。
// 循环链表找中间节点=> 快慢指针 => 相遇块指针移到首节点同步
// 注意: 相遇的时候slow 要走前一步 ,当再次相遇则为回环节点
func detectCycle(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	slow := head
	fast := head.Next
	for nil != head.Next {
		if slow == fast {
			fast = head
			slow = slow.Next
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		} else {
			fast = fast.Next.Next
			slow = slow.Next
		}
	}
	return nil
}

// evaluate-reverse-polish-notation
// 波兰表达式计算 波兰表达式计算 > 输入: ["2", "1", "+", "3", "*"] > 输出: 9
func evalRPN(tokens []string) int {
	stack := make([]string, 0)
	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			first := stack[len(stack)-2]
			second := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":

			}
		default:
			stack = append(stack, v)
		}
	}
	v, _ := strconv.Atoi(stack[0])
	return v
}

// decode-string
// 给定一个经过编码的字符串，返回它解码后的字符串。
// s = "3[a]2[bc]", 返回 "aaabcbc". s = "3[a2[c]]",
// 返回 "accaccacc". s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	st := strings.Builder{}
	for _, v := range s {
		switch v {
		case ']':
			strs := make([]byte, 0)
			for b := stack[len(stack)-1]; b != '['; stack = stack[len(stack)-1:] {
				strs = append(strs, b)
			}
			reverseBytes(strs)
			stack = stack[len(stack)-1:]
			counts := make([]byte, 0)
			for b := stack[len(stack)-1]; b > 57 && b < 69; stack = stack[len(stack)-1:] {
				counts = append(counts, b)
			}
			reverseBytes(counts)
			st.Reset()
			st.Write(counts)
			couunt, _ := strconv.Atoi(st.String())
			for i := 0; i < couunt; i++ {
				stack = append(stack, strs...)
			}
		default:
			stack = append(stack, byte(v))
		}
	}
	st.Reset()
	st.Write(stack)

	return st.String()
}

func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i <= j; {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

// 树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	result := make([]int, 0)
	for len(stack) > 0 && nil != root {

		for nil != root {
			stack = append(stack, root.Left)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[len(stack)-1:]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

// number-of-islands
// 给定一个由  '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。
// 一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
// 核心就是dfs 遍历所有情况即可
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			count += dfsCountNumbs(grid, i, j)
		}
	}
	return count
}

func dfsCountNumbs(grid [][]byte, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid) || grid[i][j] == 2 {
		return 0
	}

	if grid[i][j] == '1' {
		grid[i][j] = 2
		return 1 + dfsCountNumbs(grid, i+1, j) + dfsCountNumbs(grid, i, j+1) + dfsCountNumbs(grid, i-1, j) + dfsCountNumbs(grid, i, j-1)
	}
	return 0
}

func BFS(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if nil != node.Left {
			queue = append(queue, node.Left)
		}
		if nil != node.Right {
			queue = append(queue, node.Right)
		}
		result = append(result, node.Val)
	}

	return result
}
