/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-26 14:13 
# @File : http.go
# @Description : 
*/
package utils

import (
	"context"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
)

func ConvtMap2FastHttpArgs(params map[string]interface{}) *fasthttp.Args {
	args := &fasthttp.Args{}
	for key, value := range params {
		switch value.(type) {
		case string:
			args.Set(key, value.(string))
		case int:
			args.SetUint(key, value.(int))
		}
	}
	return args
}

// post表单的方式
func DoPostForm(ctx context.Context, postUrl string, keys []string, postValues []string) (*http.Response, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		values := url.Values{}
		l := len(keys)
		for i := 0; i < l; i++ {
			values.Add(keys[i], postValues[i])
		}
		return http.PostForm(postUrl, values)
	}
}
