package set02

import (
	"crypto/aes"
	"encoding/base64"
	"io/ioutil"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/pkg/errors"
)

// DecryptCBC decrypts the file content with AES in CBC mode using the given key and IV
func DecryptCBC(filepath string, key, iv []byte) ([]byte, error) {
	inB64, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read: %s", filepath)
	}

	b64 := base64.StdEncoding
	in := make([]byte, b64.DecodedLen(len(inB64)))
	n, err := b64.Decode(in, inB64)
	if err != nil {
		return nil, errors.Wrapf(err, "could not decode: %s", inB64)
	}
	in = in[:n]

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create cipher from key: %s", key)
	}

	decrypter := crypts.NewCBCDecrypter(c, iv)
	data := make([]byte, n)
	copy(data, in)
	decrypter.CryptBlocks(data, data)

	return data, nil
}
