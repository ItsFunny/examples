/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-10 13:31 
# @File : qr.go
# @Description : 
*/
package utils

import "github.com/skip2/go-qrcode"

func GenerateQR(url string) ([]byte, error) {
	return qrcode.Encode(url, qrcode.Medium, 256)
}
