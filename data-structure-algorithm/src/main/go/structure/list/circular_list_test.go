package list

import (
	"fmt"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

var testNormalArray = []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
var testEmptyArray = []interface{}{}
var testSingleValueArray = []interface{}{0}

func insertNodes(arr []interface{}) *CircularList {
	circularList := &CircularList{}
	for _, v := range arr {
		circularList.AddNode(v)
	}
	return circularList
}
func TestCircularList_AddNode(t *testing.T) {
	circularList := insertNodes(testNormalArray)
	values, e := circularList.IterateNode()
	if nil != e {
		panic(e)
	}
	for _, v := range values {
		fmt.Printf("%d->", v)
	}
	fmt.Println()
	successArray := []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
	assert.Equal(t, values, successArray)
}

func TestCircularList_AddNode_Empty(t *testing.T) {
	circularList := &CircularList{}
	for _, value := range testEmptyArray {
		circularList.AddNode(value)
	}
	_, e := circularList.IterateNode()
	if nil == e {
		panic("the return error should not be null")
	}
}

func TestCircularList_Remove(t *testing.T) {
	// 1,2,3,4,5,6,7,8
	circularList := insertNodes(testNormalArray)
	remove, e := circularList.Remove()
	if nil != e {
		panic(e)
	}
	fmt.Println("remove value:", remove)
	assert.Equal(t, remove, 8)
}

func TestCircularList_Remove_emptyList(t *testing.T) {
	circularList := &CircularList{}
	_, e := circularList.Remove()
	if nil == e {
		panic("should not be here")
	}
}

// 测试输入1,2,3,4,5,6,7,8 然后移除head
func TestCircularList_RemoveByIndex(t *testing.T) {
	// 1,2,3,4,5,6,7,8
	circularList := insertNodes(testNormalArray)
	removeValue, e := circularList.RemoveByIndex(1)
	if nil != e {
		panic(e)
	}
	fmt.Println("remove value:", removeValue)
	values, e := circularList.IterateNode()
	if nil != e {
		panic(e)
	}
	assert.Equal(t, removeValue, 1)
	assert.Equal(t, values, []interface{}{2, 3, 4, 5, 6, 7, 8})
}

// 测试只输入1个值,然后移除的是超过这个值
// 正常情况下应该是正常移除1 然后遍历的时候是NullPointerException
func TestCircularList_RemoveByIndex_OneValue(t *testing.T) {
	circularList := insertNodes([]interface{}{1})
	removeValue, e := circularList.RemoveByIndex(2)
	if nil != e {
		panic(e)
	}
	fmt.Println("remove value:", removeValue)
	values, e := circularList.IterateNode()
	if nil != e {
		assert.Equal(t, e.Error(), "NullPointerException")
	}
	for _, value := range values {
		fmt.Println(value)
	}
}

func TestCircularList_IterateNode(t *testing.T) {

}
