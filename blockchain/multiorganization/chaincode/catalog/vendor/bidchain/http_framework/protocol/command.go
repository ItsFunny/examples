package protocol

import (
	"bidchain/fabric/bserror"
	"bidchain/fabric/context"
	"github.com/golang/protobuf/proto"

	"reflect"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ICommand interface {
	ICommonCommand
	ICustomCommand
}

// 抽象类实现
type ICommonCommand interface {
	// request和response能够new出对方的command 并且会自动为设置好包头所有的值
	NewCouple() ICommand  // 对Request而言
	//GetCouple() ICommand // 对于Response而言
	SetHttpCommandContext(ctx *context.HttpCommandContext)
	GetHttpCommandContext() *context.HttpCommandContext

	// 请求处理时间(单位毫秒)
	SetReceiveMillisecond(millisecond int64)

	GetReceiveMillisecond() int64

	//GetHead() *types.Head
	//
	//SetHead(head *types.Head)
	ToJson() (string, error)

	ToBytes() ([]byte, error)

	GetChannelName() string
	//SetChannelName(name string)
	GetChaincodeName() string
	//SetChaincodeName(chaincodeName string)
	//SetFuncName(funcName string)
	GetFuncName() string
	// 请求类型
	GetRequestType() RequestType
	SetRequestType(requestType RequestType)


	SetErrCode(errCode int64)
	GetErrCode() int64
	SetErrDesc(errDesc string)
	GetErrDesc() string
	SetBSError(err *bserror.BSError)
	GetBSError() *bserror.BSError
	SetErrBody(body interface{})
	GetErrBody() interface{}

	// response做的 生成pb.Response
	GenerateResponse() (resp pb.Response, errCode int64, errDesc string)
}

// 子类实现
type ICustomCommand interface {
	// request和response的couple是它们对方，publish的couple是null
	GetCoupleType() reflect.Type
	// 反序列化[]byte 对应的请求参数
	//DeserializeParams(data []byte) error
	// 将相应结果反序列化为对应的[]byte
	//SerializeParams() ([]byte, error)
	GetURI() string // 路径 <chaincodeName>/<funcName>
	GetParameters() proto.Message
}

type RequestType int

const (
	RequestType_QUERY RequestType = iota
	RequestType_INVOKE
)

