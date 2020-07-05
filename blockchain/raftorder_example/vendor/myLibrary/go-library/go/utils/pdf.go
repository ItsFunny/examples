/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-05-12 11:57 
# @File : pdf.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"errors"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"io"
	"os"
	"path/filepath"
)

type PageOperationDecorator interface {
	Decorate(options *wkhtmltopdf.PageOptions) *wkhtmltopdf.PageOptions
}

type Html2PdfStreamSourceReq struct {
	PdfStoreBasePath string
	PdfNewName       string
	// 彩色还是黑白模式,true 代表黑白
	GrayMode bool

	PageDecorator         func(wkhtmltopdf.PageOptions) wkhtmltopdf.PageOptions
	PDFGeneratorDecorator func(generator *wkhtmltopdf.PDFGenerator)
}

type Html2PdfPureSourceReq struct {
	Html2PdfStreamSourceReq
	HtmlSource string
}
type Html2PdfPureReaderReq struct {
	Html2PdfStreamSourceReq
	Reader io.Reader
}
type Html2PdfResp struct {
	PdfStorePath string
}

func NewConfigableHtml2PdfReq(source string) Html2PdfPureSourceReq {
	r := Html2PdfPureSourceReq{}
	r.HtmlSource = source
	return r
}
func NewLocalHtml2PdfPureSourceReq(source string, pdfStorePath, newName string) Html2PdfPureSourceReq {
	return newPureSourceReq(source, pdfStorePath, newName)
}

func newPureSourceReq(source string, pdfStorePath, newName string) Html2PdfPureSourceReq {
	r := Html2PdfPureSourceReq{}
	r.HtmlSource = source
	r.PageDecorator = func(options wkhtmltopdf.PageOptions) wkhtmltopdf.PageOptions {
		options.FooterRight.Set("[PAGE]")
		options.FooterFontSize.Set(10)
		options.Zoom.Set(0.95)
		return options
	}
	r.PDFGeneratorDecorator = func(pdfg *wkhtmltopdf.PDFGenerator) {
		// Set global options
		pdfg.Dpi.Set(300)
		pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
		pdfg.Grayscale.Set(r.GrayMode)
	}
	r.PdfNewName = newName
	r.PdfStoreBasePath = pdfStorePath
	return r
}
func NewRemotetHtml2PdfPureSourceReq(source string, pdfStorePath, newName string) Html2PdfPureSourceReq {
	return newPureSourceReq(source, pdfStorePath, newName)
}

func Html2PdfByReader(req Html2PdfPureReaderReq) (Html2PdfResp, error) {
	var (
		result Html2PdfResp
	)
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return result, err
	}
	req.PDFGeneratorDecorator(pdfg)
	page := wkhtmltopdf.NewPageReader(req.Reader)
	page.PageOptions = req.PageDecorator(page.PageOptions)
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return result, err
	}

	filePath := req.PdfStoreBasePath + string(filepath.Separator)
	if !IsFileOrDirExists(filePath) {
		if err := CreateMultiFileDirs(filePath); nil != err {
			return result, errors.New("创建文件夹失败:" + err.Error())
		}
	}
	filePath += req.PdfNewName + ".pdf"
	// Write buffer contents to file on disk
	err = pdfg.WriteFile(filePath)
	if err != nil {
		return result, err
	}
	result.PdfStorePath = filePath

	return result, nil
}

func Html2PdfBySource(req Html2PdfPureSourceReq) (Html2PdfResp, error) {
	var (
		result Html2PdfResp
	)
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return result, err
	}
	req.PDFGeneratorDecorator(pdfg)

	page := wkhtmltopdf.NewPage(req.HtmlSource)
	page.PageOptions = req.PageDecorator(page.PageOptions)

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return result, err
	}

	filePath := req.PdfStoreBasePath + string(filepath.Separator)

	if !IsFileOrDirExists(filePath) {
		if err := CreateMultiFileDirs(filePath); nil != err {
			return result, errors.New("创建文件夹失败:" + err.Error())
		}
	}
	filePath += string(filepath.Separator) + req.PdfNewName + ".pdf"
	if IsFileOrDirExists(filePath) {
		os.Remove(filePath)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(filePath)
	if err != nil {
		return result, err
	}
	result.PdfStorePath = filePath

	return result, nil
}
