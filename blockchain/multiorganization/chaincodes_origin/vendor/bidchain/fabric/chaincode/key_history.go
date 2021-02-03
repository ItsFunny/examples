package chaincode

import (
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/context"
	"bidchain/fabric/log"
	"bidchain/protocol/transport/base"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
	"runtime"
	"time"
)

type IKeyHistoryIterator interface {
	HasNext() (bool, *queryresult.KeyModification)
	Accept(kM *queryresult.KeyModification, index int) bool
}

type IKeyHistory interface {
	GetKeyHistory(key string, keyHistoryIterator IKeyHistoryIterator) []*queryresult.KeyModification
}

type KeyHistoryIteratorImpl struct {
	Ctx                   context.IBidchainContext
	Key                   string
	HistoryQueryInterface shim.HistoryQueryIteratorInterface
	Position              int
}

func NewKeyHistoryIteratorImpl(ctx context.IBidchainContext, key string) *KeyHistoryIteratorImpl {
	iteratorInterface, err := ctx.GetHistoryForKey(key)
	if err != nil {
		panic(err)
	}
	return &KeyHistoryIteratorImpl{
		Ctx:                   ctx,
		Key:                   key,
		HistoryQueryInterface: iteratorInterface,
		Position:              0,
	}
}

func (csImpl *KeyHistoryIteratorImpl) HasNext() (bool, *queryresult.KeyModification) {
	if csImpl.HistoryQueryInterface.HasNext() {
		csImpl.Position++
		modification, _ := csImpl.HistoryQueryInterface.Next()
		return true, modification
	} else {
		if csImpl.HistoryQueryInterface != nil {
			csImpl.HistoryQueryInterface.Close()
		}
		return false, nil
	}

}

func (csImpl *KeyHistoryIteratorImpl) Accept(kM *queryresult.KeyModification, index int) bool {
	return true
}

type KeyHistoryImpl struct {
	ctx context.IBidchainContext
}

func (khImpl *KeyHistoryImpl) GetKeyHistory(key string, keyHistoryIterator IKeyHistoryIterator) []*queryresult.KeyModification {
	var ret []*queryresult.KeyModification
	index := 0
	for hasNext, keyModification := keyHistoryIterator.HasNext(); hasNext; hasNext, keyModification = keyHistoryIterator.HasNext() {
		if keyHistoryIterator.Accept(keyModification, index) {
			ret = append(ret, keyModification)
		}
		index++
	}
	return ret
}

type keyFirstHistoryIterator struct {
	*KeyHistoryIteratorImpl
}

func (kfHI *keyFirstHistoryIterator) HasNext() (bool, *queryresult.KeyModification) {
	if kfHI.Position == 0 {
		return kfHI.KeyHistoryIteratorImpl.HasNext()
	}
	return false, nil
}

func newKeyFirstHistoryIterator(ctx context.IBidchainContext, key string) *keyFirstHistoryIterator {
	var iterator keyFirstHistoryIterator
	iterator.KeyHistoryIteratorImpl = NewKeyHistoryIteratorImpl(ctx, key)
	return &iterator
}

// 获取key对应的所有历史记录
func GetKeyAllHistory(ctx context.IBidchainContext, key string) []*queryresult.KeyModification {
	var keyHistoryImpl KeyHistoryImpl
	iterator := NewKeyHistoryIteratorImpl(ctx, key)
	ret := keyHistoryImpl.GetKeyHistory(key, iterator)
	if len(ret) > 0 {
		var msg = "{\n\"" + key + "\":\t[\n"
		for _, modification := range ret {
			hashValue := sha256.Sum256(modification.Value)
			msg += fmt.Sprintf("{txID: %s, value=%s, valueHash: %s, timestamp: {%v}}\n", modification.TxId, string(modification.Value), hex.EncodeToString(hashValue[:]), modification.Timestamp)
		}
		msg += "]}"
		log.Info(bsmodule.KEY_HISTORY, msg)
	} else {
		msg := fmt.Sprintf("key[%s] doesn't exists", key)
		log.Info(bsmodule.KEY_HISTORY, msg)
	}

	return ret
}

// 获取key对应的第一条记录
//func GetKeyFirstHistory(ctx context.IBidchainContext, key string) *queryresult.KeyModification {
//	var keyHistoryImpl KeyHistoryImpl
//	iterator := newKeyFirstHistoryIterator(ctx, key)
//	data := keyHistoryImpl.GetKeyHistory(key, iterator)
//	if len(data) >= 1 {
//		modification := data[0]
//		hashValue := sha256.Sum256(modification.Value)
//		msg := fmt.Sprintf("{key: %s, txID: %s, valueHash: %s, timestampe: {%v}}\n", key, modification.TxId, hex.EncodeToString(hashValue[:]), modification.Timestamp)
//		log.Info(bsmodule.KEY_HISTORY, msg)
//		return data[0]
//	} else {
//		msg := fmt.Sprintf("key[%s] doesn't exists", key)
//		log.Info(bsmodule.KEY_HISTORY, msg)
//		return nil
//	}
//
//}

// 获取key对应的第一条记录
func GetFirstKeyUpChainTime(ctx context.IBidchainContext, key string) *base.LongWrapper {
	// 测试开发
	if runtime.GOOS == "windows" {
		v := time.Now().UnixNano() / 1000000
		return &base.LongWrapper{Value: v}
	}
	iteratorInterface, err := ctx.GetHistoryForKey(key)
	if err != nil {
		panic(err)
	}
	var ret *queryresult.KeyModification
	for iteratorInterface.HasNext() {
		modification, _ := iteratorInterface.Next()
		ret = modification
	}
	if ret == nil {
		return nil
	}
	ts := ret.Timestamp
	ms := ts.Seconds*1000 + int64(ts.Nanos/1000000)
	return &base.LongWrapper{Value: ms}
}
