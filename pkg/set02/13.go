package set02

import (
	"bytes"
	"crypto/aes"
	"fmt"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

// KVParseDecrypt takes in an ECB encrypted key-value encoded string into a map
func KVParseDecrypt(src []byte, key []byte) (map[string][]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create cipher from key: %s", key)
	}

	data := make([]byte, len(src))
	copy(data, src)
	decrypter := crypts.NewECBDecrypter(c)
	decrypter.CryptBlocks(data, data)
	data = utils.PKCS7Unpad(data)

	pairs := bytes.Split(data, []byte("&"))
	kv := make(map[string][]byte)
	for _, pair := range pairs {
		p := bytes.Split(pair, []byte("="))
		kv[string(p[0])] = p[1]
	}
	return kv, nil
}

// ProfileForEncrypt create a new key-value encoded string profile for the given email
// and encrypts it with ECB
func ProfileForEncrypt(email []byte, key []byte) ([]byte, error) {
	// remove '&' and '=' from email
	filter := func(r rune) rune {
		if r == '&' || r == '=' {
			return -1
		}
		return r
	}

	email = bytes.Map(filter, email)
	prof := []byte(fmt.Sprintf("email=%s&uid=%d&role=%s", email, 10, "user"))

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create cipher from key: %s", key)
	}
	blocksize := c.BlockSize()

	paddedLen := len(prof) + blocksize - (len(prof) % blocksize)
	data := utils.PKCS7Pad(prof, paddedLen)
	encrypter := crypts.NewECBEncrypter(c)
	encrypter.CryptBlocks(data, data)

	return data, nil
}

// FindBlockSize of the encryption used by the oracle
func FindBlockSize(key []byte, oracle func([]byte, []byte) ([]byte, error)) (int, error) {
	for i := 1; i <= 64; i++ {
		src := bytes.Repeat([]byte("A"), i)
		out, err := oracle(src, key)
		if err != nil {
			return 0, errors.Wrapf(err, "could not call oracle(%s, %s)", src, key)
		}

		// only find block size divided by 8
		for j := 16; j <= 64; j += 8 {
			if utils.HasNonOverlapDup(out, j) {
				return j, nil
			}
		}
	}
	return 0, errors.New("could not find block size")
}

// ModProfileRole takes in ECB encrypted kv-string encoded profile and changes the role
// using only the email input and output cipher
func ModProfileRole(key []byte) ([]byte, error) {
	blocksize, err := FindBlockSize(key, ProfileForEncrypt)
	if err != nil {
		return nil, errors.Wrapf(err, "could not finc blocksize for key: %x", key)
	}

	// mandatory bytes for profile kv-string
	fixedBytes := []byte("email=&uid=10&role=")
	fixedFields := bytes.SplitAfterN(fixedBytes, []byte("="), 2)

	// first block = "email=AAAAAAAAAA",
	// second block = "admin\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b",
	emailOffset := bytes.Repeat([]byte("A"), blocksize-len(fixedFields[0]))
	adminFullBlock := utils.PKCS7Pad([]byte("admin"), blocksize)

	// email = "AAAAAAAAAAadmin\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b",
	oracleEmail := make([]byte, len(emailOffset)+len(adminFullBlock))
	copy(oracleEmail[:len(emailOffset)], emailOffset)
	copy(oracleEmail[len(emailOffset):], adminFullBlock)

	// get cipher block with only "admin" and padded bytes
	oracleCipher, err := ProfileForEncrypt(oracleEmail, key)
	if err != nil {
		panic(errors.Wrapf(err, "could not encrypt: %s", oracleEmail))
	}
	adminCipherBlock := oracleCipher[blocksize : 2*blocksize]

	// first block = "email=AAAAAAAAAA",
	// second block = "AAA&uid=10&role=",
	// third clock = "user\x0c\x0c\x0c\x0c\x0c\x0c\x0c\x0c\x0c\x0c\x0c",
	adminEmail := bytes.Repeat([]byte("A"), 2*blocksize-len(fixedBytes))
	adminCipher, err := ProfileForEncrypt(adminEmail, key)
	if err != nil {
		panic(errors.Wrapf(err, "could not encrypt: %s", oracleEmail))
	}

	// replace third block with previously obtained block with "admin"
	modAdminCipher := make([]byte, 3*blocksize)
	copy(modAdminCipher[:2*blocksize], adminCipher[:2*blocksize])
	copy(modAdminCipher[2*blocksize:], adminCipherBlock)

	return modAdminCipher, nil
}
