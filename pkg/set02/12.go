package set02

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

// ECBOracle appends source with unknown bytes and encrypts it with ECB under an unknown key
func ECBOracle(src []byte) ([]byte, error) {
	suffixB64 := []byte("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg\naGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq\ndXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg\nYnkK")
	key := []byte("YELLOW SUBMARINE")

	b64 := base64.StdEncoding
	suffix := make([]byte, b64.DecodedLen(len(suffixB64)))
	suffixLen, err := b64.Decode(suffix, suffixB64)
	if err != nil {
		return nil, errors.Wrapf(err, "could not decode: %s", suffixB64)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create cipher from key: %x", key)
	}
	blocksize := c.BlockSize()

	srcLen := len(src)
	padLen := blocksize - ((srcLen + suffixLen) % blocksize)

	dst := make([]byte, srcLen+suffixLen+padLen)
	copy(dst[:srcLen], src)
	copy(dst[srcLen:srcLen+suffixLen], suffix)
	copy(dst[srcLen+suffixLen:], utils.PKCS7Pad([]byte{}, padLen))

	encrypter := crypts.NewECBEncrypter(c)
	encrypter.CryptBlocks(dst, dst)
	return dst, nil
}

// BreakECBOracle detects if the cipher text is encrypted with ECB and find the plaintext
func BreakECBOracle() ([]byte, error) {
	var textsize int
	// find cipher block size and appended text length
	for i := 1; i <= 32; i++ {
		src := bytes.Repeat([]byte("A"), i*2)
		out, err := ECBOracle(src)
		if err != nil {
			return nil, errors.Wrapf(err, "could not encrypt: %s", src)
		}

		if DetectOracle(out) {
			// output always full blocks
			textsize = len(out) - i*2
		}
	}

	var discovered []byte
	for i := 1; i <= textsize; i++ {
		src := bytes.Repeat([]byte("A"), textsize-i)
		tmp := append(src, discovered...)
		cryptDict := make(map[string]byte)
		// map block output to byte
		for b := 0x00; b <= 0xff; b++ {
			entry := append(tmp, byte(b))
			out, err := ECBOracle(entry)
			if err != nil {
				return nil, errors.Wrapf(err, "could not encrypt: %s", entry)
			}

			cryptDict[string(out[:textsize])] = byte(b)
		}

		out, err := ECBOracle(src)
		if err != nil {
			return nil, errors.Wrapf(err, "could not encrypt: %s", src)
		}

		discovered = append(discovered, cryptDict[string(out[:textsize])])
	}

	return discovered, nil
}
