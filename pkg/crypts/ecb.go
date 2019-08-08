// Electronic code block (ECB) mode

// Implementaion based on Golang officials cipher package
// https://golang.org/src/crypto/cipher/

package crypts

import "crypto/cipher"

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ebcEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code block mode
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ebcEncrypter)(newECB(b))
}

func (x *ebcEncrypter) BlockSize() int {
	return x.blockSize
}

func (x *ebcEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("input not full blocks")
	}
	if len(dst) < len(src) {
		panic("output smaller than input")
	}

	for len(src) > 0 {
		x.b.Encrypt(dst[:x.blockSize], src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ebcDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code block mode
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ebcDecrypter)(newECB(b))
}

func (x *ebcDecrypter) BlockSize() int {
	return x.blockSize
}

func (x *ebcDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("input not full blocks")
	}
	if len(dst) < len(src) {
		panic("output smaller than input")
	}

	for len(src) > 0 {
		x.b.Decrypt(dst[:x.blockSize], src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
