/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/2/1 19:17
# @File : contract.go
# @Description :
# @Attention :
*/
package testdata

//
import (
	"bidchain/chaincode/catalog/contract"
	"bidchain/chaincode/catalog/gen/command"
	"bidchain/fabric/chaincode_mock"
	"bidchain/fabric/log"
	"bidchain/http_framework/protocol"
	"bidchain/protocol/transport/catalog"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
	"time"
)

func Test_contract(t *testing.T) {

}
var (
	cmdList = make([]protocol.ICommand, 0)
)

func init() {
	cmdList = append(cmdList, &command.AddOrUpdateCatalogRequest{})
	cmdList = append(cmdList, &command.GetCatalogInfoByIdRequest{})
	cmdList = append(cmdList, &command.AddMainDataRequest{})
	cmdList = append(cmdList, &command.AddCatalogInheritDetailRequest{})
}
func mockCC() *contract.CatalogContract {
	log.SetLevel(l4g.DEBUG)
	cc := new(contract.CatalogContract)
	cc.Child = cc
	// 注册cmd
	for _, cmd := range cmdList {
		protocol.RegisterCommand(cmd)
	}

	stub := shim.NewMockStub("catalog", cc)
	cc.Init(stub)
	// ctx := cc.GetContext()
	return cc
}
func Test_aaa(t *testing.T) {
	a := struct {
		name string
	}{name: "joker"}
	datas := make([]interface{}, 0)
	datas = append(datas, "123", a)

}

// func TestMainLocal(t *testing.T) {
// }

func TestCatalogContract_AddCatalogInheritDetail(t *testing.T) {
	cc := mockCC()
	ctx := cc.GetContext()
	ctx.SetFunctionName("addCatalogInheritDetail")
	request := new(command.AddCatalogInheritDetailRequest)
	requestList := make([]*catalog.CatalogInheritDetail, 0)
	requestList = append(requestList, &catalog.CatalogInheritDetail{
		NoeId:               randomString(),
		InheritDetailId:     randomString(),
		CatalogID:           randomString(),
		ParentCatalogID:     randomString(),
		ParentUploadVersion: randomString(),
	})
	request.SetReq(requestList)

	iCommand, result := chaincode_mock.InvokeCommand(cc, ctx, request)
	fmt.Println(result)
	fmt.Println(iCommand)

	time.Sleep(time.Second * 10)
}

func randomString() string {
	return uuid.New().String()
}
func randBytes() []byte {
	return []byte(randomString())
}
