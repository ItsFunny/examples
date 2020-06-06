package list

import (
	"github.com/influxdata/influxdb/pkg/testing/assert"
	"testing"
)
/*
	测试用例List:
	通过下标删除节点:
		删除0号元素->PASS
		删除最后一个下标元素->PASS
		删除中间任意一个元素->PASS
	通过具体的值删除节点:
		删除头节点->PASS
		删除最后一个节点->PASS
		删除中间任意一个节点->PASS
 */

var doubleTestArray=[]interface{}{1,2,3,4,5,6,7,8}
var successAddArray=[]interface{}{1,2,3,4,5,6,7,8}			// 成功添加之后打印的结果

var(
	successRemoveHeadArray=[]interface{}{2,3,4,5,6,7,8} 	// 移除的元素是头元素
	successRemoveTailArray=[]interface{}{1,2,3,4,5,6,7}	// 移除的元素是尾元素
	removeIndex=4
	successRemoveMiddleArray=[]interface{}{1,2,3,4,6,7,8}	// 移除的元素是除了头和尾之外的任意元素
)

func TestDoubleSiededList_Add(t *testing.T) {
	doubleSidedList := NewDoubleSidedList()
	for _,value:=range doubleTestArray{
		doubleSidedList.Add(value)
	}
	assert.Equal(t,successAddArray,doubleSidedList.CollectResults())
}

// 测试删除的是尾节点
func TestDoubleSiededList_Remove_Head(t *testing.T) {
	doubleSidedList := NewDoubleSidedList()
	for _,value:=range doubleTestArray{
		doubleSidedList.Add(value)
	}
	doubleSidedList.Remove(0)
	assert.Equal(t,doubleSidedList.CollectResults(),successRemoveHeadArray)
}

// 测试删除的是尾节点
func TestDoubleSiededList_Remove_Tail(t *testing.T) {
	doubleSidedList := NewDoubleSidedList()
	for _,value:=range doubleTestArray{
		doubleSidedList.Add(value)
	}

	doubleSidedList.Remove(len(doubleTestArray)-1)
	assert.Equal(t,doubleSidedList.CollectResults(),successRemoveTailArray)
}
// 测试删除的是中间任意一个节点
func TestDoubleSiededList_Remove_Middle(t *testing.T) {
	doubleSidedList := NewDoubleSidedList()
	for _,value:=range doubleTestArray{
		doubleSidedList.Add(value)
	}

	doubleSidedList.Remove(4)
	assert.Equal(t,doubleSidedList.CollectResults(),successRemoveMiddleArray)
}

// 通过具体的值删除: 删除的是头节点
func TestDoubleSiededList_DeleteConcreteValue_Head(t *testing.T) {
	doubleSidedList := NewDoubleSidedList()
	for _,value:=range doubleTestArray{
		doubleSidedList.Add(value)
	}

	doubleSidedList.DeleteConcreteValue(1)
	assert.Equal(t,doubleSidedList.CollectResults(),successRemoveHeadArray)
}
// 通过具体的值删除:删除的是尾节点

func TestDoubleSiededList_DeleteConcreteValue_Tail(t *testing.T) {
	doubleSidedList := NewDoubleSidedList()
	for _,value:=range doubleTestArray{
		doubleSidedList.Add(value)
	}

	doubleSidedList.DeleteConcreteValue(8)
	assert.Equal(t,doubleSidedList.CollectResults(),successRemoveTailArray)
}

func TestDoubleSiededList_DeleteConcreteValue_Middle(t *testing.T) {
	doubleSidedList := NewDoubleSidedList()
	for _,value:=range doubleTestArray{
		doubleSidedList.Add(value)
	}

	doubleSidedList.DeleteConcreteValue(5)
	assert.Equal(t,doubleSidedList.CollectResults(),successRemoveMiddleArray)
}