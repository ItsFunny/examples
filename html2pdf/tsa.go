/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-05-16 15:41 
# @File : tsa.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"errors"
	"myLibrary/go-library/go/utils"
)

type VTsaPdfReq struct {
	TemplateHtmlPath        string
	StaticHtmlStoreBasePath string
	NewName                 string
	Data                    map[string]interface{}

	// pdf
	PdfStoreBasePath string
}

func NewVTsaPdfReq(templateHtmlPath, staticHtmlStoreBasePath, pdfStoreBasePath, newName string, data map[string]interface{}) VTsaPdfReq {
	r := VTsaPdfReq{}
	r.TemplateHtmlPath = templateHtmlPath
	r.StaticHtmlStoreBasePath = staticHtmlStoreBasePath
	r.PdfStoreBasePath = pdfStoreBasePath
	r.NewName = newName
	r.Data = data

	return r
}

type VTsaPdfResp struct {
	VTsaHtmlStorePath string
	VTsaPdfStorePath  string
}

type PdfTsaDataReq struct {
	ObjectName string `json:"objectName"`
	// 作者
	ObjectAuthor string `json:"objectAuthor"`
	// 创建时间
	CreateDate string `json:"createDate"`
	// 本地时间戳
	LocalTimeStamp    string `json:"localTimeStamp"`
	CopyrightHashCode string `json:"copyrightHashCode"`
	// 版本
	Version string `json:"version"`
	// dna
	Dna string `json:"dna"`

	// 签名
	Signature string `json:"signature"`
	// 公钥
	PubKey string `json:"pubKey"`
	// 区块编号, 既区块高度
	BlockNumber uint64 `json:"blockNumber"`
	// 交易地址,既数据保存在哪个block中
	TransactionBlockAddress string `json:"transactionBlockAddress"`
}

func BuildVTsaPdfData(req PdfTsaDataReq) map[string]interface{} {
	m := make(map[string]interface{})
	m["platform"] = "VLINK"
	m["objectName"] = req.ObjectName
	m["objectAuthor"] = req.ObjectAuthor
	m["createDate"] = req.CreateDate
	m["timeStamp"] = req.LocalTimeStamp
	m["hashCode"] = req.CopyrightHashCode
	m["version"] = req.Version
	m["dna"] = req.Dna

	return m
}

// 先生成静态html文件,再通过静态html文件生成pdf文件
func GenerateHtmlAndConv2TLocalVTsaPdf(req VTsaPdfReq) (VTsaPdfResp, error) {
	var (
		result VTsaPdfResp
	)
	htmlReq := utils.NewStaticDynamicHtmlReq(req.TemplateHtmlPath, req.StaticHtmlStoreBasePath, req.NewName, req.Data)
	resp, e := utils.StaticDynamicHtml(htmlReq)
	if nil != e {
		return result, errors.New("生成html失败:" + e.Error())
	}
	result.VTsaHtmlStorePath = resp.StaticHtmlStorePath

	pdfSourceReq := utils.NewLocalHtml2PdfPureSourceReq(resp.StaticHtmlStorePath, req.PdfStoreBasePath, req.NewName)
	pdfResp, e := utils.Html2PdfBySource(pdfSourceReq)
	if nil != e {
		return result, errors.New("html转化为pdf失败:" + e.Error())
	}
	result.VTsaPdfStorePath = pdfResp.PdfStorePath

	return result, nil
}
