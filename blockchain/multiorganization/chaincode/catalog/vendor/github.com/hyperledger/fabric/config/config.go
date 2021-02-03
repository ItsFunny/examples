/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-07 15:34 
# @File : config.go
# @Description : 
# @Attention : 
*/
package config

import (
	"github.com/hyperledger/fabric/common/flogging"
	"io/ioutil"
	"os"
)

type OnError func(err error)

func GetOrDefaultFileBytes(envFileConstants string, defaultFilePath string, logger *flogging.FabricLogger, onErr OnError) []byte {
	var configPath string
	if envF := os.Getenv(envFileConstants); len(envF) != 0 {
		configPath = envF
	} else {
		configPath = defaultFilePath
	}
	logWarnIfNotNil(logger, "加密机的配置路径为:"+configPath)
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		onErr(err)
		logWarnIfNotNil(logger, "加密机配置文件不存在")
		return nil
	}
	bytes, err := ioutil.ReadFile(configPath)
	if nil != err {
		onErr(err)
		logWarnIfNotNil(logger, "读取加密机配置失败:"+err.Error())
		return nil
	}
	logWarnIfNotNil(logger, "加密机配置为:"+string(bytes))
	return bytes
}

func GetOrDefaultFileBytesPanic(envFileConstants string, defaultFilePath string, logger *flogging.FabricLogger, errorDecorator func(err error) error) []byte {
	return GetOrDefaultFileBytes(envFileConstants, defaultFilePath, logger, func(err error) {
		panic(errorDecorator(err))
	})
}
func logWarnIfNotNil(logger *flogging.FabricLogger, msg string) {
	if nil != logger {
		logger.Warn(msg)
	}
}
