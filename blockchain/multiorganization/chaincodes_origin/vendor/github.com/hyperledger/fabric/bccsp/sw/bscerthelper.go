package sw

import (
	"github.com/hyperledger/fabric/bccsp"
	"github.com/tjfoc/gmsm/sm2"
	"io"
)

func CreateSM2Certificate(rand io.Reader, template, parent *sm2.Certificate, pub interface{}, key bccsp.Key) ([]byte, error) {
	privKey := key.(*sm2PrivateKey).privKey
	certBytes, err := sm2.CreateCertificate(rand, template, parent, pub, privKey)
	return certBytes, err
}