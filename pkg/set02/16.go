package set02

import (
	"bytes"
	"crypto/cipher"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
)

// CBCOracle appends and prepends data to given buffers and encrypts it with CBC under random key
func CBCOracle(src []byte, c cipher.Block, iv []byte) []byte {
	prefixBytes := []byte("comment1=cooking%20MCs;userdata=")
	suffixBytes := []byte(";comment2=%20like%20a%20pound%20of%20bacon")

	src = bytes.Map(
		func(r rune) rune {
			if r == ';' || r == '=' {
				return -1
			}
			return r
		},
		src,
	)

	data := make([]byte, len(src)+len(prefixBytes)+len(suffixBytes))
	copy(data[:len(prefixBytes)], prefixBytes)
	copy(data[len(prefixBytes):len(prefixBytes)+len(src)], src)
	copy(data[len(prefixBytes)+len(src):], suffixBytes)

	data = utils.PKCS7Pad(data, len(data)+c.BlockSize()-(len(data)%c.BlockSize()))

	encrypter := crypts.NewCBCEncrypter(c, iv)
	encrypter.CryptBlocks(data, data)
	return data
}

// CBCCheckAdmin decrypts the given buffer and check for the required sub-buffer
func CBCCheckAdmin(src []byte, c cipher.Block, iv []byte) bool {
	data := make([]byte, len(src))
	copy(data, src)

	decrypter := crypts.NewCBCDecrypter(c, iv)
	decrypter.CryptBlocks(data, data)

	adIden := []byte(";admin=true;")
	return bytes.Contains(data, adIden)
}

// CBCBitFlipping exploits CBC by changing to CBC cipher text bits to obtain desired plaintext
func CBCBitFlipping(c cipher.Block, iv []byte) []byte {
	// expected sub buffer "AAAA;admin=true;"
	// ';' and '=' are removed => replace with '?'j
	src := []byte("\x00admin\x00true\x00")

	enc := CBCOracle(src, c, iv)
	targetBlock := enc[16:32] // second block

	// 0x00 ^ ';' = 0x3b
	// 0x00 ^ '=' = 0x3d
	// xor the cipher block at specific position to obtain the wanted plain text
	targetBlock[0] ^= 0x3b
	targetBlock[6] ^= 0x3d
	targetBlock[11] ^= 0x3b

	copy(enc[16:32], targetBlock)
	return enc
}
