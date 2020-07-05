/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-23 16:05 
# @File : config.go
# @Description : 
# @Attention : 
*/
package config

import (
	error2 "examples/blockchain/config/common/error"
	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

func Config(path string) error2.IVlinkError {
	fmt.Println("加载配置文件")
	cf, err := LoadYaml(path)
	if nil != err {
		return err
	}
	configuration.Properties = cf
	fmt.Println("加载配置文件完毕")
	return nil
}

func LoadYaml(path string) (*VlinkProperties, error2.IVlinkError) {
	var conf VlinkProperties
	bytes, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, error2.NewConfigError(err, "读取文件错误")
	}
	err = yaml.Unmarshal(bytes, &conf)
	if nil != err {
		fmt.Println("yaml unmarshal occur err", err)
		return nil, error2.NewConfigError(err, "yaml反序列化错误")
	}
	e := conf.config()
	return &conf, e
}
