package chaincode

import (
	"bidchain/base/jsonutils"
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/context"
	"bidchain/fabric/log"
	"bidchain/protocol/transport/base"
	"bidchain/protocol/transport/datatable_metadata"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"reflect"
	"strings"
)

var (
	publicInfoKeyMap = map[string]interface{}{
		"Info":       nil,
		"PublicInfo": nil,
	}
	privateInfoKeyMap = map[string]interface{}{
		"Message":       nil,
		"CryptoMessage": nil,
	}
)

func isPublicInfo(s string) bool {
	_, ok := publicInfoKeyMap[s]
	return ok
}

func isCryptoMessage(s string) bool {
	_, ok := privateInfoKeyMap[s]
	return ok
}

func validatePublicInfo(module bsmodule.Module, publicInfo reflect.Value, message proto.Message, metadata *datatable_metadata.DataTableMetadata, skipVerifyPublicFieldList []string) pb.Response {
	// 验证publicInfo是否为空
	if publicInfo.Pointer() == 0 {
		desc := fmt.Sprintf("publicInfo is null, message=[%v]", message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}
	publicInfo = publicInfo.Elem()
	var publicFieldMap = make(map[string]*datatable_metadata.DataTableFieldMetadata)
	// 同时有公开信息和加密信息
	for _, publicFieldMetadata := range metadata.PublicFieldList {
		name := strings.Title(publicFieldMetadata.FieldName)
		publicFieldMap[name] = publicFieldMetadata
	}

	// 3个字段是protobuf自带的: XXX_NoUnkeyedLiteral、XXX_unrecognized、XXX_sizecache
	if publicInfo.NumField()-3 != len(publicFieldMap) {
		desc := fmt.Sprintf("publicInfo fieldNum[%d] differs from the metadataTable definition[%d], message=[%v]", publicInfo.NumField()-3, len(publicFieldMap), message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}

	skipVerifyPublicFieldMap := make(map[string]interface{})
	for _, name := range skipVerifyPublicFieldList {
		name := strings.Title(name) // 字段首字母大写
		skipVerifyPublicFieldMap[name] = nil
	}

	for fieldName, fieldMetadata := range publicFieldMap {
		// 是否跳过验证
		if _, ok := skipVerifyPublicFieldMap[fieldName]; ok {
			log.Debugf(module, "skip verify public Field [%s],message=[%v]", fieldName, message)
			continue
		}

		field := publicInfo.FieldByName(fieldName)
		if !field.IsValid() {
			desc := fmt.Sprintf("not found publicField[%s] in message=[%v]", fieldName, message)
			log.Warn(module, desc)
			return BadRequestErrorWithMessage(module, desc, message)
		}
		fieldType := GetFieldType(fieldMetadata.FieldType)

		//  字符串类型
		if fieldType == String {
			val := field.String()
			// 不允许为空
			if !fieldMetadata.FieldAllowNullOrEmpty {
				if val == "" {
					desc := fmt.Sprintf("field [%s] is empty while fieldAllowNullOrEmpty[%v], message=[%v]", fieldName, fieldMetadata.FieldAllowNullOrEmpty, message)
					log.Warn(module, desc)
					return BadRequestErrorWithMessage(module, desc, message)
				}
			}

			// 验证长度是否满足要求
			if len(val) > (int)(fieldMetadata.FieldMaxLength) {
				desc := fmt.Sprintf("field [%s]'s length [%d] larger than fieldMaxLength[%d], message=[%v]", fieldName, len(val), fieldMetadata.FieldMaxLength, message)
				log.Warn(module, desc)
				return BadRequestErrorWithMessage(module, desc, message)
			}

		} else {
			// 整数，布尔值不需要验证是否为空以及最大长度限制
			// 不需要验证
		}

	}
	return shim.Success(nil)
}

func validateCryptoMessage(module bsmodule.Module, cryptoMessage reflect.Value, message proto.Message, metadata *datatable_metadata.DataTableMetadata, ) pb.Response {
	// 判断是否为空值
	//if cryptoMessage.Interface() == (*transport.CryptoMessageEx)(nil) {
	if cryptoMessage.Pointer() == 0 {
		desc := fmt.Sprintf("cryptoMesage is null, message=[%v]", message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}

	cryptoMessageValue, ok := cryptoMessage.Interface().(*base.CryptoMessageEx)
	if !ok {
		desc := fmt.Sprintf("invalid cryptoMessage, type must be [*transport.CryptoMessageEx], message=[%v]", message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}

	// 平台id不能为空
	if cryptoMessageValue.PlatformId == "" {
		desc := fmt.Sprintf("cryptoMessage's platformId is empty, message=[%v]", message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}
	// 平台id长度不能太长
	if len(cryptoMessageValue.PlatformId) > 100 {
		desc := fmt.Sprintf("cryptoMessage's platformId[%s]' length larger than 100, message=[%v]", cryptoMessageValue.PlatformId, message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}

	// TODO 验证平台id的有效性

	// TODO 加密前hash长度限制

	// hash方法长度限制
	if len(cryptoMessageValue.HashMethod) > 100 {
		desc := fmt.Sprintf("cryptoMessage's HashMethod[%s]'s length larger than 100, message=[%v]", cryptoMessageValue.HashMethod, message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}
	// 加密方法长度限制
	if len(cryptoMessageValue.EncryptMethod) > 100 {
		desc := fmt.Sprintf("cryptoMessage's EncryptMethod[%s]'s length larger than 100, message=[%v]", cryptoMessageValue.EncryptMethod, message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}

	// 加密方法长度限制
	if len(cryptoMessageValue.DataType) > 100 {
		desc := fmt.Sprintf("cryptoMessage's DataType[%s]'s length larger than 100, message=[%v]", cryptoMessageValue.DataType, message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}

	// 判断密文信息是否为空
	if len(cryptoMessageValue.MessageData) == 0 {
		desc := fmt.Sprintf("cryptoMessage's MessageData's length is zero, message=[%v]", message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}

	// 验证长度
	var privateFieldTotalMaxLength int64 = 0
	for _, privateFieldMetadata := range metadata.PrivateFieldList {
		//name := strings.Title(privateFieldMetadata.FieldName)
		// privateFieldMap[name] = privateFieldMetadata
		privateFieldTotalMaxLength += privateFieldMetadata.FieldMaxLength
	}

	// 10倍以内
	if len(cryptoMessageValue.MessageData) > 10*int(privateFieldTotalMaxLength) {
		desc := fmt.Sprintf("cryptoMessage's MessageData's length[%d] is larger than privateFieldTotalMaxLength[%d]*10, message=[%v]", len(cryptoMessageValue.MessageData), privateFieldTotalMaxLength, message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}


	// 验证加密信封
	// 信封个数不能为0
	if len(cryptoMessageValue.EnvelopInfo) == 0 {
		desc := fmt.Sprintf("cryptoemessage's Envelope count is zero, message=[%v]", message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	}

	// 验证加密信封
	for index, envelope := range cryptoMessageValue.EnvelopInfo {
		if len(envelope.EnvelopeData) == 0 {
			desc := fmt.Sprintf("envelope at index[%d]'s EnvelopeData is null or empty, message=[%v]", index, message)
			log.Warn(module, desc)
			return BadRequestErrorWithMessage(module, desc, message)
		}

		// 对EnvelopeInfo长度进行限制，本质上就是用非对称加密加密的随机生成的对称秘钥
		// 10k应该足够了
		if len(envelope.EnvelopeData) > 10240 {
			desc := fmt.Sprintf("envelope at index[%d]'s EnvelopeData's length[%d] is larger than 10240, message=[%v]", index, len(envelope.EnvelopeData), message)
			log.Warn(module, desc)
			return BadRequestErrorWithMessage(module, desc, message)
		}

		// 加密方法不能为空
		if len(envelope.EncryptMethod) == 0 {
			desc := fmt.Sprintf("envelope at index[%d]'s EnvelopeData's EncrythMethod is null or empty, message=[%v]", index, message)
			log.Warn(module, desc)
			return BadRequestErrorWithMessage(module, desc, message)
		}

		if len(envelope.EncryptMethod) > 100 {
			desc := fmt.Sprintf("envelope at index[%d]'s EncryptMethod's length[%d] is  large than 100, message=[%v]", index, len(envelope.EncryptMethod), message)
			log.Warn(module, desc)
			return BadRequestErrorWithMessage(module, desc, message)
		}

		// 不能同时为空
		if len(envelope.EncryptPublicKey) == 0 && len(envelope.EnvelopeIdentifier) == 0 {
			desc := fmt.Sprintf("envelope at index[%d] both EncryptPublicKey and EnvelopeIdentifier is null or empty, message=[%v]", index, message)
			log.Warn(module, desc)
			return BadRequestErrorWithMessage(module, desc, message)
		}


		// 10k应该足够了
		if len(envelope.EncryptPublicKey) > 10240 {
			desc := fmt.Sprintf("envelope at index[%d]'s EncryptPublicKey's length[%d] is  large than 10240, message=[%v]", index, len(envelope.EncryptPublicKey), message)
			log.Warn(module, desc)
			return BadRequestErrorWithMessage(module, desc, message)
		}

		if len(envelope.EnvelopeIdentifier) > 10240 {
			desc := fmt.Sprintf("envelope at index[%d]'s EnvelopeIdentifier's length[%d] is  large than 10240, message=[%v]", index, len(envelope.EnvelopeIdentifier), message)
			log.Warn(module, desc)
			return BadRequestErrorWithMessage(module, desc, message)
		}

	}


	return shim.Success(nil)
}

func ValidateProtoMessageByDataTableMetadata(module bsmodule.Module, message proto.Message, metadata *datatable_metadata.DataTableMetadata) pb.Response {
	return ValidateProtoMessageByDataTableMetadataWithSkipList(module, message, metadata, nil)
}

func ValidateProtoMessageByDataTableMetadataWithSkipList(module bsmodule.Module, message proto.Message, metadata *datatable_metadata.DataTableMetadata, skipVerifyPublicFieldList []string) pb.Response {
	v := reflect.ValueOf(message).Elem()
	cryptoMessage := v.FieldByNameFunc(isCryptoMessage)
	publicInfo := v.FieldByNameFunc(isPublicInfo)

	// 加密字段默认最大值

	//	var privateFieldMap = make(map[string]*transport.DataTableFieldMetadata)
	if publicInfo.IsValid() && cryptoMessage.IsValid() {

		resp := validatePublicInfo(module, publicInfo, message, metadata, skipVerifyPublicFieldList)
		if resp.Status != shim.OK {
			return resp
		}
		// 验证加密数据
		resp = validateCryptoMessage(module, cryptoMessage, message, metadata)
		return resp

	} else if !publicInfo.IsValid() && !cryptoMessage.IsValid() {
		// 只有公开信息
		publicInfo = reflect.ValueOf(message)
		resp := validatePublicInfo(module, publicInfo, message, metadata, skipVerifyPublicFieldList)
		return resp

	} else if publicInfo.IsValid() && !cryptoMessage.IsValid() {
		// 只有公开字段不合理
		desc := fmt.Sprintf("invalid message format, message=[%v]", message)
		log.Warn(module, desc)
		return BadRequestErrorWithMessage(module, desc, message)
	} else {
		// 只有加密字段
		// 验证加密字段
		// 验证加密数据
		resp := validateCryptoMessage(module, cryptoMessage, message, metadata)
		return resp
	}

	return shim.Success(nil)
}

// 根据数据类型获取对应的元数据
func GetMetadataTableByProtoMessage(module bsmodule.Module, ctx context.IBidchainContext, channelName string, message proto.Message) (*datatable_metadata.DataTableMetadata, error) {
	name := GetProtoMessageName(message)
	chaincodeName := "datatable_metadata"
	funcName := "GetMetadataTableByDataType"
	resp := ctx.InvokeChaincode(chaincodeName, generateFuncParameters(funcName, name), channelName)
	if resp.Status != shim.OK {
		return nil, errors.New(resp.Message)
	}
	if resp.Payload == nil {
		return nil, nil
	}
	var ret datatable_metadata.DataTableMetadata
	err := jsonutils.ProtoUnmarshal(resp.Payload, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}