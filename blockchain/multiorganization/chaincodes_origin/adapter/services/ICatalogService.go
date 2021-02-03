/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/27 14:55
# @File : ICatalogService.go
# @Description :
# @Attention :
*/
package services

import (
	"bidchain/base/services"
	"bidchain/chaincode/catalog/adapter/models"
	"bidchain/fabric/bserror"
	"bidchain/fabric/log"
	"bidchain/micro/catalog/bo"
	"bidchain/micro/common/dto"
)

type ICatalogService interface {
	log.Logger
	services.IBaseService
	AddCatalogInheritDetail(models.CatalogInheritanceUploadReq) (models.CatalogInheritanceUploadResp, bserror.IBSError)
	AddOrUpdateCatalog(req models.CatalogUploadReq) (models.CatalogUploadResp, bserror.IBSError)
	// 主数据上链
	AddMainData(models.CatalogMainDataUploadReq) (models.CatalogMainDataUploadResp, bserror.IBSError)

	GetCatalogInfoById(id bo.GetCatalogInfoByIdReq) (dto.ResultDTO, bserror.IBSError)

	RevokeCatalogRelation(req bo.RevokeCatalogRelationReq) (dto.ResultDTO, bserror.IBSError)
	// 判断是否是后代目录
	CheckIsDesendatantCatalog(req bo.CheckIsDesendatantCatalogReq) (dto.ResultDTO, bserror.IBSError)
}


