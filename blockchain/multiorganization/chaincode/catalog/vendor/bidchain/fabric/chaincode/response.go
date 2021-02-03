package chaincode

import (
	"bidchain/base/jsonutils"
	"bidchain/fabric/bserror"
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/log"
	"bidchain/protocol/transport"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/hyperledger/fabric/protos/peer"
)


// 这是老的，以后使用command就不要调用这个了
func Error(moduleName bsmodule.Module, bsError *bserror.BSError, desc string, message proto.Message) pb.Response {
	code := bsError.Code
	var data []byte
	var err error
	if message != nil {
		data, err = jsonutils.ProtoMarshal(message)
		if err != nil {
			log.Warnf(moduleName, "failed to jsonutils.ProtoMarshal message[%s]", message)
			return ErrBlockchainInternal
		}
	}

	resp := &transport.CommonResponse{
		Code: int32(code),
		Desc: desc,
		Ret:  data,
	}
	data, err = jsonutils.ProtoMarshal(resp)
	if err != nil {
		log.Warnf(moduleName, "failed to jsonutils.ProtoMarshal message[%s]", message)
		return ErrBlockchainInternal
	}
	return pb.Response{
		Status:  int32(code),
		Message: desc,
		Payload: data,
	}
}

// 请求错误
// 如输入参数不合法: 参数个数，类型不匹配
// 序列化失败
func BadRequestError(moduleName bsmodule.Module, message string) pb.Response {
	return Error(moduleName, bserror.BAD_REQUEST, message, nil)
}

func BadRequestErrorWithMessage(moduleName bsmodule.Module, desc string, message proto.Message) pb.Response {
	return Error(moduleName, bserror.BAD_REQUEST, desc, message)
}

// 内部错误
// 比如PutState错误
func InternalServerError(moduleName bsmodule.Module, message string) pb.Response {
	return Error(moduleName, bserror.INTERNAL_SERVER_ERROR, message, nil)
}

// 403 权限不够
func ForbiddenError(moduleName bsmodule.Module, message string) pb.Response {
	return Error(moduleName, bserror.FORBIDDEN, message, nil)
}

func JsonMarshalError(moduleName bsmodule.Module, err error) pb.Response {
	msg := fmt.Sprintf("json.Marshal error: %s", err.Error())
	log.Warn(moduleName, msg)
	return InternalServerError(moduleName, msg)
}

func JsonUnMarshalError(moduleName bsmodule.Module, err error) pb.Response {
	msg := fmt.Sprintf("json.Unmarshal error: %s", err.Error())
	log.Warn(moduleName, msg)
	return InternalServerError(moduleName, msg)
}

func ProtoMarshalError(moduleName bsmodule.Module, err error) pb.Response {
	msg := fmt.Sprintf("proto.Marshal error: %s", err.Error())
	log.Warn(moduleName, msg)
	return InternalServerError(moduleName, msg)
}

func ProtoUnMarshalError(moduleName bsmodule.Module, err error) pb.Response {
	msg := fmt.Sprintf("proto.Unmarshal error: %s", err.Error())
	log.Warn(moduleName, msg)
	return InternalServerError(moduleName, msg)
}

func GetStateError(moduleName bsmodule.Module, err error) pb.Response {
	msg := fmt.Sprintf("GetState error: %s", err.Error())
	log.Warn(moduleName, msg)
	return InternalServerError(moduleName, msg)
}

func PutStateError(moduleName bsmodule.Module, err error) pb.Response {
	msg := fmt.Sprintf("PutState error: %s", err.Error())
	log.Warn(moduleName, msg)
	return InternalServerError(moduleName, msg)
}

func GetStateByPartialCompositeKeyError(moduleName bsmodule.Module, err error) pb.Response {
	msg := fmt.Sprintf("failed to GetStateByPartialCompositeKey, desc=[%s]", err.Error())
	log.Warn(moduleName, msg)
	return InternalServerError(moduleName, msg)
}

func SplitCompositeKeyError(moduleName bsmodule.Module, err error) pb.Response {
	msg := fmt.Sprintf("failed to SplitCompositeKey, desc=[%s]", err.Error())
	log.Warn(moduleName, msg)
	return InternalServerError(moduleName, msg)
}



func CreateCompositeKeyError(moduleName bsmodule.Module, err error) pb.Response {
	msg := fmt.Sprintf("failed to CreateCompositeKey, desc=[%s]", err.Error())
	log.Warn(moduleName, msg)
	return InternalServerError(moduleName, msg)
}

// 无效时间
func InvalidTimeError(module bsmodule.Module, err error, message proto.Message) pb.Response {
	msg := err.Error()
	return Error(module, bserror.INVALID_TIME, msg, message)
}

// 字段为空错误
func EmptyFieldError(module bsmodule.Module, err error, message proto.Message) pb.Response {
	desc := err.Error()
	return Error(module, bserror.EMPTY_FIELD, desc, message)
}

// 无效的加密副本信息
func InValidEncryptedMirrorData(moduleName bsmodule.Module, err error, message proto.Message) pb.Response {
	desc := err.Error()
	return Error(moduleName, bserror.INVALID_ENCRYPTED_MIRRORDATA, desc, message)
}


