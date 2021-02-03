package command

import (
	"bidchain/http_framework/protocol"
	"bidchain/protocol/transport/catalog"
	"github.com/golang/protobuf/proto"
	"reflect"
)

type AddMainDataRequest struct {
	*protocol.AbstractCommand
	params catalog.AddMainDataRequestParameters
}

func (request *AddMainDataRequest) GetCoupleType() reflect.Type {
	return reflect.TypeOf(AddMainDataResponse{})
}

func (request *AddMainDataRequest) GetURI() string {
	return "/catalog/addMainData"
}

func (request *AddMainDataRequest) GetFuncName() string {
	return "addMainData"
}

func (request *AddMainDataRequest) GetParameters() proto.Message {
	return &request.params
}

func (request *AddMainDataRequest) GetReq() []*catalog.Data {
    return request.params.Req
}

func (request *AddMainDataRequest) SetReq(req []*catalog.Data ) {
    request.params.Req = req
}

