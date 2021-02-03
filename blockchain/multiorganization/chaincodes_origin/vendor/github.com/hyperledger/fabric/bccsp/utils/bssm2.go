/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/asn1"
	"errors"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
	"strings"
)

func IsSM2Cert(certBytes []byte) (*sm2.Certificate, bool, error) {
	certificate, err := sm2.ParseCertificate(certBytes)
	if err != nil {
		return nil, false, err
	}
	isSM2Cert := strings.Contains(certificate.SignatureAlgorithm.String(), "SM2") && strings.Contains(certificate.PublicKey.(*ecdsa.PublicKey).Curve.Params().Name, "SM2")
	return certificate, isSM2Cert, nil
}

type SM2Signature struct {
	R, S *big.Int
}

var (
	// curveHalfOrders contains the precomputed curve group orders halved.
	// It is used to ensure that signature' S value is lower or equal to the
	// curve group order halved. We accept only low-S signatures.
	// They are precomputed for efficiency reasons.
	sm2CurveHalfOrders = map[elliptic.Curve]*big.Int{
		sm2.P256Sm2(): new(big.Int).Rsh(sm2.P256Sm2().Params().N, 1),
	}
)

func SM2GetCurveHalfOrdersAt(c elliptic.Curve) *big.Int {
	return big.NewInt(0).Set(sm2CurveHalfOrders[c])
}

func MarshalSM2Signature(r, s *big.Int) ([]byte, error) {
	return asn1.Marshal(SM2Signature{r, s})
}

func UnmarshalSM2Signature(raw []byte) (*big.Int, *big.Int, error) {
	// Unmarshal
	sig := new(SM2Signature)
	_, err := asn1.Unmarshal(raw, sig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed unmashalling signature [%s]", err)
	}

	// Validate sig
	if sig.R == nil {
		return nil, nil, errors.New("invalid signature, R must be different from nil")
	}
	if sig.S == nil {
		return nil, nil, errors.New("invalid signature, S must be different from nil")
	}

	if sig.R.Sign() != 1 {
		return nil, nil, errors.New("invalid signature, R must be larger than zero")
	}
	if sig.S.Sign() != 1 {
		return nil, nil, errors.New("invalid signature, S must be larger than zero")
	}

	return sig.R, sig.S, nil
}

func SM2SignatureToLowS(k *ecdsa.PublicKey, signature []byte) ([]byte, error) {
	r, s, err := UnmarshalSM2Signature(signature)
	if err != nil {
		return nil, err
	}

	s, modified, err := SM2ToLowS(k, s)
	if err != nil {
		return nil, err
	}

	if modified {
		return MarshalSM2Signature(r, s)
	}

	return signature, nil
}

// IsLow checks that s is a low-S
func SM2IsLowS(k *ecdsa.PublicKey, s *big.Int) (bool, error) {
	halfOrder, ok := sm2CurveHalfOrders[k.Curve]
	if !ok {
		return false, fmt.Errorf("curve not recognized [%s]", k.Curve)
	}

	return s.Cmp(halfOrder) != 1, nil

}

func SM2ToLowS(k *ecdsa.PublicKey, s *big.Int) (*big.Int, bool, error) {
	lowS, err := SM2IsLowS(k, s)
	if err != nil {
		return nil, false, err
	}

	if !lowS {
		// Set s to N - s that will be then in the lower part of signature space
		// less or equal to half order
		s.Sub(k.Params().N, s)

		return s, true, nil
	}

	return s, false, nil
}
