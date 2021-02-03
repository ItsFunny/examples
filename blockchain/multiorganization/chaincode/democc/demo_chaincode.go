/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-03-26 13:23 
# @File : demo_chaincode.go
# @Description : 
# @Attention : 
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"log"
	"strconv"
	"strings"
	"time"
)

type DemoChainCode struct {
}

func (this *DemoChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("init")
	return shim.Success(nil)
}

type HistoryNode struct {
	Timestamp string `json:"timestamp"`
	TxId      string `json:"txId"`
}

func (this *DemoChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	s, args := stub.GetFunctionAndParameters()
	strs := ""
	for _, v := range args {
		strs += v + ","
	}
	s = strings.ToLower(s)
	key, value := "key", "joker"
	if len(args) == 1 {
		key = args[0]
	}
	if len(args) >= 2 {
		key = args[0]
		value = args[1]
	}

	var result []byte
	funcName := "setvalue"
	var e error
	code := OK

	switch s {
	case "setvalue":
		e = stub.PutState(key, []byte(value))
		result = []byte("ok")
	case "getvalue":
		funcName = "getvalue"
		bytes, err := stub.GetState(key)
		if nil != err {
			code = FAIL
			e = err
		} else if len(bytes) != 0 {
			result = bytes
		} else {
			code = DATA_NOT_EXIST
			result = []byte("123")
		}

	case "gethistory":
		funcName = "gethistory"
		originCodes := make([]int, 0)
		for i := 0; i < 4; i++ {
			iteratorInterface, err := stub.GetHistoryForKey(key)
			if nil != err {
				e = err
				code = HISTORY_ERROR
			} else if nil == iteratorInterface {
				code = HISTORY_EMPTY
			} else {
				times := make([]HistoryNode, 0)
				if !iteratorInterface.HasNext() {
					code = HISTORY_DOESNT_HAS_NEXT
				} else {
					code = OK
					for iteratorInterface.HasNext() {
						modification, err := iteratorInterface.Next()
						if nil != err {
							code = HISTORY_ITERATOR_FAIL
							e = err
						} else {
							timestamp := strconv.FormatInt(modification.Timestamp.Seconds, 10)
							times = append(times, HistoryNode{
								Timestamp: timestamp,
								TxId:      modification.TxId,
							})
						}
					}
					bytes, _ := json.Marshal(times)
					result = bytes
				}
			}
			// 如果之前没有查到数据,则多查询几次,如果中途查到一次,则这个code 会为 OK
			if code == HISTORY_ITERATOR_FAIL || code == HISTORY_EMPTY || code == HISTORY_DOESNT_HAS_NEXT {
				originCodes = append(originCodes, code)
				code = HISTORY_CONTINUE_FIND + code
				// 如果code 为1000 ,1001, 1002,代表的是第2,第3次才找到数据
			} else {
				break
			}
		}
		if code == OK && len(originCodes) != 0 {
			// 说明发生了查询 多次才查询到这个数据,则此时用新的值替代
			code = HISTORY_FINALLY_FOUND_AFTER_SEVERALTIMES
			sList := make([]string, 0)
			for _, c := range originCodes {
				sList = append(sList, strconv.Itoa(c))
			}
			join := strings.Join(sList, ",")
			result = append(result, []byte(join)...)
		}
	}
	printMsg(funcName, key, result, e)
	msg := ""
	if e != nil {
		msg = e.Error()
	}
	return shim.Success(DefaultWith(code, msg, result))
}
func printMsg(funcN string, key string, data []byte, e error) {
	sb := strings.Builder{}
	sb.WriteString(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
	d := ""
	err := ""
	if len(data) != 0 {
		d = string(data)
	}
	if nil != e {
		err = e.Error()
	}
	sb.WriteString(fmt.Sprintf("当前时间:[ %s ]  ,方法= [ %s ]  ,key=[ %s ],data=[ %s ] ,err=[ %s ]\n", time.Now().String(), funcN, key, d, err))

	sb.WriteString("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n\n")

	log.Println(sb.String())
}

func main() {
	u := new(DemoChainCode)
	if err := shim.Start(u); nil != err {
		panic(err)
	}
}
