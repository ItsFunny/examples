/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-03-08 13:00 
# @File : tx_base_test.go.go
# @Description : 
# @Attention : 
*/
package base

import (
	"fmt"
	"testing"
)

// func TestGetRegularInfo(t *testing.T) {
// 	base.NewNeedRecordTransBaseDescription(base.USERCC_TX_BASE_USER_UPLOAD, "用户上链")
// }

func TestGetRegularInfo(t *testing.T) {
	description := NewNeedRecordTransBaseDescription(USERCOIN_TX_BASE_UPLOAD_USER_WITH_INIT_COIN, "用户上链")
	v2s := description.TransBaseType
	fmt.Println(v2s.String())
}
