package chaincode

import (
	"bidchain/fabric/context"
	"bidchain/fabric/log"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"reflect"
)

//type transport2StoreHandler func(message proto.Message) proto.Message
//type store2TransportHandler func(message proto.Message) proto.Message
//
//var (
//	transport2StoreHandlerMap = make(map[proto.Message]transport2StoreHandler)
//	store2TransportHandlerMap = make(map[proto.Message]store2TransportHandler)
//)
//
//func RegisterTransport2StoreHandler(messageType proto.Message, handler transport2StoreHandler) pb.Response {
//	if _, found := transport2StoreHandlerMap[messageType]; found {
//		msg := fmt.Sprintf("message type[%v] has already been registered", messageType)
//		log.Warn(ModuleName, msg)
//		return InternalServerError(ModuleName, msg)
//	}
//	transport2StoreHandlerMap[messageType] = handler
//	return shim.Success(nil)
//}
//
//func RegisterStore2TransportHandler(messageType proto.Message, handler store2TransportHandler) pb.Response {
//	if _, found := store2TransportHandlerMap[messageType]; found {
//		msg := fmt.Sprintf("message type[%v] has already bean registered", messageType)
//		log.Warn(ModuleName, msg)
//		return InternalServerError(ModuleName, msg)
//	}
//	store2TransportHandlerMap[messageType] = handler
//	return shim.Success(nil)
//}

// 转换/存储数据互转
type TransformHandler func(message proto.Message) proto.Message

var (
	transformHandlerMap = make(map[reflect.Type]TransformHandler)
	storeKeyPrefixMap   = make(map[reflect.Type]map[int64]string)
)

func RegisterTransformHandler(message proto.Message, handler TransformHandler) error {
	messageType := reflect.TypeOf(message)
	if _, found := transformHandlerMap[messageType]; found {
		msg := fmt.Sprintf("message type[%v] has already been registered", messageType)
		log.Warn(ModuleName, msg)
		return errors.New(msg)
	}
	transformHandlerMap[messageType] = handler
	return nil
}

// 根据数据类型，获取对应的转换函数
func GetTransformHandler(message proto.Message) TransformHandler {
	messageType := reflect.TypeOf(message)
	handler, found := transformHandlerMap[messageType]
	if !found {
		msg := fmt.Sprintf("message type[%v] has not  been registered", messageType)
		log.Warn(ModuleName, msg)
		return nil
	}
	return handler
}

// 以保函业务为例，输入状态有0,1,2(新增，变更，撤回),
//  而变更的本质是撤销+新增(原子操作), 撤销的数据和新增的数据的`OldStatus`都是1, 但是撤销数据的`ActualStatus`为2,而新增数据的`ActualStatus`为0
type ProtoMessageWrapper struct {
	proto.Message
	//// 原始状态 (根据原始状态和实际业务对数据进行验证)
	//OldStatus int64
	//// 实际状态(实际状态确定 存储前缀)
	//ActualStatus int64
	Status int64
	// 数据存储时对应的key
	StoreKey string
	// 数据存储转换函数
	TransformHandler TransformHandler
}

//func NewProtoMessageWrapper(message proto.Message) *ProtoMessageWrapper {
//	return &ProtoMessageWrapper{
//		Message: message,
//	}
//}

//func NewProtoMessageWrapper(ctx context.IBidchainContext, message proto.Message, oldStatus, actualStatus int64, getStoreKeyParams func(message proto.Message, params ...interface{}) []string, params ...interface{}) (*ProtoMessageWrapper, error) {
func NewProtoMessageWrapper(ctx context.IBidchainContext, message proto.Message, status int64, getStoreKeyParams func(message proto.Message, params ...interface{}) []string, params ...interface{}) (*ProtoMessageWrapper, error) {
	wrapper := &ProtoMessageWrapper{Message: message}
	wrapper.TransformHandler = GetTransformHandler(message)
	//wrapper.OldStatus = oldStatus
	//wrapper.ActualStatus = actualStatus
	//storeKeyParams := getStoreKeyParams(message, params)
	//storeKey, err := GetStoreKey(ctx, message, wrapper.ActualStatus, storeKeyParams)
	wrapper.Status = status
	storeKeyParams := getStoreKeyParams(message, params)
	if storeKeyParams == nil || len(storeKeyParams) == 0 {
		msg := fmt.Sprintf("storeKeyParams is nil or size equals zero")
		log.Warn(ModuleName, msg)
		return nil, errors.New(msg)
	}
	storeKey, err := GetStoreKey(ctx, message, status, storeKeyParams)
	if err != nil {
		return nil, err
	}
	wrapper.StoreKey = storeKey
	return wrapper, nil
}

func RegisterStoreKeyPrefix(message proto.Message, status int64, storeKeyPrefix string) error {
	messageType := reflect.TypeOf(message)
	if _, found := storeKeyPrefixMap[messageType]; !found {
		storeKeyPrefixMap[messageType] = make(map[int64]string)
	}
	if _, found := storeKeyPrefixMap[messageType][status]; found {
		msg := fmt.Sprintf("storeKeyPrefix[%v] 's status[%d] has alreday been registered", messageType, status)
		log.Warn(ModuleName, msg)
		return errors.New(msg)
	}
	storeKeyPrefixMap[messageType][status] = storeKeyPrefix
	return nil
}

func GetStoreKeyPrefix(message proto.Message, status int64) (string, error) {
	messageType := reflect.TypeOf(message)
	if _, found := storeKeyPrefixMap[messageType]; !found {
		msg := fmt.Sprintf("storeKeyPrefix[%v]  has not been registered", messageType)
		log.Warn(ModuleName, msg)
		return "", errors.New(msg)
	}
	storeKeyPrefix, found := storeKeyPrefixMap[messageType][status]
	if !found {
		msg := fmt.Sprintf("storeKeyPrefix[%v] 's status[%d] has not been registered", messageType, status)
		log.Warn(ModuleName, msg)
		return "", errors.New(msg)
	}
	return storeKeyPrefix, nil
}

// 获取存储数据的键
func GetStoreKey(ctx context.IBidchainContext, message proto.Message, status int64, params []string) (string, error) {
	prefix, err := GetStoreKeyPrefix(message, status)
	if err != nil {
		return "", err
	}
	key, err := ctx.CreateCompositeKey(prefix, params)
	return key, err
}
