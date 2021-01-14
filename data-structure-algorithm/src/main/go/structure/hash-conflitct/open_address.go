package hash_conflitct

import (
	"fmt"
	"github.com/mitchellh/hashstructure"
	"github.com/pkg/errors"
)

// 在hash开放地址法中用于存放值和便于判断是否删除
type HashValueWrapper struct {
	value   interface{}
	deleted bool
}

// 通过开放地址法解决hash冲突
// 名字待定
type HashOpenAddress struct {
	maxDataSize int
	// 用定长数组存储值
	datas [1024]*HashValueWrapper
	// 代表溢出池的cursor,仅用于溢出池,实际最好抽出成一个对象
	cursor int
	// 用切片充当溢出池
	overFlowPool []*HashValueWrapper
}

func NewHashOpenAddress() *HashOpenAddress {
	hashOpenAddress := &HashOpenAddress{
		maxDataSize:  1024,
		overFlowPool: nil,
	}
	return hashOpenAddress
}

func (h *HashOpenAddress) Put(value interface{}) (bool, error) {
	sprintf := fmt.Sprintf("%v", value)
	hash, err := hashstructure.Hash(sprintf, nil)
	if nil != err {
		return false, errors.New("hash occur error")
	}
	index := hash % uint64(h.maxDataSize)
	valueWrapper := &HashValueWrapper{
		value:   value,
		deleted: false,
	}
	if oldValue := h.datas[index]; nil != oldValue {
		valueWrapper := valueWrapper
		h.datas[index] = valueWrapper
		return true, nil
	}
	for i := int(index); i < h.maxDataSize; i++ {
		if nil == h.datas[i] || h.datas[i].deleted {
			h.datas[i] = valueWrapper
			return true, nil
		}
	}
	// 插入到溢出池中
	h.cursor++
	h.overFlowPool[h.cursor] = valueWrapper
	return true, nil
}

func (h *HashOpenAddress) Find(value interface{}) (bool, error) {
	hashCode, err := hashstructure.Hash(fmt.Sprintf("%v", value), nil)
	// hashString := util.Hash(fmt.Sprintf("%v", value))
	// hashCode, err := strconv.Atoi(hashString)
	if nil != err {
		return false, err
	}
	index :=hashCode % uint64(h.maxDataSize)
	valueWrapper := h.datas[index]

	// 从当前的槽位遍历
	for i := int(index); i < h.maxDataSize; i++ {
		if valueWrapper = h.datas[i]; nil != valueWrapper && !valueWrapper.deleted && valueWrapper.value == value {
			return true, nil
		}
	}
	// 说明之前的池中没有,则尝试从溢出池中判断是否存在
	for i := 0; i < len(h.overFlowPool); i++ {
		if valueWrapper = h.overFlowPool[i]; nil != valueWrapper && !valueWrapper.deleted && valueWrapper.value == value {
			return true, nil
		}
	}
	return false, nil
}
