/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-05-16 16:05 
# @File : tsa_test.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"fmt"
	"myLibrary/go-library/go/utils"
	"testing"
)

var (
	templatePath        = "/Users/joker/go/src/vlink.com/v2/vlink-file/resources/vtsa_pdf.html"
	staticStoreBasePath = "/Users/joker/go/src/vlink.com/v2/vlink-file/resources/"
	pdfStoreBasePath    = "/Users/joker/go/src/vlink.com/v2/vlink-file/resources/"
	newName             = "test"
	data                = map[string]interface{}{"test": "ddddd"}
)

func TestGenerateLocalVTsaPdf(t *testing.T) {

	req := utils.NewStaticDynamicHtmlReq(templatePath, staticStoreBasePath, newName, data)
	resp, e := utils.StaticDynamicHtml(req)
	if nil != e {
		panic(e)
	}
	fmt.Println(resp.StaticHtmlStorePath)
}

func TestGenerateHtmlAndConv2TLocalVTsaPdf(t *testing.T) {
	r := NewVTsaPdfReq(templatePath, staticStoreBasePath, pdfStoreBasePath, newName, data)
	resp, e := GenerateHtmlAndConv2TLocalVTsaPdf(r)
	if nil != e {
		panic(e)
	}
	fmt.Println(resp.VTsaPdfStorePath)
	fmt.Println(resp.VTsaHtmlStorePath)
}
