package set02

import (
	"crypto/aes"
	"math/rand"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

// EncryptOracle encrypts the input in either ECB or CBC randomly using random key
func EncryptOracle(src []byte, keylen int) ([]byte, bool, error) {
	randKey := make([]byte, keylen)
	if _, err := rand.Read(randKey); err != nil {
		return nil, false, errors.Wrapf(err, "could not create random key of size %d", keylen)
	}

	suffixLen := rand.Intn(6) + 5
	suffix := make([]byte, suffixLen)
	if _, err := rand.Read(suffix); err != nil {
		return nil, false, errors.Wrapf(err, "could not create random suffix of size %d", suffixLen)
	}

	prefixLen := rand.Intn(6) + 5
	prefix := make([]byte, prefixLen)
	if _, err := rand.Read(prefix); err != nil {
		return nil, false, errors.Wrapf(err, "could not create random prefix ofr size %d", prefixLen)
	}

	c, err := aes.NewCipher(randKey)
	if err != nil {
		return nil, false, errors.Wrapf(err, "could not create cipher from key %x", randKey)
	}
	blocksize := c.BlockSize()
	padLen := blocksize - ((prefixLen + len(src) + suffixLen) % blocksize)

	paddedSrc := make([]byte, prefixLen+len(src)+suffixLen+padLen)
	copy(paddedSrc[:prefixLen], prefix)
	copy(paddedSrc[prefixLen:prefixLen+len(src)], src)
	copy(paddedSrc[prefixLen+len(src):prefixLen+len(src)+suffixLen], suffix)
	copy(paddedSrc[prefixLen+len(src)+suffixLen:], utils.PKCS7Pad([]byte{}, padLen))

	var isECB bool
	dst := make([]byte, len(paddedSrc))
	copy(dst, paddedSrc)
	if rand.Intn(2) == 0 {
		encrypter := crypts.NewECBEncrypter(c)
		encrypter.CryptBlocks(dst, dst)
		isECB = true
	} else {
		iv := make([]byte, c.BlockSize())
		encrypter := crypts.NewCBCEncrypter(c, iv)
		encrypter.CryptBlocks(dst, dst)
		isECB = false
	}
	return dst, isECB, nil
}

// DetectOracle checks whether the cipher text is encrypted with ECB or CBC
func DetectOracle(src []byte) bool {
	for i := 16; i <= 32; i += 8 {
		if utils.HasNonOverlapDup(src, i) {
			return true
		}
	}
	return false
}
