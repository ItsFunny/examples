package hot100

// 关键:
// 1. 构建成环
// 2. 新的头节点在 n-(k%n) ,尾节点则在 n-(k%n)-1 处
// 注意点就是找新的节点时候的边界条件而已
func rotateRight(head *ListNode, k int) *ListNode {
	if nil==head{
		return nil
	}
	// 第一步,先构建成环
	loopNode := head
	l := 1
	for nil != loopNode.Next {
		loopNode = loopNode.Next
		l++
	}
	loopNode.Next = head

	// 第二步, 开始找尾节点,找到尾节点,就代表着已经找到了新的头结点
	// 注意点: 找新的head的时候,从loop开始
	walkNode := loopNode
	newTailCount := l - (k % l)
	for ;newTailCount>0;newTailCount--{
		walkNode=walkNode.Next
	}

	tail:=walkNode
	newHead:=tail.Next
	tail.Next=nil
	return newHead
}
