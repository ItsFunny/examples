package command

import (
	"bidchain/http_framework/protocol"
	"bidchain/protocol/transport/catalog"
	"github.com/golang/protobuf/proto"
	"reflect"
)

type AddMainDataResponse struct {
	*protocol.AbstractCommand
	params catalog.AddMainDataResponseParameters
}

func (request *AddMainDataResponse) GetCoupleType() reflect.Type {
	return reflect.TypeOf(AddMainDataRequest{})
}

func (request *AddMainDataResponse) GetURI() string {
	return "/catalog/addMainData"
}

func (request *AddMainDataResponse) GetFuncName() string {
	return "addMainData"
}

func (request *AddMainDataResponse) GetParameters() proto.Message {
	return &request.params
}
