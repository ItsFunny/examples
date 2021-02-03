package chaincode

import (
	"bidchain/protocol/store/store_base"
	"bidchain/protocol/transport/base"
	"fmt"
	"github.com/pkg/errors"
)

var (
	nilCryptoMessage   = errors.New("CryptoMessage is null")
	emptyPlatformId    = errors.New("platformId is empty")
	emptyEnCryptMethod = errors.New("encryptMethod is empty")
	emptyMessageData   = errors.New("messageData is empty")
	emptyDataType      = errors.New("dataType is empty")
)

func IsValidCryptoMessageEx(message *base.CryptoMessageEx, platformId string) error {
	var err error
	if message == nil {
		return nilCryptoMessage
	}
	if message.PlatformId == "" {
		return emptyPlatformId
	}
	if len(message.PlatformId) > 1000 {
		return fmt.Errorf("platformId' length[%d] too large", len(message.PlatformId))
	}
	if message.PlatformId != platformId {
		return fmt.Errorf("CryptoMessage's platformId[%s] different with platformId[%s]", message.PlatformId, platformId)
	}

	if len(message.HashBeforeEncrypt) > 10000000 {
		return fmt.Errorf("CryptoMessage's HashBeforeEncrypt's lenght[%d] too large", len(message.HashBeforeEncrypt))
	}

	if len(message.HashMethod) > 100 {
		return fmt.Errorf("CryptoMessage's hashMethod's length[%d] too large", len(message.HashMethod))
	}

	if message.EncryptMethod == "" {
		return emptyEnCryptMethod
	}
	if len(message.EncryptMethod) > 100 {
		return fmt.Errorf("CryptoMessage's EncryptMethod's length[%d] too large", len(message.EncryptMethod))
	}

	if message.DataType == "" {
		return emptyDataType
	}
	if len(message.DataType) > 100 {
		return fmt.Errorf("CryptoMessage's dataType's length[%d] too large", len(message.DataType))
	}

	if len(message.MessageData) == 0 {
		return emptyMessageData
	}
	if len(message.MessageData) > 10000000 {
		return fmt.Errorf("messageData's length[%d] too large", len(message.MessageData))
	}

	// 验证信封
	if len(message.EnvelopInfo) == 0 {
		return fmt.Errorf("no EnvelopeInfo")
	}

	for index, info := range message.EnvelopInfo {
		if err = isValidEnvelopeInfo(info); err != nil {
			err = errors.Wrapf(err, "invalid envelopeInfo at index[%d]", index)
			return err
		}
	}

	if len(message.Extension) >= 10000000 {
		return fmt.Errorf("extension's lenght[%d] too large", len(message.Extension))
	}

	return nil
}

var (
	nilEnvelopeInfo       = errors.New("EnvelopeInfo is null")
	emptyEnvelopeData     = errors.New("envelopeData is empty")
	emptyEncryptMethod    = errors.New("encryptMethod is empty")
	emptyEncryptPublicKey = errors.New("encryptPublicKey is empty")
)

func isValidEnvelopeInfo(env *base.EnvelopeInfo) error {
	if env == nil {
		return nilEnvelopeInfo
	}
	if len(env.EnvelopeData) == 0 {
		return emptyEnvelopeData
	}

	if len(env.EnvelopeData) > 10000 {
		return fmt.Errorf("EnvelopeData's length[%d] too large", len(env.EnvelopeData))
	}

	if env.EncryptMethod == "" {
		return emptyEncryptMethod
	}

	if len(env.EncryptMethod) > 100 {
		return fmt.Errorf("encryptMethod's length[%d] too large", len(env.EncryptMethod))
	}

	if env.EncryptPublicKey == "" {
		return emptyEncryptPublicKey
	}

	if len(env.EncryptPublicKey) > 10000 {
		return fmt.Errorf("encryptPublicKey's length[%d] too large", len(env.EncryptPublicKey))
	}

	if len(env.EnvelopeIdentifier) > 1000 {
		return fmt.Errorf("envelopeIdentifier's length[%d] too large", len(env.EnvelopeIdentifier))
	}
	if len(env.Extension) > 10000000 {
		return fmt.Errorf("extension's length[%d] too large", len(env.Extension))
	}
	if len(env.Description) > 1000 {
		return fmt.Errorf("description's length[%d] too large", len(env.Description))
	}
	return nil
}

func Transport2StoreEncryptedMessageEx(src *base.CryptoMessageEx) *store_base.StoreCryptoMessageEx {
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
			Extension:          transportEnvelopeInfo.Extension,
			Description:        transportEnvelopeInfo.Description,
		}
		envelopeInfoList = append(envelopeInfoList, storeEnvelopeInfo)
	}
	dest.EnvelopInfo = envelopeInfoList
	dest.Extension = src.Extension
	return &dest
}

func Store2TransportEncryptedMessageEx(src *store_base.StoreCryptoMessageEx) *base.CryptoMessageEx {
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
			Extension:          storeEnvelopeInfo.Extension,
			Description:        storeEnvelopeInfo.Description,
		}
		envelopeInfoList = append(envelopeInfoList, transportEnvelopeInfo)
	}
	dest.EnvelopInfo = envelopeInfoList
	dest.Extension = src.Extension
	return &dest
}
