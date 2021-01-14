/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-05-21 08:29 
# @File : main.go
# @Description : 
# @Attention : 
*/
package main

import (
	"flag"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"myLibrary/go-library/go/utils"
)

var (
	templatePath = flag.String("template", "/Users/joker/Desktop/fabric-ca1.4.htm", "")
)

func main() {

	flag.Parse()

	htmlReq := utils.NewStaticDynamicHtmlReq(*templatePath, "/usr/local/", "newName", map[string]interface{}{"name": "joker"})
	resp, e := utils.StaticDynamicHtml(htmlReq)
	if nil != e {
		panic(e)
	}

	pdfSourceReq := utils.NewLocalHtml2PdfPureSourceReq(resp.StaticHtmlStorePath, "/usr/local", "pdfNewName")
	pdfSourceReq.PageDecorator = func(options wkhtmltopdf.PageOptions) wkhtmltopdf.PageOptions {
		options.Zoom.Set(0.95)
		return options
	}
	_, e = utils.Html2PdfBySource(pdfSourceReq)
	if nil != e {
		panic(e)
	}
}
