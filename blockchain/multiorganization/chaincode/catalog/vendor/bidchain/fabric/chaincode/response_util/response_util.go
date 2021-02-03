package response_util

import (
	"bidchain/base/base64utils"
	"bidchain/http_framework/protocol"
	"bidchain/protocol/transport"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type errCode int

var (
	bad_request           errCode = 400
	internal_server_error errCode = 500
)

func ErrorCommandResponse(errCode int64, errDesc string, data interface{}) pb.Response {
	var body []byte
	if data != nil {
		switch data.(type) {
		case string:
			body = []byte(data.(string))
		case []byte:
			body = data.([]byte)
		default:
			msg := fmt.Sprintf("invalid data type")
			panic(msg)
		}
	}

	packet := transport.Packet{
		Header: &transport.Packet_Header{
			ErrorCode: errCode,
			ErrorDesc: errDesc,
		},
		Body: body,
	}
	marshalBytes, _ := proto.Marshal(&packet)
	s := base64utils.Base64Encode(marshalBytes)
	s = protocol.ERROR_RESPONSE_WITH_HTTP_COMMAND_PREFIX + s
	return shim.Error(s)
}

func BadRequestCommandResponse(errDesc string) pb.Response {
	return ErrorCommandResponse(int64(bad_request), errDesc, nil)
}
func BadRequestCommandResponseWithData(errDesc string, data interface{}) pb.Response {
	return ErrorCommandResponse(int64(bad_request), errDesc, data)
}

func InternalServerErrorCommandResponse(errDesc string) pb.Response {
	return ErrorCommandResponse(int64(internal_server_error), errDesc, nil)
}

func InternalServerErrorCommandResponseWithData(errDesc string, data interface{}) pb.Response {
	return ErrorCommandResponse(int64(internal_server_error), errDesc, data)
}
