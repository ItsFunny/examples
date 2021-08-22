/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/6/30 8:58 上午
# @File : jz_09_青蛙跳台阶扩展问题_test.go.go
# @Description :
# @Attention :
*/
package offer

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_jumpFloorII(t *testing.T) {
	ii := jumpFloorII(3)
	fmt.Println(ii)
}

type A struct {
	A string
}

func Test_aad(t *testing.T) {
	var v A
	checkInterface(v)
	checkInterface(&v)
}
func checkInterface(v interface{}) {
	if v == nil {
		fmt.Println(1)
	} else {
		marshal, err := json.Marshal(v)
		if nil != err {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(marshal))
		}
	}
}

func Test_Print(t *testing.T) {
	path1 := "/Users/joker/Desktop/gopath/src/github.com/hyperledger/fabric/orderer/consensus/bft/impl/1.log"
	path2 := "/Users/joker/Desktop/gopath/src/github.com/hyperledger/fabric/orderer/consensus/bft/impl/2.log"
	file, _ := ioutil.ReadFile(path1)
	file2, _ := ioutil.ReadFile(path2)

	l := len(file)
	if len(file2) < l {
		l = len(file2)
	}
	index := 0
	for i := 0; i < l; i++ {
		if file[i] != file2[i] {
			index = i
			break
		}
	}
	fmt.Println(file[index], file2[index])
	beforeSame := file[:index]
	beforeSame2 := file2[:index]
	fmt.Println(bytes.Equal(beforeSame, beforeSame2))
	str1 := string(beforeSame)
	str2 := string(beforeSame2)
	fmt.Println(str1)
	fmt.Println(str2)

	newFile := file[index:]
	newFile2 := file2[index:]
	if bytes.Equal(newFile2, newFile) {
		fmt.Println("equal")
	}
	str1 = string(newFile)
	str2 = string(newFile2)
	fmt.Println(str1)
	fmt.Println(str2)
}

func Test_asd(t *testing.T) {
	str := "PP+BAwEBCEJsb2NrTXNnAf+CAAEDAQdDaGFpbklkAQwAAQpMYXN0SGVpZ2h0AQQAAQVCYXRjaAH/iAAAACH/hwIBARJbXSpjb21tb24uRW52ZWxvcGUB/4gAAf+EAAB1/4MDAQL/hAABBgEHUGF5bG9hZAEKAAEJU2lnbmF0dXJlAQoAAQlFeHRyYUluZm8BCgABFFhYWF9Ob1Vua2V5ZWRMaXRlcmFsAf+GAAEQWFhYX3VucmVjb2duaXplZAEKAAENWFhYX3NpemVjYWNoZQEEAAAAFf+FAwEBCXN0cnVjdCB7fQH/hgAAAP6gNP+CAQ5zeXNkZW1vY2hhbm5lbAIBAf6f0A"
	// bs := []byte(str)
	// fmt.Println(string(bs))
	decodeString, _ := base64.StdEncoding.DecodeString(str)
	// fmt.Println(err)
	fmt.Println(string(decodeString))

}
