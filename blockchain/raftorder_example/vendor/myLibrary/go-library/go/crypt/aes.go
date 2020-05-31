/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-18 13:10 
# @File : aes.go
# @Description : 
*/
package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesCBCEncrypt(key []byte, originData []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(originData))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[aes.BlockSize:],
		originData)
	return ciphertext, nil
}

func AesCBCDecrypt(key []byte, encryptData []byte) (originData []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(encryptData) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := encryptData[:aes.BlockSize]
	encryptData = encryptData[aes.BlockSize:]
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(encryptData,encryptData)
	// cipher.NewCFBDecrypter(block, iv).XORKeyStream(encryptData, encryptData)
	originData = encryptData
	return
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type AesContainer struct {
	Versions         []uint64
	AesVersionKeyMap map[uint64]string
}

func (this *AesContainer) Add(version uint64, key string) {
	if nil == this.AesVersionKeyMap {
		this.AesVersionKeyMap = make(map[uint64]string)
	}
	this.AesVersionKeyMap[version] = key
	this.Versions = append(this.Versions, version)
}
func (this *AesContainer) GetEncKeyWithError(version uint64) (string, error) {
	if key, exist := this.AesVersionKeyMap[version]; !exist {
		return "", errors.New("版本号错误")
	} else {
		return key, nil
	}
}
func (this *AesContainer) GetLatest() (uint64, string) {
	version := this.Versions[len(this.Versions)-1]
	return version, this.AesVersionKeyMap[version]
}
