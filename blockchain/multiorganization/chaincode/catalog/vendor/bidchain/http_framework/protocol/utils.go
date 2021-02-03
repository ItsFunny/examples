package protocol

import (
	"bidchain/protocol/transport"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func GetRequestType(header *transport.Packet_Header) RequestType {
	if header == nil {
		return RequestType_QUERY
	}
	switch header.RequestType {
	case transport.RequestType_QUERY:
		return RequestType_QUERY
	default:
		return RequestType_INVOKE
	}
}

func DeserializeParams(cmd ICommand, data []byte) error {
	// 反序列话成packet
	var packet transport.Packet
	err := proto.Unmarshal(data, &packet)
	if err != nil {
		msg := fmt.Sprintf("invalid packet data=%v", data)
		return errors.WithMessage(err, msg)
	}

	// 获取请求类型
	reqType := GetRequestType(packet.Header)
	cmd.SetRequestType(reqType)
	if packet.Body != nil {
		// 报文反序列化
		err = proto.Unmarshal(packet.Body, cmd.GetParameters())
		if err != nil {
			return err
		}
	}
	return nil
}

func SerializeParams(cmd ICustomCommand) ([]byte, error) {
	data, err := proto.Marshal(cmd.GetParameters())
	return data, err
}
