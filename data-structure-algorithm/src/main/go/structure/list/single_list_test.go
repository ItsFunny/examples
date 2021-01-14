package list

import (
	"fmt"
	"github.com/influxdata/influxdb/pkg/testing/assert"
	"testing"
)

func TestSingleList_AddNode(t *testing.T) {
	singleList := &SingleList{}
	values := []interface{}{1, 2, 3, 4, 5, 6}
	for _, v := range values {
		singleList.AddNode(v)
	}
	resultList := singleList.IteratorNode()
	for _, v := range resultList {
		fmt.Printf("%d->", v)
	}
	assert.Equal(t, resultList, values)
}
func TestSingleList_Remove(t *testing.T) {
	singleList := &SingleList{}
	values := []interface{}{1, 2, 3, 4, 5, 6}
	for _, v := range values {
		singleList.AddNode(v)
	}

	values = []interface{}{1, 2, 3, 4, 5}
	singleList.Remove()
	resultList := singleList.IteratorNode()
	assert.Equal(t, resultList, values)
}
