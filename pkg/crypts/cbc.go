// Cipher block chaining (ECB) mode

// Extend the implemetation of ECB to CBC

package crypts

import (
	"crypto/cipher"
	"fmt"
)

type cbc struct {
	b         cipher.Block
	blockSize int
	iv        []byte
}

func newCBC(b cipher.Block, iv []byte) *cbc {
	return &cbc{
		b:         b,
		blockSize: b.BlockSize(),
		iv:        iv,
	}
}

type cbcEncrypter cbc

// NewCBCEncrypter returns a BlockMode which encrypts in cipher block chaining mode
func NewCBCEncrypter(b cipher.Block, iv []byte) cipher.BlockMode {
	if len(iv) != b.BlockSize() {
		panic(fmt.Sprintf(
			"invalid IV length:\nhave %d\nwant %d", len(iv), b.BlockSize(),
		))
	}

	return (*cbcEncrypter)(newCBC(b, iv))
}

func (x *cbcEncrypter) BlockSize() int { return x.blockSize }

func (x *cbcEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("input not full blocks")
	}
	if len(dst) < len(src) {
		panic("output smaller than input")
	}

	iv := x.iv
	for len(src) > 0 {
		tmp, _ := FixedXOR(src[:x.blockSize], iv)
		x.b.Encrypt(dst[:x.blockSize], tmp)

		iv = dst[:x.blockSize]
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type cbcDecrypter cbc

// NewCBCDecrypter returns a BlockMode which decrypts in cipher block chaining mode
func NewCBCDecrypter(b cipher.Block, iv []byte) cipher.BlockMode {
	if len(iv) != b.BlockSize() {
		panic(fmt.Sprintf(
			"invalid IV length:\nhave %d\nwant %d", len(iv), b.BlockSize(),
		))
	}

	return (*cbcDecrypter)(newCBC(b, iv))
}

func (x *cbcDecrypter) BlockSize() int { return x.blockSize }

func (x *cbcDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("input not full blocks")
	}
	if len(dst) < len(src) {
		panic("output smaller than input")
	}

	iv := make([]byte, len(src))
	copy(iv, x.iv)
	copy(iv[x.blockSize:], src[:len(src)-x.blockSize])
	for len(src) > 0 {
		x.b.Decrypt(dst[:x.blockSize], src[:x.blockSize])
		tmp, _ := FixedXOR(dst[:x.blockSize], iv[:x.blockSize])
		copy(dst[:x.blockSize], tmp)

		iv = iv[x.blockSize:]
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
