package set01

import (
	"crypto/aes"
	"encoding/base64"
	"io/ioutil"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/pkg/errors"
)

// ECBDecryptB64AES decrypts aes in ecb mode encryped text
func ECBDecryptB64AES(filepath string, key []byte) ([]byte, error) {
	srcB64, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read: %s", filepath)
	}

	b64 := base64.StdEncoding
	src := make([]byte, b64.DecodedLen(len(srcB64)))
	n, err := b64.Decode(src, srcB64)
	if err != nil {
		return nil, errors.Wrapf(err, "could not decode: %s", srcB64)
	}
	src = src[:n]

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create cipher from key: %s", key)
	}

	decrypter := crypts.NewECBDecrypter(c)
	dst := make([]byte, len(src))
	decrypter.CryptBlocks(dst, src)
	return dst, nil
}
