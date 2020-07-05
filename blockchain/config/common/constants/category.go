/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-01-13 09:50 
# @File : category.go
# @Description : 
# @Attention : 
*/
package constants

const (
	PHOTO                       = 1
	PHOTO_CATEGORY_CARTOON      = PHOTO<<1 | 1;
	PHOTO_CATEGORY_ILLUSTRATION = PHOTO_CATEGORY_CARTOON<<1 | 1
	PHOTO_CATEGORY_PHOTOGRAPH   = PHOTO_CATEGORY_ILLUSTRATION<<1 | 1

	VIDEO = 1 << 6

	MUSIC = 1 << 12

	ARTICLE = 1 << 18
)
