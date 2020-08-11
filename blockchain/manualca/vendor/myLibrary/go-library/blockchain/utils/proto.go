/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-07 06:01 
# @File : proto.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/gogo/protobuf/proto"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
	"myLibrary/go-library/common/blockchain/base"
)

type TransactionDetail struct {
	From       base.From            `json:"from"`
	To         base.To              `json:"to"`
	Token      base.Token           `json:"token"`
	Version    base.Version         `json:"version"`
	TxBaseType base.TransBaseTypeV2 `json:"txBaseType"`
	// 描述
	TransactionDescription string `json:"transactionDescription"`
	// 是否是加密数据
	Encrypted bool `json:"encrypted"`

	// 遗留字段
	LeftBytes []byte `json:"leftBytes"`

	IsShow    bool   `json:"isShow"`
	Signature string `json:"signature"`

	TransactionID string   `json:"transactionId"`
	Args          []string `json:"args"`

	CreateTime int `json:"createTime"`
}

// 从SDK中Block.BlockDara.Data中提取交易具体信息
func GetTransactionInfoFromData(chainCodeId []string, data []byte, needArgs bool, desc func(base.TransBaseTypeV2) string) (TransactionDetail, error) {
	env, err := GetEnvelopeFromBlock(data)
	if err != nil {
		return TransactionDetail{}, fmt.Errorf("error extracting Envelope from block:%s", err.Error())
	}
	if env == nil {
		return TransactionDetail{}, fmt.Errorf("nil envelope")
	}
	detail, err := GetTransactionInfoFromEnvelop(env, chainCodeId, needArgs, desc)
	if nil != err {
		return detail, err
	}
	detail.Signature = hex.EncodeToString(env.Signature)
	return detail, err

}

func GetBlockDetail(chainCodeId string, data []byte) (isLogicCC bool, createdTime int64, amount float64, err error) {
	env, err := GetEnvelopeFromBlock(data)
	if err != nil {
		return
	}
	if env == nil {
		return
	}
	payload, err := GetPayload(env)

	if err != nil {
		return
	}
	channelHeaderBytes := payload.Header.ChannelHeader
	channelHeader := &cb.ChannelHeader{}
	if err = proto.Unmarshal(channelHeaderBytes, channelHeader); err != nil {
		return
	}
	createdTime = channelHeader.Timestamp.Seconds

	// chaincodeaction->chaincodeendorseredaction->proposalresponsepayload->chaincodeaction->txreadwriteset
	tx, err := GetTransaction(payload.Data)
	if err != nil {
		return
	}

	// 1. 获取ChaincodeActionPayload
	chaincodeActionPayload, err := GetChaincodeActionPayload(tx.Actions[0].Payload)
	if err != nil {
		return
	}
	// chaincodeAction:=pb.ChaincodeAction{}
	// 2. 获取ProposalResponsePayload
	extention := peer.ProposalResponsePayload{}
	// str := string(chaincodeActionPayload.Action.ProposalResponsePayload)
	// fmt.Println(str)
	if err = proto.Unmarshal(chaincodeActionPayload.Action.ProposalResponsePayload, &extention); nil != err {
		return
	}

	// 3. 获取ChaincodeAction
	chainCodeAction := peer.ChaincodeAction{}
	if err = proto.Unmarshal(extention.Extension, &chainCodeAction); nil != err {
		return
	}
	// 4. 准备TxReadWriteSet
	rwSet := rwset.TxReadWriteSet{}
	if err = proto.Unmarshal(chainCodeAction.Results, &rwSet); nil != err {
		return
	}
	// 5. 遍历解析数据
	for _, rs := range rwSet.NsRwset {
		if rs.Namespace == chainCodeId {
			isLogicCC = true
			kvrwset := kvrwset.KVRWSet{}
			if err = proto.Unmarshal(rs.Rwset, &kvrwset); nil != err {
				return
			}
			if len(kvrwset.Writes) > 0 {
				if len(kvrwset.Writes[0].Value) > 0 {
					node, valid := BuildRegularInfo(kvrwset.Writes[0].Value)
					// node, _ := GetRegularInfo(kvrwset.Writes[0].Value)
					if valid {
						amount += float64(node.Token)
					}
				}
			}
		}
	}

	return
}
func BuildRegularInfo(data []byte) (base.BCBaseNodeInfo, bool) {
	var result base.BCBaseNodeInfo
	defer func(r *base.BCBaseNodeInfo) (base.BCBaseNodeInfo, bool) {
		if nil != recover() {
			return base.BCBaseNodeInfo{}, false
		}
		return *r, true
	}(&result)

	result, _ = base.GetRegularInfoV2(data)

	return result, true
}

func GetBlockCreatedTime(data []byte) (int64, error) {
	env, err := GetEnvelopeFromBlock(data)
	if err != nil {
		return 0, fmt.Errorf("error extracting Envelope from block:%s", err.Error())
	}
	if env == nil {
		return 0, fmt.Errorf("nil envelope")
	}
	// 如果想获取时间等信息,可以在头部获取
	payload, err := GetPayload(env)
	if nil != err {
		return 0, fmt.Errorf("无法获取到对应的payload:%s", err.Error())
	}
	channelHeader := &cb.ChannelHeader{}
	if err := proto.Unmarshal(payload.Header.ChannelHeader, channelHeader); err != nil {
		return 0, fmt.Errorf("error extracting ChannelHeader from payload:%s", err.Error())
	}

	return channelHeader.Timestamp.Seconds, nil
}

// GetEnvelopeFromBlock gets an envelope from a block's Data field.
func GetEnvelopeFromBlock(data []byte) (*cb.Envelope, error) {
	// Block always begins with an envelope
	var err error
	env := &cb.Envelope{}
	if err = proto.Unmarshal(data, env); err != nil {
		return nil, errors.Wrap(err, "error unmarshaling Envelope")
	}

	return env, nil
}

// GetPayloads gets the underlying payload objects in a TransactionAction
func GetPayload(e *cb.Envelope) (*cb.Payload, error) {
	payload := &cb.Payload{}
	err := proto.Unmarshal(e.Payload, payload)
	return payload, errors.Wrap(err, "error unmarshaling Payload")
}

// GetChaincodeActionPayload Get ChaincodeActionPayload from bytes
func GetChaincodeActionPayload(capBytes []byte) (*peer.ChaincodeActionPayload, error) {
	cap := &peer.ChaincodeActionPayload{}
	err := proto.Unmarshal(capBytes, cap)
	return cap, errors.Wrap(err, "error unmarshaling ChaincodeActionPayload")
}

// GetProposalResponsePayload gets the proposal response payload
func GetProposalResponsePayload(prpBytes []byte) (*peer.ProposalResponsePayload, error) {
	prp := &peer.ProposalResponsePayload{}
	err := proto.Unmarshal(prpBytes, prp)
	return prp, errors.Wrap(err, "error unmarshaling ProposalResponsePayload")
}

// GetChaincodeAction gets the ChaincodeAction given chaicnode action bytes
func GetChaincodeAction(caBytes []byte) (*peer.ChaincodeAction, error) {
	chaincodeAction := &peer.ChaincodeAction{}
	err := proto.Unmarshal(caBytes, chaincodeAction)
	return chaincodeAction, errors.Wrap(err, "error unmarshaling ChaincodeAction")
}

// GetTransaction Get Transaction from bytes
func GetTransaction(txBytes []byte) (*peer.Transaction, error) {
	tx := &peer.Transaction{}
	err := proto.Unmarshal(txBytes, tx)
	return tx, errors.Wrap(err, "error unmarshaling Transaction")
}

func GetTransactionInfoFromEnvelop(envelop *cb.Envelope, chainCodeId []string, needArgs bool, desc func(base.TransBaseTypeV2) string) (TransactionDetail, error) {
	var (
		result TransactionDetail
	)
	result.Signature = hex.EncodeToString(envelop.Signature)
	// 如果想获取时间等信息,可以在头部获取
	payload, err := GetPayload(envelop)

	if err != nil {
		return result, fmt.Errorf("error extracting Payload from envelope:%s", err.Error())
	}
	channelHeaderBytes := payload.Header.ChannelHeader
	channelHeader := &cb.ChannelHeader{}

	if err := proto.Unmarshal(channelHeaderBytes, channelHeader); err != nil {
		return result, fmt.Errorf("error extracting ChannelHeader from payload:%s", err.Error())
	}

	// chaincodeaction->chaincodeendorseredaction->proposalresponsepayload->chaincodeaction->txreadwriteset
	tx, err := GetTransaction(payload.Data)
	if err != nil {
		return result, fmt.Errorf("error unmarshalling transaction payload:%s", err.Error())
	}
	// 1. 获取ChaincodeActionPayload
	chaincodeActionPayload, err := GetChaincodeActionPayload(tx.Actions[0].Payload)
	if err != nil {
		return result, fmt.Errorf("error unmarshalling chaincode action payload:%s", err.Error())
	}
	// chaincodeAction:=pb.ChaincodeAction{}
	// 2. 获取ProposalResponsePayload
	extention := peer.ProposalResponsePayload{}
	if err = proto.Unmarshal(chaincodeActionPayload.Action.ProposalResponsePayload, &extention); nil != err {
		return result, err
	}

	// 3. 获取ChaincodeAction
	chainCodeAction := peer.ChaincodeAction{}
	if err = proto.Unmarshal(extention.Extension, &chainCodeAction); nil != err {
		return result, err
	}
	// 4. 准备TxReadWriteSet
	rwSet := rwset.TxReadWriteSet{}
	if err = proto.Unmarshal(chainCodeAction.Results, &rwSet); nil != err {
		return result, err
	}
	// 5. 遍历解析数据
	for _, rs := range rwSet.NsRwset {
		equals := false
		for _, cid := range chainCodeId {
			if rs.Namespace == cid {
				equals = true
			}
		}
		if !equals {
			continue
		}
		equals = false

		kvrwset := kvrwset.KVRWSet{}
		if err := proto.Unmarshal(rs.Rwset, &kvrwset); nil != err {
			return result, err
		}
		// BUG here ,when the block is system block ,the writes will be pass
		// 有效交易的定义为,必须向块中写入数据,单纯的访问数据并不会算是有效交易,
		// 因此需要过滤无write的数据
		if len(kvrwset.Writes) == 0 || result.From != "" || result.To != "" {
			continue
		}

		// for _, k := range kvrwset.Reads {
		// 	fmt.Println(k.Key)
		// }

		node, valid := BuildRegularInfo(kvrwset.Writes[0].Value)
		if valid {
			result.TxBaseType = node.TxBaseType
			result.To = node.To
			result.From = node.From
			result.Version = node.Version
			result.Token = node.Token
			result.LeftBytes = node.LeftBytes
			result.Encrypted = node.Encrypted
			result.TransactionDescription = desc(node.TxBaseType)
			result.IsShow = true
		}
		break
	}
	// TRANSACTIONTYPE_ITEM_BUIED_LIST_UPDATE

	var (
		args []string
	)
	if needArgs {
		propPayload := &peer.ChaincodeProposalPayload{}
		if err := proto.Unmarshal(chaincodeActionPayload.ChaincodeProposalPayload, propPayload); err != nil {
			return result, fmt.Errorf("error extracting ChannelHeader from payload:%s", err.Error())
		}
		invokeSpec := &peer.ChaincodeInvocationSpec{}
		err = proto.Unmarshal(propPayload.Input, invokeSpec)
		if err != nil {
			return result, fmt.Errorf("error extracting ChannelHeader from payload:%s", err.Error())
		}
		for _, v := range invokeSpec.ChaincodeSpec.Input.Args {
			args = append(args, string(v))
		}
	}

	result.Args = args
	result.TransactionID = channelHeader.TxId
	result.CreateTime = int(channelHeader.Timestamp.Seconds)

	return result, nil
}
