/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-16 14:15 
# @File : walley.go
# @Description : 
# @Attention : 
*/
package wallet

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/hbakhtiyor/schnorr"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ripemd160"
	"math/big"
	"strconv"
	"strings"
)

type Wallet struct {
	PubKey *bip32.Key
	PrvKey *bip32.Key
	// 路径,每个钱包都有根路径       0    1                2           3               4               5
	// Path的格式为: 遵循bip44协议: m / purpose(固定44)/ coin(货币编号)/ account(地址)/ change(出还是入) /address_index
	Path string
	// 其中 m 表示私钥，而 M 表示公钥。m/3/2/5 表示从主节点派生出来的第 4 个子节点的第 3 个孙节点的第 6 个重孙节点，派生过程中使用的是 CKD 函数，而 m/3'/2'/5' 则表示派生过程中使用的是 HCKD 函数。
	Position string
}

func (w Wallet) GetAddress() string {
	if w.Path != "" {
		pathes := strings.Split(w.Path, "/")
		return pathes[3]
	}
	return ""
}

// 主钱包
type HDWallet struct {
	Wallet
	// 助记词
	Mnemonic string
	// 记录子钱包的编号
	Index uint32
}

type ChildWallet struct {
	Wallet
}

func NewHDWallet(Type int, inOrOut int, passWord string) *HDWallet {
	w := new(HDWallet)
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, passWord)

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()
	w.PrvKey = masterKey
	w.PubKey = publicKey

	address := GenerateAddress(w.PubKey.Key)
	w.Path = "m/44/" + strconv.Itoa(Type) + "/" + address + "/" + strconv.Itoa(inOrOut) + "/" + strconv.Itoa(0)
	w.Position = "m"

	return w
}

func GenerateAddress(pubKeyBytes []byte) string {
	/* See https://en.bitcoin.it/wiki/Technical_background_of_Bitcoin_addresses */

	/* Convert the public key to bytes */
	pub_bytes := pubKeyBytes

	/* SHA256 Hash */
	// fmt.Println("2 - Perform SHA-256 hashing on the public key")
	sha256_h := sha256.New()
	sha256_h.Reset()
	sha256_h.Write(pub_bytes)
	pub_hash_1 := sha256_h.Sum(nil)
	// fmt.Println(byteString(pub_hash_1))
	// fmt.Println("=======================")

	/* RIPEMD-160 Hash */
	// fmt.Println("3 - Perform RIPEMD-160 hashing on the result of SHA-256")
	ripemd160_h := ripemd160.New()
	ripemd160_h.Reset()
	ripemd160_h.Write(pub_hash_1)
	pub_hash_2 := ripemd160_h.Sum(nil)
	// fmt.Println(byteString(pub_hash_2))
	// fmt.Println("=======================")
	/* Convert hash bytes to base58 check encoded sequence */
	address := b58checkencode(0x00, pub_hash_2)

	return address
}

func (w *HDWallet) SchnorrSign(message string) ([]byte, error) {
	var msg [32]byte
	copy(msg[:], []byte(message))
	bytes, e := schnorr.Sign(new(big.Int).SetBytes(w.PrvKey.Key), msg)
	if nil != e {
		return nil, e
	}
	b := make([]byte, 0)
	for i := 0; i < len(bytes); i++ {
		b = append(b, bytes[i])
	}
	return b, nil
}

func (w *HDWallet) SchnorrVerisign(exceptStr string, signatureStr string) bool {
	var (
		publicKey [33]byte
		message   [32]byte
		signature [64]byte
	)
	copy(publicKey[:], w.PubKey.Key)
	copy(message[:], []byte(exceptStr))
	bytes, e := hex.DecodeString(signatureStr)
	if nil != e {
		return false
	}
	copy(signature[:], bytes)

	b, e := schnorr.Verify(publicKey, message, signature)
	if nil != e {
		return false
	}
	return b
}

func (w *HDWallet) NewChildWallet(Type int) (*ChildWallet, error) {
	key, e := w.PrvKey.NewChildKey(w.Index + 1)
	if nil != e {
		return nil, e
	}
	indexStr := strconv.Itoa(int(w.Index))
	w.Index++
	cw := new(ChildWallet)
	cw.PubKey, cw.PrvKey = key.PublicKey(), key
	// Path的格式为: 遵循bip44协议: m / purpose(固定44)/ coin(货币编号)->这里更换为Type/ account(地址)/ change(出还是入) /address_index
	address := GenerateAddress(cw.PubKey.Key)
	// m,44,coin,account,change,address_index
	pathes := strings.Split(w.Path, "/")
	pathes[2] = strconv.Itoa(Type)
	pathes[3] = address
	pathes[5] = strconv.Itoa(int(w.Index))
	cw.Path = strings.Join(pathes, "/")
	cw.Position = w.Position + "/" + indexStr

	return cw, nil
}

// b58encode encodes a byte slice b into a base-58 encoded string.
func b58encode(b []byte) (s string) {
	/* See https://en.bitcoin.it/wiki/Base58Check_encoding */

	const BITCOIN_BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	/* Convert big endian bytes to big int */
	x := new(big.Int).SetBytes(b)

	/* Initialize */
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s = ""

	/* Convert big int to string */
	for x.Cmp(zero) > 0 {
		/* x, r = (x / 58, x % 58) */
		x.QuoRem(x, m, r)
		/* Prepend ASCII character */
		s = string(BITCOIN_BASE58_TABLE[r.Int64()]) + s
	}

	return s
}

// b58checkencode encodes version ver and byte slice b into a base-58 check encoded string.
func b58checkencode(ver uint8, b []byte) (s string) {
	/* Prepend version */
	// fmt.Println("4 - Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)")
	bcpy := append([]byte{ver}, b...)
	// fmt.Println(byteString(bcpy))
	// fmt.Println("=======================")

	/* Create a new SHA256 context */
	sha256H := sha256.New()

	/* SHA256 Hash #1 */
	// fmt.Println("5 - Perform SHA-256 hash on the extended RIPEMD-160 result")
	sha256H.Reset()
	sha256H.Write(bcpy)
	hash1 := sha256H.Sum(nil)
	// fmt.Println(byteString(hash1))
	// fmt.Println("=======================")

	/* SHA256 Hash #2 */
	// fmt.Println("6 - Perform SHA-256 hash on the result of the previous SHA-256 hash")
	sha256H.Reset()
	sha256H.Write(hash1)
	hash2 := sha256H.Sum(nil)
	// fmt.Println(byteString(hash2))
	// fmt.Println("=======================")

	/* Append first four bytes of hash */
	// fmt.Println("7 - Take the first 4 bytes of the second SHA-256 hash. This is the address checksum")
	// fmt.Println(byteString(hash2[0:4]))
	// fmt.Println("=======================")

	// fmt.Println("8 - Add the 4 checksum bytes from stage 7 at the end of extended RIPEMD-160 hash from stage 4. This is the 25-byte binary Bitcoin Address.")
	bcpy = append(bcpy, hash2[0:4]...)
	// fmt.Println(byteString(bcpy))
	// fmt.Println("=======================")

	/* Encode base58 string */
	s = b58encode(bcpy)

	/* For number of leading 0's in bytes, prepend 1 */
	for _, v := range bcpy {
		if v != 0 {
			break
		}
		s = "1" + s
	}
	// fmt.Println("9 - Convert the result from a byte string into a base58 string using Base58Check encoding. This is the most commonly used Bitcoin Address format")
	// fmt.Println(s)
	// fmt.Println("=======================")

	return s
}
