package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ripemd160"
	"math/big"
	"myLibrary/go-library/go/crypt"
)

func byteString(b []byte) (s string) {
	s = ""
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%02X", b[i])
	}
	return s
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

// paddedAppend appends the src byte slice to dst, returning the new slice.
// If the length of the source is smaller than the passed size, leading zero
// bytes are appended to the dst slice before appending src.
func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}

const version = byte(0x00)
const addressChecksumLen = 4

// Wallet stores private and public keys
type Wallet struct {
	// PrvKey          *ecdsa.PrivateKey
	PrivateKeyBytes []byte
	PublicKeyBytes  []byte
}

// NewWallet creates and returns a Wallet
func NewWallet() (*Wallet, error) {
	// prvK, err := newKeyPair()
	// if nil != err {
	// 	return nil, err
	// }
	var wallet Wallet
	prvBytes, pubBytes := newKeyPairBytes()
	wallet.PrivateKeyBytes = prvBytes
	wallet.PublicKeyBytes = pubBytes
	return &wallet, nil
}

// func (w *Wallet) GetPrvBytes() []byte {
// 	d := w.PrvKey.D.Bytes()
// 	b := make([]byte, 0, privKeyBytesLen)
// 	priKet := paddedAppend(privKeyBytesLen, b, d)
// 	return priKet
// }
// func (w *Wallet) GetPubBytes() []byte {
// 	return append(w.PrvKey.PublicKey.X.Bytes(), w.PrvKey.PublicKey.Y.Bytes()...)
// }

// GetAddress returns wallet address
func (w Wallet) GetAddress() (address string) {
	/* See https://en.bitcoin.it/wiki/Technical_background_of_Bitcoin_addresses */

	/* Convert the public key to bytes */
	pub_bytes := w.PublicKeyBytes

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
	address = b58checkencode(0x00, pub_hash_2)

	return address
}


// HashPubKey hashes public key
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

const privKeyBytesLen = 32

func newKeyPair() (*ecdsa.PrivateKey, error) {
	curve := crypto.S256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return private, nil
}

// func newKeyPairBytes(private *ecdsa.PrivateKey) ([]byte, []byte) {
// 	d := private.D.Bytes()
// 	b := make([]byte, 0, privKeyBytesLen)
// 	priKet := paddedAppend(privKeyBytesLen, b, d)
// 	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
//
// 	return priKet, pubKey
// }

func newKeyPairBytes() ([]byte, []byte) {
	// curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	// private, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	// b := make([]byte, 0, privKeyBytesLen)
	// priKet := paddedAppend(privKeyBytesLen, b, d)
	priKet := encrypt.ECDSAPrv2Bytes(private)
	pubKey := encrypt.ECDSAPub2Bytes(&private.PublicKey)

	return priKet, pubKey
}

// ToWIF converts a Bitcoin private key to a Wallet Import Format string.
func ToWIF(priv []byte) (wif string) {
	/* Convert bytes to base-58 check encoded string with version 0x80 */
	wif = b58checkencode(0x80, priv)

	return wif
}



