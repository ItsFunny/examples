package hot100

// 关键: 链表已经排序
// 还要注意,可能头节点被删除,所以1. 要有dummy 2. 开始的节点不能是dummy#Next
// 因为要删除所有重复元素,而不是只保留一个,所以 必须用next 和next.next 去匹配
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	var rmValue int
	for temp := dummy; temp.Next != nil && temp.Next.Next != nil; {
		if temp.Next.Val == temp.Next.Next.Val {
			rmValue = temp.Next.Val
			// 然后删除所有与之相同的节点
			for nil!=temp.Next && temp.Next.Val == rmValue {
				temp.Next = temp.Next.Next
			}
		} else {
			temp = temp.Next
		}
	}

	return dummy.Next
}
