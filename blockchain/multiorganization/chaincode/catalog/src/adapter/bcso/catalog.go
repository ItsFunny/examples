/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/28 16:15
# @File : catalog.go
# @Description :
# @Attention :
*/
package bcso

import (
	"bidchain/protocol/transport/catalog"
)

// key 为 RELATION前缀+父目录的ID+版本号 构成
type BCSCatalogInheriRelationInfo struct {
	*catalog.CatalogInheriRelationInfo
}


//////////////// 目录基本信息
type BCSCatalogBasicInfo struct {
	*catalog.CatalogBasicInfo
}


// 目录的编码集信息
type BCSCatalogDataTypeInfo struct {
	*catalog.CatalogDataTypeInfo
}

type BCSCatalogDataTypeInfoNode struct {
	*catalog.CatalogDataTypeInfoNode
}