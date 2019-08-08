package set01

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/pkg/errors"
)

// Challenge07 decrypts aes in ecb mode encryped text
func Challenge07(filepath string, key []byte) ([]byte, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"Could not read file: %s", filepath,
		))
	}

	b64 := base64.StdEncoding
	cipher := make([]byte, b64.DecodedLen(len(b)))
	n, err := b64.Decode(cipher, b)
	if err != nil {
		return nil, errors.Wrap(err, "Could not decode base64 text")
	}
	cipher = cipher[:n]

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(
			"Coulld not create new cipher from key: %s", key,
		))
	}

	decrypter := crypts.NewECBDecrypter(c)
	data := make([]byte, len(cipher))
	decrypter.CryptBlocks(data, cipher)
	return data, nil
}
