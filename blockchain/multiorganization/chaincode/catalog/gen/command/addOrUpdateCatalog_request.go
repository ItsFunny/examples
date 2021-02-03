package command

import (
	"bidchain/http_framework/protocol"
	"bidchain/protocol/transport/catalog"
	"github.com/golang/protobuf/proto"
	"reflect"
)

type AddOrUpdateCatalogRequest struct {
	*protocol.AbstractCommand
	params catalog.AddOrUpdateCatalogRequestParameters
}

func (request *AddOrUpdateCatalogRequest) GetCoupleType() reflect.Type {
	return reflect.TypeOf(AddOrUpdateCatalogResponse{})
}

func (request *AddOrUpdateCatalogRequest) GetURI() string {
	return "/catalog/addOrUpdateCatalog"
}

func (request *AddOrUpdateCatalogRequest) GetFuncName() string {
	return "addOrUpdateCatalog"
}

func (request *AddOrUpdateCatalogRequest) GetParameters() proto.Message {
	return &request.params
}

func (request *AddOrUpdateCatalogRequest) GetReq() []*catalog.Catalog {
    return request.params.Req
}

func (request *AddOrUpdateCatalogRequest) SetReq(req []*catalog.Catalog ) {
    request.params.Req = req
}

