/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-02 10:41 
# @File : chinese.go
# @Description : 
*/
package converter

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

var (
	DEFAULT_SORTED_BY_CHINESE = func(strI, strJ string, desc bool) bool {
		hashI, _ := UTF82GBK(strI)
		hashJ, _ := UTF82GBK(strJ)
		bLen := len(hashI)
		for idx, chr := range hashI {
			if idx > bLen-1 {
				if desc {
					return false
				} else {
					return true
				}

			}
			if chr != hashJ[idx] {
				if desc {
					return chr < hashJ[idx]
				} else {
					return chr > hashJ[idx]
				}
			}
		}
		return true
	}
)
