package chaincode

import (
	"bidchain/fabric/context"
	"bidchain/protocol/store/store_base"
	"bidchain/protocol/transport/base"
	"fmt"
	"github.com/pkg/errors"
)

func Transport2StoreAttachmentInfo(src *base.AttachmentInfo) *store_base.StoreAttachmentInfo {
	if src == nil {
		return nil
	}
	var dest store_base.StoreAttachmentInfo
	dest.IpfsFileHashId = src.IpfsFileHashId
	dest.FileHash = src.FileHash
	dest.HashMethod = src.HashMethod
	dest.ChannelName = src.ChannelName
	return &dest
}

func Store2TransportAttachmentInfo(src *store_base.StoreAttachmentInfo) *base.AttachmentInfo {
	if src == nil {
		return nil
	}
	var dest base.AttachmentInfo
	dest.IpfsFileHashId = src.IpfsFileHashId
	dest.FileHash = src.FileHash
	dest.HashMethod = src.HashMethod
	dest.ChannelName = src.ChannelName
	return &dest
}

func IsValidAttachmentInfo(ctx context.IBidchainContext, src *base.AttachmentInfo) error {
	if src == nil {
		return errors.New("attachmentInfo is null")
	}
	if src.IpfsFileHashId == "" {
		return errors.New("IpfsFileHashId is empty")
	}
	if src.FileHash == "" {
		return errors.New("FileHash is empty")
	}
	if src.HashMethod == "" {
		return errors.New("HashMethod is empty")
	}
	//if src.ChannelName == "" {
	//	return errors.New("ChannelName is empty")
	//}
	channelName := src.ChannelName
	if channelName == "" {
		channelName = "ebidsun-files"
	}

	//// 判断对应的ipfs 文件是否存在
	ipfsHashList := []string{src.IpfsFileHashId}
	existsList := IsIpfshFileHashInfoExistByChaincodeQueryInternal(ctx, channelName, ipfsHashList)
	if existsList[0] == false {
		return fmt.Errorf("ipfsHashId[%s] does not exist in channel[%s]", src.IpfsFileHashId, src.ChannelName)
	}

	return nil
}
