/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/2/1 15:24
# @File : main_test.go
# @Description :
# @Attention :
*/
package main
//
// //
// import (
// 	"bidchain/chaincode/catalog/gen/command"
// 	"bidchain/fabric/chaincode_mock"
// 	"bidchain/fabric/log"
// 	"bidchain/http_framework/protocol"
// 	"bidchain/protocol/transport/catalog"
// 	"fmt"
// 	l4g "github.com/alecthomas/log4go"
// 	"github.com/google/uuid"
// 	"github.com/hyperledger/fabric/core/chaincode/shim"
// 	"testing"
// 	"time"
// )
//
// func mockCC() *CatalogContract {
// 	log.SetLevel(l4g.DEBUG)
// 	cc := new(CatalogContract)
// 	cc.Child = cc
// 	// 注册cmd
// 	for _, cmd := range cmdList {
// 		protocol.RegisterCommand(cmd)
// 	}
//
// 	stub := shim.NewMockStub("catalog", cc)
// 	cc.Init(stub)
// 	// ctx := cc.GetContext()
// 	return cc
// }
// func Test_aaa(t *testing.T) {
// 	a := struct {
// 		name string
// 	}{name: "joker"}
// 	datas := make([]interface{}, 0)
// 	datas = append(datas, "123", a)
//
// }
//
// // func TestMainLocal(t *testing.T) {
// // }
//
// func TestCatalogContract_AddCatalogInheritDetail(t *testing.T) {
// 	cc := mockCC()
// 	ctx := cc.GetContext()
// 	ctx.SetFunctionName("addCatalogInheritDetail")
// 	request := new(command.AddCatalogInheritDetailRequest)
// 	requestList := make([]*catalog.CatalogInheritDetail, 0)
// 	requestList = append(requestList, &catalog.CatalogInheritDetail{
// 		NoeId:               randomString(),
// 		InheritDetailId:     randomString(),
// 		CatalogID:           randomString(),
// 		ParentCatalogID:     randomString(),
// 		ParentUploadVersion: randomString(),
// 	})
// 	request.SetReq(requestList)
//
// 	iCommand, result := chaincode_mock.InvokeCommand(cc, ctx, request)
// 	fmt.Println(result)
// 	fmt.Println(iCommand)
//
// 	time.Sleep(time.Second * 10)
// }
//
// func randomString() string {
// 	return uuid.New().String()
// }
// func randBytes() []byte {
// 	return []byte(randomString())
// }
