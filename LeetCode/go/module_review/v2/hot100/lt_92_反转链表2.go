package hot100

// 链表区间反转
// 关键:
// 1. 头节点是可能被反转的,所以需要dummy
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy:=&ListNode{Next: head}
	var headBeforeReverse *ListNode
	head=dummy
	i:=0
	for ;i<left;i++{
		headBeforeReverse=head
		head=head.Next
	}
	// 然后开始区间反转
	// 记录下leftNode的值,因为要与rightNode.next连接
	leftNode:=head

	var prev *ListNode
	for j:=i;j<=right;j++{
		// 链表反转
		tmp:=head.Next

		head.Next=prev
		prev=head

		head=tmp
	}
	// 开始重新连接,leftNode之前的一个节点的next 需要为right节点(既区间内的最后一个节点)
	headBeforeReverse.Next=prev
	// 此时的head是rightNode的下一个节点
	leftNode.Next=head

	return  dummy.Next
}