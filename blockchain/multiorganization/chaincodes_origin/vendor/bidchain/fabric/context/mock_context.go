package context

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type MockContext struct {
	BidchainContext
	funcName  string
	params    []string
	channelID string
}

func NewMockContext(stub shim.ChaincodeStubInterface) IBidchainContext {
	newStub, ok := stub.(*shim.MockStub)
	if !ok {
		panic("本地测试必须必须初始化MockStub")
	}
	newStub.TxID = "1"
	ctx := &MockContext{
		BidchainContext: BidchainContext{stub: newStub},
		channelID:       DEFAULT_CHANNEL_ID,
	}
	return ctx
}

func (sc *MockContext) GetFunctionAndParameters() (function string, params []string) {
	if sc.funcName == "" {
		panic("本次测试需要先调用SetFunctionName和SetParameters，再获取对应的函数名和方法名称")
	}
	// chaincode 链码间调用
	if sc.params == nil {
		return sc.stub.GetFunctionAndParameters()
	}
	return sc.funcName, sc.params
}

func (sc *MockContext) SetFunctionName(name string) {
	sc.funcName = name
}

func (sc *MockContext) SetParameters(params []string) {
	sc.params = params
}

func (sc *MockContext) GetFunctionName() string {
	return sc.funcName
}

func (sc *MockContext) GetParameters() []string {
	return sc.params
}

func (sc *MockContext) GetChannelID() string {
	return sc.channelID
}

func (sc *MockContext) SetChannelID(cid string) {
	sc.channelID = cid
}

func (sc *MockContext) InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response {
	if len(args) < 1 {
		panic("InvokeChaincode 参数不足， 必须指定调用的链码函数名称")
	}

	return sc.stub.InvokeChaincode(chaincodeName, args, channel)
}

