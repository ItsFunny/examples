package chaincode

import (
	"bidchain/fabric/chaincode/ibidchain_contract"
	"bidchain/http_framework/protocol"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func StartChaincode(cc ibidchain_contract.IBidchainContract, cmdList []protocol.ICommand) {
	cc.SetChild(cc)
	// 缓存方法
	for _, cmd := range cmdList {
		protocol.RegisterCommand(cmd)
	//	cc.(*BasisContract).CacheMethod(cmd.GetFuncName())
	}
	err := shim.Start(cc.(shim.Chaincode))
	if err != nil {
		fmt.Println("failed to start chaincode: " + err.Error())
		return
	}
}
