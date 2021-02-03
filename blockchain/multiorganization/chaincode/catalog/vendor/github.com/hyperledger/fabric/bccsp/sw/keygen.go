/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sw

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/hyperledger/fabric/bccsp"
)

type ecdsaKeyGenerator struct {
	curve elliptic.Curve
}

func (kg *ecdsaKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	privKey, err := ecdsa.GenerateKey(kg.curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("Failed generating ECDSA key for [%v]: [%s]", kg.curve, err)
	}

	return &ecdsaPrivateKey{privKey}, nil
}

type sm2KeyGenerator struct {
	curve elliptic.Curve
}

// 通过环境变量获取生成私钥的jar包路径
func genKeyPrivateKeyInternal() (*sm2.PrivateKey, error) {

	var defaultPath string
	if runtime.GOOS == "windows" {
		defaultPath = "D:/"
	} else {
		defaultPath = "/home/services/go/bin"
	}
	path := defaultPath
	tmpPath := os.Getenv("FABRIC_SM2_KEY_GENERATOR_PATH")
	if tmpPath != "" {
		path = tmpPath
	}
	var jarPath string
	if _, err := os.Stat(tmpPath); err == nil {
		jarPath = tmpPath
	} else {
		jarPath = filepath.Join(path, "sm2KeyGenerator.jar")
		if _, err := os.Stat(jarPath); os.IsNotExist(err) {
			return nil, fmt.Errorf("sm2Generator jar file[%s] not exists", jarPath)
		}
	}



	cmd := exec.Command("java", "-jar", jarPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, string(output))
	}

	var privateKey sm2.PrivateKey
	d, err := base64.StdEncoding.DecodeString(string(output))
	if err != nil {
		return nil, err
	}
	privateKey.D = new(big.Int).SetBytes(d)
	curve := sm2.P256Sm2()
	privateKey.Curve = curve
	x, y := curve.ScalarBaseMult(d)
	privateKey.X, privateKey.Y = x, y
	return &privateKey, nil
}

func (kg *sm2KeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	//privateKey, err := sm2.GenerateKey()
	privateKey, err := genKeyPrivateKeyInternal()
	if err != nil {
		return nil, fmt.Errorf("Failed generating SM2 key for [%v]: [%s]", kg.curve, err.Error())
	}
	return &sm2PrivateKey{privateKey}, nil
}

type aesKeyGenerator struct {
	length int
}

func (kg *aesKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	lowLevelKey, err := GetRandomBytes(int(kg.length))
	if err != nil {
		return nil, fmt.Errorf("Failed generating AES %d key [%s]", kg.length, err)
	}

	return &aesPrivateKey{lowLevelKey, false}, nil
}

type rsaKeyGenerator struct {
	length int
}

func (kg *rsaKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	lowLevelKey, err := rsa.GenerateKey(rand.Reader, int(kg.length))

	if err != nil {
		return nil, fmt.Errorf("Failed generating RSA %d key [%s]", kg.length, err)
	}

	return &rsaPrivateKey{lowLevelKey}, nil
}
