package protocol

import (
	"bidchain/base/base64utils"
	"bidchain/base/jsonutils"
	"bidchain/fabric/bserror"
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/context"
	"bidchain/fabric/log"
	"bidchain/http_framework/types"
	"bidchain/protocol/transport"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"reflect"
)

const (
	ERROR_RESPONSE_WITH_HTTP_COMMAND_PREFIX = "ERROR_RESPONSE_WITH_HTTP_COMMAND_PREFIX_"
)

var (
	moduleName = bsmodule.HTTP_FRAMEWORK_PROTOCOL
)



type AbstractCommand struct {
	ICommonCommand               // 通用函数
	Cmd           ICustomCommand // 子类实现函数
	head          types.Head
	couple        ICommand
	channelName   string
	chaincodeName string
	funcName      string
	//requestType        RequestType
	RequestType        RequestType
	errBody            interface{}
	receiveMilliSecond int64
	bSError            *bserror.BSError
	httpCmdContext     *context.HttpCommandContext // 上下文用来存储必要的数据
}

//func NewAbstractCommand() *AbstractCommand {
//	ac := &AbstractCommand{}
//
//	return ac
//}

//// request和response的couple是它们对方，publish的couple是null
//func (ac *AbstractCommand) GetCouple() reflect.Type {
//	// golang中没有注解， 这个让具体的command提供, command提供的结果通过元数据自动生成
//	panic("unimplemented method, must be implemented by child")
//}

// request和response能够new出对方的command 并且会自动为设置好包头所有的值
func (ac *AbstractCommand) NewCouple() ICommand {
	// 子类实现了GetCoupleType
	t := ac.Cmd.GetCoupleType()
	v := reflect.New(t)
	abstractCmd := &AbstractCommand{
		Cmd: v.Interface().(ICustomCommand),
	}
	v.Elem().FieldByName("AbstractCommand").Set(reflect.ValueOf(abstractCmd))
	coupleCmd := v.Interface().(ICommand)
	ac.couple = coupleCmd
	// TODO response和request建立关系
	// 关系的建立应该在外层了
	return coupleCmd
}

// 请求处理时间(单位毫秒)
func (ac *AbstractCommand) SetReceiveMillisecond(millisecond int64) {
	ac.receiveMilliSecond = millisecond
}

func (ac *AbstractCommand) GetReceiveMillisecond() int64 {
	return ac.receiveMilliSecond
}

//func (ac *AbstractCommand) GetHead() *types.Head {
//	return ac.head
//}
//
//func (ac *AbstractCommand) SetHead(head *types.Head) {
//	ac.head = head
//}

// go 端调试使用
func (ac *AbstractCommand) ToJson() (string, error) {
	packet, err := ac.generatePacket()
	if err != nil {
		return "", err
	}
	data, err := jsonutils.Marshal(packet)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// 响应的结果
func (ac *AbstractCommand) ToBytes() ([]byte, error) {
	packet, err := ac.generatePacket()
	if err != nil {
		return nil, err
	}
	data, err := proto.Marshal(packet)
	return data, err
}

// 生成protobuf
// 序列化，打印需要用到
func (ac *AbstractCommand) generatePacket() (proto.Message, error) {
	var body []byte
	var err error
	if ac.GetErrCode() != 0 {
		body, err = getBodyBytes(ac.GetErrBody())
		if err != nil {
			return nil, err
		}
	} else {
		//bs, err := ac.Cmd.SerializeParams()
		bs, err := SerializeParams(ac.Cmd)
		if err != nil {
			return nil, err
		}
		body = bs
	}
	//bodyBytes, err := getBodyBytes(body)
	//if err != nil {
	//	return nil, err
	//}

	message := &transport.Packet{
		Header: &transport.Packet_Header{
			ErrorCode: ac.GetErrCode(),
			ErrorDesc: ac.GetErrDesc(),
		},
		Body: body,
	}
	return message, nil
}

//func (ac *AbstractCommand) SetFuncName(funcName string) {
//	ac.FuncName = funcName
//}

//func (ac *AbstractCommand) SetChannelName(channelName string) {
//	ac.ChannelName = channelName
//}

//func (ac *AbstractCommand) SetChaincodeName(chaincodeName string) {
//	ac.ChaincodeName = chaincodeName
//}
//
//
//func (ac *AbstractCommand) GetChannelName() string {
//	return ac.ChannelName
//}
//
//func (ac *AbstractCommand) GetChaincodeName() string {
//	return ac.ChaincodeName
//}
//
//func (ac *AbstractCommand) GetFuncName() string {
//	return ac.FuncName
//}

func (ac *AbstractCommand) GetChannelName() string {
	return ac.httpCmdContext.ChannelName
}

func (ac *AbstractCommand) GetChaincodeName() string {
	return ac.httpCmdContext.ChaincodeName
}

//func (ac *AbstractCommand) GetFuncName() string {
//	return ac.HttpCmdContext.FuncName
//}

func (ac *AbstractCommand) SetHttpCommandContext(ctx *context.HttpCommandContext) {
	ac.httpCmdContext = ctx
}

func (ac *AbstractCommand) GetHttpCommandContext() *context.HttpCommandContext {
	return ac.httpCmdContext
}

// 请求类型
func (ac *AbstractCommand) GetRequestType() RequestType {
	return ac.RequestType
}

// 请求类型
//func (ac *AbstractCommand) GetRequestType() RequestType {
//	return ac.requestType
//}

func (ac *AbstractCommand) SetRequestType(requestType RequestType) {
	ac.RequestType = requestType
}

func (ac *AbstractCommand) SetErrCode(errCode int64) {
	ac.head.SetErrCode(errCode)
}

func (ac *AbstractCommand) GetErrCode() int64 {
	return ac.head.GetErrCode()
}

func (ac *AbstractCommand) SetErrDesc(errDesc string) {
	ac.head.SetErrDesc(errDesc)
}
func (ac *AbstractCommand) GetErrDesc() string {
	return ac.head.GetErrDesc()
}

func (ac *AbstractCommand) SetErrBody(body interface{}) {
	ac.errBody = body
}

func (ac *AbstractCommand) GetErrBody() interface{} {
	return ac.errBody
}

func (ac *AbstractCommand) SetBSError(err *bserror.BSError) {
	ac.bSError = err
	ac.SetErrCode(err.Code)
	ac.SetErrDesc(err.Desc)
}
func (ac *AbstractCommand) GetBSError() *bserror.BSError {
	return ac.bSError
}

func (ac *AbstractCommand) GenerateResponse() (resp pb.Response, errCode int64, errDesc string) {
	data, err := ac.ToBytes()
	if err != nil {
		panic(err)
	}
	// 为了和之前的代码做兼容 这里做统一处理
	if ac.GetErrCode() == 0 {
		return shim.Success(data), 0,""
	} else {
		s := base64utils.Base64Encode(data)
		s = ERROR_RESPONSE_WITH_HTTP_COMMAND_PREFIX + s
		return shim.Error(s), ac.GetErrCode(), ac.GetErrDesc()
	}
}

func getBodyBytes(data interface{}) ([]byte, error) {
	// 判断是否为空
	if data == nil {
		return nil, nil
	}
	// 如果是proto.Message
	if val, ok := data.(proto.Message); ok {
		return jsonutils.ProtoMarshal(val)
	}

	// 这里和java通信，序列化的数据要目是protobuf序列化的字节数组，原样返回
	// 要木是基本数据类型 字符串，整数，布尔值，浮点数
	if reflect.TypeOf(data) == reflect.TypeOf([]byte(nil)) {
		return data.([]byte), nil
	}

	//var buf proto.Buffer
	switch reflect.TypeOf(data).Kind() {
	case reflect.String:
		//buf.EncodeStringBytes(data.(string))
		//return buf.Bytes(), nil
		return []byte(data.(string)), nil
	//case reflect.Bool:
	//	// 转为字符串比较方便
	//
	//case reflect.Int32:
	//case reflect.Int64:
	//case reflect.Uint32:
	//case reflect.Uint64:
	//case reflect.Float32:
	//case reflect.Float64:
	default:
		msg := fmt.Sprintf("invalid type: %v", reflect.TypeOf(data))
		log.Warn(moduleName, msg)
		return nil, errors.New(msg)
	}
}
