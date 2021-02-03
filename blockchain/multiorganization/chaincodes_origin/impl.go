/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/2/1 14:15
# @File : impl.go
# @Description :
# @Attention :
*/
package main

//
// func argumentError(iCommand protocol.ICommand, msg string) {
// 	errorCommand(iCommand, 400, msg)
// }
// func errorCommand(iCommand protocol.ICommand, code int64, msg string) {
// 	iCommand.SetErrCode(code)
// 	iCommand.SetErrDesc(msg)
// }
//
// // func (cc *CatalogContract) AddCatalogInheritDetail(request *command.AddCatalogInheritDetailRequest, response *command.AddCatalogInheritDetailResponse) {
// // 	list := request.GetReq()
// // 	if len(list) == 0 {
// // 		response.SetErrDesc("参数为空")
// // 		response.SetErrCode(400)
// // 		return
// // 	}
// // 	reqIdList := make([]string, 0)
// // 	for _, node := range list {
// // 		reqIdList = append(reqIdList, node.InheritDetailId)
// // 	}
// // 	serviceImpl := impl.NewCatalogServiceImpl(&cc.BasisContract, strings.Join(reqIdList, ","))
// // 	serviceImpl.BeforeStart("AddCatalogInheritDetail")
// // 	defer serviceImpl.AfterEnd()
// // 	for _, node := range list {
// // 		req := models.CatalogInheritanceUploadReq{Detail: node}
// // 		if e := req.Validation(); nil != e {
// // 			marshal, _ := json.Marshal(node)
// // 			serviceImpl.Error("参数校验失败:" + e.Error() + ",请求数据为:" + string(marshal))
// // 			argumentError(response, e.Error())
// // 			return
// // 		}
// // 		_, ibsError := serviceImpl.AddCatalogInheritDetail(req)
// // 		if nil != ibsError {
// // 			bytes, e := json.Marshal(node)
// // 			msg := "上链数据失败,"
// // 			if nil == e {
// // 				msg += "失败数据为:" + string(bytes)
// // 			}
// // 			msg += ",错误为:" + ibsError.Error()
// // 			serviceImpl.Error(msg)
// // 			utils.ReturnWithError(response, ibsError)
// // 			return
// // 		}
// // 	}
// // }
//
// func (cc *CatalogContract) AddOrUpdateCatalog(request *command.AddOrUpdateCatalogRequest, response *command.AddOrUpdateCatalogResponse) {
// 	catalogList := request.GetReq()
// 	reqIdList := make([]string, 0)
// 	for _, node := range catalogList {
// 		reqIdList = append(reqIdList, node.CatalogBasicInfo.CatalogId)
// 	}
//
// 	serviceImpl := impl.NewCatalogServiceImpl(&cc.BasisContract, strings.Join(reqIdList, ","))
// 	serviceImpl.BeforeStart("AddOrUpdateCatalog")
// 	defer serviceImpl.AfterEnd()
//
// 	for _, node := range catalogList {
// 		req := models.CatalogUploadReq{Req: node}
// 		if e := req.Validation(); nil != e {
// 			marshal, _ := json.Marshal(node)
// 			serviceImpl.Error("参数校验失败:" + e.Error() + ",请求数据为:" + string(marshal))
// 			argumentError(response, e.Error())
// 			return
// 		}
//
// 		_, ibsError := serviceImpl.AddOrUpdateCatalog(req)
// 		if nil != ibsError {
// 			bytes, _ := json.Marshal(node)
// 			msg := "上链数据失败,"
// 			msg += "失败数据为:" + string(bytes)
// 			msg += ",错误为:" + ibsError.Error()
// 			serviceImpl.Error(msg)
// 			utils.ReturnWithError(response, ibsError)
// 			return
// 		}
// 	}
// }
//
// func (cc *CatalogContract) AddMainData(request *command.AddMainDataRequest, response *command.AddMainDataResponse) {
// 	catalogList := request.GetReq()
// 	reqIdList := make([]string, 0)
// 	for _, node := range catalogList {
// 		reqIdList = append(reqIdList, node.DataID)
// 	}
//
// 	serviceImpl := impl.NewCatalogServiceImpl(&cc.BasisContract, strings.Join(reqIdList, ","))
// 	serviceImpl.BeforeStart("AddMainData")
// 	defer serviceImpl.AfterEnd()
//
// 	for _, node := range catalogList {
// 		req := models.CatalogMainDataUploadReq{Req: node}
// 		if e := req.Validation(); nil != e {
// 			marshal, _ := json.Marshal(node)
// 			serviceImpl.Error("参数校验失败:" + e.Error() + ",请求数据为:" + string(marshal))
// 			argumentError(response, e.Error())
// 			return
// 		}
//
// 		_, ibsError := serviceImpl.AddMainData(req)
// 		if nil != ibsError {
// 			bytes, _ := json.Marshal(node)
// 			msg := "上链数据失败,"
// 			msg += "失败数据为:" + string(bytes)
// 			msg += ",错误为:" + ibsError.Error()
// 			serviceImpl.Error(msg)
// 			utils.ReturnWithError(response, ibsError)
// 			return
// 		}
// 	}
// }
//
//
// func (cc *CatalogContract) GetCatalogInfoById(ctx context.IBidchainContext, args []string) peer.Response {
// 	if len(args) == 0 {
// 		return shim.Error("参数错误,为空")
// 	}
// 	var req bo.GetCatalogInfoByIdReq
// 	if e := json.Unmarshal([]byte(args[0]), &req); nil != e {
// 		return shim.Error("反序列化失败:" + e.Error())
// 	}
//
// 	reqId := req.CatalogId
// 	serviceImpl := impl.NewCatalogServiceImpl(&cc.BasisContract, reqId)
// 	serviceImpl.BeforeStart("GetCatalogInfoById")
// 	dto, ibsError := serviceImpl.GetCatalogInfoById(req)
// 	if nil != ibsError {
// 		return shim.Error(ibsError.Error())
// 	}
//
// 	return dto.Result()
// }
//
// // CheckIsDesendatantCatalog(req bo.CheckIsDesendatantCatalogReq) (dto.ResultDTO, bserror.IBSError)
// func (cc *CatalogContract) CheckIsDesendatantCatalog(ctx context.IBidchainContext, args []string) peer.Response {
// 	if len(args) == 0 {
// 		return shim.Error("参数错误,为空")
// 	}
// 	var req bo.CheckIsDesendatantCatalogReq
// 	if e := json.Unmarshal([]byte(args[0]), req); nil != e {
// 		return shim.Error("反序列化失败:" + e.Error())
// 	}
//
// 	reqId := req.ParentCatalogId+","+req.ChildCatalogId
// 	serviceImpl := impl.NewCatalogServiceImpl(&cc.BasisContract, reqId)
// 	serviceImpl.BeforeStart("CheckIsDesendatantCatalog")
// 	dto, ibsError := serviceImpl.CheckIsDesendatantCatalog(req)
// 	if nil != ibsError {
// 		return shim.Error(ibsError.Error())
// 	}
//
// 	return dto.Result()
// }
//
//
// // RevokeCatalogRelation(req bo.RevokeCatalogRelationReq) (dto.ResultDTO, bserror.IBSError)
// func (cc *CatalogContract) RevokeCatalogRelation(ctx context.IBidchainContext, args []string) peer.Response {
// 	if len(args) == 0 {
// 		return shim.Error("参数错误,为空")
// 	}
// 	var req bo.RevokeCatalogRelationReq
// 	if e := json.Unmarshal([]byte(args[0]), req); nil != e {
// 		return shim.Error("反序列化失败:" + e.Error())
// 	}
//
// 	reqId := req.ParentCatalogId+","+req.ChildCatalogId
// 	serviceImpl := impl.NewCatalogServiceImpl(&cc.BasisContract, reqId)
// 	serviceImpl.BeforeStart("RevokeCatalogRelation")
// 	dto, ibsError := serviceImpl.RevokeCatalogRelation(req)
// 	if nil != ibsError {
// 		return shim.Error(ibsError.Error())
// 	}
//
// 	return dto.Result()
// }
