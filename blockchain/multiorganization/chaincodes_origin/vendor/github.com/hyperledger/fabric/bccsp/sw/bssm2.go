/*
Copyright IBM Corp. 2016 All Rights Reserved.

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
	"fmt"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
	"strings"

	"github.com/hyperledger/fabric/bccsp"
)

func signSM2(k *sm2.PrivateKey, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	r, s, err := sm2.Sm2Sign(k, digest, nil)
	if err != nil {
		return nil, err
	}

	//pubKey := &ecdsa.PublicKey{
	//	Curve: k.PublicKey.Curve,
	//	X:     k.PublicKey.X,
	//	Y:     k.PublicKey.Y,
	//}
	//
	//s, _, err = utils.SM2ToLowS(pubKey, s)
	//if err != nil {
	//	return nil, err
	//}

	//return utils.MarshalECDSASignature(r, s)
	return encodeSignature(r, s), nil
}

// 将两个大整数拼接成字符串
func encodeSignature(r, s *big.Int) []byte {
	// 缺位补足
	r1 := r.Text(16)
	if len(r1) <64 {
		for i := len(r1); i < 64; i++ {
			r1 = "0" + r1
		}
	}

	// 缺位补足
	s1 := s.Text(16)
	if len(s1) < 64 {
		for i := len(s1); i < 64; i++ {
			s1 = "0" + s1
		}
	}

	return []byte(fmt.Sprintf("%s%s", r1, s1))
}

// 将签名值转换成2个大整数
func decodeSignature(signatureBytes []byte) (*big.Int, *big.Int, error) {
	signature := string(signatureBytes)
	// 去除前导0
	signatureR := strings.TrimLeft(signature[:64], "0")
	signatureS := strings.TrimLeft(signature[64:], "0")
	var r, s big.Int
	_, err := fmt.Sscanf(signatureR, "%x", &r)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse big.Integer R from signature")
	}
	_, err = fmt.Sscanf(signatureS, "%x", &s)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse big.Integer S from signature")
	}
	return &r, &s, nil
}

func verifySM2(k *sm2.PublicKey, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	// 签名是由2个32个字节的大整数拼接而成
	r, s, err := decodeSignature(signature)
	if err != nil {
		return false, err
	}
	//r, s, err := utils.UnmarshalSM2Signature(signature)
	//if err != nil {
	//	return false, fmt.Errorf("Failed unmashalling signature [%s]", err)
	//}

	//pubKey := &ecdsa.PublicKey{
	//	Curve: k.Curve,
	//	X:     k.X,
	//	Y:     k.Y,
	//}
	//lowS, err := utils.SM2IsLowS(pubKey, s)
	//if err != nil {
	//	return false, err
	//}
	//
	//if !lowS {
	//	return false, fmt.Errorf("Invalid S. Must be smaller than half the order [%s][%s]", s, utils.GetCurveHalfOrdersAt(k.Curve))
	//}
	return sm2.Sm2Verify(k, digest, nil, r, s), nil
}

type sm2Signer struct{}

func (s *sm2Signer) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	return signSM2(k.(*sm2PrivateKey).privKey, digest, opts)
}

type sm2PrivateKeyVerifier struct{}

func (v *sm2PrivateKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return verifySM2(&(k.(*sm2PrivateKey).privKey.PublicKey), signature, digest, opts)
}

type sm2PublicKeyKeyVerifier struct{}

func (v *sm2PublicKeyKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return verifySM2(k.(*sm2PublicKey).pubKey, signature, digest, opts)
}
