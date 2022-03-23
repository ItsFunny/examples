package hot100

// 关键: 根据题意,重复的元素保留1个
// 头结点可能会被删除
// 与lt82 不同的地方在于,82是删除所有重复的元素,而 83 是保留一个, 所以不可以直接从dummy开始
func deleteDuplicates83(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for temp := dummy.Next; nil !=temp && temp.Next != nil; temp = temp.Next {
		for temp != nil && temp.Next != nil && temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		}
	}
	return dummy.Next
}