/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/2/1 16:15
# @File : catalog.go
# @Description :
# @Attention :
*/
package bo

type GetCatalogInfoByIdReq struct {
	CatalogId string `json:"catalogId"`
}

type GetCatalogInfoByIdResp struct {
	CatalogId              string `json:"catalogId"`
	CatalogName            string `json:"catalogName"`
	UploadVersion          int32 `json:"uploadVersion"`
	ShowVersionId          string `json:"showVersionId"`
	ShowVersion            string `json:"showVersion"`
	CatalogOwnerPlatformId string `json:"catalogOwnerPlatformId"`
	StateEnumId            string `json:"stateEnumId"`
	PublishTime            int64  `json:"publishTime"`
	CreateDraftTime        int64  `json:"createDraftTime"`
	StartPublicityTime     int64  `json:"startPublicityTime"`
	EndPublicityTime       int64  `json:"endPublicityTime"`
}

// 删除 父目录的时候,切断与子目录的关系
// 触发时机: 取消授权
type RevokeCatalogRelationReq struct {
	ParentCatalogId string `json:"parentCatalogId"`
	ChildCatalogId  string `json:"childCatalogId"`
}

type RevokeCatalogRelationResp struct {
}

// d
type CheckIsDesendatantCatalogReq struct {
	ParentCatalogId string `json:"parentCatalogId"`
	ChildCatalogId  string `json:"childCatalogId"`
}

type CheckIsChildCatalogResp struct {
	Desendatant bool `json:"desendatant"`
}
