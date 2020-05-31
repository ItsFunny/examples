/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-03-15 13:24 
# @File : copyright.go
# @Description : 
# @Attention : 
*/
package args

import (
	"encoding/json"
	error2 "examples/blockchain/config/common/error"
	"examples/blockchain/config/common/models/copyrightcc/models"
)

var (
	COPYRIGHT_UPLOAD_CONVERTER = func(args []string) (interface{}, error2.IVlinkError) {
		var res models.ItemUpload2ChainReqBO
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}
	COPYRIGHT_GET_COPYRIGHTINFO = func(args []string) (interface{}, error2.IVlinkError) {
		var res models.ItemGetCopyrightInfoReq
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}
)
