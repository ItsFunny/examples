package context

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

type IBidchainContext interface {
	GetStub() shim.ChaincodeStubInterface
	GetState(key string) ([]byte, error)
	GetStateWithoutError(key string) string
	GetStateProtoMessage(key string, message proto.Message, exists *bool)
	PutState(key string, value []byte) error
	PutStateWithoutError(key string, value string)
	PutStateProtoMessage(key string, message proto.Message)
	DelState(key string) error
	GetFunctionAndParameters() (function string, params []string)
	SetFunctionName(name string)
	SetParameters(params []string)
	GetFunctionName() string
	GetParameters() []string
	GetChannelID() string
	SetChannelID(cid string)
	SetChaincodeName(cn string)
	GetChaincodeName() string
	GetStateByRange(startKey, endKey string) (shim.StateQueryIteratorInterface, error)
	InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response
	SetEvent(name string, payload []byte) error
	CreateCompositeKey(objectType string, attributes []string) (string, error)
	CreateCompositeKeyWithoutError(objectType string, attributes ...string) string


	SplitCompositeKey(compositeKey string) (string, []string, error)
	GetQueryResult(query string) (shim.StateQueryIteratorInterface, error)
	GetQueryResultWithPagination(query string, pageSize int32,
		bookmark string) (shim.StateQueryIteratorInterface, *pb.QueryResponseMetadata, error)
	GetHistoryForKey(key string) (shim.HistoryQueryIteratorInterface, error)
	GetStateByPartialCompositeKey(objectType string, keys []string) (shim.StateQueryIteratorInterface, error)
	GetTxID() string
	GetCreator() ([]byte, error)
	GetTxTimestamp() (*timestamp.Timestamp, error)
	GetTxTimestampLongValue() int64
}

type BidchainContext struct {
	stub          shim.ChaincodeStubInterface
	chaincodeName string
}

func NewBidchainContext(stub shim.ChaincodeStubInterface) IBidchainContext {
	return &BidchainContext{
		stub: stub,
	}
}

func (bcc *BidchainContext) GetStub() shim.ChaincodeStubInterface {
	return bcc.stub
}

func (bcc *BidchainContext) GetState(key string) ([]byte, error) {
	return bcc.stub.GetState(key)
}

func (bcc *BidchainContext) GetStateWithoutError(key string) string {
	data, err := bcc.stub.GetState(key)
	if err != nil {
		panic(err)
	}
	if data == nil {
		return ""
	}
	return string(data)
}

func (bcc *BidchainContext) PutState(key string, value []byte) error {
	return bcc.stub.PutState(key, value)
}

func (bcc *BidchainContext) PutStateWithoutError(key string, value string) {
	err := bcc.stub.PutState(key, []byte(value))
	if err != nil {
		panic(err)
	}
	return
}

func (bcc *BidchainContext) DelState(key string) error {
	return bcc.stub.DelState(key)
}

func (bcc *BidchainContext) GetFunctionAndParameters() (function string, params []string) {
	return bcc.stub.GetFunctionAndParameters()
}

func (bcc *BidchainContext) SetFunctionName(name string) {
	//"实际部署不应该调用SetFunctionName方法!"
	return
}

func (bcc *BidchainContext) SetParameters(params []string) {
	// "实际部署不应该调用SetParameters方法!"
	return
}

func (bcc *BidchainContext) GetFunctionName() string {
	function, _ := bcc.stub.GetFunctionAndParameters()
	return function
}

func (bcc *BidchainContext) GetParameters() []string {
	_, params := bcc.stub.GetFunctionAndParameters()
	return params
}

func (bcc *BidchainContext) GetChannelID() string {
	return bcc.stub.GetChannelID()
}

func (bcc *BidchainContext) SetChannelID(cid string) {
	// "实际部署不应该调用SetChannelID方法!"
	return
}

func (bcc *BidchainContext) SetChaincodeName(cn string) {
	bcc.chaincodeName = cn
	return
}

func (bcc *BidchainContext) GetChaincodeName() string {
	return bcc.chaincodeName
}

func (bcc *BidchainContext) GetStateByRange(startKey, endKey string) (shim.StateQueryIteratorInterface, error) {
	return bcc.stub.GetStateByRange(startKey, endKey)
}

func (bcc *BidchainContext) InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response {
	return bcc.stub.InvokeChaincode(chaincodeName, args, channel)
}

func (bcc *BidchainContext) SetEvent(name string, payload []byte) error {
	return bcc.stub.SetEvent(name, payload)
}

func (bcc *BidchainContext) GetQueryResult(query string) (shim.StateQueryIteratorInterface, error) {
	return bcc.stub.GetQueryResult(query)
}

func (bcc *BidchainContext) CreateCompositeKey(objectType string, attributes []string) (string, error) {
	return bcc.stub.CreateCompositeKey(objectType, attributes)
}

func (bcc *BidchainContext) CreateCompositeKeyWithoutError(objectType string, attributes ...string) string  {
	key, err :=  bcc.stub.CreateCompositeKey(objectType, attributes)
	if err != nil {
		panic(err)
	}
	return key
}


func (bcc *BidchainContext) SplitCompositeKey(compositeKey string) (string, []string, error) {
	return bcc.stub.SplitCompositeKey(compositeKey)
}
func (bcc *BidchainContext) GetStateByPartialCompositeKey(objectType string, keys []string) (shim.StateQueryIteratorInterface, error) {
	return bcc.stub.GetStateByPartialCompositeKey(objectType, keys)
}

func (bcc *BidchainContext) GetQueryResultWithPagination(query string, pageSize int32,
	bookmark string) (shim.StateQueryIteratorInterface, *pb.QueryResponseMetadata, error) {
	return bcc.stub.GetQueryResultWithPagination(query, pageSize, bookmark)
}

func (bcc *BidchainContext) GetHistoryForKey(key string) (shim.HistoryQueryIteratorInterface, error) {
	return bcc.stub.GetHistoryForKey(key)
}

func (bcc *BidchainContext) GetTxID() string {
	return bcc.stub.GetTxID()
}

func (bcc *BidchainContext) GetCreator() ([]byte, error) {
	return bcc.stub.GetCreator()
}

func (bcc *BidchainContext) GetTxTimestamp() (*timestamp.Timestamp, error) {
	txTimestamp, err := bcc.stub.GetTxTimestamp()
	var ts timestamp.Timestamp
	ts.Seconds = txTimestamp.Seconds
	ts.Nanos = txTimestamp.Nanos
	return &ts, err
}

func (bcc *BidchainContext) GetTxTimestampLongValue() int64 {
	ts, err := bcc.stub.GetTxTimestamp()
	if err != nil {
		err = errors.Wrap(err, "GetTxTimestamp")
		panic(err)
	}
	return ts.Seconds*1000 + int64(ts.Nanos/1000000)
}

func (bcc *BidchainContext) GetStateProtoMessage(key string, message proto.Message, exists *bool) {
	data, err := bcc.stub.GetState(key)
	if err != nil {
		panic(err)
	}
	if data == nil {
		*exists = false
		return
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		panic(err)
	}
	*exists = true
}

func (bcc *BidchainContext) PutStateProtoMessage(key string, message proto.Message) {
	data, err := proto.Marshal(message)
	if err != nil {
		panic(err)
	}
	err = bcc.stub.PutState(key, data)
	if err != nil {
		panic(err)
	}
}
