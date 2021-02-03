package command

import (
	"bidchain/http_framework/protocol"
	"bidchain/protocol/transport/catalog"
	"github.com/golang/protobuf/proto"
	"reflect"
)

type AddOrUpdateCatalogResponse struct {
	*protocol.AbstractCommand
	params catalog.AddOrUpdateCatalogResponseParameters
}

func (request *AddOrUpdateCatalogResponse) GetCoupleType() reflect.Type {
	return reflect.TypeOf(AddOrUpdateCatalogRequest{})
}

func (request *AddOrUpdateCatalogResponse) GetURI() string {
	return "/catalog/addOrUpdateCatalog"
}

func (request *AddOrUpdateCatalogResponse) GetFuncName() string {
	return "addOrUpdateCatalog"
}

func (request *AddOrUpdateCatalogResponse) GetParameters() proto.Message {
	return &request.params
}
