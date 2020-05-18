/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-05 10:02 
# @File : coin.go
# @Description : 
# @Attention : 
*/
package args

import (
	"encoding/json"
	"fmt"
	error2 "vlink.com/v2/vlink-common/error"
	"vlink.com/v2/vlink-common/models/coincc"
)

var (
	// 更新用户积分
	USER_COIN_UPDATE = func(args []string) (interface{}, error2.IVlinkError) {
		for _, a := range args {
			fmt.Println("USER_COIN_UPDATE的参数:", a)
		}
		var res coincc.BCUserCoinUpdateReq
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}

	// 查询用户积分
	USER_COIN_GET = func(args []string) (interface{}, error2.IVlinkError) {
		for _, a := range args {
			fmt.Println("USER_COIN_UPDATE的参数:", a)
		}
		var res coincc.BCUserCoinGetReq
		if e := json.Unmarshal([]byte(args[0]), &res); nil != e {
			return nil, error2.NewArguError(e, "参数反序列化失败")
		}
		return res, nil
	}
)
