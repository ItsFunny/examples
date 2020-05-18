/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-07 09:56 
# @File : user.go
# @Description : 
# @Attention : 
*/
package args

import (
	"encoding/json"
	error2 "vlink.com/v2/vlink-common/error"
	"vlink.com/v2/vlink-common/models"
)

var (

	USERCOIN_UPLOAD_USER_CONVETER= func(args []string) (interface{}, error2.IVlinkError) {
		var res models.BCUploadUserAndCoinReq
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}
)
