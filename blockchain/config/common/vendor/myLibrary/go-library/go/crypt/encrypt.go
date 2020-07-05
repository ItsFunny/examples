package encrypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/prometheus/common/log"
	"github.com/wumansgy/goEncrypt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
)

type RSAEncryptModel struct {
	EncryptTime int64
	PublickKey  []byte
	PrivateKey  []byte
}

func Gen1024RSAKey() ([]byte, []byte, error) {
	return GenRsaKey(1024)
}

// RSA公钥私钥产生
func GenRsaKey(bits int) ([]byte, []byte, error) {
	var (
		privateBytes, publicBytes []byte
	)
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return nil, nil, err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return nil, nil, err
	}
	if bytes, err := ioutil.ReadFile("private.pem"); nil != err {
		fmt.Println(err)
	} else {
		// fmt.Println(hex.EncodeToString(bytes))
		// privateK = string(bytes)
		privateBytes = bytes
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return nil, nil, err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return nil, nil, err
	}
	if bytes, err := ioutil.ReadFile("public.pem"); nil != err {
		return nil, nil, err
	} else {
		// fmt.Println(hex.EncodeToString(bytes))
		// publicK = string(bytes)
		publicBytes = bytes
	}
	// go func() {
	// 	os.Remove("public.pem")
	// 	os.Remove("private.pem")
	// }()
	return privateBytes, publicBytes, nil
}

func encrypt(datas, key []byte) ([]byte, error) {
	return goEncrypt.RsaEncrypt(datas, key)
}
func decrypt(datas, key []byte) ([]byte, error) {
	return goEncrypt.RsaDecrypt(datas, key)
}

func RSASign(item interface{}, model RSAEncryptModel) ([]byte, error) {
	bytes, _ := json.Marshal(item)
	if cryptText, err := goEncrypt.RsaSign(bytes, model.PrivateKey); nil != err {
		log.Errorf("[RSAEncrypt] sign failed:%v", err.Error())
		return nil, err
	} else {
		return cryptText, nil
	}
}

func RSAVeriSign(item interface{}, signCodes []byte, model RSAEncryptModel) bool {
	bytes, _ := json.Marshal(item)
	return goEncrypt.RsaVerifySign(bytes, signCodes, model.PublickKey)
}

// rsa 加密
func RSAEncryptByPub(item interface{}, model RSAEncryptModel) ([]byte, error) {
	bytes, _ := json.Marshal(item)
	if cryptText, err := encrypt(bytes, model.PublickKey); nil != err {
		log.Errorf("[RSAEncrypt] encrypt failed:%v", err.Error())
		return nil, err
	} else {
		return cryptText, nil
	}
}

func RSAEncryptByPrv(data interface{}, model RSAEncryptModel) ([]byte, error) {
	bytes, _ := json.Marshal(data)
	if cryptText, err := encrypt(bytes, model.PrivateKey); nil != err {
		log.Errorf("[RSAEncrypt] encrypt failed:%v", err.Error())
		return nil, err
	} else {
		return cryptText, nil
	}
}

func RSADecrypt(encryptStr string, key []byte) ([]byte, error) {
	// if plainText, err := goEncrypt.RsaDecrypt([]byte(encryptStr), property.PrivateBytes); nil != err {
	bytes, e := hex.DecodeString(encryptStr)
	if nil != e {
		return nil, e
	}
	if plainText, err := goEncrypt.RsaDecrypt(bytes, key); nil != err {
		log.Errorf("[RSADecrypt]faield:%v", err.Error())
		return nil, err
	} else {
		return plainText, nil
	}
}

func MD5Encrypt(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func MD5EncryptByBytes(str string) string {
	w := md5.New()
	io.WriteString(w, str)                   // 将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil)) // w.Sum(nil)将w的hash转成[]byte格式
	return md5str2
}
func MD5EncryptFile(file *os.File) string {
	h := md5.New()
	io.Copy(h, file)
	return hex.EncodeToString(h.Sum(nil))
}

func HashHMacEncrypt(key string, data string) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(key))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string

	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println("Result: " + sha)

	return sha
}

// 对消息的散列值生成数字签名
// 结果以hex的形式打印
// func ECCSignWithHex(msg []byte, prk *ecdsa.PrivateKey) (string, error) {
func ECCSignWithHex(msgStr string, privateKeyBytes []byte) (string, error) {
	bytes, e := ECCSign(msgStr, privateKeyBytes)
	if nil != e {
		return "", e
	}
	return hex.EncodeToString(bytes), nil
}

func ECCSign(msgStr string, privateKeyBytes []byte) ([]byte, error) {
	msg := []byte(msgStr)
	return ECCSignWithBytes(msg, privateKeyBytes)
}

func ECCSignWithBytes(msgBytes []byte, privateKeyBytes []byte) ([]byte, error) {
	// 取得私钥
	prk, e := Bytes2ECDSAPrv(privateKeyBytes)

	if nil != e {
		return nil, e
	}
	// r, s, err := ecdsa.Sign(rand.Reader, prk, msg)
	// if err != nil {
	// 	return "", err
	// }
	// params := prk.Curve.Params()
	// curveOrderByteSize := params.P.BitLen() / 8
	// rBytes, sBytes := r.Bytes(), s.Bytes()
	// signature := make([]byte, curveOrderByteSize*2)
	// copy(signature[curveOrderByteSize-len(rBytes):], rBytes)
	// copy(signature[curveOrderByteSize*2-len(sBytes):], sBytes)
	// return hex.EncodeToString(signature), nil
	//
	hash := sha256.New()
	// 计算哈希值
	// 填入数据
	hash.Write(msgBytes)
	bytes := hash.Sum(nil)

	// 对哈希值生成数字签名
	r, s, err := ecdsa.Sign(rand.Reader, prk, bytes)
	if err != nil {
		return nil, err
	}
	params := prk.Curve.Params()
	curveOrderByteSize := params.P.BitLen() / 8
	rBytes, sBytes := r.Bytes(), s.Bytes()
	signature := make([]byte, curveOrderByteSize*2)
	copy(signature[curveOrderByteSize-len(rBytes):], rBytes)
	copy(signature[curveOrderByteSize*2-len(sBytes):], sBytes)

	return signature, nil
}

// 验证数字签名
// data:指的是预期的值,匹对的时候直接[]byte 强转即可
// signature: 签名, 既经过hex编码后的string,所以需要先解码
func ECCVerifySignWithHex(data string, signature string, pubBytes []byte) bool {
	sg, err := hex.DecodeString(signature)
	if nil != err {
		return false
	}
	puk := Bytes2ECDSAPub(pubBytes)
	// 计算哈希值
	hash := sha256.New()
	hash.Write([]byte(data))
	bytes := hash.Sum(nil)

	curveOrderByteSize := puk.Curve.Params().P.BitLen() / 8
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(sg[:curveOrderByteSize])
	s.SetBytes(sg[curveOrderByteSize:])
	// 验证数字签名
	return ecdsa.Verify(puk, bytes, r, s)
}

func ECCVerify(data []byte, sg []byte, pubBytes []byte) bool {
	puk := Bytes2ECDSAPub(pubBytes)
	// 计算哈希值
	hash := sha256.New()
	hash.Write(data)
	bytes := hash.Sum(nil)

	curveOrderByteSize := puk.Curve.Params().P.BitLen() / 8
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(sg[:curveOrderByteSize])
	s.SetBytes(sg[curveOrderByteSize:])
	// 验证数字签名
	return ecdsa.Verify(puk, bytes, r, s)
}

// ECC 加密
// hex 是指返回值被hex修饰
func ECCEncryptWithHex(msg string, pubKBytes []byte) (string, error) {
	pubK := Bytes2ECDSAPub(pubKBytes)
	public := ecies.ImportECDSAPublic(pubK)

	ct, err := ecies.Encrypt(rand.Reader, public, []byte(msg), nil, nil)
	if nil != err {
		return "", err
	}
	return hex.EncodeToString(ct), nil
}

// hex 是指参数需要先hex解码
func ECCDecryptWithHex(msg string, prvBytes []byte) (string, error) {
	prvK, e := Bytes2ECDSAPrv(prvBytes)
	if nil != e {
		return "", e
	}
	bytes, e := hex.DecodeString(msg)
	if nil != e {
		return "", e
	}
	eciesPrvK := ecies.ImportECDSA(prvK)
	m, e := eciesPrvK.Decrypt(bytes, nil, nil)

	return string(m), e
}

// ecdsa 公钥转为字节数组
func ECDSAPub2Bytes(pub *ecdsa.PublicKey) []byte {
	return crypto.FromECDSAPub(pub)
}

// ECDSA私钥 -> []byte
// FromECDSA exports a private key into a binary dump.
func ECDSAPrv2Bytes(priv *ecdsa.PrivateKey) []byte {
	return crypto.FromECDSA(priv)
}

// []byte 转为 ecdsa公钥
func Bytes2ECDSAPub(pub []byte) *ecdsa.PublicKey {
	if len(pub) == 0 {
		return nil
	}

	x, y := elliptic.Unmarshal(crypto.S256(), pub)
	return &ecdsa.PublicKey{Curve: crypto.S256(), X: x, Y: y,
	}
}

func Bytes2ECDSAPrv(prvBytes []byte) (*ecdsa.PrivateKey, error) {
	key, e := crypto.ToECDSA(prvBytes)

	return key, e
}
