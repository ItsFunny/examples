package chaincode

import (
	"bidchain/fabric/context"
	"bidchain/fabric/log"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// 通过chaincode 查询对应的channel文件是否存在
func IsIpfshFileHashInfoExistByChaincodeQueryInternal(ctx context.IBidchainContext, channelName string, ipfsHashList []string) []bool {
	//return []bool {true}
	chaincodeName := "ipfs_filehash"
	funcName := "IsIpfshFileHashInfoExistByChaincodeQueryInternal"

	data, err := json.Marshal(ipfsHashList)
	if err != nil {
		panic(err)
	}
	resp := ctx.InvokeChaincode(chaincodeName, [][]byte{[]byte(funcName), data}, channelName)
	if resp.Status != shim.OK {
		msg := fmt.Sprintf("call channel[%s], chaincode[%s], funcName[%s] failed,  params=[%s] desc=[%s]", channelName, chaincodeName, funcName, ipfsHashList, resp.Message)
		log.Warn(ModuleName, msg)
		panic(msg)
	}

	var existList []bool
	err = json.Unmarshal(resp.Payload, &existList)
	if err != nil {
		panic(err)
	}

	return existList
}
