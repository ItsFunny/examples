/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-12 11:58 
# @File : http_util.go
# @Description : 
# @Attention : 
*/
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

type HttpReq struct {
	TimeOutSeconds int
	MaxFailCount   int
	HttpUrl        string
	BytesData      []byte
	Headers        map[string]string
	// res,bool: 是否停止重试,error: 错误信息
	HandlerResponse func(response *http.Response) (*HttpResult, bool, error)
}


func BuildHttpRequest(url string, data []byte, heads map[string]string) (*http.Request, error) {
	request, e := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if nil != e {
		return nil, errors.New("创建httpRequest失败")
	}
	for k, v := range heads {
		request.Header.Set(k, v)
	}
	return request, nil

}

func PostPerClient(req HttpReq) (interface{}, error) {
	result, e := httpForSec(req)
	if nil != e {
		return "", e
	}
	return result.Data, nil
}

func httpForSec(req HttpReq) (*HttpResult, error) {

	client := &http.Client{}
	// defer client.CloseIdleConnections()

	// failCount := 0
	// var lastErr error
	// for failCount <= req.MaxFailCount {
	// 	httpReq, e := BuildHttpRequest(req.HttpUrl, req.BytesData, req.Headers)
	// 	if nil != e {
	// 		panic("构建http请求参数失败:" + e.Error())
	// 	}
	// 	// httpReq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// 	httpResp, e := client.Do(httpReq)
	// 	if nil != e {
	// 		failCount++
	// 		log.Println("http调用失败:" + e.Error())
	// 		lastErr = e
	// 		continue
	// 	}
	// 	if res, keepRetry, e := req.HandlerResponse(httpResp); nil != e {
	// 		log.Println("http调用失败,状态为:" + httpResp.Status + ",错误信息为:"+e.Error())
	// 		if keepRetry {
	// 			failCount++
	// 			log.Println("开始失败重试,failCount=" + strconv.Itoa(failCount))
	// 			continue
	// 		} else {
	// 			log.Println("异常停止,停止失败重试")
	// 			lastErr = e
	// 			break
	// 		}
	// 	} else{
	// 		return res,nil
	// 	}
	// }
	// return nil,lastErr

	resultChan := make(chan *HttpResult)
	errChan := make(chan error)
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*time.Duration(req.TimeOutSeconds))

	go func() {
		defer func() {
			if ePanic := recover(); nil != ePanic {
				if err, ok := ePanic.(error); ok {
					errChan <- errors.New("http时发生了panic,错误为:" + err.Error())
				} else {
					marshal, _ := json.Marshal(ePanic)
					if len(marshal) == 0 {
						marshal = []byte("unknown")
					}
					errChan <- errors.New("未知panic:" + string(marshal))
				}
			}
		}()
		failCount := 0
		var lastErr error
		for failCount <= req.MaxFailCount {
			httpReq, e := BuildHttpRequest(req.HttpUrl, req.BytesData, req.Headers)
			if nil != e {
				panic("构建http请求参数失败:" + e.Error())
			}
			// httpReq.Header.Set("Content-Type", "application/json;charset=UTF-8")
			httpResp, e := client.Do(httpReq)
			if nil != e {
				failCount++
				log.Println("http调用失败:" + e.Error())
				lastErr = e
				continue
			}
			if res, keepRetry, e := req.HandlerResponse(httpResp); nil != e {
				log.Println("http调用失败,状态为:" + httpResp.Status + ",错误信息为:")
				if keepRetry {
					failCount++
					log.Println("开始失败重试,failCount=" + strconv.Itoa(failCount))
					continue
				} else {
					log.Println("异常停止,停止失败重试")
					lastErr = e
					break
				}
			} else {
				resultChan <- res
			}
		}
		errChan <- lastErr
	}()

	select {
	case res := <-resultChan:
		cancelFunc()
		return res, nil
	case err := <-errChan:
		return nil, err
	case <-timeout.Done():
		return nil, errors.New("http请求达到超时也无法获取到aes密钥,检查Java服务器是否启动着")
	}
}
