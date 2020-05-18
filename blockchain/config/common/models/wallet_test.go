/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-07 11:01 
# @File : wallet_test.go.go
# @Description : 
# @Attention : 
*/
package models

import (
	"fmt"
	"testing"
	"vlink.com/v2/vlink-common/constants/usercc/wallet"
	"myLibrary/go-library/go/crypt"
)

func TestNewWalletTree(t *testing.T) {
	prevPwd := "123"
	cryptPrevpwd := encrypt.MD5EncryptByBytes(prevPwd)
	vlinkWallet := NewVlinkWallet(wallet.WALLET_CHANGE_IN, cryptPrevpwd)
	fmt.Println(vlinkWallet.GetAddress())
	nowpwd := "1234"
	nowcryptPrevpwd := encrypt.MD5EncryptByBytes(nowpwd)
	newVlinkWallet := NewVlinkWallet(wallet.WALLET_CHANGE_IN, nowcryptPrevpwd)
	fmt.Println(newVlinkWallet.GetAddress())
}
