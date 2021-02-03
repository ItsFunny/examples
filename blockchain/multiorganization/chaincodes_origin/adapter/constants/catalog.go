/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/28 16:06
# @File : relation.go
# @Description :
# @Attention :
*/
package constants

import "bidchain/base/config"

var (

	// 目录继承明细前缀 : 数据是这个目录的所有父目录和版本号 prefix+目录id  ,内部为嵌套数据, key(版本号)->多条记录
	// OT_RELATION_PARENT_INFO_LIST config.ObjectType = "OT_RELATION_PARENT_INFO_LIST"
	// 目录继承明细前缀: 数据是这个目录的所有子目录和版本号 prefix+目录id ,内部为嵌套数据,key(版本号)->多条子目录信息
	OT_RELEATION_CHILD_INFO_LIST config.ObjectType = "OT_RELEATION_CHILD_INFO_LIST"

	// 目录上链前缀  prefix+id+version
	OT_CATALOG_BASICINFO config.ObjectType = "OT_CATALOG_BASICINFO"
	// 目录上链-每个目录都会维护一个自身的版本号列表 prefix+id
	// OT_CATALOG_VERSION config.ObjectType="OT_CATALOG_VERSION"
	// 目录上链-编码集,prefix+id+version
	OT_CATALOG_DATATYPE config.ObjectType = "OT_CATALOG_DATATYPE"
	// 目录上链-标签, prefix+id+version
	OT_CATALOG_TAG config.ObjectType = "OT_CATALOG_TAG"
	// 目录上链-表关系 prefix+id+version
	OT_CATALOG_TABLECONSTRAINT config.ObjectType = "OT_CATALOG_TABLECONSTRAINT"

	// 主数据 : prefix+ 目录ID
	OT_MAINDATA config.ObjectType = "OT_MAINDATA"
)

var (
	// 权限跨链查询

	// 目录权限跨链查询
	OC_CATALOG_AUTH config.OverChainKey = "OC_CATALOG_AUTH"

	//
)

func init() {
	config.AddKey([]config.ObjectType{
		OT_RELEATION_CHILD_INFO_LIST, OT_RELEATION_CHILD_INFO_LIST, OT_CATALOG_BASICINFO,
		OT_CATALOG_DATATYPE, OT_CATALOG_TAG, OT_CATALOG_TABLECONSTRAINT,
		OT_MAINDATA,
	}, []config.KeyGenerator{
		config.COMMON_STRING_OR_STRING_ARRAY_KEY_GENERATOR, config.COMMON_STRING_OR_STRING_ARRAY_KEY_GENERATOR, config.COMMON_STRING_OR_STRING_ARRAY_KEY_GENERATOR,
		config.COMMON_STRING_OR_STRING_ARRAY_KEY_GENERATOR, config.COMMON_STRING_OR_STRING_ARRAY_KEY_GENERATOR, config.COMMON_STRING_OR_STRING_ARRAY_KEY_GENERATOR,
		config.COMMON_STRING_OR_STRING_ARRAY_KEY_GENERATOR,
	})
	config.AddOverChain([]config.OverChainKey{OC_CATALOG_AUTH}, []string{"1"}, []string{"1"})
}
