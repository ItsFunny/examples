package protocol

import (
	"bidchain/fabric/context"
	"bidchain/protocol/store/store_base"
	"bidchain/protocol/transport/base"
	"github.com/golang/protobuf/proto"
)

func SaveMessage(ctx context.IBidchainContext, response ICommand, key string, storeMessage proto.Message) {
	data, err := proto.Marshal(storeMessage)
	if err != nil {
		response.SetErrCode(500)
		response.SetErrDesc(err.Error())
		return
	}
//	log.Debugf(moduleName, "PutState begin, key=%s, value=%v", key, data)
	err = ctx.PutState(key, data)
	if err != nil {
		response.SetErrCode(500)
		response.SetErrDesc(err.Error())
		return
	}
//	log.Debugf(moduleName, "PutState end, key=%s, value=%v", key, data)
	return
}

func Transport2StoreEncryptedMessageEx(src *base.CryptoMessageEx) *store_base.StoreCryptoMessageEx {
	if src == nil {
		return nil
	}
	var dest store_base.StoreCryptoMessageEx
	dest.PlatformId = src.PlatformId
	dest.HashBeforeEncrypt = src.HashBeforeEncrypt
	dest.HashMethod = src.HashMethod
	dest.EncryptMethod = src.EncryptMethod
	dest.DataType = src.DataType
	dest.MessageData = src.MessageData
	envelopeInfoList := make([]*store_base.StoreEnvelopeInfo, 0)
	for _, transportEnvelopeInfo := range src.EnvelopInfo {
		storeEnvelopeInfo := &store_base.StoreEnvelopeInfo{
			EnvelopeData:       transportEnvelopeInfo.EnvelopeData,
			EncryptMethod:      transportEnvelopeInfo.EncryptMethod,
			EncryptPublicKey:   transportEnvelopeInfo.EncryptPublicKey,
			EnvelopeIdentifier: transportEnvelopeInfo.EnvelopeIdentifier,
			Extension: transportEnvelopeInfo.Extension,
			Description: transportEnvelopeInfo.Description,
		}
		envelopeInfoList = append(envelopeInfoList, storeEnvelopeInfo)
	}
	dest.EnvelopInfo = envelopeInfoList
	return &dest
}

func Store2TransportEncryptedMessageEx(src *store_base.StoreCryptoMessageEx) *base.CryptoMessageEx {
	if src == nil {
		return nil
	}
	var dest base.CryptoMessageEx
	dest.PlatformId = src.PlatformId
	dest.HashBeforeEncrypt = src.HashBeforeEncrypt
	dest.HashMethod = src.HashMethod
	dest.EncryptMethod = src.EncryptMethod
	dest.DataType = src.DataType
	dest.MessageData = src.MessageData
	envelopeInfoList := make([]*base.EnvelopeInfo, 0)
	for _, storeEnvelopeInfo := range src.EnvelopInfo {
		transportEnvelopeInfo := &base.EnvelopeInfo{
			EnvelopeData:       storeEnvelopeInfo.EnvelopeData,
			EncryptMethod:      storeEnvelopeInfo.EncryptMethod,
			EncryptPublicKey:   storeEnvelopeInfo.EncryptPublicKey,
			EnvelopeIdentifier: storeEnvelopeInfo.EnvelopeIdentifier,
			Extension: storeEnvelopeInfo.Extension,
			Description: storeEnvelopeInfo.Description,
		}
		envelopeInfoList = append(envelopeInfoList, transportEnvelopeInfo)
	}
	dest.EnvelopInfo = envelopeInfoList
	return &dest
}