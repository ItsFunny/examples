/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/28 17:08
# @File : catalog.go
# @Description :
# @Attention :
*/
package bo

// 目录权限判断
type CatalogAuthenticationReq struct {
	CatalogOwnerPlatformId string
	Authentication         int
}

type CatalogAuthenticationResp struct {
	Allowed bool
}

/*
函数名为请求的首字母小写，去掉Req
 */
// 获取目录的创建者
type GetCatalogCreatorReq struct {
	CatalogId string
}

// 返回对应的目录创建者的平台d
type GetCatalogCreatorResp struct {
	PlatformId string
}

// 判断对应平台是否包含指定平台的子目录
type HasCatalogDescendantReq struct {
	PlatformId string
	CatalogId string
	// 后代平台id (子、孙子等)
	DesendatantPlatformId string
}


type HasCatalogDescendantResp struct {
	Ok bool
}

// 取消继承关系表的一行关系
type RevokeInheritReq struct {
	PlatformId string
	CatalogId string
	// 子目录所在平台
	ChildPlatformId string
}

type RevokeInheritResp struct {

}

// 一个平台是否对某个目录有指定的权限
type HasCatalogPermissionReq struct {
	PlatformId string
	CatalogId string
	Permission int32
}

type HasCatalogPermissionResp struct {
	Ok bool
}

/*
后面可能会增加的
getCatalogIdList(platformId) // 获取所有的目录
getCatalogChildren(platformId, catalogId) // 返回子目录的平台id和catalogid []{  // 按照层级  儿子、孙子 ...
getCatalogAncenstor(platformId, catalogId) // 返回父目录的平台id和catalogid // 父亲、爷爷...
 */
