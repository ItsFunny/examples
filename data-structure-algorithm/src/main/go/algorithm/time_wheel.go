package algorithm

/*
	时间轮,内部的元素到期之后自动执行其回调方法
	V1 不提供运行时数组扩容
	V2 分层时间轮
	时间轮算法实现延迟任务的处理
	1. 数据结构:是一个环状数组,并且为了解决hash冲突,采用的是链地址法,名为Slot,并且存放的是其执行的任务(V3分发到线程池)
*/
// 因为时间轮是其内部到期之后自动执行回调
// 因此抽为接口
type SlotNodeInterfacer interface {
	CallBack() (interface{}, error)
}

type slotNode struct {
	data interface{}
	next *slotNode
}
type Slot struct {
	// 内部是链表
	root *slotNode
}

type TimeWheel struct {
	// 内部是环状数组
	// 如何定义一个环状数组:通过一个游标代替左右
	maxIndex int
	slots    []*Slot
}

func NewTimeWheel() *TimeWheel {
	timeWheel := TimeWheel{
		slots: make([]*Slot, 0),
	}
	return &timeWheel
}
func (w *TimeWheel) InsertJob() {

}
