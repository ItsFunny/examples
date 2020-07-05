/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-04 14:28 
# @File : file.go
# @Description : 
*/
package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// 判断所给路径文件/文件夹是否存在
func IsFileOrDirExists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
// 判断文件是否存在
func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func CreateMultiFileDirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 获取当前可执行文件的路径
func GetCurrentExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	// fmt.Println("path111:", path)
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}
	// fmt.Println("path222:", path)
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return "", errors.Errorf(`Can't find "/" or "\".`)
	}
	// fmt.Println("path333:", path)
	return string(path[0 : i+1]), nil
}

// 获取文件大小
func GetFileSize(path string) (int64, error) {
	// 判断文件是否存在
	exists := IsFileOrDirExists(path)
	if !exists {
		return 0, errors.New("文件不存在")
	}
	// 需要判断是否是文件
	if b := IsFile(path); !b {
		return 0, errors.New("该路径不是文件")
	}
	size := int64(0)
	if err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		size = f.Size()
		return nil
	}); nil != err {
		return 0, err
	}

	return size / 1000, nil
}

// 使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
// 创建并且写入文件
func Write2File(filePath string, bytes []byte) error {
	if exists := IsFileOrDirExists(filePath); exists {
		return errors.New("文件已经存在")
	}
	// ./a.txt  ||  /user/local/c.txt  || ./a/b.txt
	count := strings.Count(filePath, "/")
	if count > 1 {
		// 说明有文件夹
		// 判断文件夹是否存在
		dirs := SubString(filePath, strings.LastIndex(filePath, "/"))
		if !IsFileOrDirExists(dirs) {
			if err := CreateMultiFileDirs(dirs); nil != err {
				fmt.Println("创建文件夹失败:", err.Error())
				return err
			}
		}
	}

	file, cErr := os.Create(filePath)
	if nil != cErr {
		fmt.Println("创建文件失败:", cErr.Error())
		return cErr
	}
	// file, e := os.Open(filePath)
	// if nil != e {
	// 	return errors.New("打开文件失败")
	// }
	defer file.Close()
	if _, e := file.Write(bytes); nil != e {
		fmt.Println("写入到文件失败:", e.Error())
		return e
	}

	return nil
}

// 获取某个文件夹下符合某种后缀的所有文件路径
// 通过ioutil.ReadDir读取
// TODO 后期改为接口,变长数组,多个过滤
func GetFilesBelownDirFilterBySuffix(dir string, filterSuffix string) ([]string, error) {
	if !IsDir(dir) {
		return nil, errors.New(fmt.Sprintf("[%s]不是文件夹", dir))
	}
	infos, e := ioutil.ReadDir(dir)
	if nil != e {
		fmt.Println("读取文件夹错误", e.Error())
		return nil, e
	}
	res := make([]string, 0)
	for _, f := range infos {
		if path.Ext(f.Name()) == filterSuffix {
			res = append(res, dir+string(filepath.Separator)+f.Name())
		}
	}

	return res, nil
}

// 获取目录下的所有文件
func GetAllFilesUnderDir(dirPath string) ([]*os.File, error) {
	files := make([]*os.File, 0)
	if e := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		// 跳过
		if path == dirPath {
			return nil
		}
		if nil != err {
			return errors.New("读取目录失败:" + err.Error())
		}
		if f, e := os.Open(path); nil != e {
			return errors.Errorf("读取文件[%s]的时候失败:%s", path, e.Error())
		} else {
			files = append(files, f)
		}
		return nil
	}); nil != e {
		return nil, errors.New("filepath#walk 调用失败:" + e.Error())
	}

	return files, nil
}

// 删除文件或者文件夹
func DeleteFileOrDir(path string) error {
	if IsFileOrDirExists(path) {
		return errors.New(fmt.Sprintf("path=[%s]不存在", path))
	}
	return os.RemoveAll(path)
}
