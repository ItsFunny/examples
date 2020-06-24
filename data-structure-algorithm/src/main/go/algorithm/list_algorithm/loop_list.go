package list_algorithm

/*
	判断一个链表是否有环
 */
type ListNode struct {
	data interface{}
	next *ListNode
}
type LinkedList struct {
	root *ListNode
	size int
}

func (l *LinkedList)Push(value interface{}){
	newNode:=&ListNode{
		data: value,
	}
	if nil==l.root {
		l.root=newNode
	}else{
		tempNode:=l.root
		for nil!=tempNode.next{
			tempNode=tempNode.next
		}
		tempNode.next=newNode
	}
	l.size++
}

// true :loop else non-loop
// 遇到的坑: 内部的quickNode 要多走一步,而不是所初始化的时候多走一步,如果内部不多走一步是永远与slowNode相连的
func (l *LinkedList)ValidIfLoop()bool{
	if l.size==0 {
		return false
	}
	slowNode:=l.root
	quickNode:=slowNode.next
	for {
		if nil==quickNode{
			return false
		}else if quickNode==slowNode{
			return true
		}else {
			// 遇到一个坑,这里不能直接把quickNode的值给他
			// 不过遇到这个坑的缘故在于少了 内部的if那块代码,如果没有那块代码,quickNode与slowNode相连
			slowNode=slowNode.next
			quickNode=quickNode.next
			if nil!=quickNode {
				quickNode=quickNode.next
			}
		}
	}
}

// 通过快慢指针获取回环的节点
// 遇到的坑: 当break的时候代表发生了碰撞,因为我采用的方法,初始化的时候快节点会比慢节点快一步,所以当break获取
// 回环节点的时候也需要先快一步
func (l *LinkedList)GetLoopNode()*ListNode{
	// 这个方法的测试的目标是:必有回环的链表
	if l.size==0{
		return nil
	}
	slowNode:=l.root
	quickNode:=slowNode.next
	for{
		slowNode=slowNode.next
		quickNode=quickNode.next
		if nil!=quickNode {
			quickNode=quickNode.next
		}
		if nil==quickNode {
			return nil
		}else if quickNode==slowNode{
			// 说明有环碰撞了
			break
		}else{
			slowNode=slowNode.next
			quickNode=quickNode.next
			if nil!=quickNode{
				quickNode=quickNode.next
			}
		}
	}
	// 如果发生碰撞,则慢节点从表头出发,快节点从碰撞处出发,两者的运动速率是一致的,如果2者相等就是回环的节点
	// 这里有一个坑,因为最开始quickNode就比slowNode快一步(初始化的时候就快一步,所以这里也需要先提前快一步)
	slowNode=l.root
	quickNode=quickNode.next
	for quickNode!=slowNode{
		slowNode=slowNode.next
		quickNode=quickNode.next
	}
	return slowNode
}


// 通过map 额外申请内存来判断是否存在环
func (l LinkedList)ValidIfLoopByMap()bool{
	validMap:=make(map[*ListNode]struct{})
	tempNode:=l.root
	for nil!=tempNode{
		if _,ok:=validMap[tempNode] ;!ok{
			validMap[tempNode]= struct{}{}
		}else{
			return true
		}
	}
	return false
}