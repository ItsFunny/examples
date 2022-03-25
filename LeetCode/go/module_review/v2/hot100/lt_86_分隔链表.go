package hot100

// 关键:
// 2个链表,将大于x的节点都挪到另外一个链表中,最后再拼接
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	bigger := &ListNode{}
	dummy := &ListNode{Next: head}
	head = dummy
	// 然后开始遍历
	biggerTemp := bigger
	temp := head
	// 要从 temp.next 作为判断条件,原因在于,如果用temp作为判断条件,
	// 则当temp 当前值 >=x 的时候, 需要讲temp 从原先的head中删除,此时是做不到的
	// 所以只能使用temp.next
	for temp.Next != nil {
		if temp.Next.Val < x {
			temp = temp.Next
		} else {
			// 把更大的节点放到bigger中
			biggerTemp.Next = temp.Next
			biggerTemp = biggerTemp.Next
			// 然后跳到下一个节点
			temp.Next = temp.Next.Next
		}
	}

	// 最后将两个链表连接
	biggerTemp.Next = nil
	temp.Next = bigger.Next

	return dummy.Next
}
