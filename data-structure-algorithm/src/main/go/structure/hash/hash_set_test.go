package hash

import "testing"
// HashSet测试用例不全,只简单的测试了是否会添加重复元素

// 测试HashSet添加内容
func TestHashSet_Add(t *testing.T) {
	hashSet := NewHashSet()
	hashSet.Add(1)
	hashSet.Add(2)
	hashSet.Add(3)
	hashSet.show()
	hashSet.Add(1)
	hashSet.Add(4)
	hashSet.Add(3)
	hashSet.show()
}