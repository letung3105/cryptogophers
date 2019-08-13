package set02

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"math/rand"

	"github.com/letung3105/cryptogophers/pkg/crypts"
	"github.com/letung3105/cryptogophers/pkg/utils"
	"github.com/pkg/errors"
)

// ComplexECBOracle appends source with unknown bytes and encrypts it with ECB under an unknown key
func ComplexECBOracle(src []byte) ([]byte, error) {
	// seed to ensure same random output
	rand.Seed(2<<31 - 1)
	key := []byte("YELLOW SUBMARINE")

	b64 := base64.StdEncoding
	suffixB64 := []byte("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg\naGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq\ndXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg\nYnkK")
	suffix := make([]byte, b64.DecodedLen(len(suffixB64)))
	if _, err := b64.Decode(suffix, suffixB64); err != nil {
		panic(errors.Wrapf(err, "could not decode: %s", suffixB64))
	}

	prefix := make([]byte, rand.Intn(len(suffix)/2)+len(suffix)/2)
	if _, err := rand.Read(prefix); err != nil {
		panic(errors.Wrap(err, "could create random prefix"))
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create cipher from key: %x", key)
	}
	blocksize := c.BlockSize()

	data := make([]byte, len(prefix)+len(src)+len(suffix))
	copy(data[:len(prefix)], prefix)
	copy(data[len(prefix):len(prefix)+len(src)], src)
	copy(data[len(prefix)+len(src):], suffix)

	paddedData := utils.PKCS7Pad(data, len(data)+blocksize-len(data)%blocksize)
	encrypter := crypts.NewECBEncrypter(c)
	encrypter.CryptBlocks(paddedData, paddedData)
	return paddedData, nil
}

// BreakComplexECBOracle find the plaintext that was encrypt with ComplexECBOracle
func BreakComplexECBOracle() ([]byte, error) {
	firstSample, err := ComplexECBOracle([]byte{})
	if err != nil {
		return nil, errors.Wrapf(err, "could not make call to oracle")
	}

	secondSample, err := ComplexECBOracle([]byte("A"))
	if err != nil {
		return nil, errors.Wrapf(err, "could not make call to oracle")
	}

	// check for first different byte
	var diffIndex int
	for diffIndex = 0; diffIndex < minInt(len(firstSample), len(secondSample)); diffIndex++ {
		if firstSample[diffIndex] != secondSample[diffIndex] {
			break
		}
	}

	// input offset to block align the prefix
	blocksize := 16 // assume known blocksize
	var offset int
	for i := 1; i <= 48; i++ {
		sample, err := ComplexECBOracle(bytes.Repeat([]byte("A"), i))
		if err != nil {
			return nil, errors.Wrapf(err, "could not make call to oracle")
		}

		if utils.HasNonOverlapDup(sample, blocksize) {
			offset = i % blocksize
			break
		}
	}

	// find length of the prefix with offset and length of appended text
	offsetIn := bytes.Repeat([]byte("A"), offset)
	offsetOut, err := ComplexECBOracle(offsetIn)
	if err != nil {
		return nil, errors.Wrapf(err, "could not make call to oracle with input: %x", offsetIn)
	}

	var prefixLen int
	if offset != 0 {
		prefixLen = diffIndex + blocksize
	} else {
		prefixLen = diffIndex
	}
	textLen := len(offsetOut) - prefixLen

	var discovered []byte
	for i := 1; i <= textLen; i++ {
		src := bytes.Repeat([]byte("A"), textLen-i+offset)
		tmp := append(src, discovered...)
		cryptDict := make(map[string]byte)

		// map block output to byte
		for b := 0x00; b <= 0xff; b++ {
			entry := append(tmp, byte(b))
			out, err := ComplexECBOracle(entry)
			if err != nil {
				return nil, errors.Wrapf(err, "could not make call to oracle with input: %x", entry)
			}

			cryptDict[string(out[prefixLen:prefixLen+textLen])] = byte(b)
		}

		out, err := ComplexECBOracle(src)
		if err != nil {
			return nil, errors.Wrapf(err, "could not make call to oracle with input: %x", src)
		}
		discovered = append(discovered, cryptDict[string(out[prefixLen:prefixLen+textLen])])
	}

	return discovered, nil
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
