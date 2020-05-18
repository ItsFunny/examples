/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-18 15:07 
# @File : method.go
# @Description : 
# @Attention : 
*/
package method

import "vlink.com/v2/vlink-common/base/fabric"

const (
	USERCC_METHOD_UPLOAD_USER_TO_CHAIN = base.MethodName("UploadUserToChain")
	USERCC_METHOD_UPDATE_USERINFO      = base.MethodName("UpdateUserInfo")
	USERCC_METHOD_GET_USERINFO         = base.MethodName("GetUserInfo")

	USERCC_METHOD_GET_USERCOIN = base.MethodName("GetUserCoin")

	USERCC_METHOD_GET_USERWALLET  = base.MethodName("GetUserWallet")
	USERCC_METHOD_UPDATE_USER_PWD = base.MethodName("UpdateUserPwd")
)
