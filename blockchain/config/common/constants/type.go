/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-15 12:51 
# @File : type.go
# @Description : 
# @Attention : 
*/
package constants

const (
	// 单独的作品
	ITEM_TYPE_SINGLE   = 1
	ITEM_TYPE_CONTINUS = 2
)

// 文学文字类型
const (
	// 纯文字
	ARTICLE_CONTENT_TYPE_PURE_WORDS = 1
	// 文件类型
	ARTICLE_CONTENT_TYPE_FILE = 2
	// PDF
	ARTICLE_CONTENT_TYPE_PDF  = ARTICLE_CONTENT_TYPE_FILE<<1 | ARTICLE_CONTENT_TYPE_FILE
)
