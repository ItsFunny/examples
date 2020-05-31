/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-18 15:18 
# @File : user.go
# @Description : 
# @Attention : 
*/
package args

import (
	"encoding/json"
	error2 "examples/blockchain/config/common/error"
	"examples/blockchain/config/common/models"
)

var (
	USER_UPLOAD_CONVERTER = func(args []string) (interface{}, error2.IVlinkError) {
		var res models.BCUploadUserReq
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}

	USER_UPDATEINFO_CONVTER=func(args []string) (interface{}, error2.IVlinkError){
		var res models.BCUpdateUserReq
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}

	USER_GETUSERINFO_CONVTER=func(args []string) (interface{}, error2.IVlinkError){
		var res models.BCGetUserInfoReq
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}

	USER_GETUSERCOIN_CONVTER=func(args []string) (interface{}, error2.IVlinkError){
		var res models.BCUserGetCoinAmoutReq
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}

	USER_GETUSERWALLET_CONVTER=func(args []string) (interface{}, error2.IVlinkError){
		var res models.BCWalletInfoReq
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}

	USER_UPDATEPWD_CONVTER=func(args []string) (interface{}, error2.IVlinkError){
		var res models.BCUserUpdatePwdReqBO
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}

)
